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

// InitiateMobileTowerPayment : ""
func (s *Service) InitiateMobileTowerPayment(ctx *models.Context, imtpr *models.InitiateMobileTowerPaymentReq) (string, error) {
	// Start Transaction
	tnxId := ""
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		filter := new(models.MobileTowerCalcQueryFilter)
		filter.AddFy = imtpr.FYs
		filter.OmitPayedYears = true
		mtd, err := s.CalcMobileTowerDemand(ctx, imtpr.MobileTowerID, filter)
		if err != nil {
			return errors.New("Error in calculating demand - " + err.Error())
		}
		if mtd == nil {
			return errors.New("Demand is nil ")
		}

		mtp := new(models.MobileTowerPayments)
		mtp.TnxID = s.Shared.GetTransactionID(mtd.PropertyMobileTower.PropertyID, 32)
		tnxId = mtp.TnxID
		mtp.PropertyID = mtd.PropertyID
		mtp.MobileTowerID = mtd.PropertyMobileTower.UniqueID
		mtp.Scenario = constants.MOBILETOWERPAYMENTYEARLYPAYMENT

		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		mtp.FinancialYear = fy.FinancialYear

		mtp.Demand = models.MobileTowerPaymentDemand{
			Current: models.MobileTowerPaymentDemandSplitage{
				Tax:     mtd.Demand.Current.Tax,
				Penalty: mtd.Demand.Current.Penalty,
				Other:   mtd.Demand.Current.Other,
				Total:   mtd.Demand.Current.Total,
			},
			Arrear: models.MobileTowerPaymentDemandSplitage{
				Tax:     mtd.Demand.Arrear.Tax,
				Penalty: mtd.Demand.Arrear.Penalty,
				Other:   mtd.Demand.Arrear.Other,
				Total:   mtd.Demand.Arrear.Total,
			},
			Total: models.MobileTowerPaymentDemandSplitage{
				Tax:     mtd.Demand.Total.Tax,
				Penalty: mtd.Demand.Total.Penalty,
				Other:   mtd.Demand.Total.Other,
				Total:   mtd.Demand.Total.Total,
			},
		}
		mtp.Status = constants.MOBILETOWERPAYMENRSTATUSINIT
		mtp.Address = mtd.Address
		//	mtp.Address = mtd.Property.Address
		mtp.Created = models.CreatedV2{
			By:     imtpr.By,
			ByType: imtpr.ByType,
			On:     &t,
		}
		var mtpFys []models.MobileTowerPaymentsfY
		for _, v := range mtd.FY {
			var mtpFy models.MobileTowerPaymentsfY
			mtpFy.TnxID = mtp.TnxID
			mtpFy.PropertyID = mtd.PropertyID
			mtpFy.MobileTowerID = mtd.PropertyMobileTower.UniqueID
			mtpFy.FY = v
			mtpFy.Status = constants.MOBILETOWERPAYMENRSTATUSINIT
			mtpFys = append(mtpFys, mtpFy)
		}

		var mtpBasic = new(models.MobileTowerPaymentsBasics)
		mtpBasic.TnxID = mtp.TnxID
		mtpBasic.PropertyID = mtd.PropertyID
		mtpBasic.MobileTowerID = mtd.PropertyMobileTower.UniqueID
		mobiletower, err := s.Daos.GetSinglePropertyMobileTower(ctx, mtd.PropertyMobileTower.UniqueID)
		if err != nil {
			return errors.New("Error in getting property of mobile tower  - " + err.Error())
		}
		if mobiletower != nil {
			mtpBasic.MobileTower = *mobiletower
		}
		property, err := s.Daos.GetSingleProperty(ctx, mtd.PropertyID)
		if err != nil {
			return errors.New("Error in getting property of mobile tower property - " + err.Error())
		}
		mtpBasic.Property = property
		owners, err := s.Daos.GetOwnersOfProperty(ctx, mtd.PropertyID)
		if err != nil {
			return errors.New("Error in getting owners of mobile tower - " + err.Error())
		}
		mtpBasic.Owner = owners
		mtpBasic.Status = constants.MOBILETOWERPAYMENRSTATUSINIT

		err = s.Daos.SaveMobileTowerPayment(ctx, mtp)
		if err != nil {
			return errors.New("Error in saving mobile tower payment - " + err.Error())
		}
		err = s.Daos.SaveMobileTowerPaymentFYs(ctx, mtpFys)
		if err != nil {
			return errors.New("Error in saving mobile tower payment fys- " + err.Error())
		}
		err = s.Daos.SaveMobileTowerPaymentBasic(ctx, mtpBasic)
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

func (s *Service) GetSingleMobileTowerPayment(ctx *models.Context, tnxID string) (*models.RefMobileTowerPayments, error) {
	return s.Daos.GetSingleMobileTowerPayment(ctx, tnxID)
}

func (s *Service) MakeMobileTowerPayment(ctx *models.Context, mmtpr *models.MakeMobileTowerPaymentReq) error {
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		mmtpr.CompletionDate = &t
		status, dbErr := s.MobileTowerPaymentStatusSelector(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}
		mmtpr.Status = status
		mmtpr.Details.Collector.On = &t

		dbErr = s.Daos.MakeMobileTowerPayment(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
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

func (s *Service) MobileTowerPaymentStatusSelector(ctx *models.Context, mmtpr *models.MakeMobileTowerPaymentReq) (string, error) {
	if mmtpr == nil {
		return "", errors.New("Nil Payment while selecting status")
	}
	switch mmtpr.Details.MOP.Mode {
	case "Cash":
		return constants.MOBILETOWERPAYMENRSTATUSCOMPLETED, nil
	default:
		return constants.MOBILETOWERPAYMENRSTATUSPENDING, nil
	}

}

func (s *Service) FilterMobileTowerPayment(ctx *models.Context, filter *models.MobileTowerPaymentsFilter, pagination *models.Pagination) ([]models.RefMobileTowerPayments, error) {
	return s.Daos.FilterMobileTowerPayment(ctx, filter, pagination)
}

// VerifyMobileTowerPayment : ""
func (s *Service) VerifyMobileTowerPayment(ctx *models.Context, makeAction *models.MakeMobileTowerPaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.VerifyMobileTowerPayment(ctx, makeAction)
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

// NotVerifyMobileTowerPayment : ""
func (s *Service) NotVerifyMobileTowerPayment(ctx *models.Context, makeAction *models.MakeMobileTowerPaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.NotVerifyMobileTowerPayment(ctx, makeAction)
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

// RejectMobileTowerPayment : ""
func (s *Service) RejectMobileTowerPayment(ctx *models.Context, makeAction *models.MakeMobileTowerPaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.RejectMobileTowerPayment(ctx, makeAction)
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

func (s *Service) GetMobileTowerPaymentReceiptsPDF(ctx *models.Context, ID string) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.GetSingleMobileTowerPayment(ctx, ID)
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
		return nil, errors.New("Error in geting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "mobiletower_receipt.html"
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
