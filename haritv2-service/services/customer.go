package services

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveCustomer : ""
func (s *Service) SaveCustomer(ctx *models.Context, Customer *models.Customer) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Customer.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCUSTOMER)
	Customer.Status = constants.CUSTOMERSSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	Customer.Created = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveCustomer(ctx, Customer)
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

//GetSingleCustomer :""
func (s *Service) GetSingleCustomer(ctx *models.Context, UniqueID string) (*models.RefCustomer, error) {
	tower, err := s.Daos.GetSingleCustomer(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateCustomer : ""
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

// EnableCustomer : ""
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

// FilterCustomer : ""
func (s *Service) FilterCustomer(ctx *models.Context, filter *models.CustomerFilter, pagination *models.Pagination) ([]models.RefCustomer, error) {
	return s.Daos.FilterCustomer(ctx, filter, pagination)

}

//OTPLoginGenerateOTP :
func (s *Service) CustomerLoginGenerateOTP(ctx *models.Context, login *models.Login) error {
	data, err := s.Daos.GetSingleCustomerwithmobileno(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data == nil {
		return errors.New("customer Not Available")
	}
	if data.Status == constants.CUSTOMERSSTATUSINIT {
		return errors.New("Awaiting Activation")
	}
	if data.Status != constants.CUSTOMERSSTATUSACTIVE {
		return errors.New("Please Contact Administrator")
	}

	otp, err := s.GenerateOTP(constants.CUSTOMERLOGIN, login.UserName, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return errors.New("Otp Generate Error - " + err.Error())
	}
	text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", data.Name, otp)
	err = s.SendSMS(login.UserName, text)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}

	return nil
}

//OTPLoginValidateOTP :
func (s *Service) CustomerLoginValidateOTP(ctx *models.Context, login *models.OTPLogin) (*models.RefCustomer, bool, error) {

	data, err := s.Daos.GetSingleCustomerwithmobileno(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	if data == nil {
		return nil, false, errors.New("customer Not Available")
	}
	if data.Status == constants.CUSTOMERSSTATUSINIT {
		return nil, false, errors.New("Awaiting Activation")
	}
	if data.Status != constants.CUSTOMERSSTATUSACTIVE {
		return nil, false, errors.New("Please Contact Administrator")
	}

	err = s.ValidateOTP(constants.CUSTOMERLOGIN, login.Mobile, login.OTP)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}

	return data, true, nil

}

//GetSingleCustomer :""
func (s *Service) GetSingleCustomerwithprofile(ctx *models.Context, Profile string) (*models.RefCustomer, error) {
	tower, err := s.Daos.GetSingleCustomerwithprofile(ctx, Profile)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateCustomer : ""
func (s *Service) UpdateCustomerwithprofile(ctx *models.Context, Customer *models.Customer) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCustomerwithprofile(ctx, Customer)
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
