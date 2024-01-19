package services

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveFPOInventory : ""
func (s *Service) SaveFPOInventory(ctx *models.Context, fpo *models.FPOInventory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	fpo.CompanyID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFPOINVENTORY)
	fpo.Status = constants.FPOINVENTORYSTATUSACTIVE
	t := time.Now()
	created := new(models.CreatedV2)
	created.On = &t
	created.By = constants.SYSTEM
	fpo.Created = created
	fpo.IsSellingPriceUpdated = true
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFPOInventory(ctx, fpo)
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

//GetSingleFPOInventory :""
func (s *Service) GetSingleFPOInventory(ctx *models.Context, UniqueID string) (*models.RefFPOINVENTORY, error) {
	tower, err := s.Daos.GetSingleFPOInventory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateFPOInventory : ""
func (s *Service) UpdateFPOInventory(ctx *models.Context, fpo *models.FPOInventory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFPOInventory(ctx, fpo)
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

// EnableFPOInventory : ""
func (s *Service) EnableFPOInventory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFPOInventory(ctx, UniqueID)
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

//DisableFPOInventory : ""
func (s *Service) DisableFPOInventory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFPOInventory(ctx, UniqueID)
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

//DeleteFPOInventory : ""
func (s *Service) DeleteFPOInventory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFPOInventory(ctx, UniqueID)
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

// // FilterFPOInventory : ""
// func (s *Service) FilterFPOInventory(ctx *models.Context, filter *models.FPOInventoryFilter, pagination *models.Pagination) ([]models.RefFPOInventory, error) {
// 	return s.Daos.FilterFPOInventory(ctx, filter, pagination)

// }

// FPOInventoryQuantityUpdate : ""
func (s *Service) FPOInventoryQuantityUpdate(ctx *models.Context, fpo *models.FPOInventory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFPOInventory(ctx, fpo)
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

// FPOInventoryQuantityUpdate : ""
func (s *Service) FPOInventoryPriceUpdate(ctx *models.Context, fpo *models.FPOInventory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.FPOInventoryPriceUpdate(ctx, fpo)
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

//GetSingleFPOInventory :""
func (s *Service) GetSingleFPOInventoryWithCompalyID(ctx *models.Context, UniqueID string) (*models.RefFPOINVENTORY, error) {
	tower, err := s.Daos.GetSingleFPOInventoryWithCompalyID(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}
