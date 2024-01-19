package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveLegacy :""
func (s *Service) SaveLegacy(ctx *models.Context, legacy *models.RegLegacyProperty) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	s.PreSaveLegacy(ctx, legacy)
	// legacy.LegacyProperty.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEGACY)

	// //bankDeposit.Status = constants.BANKDEPOSITSTATUSPENDING
	// t := time.Now()
	// created := models.CreatedV2{}
	// created.On = &t
	// created.By = legacy.CreatedBy
	// created.ByType = legacy.CreatedType
	// log.Println("b4 user.created")
	// legacy.LegacyProperty.Created = created
	// legacy.LegacyProperty.Status = constants.LEGACYPROPERTYSTATUSACTIVE
	// log.Println("b4 user.created")
	// for k := range legacy.LegacyPropertyFy {
	// 	legacy.LegacyPropertyFy[k].PropertyID = legacy.LegacyProperty.PropertyID
	// 	legacy.LegacyPropertyFy[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEGACYYEAR)
	// 	legacy.LegacyPropertyFy[k].LegacyID = legacy.LegacyProperty.UniqueID
	// 	legacy.LegacyPropertyFy[k].Status = constants.LEGACYPROPERTYFYSTATUSACTIVE
	// }
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveLegacy(ctx, legacy)
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

// PreSaveLegacy : ""
func (s *Service) PreSaveLegacy(ctx *models.Context, legacy *models.RegLegacyProperty) {

	legacy.LegacyProperty.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEGACY)

	//bankDeposit.Status = constants.BANKDEPOSITSTATUSPENDING
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = legacy.CreatedBy
	created.ByType = legacy.CreatedType
	log.Println("b4 user.created")
	legacy.LegacyProperty.Created = created
	legacy.LegacyProperty.Status = constants.LEGACYPROPERTYSTATUSACTIVE
	log.Println("b4 user.created")
	for k := range legacy.LegacyPropertyFy {
		legacy.LegacyPropertyFy[k].PropertyID = legacy.LegacyProperty.PropertyID
		legacy.LegacyPropertyFy[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEGACYYEAR)
		legacy.LegacyPropertyFy[k].LegacyID = legacy.LegacyProperty.UniqueID
		legacy.LegacyPropertyFy[k].Status = constants.LEGACYPROPERTYFYSTATUSACTIVE
	}

}

// GetLegacyForAProperty : ""
func (s *Service) GetLegacyForAProperty(ctx *models.Context, propertyID string) (*models.RefLegacyPropertyPayment, error) {
	return s.Daos.GetLegacyForAProperty(ctx, propertyID)
}

func (s *Service) UpdateLegacyForAProperty(ctx *models.Context, legacy *models.RegLegacyProperty) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for k := range legacy.LegacyPropertyFy {
			if legacy.LegacyPropertyFy[k].UniqueID == "" {
				legacy.LegacyPropertyFy[k].PropertyID = legacy.LegacyProperty.PropertyID
				legacy.LegacyPropertyFy[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEGACYYEAR)
				legacy.LegacyPropertyFy[k].LegacyID = legacy.LegacyProperty.UniqueID
				legacy.LegacyPropertyFy[k].Status = constants.LEGACYPROPERTYFYSTATUSACTIVE
			}
		}
		err := s.Daos.UpdateLegacyForAProperty(ctx, legacy)
		if err != nil {
			return nil
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

//GetFinancialYearsForLegacyPayments : ""
func (s *Service) GetFinancialYearsForLegacyPayments(ctx *models.Context, propertyId string) ([]models.RefV2LegacyPropertyFy, error) {
	return s.Daos.GetFinancialYearsForLegacyPayments(ctx, propertyId)
}

//GetReqFinancialYearForLegacy : ""
func (s *Service) GetReqFinancialYearForLegacy(ctx *models.Context, grfy *models.GetReqFinancialYear) ([]models.RefFinancialYear, error) {
	return s.Daos.GetReqFinancialYearForLegacy(ctx, grfy.Doa)
}

// FilterLegacy : ""
func (s *Service) FilterLegacy(ctx *models.Context, filter *models.LegacyPropertyFilter, pagination *models.Pagination) ([]models.RefLegacyPropertyPayment, error) {
	return s.Daos.FilterLegacy(ctx, filter, pagination)

}
func (s *Service) LegacyReportExcel(ctx *models.Context, filter *models.LegacyPropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterLegacy(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Legacy Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "E3")
		excel.MergeCell(sheet1, "A4", "E5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "E3")
		excel.MergeCell(sheet1, "C4", "E5")
	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v", "S.No"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v", "PropertyId"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", "Finacial Year"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v", "Payment Date"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v", "Tax Amount"))
	rowNo++

	var totalAmount float64
	for _, res := range data {
		sno := 1
		for _, v := range res.LegacyPropertyFy {

			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sno)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
				if res.PropertyID != "" {
					return res.PropertyID
				}
				return "NA"
			}())

			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
				if v.Ref.Fy != nil {
					return v.Ref.Fy.Name
				}
				return "NA"
			}())
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
				if res.PaymentDate != nil {
					return res.PaymentDate.Format("2006-01-02")
				}
				return "NA"
			}())
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.TaxAmount)

			totalAmount = totalAmount + v.TaxAmount
			rowNo++
			sno++
		}

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}
