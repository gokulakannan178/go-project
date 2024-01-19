package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// InitiateTradeLicensePaymentPart2 : ""
func (s *Service) InitiateTradeLicensePaymentPart2(ctx *models.Context, filter *models.TradeLicenseDemandPart2Filter) (string, error) {
	var transactionID string
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		refTlDemand, err := s.TradeLicenseDemandCalcPart2(ctx, filter)
		if err != nil {
			return err
		}
		if refTlDemand == nil {
			return errors.New("nil demand for this trade license")
		}
		tlpp2 := new(models.TradeLicensePaymentsPart2)
		tlpp2.TnxID = s.Shared.GetTransactionID(refTlDemand.TradeLicense.UniqueID, 32)
		transactionID = tlpp2.TnxID
		tlpp2.TradeLicenseID = refTlDemand.TradeLicense.UniqueID
		refFy, err := s.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("error in getting financial year")
		}
		tlpp2.FinancialYear = refFy.FinancialYear
		tlpp2.Status = constants.TRADELICENSEPAYMENRSTATUSINIT
		t := time.Now()

		tlpp2.Created = models.CreatedV2{
			By:     filter.By,
			ByType: filter.ByType,
			On:     &t,
		}
		tlpp2.Address = refTlDemand.Address
		tlpp2.Demand = models.TradeLicensePaymentDemand{
			Current: models.TradeLicensePaymentDemandSplitage{
				Tax:     refTlDemand.Demand.Current.Tax,
				Penalty: refTlDemand.Demand.Current.Penalty,
				Rebate:  refTlDemand.Demand.Current.Rebate,
				Other:   refTlDemand.Demand.Current.Other,
				Total:   refTlDemand.Demand.Current.Total,
			},
			Arrear: models.TradeLicensePaymentDemandSplitage{
				Tax:     refTlDemand.Demand.Arrear.Tax,
				Penalty: refTlDemand.Demand.Arrear.Penalty,
				Rebate:  refTlDemand.Demand.Arrear.Rebate,
				Other:   refTlDemand.Demand.Arrear.Other,
				Total:   refTlDemand.Demand.Arrear.Total,
			},
			Total: models.TradeLicensePaymentDemandSplitage{
				Tax:     refTlDemand.Demand.Total.Tax,
				Penalty: refTlDemand.Demand.Total.Penalty,
				Rebate:  refTlDemand.Demand.Total.Rebate,
				Other:   refTlDemand.Demand.Total.Other,
				Total:   refTlDemand.Demand.Total.Total,
			},
		}
		refTl, err := s.Daos.GetSingleTradeLicense(ctx, filter.TradeLicenseID)
		if err != nil {
			return errors.New("error in getting tradelicense data - " + err.Error())
		}
		var tlpfys []models.TradeLicensePaymentsfY
		for _, v := range refTlDemand.FYs {
			var tlpfy models.TradeLicensePaymentsfY
			tlpfy.TradeLicenseID = refTlDemand.TradeLicense.UniqueID
			tlpfy.TnxID = transactionID
			tlpfy.FY.FinancialYear = v.FinancialYear
			tlpfy.FY.TradeLicenseID = refTlDemand.TradeLicense.UniqueID
			tlpfy.FY.Status = constants.TRADELICENSEPAYMENRSTATUSINIT
			tlpfy.FY.Details.Penalty = v.FYTLPenalty
			tlpfy.FY.Details.PenaltyValue = v.PenaltyValue
			tlpfy.FY.Details.RebateValue = v.RebateValue
			tlpfy.FY.Details.Rebate = v.FYRebate
			tlpfy.FY.Details.AfterRebate = v.FYAfterRebate
			tlpfy.FY.Details.Tax = v.FYTax
			tlpfy.FY.Details.TotalTaxAmount = v.FYTotal
			tlpfys = append(tlpfys, tlpfy)
		}

		var tlpb = new(models.TradeLicensePaymentsBasicsPart2)
		tlpb.TnxID = transactionID
		tlpb.TradeLicenseID = refTlDemand.TradeLicense.UniqueID
		tlpb.TradeLicense = *refTl
		err = s.Daos.SaveTradeLicensePaymentsPart2(ctx, tlpp2)
		if err != nil {
			return errors.New("error in saving tradelicensepaymentspart2 data - " + err.Error())
		}

		err = s.Daos.SaveTradeLicensePaymentsFYsPart2(ctx, tlpfys)
		if err != nil {
			return errors.New("error in saving tradelicensepaymentsfyspart2 data - " + err.Error())
		}
		err = s.Daos.SaveTradeLicensePaymentsBasicsPart2(ctx, tlpb)
		if err != nil {
			return errors.New("error in saving tradelicensepaymentsbasicspart2 data - " + err.Error())
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
	return transactionID, nil

}

