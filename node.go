package json

import (
	"encoding/json"
	"errors"
	"fmt"
)

type NodeType uint32

const (
	ErrorNode NodeType = iota
	ObjectNode
	ArrayNode
	StringNode
	NumberNode
	BoolNode
	NullNode
)

func (typ NodeType) String() string {
	switch typ {
	case ObjectNode:
		return "object"
	case ArrayNode:
		return "array"
	case StringNode:
		return "string"
	case NumberNode:
		return "number"
	case BoolNode:
		return "bool"
	case NullNode:
		return "null"
	default:
		return "error"
	}
}

func tokenToNodeType(token json.Token) NodeType {
	switch token.(type) {
	case json.Delim:
		if token == json.Delim('{') {
			return ObjectNode
		} else if token == json.Delim('[') {
			return ArrayNode
		}
		return ErrorNode
	case string:
		return StringNode
	case bool:
		return BoolNode
	case float64:
		return NumberNode
	case json.Number:
		return NumberNode
	case nil:
		return NullNode
	default:
		return ErrorNode
	}
}

type (
	Node struct {
		Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

		Path   string
		Key    string
		Value  interface{}
		Type   NodeType
		Length int
		Index  int
		Depth  int
	}
)

func (node *Node) Root() *Node {
	root := node
	for root.Parent != nil {
		root = root.Parent
	}
	return root
}

func (node *Node) SetParent(parent *Node) {
	if node.Parent != nil {
		node.Parent.SetLastChild(parent)
	}
	parent.SetLastChild(node)
}

func (node *Node) SetFirstChild(first *Node) {
	if node.FirstChild != nil {
		node.FirstChild.PrevSibling = first
		first.NextSibling = node.FirstChild
	}
	node.FirstChild = first
	first.Parent = node
}

func (node *Node) SetLastChild(last *Node) {
	if node.LastChild != nil {
		node.LastChild.NextSibling = last
		last.PrevSibling = node.LastChild
	}
	node.LastChild = last
	last.Parent = node
}

func (node *Node) SetPrevSibling(prev *Node) {
	if node.PrevSibling != nil {
		prev.PrevSibling = node.PrevSibling
		node.PrevSibling.NextSibling = prev
	}
	node.PrevSibling = prev
	prev.NextSibling = node
}

func (node *Node) SetNextSibling(next *Node) {
	if node.NextSibling != nil {
		next.NextSibling = node.NextSibling
		node.NextSibling.PrevSibling = next
	}
	node.NextSibling = next
	next.PrevSibling = node
}

func (node *Node) String() string {
	s := fmt.Sprintf(`path: "%s", type: "%s"`, node.Path, node.Type)
	if node.Value != nil {
		var v interface{}
		switch node.Type {
		case StringNode:
			if vv, ok := node.Value.(string); ok {
				if len(vv) > 100 {
					v = `"` + vv[:100] + `..."`

				} else {
					v = `"` + vv + `"`
				}
			}
		default:
			v = node.Value
		}
		s = s + fmt.Sprintf(`, key: "%s", value: %v`, node.Key, v)
	}
	return s
}

type WalkFunc func(node *Node) error

func SkipArrayNotZero(handler func(*Node)) WalkFunc {
	return func(node *Node) error {
		parent := node.Parent
		if parent != nil {
			if parent.Type == ArrayNode {
				if node.Index > 0 {
					return SkipNode
				}
			}
		}
		handler(node)
		return nil
	}
}

func (node *Node) Walk(fn WalkFunc) error {
	if node == nil {
		return fmt.Errorf("nil root node")
	}
	return walk(node, fn)
}

//lint:ignore ST1012 skip node
var SkipNode = errors.New("skip node")

func walk(node *Node, walkFn WalkFunc) error {
	if node == nil {
		return nil
	}

	if err := walkFn(node); err != nil {
		return err
	}
	if err := walk(node.FirstChild, walkFn); err != nil {
		if err != SkipNode {
			return err
		}
	}

	curr := node.NextSibling
	for curr != nil {
		if err := walkFn(curr); err != nil {
			return err
		}
		if err := walk(curr.FirstChild, walkFn); err != nil {
			if err != SkipNode {
				return err
			}
		}
		curr = curr.NextSibling
	}
	return nil
}
