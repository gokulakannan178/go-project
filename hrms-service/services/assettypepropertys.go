package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveAssetTypePropertys :""
func (s *Service) SaveAssetTypePropertys(ctx *models.Context, assetTypePropertys *models.AssetTypePropertys) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	assetTypePropertys.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETTYPEPROPERTYS)
	assetTypePropertys.Status = constants.ASSETTYPEPROPERTYSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AssetTypePropertys.created")
	assetTypePropertys.Created = created
	log.Println("b4 AssetTypePropertys.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAssetTypePropertys(ctx, assetTypePropertys)
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

//GetSingleAssetTypePropertys :""
func (s *Service) GetSingleAssetTypePropertys(ctx *models.Context, UniqueID string) (*models.RefAssetTypePropertys, error) {
	AssetTypePropertys, err := s.Daos.GetSingleAssetTypePropertys(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return AssetTypePropertys, nil
}

//UpdateAssetTypePropertys : ""
func (s *Service) UpdateAssetTypePropertys(ctx *models.Context, assetTypePropertys *models.AssetTypePropertys) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAssetTypePropertys(ctx, assetTypePropertys)
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

//EnableAssetTypePropertys : ""
func (s *Service) EnableAssetTypePropertys(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAssetTypePropertys(ctx, UniqueID)
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

//DisableAssetTypePropertys : ""
func (s *Service) DisableAssetTypePropertys(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAssetTypePropertys(ctx, UniqueID)
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

//DeleteAssetTypePropertys : ""
func (s *Service) DeleteAssetTypePropertys(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAssetTypePropertys(ctx, UniqueID)
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

//FilterAssetTypePropertys :""
func (s *Service) FilterAssetTypePropertys(ctx *models.Context, assetTypePropertysFilter *models.FilterAssetTypePropertys, pagination *models.Pagination) ([]models.RefAssetTypePropertys, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterAssetTypePropertys(ctx, assetTypePropertysFilter, pagination)

}
func (s *Service) SaveAssetTypePropertysWithoutTransaction(ctx *models.Context, assetTypePropertys *models.AssetTypePropertys) error {
	log.Println("transaction start")
	// Start Transaction

	assetTypePropertys.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETTYPEPROPERTYS)
	assetTypePropertys.Status = constants.ASSETTYPEPROPERTYSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnboardingCheckList.created")
	assetTypePropertys.Created = created
	log.Println("b4 OnboardingCheckList.created")

	dberr := s.Daos.SaveAssetTypePropertys(ctx, assetTypePropertys)
	if dberr != nil {
		return dberr
	}
	return nil

}