// GetSingleTradeLicensePaymentPart2 : ""
func (s *Service) GetSingleTradeLicensePaymentPart2(ctx *models.Context, tnxID string) (*models.RefTradeLicensePayments, error) {
	return s.Daos.GetSingleTradeLicensePaymentPart2(ctx, tnxID)
}

// MakeTradeLicensePaymentPart2 : ""
func (s *Service) MakeTradeLicensePaymentPart2(ctx *models.Context, mtlprp2 *models.MakeTradeLicensePaymentReqPart2) (string, error) {
	tradeLicenseID := ""
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		mtlprp2.CompletionDate = &t
		fmt.Println("value of completion date ======> ", mtlprp2.CompletionDate)
		status, dbErr := s.TradeLicensePaymentStatusSelector(ctx, &mtlprp2.MakeTradeLicensePaymentReq)
		if dbErr != nil {
			return dbErr
		}
		mtlprp2.Status = status
		mtlprp2.Details.Collector.On = &t
		dbErr = s.Daos.MakeTradeLicensePaymentPart2(ctx, mtlprp2)
		if dbErr != nil {
			return dbErr
		}

		payment, dbErr := s.Daos.GetSingleTradeLicensePaymentPart2(ctx, mtlprp2.TnxID)
		if dbErr != nil {
			return dbErr
		}
		if status == "Completed" {
			err := s.UpdateLicenseExpiryPart2(ctx, mtlprp2.TnxID)
			if err != nil {
				return err
			}
			fmt.Println("Updatetradelicenseexpirypart2")

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

// FilterTradeLicensePaymentPart2 : ""
func (s *Service) FilterTradeLicensePaymentPart2(ctx *models.Context, filter *models.TradeLicensePaymentsFilterPart2, pagination *models.Pagination) ([]models.RefTradeLicensePaymentsPart2, error) {
	return s.Daos.FilterTradeLicensePaymentPart2(ctx, filter, pagination)
}

// GetTradeLicensePaymentReceiptsPDFPart2 : ""
func (s *Service) GetTradeLicensePaymentReceiptsPDFPart2(ctx *models.Context, ID string) ([]byte, error) {
	r := NewRequestPdf("")
	data, err := s.GetSingleTradeLicensePaymentPart2(ctx, ID)
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
	templatePath := templatePathStart + "tradelicense_receipt_part2.html"
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

//GetSingleTradeLicenseV2 :""
func (s *Service) GetSingleTradeLicenseV2Part2(ctx *models.Context, UniqueID string) (*models.RefTradeLicense, error) {
	tower, err := s.Daos.GetSingleTradeLicenseV2Part2(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// GetTradeLicensePaymentReceiptsPDFV2Part2 : ""
func (s *Service) GetTradeLicensePaymentReceiptsPDFV2Part2(ctx *models.Context, ID string) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.GetSingleTradeLicenseV2Part2(ctx, ID)
	if err != nil {
		return nil, err
	}
	var varFy models.TradeLicensePaymentsfY
	if data != nil {
		strs := strings.Split(data.Ref.TradeLicenseCategoryType.Name, ",")

		data.Ref.CategoryNames = strs
		if len(data.Ref.Payments) > 0 {
			if len(data.Ref.Payments[0].FYs) > 0 {
				varFy = data.Ref.Payments[0].FYs[0]
			}
		}

		for _, v := range data.Ref.Payments {
			for _, v2 := range v.FYs {
				fmt.Println("v2.FY.From", v2.FY.From)
				fmt.Println("varFy.FY.From", varFy.FY.From)
				fmt.Println("**********", *varFy.FY.From)
				if v2.FY.From.After(*varFy.FY.From) {
					varFy = v2
				}
				fmt.Println("**********", *varFy.FY.From)
				fmt.Println("=========================================")
			}
		}
	}

	// fmt.Println(data.UniqueID)
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["tradeLicense"] = data
	fmt.Println("varFy =======>", varFy)
	m["recentPayment"] = varFy
	m2["currentDate"] = time.Now()
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in geting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	// pdfdata.Data = m3
	pdfdata.Config = productConfig.ProductConfiguration

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "tradelicense_certificate_part2.html"
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

// UpdateLicenseExpiry : ""
func (s *Service) UpdateLicenseExpiryPart2(ctx *models.Context, tnxID string) error {
	res, err := s.Daos.GetSingleTradeLicensePaymentPart2(ctx, tnxID)
	if err != nil {
		return err
	}

	res1, err := s.Daos.GetSingleTradeLicenseV2Part2(ctx, res.TradeLicenseID)
	if err != nil {
		return err
	}
	var varFy models.TradeLicensePaymentsfY

	if res1 != nil {

		if len(res1.Ref.Payments) > 0 {
			if len(res1.Ref.Payments[0].FYs) > 0 {
				varFy = res1.Ref.Payments[0].FYs[0]
			}
		}

		for _, v := range res1.Ref.Payments {
			for _, v2 := range v.FYs {
				fmt.Println("v2.FY.From", v2.FY.From)
				fmt.Println("varFy.FY.From", varFy.FY.From)
				fmt.Println("**********", *varFy.FY.From)
				if v2.FY.From.After(*varFy.FY.From) {
					varFy = v2
				}
				fmt.Println("**********", *varFy.FY.From)
				fmt.Println("=========================================")
				fmt.Println("**********", varFy.FY.From)
				fmt.Println("**********", varFy.FY.To)

				err := s.Daos.UpdateExpiryDatePart2(ctx, res.TradeLicenseID, varFy.FY.To, varFy.FY.From, constants.TRADELICENSESTATUSACTIVE)
				if err != nil {
					return err
				}
			}
		}
	}

	// if res1 != nil {
	// 	if len(res1.Ref.Payments) > 0 {
	// 		payment := res1.Ref.Payments[0]
	// 		if len(payment.FYs) > 0 {
	// 			fy := payment.FYs[0]
	// 			fmt.Println("fy value", fy.FY.FinancialYear.To)
	// 			err := s.Daos.UpdateExpiryDatePart2(ctx, res.TradeLicenseID, fy.FY.FinancialYear.To, fy.FY.FinancialYear.From, constants.TRADELICENSESTATUSACTIVE)
	// 			if err != nil {
	// 				return err
	// 			}
	// 		}

	// 	}

	// }

	return nil
}

// UpdateTradeLicenseExpiryOnRejectPart2 : ""
func (s *Service) UpdateTradeLicenseExpiryOnRejectPart2(ctx *models.Context, tnxID string) error {
	res, err := s.Daos.GetSingleTradeLicensePaymentPart2(ctx, tnxID)
	if err != nil {
		return err
	}

	err = s.Daos.UpdateExpiryDatePart2(ctx, res.TradeLicenseID, nil, nil, constants.TRADELICENSESTATUSNOTVERIFIED)
	if err != nil {
		return err
	}
	return nil
}

// VerifyTradeLicensePaymentPart2 : ""
func (s *Service) VerifyTradeLicensePaymentPart2(ctx *models.Context, makeAction *models.MakeTradeLicensePaymentsActionPart2) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.VerifyTradeLicensePaymentPart2(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		dberr1 := s.UpdateLicenseExpiryPart2(ctx, makeAction.TnxID)
		if dberr1 != nil {
			return dberr1
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

// NotVerifyTradeLicensePaymentPart2 : ""
func (s *Service) NotVerifyTradeLicensePaymentPart2(ctx *models.Context, makeAction *models.MakeTradeLicensePaymentsActionPart2) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.NotVerifyTradeLicensePaymentPart2(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		dberr1 := s.UpdateTradeLicenseExpiryOnRejectPart2(ctx, makeAction.TnxID)
		if dberr1 != nil {
			return dberr1
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

// RejectTradeLicensePaymentPart2 : ""
func (s *Service) RejectTradeLicensePaymentPart2(ctx *models.Context, makeAction *models.MakeTradeLicensePaymentsActionPart2) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.RejectTradeLicensePaymentPart2(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		dberr1 := s.UpdateTradeLicenseExpiryOnRejectPart2(ctx, makeAction.TnxID)
		if dberr1 != nil {
			return dberr1
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

// BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2 : ""
func (s *Service) BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2(ctx *models.Context, rbtlulp2 *models.RefBasicTradeLicenseUpdateLogV2Part2) ([]models.RefTradeLicensePaymentsPart2, error) {
	return s.Daos.BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2(ctx, rbtlulp2)
}
