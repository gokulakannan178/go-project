package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTradeLicense : ""
func (s *Service) SaveTradeLicense(ctx *models.Context, trade *models.TradeLicense) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	trade.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSE)
	trade.Status = constants.TRADELICENSESTATUSPENDING
	t := time.Now()
	trade.Created.On = &t
	trade.Created.By = constants.SYSTEM
	trade.LicenseDate = &t
	expiryDate := t.AddDate(1, 0, 0)
	trade.LicenseExpiryDate = &expiryDate
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTradeLicense(ctx, trade)
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

// SaveTradeLicenseV2 : ""
func (s *Service) SaveTradeLicenseV2(ctx *models.Context, trade *models.TradeLicense) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return err1
	}
	trade.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSE)
	if resPD.LocationID == "Dev" {
		trade.Status = constants.TRADELICENSESTATUSINIT
	}
	if resPD.LocationID == "Bhagalpur" {
		trade.Status = constants.TRADELICENSESTATUSINIT
	} else {
		trade.Status = constants.TRADELICENSESTATUSPENDING
	}

	t := time.Now()
	trade.Created.On = &t
	//trade.Created.By = constants.SYSTEM

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTradeLicense(ctx, trade)
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

//GetSingleTradeLicense :""
func (s *Service) GetSingleTradeLicense(ctx *models.Context, UniqueID string) (*models.RefTradeLicense, error) {
	tower, err := s.Daos.GetSingleTradeLicense(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateTradeLicense : ""
func (s *Service) UpdateTradeLicense(ctx *models.Context, trade *models.TradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	// t := time.Now()
	// trade.Created.On = &t
	// trade.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTradeLicense(ctx, trade)
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

// EnableTradeLicense : ""
func (s *Service) EnableTradeLicense(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTradeLicense(ctx, UniqueID)
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

//DisableTradeLicense : ""
func (s *Service) DisableTradeLicense(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTradeLicense(ctx, UniqueID)
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

//DeleteTradeLicense : ""
func (s *Service) DeleteTradeLicense(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTradeLicense(ctx, UniqueID)
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

// RejectedTradeLicense : ""
func (s *Service) RejectedTradeLicense(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectedTradeLicense(ctx, UniqueID)
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

// FilterTradeLicense : ""
func (s *Service) FilterTradeLicense(ctx *models.Context, filter *models.TradeLicenseFilter, pagination *models.Pagination) ([]models.RefTradeLicense, error) {
	return s.Daos.FilterTradeLicense(ctx, filter, pagination)
}

// VerifyTradeLicensePayment : ""
func (s *Service) VerifyTradeLicensePayment(ctx *models.Context, makeAction *models.MakeTradeLicensePaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.VerifyTradeLicensePayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		dberr1 := s.UpdateLicenseExpiry(ctx, makeAction.TnxID)
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

// NotVerifyTradeLicensePayment : ""
func (s *Service) NotVerifyTradeLicensePayment(ctx *models.Context, makeAction *models.MakeTradeLicensePaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.NotVerifyTradeLicensePayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		dberr1 := s.UpdateTradeLicenseExpiryOnReject(ctx, makeAction.TnxID)
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

// RejectTradeLicensePayment : ""
func (s *Service) RejectTradeLicensePayment(ctx *models.Context, makeAction *models.MakeTradeLicensePaymentsAction) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		dberr := s.Daos.RejectTradeLicensePayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
		dberr1 := s.UpdateTradeLicenseExpiryOnReject(ctx, makeAction.TnxID)
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

func (s *Service) UpdateTradeLicenseDemandAndCollections(ctx *models.Context, shopID string) error {

	demand, err := s.TradeLicenseGetOutstandingDemand(ctx, shopID)
	if err != nil {
		return errors.New("Eror in calc demand - " + err.Error())
	}
	demand.TradeLicense.Collections = models.TradeLicenseTotalCollection{}
	demand.TradeLicense.PendingCollections = models.TradeLicenseTotalCollection{}
	demand.TradeLicense.OutStanding = models.TradeLicenseTotalOutStanding{}
	mainpipeline, err := demand.CalcCollectionQuery()
	if err != nil {
		return errors.New("Error in Calc generating Query - " + err.Error())
	}
	payments, err := s.Daos.CalcTradeLicensePaymens(ctx, mainpipeline)
	if err != nil {
		return errors.New("Error in Payments - " + err.Error())
	}
	err = demand.CalcCollection(payments)
	if err != nil {
		return errors.New("Error in Calc Collection - " + err.Error())
	}
	mainpipeline2, err := demand.CalcPendingCollectionQuery()
	pendingpayments, err := s.Daos.CalcTradeLicensePendingPaymens(ctx, mainpipeline2)
	if err != nil {
		return errors.New("Error in Pending Payments - " + err.Error())
	}
	err = demand.CalcPendingCollection(pendingpayments)
	if err != nil {
		return errors.New("Error in Calc pending Collection - " + err.Error())
	}
	err = demand.CalcOutStanding()
	if err != nil {
		return errors.New("Error in Calc outstanding demand - " + err.Error())
	}
	err = s.Daos.UpdateTradeLicenseCalc(ctx, demand)
	if err != nil {
		return errors.New("Error in updating demand - " + err.Error())
	}
	return nil

}

func (s *Service) UpdateTradeLicenseDemandAndCollectionsWithTnx(ctx *models.Context, shopID string) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dbErr := s.UpdateTradeLicenseDemandAndCollectionsWithTnx(ctx, shopID)
		if dbErr != nil {
			return dbErr
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

// ApproveTradeLicense : ""
func (s *Service) ApproveTradeLicense(ctx *models.Context, approve *models.ApproveTradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		approve.On = &t

		err := s.Daos.ApproveTradeLicense(ctx, approve)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// NotApproveTradeLicense : ""
func (s *Service) NotApproveTradeLicense(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		notApprove.On = &t

		err := s.Daos.NotApproveTradeLicense(ctx, notApprove)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

func (s *Service) GetTradeLicenseSAFDashboard(ctx *models.Context, filter *models.GetTradeLicenseSAFDashboardFilter) (*models.TradeLicenseSAFDashboard, error) {
	return s.Daos.GetTradeLicenseSAFDashboard(ctx, filter)
}

// ApproveTradeLicense : ""
func (s *Service) VerifyTradeLicense(ctx *models.Context, approve *models.ApproveTradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		approve.On = &t

		err := s.Daos.VerifyTradeLicense(ctx, approve)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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
