package service

import (
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

// SaveHouseVisited : ""
func (s *Service) SaveHouseVisited(ctx *models.Context, HouseVisited *models.HouseVisited) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	HouseVisited.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONHOUSEVISITED)
	if HouseVisited.IsAvailable == "C" {
		HouseVisited.IsStatus = constants.HOUSEVISITEDSTATUSCOLLECTED
	}
	if HouseVisited.IsAvailable == "NA" {
		HouseVisited.IsStatus = constants.HOUSEVISITEDSTATUSNOTAVAILABLE
	}
	HouseVisited.Status = constants.HOUSEVISITEDSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 HouseVisited.created")
	HouseVisited.Created = &created
	HouseVisited.Date = &t
	log.Println("b4 HouseVisited.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		property, err := s.Daos.GetSinglePropertiesWithHouseUID(ctx, HouseVisited.HouseUID)
		if err != nil {
			return err
		}
		if property == nil {
			return err
		}
		fmt.Println("property=======>", property)
		HouseVisited.Property.Id = property.UniqueID
		HouseVisited.Property.Name = property.Name
		HouseVisited.Property.HoldingNumber = property.HoldingNumber
		HouseVisited.CircleNo = property.CircleCode
		HouseVisited.WardNo = property.WardCode
		HouseVisited.SectorCode = property.SectorCode
		dberr := s.Daos.SaveHouseVisited(ctx, HouseVisited)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		totalproperty, err := s.Daos.GetPropertiesCountWithWard(ctx, HouseVisited.WardNo)
		if err != nil {
			return err
		}
		fmt.Println("totalproperty=========>", totalproperty)
		tpcw, err := s.Daos.GetPropertiesCountWithCircle(ctx, HouseVisited.CircleNo)
		if err != nil {
			return err
		}
		fmt.Println("tpcw=========>", tpcw)

		datestr := HouseVisited.Date.Format("2006-01-02")
		wardwisehousevisited := new(models.WardWiseHouseVisited)
		wwhv, err := s.Daos.GetSingleWardWiseHouseVisitedWithDate(ctx, datestr, HouseVisited.WardNo)
		if err != nil {
			return err
		}
		if wwhv == nil {
			wardwisehousevisited.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWARDWISEHOUSEVISITED)
			wardwisehousevisited.CircleCode = HouseVisited.CircleNo
			wardwisehousevisited.WardCode = HouseVisited.WardNo
			wardwisehousevisited.Date = HouseVisited.Date
			wardwisehousevisited.Datestr = datestr
			wardwisehousevisited.Status = constants.WARDWISEHOUSEVISITEDSTATUSACTIVE
			if totalproperty != nil {
				wardwisehousevisited.TotalProperties = totalproperty.Quantity
			}
			if totalproperty == nil {
				wardwisehousevisited.TotalProperties = 0
			}
			wardwisehousevisited.TodayCollection = 1
			dberr := s.Daos.SaveWardWiseHouseVisited(ctx, wardwisehousevisited)
			if dberr != nil {
				return dberr
			}
		} else {
			wardwisehousevisited.CircleCode = HouseVisited.CircleNo
			wardwisehousevisited.WardCode = HouseVisited.WardNo
			wardwisehousevisited.Date = HouseVisited.Date
			wardwisehousevisited.Datestr = datestr
			dberr := s.Daos.IncreaseWardWiseHouseVisitedCount(ctx, wardwisehousevisited.Datestr, wardwisehousevisited.WardCode)
			if dberr != nil {
				return dberr
			}
		}
		circlehousevisited := new(models.CircleWiseHouseVisited)
		cwhv, err := s.Daos.GetSingleCircleWiseHouseVisitedWithDate(ctx, datestr, HouseVisited.CircleNo)
		if err != nil {
			return err
		}
		if cwhv == nil {
			circlehousevisited.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCIRCLEWISEHOUSEVISITED)
			circlehousevisited.CircleCode = HouseVisited.CircleNo
			circlehousevisited.Date = HouseVisited.Date
			circlehousevisited.Datestr = datestr
			circlehousevisited.Status = constants.WARDWISEHOUSEVISITEDSTATUSACTIVE
			circlehousevisited.TodayCollection = 1
			if tpcw != nil {
				circlehousevisited.TotalProperties = tpcw.Quantity

			}
			if tpcw == nil {
				circlehousevisited.TotalProperties = 0

			}
			dberr := s.Daos.SaveCircleWiseHouseVisited(ctx, circlehousevisited)
			if dberr != nil {
				return dberr
			}
		} else {
			circlehousevisited.CircleCode = HouseVisited.CircleNo
			circlehousevisited.Date = HouseVisited.Date
			circlehousevisited.Datestr = datestr
			dberr := s.Daos.IncreaseCircleWiseHouseVisitedCount(ctx, circlehousevisited.Datestr, circlehousevisited.CircleCode)
			if dberr != nil {
				return dberr
			}
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

// GetSingleHouseVisited : ""
func (s *Service) GetSingleHouseVisited(ctx *models.Context, UniqueID string) (*models.RefHouseVisited, error) {
	HouseVisited, err := s.Daos.GetSingleHouseVisited(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return HouseVisited, nil
}

//UpdateHouseVisited : ""
func (s *Service) UpdateHouseVisited(ctx *models.Context, HouseVisited *models.HouseVisited) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateHouseVisited(ctx, HouseVisited)
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

// EnableHouseVisited : ""
func (s *Service) EnableHouseVisited(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableHouseVisited(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableHouseVisited : ""
func (s *Service) DisableHouseVisited(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableHouseVisited(ctx, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}

//DeleteHouseVisited : ""
func (s *Service) DeleteHouseVisited(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteHouseVisited(ctx, UniqueID)
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

// FilterHouseVisited : ""
func (s *Service) FilterHouseVisited(ctx *models.Context, HouseVisited *models.FilterHouseVisited, pagination *models.Pagination) (HouseVisiteds []models.RefHouseVisited, err error) {
	return s.Daos.FilterHouseVisited(ctx, HouseVisited, pagination)
}

func (s *Service) CollectedHouseVisited(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.CollectedHouseVisited(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

func (s *Service) HouseVisitedReportExcel(ctx *models.Context, filter *models.FilterHouseVisited, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterHouseVisited(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "HouseVisitedReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "I1")
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
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	rowNo++
	//
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "WardNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "HoldingNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "CitizenName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "CollectorName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "ManagerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Status")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Remarks")
	rowNo++
	//	var totalAmount float64
	for k, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Date.Format("02-January-2006"))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.WardNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Property.HoldingNumber)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Property.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.GCUser.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.MinUser.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Status)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Remark)
		rowNo++
	}
	//excel.MergeCell(sheet1, "A", "D")
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	a := fmt.Sprintf("A%v", rowNo)
	c := fmt.Sprintf("D%v", rowNo)
	excel.MergeCell(sheet1, a, c)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Total")

	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
