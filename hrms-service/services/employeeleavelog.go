package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeLeaveLog : ""
func (s *Service) SaveEmployeeLeaveLog(ctx *models.Context, employeeLeaveLog *models.EmployeeLeaveLog) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeLeaveLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVELOG)
	employeeLeaveLog.Status = constants.EMPLOYEELEAVELOGSTATUSACTIVE
	t := time.Now()
	employeeLeaveLog.CreateDate = &t
	employeeLeaveLog.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeLeaveLog.created")
	employeeLeaveLog.Created = &created
	log.Println("b4 EmployeeLeaveLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeLeaveLog(ctx, employeeLeaveLog)
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

// GetSingleEmployeeLeaveLog : ""
func (s *Service) GetSingleEmployeeLeaveLog(ctx *models.Context, UniqueID string) (*models.RefEmployeeLeaveLog, error) {
	EmployeeLeaveLog, err := s.Daos.GetSingleEmployeeLeaveLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return EmployeeLeaveLog, nil
}

// EmployeeLeaveLogCount : ""
func (s *Service) EmployeeLeaveLogCount(ctx *models.Context, EmployeeLeaveLogCount *models.EmployeeLeaveLogCount) (*models.RefEmployeeLeaveLogCount, error) {
	EmployeeLeaveLog, err := s.Daos.EmployeeLeaveLogCount(ctx, EmployeeLeaveLogCount.EmployeeId, EmployeeLeaveLogCount.OrganisationId, EmployeeLeaveLogCount.LeaveType)
	if err != nil {
		return nil, err
	}
	return EmployeeLeaveLog, nil
}

//UpdateEmployeeLeaveLog : ""
func (s *Service) UpdateEmployeeLeaveLog(ctx *models.Context, EmployeeLeaveLog *models.EmployeeLeaveLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeLeaveLog(ctx, EmployeeLeaveLog)
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

// EnableEmployeeLeaveLog : ""
func (s *Service) EnableEmployeeLeaveLog(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeLeaveLog(ctx, uniqueID)
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

// DisableEmployeeLeaveLog : ""
func (s *Service) DisableEmployeeLeaveLog(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeLeaveLog(ctx, uniqueID)
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

//DeleteEmployeeLeaveLog : ""
func (s *Service) DeleteEmployeeLeaveLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeLeaveLog(ctx, UniqueID)
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

// FilterEmployeeLeaveLog : ""
func (s *Service) FilterEmployeeLeaveLog(ctx *models.Context, employeeLeaveLog *models.FilterEmployeeLeaveLog, pagination *models.Pagination) (employeeLeaveLogs []models.RefEmployeeLeaveLog, err error) {
	return s.Daos.FilterEmployeeLeaveLog(ctx, employeeLeaveLog, pagination)
}
