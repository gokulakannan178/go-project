package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveWaterTaxArv : ""
func (s *Service) SaveWaterTaxArv(ctx *models.Context, watertaxarv *models.WaterTaxArv) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	watertaxarv.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWATERTAXARV)
	watertaxarv.Status = constants.WATERTAXARVSTATUSACTIVE
	//WaterTaxArv.PaymentStatus = constants.WaterTaxArvPAYMENDSTATUS
	t := time.Now()
	Created := new(models.CreatedV2)
	Created.On = &t
	//WaterTaxArv.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWaterTaxArv(ctx, watertaxarv)
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

//GetSingleWaterTaxArv :""
func (s *Service) GetSingleWaterTaxArv(ctx *models.Context, UniqueID string) (*models.RefWaterTaxArv, error) {
	WaterTaxArv, err := s.Daos.GetSingleWaterTaxArv(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return WaterTaxArv, nil
}

// UpdateWaterTaxArv : ""
func (s *Service) UpdateWaterTaxArv(ctx *models.Context, watertaxarv *models.WaterTaxArv) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateWaterTaxArv(ctx, watertaxarv)
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

// EnableWaterTaxArv : ""
func (s *Service) EnableWaterTaxArv(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableWaterTaxArv(ctx, UniqueID)
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

//DisableWaterTaxArv : ""
func (s *Service) DisableWaterTaxArv(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableWaterTaxArv(ctx, UniqueID)
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

//DeleteWaterTaxArv : ""
func (s *Service) DeleteWaterTaxArv(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWaterTaxArv(ctx, UniqueID)
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

// FilterWaterTaxArv : ""
func (s *Service) FilterWaterTaxArv(ctx *models.Context, filter *models.WaterTaxArvFilter, pagination *models.Pagination) ([]models.RefWaterTaxArv, error) {
	return s.Daos.FilterWaterTaxArv(ctx, filter, pagination)

}
