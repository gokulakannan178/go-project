package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeFamilyMembers : ""
func (s *Service) SaveEmployeeFamilyMembers(ctx *models.Context, employeeFamilyMembers *models.EmployeeFamilyMembers) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeFamilyMembers.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEFAMILYMEMBERS)
	employeeFamilyMembers.Status = constants.EMPLOYEEFAMILYMEMBERSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeFamilyMembers.created")
	employeeFamilyMembers.Created = &created
	log.Println("b4 EmployeeFamilyMembers.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeFamilyMembers(ctx, employeeFamilyMembers)
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

// GetSingleEmployeeFamilyMembers : ""
func (s *Service) GetSingleEmployeeFamilyMembers(ctx *models.Context, UniqueID string) (*models.RefEmployeeFamilyMembers, error) {
	EmployeeFamilyMembers, err := s.Daos.GetSingleEmployeeFamilyMembers(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return EmployeeFamilyMembers, nil
}

//UpdateEmployeeFamilyMembers : ""
func (s *Service) UpdateEmployeeFamilyMembers(ctx *models.Context, employeeFamilyMembers *models.EmployeeFamilyMembers) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeFamilyMembers(ctx, employeeFamilyMembers)
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

// EnableEmployeeFamilyMembers : ""
func (s *Service) EnableEmployeeFamilyMembers(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeFamilyMembers(ctx, uniqueID)
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

// DisableEmployeeFamilyMembers : ""
func (s *Service) DisableEmployeeFamilyMembers(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeFamilyMembers(ctx, uniqueID)
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

//DeleteEmployeeFamilyMembers : ""
func (s *Service) DeleteEmployeeFamilyMembers(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeFamilyMembers(ctx, UniqueID)
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

// FilterEmployeeFamilyMembers : ""
func (s *Service) FilterEmployeeFamilyMembers(ctx *models.Context, employeeFamilyMembers *models.FilterEmployeeFamilyMembers, pagination *models.Pagination) (employeeFamilyMemberss []models.RefEmployeeFamilyMembers, err error) {
	return s.Daos.FilterEmployeeFamilyMembers(ctx, employeeFamilyMembers, pagination)
}
