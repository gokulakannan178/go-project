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

//OTPLoginGenerateOTP :
func (s *Service) CustomerregistrationGenerateOTP(ctx *models.Context, login *models.RegistrationCustomer) error {
	data, err := s.Daos.GetSingleCustomerwithmobileno(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data != nil {
		return errors.New("Mobile number already registered, please login")
	}
	if data == nil {

		otp, err := s.GenerateOTP(constants.CUSTOMERREGISTRATION, login.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
		if err != nil {
			return errors.New("Otp Generate Error - " + err.Error())
		}

		text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", login.Name, otp)
		err = s.SendSMS(login.Mobile, text)
		if err != nil {
			return errors.New("Sms Sending Error - " + err.Error())
		}
	}
	return nil
}

//OTPLoginValidateOTP :
func (s *Service) CustomerregistrationValidateOTP(ctx *models.Context, login *models.RegistrationCustomer) (*models.RefCustomer, bool, error) {
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, false, err
	}

	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		data, err := s.Daos.GetSingleCustomerwithmobileno(ctx, login.Mobile)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if data != nil {
			return errors.New("Mobile number already registered, please login")
		}
		if data == nil {
			err = s.ValidateOTP(constants.CUSTOMERREGISTRATION, login.Mobile, login.OTP)
			if err != nil {
				fmt.Println(err)
				return err
			}
			t := time.Now()
			login.Customer.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCUSTOMER)
			login.Customer.Status = constants.CUSTOMERSSTATUSACTIVE
			login.Customer.Created.On = &t
			login.Customer.Created.By = constants.SYSTEM

			reg := s.Daos.SaveCustomer(ctx, &login.Customer)
			if reg != nil {
				fmt.Println(reg)
				return reg

			}

		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return nil, false, errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return nil, false, err
	}
	data, err := s.Daos.GetSingleCustomer(ctx, login.UniqueID)
	if err != nil {
		return nil, false, err
	}
	return data, true, nil

}
