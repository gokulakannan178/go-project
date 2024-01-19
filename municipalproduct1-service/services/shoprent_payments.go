package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// InitiateShopRentPayment : ""
func (s *Service) InitiateShopRentPayment(ctx *models.Context, ipmtr *models.InitiateShopRentPaymentReq) (string, error) {
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
		filter := new(models.ShopRentCalcQueryFilter)
		filter.ShopRentID = ipmtr.ShopRentID
		filter.AddFy = ipmtr.FYs
		demand, err := s.CalcShopRentDemandForParticulars(ctx, filter)
		if err != nil {
			return errors.New("Error in calculating demand - " + err.Error())
		}
		if demand == nil {
			return errors.New("Demand is nil ")
		}

		pmt := new(models.ShopRentPayments)
		pmt.TnxID = s.Shared.GetTransactionID(demand.ShopRent.UniqueID, 32)
		tnxId = pmt.TnxID
		pmt.ShopRentID = demand.ShopRent.UniqueID

		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		pmt.FinancialYear = fy.FinancialYear

		pmt.Demand = models.ShopRentPaymentDemand{
			Current: models.ShopRentPaymentDemandSplitage{
				Tax:     demand.Demand.Current.Tax,
				Penalty: demand.Demand.Current.Penalty,
				Other:   demand.Demand.Current.Other,
				Total:   demand.Demand.Current.Total,
			},
			Arrear: models.ShopRentPaymentDemandSplitage{
				Tax:     demand.Demand.Arrear.Tax,
				Penalty: demand.Demand.Arrear.Penalty,
				Other:   demand.Demand.Arrear.Other,
				Total:   demand.Demand.Arrear.Total,
			},
			Total: models.ShopRentPaymentDemandSplitage{
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
		var pmtFys []models.ShopRentPaymentsfY
		for _, v := range demand.FY {
			var pmtFy models.ShopRentPaymentsfY
			pmtFy.TnxID = pmt.TnxID
			pmtFy.ShopRentID = demand.ShopRent.UniqueID
			pmtFy.FY = v
			pmtFy.Status = constants.SHOPRENTPAYMENTSTATUSINIT
			pmtFys = append(pmtFys, pmtFy)
		}

		var pmtBasic = new(models.ShopRentPaymentsBasics)
		pmtBasic.TnxID = pmt.TnxID
		pmtBasic.ShopRentID = demand.ShopRent.UniqueID
		sr, err := s.Daos.GetSingleShopRent(ctx, ipmtr.ShopRentID)
		if err != nil {
			return err
		}
		pmt.Address = demand.ShopRent.Address
		if sr != nil {
			pmtBasic.ShopRent = *sr
		}
		err = s.Daos.SaveShopRentPayment(ctx, pmt)
		if err != nil {
			return errors.New("Error in saving mobile tower payment - " + err.Error())
		}
		err = s.Daos.SaveShopRentPaymentFYs(ctx, pmtFys)
		if err != nil {
			return errors.New("Error in saving mobile tower payment fys- " + err.Error())
		}
		err = s.Daos.SaveShopRentPaymentBasic(ctx, pmtBasic)
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

func (s *Service) GetSingleShopRentPayment(ctx *models.Context, tnxID string) (*models.RefShopRentPayments, error) {
	return s.Daos.GetSingleShopRentPayment(ctx, tnxID)
}

func (s *Service) MakeShopRentPayment(ctx *models.Context, mmtpr *models.MakeShopRentPaymentReq) (string, error) {
	shopRentID := ""
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		mmtpr.CompletionDate = &t
		status, dbErr := s.ShopRentPaymentStatusSelector(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}
		mmtpr.Status = status
		mmtpr.Details.Collector.On = &t
		dbErr = s.Daos.MakeShopRentPayment(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}

		payment, dbErr := s.Daos.GetSingleShopRentPayment(ctx, mmtpr.TnxID)
		if dbErr != nil {
			return dbErr
		}
		shopRentID = payment.ShopRentID
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
	return shopRentID, nil

}

func (s *Service) ShopRentPaymentStatusSelector(ctx *models.Context, mmtpr *models.MakeShopRentPaymentReq) (string, error) {
	if mmtpr == nil {
		return "", errors.New("Nil Payment while selecting status")
	}
	switch mmtpr.Details.MOP.Mode {
	case "Cash":
		return constants.SHOPRENTPAYMENTSTATUSCOMPLETED, nil
	default:
		return constants.SHOPRENTPAYMENTSTATUSPENDING, nil
	}

}

// FilterShopRentPayment : ""
func (s *Service) FilterShopRentPayment(ctx *models.Context, filter *models.ShopRentPaymentsFilter, pagination *models.Pagination) ([]models.RefShopRentPayments, error) {
	return s.Daos.FilterShopRentPayment(ctx, filter, pagination)
}

// GetShopRentPaymentReceiptsPDF : ""
func (s *Service) GetShopRentPaymentReceiptsPDF(ctx *models.Context, ID string) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.GetSingleShopRentPayment(ctx, ID)
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
	templatePath := templatePathStart + "shoprent_receipt.html"
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

func (s *Service) ShoprentBouncePayment(ctx *models.Context, bp *models.BouncePayment) (string, error) {
	t := time.Now()
	bp.ActionDate = &t
	if bp.Date == nil {
		bp.Date = &t
	}
	err := s.Daos.ShoprentBouncePayment(ctx, bp)
	if err != nil {
		return "", err
	}
	propertypayment, err := s.Daos.GetSingleShoprentPaymentWithTxtID(ctx, bp.TnxID)
	if err != nil {
		return "", err
	}
	return propertypayment.PropertyID, err
}
