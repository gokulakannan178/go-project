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

// SaveEmployeeTimeOff : ""
func (s *Service) SaveEmployeeTimeOff(ctx *models.Context, employeeTimeOff *models.EmployeeTimeOff) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeTimeOff.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEETIMEOFF)
	employeeTimeOff.Status = constants.EMPLOYEETIMEOFFSTATUSREQUEST
	employeeTimeOff.Revoke = constants.EMPLOYEETIMEOFFSTATUSREQUEST
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeTimeOff.created")
	employeeTimeOff.Created = &created
	loc, _ := time.LoadLocation("Asia/Kolkata")
	log.Println("b4 EmployeeTimeOff.created")
	employeeTimeOff.StartDate.In(loc)
	employeeTimeOff.EndDate.In(loc)
	startDate := employeeTimeOff.StartDate
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 1, startDate.Location())
	//sd := s.Shared.BeginningOfMonth(*employeeTimeOff.StartDate)
	ed := s.Shared.EndOfMonth(*employeeTimeOff.StartDate)
	endDate := employeeTimeOff.EndDate
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 23, 59, 59, 0, endDate.Location())
	fmt.Println("employeeTimeOff.StartDate==>", employeeTimeOff.StartDate)
	fmt.Println(" employeeTimeOff.EndDate==>", employeeTimeOff.EndDate)
	fmt.Println("start==>", start)
	fmt.Println("end==>", end)
	fmt.Println("ed==>", ed)
	if endDate != nil {
		// result := end.Day() - (start.Day())
		// if ed.Day() == endDate.Day() {
		// 	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day()-1, 0, 0, 0, 1, startDate.Location())
		// 	result = endDate.Day() - (start.Day())
		// }
		fmt.Println("end====>", end)
		fmt.Println("start====>", start)

		result := end.Sub(start)
		fmt.Println("result====>", result)
		s := int(result.Hours()) / 24
		fmt.Println("s====>", s)
		fmt.Println("Folad====>", float64(result.Hours())/24)

		employeeTimeOff.NumberOfDays = int64(s)
	} else {
		employeeTimeOff.NumberOfDays = 1
	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.SaveEmployeeTimeOff(ctx, employeeTimeOff)
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

// GetSingleEmployeeTimeOff : ""
func (s *Service) GetSingleEmployeeTimeOff(ctx *models.Context, UniqueID string) (*models.RefEmployeeTimeOff, error) {
	EmployeeTimeOff, err := s.Daos.GetSingleEmployeeTimeOff(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return EmployeeTimeOff, nil
}

// EmployeeTimeOffCount : ""
func (s *Service) EmployeeTimeOffCount(ctx *models.Context, employeeTimeOffCount *models.EmployeeTimeOffCount) (*models.RefEmployeeTimeOffCount, error) {
	EmployeeTimeOff, err := s.Daos.EmployeeTimeOffCount(ctx, employeeTimeOffCount.EmployeeId, employeeTimeOffCount.OrganisationId, employeeTimeOffCount.TimeOffType)
	if err != nil {
		return nil, err
	}
	return EmployeeTimeOff, nil
}

// UpdateEmployeeTimeOff : ""
func (s *Service) UpdateEmployeeTimeOff(ctx *models.Context, employeeTimeOff *models.EmployeeTimeOff) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeTimeOff.Status = constants.EMPLOYEETIMEOFFSTATUSREQUEST
	employeeTimeOff.Revoke = constants.EMPLOYEETIMEOFFSTATUSREQUEST
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeTimeOff.created")
	employeeTimeOff.Created = &created
	log.Println("b4 EmployeeTimeOff.created")
	startDate := employeeTimeOff.StartDate
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 1, startDate.Location())
	//sd := s.Shared.BeginningOfMonth(*employeeTimeOff.StartDate)
	//	ed := s.Shared.EndOfMonth(*employeeTimeOff.StartDate)
	endDate := employeeTimeOff.EndDate
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 23, 59, 59, 0, endDate.Location())

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if endDate != nil {
			// result := end.Day() - (start.Day())
			// if ed.Day() == endDate.Day() {
			// 	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day()-1, 0, 0, 0, 1, startDate.Location())
			// 	result = endDate.Day() - (start.Day())
			// }
			fmt.Println("end====>", end)
			fmt.Println("start====>", start)

			result := end.Sub(start)
			fmt.Println("result====>", result)
			s := int(result.Hours()) / 24
			fmt.Println("s====>", s)
			fmt.Println("Folad====>", float64(result.Hours())/24)

			employeeTimeOff.NumberOfDays = int64(s)
		}
		err := s.Daos.UpdateEmployeeTimeOff(ctx, employeeTimeOff)
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

