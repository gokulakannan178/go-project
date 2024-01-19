package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveParkingPenalty : ""
func (s *Service) SaveParkingPenalty(ctx *models.Context, parkingPenalty *models.ParkingPenalty) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	parkingPenalty.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPARKINGPENALTIES)
	parkingPenalty.Status = constants.PARKINGPENALTYSTATUSACTIVE
	t := time.Now()
	parkingPenalty.Created = new(models.CreatedV2)
	parkingPenalty.Created.On = &t
	parkingPenalty.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveParkingPenalty(ctx, parkingPenalty)
		if dberr != nil {
			return dberr
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

//GetSingleParkingPenalty :""
func (s *Service) GetSingleParkingPenalty(ctx *models.Context, UniqueID string) (*models.RefParkingPenalty, error) {
	tower, err := s.Daos.GetSingleParkingPenalty(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateParkingPenalty : ""
func (s *Service) UpdateParkingPenalty(ctx *models.Context, parkingPenalty *models.ParkingPenalty) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateParkingPenalty(ctx, parkingPenalty)
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

// EnableParkingPenalty : ""
func (s *Service) EnableParkingPenalty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableParkingPenalty(ctx, UniqueID)
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

//DisableParkingPenalty : ""
func (s *Service) DisableParkingPenalty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableParkingPenalty(ctx, UniqueID)
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

//DeleteParkingPenalty : ""
func (s *Service) DeleteParkingPenalty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteParkingPenalty(ctx, UniqueID)
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

// FilterParkingPenalty : ""
func (s *Service) FilterParkingPenalty(ctx *models.Context, filter *models.ParkingPenaltyFilter, pagination *models.Pagination) ([]models.RefParkingPenalty, error) {
	return s.Daos.FilterParkingPenalty(ctx, filter, pagination)

}
