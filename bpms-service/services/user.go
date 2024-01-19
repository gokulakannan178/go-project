package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveUser(ctx *models.Context, user *models.User) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	user.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDUSER)
	user.Status = constants.USERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	user.Created = created
	user.Password = constants.USERSDEFAULTPASSWORD
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUser(ctx, user)
		if dberr != nil {
			return dberr
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

//UpdateUser : ""
func (s *Service) UpdateUser(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUser(ctx, user)
		if err != nil {
			return err
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

//EnableUser : ""
func (s *Service) EnableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUser(ctx, UniqueID)
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

//DisableUser : ""
func (s *Service) DisableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUser(ctx, UniqueID)
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

//DeleteUser : ""
func (s *Service) DeleteUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUser(ctx, UniqueID)
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

//GetSingleUser :""
func (s *Service) GetSingleUser(ctx *models.Context, UniqueID string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		fmt.Println("checking user types")
		switch user.Type {

		case constants.USERTYPEULB:
			fmt.Println("in ulb")
			user.Ref.ULB, err = s.Daos.GetSingleULB(ctx, user.OrganisationID)
			if err != nil {
				fmt.Println("error in finding ulb - " + err.Error())
			}

		case constants.USERTYPESUPERADMIN:
			fmt.Println("in superadmin")
		case constants.USERTYPEDEPARTMENT:
			fmt.Println("in dept")
			user.Ref.Department, err = s.Daos.GetSingleDepartment(ctx, user.OrganisationID)
			if err != nil {
				fmt.Println("error in finding dept - " + err.Error())
			}
		case constants.USERTYPEAPPLICANT:
			fmt.Println("in applicant")
			user.Ref.Applicant, err = s.Daos.GetSingleApplicant(ctx, user.UserName)
			if err != nil {
				fmt.Println("error in finding applicant - " + err.Error())
			}
		}
	}
	return user, nil
}

//FilterUser :""
func (s *Service) FilterUser(ctx *models.Context, userfilter *models.UserFilter, pagination *models.Pagination) (user []models.RefUser, err error) {
	return s.Daos.FilterUser(ctx, userfilter, pagination)
}
