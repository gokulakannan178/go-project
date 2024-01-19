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

// InitiateTradeLicensePayment : ""
func (s *Service) InitiateTradeLicensePayment(ctx *models.Context, ipmtr *models.InitiateTradeLicensePaymentReq) (string, error) {
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	tnxId := ""
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("ipmtr", ipmtr)
		filter := new(models.TradeLicenseCalcQueryFilter)
		filter.TradeLicenseID = ipmtr.TradeLicenseID
		filter.AddFy = ipmtr.FYs
		demand, err := s.CalcTradeLicenseDemandForParticulars(ctx, filter)
		if err != nil {
			return errors.New("Error in calculating demand - " + err.Error())
		}
		if demand == nil {
			return errors.New("Demand is nil ")
		}

		pmt := new(models.TradeLicensePayments)
		pmt.TnxID = s.Shared.GetTransactionID(demand.TradeLicense.UniqueID, 32)

		pmt.ReciptNo = fmt.Sprintf("%v%v%v_%v_%v", "TL", t.Day(), int(t.Month()), t.Year(), s.Daos.GetUniqueID(ctx, "tl_recipt"))

		tnxId = pmt.TnxID
		pmt.TradeLicenseID = demand.TradeLicense.UniqueID

		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		pmt.FinancialYear = fy.FinancialYear

		pmt.Demand = models.TradeLicensePaymentDemand{
			Current: models.TradeLicensePaymentDemandSplitage{
				Tax:     demand.Demand.Current.Tax,
				Penalty: demand.Demand.Current.Penalty,
				Other:   demand.Demand.Current.Other,
				Total:   demand.Demand.Current.Total,
			},
			Arrear: models.TradeLicensePaymentDemandSplitage{
				Tax:     demand.Demand.Arrear.Tax,
				Penalty: demand.Demand.Arrear.Penalty,
				Other:   demand.Demand.Arrear.Other,
				Total:   demand.Demand.Arrear.Total,
			},
			Total: models.TradeLicensePaymentDemandSplitage{
				Tax:     demand.Demand.Total.Tax,
				Penalty: demand.Demand.Total.Penalty,
				Other:   demand.Demand.Total.Other,
				Total:   demand.Demand.Total.Total,
			},
		}
		pmt.Status = constants.SHOPRENTPAYMENTSTATUSINIT

		pmt.Created = models.CreatedV2{
			By:     ipmtr.By,
			ByType: ipmtr.ByType,
			On:     &t,
		}
		var pmtFys []models.TradeLicensePaymentsfY
		for _, v := range demand.FY {
			var pmtFy models.TradeLicensePaymentsfY
			pmtFy.TnxID = pmt.TnxID
			pmtFy.TradeLicenseID = demand.TradeLicense.UniqueID
			pmtFy.FY = v
			pmtFy.Status = constants.SHOPRENTPAYMENTSTATUSINIT
			pmtFys = append(pmtFys, pmtFy)
		}

		var pmtBasic = new(models.TradeLicensePaymentsBasics)
		pmtBasic.TnxID = pmt.TnxID
		pmtBasic.TradeLicenseID = demand.TradeLicense.UniqueID
		sr, err := s.Daos.GetSingleTradeLicense(ctx, ipmtr.TradeLicenseID)
		if err != nil {
			return err
		}
		pmt.Address = demand.TradeLicense.Address
		if sr != nil {
			pmtBasic.TradeLicense = *sr
		}
		err = s.Daos.SaveTradeLicensePayment(ctx, pmt)
		if err != nil {
			return errors.New("Error in saving mobile tower payment - " + err.Error())
		}
		err = s.Daos.SaveTradeLicensePaymentFYs(ctx, pmtFys)
		if err != nil {
			return errors.New("Error in saving mobile tower payment fys- " + err.Error())
		}
		err = s.Daos.SaveTradeLicensePaymentBasic(ctx, pmtBasic)
		if err != nil {
			return errors.New("Error in saving mobile tower payment basics- " + err.Error())
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return "", errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return "", err
	}
	return tnxId, nil
}

// GetSingleTradeLicensePayment : ""
func (s *Service) GetSingleTradeLicensePayment(ctx *models.Context, tnxID string) (*models.RefTradeLicensePayments, error) {
	return s.Daos.GetSingleTradeLicensePayment(ctx, tnxID)
}

