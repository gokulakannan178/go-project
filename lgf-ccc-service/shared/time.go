package shared

import (
	"errors"
	"fmt"
	"time"
)

func (s *Shared) UniqueDateStr(date *time.Time) (string, error) {
	if date == nil {
		return "", errors.New("Date is nil")
	}
	str := fmt.Sprintf("%v_%v_%v", date.Day(), date.Month().String(), date.Year())
	return str, nil
}

//Take the Date and Give StartDateToEndDate
func (s *Shared) StartDateToEndDate(date *time.Time) (*time.Time, *time.Time) {
	if date == nil {

		return nil, nil
	}

	FromDate := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	ToDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
	return &FromDate, &ToDate
}
