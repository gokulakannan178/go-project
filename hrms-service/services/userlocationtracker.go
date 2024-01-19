package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveUserLocationTracker : ""
func (s *Service) SaveUserLocationTracker(ctx *models.Context, userLocationTracker *models.UserLocationTracker) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	userLocationTracker.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERLOCATIONTRACKER)
	userLocationTracker.Status = constants.USERLOCATIONTRACKERSTATUSACTIVE
	t := time.Now()
	userLocationTracker.TimeStamp = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 UserLocationTracker.created")
	userLocationTracker.Created = &created
	log.Println("b4 UserLocationTracker.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUserLocationTracker(ctx, userLocationTracker)
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

// GetSingleUserLocationTracker : ""
func (s *Service) GetSingleUserLocationTracker(ctx *models.Context, UniqueID string) (*models.RefUserLocationTracker, error) {
	UserLocationTracker, err := s.Daos.GetSingleUserLocationTracker(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return UserLocationTracker, nil
}

//UpdateUserLocationTracker : ""
func (s *Service) UpdateUserLocationTracker(ctx *models.Context, userLocationTracker *models.UserLocationTracker) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserLocationTracker(ctx, userLocationTracker)
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
func (s *Service) EnableUserLocationTracker(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableUserLocationTracker(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableUserLocationTracker : ""
func (s *Service) DisableUserLocationTracker(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableUserLocationTracker(ctx, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
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
func (s *Service) FilterUserLocationTracker(ctx *models.Context, userLocationTracker *models.FilterUserLocationTracker, pagination *models.Pagination) (userLocationTrackers []models.RefUserLocationTracker, err error) {
	return s.Daos.FilterUserLocationTracker(ctx, userLocationTracker, pagination)
}
