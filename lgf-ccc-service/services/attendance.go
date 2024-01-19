package service

import (
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

// SaveAttendance : ""
func (s *Service) SaveAttendance(ctx *models.Context, att *models.Attendance) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if att.UniqueID == "" {
		att.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONATTENDANCE)
	}
	att.Status = constants.ATTENDANCESTATUSPENDING
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Attendance.created")
	att.Created = &created
	log.Println("b4 Attendance.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAttendance(ctx, att)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}
func (s *Service) SaveAttendanceWithEditEmployee(ctx *models.Context, attandeace *models.Attendance) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	currentTime := attandeace.Date
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)
	uniqueIdcurrentformat := strDay + strMonth + strYear
	attandeace.UniqueID = uniqueIdcurrentformat
	fmt.Println("attandeace.UniqueID===>", attandeace.UniqueID)
	attandeace.Status = constants.ATTENDANCESTATUSPENDING
	attandeace.LoginMode = "Edit"

	var res1 []string
	if attandeace.TotalWorkingHoursStr != "" {
		res1 = strings.Split(attandeace.TotalWorkingHoursStr, ":")
		WorkingHours, err := strconv.ParseFloat(res1[0], 64)
		if err != nil {
			return err
		}
		WorkingHoursMins, err := strconv.ParseFloat(res1[1], 64)
		if err != nil {
			return err
		}
		fmt.Println("WorkingHours====>", WorkingHours)
		fmt.Println("WorkingHoursMins====>", WorkingHoursMins)
		attandeace.WorkingHours = WorkingHours
	}
	fmt.Println("TotalWorkingHoursStr====>", attandeace.TotalWorkingHoursStr)
	fmt.Println("res1====>", res1)
	// var WorkingHours float64
	// var WorkingHoursMins float64

	fmt.Println("TotalOverTimehoursstr====>", attandeace.TotalOverTimehoursstr)
	var res2 []string
	if attandeace.TotalOverTimehoursstr != "" {
		res2 = strings.Split(attandeace.TotalOverTimehoursstr, ":")
		fmt.Println("TotalWorkingHoursStr====>", attandeace.TotalOverTimehoursstr)
		fmt.Println("res1====>", res1)
		// var WorkingHours float64
		// var WorkingHoursMins float64
		OvertimeWorkingHours, err := strconv.ParseFloat(res2[0], 64)
		if err != nil {
			return err
		}
		OvertimeWorkingmins, err := strconv.ParseFloat(res2[1], 64)
		if err != nil {
			return err
		}
		attandeace.OverTime = OvertimeWorkingHours
		attandeace.OverTimeMins = OvertimeWorkingmins
	}

	attandeace.TotalWorkingMins = (attandeace.OverTime * 60) + (attandeace.WorkingHours * 60) + attandeace.WorkingHoursMins + attandeace.OverTimeMins
	if attandeace.TotalWorkingMins != 0 {
		hours := attandeace.TotalWorkingMins / 60
		Minutes := int64(attandeace.TotalWorkingMins) % 60
		attandeace.TotalWorkingHoursStr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
	}
	if attandeace.OverTime != 0 {
		hours := attandeace.OverTime
		Minutes := attandeace.OverTimeMins
		attandeace.TotalOverTimehoursstr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
		attandeace.OverTime = hours * 60

	}
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Attendance.created")
	attandeace.Created = &created
	log.Println("b4 Attendance.created")
	//att.PayRoll = constants.ATTENDANCESTATUSLOP
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.SaveAttendanceWithUpsert(ctx, attandeace)
		if err != nil {
			return errors.New("Edited Not Updated")
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		//	var attendancelog *models.AttendanceLog
		attendancelog := new(models.AttendanceLog)
		fmt.Println("uniqueIdcurrentformat===>", uniqueIdcurrentformat)
		fmt.Println("attandeace.Employeeid===>", attandeace.EmployeeId)
		attendancelog.UniqueID = uniqueIdcurrentformat
		attendancelog.EmployeeId = attandeace.EmployeeId
		attendancelog.WorkingMins = attandeace.WorkingHours * 60
		attendancelog.OverTime = attandeace.OverTime * 60
		attendancelog.Status = attandeace.Status
		attendancelog.Date = attandeace.Date
		t := time.Now()
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		log.Println("b4 Attendance.created")
		attendancelog.Created = &created
		err = s.Daos.SaveAttendanceLog(ctx, attendancelog)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

// GetSingleAttendance : ""
func (s *Service) GetSingleAttendance(ctx *models.Context, UniqueID string) (*models.RefAttendance, error) {
	attendance, err := s.Daos.GetSingleAttendance(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

// UpdateAttendance : ""
func (s *Service) UpdateAttendance(ctx *models.Context, attendance *models.Attendance) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	attendance.Status = constants.ATTENDANCEMANUALSTATUSLOGOUT
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAttendance(ctx, attendance)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

// EnableAttendance : ""
func (s *Service) EnableAttendance(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableAttendance(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableAttendance : ""
func (s *Service) DisableAttendance(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableAttendance(ctx, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}

// FilterAttendance : ""
func (s *Service) FilterAttendance(ctx *models.Context, ft *models.FilterAttendance, pagination *models.Pagination) (property []models.RefAttendance, err error) {
	return s.Daos.FilterAttendance(ctx, ft, pagination)
}

// ClockinAttendance : ""
func (s *Service) ClockinAttendance(ctx *models.Context, att *models.Attendance) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		loc, _ := time.LoadLocation("Asia/Kolkata")
		//	now := time.Now()
		currentTime := time.Now().In(loc)
		Day := currentTime.Day()
		Month := currentTime.Month()
		Year := currentTime.Year()
		strDay := strconv.Itoa(Day)
		strMonth := Month.String()
		strYear := strconv.Itoa(Year)
		uniqueIdcurrentformat := strDay + strMonth + strYear
		currentdate := time.Now()
		currentdate.Format(constants.DDMMYYYY)

		// user, err := s.Daos.GetSingleUser(ctx, att.EmployeeId)
		// if err != nil {
		// 	return err
		// }

		var workschedulehouse float64

		AttendanceVaild, err := s.Daos.GetSingleAttendanceByEmployeeIdAndStateValue(ctx, uniqueIdcurrentformat, att.EmployeeId, constants.ATTENDANCESTATUSLOGIN)
		if AttendanceVaild != nil {
			return errors.New("Already login")
		}
		Attendancecheck, err := s.Daos.GetSingleAttendanceByEmployeeIdAndStateValue(ctx, uniqueIdcurrentformat, att.EmployeeId, constants.ATTENDANCESTATUSLOGOUT)
		if err != nil {
			fmt.Println(err)
			return err
		}

		//find a date to days
		fmt.Println("currentTime==>", currentTime)
		var diffvalue time.Duration
		if Attendancecheck != nil {
			fmt.Println("Attendancecheck==>", Attendancecheck.EmployeeId)

			att.UniqueID = uniqueIdcurrentformat
			att.LoginMode = constants.ATTENDANCESTATUSLOGIN
			att.Status = constants.ATTENDANCESTATUSPENDING
			//		att.PunchIn = &currentTime
			att.Date = &currentdate
			//att.WorkingHours =
			a1 := Attendancecheck.PunchOut
			c2 := currentTime
			diffvalue = c2.Sub(*a1)
			//fmt.Println("diffvalue ==>", diffvalue)
			//att.BreakMins = diffvalue.Minutes()
			att.TotalBreakMins = Attendancecheck.TotalBreakMins + diffvalue.Minutes()
			if att.OverTime != 0 {
				if workschedulehouse == att.WorkingHours {
					att.OverTime = 0
				}
				if workschedulehouse < att.WorkingHours {
					att.OverTime = att.WorkingHours - workschedulehouse
				}
				if workschedulehouse > att.WorkingHours {
					att.OverTime = 0
				}
			}
			//	fmt.Println("attendenace over time", att.OverTime)

		} else {
			att.UniqueID = uniqueIdcurrentformat
			att.LoginMode = constants.ATTENDANCESTATUSLOGIN
			att.PayRoll = constants.ATTENDANCESTATUSLOP
			att.PunchIn = &currentTime
			att.Date = &currentdate

		}
		dberr := s.Daos.ClockinAttendancev2(ctx, att)
		if dberr != nil {
			return dberr
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		//AttendanceLog
		AttendanceLog := new(models.AttendanceLog)
		AttendanceLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONATTENDANCELOG)
		AttendanceLog.AttendanceID = uniqueIdcurrentformat
		AttendanceLog.EmployeeId = att.EmployeeId
		AttendanceLog.OrganisationId = att.OrganisationId
		AttendanceLog.PunchinTime = &currentTime
		AttendanceLog.Date = &currentdate
		AttendanceLog.LoginMode = constants.ATTENDANCESTATUSLOGIN
		AttendanceLog.Status = constants.ATTENDANCESTATUSACTIVE
		AttendanceLog.Notes = att.Notes
		AttendanceLog.BreakMins = diffvalue.Minutes()
		dberr = s.Daos.SaveAttendanceLog(ctx, AttendanceLog)
		if dberr != nil {
			return dberr
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

// LogoutAttendance : ""
func (s *Service) ClockoutAttendance(ctx *models.Context, att *models.ClockoutAttendance) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		currentTime := time.Now()
		Day := currentTime.Day()
		Month := currentTime.Month()
		Year := currentTime.Year()
		strDay := strconv.Itoa(Day)
		strMonth := Month.String()
		strYear := strconv.Itoa(Year)
		uniqueIdcurrentformat := strDay + strMonth + strYear
		//fmt.Println("uniqueIdcurrentformat ==>", uniqueIdcurrentformat)
		currentdate := time.Now()
		currentdate.Format(constants.DDMMYYYY)
		Attendance, err := s.Daos.GetSingleAttendanceByEmployeeIdAndStateValue(ctx, uniqueIdcurrentformat, att.EmployeeId, constants.ATTENDANCESTATUSLOGIN)
		if err != nil {
			fmt.Println(err)
			return errors.New("Attendance Not Found-Pls Clock in")
		}
		Attendance.LoginMode = constants.ATTENDANCESTATUSLOGOUT
		Attendance.PayRoll = constants.ATTENDANCESTATUSPAID
		Attendance.Status = constants.ATTENDANCESTATUSPENDING
		Attendance.PunchOut = &currentTime
		fmt.Println("attendenace over time", Attendance.OverTime)
		days := time.Now().Weekday().String()
		fmt.Println("days string====>", days)
		RefAttendancelog, err := s.Daos.GetSingleLastAttendanceLogByEmployeeIdAndState(ctx, uniqueIdcurrentformat, att.EmployeeId, constants.ATTENDANCESTATUSLOGIN)
		if err != nil {
			return err
		}
		diff := Attendance.PunchOut.Sub(*RefAttendancelog.PunchinTime)
		Attendance.OverTime = diff.Minutes() + Attendance.OverTime
		Attendance.WorkingHours = 0
		Attendance.TotalWorkingMins = Attendance.OverTime
		if Attendance.TotalWorkingMins != 0 {
			hours := Attendance.TotalWorkingMins / 60
			Minutes := int64(Attendance.TotalWorkingMins) % 60
			Attendance.TotalWorkingHoursStr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
		}
		if Attendance.TotalBreakMins != 0 {
			hours := Attendance.TotalBreakMins / 60
			Minutes := int64(Attendance.TotalBreakMins) % 60
			Attendance.TotalBreakHoursStr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
		}
		if Attendance.OverTime != 0 {
			hours := Attendance.OverTime / 60
			Minutes := int64(Attendance.OverTime) % 60
			Attendance.TotalOverTimehoursstr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
		}
		dberr := s.Daos.ClockoutAttendance(ctx, &Attendance.Attendance)
		if dberr != nil {
			return dberr
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		AttendanceLog := new(models.AttendanceLog)
		if RefAttendancelog != nil {
			AttendanceLog.UniqueID = RefAttendancelog.UniqueID
			AttendanceLog.AttendanceID = RefAttendancelog.AttendanceID
			AttendanceLog.PunchoutTime = &currentTime
			diff2 := currentTime.Sub(*RefAttendancelog.PunchinTime)
			AttendanceLog.LoginMode = constants.ATTENDANCESTATUSLOGOUT
			AttendanceLog.Status = constants.ATTENDANCESTATUSACTIVE
			AttendanceLog.Date = &currentdate
			AttendanceLog.Notes = att.Notes
			//	AttendanceLog.OverTime = Attendance.OverTime
			AttendanceLog.WorkingMins = diff2.Minutes()

			dberr = s.Daos.UpdateAttendanceLog(ctx, AttendanceLog)
			if dberr != nil {
				return dberr
			}
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

// EmployeeAttendanceTodayStatus : ""
func (s *Service) EmployeeAttendanceTodayStatus(ctx *models.Context, EmployeeId string, UniqueID string) (*models.EmployeeAttendanceTodayStatus, error) {
	Attendance, err := s.Daos.GetSingleEmployeeAttendanceTodayStatus(ctx, UniqueID, EmployeeId)
	if err != nil {
		return nil, err
	}

	AttendanceLog, err := s.Daos.AttendanceEmployeeTodaystatuswithouttotaltime(ctx, EmployeeId, UniqueID)
	if err != nil {
		return nil, err
	}
	Employee, err := s.Daos.GetSingleUser(ctx, EmployeeId)
	if err != nil {
		return nil, err
	}
	if Employee == nil {
		return nil, errors.New("Employee id is not found")
	}

	// if WorkSchedule == nil {
	// 	return nil, errors.New("No work schedule is access to this employee id")
	// }
	RecentRecord, err := s.Daos.GetSingleAttendanceLoglast(ctx, EmployeeId, UniqueID)
	if err != nil {
		return nil, err
	}
	if RecentRecord == nil {
		return nil, errors.New("No work schedule is access to this employee id")
	}
	Attendance.RecentPunchinTime = RecentRecord.PunchinTime
	Attendance.RecentPunchoutTime = RecentRecord.PunchoutTime
	Attendance.FirstPunchinTime = AttendanceLog.FirstPunchinTime
	Attendance.LastpunchoutTime = AttendanceLog.LastpunchoutTime

	fmt.Println("attendenace over time", AttendanceLog.OverTime)
	return Attendance, nil
}

// AttendanceEmployeeLeaveRequest : ""
func (s *Service) AttendanceEmployeeLeaveRequest(ctx *models.Context, att *models.Attendance) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	currentTime := time.Now()
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)
	uniqueIdcurrentformat := strDay + strMonth + strYear
	currentdate := time.Now()
	currentdate.Format(constants.DDMMYYYY)

	att.UniqueID = uniqueIdcurrentformat
	att.Status = constants.ATTENDANCESTATUSLEAVE
	att.LoginMode = constants.ATTENDANCESTATUSREQUEST
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Attendance Employee Leave.created")
	att.Created = &created
	log.Println("b4 Attendance Employee Leave.created")

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAttendance(ctx, att)
		if dberr != nil {
			return dberr
		}

		employeeleavelog := new(models.EmployeeLeave)
		employeeleavelog.Description = att.Notes
		employeeleavelog.Name = att.EmployeeId + constants.ATTENDANCESTATUSREQUEST
		employeeleavelog.LeaveType = att.LeaveType
		employeeleavelog.EmployeeId = att.EmployeeId
		employeeleavelog.Date = att.Date
		employeeleavelog.Status = constants.ATTENDANCESTATUSREQUEST
		// dberr = s.Daos.SaveEmployeeLeave(ctx, employeeleavelog)
		// if dberr != nil {
		// 	return dberr
		// }

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}
func (s *Service) DayWiseAttendanceReport(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (att *models.DayWiseAttendanceReport, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	Employee, err := s.Daos.GetSingleUser(ctx, filter.EmployeeId)
	if err != nil {
		return nil, err
	}
	if Employee == nil {
		return nil, errors.New("Employee id is not found")
	}
	t := time.Now()
	if Employee.JoiningDate != nil {
		if Employee.JoiningDate.Year() == t.Year() && Employee.JoiningDate.Month() == t.Month() {
			filter.StartDate = Employee.JoiningDate
			fmt.Println("JoiningDate===>", Employee.JoiningDate)

		} else {
			sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
			filter.StartDate = &sd
		}
	}
	fmt.Println("startDate===>", filter.StartDate)
	attendancereport, err := s.Daos.DayWiseAttendanceReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	if attendancereport != nil {
		attendancereport.EmployeeName = Employee.Name

	}

	n := len(attendancereport.Days)
	fmt.Println("length of attandeace===>", n)
	if len(attendancereport.Days) > 0 {
		var Paidtime float64
		var Deficttime float64
		for k, v := range attendancereport.Days {

			if v.TotalWorkingMins < (attendancereport.Days[k].PaidTime * 60) {
				attendancereport.Days[k].Deficit = v.TotalWorkingMins - (attendancereport.Days[k].PaidTime * 60)
				hours := attendancereport.Days[k].Deficit / 60
				Minutes := -(int64(attendancereport.Days[k].Deficit) % 60)
				attendancereport.Days[k].DeficitStr = fmt.Sprintf("%vh %vmins", int64(hours), Minutes)
			} else {
				attendancereport.Days[k].DeficitStr = fmt.Sprintf("%vh %vmins", 0, 0)
			}
			//Calc TotalWorkingMinsStr
			hours := v.TotalWorkingMins / 60
			Minutes := int64(v.TotalWorkingMins) % 60
			attendancereport.Days[k].TotalWorkingMinsStr = fmt.Sprintf("%vh %vmins", int64(hours), Minutes)
			//Calc TotalOvertimeStr
			hours = v.OverTime / 60
			Minutes = int64(v.OverTime) % 60
			attendancereport.Days[k].TotalOvertimeStr = fmt.Sprintf("%vh %vmins", int64(hours), Minutes)
			//Calc
			Paidtime = Paidtime + attendancereport.Days[k].PaidTime
			//calc defecit
			Deficttime = Deficttime + attendancereport.Days[k].Deficit
		}
		attendancereport.PaidTime = Paidtime
		attendancereport.WorkSchedule = Paidtime
		attendancereport.Deficit = Deficttime
	}
	// if len(attendancereport.Days) > 0 {

	// 	fmt.Println("PaidTime==", attendancereport.PaidTime)
	// 	fmt.Println("TotalWorkingMins==", attendancereport.TotalWorkingMins/60)
	// 	// if attendancereport.PaidTime > (attendancereport.TotalWorkingMins / 60) {
	// 	// 	attendancereport.Deficit = attendancereport.TotalWorkingMins - (attendancereport.PaidTime * 60)
	// 	// }
	// }
	for k, v := range attendancereport.Days {
		attendancereport.EmployeeName = Employee.Name
		currentTime := v.Date
		var date string
		Day := currentTime.Day()
		Month := currentTime.Month()
		Year := currentTime.Year()
		strDay := strconv.Itoa(Day)
		strMonth := Month.String()
		strYear := strconv.Itoa(Year)
		uniqueIdcurrentformat := strDay + strMonth + strYear
		//	date = fmt.Sprintf("%v-%v-%v", v.Date.Day(), int(v.Date.Month()), v.Date.Year())
		attendancereport.Days[k].UniqueID = uniqueIdcurrentformat
		if v.Date.Month() < 10 {
			if v.Date.Day() < 10 {
				date = fmt.Sprintf("0%v-0%v-%v", v.Date.Day(), int(v.Date.Month()), v.Date.Year())
			} else {
				date = fmt.Sprintf("%v-0%v-%v", v.Date.Day(), int(v.Date.Month()), v.Date.Year())
			}
		} else {
			if v.Date.Day() < 10 {
				date = fmt.Sprintf("0%v-%v-%v", v.Date.Day(), int(v.Date.Month()), v.Date.Year())
			} else {
				date = fmt.Sprintf("%v-%v-%v", v.Date.Day(), int(v.Date.Month()), v.Date.Year())
			}
		}
		attendancereport.Days[k].DateStr = date
	}
	//Calc TotalWorkingMinsStr
	hours := attendancereport.TotalWorkingMins / 60
	Minutes := int64(attendancereport.TotalWorkingMins) % 60
	attendancereport.TotalWorkingMinsStr = fmt.Sprintf("%vh %vmins", int64(hours), Minutes)
	fmt.Println("TotalWorkingMinsStr==>", attendancereport.TotalWorkingMinsStr)
	//Calc TotalOvertimeStr
	hours = attendancereport.TotalOvertime / 60
	Minutes = int64(attendancereport.TotalOvertime) % 60
	attendancereport.TotalOvertimeStr = fmt.Sprintf("%vh %vmins", int64(hours), Minutes)
	fmt.Println("TotalOvertimeStr==>", attendancereport.TotalOvertimeStr)
	//Calc DeficitStr
	hours = attendancereport.Deficit / 60
	Minutes = -(int64(attendancereport.Deficit) % 60)
	attendancereport.DeficitStr = fmt.Sprintf("%vh %vmins", int64(hours), Minutes)
	fmt.Println("DeficitStr==>", attendancereport.DeficitStr)

	return attendancereport, err
}
func (s *Service) WeeklyWiseAttendanceReport(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (*models.DayWiseAttendanceReport, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	attendancereport, err := s.Daos.WeaklyWiseAttendanceReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	Employee, err := s.Daos.GetSingleUser(ctx, filter.EmployeeId)
	if err != nil {
		return nil, err
	}
	if Employee == nil {
		return nil, errors.New("Employee id is not found")
	}

	n := len(attendancereport.Days)
	fmt.Println("Weekly length of attandeace===>", n)
	if len(attendancereport.Days) > 0 {
		// attendancereport.PaidTime = WorkSchedule.WorkingHoursinDay * float64(n)
		// attendancereport.WorkSchedule = WorkSchedule.WorkingHoursinDay * float64(n)
		fmt.Println("Weekly PaidTime==", attendancereport.PaidTime)
		fmt.Println("Weekly WorkSchedule==", attendancereport.WorkSchedule)
		fmt.Println("Weekly TotalWorkingMins==", attendancereport.TotalWorkingMins/60)
		if attendancereport.PaidTime > (attendancereport.TotalWorkingMins / 60) {

			attendancereport.Deficit = attendancereport.TotalWorkingMins - (attendancereport.TotalWorkingMins * 60)
		}
	}

	return attendancereport, err
}

// AttendanceEmployeeStatistics : ""
func (s *Service) AttendanceEmployeeStatistics(ctx *models.Context, EmployeeId string) (*models.AttendanceEmployeeStatistics, error) {
	currentTime := time.Now()
	Day := currentTime.Day()
	Months := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Months.String()
	strYear := strconv.Itoa(Year)
	uniqueIdcurrentformat := strDay + strMonth + strYear
	AttendanceLog := new(models.AttendanceEmployeeStatistics)
	AttendanceLogs, err := s.Daos.AttendanceEmployeeStatistics(ctx, EmployeeId, uniqueIdcurrentformat)
	if err != nil {
		return nil, err
	}
	RecentLoginCheck, err := s.Daos.RecentLoginCheck(ctx, EmployeeId, uniqueIdcurrentformat)
	if err != nil {
		return nil, err
	}
	fmt.Println("RecentLoginCheck===>", RecentLoginCheck)
	filter := new(models.DayWiseAttendanceReportFilter)
	t := time.Now()
	filter.StartDate = &t
	filter.EmployeeId = EmployeeId
	Month, err := s.DayWiseAttendanceReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	week, err := s.WeeklyWiseAttendanceReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	Employee, err := s.Daos.GetSingleUser(ctx, EmployeeId)
	if err != nil {
		return nil, err
	}
	if Employee == nil {
		return nil, errors.New("Employee id is not found")
	}

	if Month != nil {
		fmt.Println("Deficit===>", Month.Deficit)

		AttendanceLog.MonthTotalWorkingHours = Month.WorkSchedule
		AttendanceLog.MonthWorkingHours = Month.TotalWorkingMins / 60
		AttendanceLog.RemainingWorkingHours = Month.Deficit / 60
		AttendanceLog.TotalOvertime = Month.TotalOvertime / 60
	}
	if week != nil {
		AttendanceLog.WeekTotalWorkingHours = week.WorkSchedule
		AttendanceLog.WeekWorkingHours = week.TotalWorkingMins / 60
	}

	if AttendanceLogs == nil {
		AttendanceLog.TotalWorkingMins = 0
	} else {
		if RecentLoginCheck != nil {
			t := time.Now()
			fmt.Println("currentTime", t)
			fmt.Println("RecentLoginCheck.currentTime", RecentLoginCheck.PunchinTime)
			currentTime := t.Sub(*RecentLoginCheck.PunchinTime)
			fmt.Println("RecentLoginCheck.currentTime", currentTime)
			AttendanceLog.TotalWorkingMins = AttendanceLogs.TotalWorkingMins + currentTime.Minutes()
			hours := AttendanceLog.TotalWorkingMins / 60
			//	Minutes := (int64(AttendanceLog.TotalWorkingMins) % 60)
			AttendanceLog.TotalWorkingMinsStr = fmt.Sprintf("%.2f", hours)
		} else {
			AttendanceLog.TotalWorkingMins = AttendanceLogs.TotalWorkingMins
			hours := AttendanceLog.TotalWorkingMins / 60
			//	Minutes := (int64(AttendanceLog.TotalWorkingMins) % 60)
			AttendanceLog.TotalWorkingMinsStr = fmt.Sprintf("%.2f", hours)
		}

	}
	fmt.Println("MonthTotalWorkingHours===>", AttendanceLog.MonthTotalWorkingHours)
	fmt.Println("MonthWorkingHours===>", AttendanceLog.MonthWorkingHours)
	fmt.Println("RemainingWorkingHours===>", AttendanceLog.RemainingWorkingHours)
	fmt.Println("TotalOvertime===>", AttendanceLog.TotalOvertime)
	fmt.Println("WeekTotalWorkingHours===>", AttendanceLog.WeekTotalWorkingHours)
	fmt.Println("WeekWorkingHours===>", AttendanceLog.WeekWorkingHours)
	fmt.Println("TodayTotalWorkingHours===>", AttendanceLog.TodayTotalWorkingHours)
	fmt.Println("TotalWorkingMins===>", AttendanceLog.TotalWorkingMins)
	fmt.Println("TotalWorkingMinsstr===>", AttendanceLog.TotalWorkingMinsStr)
	AttendanceLog.TotalOvertimeStr = fmt.Sprintf("%.2f", AttendanceLog.TotalOvertime)
	AttendanceLog.TodayTotalWorkingMinsStr = fmt.Sprintf("%.2f", AttendanceLog.TodayTotalWorkingHours)
	AttendanceLog.WeekWorkingHoursStr = fmt.Sprintf("%.2f", AttendanceLog.WeekWorkingHours)
	AttendanceLog.MonthWorkingHoursStr = fmt.Sprintf("%.2f", AttendanceLog.MonthWorkingHours)
	AttendanceLog.RemainingWorkingHoursStr = fmt.Sprintf("%.2f", AttendanceLog.RemainingWorkingHours)

	return AttendanceLog, nil
}

func (s *Service) TodayEmployessLeave(ctx *models.Context, UniqueID string) (*models.TodayEmployessLeave, error) {
	attendance, err := s.Daos.TodayEmployessLeave(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}
func (s *Service) DayWiseAttendanceReportExcel(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (*excelize.File, error) {
	t := time.Now()
	data, err := s.DayWiseAttendanceReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "EmployeeDayWiseReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "H1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	// title :=
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v-%v", sheet1, data.EmployeeName))
	Kitchen := "3:04PM"
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "ClockIn")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "ClockOut")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "LoggedTime")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "PaidTime")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "PaidTime")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "OverTime")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Deficit")
	rowNo++

	//	var totalAmount float64
	for k, v := range data.Days {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year()))
		if v.PunchIn != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(v.PunchIn.Format(Kitchen)))
		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "00:00")
		}
		if v.PunchOut != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(v.PunchOut.Format(Kitchen)))
		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "00:00")
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.TotalWorkingMinsStr)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%vh", v.PaidTime))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.TotalOvertimeStr)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.DeficitStr)
		rowNo++
	}
	//excel.MergeCell(sheet1, "A", "D")
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	a := fmt.Sprintf("A%v", rowNo)
	c := fmt.Sprintf("D%v", rowNo)
	excel.MergeCell(sheet1, a, c)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), data.TotalWorkingMinsStr)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%vh", data.PaidTime))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), data.TotalOvertimeStr)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), data.DeficitStr)

	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) EmployeeAttendanceApprove(ctx *models.Context, Attendance *models.EmployeeAttendanceApprove) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range Attendance.UniqueID {
			dberr := s.Daos.EmployeeAttendanceApprove(ctx, Attendance.EmployeeId, v)
			if dberr != nil {
				return dberr
			}
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

func (s *Service) GetTodayEmployeeTimeOff(ctx *models.Context) ([]models.TodayTimeoff, error) {
	attendance, err := s.Daos.GetTodayEmployeeTimeOff(ctx)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}
func (s *Service) EmployeeAttendanceRejected(ctx *models.Context, Attendance *models.EmployeeAttendanceApprove) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EmployeeAttendanceRejected(ctx, Attendance)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}
func (s *Service) GetTodayEmployeePunchin(ctx *models.Context) ([]models.TodayPunchin, error) {
	attendance, err := s.Daos.GetTodayEmployeePunchin(ctx)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}
