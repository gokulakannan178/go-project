package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeShift : ""
func (s *Service) SaveEmployeeShift(ctx *models.Context, employeeshift *models.EmployeeShift) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	//EmployeeShift.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEmployeeShift)
	employeeshift.Status = constants.EMPLOYEESHIFTSTATUSACTIVE
	t := time.Now()
	//	EmployeeShift.RegisterDate = &t
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeShift.created")
	//EmployeeShift.Created = &created
	log.Println("b4 EmployeeShift.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeShift(ctx, employeeshift)
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

// GetSingleEmployeeShift : ""
func (s *Service) GetSingleEmployeeShift(ctx *models.Context, UniqueID string) (*models.RefEmployeeShift, error) {
	employeeshift, err := s.Daos.GetSingleEmployeeShift(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeshift, nil
}

// func (s *Service) GetSingleEmployeeShiftWithHoldingNumber(ctx *models.Context, holdingNumber string) (*models.RefEmployeeShift, error) {
// 	EmployeeShift, err := s.Daos.GetSingleEmployeeShiftWithHoldingNumber(ctx, holdingNumber)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return EmployeeShift, nil
// }

// UpdateEmployeeShift : ""
func (s *Service) UpdateEmployeeShift(ctx *models.Context, employeeshift *models.EmployeeShift) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeShift(ctx, employeeshift)
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

// EnableEmployeeShift : ""
func (s *Service) EnableEmployeeShift(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeShift(ctx, uniqueID)
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

// DisableEmployeeShift : ""
func (s *Service) DisableEmployeeShift(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeShift(ctx, uniqueID)
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

// DeleteEmployeeShift : ""
func (s *Service) DeleteEmployeeShift(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeShift(ctx, UniqueID)
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

// FilterEmployeeShift : ""
func (s *Service) FilterEmployeeShift(ctx *models.Context, filter *models.FilterEmployeeShift, pagination *models.Pagination) (EmployeeShifts []models.RefEmployeeShift, err error) {
	return s.Daos.FilterEmployeeShift(ctx, filter, pagination)
}
