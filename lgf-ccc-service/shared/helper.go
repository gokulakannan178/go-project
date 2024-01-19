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

func (s *Shared) StartDayOfWeek(t time.Time) time.Time { //get monday 00:00:00
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := t.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay.Add(1 * (weekday - 2) * 24 * time.Hour)
}

func (s *Shared) EndDayOfWeek(t time.Time) time.Time {
	return s.StartDayOfWeek(t).AddDate(0, 0, 7).Add(-time.Second)
}
