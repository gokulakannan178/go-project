package service

import (
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"log"
)

//GenerateOTP :
func (s *Service) GenerateOTP(scenario string, uniqueKey string, otplength int, ttl int) (string, error) {
	// prefix, err := s.OTPScenario(scenario)
	// if err != nil {
	// 	return "", err
	// }
	// otp := "9999"
	// fmt.Println("Env==>", s.Shared.GetCmdArg(constants.ENV))
	// if s.Shared.GetCmdArg(constants.ENV) != "development" || s.Shared.GetCmdArg(constants.ENV) != "prod" {
	// 	key := fmt.Sprintf("%v%v", prefix, uniqueKey)
	// 	otp = s.Shared.GetRandomOTP(otplength)

	// 	err = s.Redis.SetValue(key, otp, ttl)
	// 	log.Println("Key==>", key, "otp==>", otp)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	otp = "9999"

	// }

	// return otp, err
	prefix, err := s.OTPScenario(scenario)
	if err != nil {
		return "", err
	}
	key := fmt.Sprintf("%v%v", prefix, uniqueKey)
	otp := s.Shared.GetRandomOTP(otplength)
	if s.Shared.GetCmdArg(constants.ENV) == "development" {
		otp = "9999"
	} else {
		// if s.Shared.GetCmdArg(constants.ENV) == "prod" {
		// 	otp = "9999"
		// }
		err = s.Redis.SetValue(key, otp, ttl)
		log.Println("Key==>", key, "otp==>", otp)
		if err != nil {
			return "", err
		}
	}
	return otp, err
}

//OTPScenario : ""
func (s *Service) OTPScenario(scenario string) (string, error) {
	switch scenario {
	case constants.CONSUMERLOGIN:
		return "CONSLOGIN_", nil
	case constants.OTPSCENARIOPASSWORD:
		return "FRGPWDOTP_", nil
	case constants.OTPSCENARIOTOKEN:
		return "FRGPWDOTPTOKEN_", nil
	case constants.USERLOGIN:
		return "USRLOGIN_", nil
	case constants.USERREGISTRATION:
		return "USRREGISTRATION", nil

	default:
		return "", errors.New("No such scenario")
	}
}

//ValidateOTP :
func (s *Service) ValidateOTP(scenario string, uniqueKey string, otp string) error {
	// prefix, err := s.OTPScenario(scenario)
	// if err != nil {
	// 	return err
	// }
	// key := fmt.Sprintf("%v%v", prefix, uniqueKey)
	// rotp := "9999"
	// fmt.Println("Env==>", s.Shared.GetCmdArg(constants.ENV))
	// if s.Shared.GetCmdArg(constants.ENV) != "development" || s.Shared.GetCmdArg(constants.ENV) != "prod" {
	// 	var ok bool
	// 	redisToken := s.Redis.GetValue(key)
	// 	rotp, ok = redisToken.(string)
	// 	if !ok {
	// 		log.Println("Cannot type cast from redis to string")
	// 		return errors.New(constants.INTERNALSERVERERROR)
	// 	}

	// 	log.Println("Key==>", key, "otp==>", rotp)
	// 	validateOtp := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.VALIDATEOTP)
	// 	if validateOtp == "NO" {
	// 		return nil
	// 	}
	// 	rotp = "9999"
	// 	if otp != rotp {
	// 		log.Println("Not a valid otp")
	// 		return errors.New(constants.INVALIDOTP)
	// 	}
	// }

	// return nil
	if s.Shared.GetCmdArg(constants.ENV) == "development" {
		return nil
	}
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
