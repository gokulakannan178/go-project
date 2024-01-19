package services

import (
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveAttendanceLog : ""
func (s *Service) SaveAttendanceLog(ctx *models.Context, att *models.AttendanceLog) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	att.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONATTENDANCELOG)
	att.Status = constants.ATTENDANCELOGSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AttendanceLog.created")
	att.Created = &created
	log.Println("b4 AttendanceLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAttendanceLog(ctx, att)
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

// GetSingleAttendanceLog : ""
func (s *Service) GetSingleAttendanceLog(ctx *models.Context, UniqueID string) (*models.RefAttendanceLog, error) {
	AttendanceLog, err := s.Daos.GetSingleAttendanceLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return AttendanceLog, nil
}

// GetSingleAttendanceLoglast : ""
func (s *Service) GetSingleAttendanceLoglast(ctx *models.Context, EmployeeId string, UniqueID string) (*models.RefAttendanceLog, error) {
	AttendanceLog, err := s.Daos.GetSingleAttendanceLoglast(ctx, EmployeeId, UniqueID)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	beginvalue := s.Shared.BeginningOfMonth(t)
	fmt.Println("Beginvalue :", beginvalue)
	endvalue := s.Shared.EndOfMonth(t)
	fmt.Println("Endvalue :", endvalue)
	beginvalues := beginvalue.Format(constants.DDMMYYYY)
	fmt.Println("beginvalues :", beginvalues)
	beginvaluestime, error := time.Parse(constants.DDMMYYYY, beginvalues)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	fmt.Println("beginvaluestime :", beginvaluestime)
	endvalues := endvalue.Format(constants.DDMMYYYY)
	fmt.Println("endvalues :", endvalues)
	endvaluestime, error := time.Parse(constants.DDMMYYYY, endvalues)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	fmt.Println("endvaluestime :", endvaluestime)

	monthvalue := s.Shared.MonthdaysInArray(beginvaluestime, endvaluestime)
	fmt.Println("Monthvalue :", monthvalue)
	return AttendanceLog, nil
}

// AttendanceEmployeeTodayStatus : ""
func (s *Service) AttendanceEmployeeTodayStatus(ctx *models.Context, EmployeeId string) (*models.AttendanceEmployeeTodayStatus, error) {
	currentTime := time.Now()
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)
	UniqueID := strDay + strMonth + strYear
	AttendanceLog := new(models.AttendanceEmployeeTodayStatus)
	fmt.Println("date==>", UniqueID)
	AttendanceLogs, err := s.Daos.AttendanceEmployeeTodayStatus(ctx, EmployeeId, UniqueID)
	if err != nil {
		return nil, err
	}
	Attendance, err := s.Daos.GetSingleAttendanceByEmployeeId(ctx, UniqueID, EmployeeId)
	if err != nil {
		return nil, err
	}
	RecentRecord, err := s.Daos.GetSingleAttendanceLoglast(ctx, EmployeeId, UniqueID)
	if err != nil {
		return nil, err
	}

	if AttendanceLogs != nil {
		AttendanceLog.RecentPunchinTime = RecentRecord.PunchinTime
		AttendanceLog.RecentPunchoutTime = RecentRecord.PunchoutTime
		AttendanceLog.FirstPunchinTime = AttendanceLogs.FirstPunchinTime
		AttendanceLog.LastpunchoutTime = AttendanceLogs.LastpunchoutTime
		AttendanceLog.TotalWorkingHours = Attendance.TotalWorkingMins
		if AttendanceLog.FirstPunchinTime != nil {
			if AttendanceLog.LastpunchoutTime == nil {
				AttendanceLog.CurrenntStatus = constants.ATTENDANCESTATUSLOGIN
			} else {
				AttendanceLog.CurrenntStatus = constants.ATTENDANCESTATUSLOGOUT
			}
		} else {
			AttendanceLog.CurrenntStatus = constants.ATTENDANCESTATUSLOGOUT

		}

		if Attendance != nil {
			if Attendance.LoginMode == constants.ATTENDANCESTATUSTIMEOFF {
				AttendanceLog.IsTimeoffCheck = true
			}
			if AttendanceLog.CurrenntStatus == constants.ATTENDANCESTATUSLOGIN {
				t := time.Now()
				fmt.Println("currentTime", t)
				fmt.Println("RecentLoginCheck.currentTime", AttendanceLog.RecentPunchinTime)
				currentTime := t.Sub(*AttendanceLog.RecentPunchinTime)
				fmt.Println("RecentLoginCheck.currentTime", currentTime)
				AttendanceLog.TotalWorkingHours = Attendance.TotalWorkingMins + currentTime.Minutes()
				Attendance.TotalWorkingMins = AttendanceLog.TotalWorkingHours
			} else {
				AttendanceLog.TotalWorkingHours = Attendance.TotalWorkingMins

			}
			hours := Attendance.TotalWorkingMins / 60
			Minutes := int64(Attendance.TotalWorkingMins) % 60
			AttendanceLog.TotalWorkingHoursStr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
			fmt.Println("TotalWorkingMinsStr==>", AttendanceLog.TotalWorkingHoursStr)
		}
		AttendanceLog.OverTime = Attendance.OverTime
		if Attendance.OverTime != 0 {
			hours := Attendance.OverTime / 60
			Minutes := int64(Attendance.OverTime) % 60
			AttendanceLog.OverTimeStr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
			fmt.Println("OverTimeStr==>", AttendanceLog.OverTimeStr)
		}
		AttendanceLog.TotalBreakHours = Attendance.TotalBreakMins
		if Attendance.TotalBreakMins != 0 {
			hours := Attendance.TotalBreakMins / 60
			Minutes := int64(Attendance.TotalBreakMins) % 60
			AttendanceLog.TotalBreakHoursStr = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
			fmt.Println("TotalBreakHours==>", AttendanceLog.TotalBreakHoursStr)
		}
	} else {
		AttendanceLog.RecentPunchinTime = nil
		AttendanceLog.RecentPunchoutTime = nil
		AttendanceLog.FirstPunchinTime = nil
		AttendanceLog.LastpunchoutTime = nil
		AttendanceLog.TotalWorkingHours = 0
		AttendanceLog.TotalBreakHours = 0
		AttendanceLog.OverTime = 0

	}
	fmt.Println("attendenace over time", AttendanceLog.OverTime)
	loc, _ := time.LoadLocation("Asia/Kolkata")
	if AttendanceLog.RecentPunchinTime != nil {
		*AttendanceLog.RecentPunchinTime = AttendanceLog.RecentPunchinTime.In(loc)
	}
	if AttendanceLog.RecentPunchoutTime != nil {
		*AttendanceLog.RecentPunchoutTime = AttendanceLog.RecentPunchoutTime.In(loc)

	}
	if AttendanceLog.FirstPunchinTime != nil {
		*AttendanceLog.FirstPunchinTime = AttendanceLog.FirstPunchinTime.In(loc)
	}
	if AttendanceLog.LastpunchoutTime != nil {
		*AttendanceLog.LastpunchoutTime = AttendanceLog.LastpunchoutTime.In(loc)

	}

	return AttendanceLog, nil
}
func (s *Service) AttendanceEmployeeTodayLogs(ctx *models.Context, EmployeeId string) ([]models.AttendanceLog, error) {
	currentTime := time.Now()
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)
	UniqueID := strDay + strMonth + strYear
	//	AttendanceLog := new(models.AttendanceEmployeeTodayStatus)
	fmt.Println("date==>", UniqueID)
	AttendanceLogs, err := s.Daos.AttendanceEmployeeTodayLogs(ctx, EmployeeId, UniqueID)
	if err != nil {
		return nil, err
	}

	return AttendanceLogs, nil
}

// UpdateAttendanceLog : ""
func (s *Service) UpdateAttendanceLog(ctx *models.Context, AttendanceLog *models.AttendanceLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAttendanceLog(ctx, AttendanceLog)
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

// EnableAttendanceLog : ""
func (s *Service) EnableAttendanceLog(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableAttendanceLog(ctx, uniqueID)
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

// DisableAttendanceLog : ""
func (s *Service) DisableAttendanceLog(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableAttendanceLog(ctx, uniqueID)
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

// DisableAttendanceLog : ""
func (s *Service) DeleteAttendanceLog(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DeleteAttendanceLog(ctx, uniqueID)
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

// FilterAttendanceLog : ""
func (s *Service) FilterAttendanceLog(ctx *models.Context, attendance *models.FilterAttendanceLog, pagination *models.Pagination) (attendances []models.RefAttendanceLog, err error) {
	return s.Daos.FilterAttendanceLog(ctx, attendance, pagination)
}
