package services

import (
	"errors"
	"fmt"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveAsset :""
func (s *Service) SaveAsset(ctx *models.Context, asset *models.Asset) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	asset.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSET)
	asset.Status = constants.ASSETSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Asset.created")
	asset.Created = created
	log.Println("b4 Asset.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAsset(ctx, asset)
		if dberr != nil {
			return dberr
		}
		for k, v := range asset.AssetPropertysId {
			AssetProperty, err := s.Daos.GetSingleAssetTypePropertysWithAssetId(ctx, v.AssetPropertyId, asset.AssetTypeId)
			if err != nil {
				return err
			}
			asset.AssetPropertysId[k].AssetTypeID = asset.AssetTypeId
			asset.AssetPropertysId[k].AssetID = asset.UniqueID
			asset.AssetPropertysId[k].Name = AssetProperty.Name
			asset.AssetPropertysId[k].AssetPropertyId = v.AssetPropertyId
			asset.AssetPropertysId[k].OrganisationID = asset.OrganisationID
			asset.AssetPropertysId[k].Description = asset.Description
			err = s.SaveAssetPropertysWithoutTransaction(ctx, &asset.AssetPropertysId[k])
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

//GetSingleAsset :""
func (s *Service) GetSingleAsset(ctx *models.Context, UniqueID string) (*models.RefAsset, error) {
	Asset, err := s.Daos.GetSingleAsset(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Asset, nil
}

// GetSingleAssetUsingEmpID : ""
// func (s *Service) GetSingleAssetUsingEmpID(ctx *models.Context, UniqueID string) (*models.RefAsset, error) {
// 	Asset, err := s.Daos.GetSingleAssetUsingEmpID(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Asset, nil
// }

//UpdateAsset : ""
func (s *Service) UpdateAsset(ctx *models.Context, asset *models.Asset) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		var uniqueIds []string
		for k, v := range asset.AssetPropertysId {
			fmt.Println("AssetPropertysId===>", v.UniqueID)
			if v.UniqueID == "" {
				asset.AssetPropertysId[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETPROPERTYS)

			}
			uniqueIds = append(uniqueIds, v.UniqueID)

		}

		err := s.Daos.UpdateAsset(ctx, asset)
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		err = s.Daos.AssetPropertysRemoveNotPresentValue(ctx, asset.UniqueID, uniqueIds)
		if err != nil {
			return err
		}
		err = s.Daos.AssetPropertysUpsert(ctx, asset)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return errors.New("Transaction Aborted - " + err.Error())
	}
	return nil
}

//EnableAsset : ""
func (s *Service) EnableAsset(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAsset(ctx, UniqueID)
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

//DisableAsset : ""
func (s *Service) DisableAsset(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAsset(ctx, UniqueID)
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

//DeleteAsset : ""
func (s *Service) DeleteAsset(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAsset(ctx, UniqueID)
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

//FilterAsset :""
func (s *Service) FilterAsset(ctx *models.Context, assetFilter *models.FilterAsset, pagination *models.Pagination) ([]models.RefAsset, error) {
	err := s.AssetDataAccess(ctx, assetFilter)
	if err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterAsset(ctx, assetFilter, pagination)

}
func (s *Service) AssetDataAccess(ctx *models.Context, filter *models.FilterAsset) (err error) {
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

func (s *Service) AssetAssign(ctx *models.Context, asset *models.AssetAssign) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		refAsset, err := s.Daos.GetSingleAssetUsingUniqueId(ctx, asset.AssetId)
		if err != nil {
			return errors.New("error in getting the assetlog- " + err.Error())
		}
		fmt.Println("assetid=============", refAsset)
		if refAsset != nil {
			refAsset.EmployeeId = asset.EmployeeId
			refAsset.Status = constants.ASSETASSIGNSTATUS
			err = s.Daos.UpdateAsset(ctx, &refAsset.Asset)
			if err != nil {
				return errors.New("error in updating the assetlog" + err.Error())
			}
		}

		//Employee
		refAssetLog, err := s.Daos.GetSingleAssetLogUsingEmpID(ctx, asset.AssetId)
		if err != nil {
			return errors.New("error in getting the assetlog- " + err.Error())
		}

		if refAssetLog != nil {
			refAssetLog.Status = constants.ASSETREVOKESTATUS
			refAssetLog.EndDate = &t
			err = s.Daos.UpdateAssetLog(ctx, &refAssetLog.AssetLog)
			if err != nil {
				return errors.New("error in updating the assetlog" + err.Error())
			}
		}

		assetLog := new(models.AssetLog)
		//assetLog.Name = asset.Name
		assetLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONASSETLOG)
		//assetLog.OrganisationID = asset.OrganisationID
		assetLog.EmployeeId = asset.EmployeeId
		assetLog.Action.UserID = asset.AssignId
		assetLog.AssetId = asset.AssetId
		assetLog.Action.Date = &t
		assetLog.Status = constants.ASSETASSIGNSTATUS
		assetLog.Remark = asset.Remark
		//assetLog.IsLog = constants.ASSETASSIGNSTATUSYES
		// assetLog.AssetId = refAssetLog.AssetId
		assetLog.StartDate = &t
		err = s.Daos.SaveAssetLog(ctx, assetLog)
		if err != nil {
			return err

		}

		dberr := s.Daos.AssetAssign(ctx, asset)
		if dberr != nil {
			return dberr
		}
		apptoken, err := s.Daos.GetRegTokenWithParticulars(ctx, refAsset.EmployeeId)
		if err != nil {
			return err
		}
		if apptoken != nil {
			fmt.Println("apptoken===>", apptoken.RegistrationToken)
			var token []string
			token = append(token, apptoken.RegistrationToken)

			fmt.Println("appToken===>", apptoken.RegistrationToken)
			topic := ""
			tittle := "Asset -" + refAsset.Name + "Asset Assign"
			Body := refAsset.Name
			//	var image string
			//if len(employeeTimeOff.) > 0 {
			image := ""
			//	}
			data := make(map[string]string)
			data["notificationType"] = "ViewAsset"
			data["id"] = refAsset.UniqueID
			err := s.SendNotification(topic, tittle, Body, image, token, data)
			if err != nil {
				log.Println(apptoken.RegistrationToken + " " + err.Error())
			}
			if err == nil {
				t := time.Now()
				ToNotificationLog := new(models.ToNotificationLog)
				notifylog := new(models.NotificationLog)
				ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
				ToNotificationLog.Name = refAsset.Name
				ToNotificationLog.UserName = refAsset.EmployeeId
				ToNotificationLog.UserType = "Employee"
				notifylog.Body = Body
				notifylog.Tittle = tittle
				notifylog.Topic = topic
				notifylog.Image = image
				notifylog.IsJob = false
				notifylog.Message = Body
				notifylog.SentDate = &t
				notifylog.SentFor = topic
				notifylog.Data = data
				notifylog.Status = "Active"
				notifylog.To = *ToNotificationLog
				err = s.Daos.SaveNotificationLog(ctx, notifylog)
				if err != nil {
					return err
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

func (s *Service) RevokeAsset(ctx *models.Context, asset *models.Asset) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RevokeAsset(ctx, asset)
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
func (s *Service) AssetReportExcel(ctx *models.Context, filter *models.FilterAsset) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterAsset(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "AssetReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "F1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	// title :=
	loc, _ := time.LoadLocation("Asia/Kolkata")
	currentTime := time.Now().In(loc)
	day := fmt.Sprintf("%v/%v/%v", currentTime.Day(), currentTime.Month(), currentTime.Year())
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v-%v", sheet1, day))
	//Kitchen := "3:04PM"
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "AssetId")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "AssetType")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "AssignedEmployee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "AssignedOn")

	rowNo++
	for k, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Ref.AssetTypeId.Name)
		if v.Ref.Employee.Name != "" {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.Employee.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Assigned")

		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "NA")
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Not Assigned")

		}
		rowNo++
	}
	//	var totalAmount float64

	//excel.MergeCell(sheet1, "A", "D")
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
