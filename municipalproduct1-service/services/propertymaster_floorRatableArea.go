package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFloorRatableArea :""
func (s *Service) SaveFloorRatableArea(ctx *models.Context, floorRatableArea *models.FloorRatableArea) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	floorRatableArea.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFLOORRATABLEAREA)
	floorRatableArea.Status = constants.FLOORRATABLEAREASTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 floorRatableArea.created")
	floorRatableArea.Created = created
	log.Println("b4 floorRatableArea.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFloorRatableArea(ctx, floorRatableArea)
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

//UpdateFloorRatableArea : ""
func (s *Service) UpdateFloorRatableArea(ctx *models.Context, floorRatableArea *models.FloorRatableArea) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFloorRatableArea(ctx, floorRatableArea)
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

//EnableFloorRatableArea : ""
func (s *Service) EnableFloorRatableArea(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFloorRatableArea(ctx, UniqueID)
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

//DisableFloorRatableArea : ""
func (s *Service) DisableFloorRatableArea(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFloorRatableArea(ctx, UniqueID)
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

//DeleteFloorRatableArea : ""
func (s *Service) DeleteFloorRatableArea(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFloorRatableArea(ctx, UniqueID)
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

//GetSingleFloorRatableArea :""
func (s *Service) GetSingleFloorRatableArea(ctx *models.Context, UniqueID string) (*models.RefFloorRatableArea, error) {
	floorRatableArea, err := s.Daos.GetSingleFloorRatableArea(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return floorRatableArea, nil
}

//FilterFloorRatableArea :""
func (s *Service) FilterFloorRatableArea(ctx *models.Context, floorRatableAreafilter *models.FloorRatableAreaFilter, pagination *models.Pagination) (floorRatableArea []models.RefFloorRatableArea, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterFloorRatableArea(ctx, floorRatableAreafilter, pagination)
}
