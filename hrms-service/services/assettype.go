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

//SaveAssetType :""
func (s *Service) SaveAssetType(ctx *models.Context, assetType *models.AssetType) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	assetType.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETTYPE)
	assetType.Status = constants.ASSETTYPESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AssetType.created")
	assetType.Created = created
	log.Println("b4 AssetType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAssetType(ctx, assetType)
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
func (s *Service) SaveAssetTypeWithPropertys(ctx *models.Context, assetType *models.AssetType) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	assetType.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETTYPE)
	assetType.Status = constants.ASSETTYPESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 AssetType.created")
	assetType.Created = created
	log.Println("b4 AssetType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAssetTypeWithPropertys(ctx, assetType)
		if dberr != nil {
			return dberr
		}
		if len(assetType.AssetTypePropertysId) > 0 {
			var assetTypeProperty []models.AssetTypePropertys
			for _, v := range assetType.AssetTypePropertysId {
				var assetTypePropertys models.AssetTypePropertys
				assetTypePropertys.Status = constants.ASSETTYPEPROPERTYSSTATUSACTIVE
				assetTypePropertys.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETTYPEPROPERTYS)
				assetTypePropertys.Name = v
				assetTypePropertys.AssetTypeID = assetType.UniqueID
				assetTypePropertys.OrganisationID = assetType.OrganisationID
				assetTypeProperty = append(assetTypeProperty, assetTypePropertys)
			}
			dberr = s.Daos.SaveAssetTypeProperty(ctx, assetTypeProperty)
			if dberr != nil {

				log.Println("err in abort out")
				return errors.New("Transaction Aborted - " + dberr.Error())
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

//GetSingleAssetType :""
func (s *Service) GetSingleAssetType(ctx *models.Context, UniqueID string) (*models.RefAssetType, error) {
	AssetType, err := s.Daos.GetSingleAssetType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return AssetType, nil
}

//UpdateAssetType : ""
func (s *Service) UpdateAssetType(ctx *models.Context, assetType *models.AssetType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAssetType(ctx, assetType)
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

//EnableAssetType : ""
func (s *Service) EnableAssetType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAssetType(ctx, UniqueID)
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

//DisableAssetType : ""
func (s *Service) DisableAssetType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAssetType(ctx, UniqueID)
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

//DeleteAssetType : ""
func (s *Service) DeleteAssetType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAssetType(ctx, UniqueID)
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

//FilterAssetType :""
func (s *Service) FilterAssetType(ctx *models.Context, assetTypeFilter *models.FilterAssetType, pagination *models.Pagination) ([]models.RefAssetType, error) {
	err := s.AssetTypeDataAccess(ctx, assetTypeFilter)
	if err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterAssetType(ctx, assetTypeFilter, pagination)

}
func (s *Service) AssetTypeDataAccess(ctx *models.Context, filter *models.FilterAssetType) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationID = append(filter.OrganisationID, v.UniqueID)
				}
			}

		}

	}
	return err
}

// func (s *Service) GetAssetTypePropertys(ctx *models.Context, UniqueID string) (*models.RefAssetType, error) {
// 	AssetType, err := s.Daos.GetAssetTypePropertys(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return AssetType, nil
// }
func (s *Service) UpdateAssetTypeWithProperty(ctx *models.Context, assetType *models.UpdateAssetType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	// for k, v := range assetType.AssetTypePropertys {
	// 	if v.UniqueID == "" {
	// 		assetType.AssetTypePropertys[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETTYPEPROPERTYS)
	// 	}
	// 	fmt.Println("AssetTypePropertys==>", v.UniqueID)

	// }
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.UpdateAssetType(ctx, &assetType.AssetType)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if assetType != nil {
			if len(assetType.AssetTypePropertys) > 0 {
				for _, v := range assetType.AssetTypePropertys {
					fmt.Println("inside for for loop ====>")
					if v.UniqueID == "" {
						v.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETTYPEPROPERTYS)
						v.OrganisationID = assetType.OrganisationID
						v.AssetTypeID = assetType.UniqueID
					}
					err := s.Daos.SaveAssetTypePropertysWithUpsert(ctx, &v)
					if err != nil {
						return errors.New("UpdateAssetTypePropertys Error" + dberr.Error())
					}
				}
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
