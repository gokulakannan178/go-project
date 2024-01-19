package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveAssetPolicyAssets :""
func (s *Service) SaveAssetPolicyAssets(ctx *models.Context, assetPolicyAssets *models.AssetPolicyAssets) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	assetPolicyAssets.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETPOLICYASSETS)
	assetPolicyAssets.Status = constants.ASSETPOLICYASSETSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AssetPolicyAssets.created")
	assetPolicyAssets.Created = created
	log.Println("b4 AssetPolicyAssets.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAssetPolicyAssets(ctx, assetPolicyAssets)
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

//SaveAssetPolicyAssets :""
func (s *Service) SaveAssetPolicyAssetsWithoutTransaction(ctx *models.Context, assetPolicyAssets *models.AssetPolicyAssets) error {
	log.Println("transaction start")
	//Start Transaction
	assetPolicyAssets.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETPOLICYASSETS)
	assetPolicyAssets.Status = constants.ASSETPOLICYASSETSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AssetPolicyAssets.created")
	assetPolicyAssets.Created = created
	log.Println("b4 AssetPolicyAssets.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAssetPolicyAssets(ctx, assetPolicyAssets)
		if dberr != nil {
			return dberr
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

//GetSingleAssetPolicyAssets :""
func (s *Service) GetSingleAssetPolicyAssets(ctx *models.Context, UniqueID string) (*models.RefAssetPolicyAssets, error) {
	assetPolicyAssets, err := s.Daos.GetSingleAssetPolicyAssets(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return assetPolicyAssets, nil
}

//UpdateAssetPolicyAssets : ""
func (s *Service) UpdateAssetPolicyAssets(ctx *models.Context, assetPolicyAssets *models.AssetPolicyAssets) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAssetPolicyAssets(ctx, assetPolicyAssets)
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

//EnableAssetPolicyAssets : ""
func (s *Service) EnableAssetPolicyAssets(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAssetPolicyAssets(ctx, UniqueID)
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

//DisableAssetPolicyAssets : ""
func (s *Service) DisableAssetPolicyAssets(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAssetPolicyAssets(ctx, UniqueID)
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

//DeleteAssetPolicyAssets : ""
func (s *Service) DeleteAssetPolicyAssets(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAssetPolicyAssets(ctx, UniqueID)
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

//FilterAssetPolicyAssets :""
func (s *Service) FilterAssetPolicyAssets(ctx *models.Context, assetPolicyAssetsFilter *models.FilterAssetPolicyAssets, pagination *models.Pagination) ([]models.RefAssetPolicyAssets, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterAssetPolicyAssets(ctx, assetPolicyAssetsFilter, pagination)

}
