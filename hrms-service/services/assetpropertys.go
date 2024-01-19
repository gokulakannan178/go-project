package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveAssetPropertys :""
func (s *Service) SaveAssetPropertys(ctx *models.Context, assetPropertys *models.AssetPropertys) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	assetPropertys.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETPROPERTYS)
	assetPropertys.Status = constants.ASSETPROPERTYSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AssetPropertys.created")
	assetPropertys.Created = created
	log.Println("b4 AssetPropertys.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAssetPropertys(ctx, assetPropertys)
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

//GetSingleAssetPropertys :""
func (s *Service) GetSingleAssetPropertys(ctx *models.Context, UniqueID string) (*models.RefAssetPropertys, error) {
	AssetPropertys, err := s.Daos.GetSingleAssetPropertys(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return AssetPropertys, nil
}

//UpdateAssetPropertys : ""
func (s *Service) UpdateAssetPropertys(ctx *models.Context, assetPropertys *models.AssetPropertys) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAssetPropertys(ctx, assetPropertys)
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

//EnableAssetPropertys : ""
func (s *Service) EnableAssetPropertys(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAssetPropertys(ctx, UniqueID)
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

//DisableAssetPropertys : ""
func (s *Service) DisableAssetPropertys(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAssetPropertys(ctx, UniqueID)
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

//DeleteAssetPropertys : ""
func (s *Service) DeleteAssetPropertys(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAssetPropertys(ctx, UniqueID)
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

//FilterAssetPropertys :""
func (s *Service) FilterAssetPropertys(ctx *models.Context, assetPropertysFilter *models.FilterAssetPropertys, pagination *models.Pagination) ([]models.RefAssetPropertys, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterAssetPropertys(ctx, assetPropertysFilter, pagination)

}

func (s *Service) SaveAssetPropertysWithoutTransaction(ctx *models.Context, assetPropertys *models.AssetPropertys) error {
	log.Println("transaction start")
	// Start Transaction

	assetPropertys.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETPROPERTYS)
	assetPropertys.Status = constants.ASSETPROPERTYSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnboardingCheckList.created")
	assetPropertys.Created = created
	log.Println("b4 OnboardingCheckList.created")

	dberr := s.Daos.SaveAssetPropertys(ctx, assetPropertys)
	if dberr != nil {
		return dberr
	}
	return nil

}
