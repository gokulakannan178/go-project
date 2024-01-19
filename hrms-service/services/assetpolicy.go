package services

import (
	"errors"
	"fmt"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveAssetPolicy :""
func (s *Service) SaveAssetPolicy(ctx *models.Context, assetPolicy *models.AssetPolicy) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	assetPolicy.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETPOLICY)
	assetPolicy.Status = constants.ASSETPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AssetPolicy.created")
	assetPolicy.Created = created
	log.Println("b4 AssetPolicy.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAssetPolicy(ctx, assetPolicy)
		if dberr != nil {
			return dberr
		}

		refAssetPolicy, err := s.Daos.GetSingleAssetPolicy(ctx, assetPolicy.UniqueID)
		if err != nil {
			return err
		}
		fmt.Println("refAssetPolicy", refAssetPolicy.AssetMasterId)

		for _, v := range assetPolicy.AssetMasterId {
			assetPolicyAssets := new(models.AssetPolicyAssets)
			documentmaster, err := s.Daos.GetSingleDocumentMasterWithActive(ctx, v, constants.DOCUMENTMASTERSTATUSACTIVE)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println("documentmaster", documentmaster)
			if refAssetPolicy != nil {
				assetPolicyAssets.AssetPolicyID = refAssetPolicy.UniqueID
			}
			assetPolicyAssets.AssetMasterID = v
			assetPolicyAssets.Name = assetPolicy.Name
			err = s.SaveAssetPolicyAssetsWithoutTransaction(ctx, assetPolicyAssets)
			if err != nil {
				return err
			}
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

//GetSingleAssetPolicy :""
func (s *Service) GetSingleAssetPolicy(ctx *models.Context, UniqueID string) (*models.RefAssetPolicy, error) {
	assetPolicy, err := s.Daos.GetSingleAssetPolicy(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return assetPolicy, nil
}

//UpdateAssetPolicy : ""
func (s *Service) UpdateAssetPolicy(ctx *models.Context, assetPolicy *models.AssetPolicy) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.AssetPolicyAssetsRemoveNotPresentValue(ctx, assetPolicy.UniqueID, assetPolicy.AssetMasterId)
		if err != nil {
			return err
		}
		err = s.Daos.AssetPolicyAssetsUpsert(ctx, assetPolicy.UniqueID, assetPolicy.AssetMasterId, assetPolicy.Name)
		if err != nil {
			return err
		}

		fmt.Println("error==>", err)

		err = s.Daos.UpdateAssetPolicy(ctx, assetPolicy)
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

//EnableAssetPolicy : ""
func (s *Service) EnableAssetPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAssetPolicy(ctx, UniqueID)
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

//DisableAssetPolicy : ""
func (s *Service) DisableAssetPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAssetPolicy(ctx, UniqueID)
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

//DeleteAssetPolicy : ""
func (s *Service) DeleteAssetPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAssetPolicy(ctx, UniqueID)
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

//FilterAssetPolicy :""
func (s *Service) FilterAssetPolicy(ctx *models.Context, assetPolicyFilter *models.FilterAssetPolicy, pagination *models.Pagination) ([]models.RefAssetPolicy, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterAssetPolicy(ctx, assetPolicyFilter, pagination)

}
