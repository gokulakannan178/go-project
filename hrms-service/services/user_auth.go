package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"

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
	data, err := s.Daos.GetSingleUserWithLogin(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return "dal err", false, err
	}
	if data == nil {
		return "", false, errors.New("Pls Check Username")
	}
	if ok := data.Password == login.PassWord; !ok {
		log.Println("Data password ==>", data.Password)
		log.Println("login password ==>", login.PassWord)
		return "Passs false", false, nil
	}
	if data.Status == constants.USERSTATUSINIT || data.Status == constants.EMPLOYEESTATUSONBORADING {
		return "", false, errors.New("Awaiting Activation")
	}

	if data.Status == constants.EMPLOYEESTATUSREJECT || data.Status == constants.EMPLOYEESTATUSRELIEVE || data.Status == constants.EMPLOYEESTATUSOFFBOARD {
		return "", false, errors.New("pls Contact Adminisator")

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
func (s *Service) GetSingleProfile(ctx *models.Context, UniqueID string) (interface{}, error) {

	employee, err := s.Daos.GetSingleEmployee(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
