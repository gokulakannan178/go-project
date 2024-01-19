package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyFixedArvLog : ""
func (s *Service) SavePropertyFixedArvLog(ctx *models.Context, propertyfixedarvlog *models.PropertyFixedArvLog) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	propertyfixedarvlog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARVLOG)
	propertyfixedarvlog.Status = constants.PROPERTYFIXEDARVLOGSTATUSACTIVE
	//PropertyFixedArvLog.PaymentStatus = constants.PropertyFixedArvLogPAYMENDSTATUS
	t := time.Now()
	Created := new(models.CreatedV2)
	Created.On = &t
	//PropertyFixedArvLog.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePropertyFixedArvLog(ctx, propertyfixedarvlog)
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

//GetSinglePropertyFixedArvLog :""
func (s *Service) GetSinglePropertyFixedArvLog(ctx *models.Context, UniqueID string) (*models.RefPropertyFixedArvLog, error) {
	propertyfixedarvlog, err := s.Daos.GetSinglePropertyFixedArvLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertyfixedarvlog, nil
}

// UpdatePropertyFixedArvLog : ""
func (s *Service) UpdatePropertyFixedArvLog(ctx *models.Context, propertyfixedarvlog *models.PropertyFixedArvLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyFixedArvLog(ctx, propertyfixedarvlog)
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

// EnablePropertyFixedArvLog : ""
func (s *Service) EnablePropertyFixedArvLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyFixedArvLog(ctx, UniqueID)
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

//DisablePropertyFixedArvLog : ""
func (s *Service) DisablePropertyFixedArvLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyFixedArvLog(ctx, UniqueID)
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

//DeletePropertyFixedArvLog : ""
func (s *Service) DeletePropertyFixedArvLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyFixedArvLog(ctx, UniqueID)
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

// FilterPropertyFixedArvLog : ""
func (s *Service) FilterPropertyFixedArvLog(ctx *models.Context, filter *models.PropertyFixedArvLogFilter, pagination *models.Pagination) ([]models.RefPropertyFixedArvLog, error) {
	return s.Daos.FilterPropertyFixedArvLog(ctx, filter, pagination)

}
