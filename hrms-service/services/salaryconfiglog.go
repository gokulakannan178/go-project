package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSalaryConfigLog : ""
func (s *Service) SaveSalaryConfigLog(ctx *models.Context, salaryConfigLog *models.SalaryConfigLog) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	salaryConfigLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSALARYCONFIGLOG)
	salaryConfigLog.Status = constants.SALARYCONFIGLOGSTATUSACTIVE
	t := time.Now()
	salaryConfigLog.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 SalaryConfigLog.created")
	salaryConfigLog.Created = &created
	log.Println("b4 SalaryConfigLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSalaryConfigLog(ctx, salaryConfigLog)
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

// GetSingleSalaryConfigLog : ""
func (s *Service) GetSingleSalaryConfigLog(ctx *models.Context, UniqueID string) (*models.RefSalaryConfigLog, error) {
	SalaryConfigLog, err := s.Daos.GetSingleSalaryConfigLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return SalaryConfigLog, nil
}

//UpdateSalaryConfigLog : ""
func (s *Service) UpdateSalaryConfigLog(ctx *models.Context, salaryConfigLog *models.SalaryConfigLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateSalaryConfigLog(ctx, salaryConfigLog)
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

// EnableSalaryConfigLog : ""
func (s *Service) EnableSalaryConfigLog(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableSalaryConfigLog(ctx, uniqueID)
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

// DisableSalaryConfigLog : ""
func (s *Service) DisableSalaryConfigLog(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableSalaryConfigLog(ctx, uniqueID)
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

//DeleteSalaryConfigLog : ""
func (s *Service) DeleteSalaryConfigLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSalaryConfigLog(ctx, UniqueID)
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

// FilterSalaryConfigLog : ""
func (s *Service) FilterSalaryConfigLog(ctx *models.Context, salaryConfigLog *models.FilterSalaryConfigLog, pagination *models.Pagination) (salaryConfigLogs []models.RefSalaryConfigLog, err error) {
	return s.Daos.FilterSalaryConfigLog(ctx, salaryConfigLog, pagination)
}
