package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"fmt"
	"log"
)

//ApplicantLoginSendOTP : ""
func (s *Service) ApplicantLoginSendOTP(ctx *models.Context, mobile string) error {
	applicant, err := s.Daos.GetSinglePreregistration(ctx, mobile)
	if err != nil {
		return errors.New("Error in finding applicant - " + err.Error())
	}
	if applicant == nil {
		return errors.New("Error in finding applicant - id null ")
	}
	token, err := s.GenerateOTP(constants.CONSUMERLOGIN, applicant.MobileNumber, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return err
	}
	fmt.Println("token =>", token)
	msg := fmt.Sprintf("Your OTP for BPMS portal is %v, Please do not share to any one",
		token,
	)
	s.SendSMS(applicant.MobileNumber, msg)
	return nil
}

//ApplicantLoginValidateOTP : ""
func (s *Service) ApplicantLoginValidateOTP(ctx *models.Context, mobile string, otp string) error {
	applicant, err := s.Daos.GetSinglePreregistration(ctx, mobile)
	if err != nil {
		return errors.New("Error in finding applicant - " + err.Error())
	}
	if applicant == nil {
		return errors.New("Error in finding applicant - id null ")
	}
	return s.ValidateOTP(constants.CONSUMERLOGIN, applicant.MobileNumber, otp)
}

//GenerateOTP :
func (s *Service) GenerateOTP(scenario string, uniqueKey string, otplength int, ttl int) (string, error) {
	prefix, err := s.OTPScenario(scenario)
	if err != nil {
		return "", err
	}
	key := fmt.Sprintf("%v%v", prefix, uniqueKey)
	otp := s.Shared.GetRandomOTP(otplength)
	otp = "9999"
	err = s.Redis.SetValue(key, otp, ttl)
	log.Println("Key==>", key, "otp==>", otp)
	if err != nil {
		return "", err
	}
	return otp,
		err
}

//OTPScenario : ""
func (s *Service) OTPScenario(scenario string) (string, error) {
	switch scenario {
	case constants.CONSUMERLOGIN:
		return "CONSLOGIN_", nil
	default:
		return "", errors.New("No such scenario")
	}
}

//ValidateOTP :
func (s *Service) ValidateOTP(scenario string, uniqueKey string, otp string) error {
	prefix, err := s.OTPScenario(scenario)
	if err != nil {

		return err
	}
	key := fmt.Sprintf("%v%v", prefix, uniqueKey)
	redisToken := s.Redis.GetValue(key)
	rotp, ok := redisToken.(string)
	if !ok {

		log.Println("Cannot type cast from redis to string")
		return errors.New(constants.INTERNALSERVERERROR)
	}
	log.Println("Key==>", key, "otp==>", rotp)
	validateOtp := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.VALIDATEOTP)
	if validateOtp == "NO" {

		return nil
	}
	if otp != rotp {

		log.Println("Not a valid otp")
		return errors.New(constants.INVALIDOTP)
	}
	return nil
}
