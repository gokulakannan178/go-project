package services

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
)

//CustomerLogin :
func (s *Service) CustomerLogin(ctx *models.Context, login *models.Login) (string, bool, error) {
	data, err := s.Daos.GetSingleCustomer(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return "dal err", false, err
	}
	if ok := data.Password == login.PassWord; !ok {
		log.Println("Data password ==>", data.Password)
		log.Println("login password ==>", login.PassWord)
		return "Passs false", false, nil
	}
	if data.Status == constants.CUSTOMERSTATUSINIT {
		return "", false, errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return "", false, errors.New("Please Contact Administrator")
	}
	// var auth models.Authentication
	// auth.UserID = data.ID
	// auth.UserName = data.UserName

	// auth.Status = data.Status
	// auth.Role = data.Role
	// fmt.Println("auth user ==>", auth, data)
	// token, err := CreateToken(&auth)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return "", false, errors.New("Error in Generating Token - " + err.Error())
	// }
	// data.Token = token
	// data.CurrentLocation = login.Location
	// err = s.Daos.UpdateUserWithUniqueID(data.UserName, data)
	// if err != nil {
	// 	log.Println("Error in saving token - " + err.Error())
	// 	return "", false, errors.New(constants.INTERNALSERVERERROR)
	// }
	return "", true, nil
}

//CustomerOTPLoginGenerateOTP :
func (s *Service) CustomerOTPLoginGenerateOTP(ctx *models.Context, login *models.Login) error {
	data, err := s.Daos.GetSingleGetUsingMobileNumber(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data == nil {
		return errors.New("Customer Not Available")
	}
	if data.Status == constants.USERSTATUSINIT {
		return errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return errors.New("Please Contact Administrator")
	}

	otp, err := s.GenerateOTP(constants.USERLOGIN, login.UserName, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
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

//CustomerOTPLoginValidateOTP :
func (s *Service) CustomerOTPLoginValidateOTP(ctx *models.Context, login *models.OTPLogin) (*models.RefCustomer, bool, error) {

	data, err := s.Daos.GetSingleGetUsingMobileNumber(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	if data == nil {
		return nil, false, errors.New("Customer Not Available")
	}
	if data.Status == constants.USERSTATUSINIT {
		return nil, false, errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return nil, false, errors.New("Please Contact Administrator")
	}

	err = s.ValidateOTP(constants.USERLOGIN, login.Mobile, login.OTP)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}

	var auth models.Authentication
	auth.UserID = data.ID
	auth.UserName = data.Name
	//auth.Type = data.Type
	auth.Status = data.Status
	//auth.Role = data.Role

	return data, true, nil

}
