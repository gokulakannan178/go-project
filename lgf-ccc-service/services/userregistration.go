package service

import (
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

//OTPLoginGenerateOTP :
func (s *Service) UserRegistrationGenerateOTP(ctx *models.Context, login *models.RegistrationUser) error {
	data, err := s.Daos.GetSingleUserWithMobileNo(ctx, login.Mobile, login.Type)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data != nil {
		return errors.New("Mobile number already registered, please login")
	}
	if data == nil {

		otp, err := s.GenerateOTP(constants.REGISTRATIONUSER, login.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
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
func (s *Service) UserRegistrationValidateOTP(ctx *models.Context, login *models.RegistrationUser) (*models.RefUser, bool, error) {
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
			err = s.ValidateOTP(constants.REGISTRATIONUSER, login.Mobile, login.OTP)
			if err != nil {
				fmt.Println(err)
				return err
			}
			//t := time.Now()
			login.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
			login.Status = "Active"
			// if login.Created != nil {
			// 	login.Created.On = &t
			// }
			// if login.Created != nil {
			// 	login.Created.By = constants.SYSTEM
			// }
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
	data, err := s.Daos.GetSingleUser(ctx, login.UserName)
	if err != nil {
		return nil, false, err
	}
	return data, true, nil

}
