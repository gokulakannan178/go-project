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

func (s *Service) InitiateSolidWasteChargeMonthlyPaymentV2(ctx *models.Context, ipmtr *models.InitiateSolidWasteChargeMonthlyPaymentReq) (string, error) {
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
		filter := new(models.SolidWasteUserChargeDemandCalcQueryFilter)
		filter.SolidWasteChargeID = ipmtr.SolidWasteChargeID
		filter.AddFy = ipmtr.Months

		demand, err := s.SolidWasteUserChargeDemandForParticulars(ctx, filter)
		if err != nil {
			return errors.New("Error in calculating demand 1- " + err.Error())
		}
		if demand == nil {
			return errors.New("Demand is nil ")
		}
		pmt := new(models.SolidWasteUserChargeMonthlyPayments)
		pmt.TnxID = s.Shared.GetTransactionID(demand.SolidWasteUserCharge.UniqueID, 32)
		tnxId = pmt.TnxID
		pmt.SolidWasteUserChargeID = demand.SolidWasteUserCharge.UniqueID

		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		pmt.FinancialYear = fy.FinancialYear
		pmt.Demand = demand.Demand
		pmt.Status = constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSINIT
		pmt.Scenario = constants.SOLIDWASTEUSERCHARGEPAYMENTSCENARIOMONTHLY
		pmt.Created = models.CreatedV2{
			By:     ipmtr.By,
			ByType: ipmtr.ByType,
			On:     &t,
		}

		var pmtFys []models.SolidWasteChargeMonthlyPaymentsfY
		for _, v := range demand.FY {
			var pmtFy models.SolidWasteChargeMonthlyPaymentsfY
			pmtFy.TnxID = pmt.TnxID
			pmtFy.SolidWasteUserChargeID = demand.SolidWasteUserCharge.UniqueID
			pmtFy.FY = v
			pmtFy.Status = constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSINIT
			pmtFys = append(pmtFys, pmtFy)
		}

		var pmtBasic = new(models.SolidWasteChargeMonthlyPaymentsBasics)
		pmtBasic.TnxID = pmt.TnxID
		pmtBasic.SolidWasteUserChargeID = demand.SolidWasteUserCharge.UniqueID
		sr, err := s.Daos.GetSingleSolidWasteUserCharge(ctx, ipmtr.SolidWasteChargeID)
		if err != nil {
			return err
		}
		pmt.Address = demand.SolidWasteUserCharge.Address
		if sr != nil {
			pmtBasic.SolidWasteUserCharge = *sr
		}
		return nil
	}); err != nil {

	}

	return tnxId, nil

}

