package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// MakeMobileTowerLegacyRegistrationPayment : ""
func (s *Service) MakeMobileTowerLegacyRegistrationPayment(ctx *models.Context, mobile *models.MobileTowerPayments) error {
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// mobile tower payment basics
		mt, err := s.Daos.GetSingleMobileTower(ctx, mobile.MobileTowerID)
		if err != nil {
			return errors.New("Error in calculating demand - " + err.Error())
		}
		if mt == nil {
			return errors.New("Mobile Tower is nil ")
		}
		var mtpBasic = new(models.MobileTowerPaymentsBasics)
		mtpBasic.TnxID = mobile.TnxID
		mtpBasic.MobileTower = *mt
		mtpBasic.Status = constants.MOBILETOWERPAYMENRSTATUSCOMPLETED
		if mt.PropertyID != "" {
			mtpBasic.PropertyID = mt.PropertyID
			mtpBasic.MobileTowerID = mt.PropertyMobileTower.UniqueID
			property, err := s.Daos.GetSingleProperty(ctx, mt.PropertyID)
			if err != nil {
				return errors.New("Error in getting property of mobile tower - " + err.Error())
			}
			mtpBasic.Property = property
		}
		//

		mobile.TnxID = s.Shared.GetTransactionID(mobile.MobileTowerID, 32)
		mobile.ReciptNo = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERPAYMENTSRECEIPT)
		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		mobile.FinancialYear = fy.FinancialYear
		mobile.Status = constants.MOBILETOWERPAYMENRSTATUSCOMPLETED
		res, err := s.Daos.GetSinglePropertyMobileTower(ctx, mobile.MobileTowerID)
		if err != nil {
			return errors.New("Error in geting mobile tower - " + err.Error())
		}
		mobile.Address = res.Address
		created := models.CreatedV2{}
		t := time.Now()
		created.On = &t
		created.By = constants.SYSTEM
		mobile.Created = created
		mobile.Scenario = constants.MOBILETOWERPAYMENTLEGACYREGISTRATIONPAYMENT
		if mobile.CompletionDate == nil {
			mobile.CompletionDate = &t
		}

		//=================
		err = s.Daos.EnablePropertyMobileTower(ctx, mobile.MobileTowerID)
		if err != nil {
			return err
		}

		err = s.Daos.SaveMobileTowerPayment(ctx, mobile)
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
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil

}
