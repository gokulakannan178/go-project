package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
)

//OTPLoginGenerateOTP :
func (s *Service) LoginGenerateotpFarmer(ctx *models.Context, login *models.FarmerLogin) error {
	data, err := s.Daos.GetSingleFarmerWithMobileno(ctx, login.MobileNumber)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data == nil {
		return errors.New("Farmer Not Available")
	}
	if data.Status == constants.FARMERSTATUSINIT {
		return errors.New("Awaiting Activation")
	}
	if data.Status != constants.FARMERSTATUSACTIVE {
		return errors.New("Please Contact Administrator")
	}

	//otp, err := s.GenerateOTP(constants.FARMERLOGIN, login.MobileNumber, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	key := fmt.Sprintf("%v_%v", constants.FARMERLOGIN, login.MobileNumber)
	var otp models.Otp
	otp.Otp = "9999"
	err = s.SetValueCacheMemory(key, otp, 1000)
	if err != nil {
		return err
	}
	// if err != nil {
	// 	return errors.New("Otp Generate Error - " + err.Error())
	// }
	// text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", data.Name, otp)
	// err = s.SendSMS(login.MobileNumber, text)
	// if err != nil {
	// 	return errors.New("Sms Sending Error - " + err.Error())
	// }
	//loginurl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.LOGINURLV2)

	//msg := fmt.Sprintf(constants.COMMONTEMPLATE, data.Name, "NICESSM", "user password updated", "please login:"+loginurl+"", "https://nicessm.org/")
	msg := fmt.Sprintf(constants.COMMONTEMPLATE, data.Name, "NICESSM", "otp for login", "OTP for NICESSM forgot password is-"+otp.Otp+"", "https://nicessm.org/")

	err = s.SendSMSV2(ctx, login.MobileNumber, msg)
	if err != nil {
		log.Println(login.MobileNumber + " " + err.Error())
	}
	if err == errors.New(constants.INSUFFICIENTBALANCE) {
		return err
	}
	return nil
}

//OTPLoginValidateOTP :
func (s *Service) LoginValidateOTPFarmer(ctx *models.Context, login *models.FarmerOTPLogin) (*models.RefFarmer, bool, error) {

	data, err := s.Daos.GetSingleFarmerWithMobileno(ctx, login.MobileNumber)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	if data == nil {
		return nil, false, errors.New("Farmer Not Available")
	}
	if data.Status == constants.FARMERSTATUSINIT {
		return nil, false, errors.New("Awaiting Activation")
	}
	if data.Status != constants.FARMERSTATUSACTIVE {
		return nil, false, errors.New("Please Contact Administrator")
	}
	// err = s.ValidateOTP(constants.FARMERLOGIN, login.MobileNumber, login.OTP)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, false, err
	// }
	key := fmt.Sprintf("%v_%v", constants.FARMERLOGIN, login.MobileNumber)
	otp := new(models.Otp)
	err = s.GetValueCacheMemory(key, otp)
	if err != nil {
		return nil, false, err
	}
	fmt.Println("Otp===>", otp.Otp)
	if otp.Otp != login.OTP {
		return nil, false, errors.New("Invaild Otp")
	}
	var auth models.FarmerAuthentication
	auth.FarmerID = data.ID
	auth.Name = data.Name
	auth.Status = data.Status
	// token, err := CreateTokenV2(&auth)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return nil, false, errors.New("Error in Generating Token - " + err.Error())
	// }
	// //data.User.Token = token
	// data.Token = token

	return data, true, nil

}
