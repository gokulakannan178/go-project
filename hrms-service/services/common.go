package services

import (
	"fmt"
	"hrms-services/models"
	"math"
	"time"
)

func (s *Service) FindWeekStartAndEndDate(ctx *models.Context, StartDate *time.Time) (*time.Time, *time.Time, error) {
	var sd, ed time.Time
	fmt.Println("StartDate.Weekday().String()===>", StartDate.Weekday().String())

	switch StartDate.Weekday().String() {
	case "Sunday":
		sd = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day(), 0, 0, 0, 0, StartDate.Location())
		ed = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()+6, 23, 59, 59, 999999999, StartDate.Location())
		fmt.Println("StartDate====>", sd)
		fmt.Println("Enddate====>", ed)
	case "Monday":
		sd = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()-1, 0, 0, 0, 0, StartDate.Location())
		ed = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()+5, 23, 59, 59, 999999999, StartDate.Location())
		fmt.Println("StartDate====>", sd)
		fmt.Println("Enddate====>", ed)
	case "Tuesday":
		sd = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()-2, 0, 0, 0, 0, StartDate.Location())
		ed = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()+4, 23, 59, 59, 999999999, StartDate.Location())
		fmt.Println("StartDate====>", sd)
		fmt.Println("Enddate====>", ed)
	case "Wednesday":
		sd = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()-3, 0, 0, 0, 0, StartDate.Location())
		ed = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()+3, 23, 59, 59, 999999999, StartDate.Location())
		fmt.Println("StartDate====>", sd)
		fmt.Println("Enddate====>", ed)
	case "Thursday":
		sd = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()-4, 0, 0, 0, 0, StartDate.Location())
		ed = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()+2, 23, 59, 59, 999999999, StartDate.Location())
		fmt.Println("StartDate====>", sd)
		fmt.Println("Enddate====>", ed)
	case "Friday":
		sd = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()-5, 0, 0, 0, 0, StartDate.Location())
		ed = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()+1, 23, 59, 59, 999999999, StartDate.Location())
		fmt.Println("StartDate====>", sd)
		fmt.Println("Enddate====>", ed)
	case "Saturday":
		sd = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day()-6, 0, 0, 0, 0, StartDate.Location())
		ed = time.Date(StartDate.Year(), StartDate.Month(), StartDate.Day(), 23, 59, 59, 999999999, StartDate.Location())
		fmt.Println("StartDate====>", sd)
		fmt.Println("Enddate====>", ed)
	}
	return &sd, &ed, nil
}
func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
