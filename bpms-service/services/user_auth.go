package services

import (
	"bpms-service/models"
	"errors"
	"fmt"
)

//Login : ""
func (s *Service) Login(ctx *models.Context, login *models.Login) (string, bool, *models.RefUser, error) {
	refUser, err := s.GetSingleUser(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return "dal err in finding user", false, nil, err
	}
	if refUser == nil {
		return "", false, nil, errors.New("user Not Found")
	}
	if ok := refUser.Password == login.PassWord; !ok {
		return "Password false", false, nil, nil
	}
	return "", true, refUser, nil
}
