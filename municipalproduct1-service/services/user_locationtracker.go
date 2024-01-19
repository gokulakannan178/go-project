package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveUserLocationTracker : ""
func (s *Service) SaveUserLocationTracker(ctx *models.Context, tracker *models.UserLocationTracker) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	tracker.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERLOCATIONTRACKER)
	tracker.Status = constants.USERLOCATIONTRACKERSTATUSACTIVE
	t := time.Now()
	tracker.Created = new(models.CreatedV2)
	tracker.Created.On = &t
	tracker.Created.By = constants.SYSTEM
	if tracker.TimeStamp == nil {
		tracker.TimeStamp = &t
	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUserLocationTracker(ctx, tracker)
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

//GetSingleUserLocationTracker :""
func (s *Service) GetSingleUserLocationTracker(ctx *models.Context, UniqueID string) (*models.RefUserLocationTracker, error) {
	tower, err := s.Daos.GetSingleUserLocationTracker(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateUserLocationTracker : ""
func (s *Service) UpdateUserLocationTracker(ctx *models.Context, tracker *models.UserLocationTracker) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserLocationTracker(ctx, tracker)
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

// EnableUserLocationTracker : ""
func (s *Service) EnableUserLocationTracker(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUserLocationTracker(ctx, UniqueID)
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

//DisableUserLocationTracker : ""
func (s *Service) DisableUserLocationTracker(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUserLocationTracker(ctx, UniqueID)
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

//DeleteUserLocationTracker : ""
func (s *Service) DeleteUserLocationTracker(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUserLocationTracker(ctx, UniqueID)
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

// FilterUserLocationTracker : ""
func (s *Service) FilterUserLocationTracker(ctx *models.Context, filter *models.UserLocationTrackerFilter, pagination *models.Pagination) ([]models.RefUserLocationTracker, error) {
	return s.Daos.FilterUserLocationTracker(ctx, filter, pagination)

}

//GetSingleUserLocationTracker :""
func (s *Service) UserLocationTrackerCoordinates(ctx *models.Context, coordinates *models.UserLocationTrackerCoordinates) (*models.RefUserLocationTracker, error) {
	tower, err := s.Daos.UserLocationTrackerCoordinates(ctx, coordinates)
	if err != nil {
		return nil, err
	}
	return tower, nil
}
