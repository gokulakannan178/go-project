package service

// import (
// 	"errors"
// 	"fmt"

// 	"lgf-ccc-service/constants"
// 	"lgf-ccc-service/models"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// //SaveAsset :""
// func (s *Service) SaveAsset(ctx *models.Context, asset *models.Asset) error {
// 	log.Println("transaction start")
// 	//Start Transaction
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	asset.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSET)
// 	asset.Status = constants.ASSETSTATUSACTIVE
// 	t := time.Now()
// 	created := models.Created{}
// 	created.On = &t
// 	created.By = constants.SYSTEM
// 	log.Println("b4 Asset.created")
// 	asset.Created = created
// 	log.Println("b4 Asset.created")
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		dberr := s.Daos.SaveAsset(ctx, asset)
// 		if dberr != nil {
// 			return dberr
// 		}
// 		for k, v := range asset.AssetPropertysId {
// 			//assetpropertys := new(models.AssetPropertys)
// 			// assetTypeProperty, err := s.Daos.GetSingleAssetTypePropertysWithActive(ctx, v, constants.ASSETTYPEPROPERTYSSTATUSACTIVE)
// 			// if err != nil {
// 			// 	fmt.Println(err)
// 			// 	return err
// 			// }

// 			//fmt.Println("onboardingchecklistmaster=======", assetTypeProperty)

// 			asset.AssetPropertysId[k].AssetTypeID = asset.AssetTypeId
// 			asset.AssetPropertysId[k].AssetID = asset.UniqueID
// 			asset.AssetPropertysId[k].Name = asset.Name
// 			asset.AssetPropertysId[k].AssetPropertyId = v.AssetPropertyId
// 			asset.AssetPropertysId[k].OrganisationID = asset.OrganisationID
// 			asset.AssetPropertysId[k].Description = asset.Description
// 			err := s.SaveAssetPropertysWithoutTransaction(ctx, &asset.AssetPropertysId[k])
// 			if err != nil {
// 				return err
// 			}

// 		}
// 		if err := ctx.Session.CommitTransaction(sc); err != nil {
// 			return errors.New("Not able to commit - " + err.Error())
// 		}
// 		return nil
// 	}); err != nil {
// 		log.Println("Transaction start aborting")
// 		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
// 			return errors.New("Error while aborting transaction" + abortError.Error())
// 		}
// 		log.Println("Transaction aborting completed successfully")
// 		return err
// 	}
// 	return nil
// }

// //GetSingleAsset :""
// func (s *Service) GetSingleAsset(ctx *models.Context, UniqueID string) (*models.RefAsset, error) {
// 	Asset, err := s.Daos.GetSingleAsset(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Asset, nil
// }

// // GetSingleAssetUsingEmpID : ""
// // func (s *Service) GetSingleAssetUsingEmpID(ctx *models.Context, UniqueID string) (*models.RefAsset, error) {
// // 	Asset, err := s.Daos.GetSingleAssetUsingEmpID(ctx, UniqueID)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return Asset, nil
// // }

// //UpdateAsset : ""
// func (s *Service) UpdateAsset(ctx *models.Context, asset *models.Asset) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		var uniqueIds []string
// 		for k, v := range asset.AssetPropertysId {
// 			fmt.Println("AssetPropertysId===>", v.UniqueID)
// 			if v.UniqueID == "" {
// 				asset.AssetPropertysId[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETPROPERTYS)

// 			}
// 			uniqueIds = append(uniqueIds, v.UniqueID)

// 		}

// 		err := s.Daos.UpdateAsset(ctx, asset)
// 		if err := ctx.Session.CommitTransaction(sc); err != nil {
// 			return errors.New("Not able to commit - " + err.Error())
// 		}
// 		err = s.Daos.AssetPropertysRemoveNotPresentValue(ctx, asset.UniqueID, uniqueIds)
// 		if err != nil {
// 			return err
// 		}
// 		err = s.Daos.AssetPropertysUpsert(ctx, asset)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}); err != nil {
// 		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
// 			return errors.New("Transaction Aborted with error" + err.Error())
// 		}
// 		return errors.New("Transaction Aborted - " + err.Error())
// 	}
// 	return nil
// }

