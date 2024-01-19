package services

import (
	"errors"
	"nicessm-api-service/models"
)

func (s *Service) GetAccessPrivillege(ctx *models.Context, userAccess models.UserAccess) (user *models.RefUser, err error) {
	if userAccess.UserName == "" {
		return nil, errors.New("username is missing")
	}
	user, err = s.Daos.GetSingleUserWithUserName(ctx, userAccess.UserName)
	if err != nil {
		return nil, errors.New("error getting access user - " + err.Error())
	}
	if user == nil {
		return nil, errors.New("access user is nil ")

	}
	return

}
