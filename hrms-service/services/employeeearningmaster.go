package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeEarningMaster : ""
func (s *Service) SaveEmployeeEarningMaster(ctx *models.Context, employeeEarningMaster *models.EmployeeEarningMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeEarningMaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEEARNINGMASTER)
	employeeEarningMaster.Status = constants.EMPLOYEEEARNINGMASTERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeEarningMaster.created")
	employeeEarningMaster.Created = &created
	log.Println("b4 EmployeeEarningMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeEarningMaster(ctx, employeeEarningMaster)
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

// GetSingleEmployeeEarningMaster : ""
func (s *Service) GetSingleEmployeeEarningMaster(ctx *models.Context, UniqueID string) (*models.RefEmployeeEarningMaster, error) {
	employeeEarningMaster, err := s.Daos.GetSingleEmployeeEarningMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeEarningMaster, nil
}

//UpdateEmployeeEarningMaster : ""
func (s *Service) UpdateEmployeeEarningMaster(ctx *models.Context, employeeEarningMaster *models.EmployeeEarningMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeEarningMaster(ctx, employeeEarningMaster)
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

// EnableEmployeeEarningMaster : ""
func (s *Service) EnableEmployeeEarningMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeEarningMaster(ctx, uniqueID)
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

// DisableEmployeeEarningMaster : ""
func (s *Service) DisableEmployeeEarningMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeEarningMaster(ctx, uniqueID)
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

//DeleteEmployeeEarningMaster : ""
func (s *Service) DeleteEmployeeEarningMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeEarningMaster(ctx, UniqueID)
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

// FilterEmployeeEarningMaster : ""
func (s *Service) FilterEmployeeEarningMaster(ctx *models.Context, employeeEarningMaster *models.FilterEmployeeEarningMaster, pagination *models.Pagination) (employeeEarningMasters []models.RefEmployeeEarningMaster, err error) {
	return s.Daos.FilterEmployeeEarningMaster(ctx, employeeEarningMaster, pagination)
}
