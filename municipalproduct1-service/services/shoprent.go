package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveShopRent :""
func (s *Service) SaveShopRent(ctx *models.Context, shoprent *models.ShopRent) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	shoprent.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSHOPRENT)
	shoprent.Status = constants.SHOPRENTSTATUSPENDING
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	if shoprent.Created != nil {
		if shoprent.Created.By != "" {
			created.By = shoprent.Created.By
		}

	}

	log.Println("b4 shoprent.created")
	shoprent.Created = &created
	log.Println("b4 shoprent.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveShopRent(ctx, shoprent)
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

//GetSingleShopRent :""
func (s *Service) GetSingleShopRent(ctx *models.Context, UniqueID string) (*models.RefShopRent, error) {
	shoprent, err := s.Daos.GetSingleShopRent(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return shoprent, nil
}

//UpdateShopRent : ""
func (s *Service) UpdateShopRent(ctx *models.Context, shoprent *models.ShopRent) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateShopRent(ctx, shoprent)
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

//EnableShopRent : ""
func (s *Service) EnableShopRent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableShopRent(ctx, UniqueID)
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

// RejectedShopRent : ""
func (s *Service) RejectedShopRent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectedShopRent(ctx, UniqueID)
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

//DisableShopRent : ""
func (s *Service) DisableShopRent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableShopRent(ctx, UniqueID)
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

//DeleteShopRent : ""
func (s *Service) DeleteShopRent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteShopRent(ctx, UniqueID)
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

//FilterShopRent :""
func (s *Service) FilterShopRent(ctx *models.Context, filter *models.ShopRentFilter, pagination *models.Pagination) ([]models.RefShopRent, error) {
	return s.Daos.FilterShopRent(ctx, filter, pagination)
}

// VerifyShopRentPayment : ""
func (s *Service) VerifyShopRentPayment(ctx *models.Context, makeAction *models.MakeShopRentPaymentsAction) (string, error) {
	log.Println("transaction start")
	shoprentID := ""
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		var dberr error

		shoprentID, dberr = s.Daos.VerifyShopRentPayment(ctx, makeAction)
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
	return shoprentID, nil
}

// NotVerifyShopRentPayment : ""
func (s *Service) NotVerifyShopRentPayment(ctx *models.Context, makeAction *models.MakeShopRentPaymentsAction) (string, error) {
	shoprentID := ""
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

		shoprentID, dberr = s.Daos.NotVerifyShopRentPayment(ctx, makeAction)
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
	return shoprentID, nil
}

// RejectShopRentPayment : ""
func (s *Service) RejectShopRentPayment(ctx *models.Context, makeAction *models.MakeShopRentPaymentsAction) (string, error) {
	log.Println("transaction start")
	shoprentID := ""
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		var dberr error
		shoprentID, dberr = s.Daos.RejectShopRentPayment(ctx, makeAction)
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
	return shoprentID, nil
}

func (s *Service) UpdateShopRentDemandAndCollections(ctx *models.Context, shopID string) error {

	demand, err := s.ShopRentGetOutstandingDemand(ctx, shopID)
	if err != nil {
		return errors.New("Eror in calc demand - " + err.Error())
	}
	demand.ShopRent.Collections = models.ShopRentTotalCollection{}
	demand.ShopRent.PendingCollections = models.ShopRentTotalCollection{}
	demand.ShopRent.OutStanding = models.ShopRentTotalOutStanding{}
	mainpipeline, err := demand.CalcCollectionQuery()
	if err != nil {
		return errors.New("Error in Calc generating Query - " + err.Error())
	}
	payments, err := s.Daos.CalcShopRentPaymens(ctx, mainpipeline)
	if err != nil {
		return errors.New("Error in Payments - " + err.Error())
	}
	err = demand.CalcCollection(payments)
	if err != nil {
		return errors.New("Error in Calc Collection - " + err.Error())
	}
	mainpipeline2, err := demand.CalcPendingCollectionQuery()
	pendingpayments, err := s.Daos.CalcShopRentPendingPaymens(ctx, mainpipeline2)
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
	err = s.Daos.UpdateShopRentCalc(ctx, demand)
	if err != nil {
		return errors.New("Error in updating demand - " + err.Error())
	}
	return nil

}

func (s *Service) UpdateShopRentDemandAndCollectionsWithTnx(ctx *models.Context, shopID string) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dbErr := s.UpdateShopRentDemandAndCollectionsWithTnx(ctx, shopID)
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
