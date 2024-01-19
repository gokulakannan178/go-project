package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePropertyWallet :""
func (s *Service) SavePropertyWallet(ctx *models.Context, propertyWallet *models.PropertyWallet) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	propertyWallet.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYWALLET)
	propertyWallet.Status = constants.PROPERTYWALLETSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 propertyWallet.created")
	propertyWallet.Created = &created
	log.Println("b4 propertyWallet.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePropertyWallet(ctx, propertyWallet)
		if dberr != nil {
			return dberr
		}

		walletLog := new(models.PropertyWalletLog)
		walletLog.Status = constants.PROPERTYWALLETSTATUSACTIVE
		walletLog.Created = &created
		walletLog.Date = &t
		walletLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYWALLETLOG)
		walletLog.WalletID = propertyWallet.UniqueID
		walletLog.PropertyID = propertyWallet.PropertyID
		walletLog.Amount = propertyWallet.Amount
		walletLog.MobileNo = propertyWallet.MobileNo
		walletLog.OwnerName = propertyWallet.OwnerName
		err := s.Daos.SavePropertyWalletLog(ctx, walletLog)
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

//GetSinglePropertyWallet :""
func (s *Service) GetSinglePropertyWallet(ctx *models.Context, UniqueID string) (*models.RefPropertyWallet, error) {
	propertyWallet, err := s.Daos.GetSinglePropertyWallet(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertyWallet, nil
}

//UpdatePropertyWallet : ""
func (s *Service) UpdatePropertyWallet(ctx *models.Context, propertyWallet *models.PropertyWallet) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyWallet(ctx, propertyWallet)
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

//EnablePropertyWallet : ""
func (s *Service) EnablePropertyWallet(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyWallet(ctx, UniqueID)
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

//DisablePropertyWallet : ""
func (s *Service) DisablePropertyWallet(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyWallet(ctx, UniqueID)
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

//DeletePropertyWallet : ""
func (s *Service) DeletePropertyWallet(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyWallet(ctx, UniqueID)
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

//FilterPropertyWallet :""
func (s *Service) FilterPropertyWallet(ctx *models.Context, filter *models.PropertyWalletFilter, pagination *models.Pagination) (propertyWallet []models.RefPropertyWallet, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPropertyWallet(ctx, filter, pagination)
}
