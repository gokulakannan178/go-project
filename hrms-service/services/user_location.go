package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUserLocation :""
func (s *Service) SaveUserLocation(ctx *models.Context, userLocation *models.UserLocation) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	userLocation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERLOCATION)

	userLocation.EntryType = constants.USERLOCATIONENTRYSTATUSGENERAL

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		user, usrErr := s.Daos.GetSingleUser(ctx, userLocation.UserName)
		if user == nil {
			userLocation.ErrMsg = "User not found"
			userLocation.Status = constants.USERLOCATIONSTATUSERROR
		}
		if usrErr != nil {
			userLocation.ErrMsg = "error in getting user - " + usrErr.Error()
			userLocation.Status = constants.USERLOCATIONSTATUSERROR
		}
		if user != nil {
			userLocation.Name = user.Name
			userLocation.UserType = user.Type
			userLocation.Status = constants.USERLOCATIONSTATUSACTIVE
		}
		t := time.Now()
		userLocation.Time = &t
		dberr := s.Daos.SaveUserLocation(ctx, userLocation)
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
