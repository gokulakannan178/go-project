package shared

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
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

func (s *Shared) BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func (s *Shared) EndOfMonth(t time.Time) time.Time {
	return s.BeginningOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}
func (s *Shared) RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
