package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

const (
	timeOutInSeconds = 10000000
	timeOutInMinues  = 60
	sessionSecret    = "yourSe$$ion$ecret"
)

//CreateToken : ""
func CreateToken(authentication *models.Authentication) (string, error) {
	var claims = jws.Claims{
		"UserID":   authentication.UserID,
		"UserName": authentication.UserName,
		"Status":   authentication.Status,
		"Role":     authentication.Role,
	}
	// claims.SetIssuedAt(time.Now())
	// claims.SetExpiration(time.Now().Add(time.Duration(timeOutInMinues) * time.Minute))
	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)
	jwtToken, err := jwt.Serialize([]byte(sessionSecret))
	log.Println("TOKKKKKKKEN ", string(jwtToken))
	return string(jwtToken), err
}

//ValidateToken : ""
func ValidateToken(token string) (*models.Authentication, error) {
	// TODO: Check if the access token is available on redis store
	// if not? then simply return unauthorized

	parsedToken, err := jws.ParseJWT([]byte(string(token)))
	if err != nil {
		return nil, err
	}
	// err = (parsedToken.Validate([]byte(sessionSecret), crypto.SigningMethodHS256))
	// if err != nil {
	// 	if err.Error() == "token is expired" {
	// 		log.Println("token is expired")
	// 		return nil, err
	// 	}
	// }
	cbytes, err1 := json.Marshal(parsedToken.Claims())
	if err1 != nil {
		return nil, err1
	}
	authenticationData := new(models.Authentication)
	json.Unmarshal(cbytes, &authenticationData)
	return authenticationData, nil
}

//Login :
func (s *Service) Login(ctx *models.Context, login *models.Login) (string, bool, error) {
	data, err := s.Daos.GetSingleUser(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return "dal err", false, err
	}
	// resPD, err := s.GetSingleProductConfiguration(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "production configuration err", false, err
	// }
	if login.Environment == "Web" && data.AllowWebLogin == "No" {
		return "", false, errors.New("Please Contact Administrator")
	} else {
		if ok := data.Password == login.PassWord; !ok {
			log.Println("Data password ==>", data.Password)
			log.Println("login password ==>", login.PassWord)
			return "Passs false", false, nil
		}
		err = s.Daos.UpdateForcedLogout(ctx, login.UserName, "No")
		if err != nil {
			return "", false, err
		}
		if data.Status == constants.USERSTATUSINIT {
			return "", false, errors.New("Awaiting Activation")
		}
		if data.Status == constants.USERSTATUSACTIVE || data.Status == constants.USERSTATUSTESTUSER {
			webToken := s.Shared.GetUserToken(login.UserName, 64)
			fmt.Println("webToken ======> ", webToken)
			err = s.Daos.UpdateUserToken(ctx, login.UserName, webToken)
			if err != nil {
				return "", false, err
			}
			return webToken, true, nil
		}

		// if data.Status != constants.USERSTATUSACTIVE || data.Status != constants.USERSTATUSTESTUSER {
		return "", false, errors.New("Please Contact Administrator")
		// }
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
	// return "", true, nil
}

// //Login :
func (s *Service) LoginGenerateOTPV2(ctx *models.Context, mobileNo string, versionKey string) error {
	// key, err := strconv.ParseFloat(versionKey, 32)
	// if err != nil {
	// 	return err
	// }
	// productConfig, err := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	// if err != nil {
	// 	return errors.New("Error in getting config -" + err.Error())
	// }
	// if productConfig.AppVersion > key {
	// 	return errors.New("please update app version")

	// }
	user, err := s.Daos.GetSingleActiveUserWithUniqueID(ctx, mobileNo)
	if err != nil {
		return errors.New("Error in geting user - " + err.Error())
	}
	if user == nil {
		return errors.New("User not available")
	}
	if user.Status == constants.USERSTATUSDISABLED {
		return errors.New("Please contact administrator")
	}
	token, err := s.GenerateOTP(constants.USERLOGIN, mobileNo, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return err
	}
	fmt.Println("token =>", token)
	// msg := fmt.Sprintf("Your OTP for Kochas Municipal Corporation Consumer login is %v. This OTP is valid only for 3 minutes. Please do not share OTP to anyone",

	// 	token)
	// msg := fmt.Sprintf("Your OTP is %v KVK App", token)
	// COMMONSMS//TEMPLATE        = "Hi %v,You have received a notification from %v regarding login OTP\n Notification -OTP -  %v  This OTP is valid only for 3 minutes. Please do not share OTP to anyone\nRegards,\nFrom %v\nFor help/queries contact %v"
	msg := fmt.Sprintf("Hi %v,You have received a notification from %v regarding login OTP\n Notification -OTP -  %v  This OTP is valid only for 3 minutes. Please do not share OTP to anyone\nRegards,\nFrom %v\nFor help/queries contact %v",
		// msg := fmt.Sprintf("Dear %v your OTP for %v tax collector app login is %v. This OTP is valid only for 3 minutes. Please do not share OTP to anyone",
		//
		user.Name, ctx.ProductConfig.Name, token, "BRMNCP", ctx.ProductConfig.UIURL)
	//msg := fmt.Sprintf("Your password is %v NICESM", token)

	if err := s.SendSMS(mobileNo, msg); err != nil {
		log.Println(err)
	}
	return nil
}
