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

// InitiateTradeLicensePayment : ""
func (s *Service) InitiateTradeLicensePaymentV2(ctx *models.Context, ipmtr *models.InitiateTradeLicensePaymentReq) (string, error) {
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
		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		filter := new(models.TradeLicenseCalcQueryFilter)
		filter.TradeLicenseID = ipmtr.TradeLicenseID
		filter.AddFy = []string{fy.UniqueID}

		pmt := new(models.TradeLicensePayments)

		demand, err := s.CalcTradeLicenseDemandForParticulars(ctx, filter)
		if err != nil {
			return errors.New("Error in calculating demand - " + err.Error())
		}
		if demand == nil {
			return errors.New("Demand is nil ")
		}

		pmt.TnxID = s.Shared.GetTransactionID(demand.TradeLicense.UniqueID, 32)
		tnxId = pmt.TnxID
		pmt.TradeLicenseID = demand.TradeLicense.UniqueID
		pmt.ReciptNo = fmt.Sprintf("%v%v%v_%v_%v", "TL", t.Day(), int(t.Month()), t.Year(), s.Daos.GetUniqueID(ctx, "tl_recipt"))

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

func (s *Service) GetTradeLicensePaymentReceiptsPDFV2(ctx *models.Context, ID string) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.GetSingleTradeLicenseV2(ctx, ID)
	if err != nil {
		return nil, err
	}
	if data != nil {
		strs := strings.Split(data.Ref.TradeLicenseCategoryType.Name, ",")

		data.Ref.CategoryNames = strs
	}
	// fmt.Println(data.UniqueID)
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["tradeLicense"] = data
	m2["currentDate"] = time.Now()
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in geting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "tradelicense_certificate.html"
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
func (s *Service) UpdateLicenseExpiry(ctx *models.Context, tnxID string) error {
	res, err := s.Daos.GetSingleTradeLicensePayment(ctx, tnxID)
	if err != nil {
		return err
	}

	res1, err := s.Daos.GetSingleTradeLicenseV2(ctx, res.TradeLicenseID)
	if err != nil {
		return err
	}
	t := time.Now()
	if res1 != nil {
		if len(res1.Ref.Payments) > 0 {
			payment := res1.Ref.Payments[0]
			if len(payment.FYs) > 0 {
				fy := payment.FYs[0]
				fmt.Println("fy value", fy.FY.FinancialYear.To)
				status := constants.TRADELICENSESTATUSACTIVE
				if ctx.ProductConfig.LocationID == "Bhagalpur" {
					status = constants.TRADELICENSESTATUSPENDING
				}
				err := s.Daos.UpdateExpiryDate(ctx, res.TradeLicenseID, fy.FY.FinancialYear.To, &t, status)
				if err != nil {
					return err
				}

			}

		}

	}

	return nil
}

// UpdateTradeLicenseExpiryOnReject : ""
func (s *Service) UpdateTradeLicenseExpiryOnReject(ctx *models.Context, tnxID string) error {
	res, err := s.Daos.GetSingleTradeLicensePayment(ctx, tnxID)
	if err != nil {
		return err
	}

	err = s.Daos.UpdateExpiryDate(ctx, res.TradeLicenseID, nil, nil, constants.TRADELICENSESTATUSNOTVERIFIED)
	if err != nil {
		return err
	}
	return nil
}

// // GetTradeLicenseCurrentFinancialYearMarketYear : ""
// func (s *Service) GetTradeLicenseCurrentFinancialYearMarketYear(ctx *models.Context) error {
// 	fy, err := s.Daos.GetCurrentFinancialYear(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("fy value", fy)
// 	err = s.Daos.GetTradeLicenseCurrentFinancialYearMarketYear(ctx, fy.From, fy.To)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
