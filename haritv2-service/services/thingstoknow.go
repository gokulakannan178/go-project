package services

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveThingsToKnow :""
func (s *Service) SaveThingsToKnow(ctx *models.Context, ThingsToKnow *models.ThingsToKnow) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	ThingsToKnow.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTHINGSTOKNOW)
	ThingsToKnow.Status = constants.THINGSTOKNOWSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ThingsToKnow.created")
	ThingsToKnow.Created = &created
	log.Println("b4 ThingsToKnow.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveThingsToKnow(ctx, ThingsToKnow)
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

//UpdateThingsToKnow : ""
func (s *Service) UpdateThingsToKnow(ctx *models.Context, ThingsToKnow *models.ThingsToKnow) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateThingsToKnow(ctx, ThingsToKnow)
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

//EnableThingsToKnow : ""
func (s *Service) EnableThingsToKnow(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableThingsToKnow(ctx, UniqueID)
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

//DisableThingsToKnow : ""
func (s *Service) DisableThingsToKnow(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableThingsToKnow(ctx, UniqueID)
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

//DeleteThingsToKnow : ""
func (s *Service) DeleteThingsToKnow(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteThingsToKnow(ctx, UniqueID)
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

//GetSingleThingsToKnow :""
func (s *Service) GetSingleThingsToKnow(ctx *models.Context, UniqueID string) (*models.ThingsToKnow, error) {
	ThingsToKnow, err := s.Daos.GetSingleThingsToKnow(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ThingsToKnow, nil
}

//FilterThingsToKnow :""
func (s *Service) FilterThingsToKnow(ctx *models.Context, FilterThingsToKnow *models.FilterThingsToKnow, pagination *models.Pagination) (ThingsToKnow []models.ThingsToKnow, err error) {

	return s.Daos.FilterThingsToKnow(ctx, FilterThingsToKnow, pagination)

}