// MakeTradeLicensePayment : ""
func (s *Service) MakeTradeLicensePayment(ctx *models.Context, mmtpr *models.MakeTradeLicensePaymentReq) (string, error) {
	tradeLicenseID := ""
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		mmtpr.CompletionDate = &t
		status, dbErr := s.TradeLicensePaymentStatusSelector(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}
		mmtpr.Status = status
		if ctx.ProductConfig.LocationID == "Bhagalpur" {
			mmtpr.Status = constants.TRADELICENSEPAYMENRSTATUSCOMPLETED
		}
		//mmtpr.Status = status
		mmtpr.Details.Collector.On = &t
		dbErr = s.Daos.MakeTradeLicensePayment(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}

		payment, dbErr := s.Daos.GetSingleTradeLicensePayment(ctx, mmtpr.TnxID)
		if dbErr != nil {
			return dbErr
		}
		if status == "Completed" {
			err := s.UpdateLicenseExpiry(ctx, mmtpr.TnxID)
			if err != nil {
				return err
			}
			fmt.Println("Updatetradelicenseexpiry")

		}
		tradeLicenseID = payment.TradeLicenseID
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return "", errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return "", err
	}
	return tradeLicenseID, nil

}

func (s *Service) TradeLicensePaymentStatusSelector(ctx *models.Context, mmtpr *models.MakeTradeLicensePaymentReq) (string, error) {
	if mmtpr == nil {
		return "", errors.New("Nil Payment while selecting status")
	}
	switch mmtpr.Details.MOP.Mode {
	case "Cash":
		return constants.TRADELICENSEPAYMENRSTATUSCOMPLETED, nil
	default:
		return constants.TRADELICENSEPAYMENRSTATUSPENDING, nil
	}

}

// FilterTradeLicensePayment : ""
func (s *Service) FilterTradeLicensePayment(ctx *models.Context, filter *models.TradeLicensePaymentsFilter, pagination *models.Pagination) ([]models.RefTradeLicensePayments, error) {
	return s.Daos.FilterTradeLicensePayment(ctx, filter, pagination)
}

func (s *Service) GetTradeLicensePaymentReceiptsPDF(ctx *models.Context, ID string) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.GetSingleTradeLicensePayment(ctx, ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(data.ReciptNo)
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["Payment"] = data
	m2["currentDate"] = time.Now()
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "tradelicense_receipt.html"
	err = r.ParseTemplate(templatePath, pdfdata)
	if err != nil {
		return nil, err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	fmt.Println(ok, "pdf generated successfully")

	return file, nil
}

func (s *Service) TradeLicenseBouncePayment(ctx *models.Context, bp *models.BouncePayment) (string, error) {
	t := time.Now()
	bp.ActionDate = &t
	if bp.Date == nil {
		bp.Date = &t
	}
	err := s.Daos.TradeLicenseBouncePayment(ctx, bp)
	if err != nil {
		return "", err
	}
	propertypayment, err := s.Daos.GetSingleTradeLicensePaymentWithTxtID(ctx, bp.TnxID)
	if err != nil {
		return "", err
	}
	return propertypayment.PropertyID, err
}

func (s *Service) DateRangeWiseTradeLisencePaymentReport(ctx *models.Context, filter *models.DateWiseTradeLicenseReportFilter) (*models.RefDateWiseTradeLicensePaymentReport, error) {
	return s.Daos.DateRangeWiseTradeLisencePaymentReport(ctx, filter)
}

// GetSinglePropertyPaymentTxtID : ""
func (s *Service) GetSingleTradelicensePaymentTxtID(ctx *models.Context, id string) (*models.RefTradeLicensePayments, error) {
	refPropertyPayment := new(models.RefTradeLicensePayments)
	payment, err := s.Daos.GetSingleTradeLicensePaymentWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment - " + err.Error())
	}
	propertyDemandBasic, err := s.Daos.GetSingleTradelicencePaymentDemandBasicWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment damand basic")
	}
	propertyDemandFys, err := s.Daos.GetTradelicencePaymentDemandFycWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment damand fys")
	}
	// ppFilter := new(models.PropertyPartPaymentFilter)
	// ppFilter.TnxID = []string{id}
	// ppFilter.Status = []string{constants.PROPERTYPAYMENTCOMPLETED}
	//FilterPropertyPartPayment(ctx *models.Context, filter *models.PropertyPartPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyPartPayment, error) {
	// refPartPayments, _ := s.Daos.FilterPropertyPartPayment(ctx, ppFilter, nil)
	// if refPartPayments != nil {
	// 	refPropertyPayment.Ref.PartPayments = refPartPayments
	// 	for _, v := range refPartPayments {
	// 		refPropertyPayment.Ref.PartAmountCollected = refPropertyPayment.Ref.PartAmountCollected + v.Details.Amount
	// 	}
	// }
	refPropertyPayment.TradeLicensePayments = *payment
	refPropertyPayment.Basic = propertyDemandBasic
	refPropertyPayment.FYs = propertyDemandFys
	state, err := s.Daos.GetSingleState(ctx, payment.Address.StateCode)
	if state != nil {
		refPropertyPayment.Ref.Address.State = &state.State
	}
	fmt.Println(err)
	district, err := s.Daos.GetSingleDistrict(ctx, payment.Address.DistrictCode)
	if district != nil {
		refPropertyPayment.Ref.Address.District = &district.District
	}
	fmt.Println(err)
	village, err := s.Daos.GetSingleVillage(ctx, payment.Address.VillageCode)
	if village != nil {
		refPropertyPayment.Ref.Address.Village = &village.Village
	}
	fmt.Println(err)
	zone, err := s.Daos.GetSingleZone(ctx, payment.Address.ZoneCode)
	if zone != nil {
		refPropertyPayment.Ref.Address.Zone = &zone.Zone
	}
	fmt.Println(err)
	ward, err := s.Daos.GetSingleWard(ctx, payment.Address.WardCode)
	if ward != nil {
		refPropertyPayment.Ref.Address.Ward = &ward.Ward
	}
	fmt.Println(err)

	collector, err := s.Daos.GetSingleUser(ctx, payment.Details.Collector.By)
	if collector != nil {
		refPropertyPayment.Ref.Collector = collector.User
	}
	fmt.Println(err)

	return refPropertyPayment, nil
}

