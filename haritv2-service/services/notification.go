package services

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveNotification :""
func (s *Service) SaveNotification(ctx *models.Context, notification *models.Notification) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	notification.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONNOTIFICATION)
	notification.Status = constants.NOTIFICATIONSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Notification.created")
	notification.Created = &created
	log.Println("b4 Notification.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveNotification(ctx, notification)
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

//UpdateNotification : ""
func (s *Service) UpdateNotification(ctx *models.Context, notification *models.Notification) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateNotification(ctx, notification)
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

//EnableNotification : ""
func (s *Service) EnableNotification(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableNotification(ctx, UniqueID)
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

//DisableNotification : ""
func (s *Service) DisableNotification(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableNotification(ctx, UniqueID)
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

//DeleteNotification : ""
func (s *Service) DeleteNotification(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteNotification(ctx, UniqueID)
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

//GetSingleNotification :""
func (s *Service) GetSingleNotification(ctx *models.Context, UniqueID string) (*models.RefNotification, error) {
	Notification, err := s.Daos.GetSingleNotification(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Notification, nil
}

//FilterNotification :""
func (s *Service) FilterNotification(ctx *models.Context, filter *models.NotificationFilter, pagination *models.Pagination) ([]models.RefNotification, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterNotification(ctx, filter, pagination)

}
