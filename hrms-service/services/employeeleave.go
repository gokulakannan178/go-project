package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeLeave : ""
func (s *Service) SaveEmployeeLeave(ctx *models.Context, employeeLeave *models.EmployeeLeave) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeLeave.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVE)
	employeeLeave.Status = constants.EMPLOYEELEAVESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeLeave.created")
	employeeLeave.Created = &created
	log.Println("b4 EmployeeLeave.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeLeave(ctx, employeeLeave)
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
func (s *Service) SaveEmployeeLeaveWithLog(ctx *models.Context, employeeLeave *models.EmployeeLeave) error {

	employeeLeave.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVE)
	employeeLeave.Status = constants.EMPLOYEELEAVESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeLeave.created")
	employeeLeave.Created = &created
	log.Println("b4 EmployeeLeave.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeLeave(ctx, employeeLeave)
		if dberr != nil {
			return dberr
		}
		employeeleavelog := new(models.EmployeeLeaveLog)
		employeeleavelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVELOG)
		employeeleavelog.EmployeeId = employeeLeave.EmployeeId
		employeeleavelog.LeaveType = employeeLeave.LeaveType
		employeeleavelog.Value = employeeLeave.Value
		employeeleavelog.Date = &t
		employeeleavelog.CreateDate = &t
		employeeleavelog.Status = constants.EMPLOYEELEAVESTATUSACTIVE
		employeeleavelog.Created = &created
		err := s.Daos.SaveEmployeeLeaveLog(ctx, employeeleavelog)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		return err
	}
	return nil
}