// FilterPropertyPaymentExcel : ""
func (s *Service) FilterTradeLicenseMonthlyPaymentExcel(ctx *models.Context, filter *models.TradeLicensePaymentsFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterTradeLicensePayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Payments"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "K3")
	excel.MergeCell(sheet1, "C4", "K5")
	excel.MergeCell(sheet1, "A6", "K6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Usercharge Payment List")
	rowNo++
	rowNo++

	//
	reportFromMsg := "Report"
	t := time.Now()
	toDate := t.Format("02-January-2006")
	if filter != nil {
		if filter.CompletionDate.From != nil {
			fmt.Println(filter.CompletionDate.From, filter.CompletionDate.To)
			if filter.CompletionDate.From != nil && filter.CompletionDate.To == nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.CompletionDate.From.Day(), filter.CompletionDate.From.Month(), filter.CompletionDate.From.Year()) + " To " + toDate
			}
			if filter.CompletionDate.From != nil && filter.CompletionDate.To != nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.CompletionDate.From.Day(), filter.CompletionDate.From.Month(), filter.CompletionDate.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.CompletionDate.To.Day(), filter.CompletionDate.To.Month(), filter.CompletionDate.To.Year())
			}
			if filter.CompletionDate.From == nil && filter.CompletionDate.To == nil {
				fmt.Println("everything is nil")
			}

		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++
	//
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "ReceiptNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Payee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Mode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Payment Made At")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Collected By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Amount")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Rejected By")

	fmt.Println("'res length==>'", len(res))
	var totalAmount float64
	for i, v := range res {
		totalAmount = totalAmount + func() float64 {
			if v.Details.Amount != 0 {
				return v.Details.Amount
			}
			return 0
		}()

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		if v.CompletionDate != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.CompletionDate.Format("2006-01-02"))
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Basic.TradeLicenseID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.ReciptNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Details.PayeeName != "" {
				return v.Details.PayeeName
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Details.MOP.Mode != "" {
				return v.Details.MOP.Mode
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Details.MadeAt.At != "" {
				if v.Details.MadeAt != nil {
					return v.Details.MadeAt.At
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Ref.Collector.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() interface{} {
			if v.Details.Amount != 0 {
				return v.Details.Amount
			}
			return "NA"
		}())
	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%.0f", totalAmount))

	return excel, nil
}
