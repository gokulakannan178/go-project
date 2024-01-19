package services

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveCustomer :""
func (s *Service) SaveCustomer(ctx *models.Context, Customer *models.Customer) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Customer.Status = constants.CUSTOMERSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		tempUser, dberr := s.Daos.GetSingleCustomerWithCondition(ctx, "mobile", Customer.Mobile)
		if dberr != nil {
			return errors.New("Db Error" + dberr.Error())
		}
		if tempUser != nil {
			return errors.New("mobile no already in use")
		}
		tempUser, dberr = s.Daos.GetSingleCustomerWithCondition(ctx, "email", Customer.Mobile)
		if dberr != nil {
			return errors.New("Db Error" + dberr.Error())
		}
		if tempUser != nil {
			return errors.New("email already in use")
		}
		dberr = s.Daos.SaveCustomer(ctx, Customer)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
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
func (s *Service) SaveCustomerWithOutTransaction(ctx *models.Context, Customer *models.Customer) error {
	Customer.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCUSTOMER)
	Customer.Status = constants.CUSTOMERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Customer.created")
	Customer.Created = &created
	log.Println("b4 Customer.created")
	dberr := s.Daos.SaveCustomer(ctx, Customer)
	return dberr
}

//UpdateCustomer : ""
func (s *Service) UpdateCustomer(ctx *models.Context, Customer *models.Customer) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCustomer(ctx, Customer)
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

//EnableCustomer : ""
func (s *Service) EnableCustomer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableCustomer(ctx, UniqueID)
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

//DisableCustomer : ""
func (s *Service) DisableCustomer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableCustomer(ctx, UniqueID)
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

//DeleteCustomer : ""
func (s *Service) DeleteCustomer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCustomer(ctx, UniqueID)
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

//GetSingleCustomer :""
func (s *Service) GetSingleCustomer(ctx *models.Context, UniqueID string) (*models.RefCustomer, error) {
	Customer, err := s.Daos.GetSingleCustomer(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Customer, nil
}

//GetSingleGetUsingMobileNumber :""
func (s *Service) GetSingleGetUsingMobileNumber(ctx *models.Context, Mobile string) (*models.RefCustomer, error) {
	Customer, err := s.Daos.GetSingleGetUsingMobileNumber(ctx, Mobile)
	if err != nil {
		return nil, err
	}
	return Customer, nil
}

//FilterCustomer :""
func (s *Service) FilterCustomer(ctx *models.Context, Customerfilter *models.CustomerFilter, pagination *models.Pagination) (Customer []models.RefCustomer, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterCustomer(ctx, Customerfilter, pagination)

}
