package services

import (
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveApptoken :""
func (s *Service) SaveApptoken(ctx *models.Context, apptoken *models.Apptoken) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	apptoken.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONAPPTOKEN)
	apptoken.Status = constants.APPTOKENSTATUSACTIVE
	t := time.Now()
	// created := models.Created{}
	// created.On = &t
	// created.By = constants.SYSTEM
	// log.Println("b4 PkgType.created")
	apptoken.Currenttime = &t
	log.Println("b4 PkgType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//	user, _ := s.Daos.GetRegTokenWithUserId(ctx, apptoken.UserID)
		fmt.Println("userId===>", apptoken.UserID)
		fmt.Println("userType===>", apptoken.Usertype)
		dberr := s.Daos.SaveApptoken2(ctx, apptoken)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}

		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//UpdateApptoken : ""
func (s *Service) UpdateApptoken(ctx *models.Context, apptoken *models.Apptoken) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateApptoken(ctx, apptoken)
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

//EnableApptoken : ""
func (s *Service) EnableApptoken(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableApptoken(ctx, UniqueID)
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

//DisableApptoken : ""
func (s *Service) DisableApptoken(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableApptoken(ctx, UniqueID)
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

//DeleteApptoken : ""
func (s *Service) DeleteApptoken(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteApptoken(ctx, UniqueID)
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

//GetSingleAppptoken :""
func (s *Service) GetSingleApptoken(ctx *models.Context, UniqueID string) (*models.RefApptoken, error) {
	apptoken, err := s.Daos.GetSingleApptoken(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return apptoken, nil
}

//FilterApptoken :""
func (s *Service) FilterApptoken(ctx *models.Context, filter *models.ApptokenFilter, pagination *models.Pagination) ([]models.RefApptoken, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterApptoken(ctx, filter, pagination)

}

// //SendAppTokenNotification :""
// func (s *Service) SendAppTokenNotification(ctx *models.Context, appToken *models.AppTokenNotification) error {
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	res, err := s.Daos.SendAppTokenNotification(ctx, appToken)
// 	if err != nil {
// 		return err
// 	}
// 	if res == nil {
// 		return nil
// 	}
// 	if len(res.RegistrationToken) > 0 {
// 		notification, err := s.Daos.GetSingleNotification(ctx, appToken.Notification)
// 		if err != nil {
// 			return err
// 		}
// 		if notification == nil {
// 			return errors.New("Notification cannot be nil")
// 		}
// 		if err = s.SendNotification(notification.Title, notification.Content, notification.Img, res.RegistrationToken); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
func (s *Service) LogoutApptoken(ctx *models.Context, UniqueID string) (err error) {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err = s.Daos.LogoutApptoken(ctx, UniqueID)
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}

		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}
func (s *Service) GetSingleApptokenWithUniqueCheck(ctx *models.Context, UniqueID string) (*models.RefApptoken, error) {
	apptoken, err := s.Daos.GetSingleApptokenWithUniqueCheck(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return apptoken, nil
}
