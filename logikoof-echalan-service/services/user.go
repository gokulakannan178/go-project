package services

import (
	"errors"
	"fmt"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveUser(ctx *models.Context, user *models.User) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	user.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
	user.Status = constants.USERSTATUSACTIVE
	user.Password = "#nature32" //Default Password
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	user.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUser(ctx, user)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//UpdateUser : ""
func (s *Service) UpdateUser(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUser(ctx, user)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//EnableUser : ""
func (s *Service) EnableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUser(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DisableUser : ""
func (s *Service) DisableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUser(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DeleteUser : ""
func (s *Service) DeleteUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUser(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//GetSingleUser :""
func (s *Service) GetSingleUser(ctx *models.Context, UniqueID string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//FilterUser :""
func (s *Service) FilterUser(ctx *models.Context, userfilter *models.UserFilter, pagination *models.Pagination) (user []models.RefUser, err error) {
	return s.Daos.FilterUser(ctx, userfilter, pagination)
}

//ResetUserPassword : ""
func (s *Service) ResetUserPassword(ctx *models.Context, userName string) error {
	return s.Daos.ResetUserPassword(ctx, userName, "#nature32")
}

//ChangePassword : ""
func (s *Service) ChangePassword(ctx *models.Context, cp *models.UserChangePassword) (bool, string, error) {
	user, err := s.Daos.GetSingleUser(ctx, cp.UserName)
	if err != nil {
		return false, "", err
	}
	if user.Password != cp.OldPassword {
		return false, "Wrong Password", nil
	}
	err = s.Daos.ResetUserPassword(ctx, cp.UserName, cp.NewPassword)
	if err != nil {
		return false, "", err
	}
	return true, "", nil
}

//ConsumerLoginSendOTP : ""
func (s *Service) ConsumerLoginSendOTP(ctx *models.Context, regNo string) error {
	vehicle, err := s.Daos.GetSingleVehicle(ctx, regNo)
	if err != nil {
		return errors.New("Error in finding vehicle - " + err.Error())
	}
	if vehicle == nil {
		return errors.New("Error in finding vehicle - id null ")
	}
	token, err := s.GenerateOTP(constants.CONSUMERLOGIN, vehicle.Mobile, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return err
	}
	fmt.Println("token =>", token)
	msg := fmt.Sprintf("Your OTP for Logikoof Echallan portal is %v, Please do not share to any one",
		token,
	)
	s.SendSMS(vehicle.Mobile, msg)
	return nil
}

//ConsumerLoginValidateOTP : ""
func (s *Service) ConsumerLoginValidateOTP(ctx *models.Context, regNo string, otp string) error {
	vehicle, err := s.Daos.GetSingleVehicle(ctx, regNo)
	if err != nil {
		return errors.New("Error in finding vehicle - " + err.Error())
	}
	if vehicle == nil {
		return errors.New("Error in finding vehicle - id null ")
	}
	return s.ValidateOTP(constants.CONSUMERLOGIN, vehicle.Mobile, otp)
}
