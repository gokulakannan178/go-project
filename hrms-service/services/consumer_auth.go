package services

import (
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
)

//SendOTPConsumerLogin : ""
func (s *Service) SendOTPConsumerLogin(ctx *models.Context, mobileNo string) error {

	token, _ := s.GenerateOTP(constants.CONSUMERLOGIN, mobileNo, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return err
	// }
	fmt.Println("token =>", token)
	msg := fmt.Sprintf("Your OTP for HRMS Consumer login is %v. This OTP is valid only for 3 minutes. Please do not share OTP to anyone",
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
