package json

import (
	"fmt"
	"os"
)

func Example_parse2() {
	r, err := os.Open("./data/context-type.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close()

	root, err := Parse(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	root.Walk(func(node *Node) error {
		parent := node.Parent
		if parent != nil {
			if parent.Type == ArrayNode {
				if node.Index > 0 {
					return SkipNode
				}
				if node.Type != ObjectNode {
					return SkipNode
				}
				return nil
			}
		}
		fmt.Println(node)
		return nil
	})
	// Output:
	// path: "$", type: "object"
	// path: "$.snssai", type: "object"
	// path: "$.snssai.type", type: "string", key: "type", value: "Property"
	// path: "$.snssai.values", type: "array"
	// path: "$.dnn", type: "object"
	// path: "$.dnn.type", type: "string", key: "type", value: "Property"
	// path: "$.dnn.values", type: "array"
	// path: "$.end_time", type: "object"
	// path: "$.end_time.type", type: "string", key: "type", value: "Property"
	// path: "$.end_time.values", type: "array"
	// path: "$.jsonString", type: "string", key: "jsonString", value: "{
	//   "snssai": {
	//     "type": "Property",
	//     "values": [
	//       [
	//         "1:223344",
	//         "2022-12" ...
	// path: "$.session_id", type: "object"
	// path: "$.session_id.type", type: "string", key: "type", value: "Property"
	// path: "$.session_id.values", type: "array"
	// path: "$.imsi", type: "object"
	// path: "$.imsi.type", type: "string", key: "type", value: "Property"
	// path: "$.imsi.values", type: "array"
	// path: "$.file_datetime", type: "object"
	// path: "$.file_datetime.type", type: "string", key: "type", value: "Property"
	// path: "$.file_datetime.values", type: "array"
	// path: "$.type", type: "string", key: "type", value: "vCoreTest"
	// path: "$.dl_mesurement", type: "object"
	// path: "$.dl_mesurement.type", type: "string", key: "type", value: "Property"
	// path: "$.dl_mesurement.values", type: "array"
	// path: "$.class_no", type: "object"
	// path: "$.class_no.type", type: "string", key: "type", value: "Property"
	// path: "$.class_no.values", type: "array"
	// path: "$.group_no", type: "object"
	// path: "$.group_no.type", type: "string", key: "type", value: "Property"
	// path: "$.group_no.values", type: "array"
	// path: "$.time_of_last_packet", type: "object"
	// path: "$.time_of_last_packet.type", type: "string", key: "type", value: "Property"
	// path: "$.time_of_last_packet.values", type: "array"
	// path: "$.ul_mesurement", type: "object"
	// path: "$.ul_mesurement.type", type: "string", key: "type", value: "Property"
	// path: "$.ul_mesurement.values", type: "array"
	// path: "$.start_time", type: "object"
	// path: "$.start_time.type", type: "string", key: "type", value: "Property"
	// path: "$.start_time.values", type: "array"
	// path: "$.time_of_first_packet", type: "object"
	// path: "$.time_of_first_packet.type", type: "string", key: "type", value: "Property"
	// path: "$.time_of_first_packet.values", type: "array"
	// path: "$.mdn", type: "object"
	// path: "$.mdn.type", type: "string", key: "type", value: "Property"
	// path: "$.mdn.values", type: "array"
	// path: "$.ue_ip", type: "object"
	// path: "$.ue_ip.type", type: "string", key: "type", value: "Property"
	// path: "$.ue_ip.values", type: "array"
	// path: "$.temporalDomain", type: "object"
	// path: "$.temporalDomain.name", type: "null"
	// path: "$.temporalDomain.id", type: "string", key: "id", value: "urn:ngsi-ld:vCoreTest:1111"
	// path: "$.temporalDomain.type", type: "null"
	// path: "$.temporalDomain.modifiedTime", type: "null"
	// path: "$.temporalDomain.createdTime", type: "null"
	// path: "$.temporalDomain.value", type: "null"
	// path: "$.temporalDomain.lastN", type: "number", key: "lastN", value: 10
	// path: "$.temporalDomain.updateCount", type: "null"
	// path: "$.temporalDomain.q", type: "null"
	// path: "$.temporalDomain.limit", type: "number", key: "limit", value: 0
	// path: "$.temporalDomain.offest", type: "number", key: "offest", value: 0
	// path: "$.temporalDomain.options", type: "null"
	// path: "$.temporalDomain.count", type: "null"
	// path: "$.temporalDomain.local", type: "null"
	// path: "$.temporalDomain.startDate", type: "null"
	// path: "$.temporalDomain.startHour", type: "null"
	// path: "$.temporalDomain.startMin", type: "null"
	// path: "$.temporalDomain.startTime", type: "null"
	// path: "$.temporalDomain.endDate", type: "null"
	// path: "$.temporalDomain.endHour", type: "null"
	// path: "$.temporalDomain.endMin", type: "null"
	// path: "$.temporalDomain.endTime", type: "null"
	// path: "$.temporalDomain.userLevelId", type: "null"
	// path: "$.temporalDomain.levelId", type: "null"
	// path: "$.temporalDomain.searchType", type: "null"
	// path: "$.temporalDomain.searchDateType", type: "null"
	// path: "$.temporalDomain.searchText", type: "null"
	// path: "$.temporalDomain.searchStatus", type: "null"
	// path: "$.temporalDomain.userId", type: "null"
	// path: "$.temporalDomain.userName", type: "null"
	// path: "$.temporalDomain.search_type", type: "null"
	// path: "$.temporalDomain.search_keyword", type: "null"
	// path: "$.temporalDomain.orderBy", type: "null"
	// path: "$.temporalDomain.titleName", type: "null"
	// path: "$.temporalDomain.group_no", type: "null"
	// path: "$.temporalDomain.page_type", type: "null"
	// path: "$.temporalDomain.companyId", type: "null"
	// path: "$.temporalDomain.loginId", type: "null"
	// path: "$.id", type: "string", key: "id", value: "urn:ngsi-ld:vCoreTest:1111"
	// path: "$.temporalMap", type: "object"
	// path: "$.temporalMap.lastN", type: "number", key: "lastN", value: 10
	// path: "$.temporalMap.id", type: "string", key: "id", value: "urn:ngsi-ld:vCoreTest:1111"
	// path: "$.temporalMap.type", type: "string", key: "type", value: "vCoreTest"
	// path: "$.location_info", type: "object"
	// path: "$.location_info.type", type: "string", key: "type", value: "Property"
	// path: "$.location_info.values", type: "array"
}
