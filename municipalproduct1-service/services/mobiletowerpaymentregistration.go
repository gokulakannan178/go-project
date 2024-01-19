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

// InitiateMobileTowerRegisterPayment : ""
func (s *Service) InitiateMobileTowerRegisterPayment(ctx *models.Context, imtpr *models.InitiateMobileTowerPaymentReq) (string, error) {
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	tnxId := ""

	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		mt, err := s.GetSingleMobileTower(ctx, imtpr.MobileTowerID)
		if err != nil {
			return errors.New("Error in calculating demand - " + err.Error())
		}
		if mt == nil {
			return errors.New("Demand is nil ")
		}

		mtp := new(models.MobileTowerPayments)
		mtp.TnxID = s.Shared.GetTransactionID(mt.PropertyMobileTower.PropertyID, 32)
		tnxId = mtp.TnxID

		mtp.PropertyID = mt.PropertyID
		mtp.MobileTowerID = mt.PropertyMobileTower.UniqueID
		mtp.Scenario = constants.MOBILETOWERPAYMENTREGISTRATIONPAYMENT

		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		mtp.FinancialYear = fy.FinancialYear
		ref, err := s.Daos.GetSingleDefaultMobileTowerRegistrationTax(ctx)
		if err != nil {
			return errors.New("Error in getting default mobiletower registration tax - " + err.Error())
		}
		if ref == nil {
			return errors.New("mobile tower registration tax is nil")
		}

		mtp.Demand = models.MobileTowerPaymentDemand{

			Total: models.MobileTowerPaymentDemandSplitage{
				Tax:     ref.MobileTowerRegistrationTax.Value,
				Penalty: 0,
				Other:   0,
				Total:   ref.MobileTowerRegistrationTax.Value,
			},
		}
		mtp.Status = constants.MOBILETOWERPAYMENRSTATUSINIT
		mtp.Address = mt.Address
		mtp.Created = models.CreatedV2{
			By:     imtpr.By,
			ByType: imtpr.ByType,
			On:     &t,
		}

		var mtpBasic = new(models.MobileTowerPaymentsBasics)
		mtpBasic.TnxID = mtp.TnxID
		if mt.PropertyID != "" {
			mtpBasic.PropertyID = mt.PropertyID
			mtpBasic.MobileTowerID = mt.PropertyMobileTower.UniqueID
			property, err := s.Daos.GetSingleProperty(ctx, mt.PropertyID)
			if err != nil {
				return errors.New("Error in getting property of mobile tower - " + err.Error())
			}
			mtpBasic.Property = property
		}

		mobiletower, err := s.Daos.GetSinglePropertyMobileTower(ctx, mt.PropertyMobileTower.UniqueID)
		if err != nil {
			return errors.New("Error in getting property of mobile tower  - " + err.Error())
		}
		if mobiletower != nil {
			mtpBasic.MobileTower = *mobiletower
		}
		err = s.Daos.SaveMobileTowerPayment(ctx, mtp)
		if err != nil {
			return errors.New("Error in saving mobile tower payment - " + err.Error())
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

// MakeMobileTowerPaymentForRegistration : ""
func (s *Service) MakeMobileTowerPaymentForRegistration(ctx *models.Context, mmtpr *models.MakeMobileTowerPaymentReq) error {
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
		refPayment, dberr := s.Daos.GetSingleMobileTowerPayment(ctx, mmtpr.TnxID)
		if dberr != nil {
			return dberr
		}
		if status == "Completed" {

			dberr = s.Daos.EnableMobileTowerPayments(ctx, refPayment.MobileTowerID)
			if dberr != nil {
				return dbErr
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

// VerifyMobileTowerPayment : ""
func (s *Service) VerifyMobileTowerRegistrationPayment(ctx *models.Context, makeAction *models.MakeMobileTowerPaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.VerifyMobileTowerRegistrationPayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		refPayment, dberr := s.Daos.GetSingleMobileTowerPayment(ctx, makeAction.TnxID)
		if dberr != nil {
			return dberr
		}
		dberr = s.Daos.EnableMobileTowerPayments(ctx, refPayment.MobileTowerID)
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
func (s *Service) NotVerifyMobileTowerRegistrationPayment(ctx *models.Context, makeAction *models.MakeMobileTowerPaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.NotVerifyMobileTowerRegistrationPayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		refPayment, dberr := s.Daos.GetSingleMobileTowerPayment(ctx, makeAction.TnxID)
		if dberr != nil {
			return dberr
		}
		dberr = s.Daos.PendingMobileTowerPayments(ctx, refPayment.MobileTowerID)
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
func (s *Service) RejectMobileTowerRegistrationPayment(ctx *models.Context, makeAction *models.MakeMobileTowerPaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.RejectMobileTowerRegistrationPayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}

		refPayment, dberr := s.Daos.GetSingleMobileTowerPayment(ctx, makeAction.TnxID)
		if dberr != nil {
			return dberr
		}
		dberr = s.Daos.PendingMobileTowerPayments(ctx, refPayment.MobileTowerID)
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

func (s *Service) GetMobileTowerRegistartionPaymentReceiptsPDF(ctx *models.Context, ID string) ([]byte, error) {

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
	templatePath := templatePathStart + "mobiletower_registrationreceipt.html"
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
