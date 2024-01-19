package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveUserChargeRateMaster : ""
func (s *Service) SaveUserChargeRateMaster(ctx *models.Context, userchargeratemaster *models.UserChargeRateMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	userchargeratemaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGERATEMASTER)
	userchargeratemaster.Status = constants.USERCHARGERATEMASTERSTATUSACTIVE
	t := time.Now()
	//UserChargeRateMaster.Created = new(models.CreatedV2)
	userchargeratemaster.Created.On = &t
	userchargeratemaster.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUserChargeRateMaster(ctx, userchargeratemaster)
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

//GetSingleUserChargeRateMaster :""
func (s *Service) GetSingleUserChargeRateMaster(ctx *models.Context, UniqueID string) (*models.RefUserChargeRateMaster, error) {
	tower, err := s.Daos.GetSingleUserChargeRateMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateUserChargeRateMaster : ""
func (s *Service) UpdateUserChargeRateMaster(ctx *models.Context, userchargeratemaster *models.UserChargeRateMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserChargeRateMaster(ctx, userchargeratemaster)
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

// EnableUserChargeRateMaster : ""
func (s *Service) EnableUserChargeRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUserChargeRateMaster(ctx, UniqueID)
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

//DisableUserChargeRateMaster : ""
func (s *Service) DisableUserChargeRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUserChargeRateMaster(ctx, UniqueID)
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

//DeleteUserChargeRateMaster : ""
func (s *Service) DeleteUserChargeRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUserChargeRateMaster(ctx, UniqueID)
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

// FilterUserChargeRateMaster : ""
func (s *Service) FilterUserChargeRateMaster(ctx *models.Context, filter *models.UserChargeRateMasterFilter, pagination *models.Pagination) ([]models.RefUserChargeRateMaster, error) {
	return s.Daos.FilterUserChargeRateMaster(ctx, filter, pagination)

}

func (s *Service) GetSingleUserChargeRateMasterWithCategoryId(ctx *models.Context, UniqueID string) (*models.RefUserChargeRateMaster, error) {
	tower, err := s.Daos.GetSingleUserChargeRateMasterWithCategoryId(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}
