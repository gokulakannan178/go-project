package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeDeductionMaster : ""
func (s *Service) SaveEmployeeDeductionMaster(ctx *models.Context, employeeDeductionMaster *models.EmployeeDeductionMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeDeductionMaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER)
	employeeDeductionMaster.Status = constants.EMPLOYEEDEDUCTIONMASTERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeDeductionMaster.created")
	employeeDeductionMaster.Created = &created
	log.Println("b4 EmployeeDeductionMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeDeductionMaster(ctx, employeeDeductionMaster)
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

// GetSingleEmployeeDeductionMaster : ""
func (s *Service) GetSingleEmployeeDeductionMaster(ctx *models.Context, UniqueID string) (*models.RefEmployeeDeductionMaster, error) {
	employeeDeductionMaster, err := s.Daos.GetSingleEmployeeDeductionMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeDeductionMaster, nil
}

//UpdateEmployeeDeductionMaster : ""
func (s *Service) UpdateEmployeeDeductionMaster(ctx *models.Context, employeeDeductionMaster *models.EmployeeDeductionMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeDeductionMaster(ctx, employeeDeductionMaster)
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

// EnableEmployeeDeductionMaster : ""
func (s *Service) EnableEmployeeDeductionMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeDeductionMaster(ctx, uniqueID)
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

// DisableEmployeeDeductionMaster : ""
func (s *Service) DisableEmployeeDeductionMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeDeductionMaster(ctx, uniqueID)
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

//DeleteEmployeeDeductionMaster : ""
func (s *Service) DeleteEmployeeDeductionMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeDeductionMaster(ctx, UniqueID)
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

// FilterEmployeeDeductionMaster : ""
func (s *Service) FilterEmployeeDeductionMaster(ctx *models.Context, employeeDeductionMaster *models.FilterEmployeeDeductionMaster, pagination *models.Pagination) (employeeDeductionMasters []models.RefEmployeeDeductionMaster, err error) {
	return s.Daos.FilterEmployeeDeductionMaster(ctx, employeeDeductionMaster, pagination)
}
