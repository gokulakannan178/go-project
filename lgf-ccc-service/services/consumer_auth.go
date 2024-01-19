package service

import (
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SendOTPConsumerLogin : ""
func (s *Service) SendOTPConsumerLogin(ctx *models.Context, mobileNo string) error {

	token, _ := s.GenerateOTP(constants.CONSUMERLOGIN, mobileNo, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return err
	// }
	fmt.Println("token =>", token)
	msg := fmt.Sprintf("Your OTP for lgf-ccc Consumer login is %v. This OTP is valid only for 3 minutes. Please do not share OTP to anyone",
		token)
	if err := s.SendSMS(mobileNo, msg); err != nil {
		log.Println(err)
	}
	return nil
}

//ConsumerLoginValidateOTP : ""
func (s *Service) ConsumerLoginValidateOTP(ctx *models.Context, mobileNo string, otp string) ([]string, error) {

	return nil, s.ValidateOTP(constants.CONSUMERLOGIN, mobileNo, otp)
}

func (s *Service) CitizenregistrationGenerateOTP(ctx *models.Context, login *models.RegistrationUser) error {
	data, err := s.Daos.GetSingleUserWithMobileNo(ctx, login.Mobile, login.Type)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data != nil {
		return errors.New("Mobile number already registered, please login")
	}
	if data == nil {

		otp, err := s.GenerateOTP(constants.USERREGISTRATION, login.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
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
func (s *Service) CitizenregistrationValidateOTP(ctx *models.Context, login *models.RegistrationUser) (*models.RefUser, bool, error) {
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, false, err
	}

	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		data, err := s.Daos.GetSingleUserWithMobileNo(ctx, login.Mobile, login.Type)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if data != nil {
			return errors.New("Mobile number already registered, please login")
		}
		if data == nil {
			err = s.ValidateOTP(constants.USERREGISTRATION, login.Mobile, login.OTP)
			if err != nil {
				fmt.Println(err)
				return err
			}
			t := time.Now()
			login.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
			login.Status = "Active"
			login.Type = "Citizen"
			if login.Created.On != nil {
				login.Created.On = &t
			}
			if login.Created.By != "" {
				login.Created.By = constants.SYSTEM
			}
			reg := s.Daos.SaveUser(ctx, &login.User)
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
	data, err := s.Daos.GetSingleUser(ctx, login.UniqueID)
	if err != nil {
		return nil, false, err
	}
	return data, true, nil

}
