package services

import (
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeOnboardingCheckList : ""
func (s *Service) SaveEmployeeOnboardingCheckList(ctx *models.Context, employeeonboardingchecklist *models.EmployeeOnboardingCheckList) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeonboardingchecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST)
	employeeonboardingchecklist.Status = constants.EMPLOYEEONBOARDINGCHECKLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeOnboardingCheckList.created")
	employeeonboardingchecklist.Created = &created
	log.Println("b4 EmployeeOnboardingCheckList.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeOnboardingCheckList(ctx, employeeonboardingchecklist)
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

// GetSingleEmployeeOnboardingCheckList : ""
func (s *Service) GetSingleEmployeeOnboardingCheckList(ctx *models.Context, UniqueID string) (*models.RefEmployeeOnboardingCheckList, error) {
	employeeonboardingchecklist, err := s.Daos.GetSingleEmployeeOnboardingCheckList(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeonboardingchecklist, nil
}

// EmployeeOnboardingCheckListFinal : ""
func (s *Service) EmployeeOnboardingCheckListFinal(ctx *models.Context, EmployeeID string) (*models.RefEmployeeOnboardingCheckListv2, error) {
	employee, err := s.Daos.GetSingleEmployee(ctx, EmployeeID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	employeeonboardingchecklist, err := s.Daos.EmployeeOnboardingCheckListFinal(ctx, EmployeeID, employee.OnboardingpolicyID)
	if err != nil {
		return nil, err
	}

	return employeeonboardingchecklist, nil
}

//UpdateEmployeeOnboardingCheckList : ""
func (s *Service) UpdateEmployeeOnboardingCheckList(ctx *models.Context, employeeonboardingchecklist *models.EmployeeOnboardingCheckList) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeOnboardingCheckList(ctx, employeeonboardingchecklist)
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

// EnableEmployeeOnboardingCheckList : ""
func (s *Service) EnableEmployeeOnboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeOnboardingCheckList(ctx, uniqueID)
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

// DisableEmployeeOnboardingCheckList : ""
func (s *Service) DisableEmployeeOnboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeOnboardingCheckList(ctx, uniqueID)
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

//DeleteEmployeeOnboardingCheckList : ""
func (s *Service) DeleteEmployeeOnboardingCheckList(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeOnboardingCheckList(ctx, UniqueID)
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

// FilterEmployeeOnboardingCheckList : ""
func (s *Service) FilterEmployeeOnboardingCheckList(ctx *models.Context, employeeonboardingchecklist *models.FilterEmployeeOnboardingCheckList, pagination *models.Pagination) (employeeonboardingchecklists []models.RefEmployeeOnboardingCheckList, err error) {
	return s.Daos.FilterEmployeeOnboardingCheckList(ctx, employeeonboardingchecklist, pagination)
}
