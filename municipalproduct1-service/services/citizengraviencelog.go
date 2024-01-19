package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveCitizenGraviansLog : ""
func (s *Service) SaveCitizenGraviansLog(ctx *models.Context, citizengravianslog *models.CitizenGraviansLog) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	citizengravianslog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
	citizengravianslog.Status = constants.CITIZENGRAVIANSLOGSTATUSACTIVE
	t := time.Now()
	//CitizenGraviansLog.Created = new(models.CreatedV2)
	citizengravianslog.On = &t
	citizengravianslog.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveCitizenGraviansLog(ctx, citizengravianslog)
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

//GetSingleCitizenGraviansLog :""
func (s *Service) GetSingleCitizenGraviansLog(ctx *models.Context, UniqueID string) (*models.RefCitizenGraviansLog, error) {
	tower, err := s.Daos.GetSingleCitizenGraviansLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateCitizenGraviansLog : ""
func (s *Service) UpdateCitizenGraviansLog(ctx *models.Context, CitizenGraviansLog *models.CitizenGraviansLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCitizenGraviansLog(ctx, CitizenGraviansLog)
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

// EnableCitizenGraviansLog : ""
func (s *Service) EnableCitizenGraviansLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableCitizenGraviansLog(ctx, UniqueID)
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

//DisableCitizenGraviansLog : ""
func (s *Service) DisableCitizenGraviansLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableCitizenGraviansLog(ctx, UniqueID)
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

//DeleteCitizenGraviansLog : ""
func (s *Service) DeleteCitizenGraviansLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCitizenGraviansLog(ctx, UniqueID)
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

// FilterCitizenGraviansLog : ""
func (s *Service) FilterCitizenGraviansLog(ctx *models.Context, filter *models.CitizenGraviansLogFilter, pagination *models.Pagination) ([]models.RefCitizenGraviansLog, error) {
	return s.Daos.FilterCitizenGraviansLog(ctx, filter, pagination)

}