// GetSingleEmployeeLeave : ""
func (s *Service) GetSingleEmployeeLeave(ctx *models.Context, UniqueID string) (*models.RefEmployeeLeave, error) {
	employeeLeave, err := s.Daos.GetSingleEmployeeLeave(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeLeave, nil
}

//UpdateEmployeeLeave : ""
func (s *Service) UpdateEmployeeLeave(ctx *models.Context, employeeLeave *models.EmployeeLeave) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeLeave(ctx, employeeLeave)
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

// EnableEmployeeLeave : ""
func (s *Service) EnableEmployeeLeave(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeLeave(ctx, uniqueID)
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

// DisableEmployeeLeave : ""
func (s *Service) DisableEmployeeLeave(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeLeave(ctx, uniqueID)
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

//DeleteEmployeeLeave : ""
func (s *Service) DeleteEmployeeLeave(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeLeave(ctx, UniqueID)
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

// FilterEmployeeLeave : ""
func (s *Service) FilterEmployeeLeave(ctx *models.Context, employeeLeave *models.FilterEmployeeLeave, pagination *models.Pagination) (employeeLeaves []models.RefEmployeeLeave, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterEmployeeLeave(ctx, employeeLeave, pagination)
}

// GetEmployeeLeaveCount : ""
func (s *Service) GetEmployeeLeaveCount(ctx *models.Context, employeeleavecount *models.EmployeeLeaveCount) (employeeleavecounts []models.RefEmployeeLeaveCount, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.GetEmployeeLeaveCount(ctx, employeeleavecount)
}
func (s *Service) UpdateEmployeeLeaveFromTimeOff(ctx *models.Context, employeeLeave *models.UpdateEmployeeLeave) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.UpdateEmployeeLeaveFromTimeOffWithOutTranscation(ctx, employeeLeave)
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
func (s *Service) UpdateEmployeeLeaveFromTimeOffWithOutTranscation(ctx *models.Context, employeeLeave *models.UpdateEmployeeLeave) error {

	_, err := s.Daos.GetSingleEmployeeLeaveWithEmployeeId(ctx, employeeLeave.EmployeeId)
	if err != nil {
		return err
	}
	dberr := s.Daos.UpdateEmployeeLeaveFromTimeOff(ctx, employeeLeave)
	if dberr != nil {
		return dberr
	}
	employeeLeavelog := new(models.EmployeeLeaveLog)
	employeeLeavelog.EmployeeId = employeeLeave.EmployeeId
	employeeLeavelog.LeaveType = employeeLeave.LeaveType
	employeeLeavelog.Value = employeeLeave.Value
	employeeLeavelog.Remarks = employeeLeave.Remarks
	err = s.SaveEmployeeLeaveLog(ctx, employeeLeavelog)
	if err != nil {
		return err
	}
	return nil

}
func (s *Service) RevertEmployeeLeaveFromTimeOffWithOutTranscation(ctx *models.Context, employeeLeave *models.UpdateEmployeeLeave) error {

	_, err := s.Daos.GetSingleEmployeeLeaveWithEmployeeId(ctx, employeeLeave.EmployeeId)
	if err != nil {
		return err
	}
	dberr := s.Daos.RevertEmployeeLeaveFromTimeOff(ctx, employeeLeave)
	if dberr != nil {
		return dberr
	}
	employeeLeavelog := new(models.EmployeeLeaveLog)
	employeeLeavelog.EmployeeId = employeeLeave.EmployeeId
	employeeLeavelog.LeaveType = employeeLeave.LeaveType
	employeeLeavelog.Value = employeeLeave.Value
	employeeLeavelog.Remarks = employeeLeave.Remarks
	employeeLeavelog.Revert = true
	err = s.SaveEmployeeLeaveLog(ctx, employeeLeavelog)
	if err != nil {
		return err
	}
	return nil

}
func (s *Service) EmployeeleaveList(ctx *models.Context, filter *models.FilterEmployeeLeave) ([]models.EmployeeLeaveListV2, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.EmployeeLeaveList(ctx, filter)

}
func (s *Service) EmployeeLeaveListV2(ctx *models.Context, filter *models.FilterEmployeeLeaveList) (*models.EmployeeLeaveList, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	employeeleave, err := s.Daos.EmployeeLeaveListV3(ctx, filter)
	if err != nil {
		return nil, err
	}
	if employeeleave != nil {
		for k, v := range employeeleave.EmployeeLeave {
			//	fmt.Println("v.NumberOfDays===>", v.NumberOfDays)

			if v.NumberOfDays > 0 {
				///.Println("beforeemployeeleave.EmployeeLeave[k].Value===>", employeeleave.EmployeeLeave[k].Value)
				//
				employeeleave.EmployeeLeave[k].Value = employeeleave.EmployeeLeave[k].Value - v.NumberOfDays
				employeeleave.EmployeeLeave[k].NumberOfDays = 0
				//fmt.Println("afteremployeeleave.EmployeeLeave[k].Value===>", employeeleave.EmployeeLeave[k].Value)
			}
		}
	}
	return employeeleave, nil
}
func (s *Service) GetAllEmployeeLeaveList(ctx *models.Context, filter *models.FilterEmployeeLeaveList) ([]models.EmployeeLeaveList, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	employeeleave, err := s.Daos.GetAllEmployeeLeaveList(ctx, filter)
	if err != nil {
		return nil, err
	}
	for _, v2 := range employeeleave {
		for k, v := range v2.EmployeeLeave {
			//	fmt.Println("v.NumberOfDays===>", v.NumberOfDays)

			if v.NumberOfDays > 0 {
				///.Println("beforeemployeeleave.EmployeeLeave[k].Value===>", employeeleave.EmployeeLeave[k].Value)
				//
				v2.EmployeeLeave[k].Value = v2.EmployeeLeave[k].Value - v.NumberOfDays
				v2.EmployeeLeave[k].NumberOfDays = 0
				//fmt.Println("afteremployeeleave.EmployeeLeave[k].Value===>", employeeleave.EmployeeLeave[k].Value)
			}
		}
	}

	return employeeleave, nil
}
func (s *Service) UpdateEmployeeLeaveWithEmployeeId(ctx *models.Context, employeeLeave *models.EmployeeLeaveLog) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.UpdateEmployeeLeaveWithEmployeeId(ctx, employeeLeave.EmployeeId, employeeLeave.LeaveType, employeeLeave.Value)
		if dberr != nil {
			return dberr
		}
		employeeLeave.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVELOG)
		employeeLeave.Status = constants.EMPLOYEELEAVELOGSTATUSACTIVE
		t := time.Now()
		employeeLeave.CreateDate = &t
		employeeLeave.Date = &t
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		log.Println("b4 EmployeeLeaveLog.created")
		employeeLeave.Created = &created
		log.Println("b4 EmployeeLeaveLog.created")
		err := s.Daos.SaveEmployeeLeaveLog(ctx, employeeLeave)
		if err != nil {
			return err
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