// EnableEmployeeTimeOff : ""
func (s *Service) EnableEmployeeTimeOff(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeTimeOff(ctx, uniqueID)
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

// DisableEmployeeTimeOff : ""
func (s *Service) DisableEmployeeTimeOff(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeTimeOff(ctx, uniqueID)
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

// DeleteEmployeeTimeOff : ""
func (s *Service) DeleteEmployeeTimeOff(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeTimeOff(ctx, UniqueID)
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

// FilterEmployeeTimeOff : ""
func (s *Service) FilterEmployeeTimeOff(ctx *models.Context, employeeTimeOff *models.FilterEmployeeTimeOff, pagination *models.Pagination) (employeeTimeOffs []models.RefEmployeeTimeOff, err error) {
	err = s.EmployeeTimeOffAccess(ctx, employeeTimeOff)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterEmployeeTimeOff(ctx, employeeTimeOff, pagination)
}
func (s *Service) EmployeeTimeOffAccess(ctx *models.Context, filter *models.FilterEmployeeTimeOff) (err error) {
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

// EmployeeTimeOffRequest : ""
func (s *Service) EmployeeTimeOffRequest(ctx *models.Context, employeeTimeOff *models.EmployeeTimeOff) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeTimeOff.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEETIMEOFF)
	employeeTimeOff.Status = constants.EMPLOYEETIMEOFFSTATUSREQUEST
	t := time.Now()
	employeeTimeOff.RequestDate = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeTimeOffRequest.created")
	employeeTimeOff.Created = &created
	log.Println("b4 EmployeeTimeOffRequest.created")
	startDate := employeeTimeOff.StartDate
	endDate := employeeTimeOff.EndDate
	if endDate != nil {
		result := endDate.Sub(*startDate)
		days := result / (24 * time.Hour)
		day := days.String()
		dayInNumber := day[:len(day)-2]
		dayInNumbers, err := strconv.Atoi(dayInNumber)
		if err != nil {
			return err
		}

		employeeTimeOff.NumberOfDays = int64(dayInNumbers)
	} else {
		employeeTimeOff.NumberOfDays = 1
	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EmployeeTimeOffRequest(ctx, employeeTimeOff)
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

// EmployeeTimeOffApprove : ""
func (s *Service) EmployeeTimeOffApprove(ctx *models.Context, employeeTimeOff *models.ReviewedEmployeeTimeOff) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EmployeeTimeOffApprove(ctx, employeeTimeOff)
		if dberr != nil {
			return dberr
		}

		employeeTimeOff, err := s.Daos.GetSingleEmployeeTimeOff(ctx, employeeTimeOff.EmployeeTimeOff)
		if err != nil {
			return err
		}
		employeeleave := new(models.UpdateEmployeeLeave)
		employeeleave.EmployeeId = employeeTimeOff.EmployeeId
		employeeleave.LeaveType = employeeTimeOff.LeaveType
		employeeleave.Value = int64(employeeTimeOff.NumberOfDays)
		employeeleave.Remarks = employeeTimeOff.Remarks
		err = s.UpdateEmployeeLeaveFromTimeOffWithOutTranscation(ctx, employeeleave)
		if err != nil {
			return err
		}
		date := employeeTimeOff.StartDate
		year, month, day := date.Date()
		var PayRoll string
		fmt.Println("day===>", day)
		for day <= employeeTimeOff.EndDate.Day() {
			if employeeTimeOff.PaidType == "Paid" {
				PayRoll = constants.ATTENDANCESTATUSLEAVEPAID
			}
			if employeeTimeOff.PaidType == "UnPaid" {
				PayRoll = constants.ATTENDANCESTATUSLEAVELOP
			}
			if employeeTimeOff.PaidType == "PatricalPaid" {
				PayRoll = constants.ATTENDANCESTATUSLEAVEPARTIALPAID
			}
			dates := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			strDay := strconv.Itoa(day)
			strMonth := month.String()
			strYear := strconv.Itoa(year)
			uniqueIdcurrentformat := strDay + strMonth + strYear
			Attendance := new(models.Attendance)
			Attendance.UniqueID = uniqueIdcurrentformat
			Attendance.EmployeeId = employeeTimeOff.EmployeeId
			Attendance.CaseLOP = constants.EMPLOYEELOPCASEPLANNED
			Attendance.Date = &dates
			Attendance.LoginMode = constants.ATTENDANCESTATUSTIMEOFF
			Attendance.LeaveType = employeeTimeOff.LeaveType
			Attendance.Status = constants.ATTENDANCESTATUSREQUEST
			Attendance.PayRoll = PayRoll

			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			Attendance.Created = &created
			log.Println("b4 Attendance.created")
			Attendancecheck, err := s.Daos.GetSingleAttendanceByEmployeeIdAndStateValue(ctx, uniqueIdcurrentformat, employeeTimeOff.EmployeeId, constants.ATTENDANCESTATUSLOGIN)
			if err != nil {
				fmt.Println(err)
				return errors.New("Attendance Not Found-Pls Clock in")
			}
			if Attendancecheck != nil {
				Attendance.PunchOut = &t
				RefAttendancelog, err := s.Daos.GetSingleLastAttendanceLogByEmployeeIdAndState(ctx, uniqueIdcurrentformat, employeeTimeOff.EmployeeId, constants.ATTENDANCESTATUSLOGIN)
				if err != nil {
					return err
				}
				diff := Attendance.PunchOut.Sub(*RefAttendancelog.PunchinTime)
				//Attendance.WorkingHours = diff.Minutes()
				Attendance.WorkingHours = Attendance.TotalWorkingMins + diff.Minutes()
				Attendance.TotalWorkingMins = Attendance.WorkingHours
				Attendance.Notes = "Employee Timeoff"
				workschedulehouse := Attendancecheck.WorkscheduleHours * 60
				if workschedulehouse >= Attendance.WorkingHours {
					Attendance.OverTime = 0
				}
				if workschedulehouse < Attendance.WorkingHours {
					Attendance.OverTime = Attendance.WorkingHours - workschedulehouse
				}
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
				AttendanceLog := new(models.AttendanceLog)

				if RefAttendancelog != nil {
					AttendanceLog.UniqueID = RefAttendancelog.UniqueID
					AttendanceLog.AttendanceID = RefAttendancelog.AttendanceID
					AttendanceLog.PunchoutTime = &t
					diff2 := t.Sub(*RefAttendancelog.PunchinTime)
					AttendanceLog.LoginMode = constants.ATTENDANCESTATUSLOGOUT
					AttendanceLog.Status = constants.ATTENDANCESTATUSACTIVE
					AttendanceLog.Date = &t
					AttendanceLog.Notes = "Employee Timeoff"
					//	AttendanceLog.OverTime = Attendance.OverTime
					AttendanceLog.WorkingMins = diff2.Minutes()

					dberr = s.Daos.UpdateAttendanceLog(ctx, AttendanceLog)
					if dberr != nil {
						return dberr
					}
				}
			}
			err = s.Daos.SaveAttendanceWithUpsert(ctx, Attendance)
			if err != nil {
				log.Panicln("Attendance Not Save")
			}
			day++

			fmt.Println("day===>", day)
			fmt.Println("date===>", dates)

		}
		employee, err := s.Daos.GetSingleEmployee(ctx, employeeTimeOff.EmployeeId)
		if err != nil {
			return err
		}

		apptoken, err := s.Daos.GetRegTokenWithParticulars(ctx, employeeTimeOff.EmployeeId)
		if err != nil {
			return err
		}
		if apptoken != nil {
			fmt.Println("apptoken===>", apptoken.RegistrationToken)
			var token []string
			token = append(token, apptoken.RegistrationToken)

			fmt.Println("appToken===>", apptoken.RegistrationToken)
			topic := ""
			tittle := "Employee -" + employeeTimeOff.UniqueID + "TimeOff Approved"
			Body := employeeTimeOff.Description
			//	var image string
			//if len(employeeTimeOff.) > 0 {
			image := ""
			//	}
			data := make(map[string]string)
			data["notificationType"] = "ViewEmployeeTimeOff"
			data["id"] = employeeTimeOff.UniqueID
			err := s.SendNotification(topic, tittle, Body, image, token, data)
			if err != nil {
				log.Println(apptoken.RegistrationToken + " " + err.Error())
			}
			if err == nil {
				t := time.Now()
				ToNotificationLog := new(models.ToNotificationLog)
				notifylog := new(models.NotificationLog)
				ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
				ToNotificationLog.Name = employee.Name
				ToNotificationLog.UserName = employee.UniqueID
				ToNotificationLog.UserType = "Employee"
				notifylog.Body = Body
				notifylog.Tittle = tittle
				notifylog.Topic = topic
				notifylog.Image = image
				notifylog.IsJob = false
				notifylog.Message = Body
				notifylog.SentDate = &t
				notifylog.SentFor = topic
				notifylog.Data = data
				notifylog.Status = "Active"
				notifylog.To = *ToNotificationLog
				err = s.Daos.SaveNotificationLog(ctx, notifylog)
				if err != nil {
					return err
				}
			}
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

// EmployeeTimeOffRevoke : ""
func (s *Service) EmployeeTimeOffRevoke(ctx *models.Context, EmployeeTimeOff *models.ReviewedEmployeeTimeOff) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EmployeeTimeOffRevoke(ctx, EmployeeTimeOff)
		if dberr != nil {
			return dberr
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		employeeTimeOff, err := s.Daos.GetSingleEmployeeTimeOff(ctx, EmployeeTimeOff.EmployeeTimeOff)
		if err != nil {
			return err
		}
		Employee, err := s.Daos.GetSingleEmployee(ctx, employeeTimeOff.EmployeeId)
		if err != nil {
			return err
		}
		employeeleave := new(models.UpdateEmployeeLeave)
		employeeleave.EmployeeId = employeeTimeOff.EmployeeId
		employeeleave.LeaveType = employeeTimeOff.LeaveType
		employeeleave.Value = int64(employeeTimeOff.NumberOfDays)
		employeeleave.Remarks = employeeTimeOff.Remarks
		err = s.RevertEmployeeLeaveFromTimeOffWithOutTranscation(ctx, employeeleave)
		if err != nil {
			return err
		}
		date := employeeTimeOff.StartDate
		year, month, day := date.Date()
		//	var PayRoll string
		Attendance := new(models.Attendance)
		fmt.Println("day===>", day)
		for day <= employeeTimeOff.EndDate.Day() {

			dates := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if Employee.WorkScheduleID != "" {
				WorkSchedule, err := s.Daos.GetSingleWorkSchedule(ctx, Employee.WorkScheduleID)
				if err != nil {
					log.Println(err)
				}
				currentTime := dates
				Day := currentTime.Day()
				Month := currentTime.Month()
				Year := currentTime.Year()
				strDay := strconv.Itoa(Day)
				strMonth := Month.String()
				strYear := strconv.Itoa(Year)
				uniqueIdcurrentformat := strDay + strMonth + strYear
				Attendance.UniqueID = uniqueIdcurrentformat
				Attendance.EmployeeId = employeeTimeOff.EmployeeId
				Attendance.Status = constants.ATTENDANCESTATUSPENDING
				days := currentTime.Weekday().String()
				holiday, err := s.Daos.GetSingleHolidays(ctx, uniqueIdcurrentformat)
				if err != nil {
					log.Panicln("Hoildays--->", err)
				}
				if holiday != nil {
					Attendance.CaseLOP = "Nil"
					Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
					Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
				} else {
					switch days {
					case "Sunday":
						if WorkSchedule.Sunday != true {
							Attendance.CaseLOP = "Nil"
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
						} else {
							Attendance.CaseLOP = constants.EMPLOYEELOPCASEUNPLANNED
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSLOP
							Attendance.WorkscheduleHours = WorkSchedule.WorkingHoursinDay
						}
					case "Monday":
						if WorkSchedule.Monday != true {
							Attendance.CaseLOP = "Nil"
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
						} else {
							Attendance.CaseLOP = constants.EMPLOYEELOPCASEUNPLANNED
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSLOP
							Attendance.WorkscheduleHours = WorkSchedule.WorkingHoursinDay
						}
					case "Tuesday":
						if WorkSchedule.Tuesday != true {
							Attendance.CaseLOP = "Nil"
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
						} else {
							Attendance.CaseLOP = constants.EMPLOYEELOPCASEUNPLANNED
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSLOP
							Attendance.WorkscheduleHours = WorkSchedule.WorkingHoursinDay
						}
					case "Wednesday":
						if WorkSchedule.Wednesday != true {
							Attendance.CaseLOP = "Nil"
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
						} else {
							Attendance.CaseLOP = constants.EMPLOYEELOPCASEUNPLANNED
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSLOP
							Attendance.WorkscheduleHours = WorkSchedule.WorkingHoursinDay
						}
					case "Thursday":
						if WorkSchedule.Thursday != true {
							Attendance.CaseLOP = "Nil"
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
						} else {
							Attendance.CaseLOP = constants.EMPLOYEELOPCASEUNPLANNED
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSLOP
							Attendance.WorkscheduleHours = WorkSchedule.WorkingHoursinDay
						}
					case "Friday":
						if WorkSchedule.Friday != true {
							Attendance.CaseLOP = "Nil"
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
						} else {
							Attendance.CaseLOP = constants.EMPLOYEELOPCASEUNPLANNED
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSLOP
							Attendance.WorkscheduleHours = WorkSchedule.WorkingHoursinDay
						}
					case "Saturday":
						if WorkSchedule.Saturday != true {
							Attendance.CaseLOP = "Nil"
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSHOLIDAY
						} else {
							Attendance.CaseLOP = constants.EMPLOYEELOPCASEUNPLANNED
							Attendance.LoginMode = constants.ATTENDANCESTATUSAUTO
							Attendance.PayRoll = constants.ATTENDANCESTATUSLOP
							Attendance.WorkscheduleHours = WorkSchedule.WorkingHoursinDay
						}
					}
				}
				t := time.Now()
				created := models.Created{}
				created.On = &t
				created.By = constants.SYSTEM
				Attendance.Date = &dates
				Attendance.Created = &created
				log.Println("b4 Attendance.created")

				err = s.Daos.SaveAttendanceWithUpsert(ctx, Attendance)
				if err != nil {
					return errors.New("Attendance Not Save")
				}

			}
			day++

			fmt.Println("day===>", day)
			fmt.Println("date===>", dates)

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

// EmployeeTimeOffReject : ""
func (s *Service) EmployeeTimeOffReject(ctx *models.Context, employeeTimeOff *models.ReviewedEmployeeTimeOff) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// t := time.Now()
		// employeeTimeOff.RejectedDate = &t
		// employeeTimeOff.Status = constants.EMPLOYEETIMEOFFSTATUSREJECT

		err := s.Daos.EmployeeTimeOffReject(ctx, employeeTimeOff)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		employeeTimeOff, err := s.Daos.GetSingleEmployeeTimeOff(ctx, employeeTimeOff.EmployeeTimeOff)
		if err != nil {
			return err
		}
		employee, err := s.Daos.GetSingleEmployee(ctx, employeeTimeOff.EmployeeId)
		if err != nil {
			return err
		}

		apptoken, err := s.Daos.GetRegTokenWithParticulars(ctx, employeeTimeOff.EmployeeId)
		if err != nil {
			return err
		}
		if apptoken != nil {
			fmt.Println("apptoken===>", apptoken.RegistrationToken)
			var token []string
			token = append(token, apptoken.RegistrationToken)

			fmt.Println("appToken===>", apptoken.RegistrationToken)
			topic := ""
			tittle := "Employee -" + employeeTimeOff.UniqueID + "TimeOff Rejected"
			Body := employeeTimeOff.Description
			//	var image string
			//if len(employeeTimeOff.) > 0 {
			image := ""
			//	}
			data := make(map[string]string)
			data["notificationType"] = "ViewEmployeeTimeOff"
			data["id"] = employeeTimeOff.UniqueID
			err := s.SendNotification(topic, tittle, Body, image, token, data)
			if err != nil {
				log.Println(apptoken.RegistrationToken + " " + err.Error())
			}
			if err == nil {
				t := time.Now()
				ToNotificationLog := new(models.ToNotificationLog)
				notifylog := new(models.NotificationLog)
				ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
				ToNotificationLog.Name = employee.Name
				ToNotificationLog.UserName = employee.UniqueID
				ToNotificationLog.UserType = "Employee"
				notifylog.Body = Body
				notifylog.Tittle = tittle
				notifylog.Topic = topic
				notifylog.Image = image
				notifylog.IsJob = false
				notifylog.Message = Body
				notifylog.SentDate = &t
				notifylog.SentFor = topic
				notifylog.Data = data
				notifylog.Status = "Active"
				notifylog.To = *ToNotificationLog
				err = s.Daos.SaveNotificationLog(ctx, notifylog)
				if err != nil {
					return err
				}
			}
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

// CancelEmployeeTimeOff
func (s *Service) CancelEmployeeTimeOff(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.CancelEmployeeTimeOff(ctx, UniqueID)
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

// RevokeRequestEmployeeTimeOff
func (s *Service) RevokeRequestEmployeeTimeOff(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RevokeRequestEmployeeTimeOff(ctx, UniqueID)
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

// EmployeeTimeOffCount : ""
func (s *Service) EmployeeTimeoffDateCheck(ctx *models.Context, employeeTimeOff *models.DayWiseAttendanceReportFilter) error {
	EmployeeTimeOff, err := s.Daos.EmployeeTimeoffDateCheck(ctx, employeeTimeOff)
	if err != nil {
		return err
	}
	if EmployeeTimeOff != nil {
		return errors.New("Employee Timeoff Found")
	}
	return nil
}
