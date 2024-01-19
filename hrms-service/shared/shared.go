package shared

import (
	"fmt"
	"hrms-services/config"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Shared : ""
type Shared struct {
	commandArgs map[string]string
	Config      *config.ViperConfigReader
}

//NewShared : Shared Factory
func NewShared(commandArgs map[string]string, configuration *config.ViperConfigReader) *Shared {
	return &Shared{commandArgs: commandArgs, Config: configuration}
}

//GetCmdArg : ""
func (sh *Shared) GetCmdArg(key string) string {
	return sh.commandArgs[key]
}

//SplitCmdArguments : ""
func SplitCmdArguments(args []string) map[string]string {
	// fmt.Println(args)
	m := make(map[string]string)
	for _, v := range args {
		strs := strings.Split(v, "=")
		if len(strs) == 2 {
			m[strs[0]] = strs[1]
		} else {
			log.Println("not proper arguments", strs)
		}
	}
	// fmt.Print(args, m)
	return m
}

//GetTransactionID : ""
func (sh *Shared) GetTransactionID(charset string, length int) string {
	charset = "ABCDEFFHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//Get : ""
func (sh *Shared) Get(url string, h map[string]string) (resp *http.Response, e error) {
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		return nil, err1
	}
	for key, val := range h {
		req.Header.Add(key, val)
	}
	return client.Do(req)
}

// GetRandomOTP : returns random numeric string
// @param n length
func (sh *Shared) GetRandomOTP(n int) string {
	var x = []byte("0123456789")
	return genRandomStr(n, x)
}

func random(min, max int) int {
	// Source of new Random
	return rand.Intn(max-min) + min
}

func genRandomStr(n int, x []byte) string {
	var t = ""
	for i := 0; i < n; i++ {
		rr := random(0, len(x)-1)
		t += string(x[rr])
	}
	return t
}

func (sh *Shared) GetCurrentWeek(d int) (startWeek, endWeek time.Time) {
	t := time.Now()
	y, w := t.ISOWeek()
	startWeek, endWeek = WeekRange(y, w, d)
	return startWeek, endWeek
}

func WeekRange(year, week int, day int) (start, end time.Time) {
	start = WeekStart(year, week)
	end = start.AddDate(0, 0, day-1)
	return
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func (sh *Shared) MonthDays(beginDate time.Time, endDate time.Time, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool) (days int) {

	days = 0
	mondaycount := 0
	tuesdaycount := 0
	wednesdaycount := 0
	thursdaycount := 0
	fridaycount := 0
	saturdaycount := 0
	sundaycount := 0

	// monday = true
	// tuesday = true
	// wednesday = true
	// thursday = true
	// friday = true
	// saturday = false
	// sunday = false
	for {

		if beginDate.Equal(endDate) {
			break
		}
		if monday == true {
			if beginDate.Weekday() != 6 && beginDate.Weekday() != 5 && beginDate.Weekday() != 4 && beginDate.Weekday() != 3 && beginDate.Weekday() != 2 && beginDate.Weekday() != 1 {
				mondaycount++
			}
		} else {
			mondaycount = 0
		}
		if tuesday == true {
			if beginDate.Weekday() != 6 && beginDate.Weekday() != 5 && beginDate.Weekday() != 4 && beginDate.Weekday() != 3 && beginDate.Weekday() != 2 && beginDate.Weekday() != 0 {
				tuesdaycount++
			}
		} else {
			tuesdaycount = 0
		}

		if wednesday == true {
			if beginDate.Weekday() != 6 && beginDate.Weekday() != 5 && beginDate.Weekday() != 4 && beginDate.Weekday() != 3 && beginDate.Weekday() != 1 && beginDate.Weekday() != 0 {
				wednesdaycount++
			}
		} else {

			wednesdaycount = 0
		}

		if thursday == true {
			if beginDate.Weekday() != 6 && beginDate.Weekday() != 5 && beginDate.Weekday() != 4 && beginDate.Weekday() != 2 && beginDate.Weekday() != 1 && beginDate.Weekday() != 0 {
				thursdaycount++
			}
		} else {
			thursdaycount = 0
		}

		if friday == true {
			if beginDate.Weekday() != 6 && beginDate.Weekday() != 5 && beginDate.Weekday() != 3 && beginDate.Weekday() != 2 && beginDate.Weekday() != 1 && beginDate.Weekday() != 0 {
				fridaycount++
			}
		} else {
			fridaycount = 0
		}
		if saturday == true {
			if beginDate.Weekday() != 6 && beginDate.Weekday() != 4 && beginDate.Weekday() != 3 && beginDate.Weekday() != 2 && beginDate.Weekday() != 1 && beginDate.Weekday() != 0 {
				saturdaycount++
			}
		} else {
			saturdaycount = 0
		}
		if sunday == true {
			if beginDate.Weekday() != 5 && beginDate.Weekday() != 4 && beginDate.Weekday() != 3 && beginDate.Weekday() != 2 && beginDate.Weekday() != 1 && beginDate.Weekday() != 0 {
				sundaycount++
			}
		} else {
			sundaycount = 0
		}
		beginDate = beginDate.Add(time.Hour * 24)
	}
	days = mondaycount + tuesdaycount + wednesdaycount + thursdaycount + fridaycount + saturdaycount + sundaycount + 1
	return days
}

func (sh *Shared) BoolsToInt(flags ...bool) (value int) {
	value = 0
	for _, flag := range flags {
		if flag {
			value++
		}
	}
	return value
}

func (sh *Shared) MonthdaysInArray(firstOfMonth time.Time, lastOfMonth time.Time) (sv []string) {
	currentYear, currentMonth, _ := firstOfMonth.Date()
	currentLocation := firstOfMonth.Location()
	t := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)

	fmt.Println(firstOfMonth)
	fmt.Println(lastOfMonth)
	currentTimev2 := firstOfMonth
	Dayv2 := currentTimev2.Day()
	Monthv2 := currentTimev2.Month()
	Yearv2 := currentTimev2.Year()
	strDayv2 := strconv.Itoa(Dayv2)
	strMonthv2 := Monthv2.String()
	strYearv2 := strconv.Itoa(Yearv2)
	uniqueIdcurrentformatv2 := strDayv2 + strMonthv2 + strYearv2

	days := 0
	uniqueIdcurrentformat := ""
	var s []string

	for {
		if firstOfMonth.Equal(lastOfMonth) {
			days++
			break
		}
		if firstOfMonth.Weekday() != 8 && firstOfMonth.Weekday() != 9 {
			days++

		}
		daysvalue := -days
		listdays := t.AddDate(0, 1, daysvalue)
		//fmt.Println("number of listdays", listdays)
		currentTime := listdays
		Day := currentTime.Day()
		Month := currentTime.Month()
		Year := currentTime.Year()
		strDay := strconv.Itoa(Day)
		strMonth := Month.String()
		strYear := strconv.Itoa(Year)
		uniqueIdcurrentformat = strDay + strMonth + strYear + ","
		//fmt.Println("uniqueIdcurrentformat ==>", uniqueIdcurrentformat)
		s = append(s, uniqueIdcurrentformat)
		//fmt.Println("s value", s)

		firstOfMonth = firstOfMonth.Add(time.Hour * 24)
	}

	sv = append(s, uniqueIdcurrentformat)
	sv = append(s, uniqueIdcurrentformatv2)
	fmt.Println("sv value", sv)

	fmt.Println("number of days", days)
	return sv
}
