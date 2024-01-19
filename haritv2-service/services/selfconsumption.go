package services

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSelfConsumption : ""
func (s *Service) SaveSelfConsumption(ctx *models.Context, selfconsumption *models.SelfConsumption) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.SaveSelfConsumptionWithoutTransaction(ctx, selfconsumption)
		if err != nil {
			return err
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

//GetSingleSelfConsumption :""
func (s *Service) GetSingleSelfConsumption(ctx *models.Context, UniqueID string) (*models.RefSelfConsumption, error) {
	tower, err := s.Daos.GetSingleSelfConsumption(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateSelfConsumption : ""
func (s *Service) UpdateSelfConsumption(ctx *models.Context, selfconsumption *models.SelfConsumption) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateSelfConsumption(ctx, selfconsumption)
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

// EnableSelfConsumption : ""
func (s *Service) EnableSelfConsumption(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableSelfConsumption(ctx, UniqueID)
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

//DisableSelfConsumption : ""
func (s *Service) DisableSelfConsumption(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableSelfConsumption(ctx, UniqueID)
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

//DeleteSelfConsumption : ""
func (s *Service) DeleteSelfConsumption(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSelfConsumption(ctx, UniqueID)
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

// FilterSelfConsumption : ""
func (s *Service) FilterSelfConsumption(ctx *models.Context, filter *models.SelfConsumptionFilter, pagination *models.Pagination) ([]models.RefSelfConsumption, error) {
	return s.Daos.FilterSelfConsumption(ctx, filter, pagination)

}
func (s *Service) DecreaseInventoryForULBandFPO(ctx *models.Context, selfconsumption *models.SelfConsumption) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.SaveSelfConsumptionWithoutTransaction(ctx, selfconsumption)
		if err != nil {
			return err
		}

		// Decreasing ULB Inventory Collection
		switch selfconsumption.Type {
		case "ULB":
			ULBInventory, err := s.Daos.GetSingleULBInventoryWithCompalyID(ctx, selfconsumption.CompanyID)
			if err != nil {
				return errors.New("Error in finding ULB Inventory customer - " + err.Error())
			}
			fmt.Println("ulb inventory found")

			if ULBInventory != nil {
				ULBInventory.Quantity = ULBInventory.Quantity - selfconsumption.Quantity
				err := s.Daos.UpdateULBInventoryDeliverSale(ctx, ULBInventory)
				if err != nil {
					return errors.New("Unable to Update quantity in ULB" + err.Error())
				}
			}
			fmt.Println("updated ulb inventory")

			// sale := refConsumption.CompanyID
			// sale.ID = primitive.NewObjectID()
			// sale.Status = constants.SALESTATUSACTIVE
			// sale.Transport.Status = constants.SALETRANSPORTSTATUSDELIVERED
			// sale.PaymentStatus = constants.SALEPAYMENTSTATUSCOMPLETED
			// err = s.Daos.SaveSale(ctx, &sale)
			// if err != nil {
			// 	return errors.New("Error in saving sale" + err.Error())
			// }
			// fmt.Println("sale saved")
		case "FPO":
			FPOInventory, err := s.Daos.GetSingleFPOInventoryWithCompalyID(ctx, selfconsumption.CompanyID)
			if err != nil {
				return errors.New("Error in finding FPO Inventory customer - " + err.Error())
			}
			fmt.Println("FPO inventory found")

			if FPOInventory != nil {
				FPOInventory.Quantity = FPOInventory.Quantity - selfconsumption.Quantity
				err := s.Daos.UpdateFPOInventoryDeliverSale(ctx, FPOInventory)
				if err != nil {
					return errors.New("Unable to Update quantity in ULB" + err.Error())
				}
			}
			fmt.Println("updated FPO inventory")

			// sale := refConsumption.CompanyID
			// sale.ID = primitive.NewObjectID()
			// sale.Status = constants.SALESTATUSACTIVE
			// sale.Transport.Status = constants.SALETRANSPORTSTATUSDELIVERED
			// sale.PaymentStatus = constants.SALEPAYMENTSTATUSCOMPLETED
			// err = s.Daos.SaveSale(ctx, &sale)
			// if err != nil {
			// 	return errors.New("Error in saving sale" + err.Error())
			// }
			// fmt.Println("sale saved")

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
func (s *Service) SaveSelfConsumptionWithoutTransaction(ctx *models.Context, selfconsumption *models.SelfConsumption) error {
	log.Println("transaction start")
	selfconsumption.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSELFCONSUMPTION)
	selfconsumption.Status = constants.SELFCONSUMPTIONSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	selfconsumption.Date = &t
	created.By = constants.SYSTEM
	selfconsumption.Created = created
	dberr := s.Daos.SaveSelfConsumption(ctx, selfconsumption)
	if dberr != nil {
		return dberr
	}
	return nil
}

func (s *Service) SelfConsumptionExcel(ctx *models.Context, filter *models.SelfConsumptionFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterSelfConsumption(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "SelfConsumption Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "E1")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "#")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Quantity")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Status")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Quantity)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ref.ULBID.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Date)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.By)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.ULBID.Status)
		rowNo++
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