// InitiateSolidWasteChargexMonthlyPayment : ""
func (s *Service) InitiateSolidWasteChargeMonthlyPayment(ctx *models.Context, ipmtr *models.InitiateSolidWasteChargeMonthlyPaymentReq) (string, error) {
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

		filter := new(models.SolidWasteUserChargeDemandCalcQueryFilter)
		filter.SolidWasteChargeID = ipmtr.SolidWasteChargeID
		fmt.Println("ipmtr.SolidWasteChargeID========>", ipmtr.SolidWasteChargeID)
		fmt.Println("filter.SolidWasteChargeID========>", filter.SolidWasteChargeID)
		filter.AddFy = ipmtr.Months
		demand, err := s.SolidWasteUserChargeDemandForParticulars(ctx, filter)
		if err != nil {
			return errors.New("Error in calculating demand 1- " + err.Error())
		}
		if demand == nil {
			return errors.New("Demand is nil ")
		}

		pmt := new(models.SolidWasteUserChargeMonthlyPayments)
		pmt.TnxID = s.Shared.GetTransactionID(demand.SolidWasteUserCharge.UniqueID, 32)
		tnxId = pmt.TnxID
		// pmt.SolidWasteUserChargeID = demand.SolidWasteUserCharge.UniqueID
		pmt.SolidWasteUserChargeID = ipmtr.SolidWasteChargeID

		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		pmt.FinancialYear = fy.FinancialYear
		pmt.Demand = demand.Demand
		pmt.Status = constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSINIT
		pmt.Scenario = constants.SOLIDWASTEUSERCHARGEPAYMENTSCENARIOMONTHLY
		pmt.Created = models.CreatedV2{
			By:     ipmtr.By,
			ByType: ipmtr.ByType,
			On:     &t,
		}

		var pmtFys []models.SolidWasteChargeMonthlyPaymentsfY
		for _, v := range demand.FY {
			var pmtFy models.SolidWasteChargeMonthlyPaymentsfY
			pmtFy.TnxID = pmt.TnxID
			// pmtFy.SolidWasteUserChargeID = demand.SolidWasteUserCharge.UniqueID
			pmtFy.SolidWasteUserChargeID = ipmtr.SolidWasteChargeID

			pmtFy.FY = v
			pmtFy.Status = constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSINIT
			pmtFys = append(pmtFys, pmtFy)
		}

		var pmtBasic = new(models.SolidWasteChargeMonthlyPaymentsBasics)
		pmtBasic.TnxID = pmt.TnxID
		// pmtBasic.SolidWasteUserChargeID = demand.SolidWasteUserCharge.UniqueID
		pmtBasic.SolidWasteUserChargeID = ipmtr.SolidWasteChargeID
		sr, err := s.Daos.GetSingleSolidWasteUserCharge(ctx, ipmtr.SolidWasteChargeID)
		if err != nil {
			return err
		}
		pmt.Address = demand.SolidWasteUserCharge.Address
		if sr != nil {
			pmtBasic.SolidWasteUserCharge = *sr
		}

		err = s.Daos.SaveSolidWasteUserChargeMonthlyPayment(ctx, pmt)
		if err != nil {
			return errors.New("Error in saving solid waste user charge payment - " + err.Error())
		}
		err = s.Daos.SaveSolidWasteUserChargeMonthlyPaymentFYs(ctx, pmtFys)
		if err != nil {
			return errors.New("Error in saving solid waste user charge payment fys- " + err.Error())
		}
		err = s.Daos.SaveSolidWasteUserChargePaymentBasic(ctx, pmtBasic)
		if err != nil {
			return errors.New("Error in saving solid waste user charge payment basics- " + err.Error())
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

// GetSingleSolidWasteUserChargePayment : ""
func (s *Service) GetSingleSolidWasteUserChargePayment(ctx *models.Context, tnxID string) (*models.RefSolidWasteChargeMonthlyPayments, error) {
	return s.Daos.GetSingleSolidWasteUserChargePayment(ctx, tnxID)
}

// SolidWasteUserChargePaymentStatusSelector : ""
func (s *Service) SolidWasteUserChargePaymentStatusSelector(ctx *models.Context, mmtpr *models.MakeSolidWasteUserChargePaymentReq) (string, error) {
	if mmtpr == nil {
		return "", errors.New("Nil Payment while selecting status")
	}
	switch mmtpr.Details.MOP.Mode {
	case "Cash":
		return constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSCOMPLETED, nil
	default:
		return constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSINIT, nil
	}

}

// MakeSolidWasteUserChargePayment : ""
func (s *Service) MakeSolidWasteUserChargePayment(ctx *models.Context, mmtpr *models.MakeSolidWasteUserChargePaymentReq) (string, error) {
	SolidWasteUserChargeID := ""
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		mmtpr.CompletionDate = &t
		status, dbErr := s.SolidWasteUserChargePaymentStatusSelector(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}
		mmtpr.Status = status
		mmtpr.Details.Collector.On = &t
		dbErr = s.Daos.MakeSolidWasteUserChargePayment(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}

		payment, dbErr := s.Daos.GetSingleSolidWasteUserChargePayment(ctx, mmtpr.TnxID)
		if dbErr != nil {
			return dbErr
		}
		SolidWasteUserChargeID = payment.SolidWasteUserChargeID
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
	return SolidWasteUserChargeID, nil

}

// FilterSolidWasteUserChargePayment : ""
func (s *Service) FilterSolidWasteUserChargePayment(ctx *models.Context, filter *models.SolidWasteUserChargePaymentsFilter, pagination *models.Pagination) ([]models.RefSolidWasteChargeMonthlyPayments, error) {
	return s.Daos.FilterSolidWasteUserChargePayment(ctx, filter, pagination)
}

// GetSolidWasteUserChargePaymentReceiptsPDF : ""
func (s *Service) GetSolidWasteUserChargePaymentReceiptsPDF(ctx *models.Context, ID string) ([]byte, error) {
	r := NewRequestPdf("")
	data, err := s.GetSingleSolidWasteUserChargePayment(ctx, ID)
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
	templatePath := templatePathStart + "solidwaste_paymentreceipt.html"
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

// VerifySolidWasteUserChargePayment : ""
func (s *Service) VerifySolidWasteUserChargePayment(ctx *models.Context, action *models.MakeSolidWasteUserChargePaymentsAction) (string, error) {
	log.Println("transaction start")
	solidWasteUserChargeID := ""
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		action.ActionDate = &t
		var dberr error

		solidWasteUserChargeID, dberr = s.Daos.VerifySolidWasteUserChargePayment(ctx, action)
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
			return "", errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return "", err
	}
	return solidWasteUserChargeID, nil
}

// NotVerifySolidWasteUserChargePayment : ""
func (s *Service) NotVerifySolidWasteUserChargePayment(ctx *models.Context, makeAction *models.MakeSolidWasteUserChargePaymentsAction) (string, error) {
	solidWasteUserChargeID := ""
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		var dberr error

		solidWasteUserChargeID, dberr = s.Daos.NotVerifySolidWasteUserChargePayment(ctx, makeAction)
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
			return "", errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return "", err
	}
	return solidWasteUserChargeID, nil
}

// RejectSolidWasteUserChargePayment : ""
func (s *Service) RejectSolidWasteUserChargePayment(ctx *models.Context, makeAction *models.MakeSolidWasteUserChargePaymentsAction) (string, error) {
	log.Println("transaction start")
	solidWasteUserChargeID := ""
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		var dberr error
		solidWasteUserChargeID, dberr = s.Daos.RejectSolidWasteUserChargePayment(ctx, makeAction)
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
			return "", errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return "", err
	}
	return solidWasteUserChargeID, nil
}