func (s *Service) EmployeeDayWiseAttendanceReportExcel(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (*excelize.File, error) {
	t := time.Now()
	data, err := s.EmployeeDayWiseAttendanceReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()
	// var k []string
	// for _, v := range data {
	// 	if len(v.Days) > 0 {
	// 		switch len(v.Days) {
	// 		case 31:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG")
	// 		case 30:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF")
	// 		case 28:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD")
	// 		default:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE")
	// 		}
	// 	}
	// }
	// n := len(k)
	excel := excelize.NewFile()
	sheet1 := "EmployeesDayWiseReport"
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	rowNo := 1
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	Lop, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FF0000"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":false}}`)
	if err != nil {
		fmt.Println(err)
	}
	paid, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#00FF00"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":false}}`)
	if err != nil {
		fmt.Println(err)
	}

	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	// title :=

	//Kitchen := "22-08-2022"
	rowNo++
	rowNo++
	if len(data) > 0 {
		ch := 'A'
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), "S.No")
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), "EmployeeName")
		ch++
		var ch1 rune
		var ch2 rune
		ch1 = 'A'
		ch2 = 'A'
		column := ""
		for _, v := range data[0].Days {
			rowNo = 3
			if ch < 'Z' {
				column = fmt.Sprintf("%c", ch)
				day := fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)

			} else if ch == 'Z' {
				column = fmt.Sprintf("%c", ch)
				day := fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)
			}
			if ch > 'Z' {
				CH1 := fmt.Sprintf("%c", ch1)
				CH2 := fmt.Sprintf("%c", ch2)
				fmt.Println("ch1===>", CH1)
				fmt.Println("ch2===>", CH2)
				column = fmt.Sprintf("%c%c", ch1, ch2)
				day := fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)

				//CH := fmt.Sprintf("%c", ch)
				fmt.Println("column===>", column)
			}
			//CH := fmt.Sprintf("%c", ch)
			fmt.Println("column===>", column)
			if ch > 'Z' {
				ch2++
			}
			ch++
		}
		CH1 := fmt.Sprintf("%c", ch1)
		CH2 := fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.Paid")
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.PartialPay")
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.Lop")
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.Holidays")
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", column, rowNo), style1)

		column1 := fmt.Sprintf("%v%v", column, 1)
		excel.MergeCell(sheet1, "A1", column1)
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", 1), fmt.Sprintf("%v%v", column, 1), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", 1), fmt.Sprintf("%v-%v%v", sheet1, filter.StartDate.Month(), filter.StartDate.Year()))

	}
	rowNo++
	rowNo++
	//	var totalAmount float64
	for k, v := range data {
		ch := 'A'
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), k+1)
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), v.Name)
		ch++
		var ch1 rune
		var ch2 rune
		ch1 = 'A'
		ch2 = 'A'
		column := ""
		for _, v2 := range v.Days {
			if ch < 'Z' {
				column = fmt.Sprintf("%c", ch)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v2.Attendance.PayRoll)
				if v2.Attendance.PayRoll == "LOP" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), Lop)
				} else if v2.Attendance.PayRoll == "Paid" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), paid)
				}

			} else if ch == 'Z' {
				column = fmt.Sprintf("%c", ch)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v2.Attendance.PayRoll)
				if v2.Attendance.PayRoll == "LOP" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), Lop)
				} else if v2.Attendance.PayRoll == "Paid" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), paid)
				}
			}
			if ch > 'Z' {
				CH1 := fmt.Sprintf("%c", ch1)
				CH2 := fmt.Sprintf("%c", ch2)
				fmt.Println("ch1===>", CH1)
				fmt.Println("ch2===>", CH2)
				column = fmt.Sprintf("%c%c", ch1, ch2)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v2.Attendance.PayRoll)
				if v2.Attendance.PayRoll == constants.ATTENDANCESTATUSLOGIN {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), Lop)
				} else if v2.Attendance.PayRoll == constants.ATTENDANCESTATUSLOGOUT {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), paid)
				}
				//CH := fmt.Sprintf("%c", ch)
				if ch > 'Z' {
					ch2++
				}

			}
			fmt.Println("column Row===>", column)

			ch++
		}
		CH1 := fmt.Sprintf("%c", ch1)
		CH2 := fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfPaid)
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfParticalPaid)
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfLOP)
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfHolidays)
		rowNo++
	}
	//excel.MergeCell(sheet1, "A", "D")
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) EmployeeDayWiseAttendanceReport(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter, pagination *models.Pagination) ([]models.EmployeeDayWiseAttendanceReport, error) {
	err := s.EmployeeDayWiseAttendanceAccess(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employereport, err := s.Daos.EmployeeDayWiseAttendanceReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	for k, v := range employereport {
		var overtime float64
		for k2, v2 := range v.Days {
			overtime = overtime + v2.Attendance.OverTime
			v.Days[k2].Dates = int64(k2) + 1
			date := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), k2+1, 0, 0, 0, 0, filter.StartDate.Location())
			v.Days[k2].Date = &date
			///fmt.Println("Payroll===>", v2.Attendance.PayRoll)
			if v2.Attendance.TotalWorkingHoursStr == "" {
				v.Days[k2].Attendance.TotalWorkingHoursStr = fmt.Sprintf("%vh%vmins", 0, 0)
			}
			if v2.Attendance.Status == constants.ATTENDANCESTATUSPENDING {
				employereport[k].PendingStatus = employereport[k].PendingStatus + 1
			}
		}
		hours := overtime / 60
		Minutes := int64(overtime) % 60
		employereport[k].Overtime = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
		fmt.Println("Employee====Dayes", v.Name, len(v.Days))
	}

	return employereport, nil

}
func (s *Service) EmployeeDayWiseAttendanceAccess(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationId = append(filter.OrganisationId, v.UniqueID)
				}
			}
			fmt.Println(" dataaccess.SuperAdmin===>", dataaccess.SuperAdmin)
			if dataaccess.SuperAdmin {
				filter.Manager = ""
			}

		}

	}
	return err
}
