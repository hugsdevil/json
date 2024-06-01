package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func Parse(r io.Reader) (*Node, error) {
	p := newParser(r)
	node, err := p.parse()
	if err != nil {
		return nil, err
	}
	return node, nil
}

type (
	parser struct {
		tokenizer *json.Decoder
	}
)

func newParser(r io.Reader) *parser {
	return &parser{
		tokenizer: json.NewDecoder(r),
	}
}

func (p *parser) parse() (*Node, error) {
	root := &Node{
		Index: 0,
		Path:  "$",
		Depth: 0,
	}
	node, err := parseValue(p.tokenizer, root)
	if err != nil {
		return nil, err
	}
	return node, nil
}

//lint:ignore ST1012 end of bracket
var eob = errors.New("end of bracket")

func parseKey(decoder *json.Decoder) (key string, err error) {
	token, err := decoder.Token()
	if err != nil {
		return "", err
	}

	if token == json.Delim('}') {
		return "", eob
	}

	key, ok := token.(string)
	if !ok {
		return "", fmt.Errorf("invalid key string: %s", token)
	}
	return key, nil
}

func parseValue(dec *json.Decoder, parent *Node) (*Node, error) {
	token, err := dec.Token()
	if err != nil {
		return nil, err
	}

	switch token {
	case json.Delim('{'):
		parent.Type = tokenToNodeType(token)
		var prev *Node
		for i := 0; ; i++ {
			key, err := parseKey(dec)
			if err != nil {
				if err == eob {
					break
				}
				return nil, err
			}

			node := &Node{
				Parent: parent,
				Path:   parent.Path + "." + key,
				Key:    key,
				Index:  i,
				Depth:  parent.Depth + 1,
			}

			curr, err := parseValue(dec, node)
			if err != nil {
				if err == eob {
					break
				}
				return nil, err
			}

			if i == 0 {
				parent.FirstChild = curr
			} else {
				prev.NextSibling = curr
				curr.PrevSibling = prev
			}
			prev = curr
		}
		parent.LastChild = prev
		return parent, nil
	case json.Delim('}'):
		return nil, eob
	case json.Delim('['):
		parent.Type = tokenToNodeType(token)
		var prev *Node
		for i := 0; ; i++ {
			node := &Node{
				Parent: parent,
				Path:   fmt.Sprintf("%s[%d]", parent.Path, i),
				Index:  i,
				Depth:  parent.Depth + 1,
			}

			curr, err := parseValue(dec, node)
			if err != nil {
				if err == eob {
					break
				}
				return nil, err
			}

			if i == 0 {
				parent.FirstChild = curr
			} else {
				prev.NextSibling = curr
				curr.PrevSibling = prev
			}
			prev = curr
		}
		parent.LastChild = prev
		return parent, nil
	case json.Delim(']'):
		return nil, eob
	default:
		parent.Value = token
		parent.Type = tokenToNodeType(token)
		return parent, nil
	}
}
