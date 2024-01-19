package shared

import (
	"encoding/json"
	"fmt"
)

//BsonToJSONPrint : ""
func (s *Shared) BsonToJSONPrint(d interface{}) {
	b, err1 := json.Marshal(d)
	fmt.Println("err1", err1, string(b))
}

//BsonToJSONPrintV2 : ""
func (s *Shared) BsonToJSONPrintTag(tag string, d interface{}) {
	b, err1 := json.Marshal(d)
	fmt.Println("err1==>", err1, tag, "==>", string(b))
}
