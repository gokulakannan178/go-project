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

// SaveEmployeeOffboardingCheckList : ""
func (s *Service) SaveEmployeeOffboardingCheckList(ctx *models.Context, employeeoffboardingchecklist *models.EmployeeOffboardingCheckList) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeoffboardingchecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST)
	employeeoffboardingchecklist.Status = constants.EMPLOYEEOFFBOARDINGCHECKLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeOffboardingCheckList.created")
	employeeoffboardingchecklist.Created = &created
	log.Println("b4 EmployeeOffboardingCheckList.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeOffboardingCheckList(ctx, employeeoffboardingchecklist)
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

// GetSingleEmployeeOffboardingCheckList : ""
func (s *Service) GetSingleEmployeeOffboardingCheckList(ctx *models.Context, UniqueID string) (*models.RefEmployeeOffboardingCheckList, error) {
	employeeoffboardingchecklist, err := s.Daos.GetSingleEmployeeOffboardingCheckList(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeoffboardingchecklist, nil
}

// EmployeeOffboardingCheckListFinal : ""
func (s *Service) EmployeeOffboardingCheckListFinal(ctx *models.Context, EmployeeID string) (*models.RefEmployeeOffboardingCheckListv2, error) {
	employee, err := s.Daos.GetSingleEmployee(ctx, EmployeeID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	employeeoffboardingchecklist, err := s.Daos.EmployeeOffboardingCheckListFinal(ctx, EmployeeID, employee.OffboardingpolicyID)
	if err != nil {
		return nil, err
	}

	return employeeoffboardingchecklist, nil
}

//UpdateEmployeeOffboardingCheckList : ""
func (s *Service) UpdateEmployeeOffboardingCheckList(ctx *models.Context, employeeoffboardingchecklist *models.EmployeeOffboardingCheckList) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeOffboardingCheckList(ctx, employeeoffboardingchecklist)
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

// EnableEmployeeOffboardingCheckList : ""
func (s *Service) EnableEmployeeOffboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeOffboardingCheckList(ctx, uniqueID)
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

// DisableEmployeeOffboardingCheckList : ""
func (s *Service) DisableEmployeeOffboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeOffboardingCheckList(ctx, uniqueID)
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

//DeleteEmployeeOffboardingCheckList : ""
func (s *Service) DeleteEmployeeOffboardingCheckList(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeOffboardingCheckList(ctx, UniqueID)
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

// FilterEmployeeOffboardingCheckList : ""
func (s *Service) FilterEmployeeOffboardingCheckList(ctx *models.Context, employeeoffboardingchecklist *models.FilterEmployeeOffboardingCheckList, pagination *models.Pagination) (employeeoffboardingchecklists []models.RefEmployeeOffboardingCheckList, err error) {
	return s.Daos.FilterEmployeeOffboardingCheckList(ctx, employeeoffboardingchecklist, pagination)
}
