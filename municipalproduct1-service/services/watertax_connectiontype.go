package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveWaterTaxConnectionType : ""
func (s *Service) SaveWaterTaxConnectionType(ctx *models.Context, watertaxconnectiontype *models.WaterTaxConnectionType) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	watertaxconnectiontype.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWATERTAXCONNECTIONTYPE)
	watertaxconnectiontype.Status = constants.WATERTAXCONNECTIONTYPESTATUSACTIVE
	//WaterTaxConnectionType.PaymentStatus = constants.WaterTaxConnectionTypePAYMENDSTATUS
	t := time.Now()
	Created := new(models.CreatedV2)
	Created.On = &t
	//WaterTaxConnectionType.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWaterTaxConnectionType(ctx, watertaxconnectiontype)
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

//GetSingleWaterTaxConnectionType :""
func (s *Service) GetSingleWaterTaxConnectionType(ctx *models.Context, UniqueID string) (*models.RefWaterTaxConnectionType, error) {
	watertaxconnectiontype, err := s.Daos.GetSingleWaterTaxConnectionType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return watertaxconnectiontype, nil
}

// UpdateWaterTaxConnectionType : ""
func (s *Service) UpdateWaterTaxConnectionType(ctx *models.Context, WaterTaxConnectionType *models.WaterTaxConnectionType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateWaterTaxConnectionType(ctx, WaterTaxConnectionType)
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

// EnableWaterTaxConnectionType : ""
func (s *Service) EnableWaterTaxConnectionType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableWaterTaxConnectionType(ctx, UniqueID)
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

//DisableWaterTaxConnectionType : ""
func (s *Service) DisableWaterTaxConnectionType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableWaterTaxConnectionType(ctx, UniqueID)
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

//DeleteWaterTaxConnectionType : ""
func (s *Service) DeleteWaterTaxConnectionType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWaterTaxConnectionType(ctx, UniqueID)
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

// FilterWaterTaxConnectionType : ""
func (s *Service) FilterWaterTaxConnectionType(ctx *models.Context, filter *models.WaterTaxConnectionTypeFilter, pagination *models.Pagination) ([]models.RefWaterTaxConnectionType, error) {
	return s.Daos.FilterWaterTaxConnectionType(ctx, filter, pagination)

}
