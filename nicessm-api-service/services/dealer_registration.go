package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//OTPLoginGenerateOTP :
func (s *Service) DealerregistrationGenerateOTP(ctx *models.Context, login *models.RegistrationDealer) error {
	data, err := s.Daos.GetSingleDealerWithMobileNo(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data != nil {
		return errors.New("Dealer already Available")
	}
	if data == nil {

		//otp, _ := s.GenerateOTP(constants.DEALERREGISTRATION, login.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
		// if err != nil {
		// 	return errors.New("Otp Generate Error - " + err.Error())
		// }
		key := fmt.Sprintf("%v_%v", constants.DEALERREGISTRATION, login.Mobile)
		var otp models.Otp
		otp.Otp = "9999"
		err = s.SetValueCacheMemory(key, otp, 1000)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", login.Name, otp)
		err = s.SendSMSV2(ctx, login.Mobile, text)
		if err != nil {
			return errors.New("Sms Sending Error - " + err.Error())
		}
		if err == errors.New(constants.INSUFFICIENTBALANCE) {
			return err
		}
	}
	return nil
}

//OTPLoginValidateOTP :
func (s *Service) DealerregistrationValidateOTP(ctx *models.Context, login *models.RegistrationDealer) (*models.RefDealer, bool, error) {
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, false, err
	}

	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		data, err := s.Daos.GetSingleDealerWithMobileNo(ctx, login.Mobile)
		if err != nil {
			fmt.Println("GetSingleDealerWithMobileNo -", err)
			return err
		}
		if data != nil {
			return errors.New("Dealer already registered, please login")
		}
		if data == nil {
			key := fmt.Sprintf("%v_%v", constants.DEALERREGISTRATION, login.Mobile)
			otp := new(models.Otp)
			err = s.GetValueCacheMemory(key, otp)
			if err != nil {
				return err
			}
			fmt.Println("Otp===>", otp.Otp)
			if otp.Otp != login.OTP {
				return errors.New("Invaild Otp")
			}
			t := time.Now()
			login.Dealer.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDEALER)
			login.Dealer.Status = "Active"
			login.Dealer.Created.On = &t
			login.Dealer.Created.By = constants.SYSTEM
			reg := s.Daos.SaveDealer(ctx, &login.Dealer)
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
	data, err := s.Daos.GetSingleDealer(ctx, login.ID.Hex())
	if err != nil {
		return nil, false, err
	}
	return data, true, nil

}