// //EnableAsset : ""
// func (s *Service) EnableAsset(ctx *models.Context, UniqueID string) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.EnableAsset(ctx, UniqueID)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// //DisableAsset : ""
// func (s *Service) DisableAsset(ctx *models.Context, UniqueID string) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.DisableAsset(ctx, UniqueID)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// //DeleteAsset : ""
// func (s *Service) DeleteAsset(ctx *models.Context, UniqueID string) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.DeleteAsset(ctx, UniqueID)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// //FilterAsset :""
// func (s *Service) FilterAsset(ctx *models.Context, assetFilter *models.FilterAsset, pagination *models.Pagination) ([]models.RefAsset, error) {
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	return s.Daos.FilterAsset(ctx, assetFilter, pagination)

// }
// func (s *Service) AssetAssign(ctx *models.Context, asset *models.AssetAssign) error {
// 	log.Println("transaction start")
// 	//Start Transaction
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	t := time.Now()
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		refAsset, err := s.Daos.GetSingleAssetUsingUniqueId(ctx, asset.AssetId)
// 		if err != nil {
// 			return errors.New("error in getting the assetlog- " + err.Error())
// 		}
// 		fmt.Println("assetid=============", refAsset)
// 		if refAsset != nil {
// 			refAsset.EmployeeId = asset.EmployeeId
// 			refAsset.Status = constants.ASSETASSIGNSTATUS
// 			err = s.Daos.UpdateAsset(ctx, &refAsset.Asset)
// 			if err != nil {
// 				return errors.New("error in updating the assetlog" + err.Error())
// 			}
// 		}

// 		//Employee
// 		refAssetLog, err := s.Daos.GetSingleAssetLogUsingEmpID(ctx, asset.AssetId)
// 		if err != nil {
// 			return errors.New("error in getting the assetlog- " + err.Error())
// 		}

// 		if refAssetLog != nil {
// 			refAssetLog.Status = constants.ASSETREVOKESTATUS
// 			refAssetLog.EndDate = &t
// 			err = s.Daos.UpdateAssetLog(ctx, &refAssetLog.AssetLog)
// 			if err != nil {
// 				return errors.New("error in updating the assetlog" + err.Error())
// 			}
// 		}

// 		assetLog := new(models.AssetLog)
// 		//assetLog.Name = asset.Name
// 		assetLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETLOG)
// 		//assetLog.OrganisationID = asset.OrganisationID
// 		assetLog.EmployeeId = asset.EmployeeId
// 		assetLog.Action.UserID = asset.AssignId
// 		assetLog.AssetId = asset.AssetId
// 		assetLog.Action.Date = &t
// 		assetLog.Status = constants.ASSETASSIGNSTATUS
// 		assetLog.Remark = asset.Remark
// 		//assetLog.IsLog = constants.ASSETASSIGNSTATUSYES
// 		// assetLog.AssetId = refAssetLog.AssetId
// 		assetLog.StartDate = &t
// 		err = s.Daos.SaveAssetLog(ctx, assetLog)
// 		if err != nil {
// 			return err

// 		}
// 		// if refAssetLog == nil {

// 		// }
// 		dberr := s.Daos.AssetAssign(ctx, asset)
// 		if dberr != nil {
// 			return dberr
// 		}
// 		if err := ctx.Session.CommitTransaction(sc); err != nil {
// 			return errors.New("Not able to commit - " + err.Error())
// 		}
// 		return nil
// 	}); err != nil {
// 		log.Println("Transaction start aborting")
// 		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
// 			return errors.New("Error while aborting transaction" + abortError.Error())
// 		}
// 		log.Println("Transaction aborting completed successfully")
// 		return err
// 	}

// 	return nil
// }

// func (s *Service) RevokeAsset(ctx *models.Context, asset *models.Asset) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.RevokeAsset(ctx, asset)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }
