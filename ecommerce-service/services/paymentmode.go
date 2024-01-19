package services

import (
	"errors"
	"log"
	"time"

	"ecommerce-service/constants"
	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePaymentMode : ""
func (s *Service) SavePaymentMode(ctx *models.Context, PaymentMode *models.PaymentMode) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	PaymentMode.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYMENTMODE)
	PaymentMode.Status = constants.PAYMENTMODESTATUSACTIVE
	t := time.Now()

	created := new(models.Created)
	created.On = &t
	created.By = constants.SYSTEM
	PaymentMode.Created = *created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePaymentMode(ctx, PaymentMode)
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

//GetSinglePaymentMode :""
func (s *Service) GetSinglePaymentMode(ctx *models.Context, UniqueID string) (*models.RefPaymentMode, error) {
	PaymentMode, err := s.Daos.GetSinglePaymentMode(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return PaymentMode, nil
}

// UpdatePaymentMode : ""
func (s *Service) UpdatePaymentMode(ctx *models.Context, PaymentMode *models.PaymentMode) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePaymentMode(ctx, PaymentMode)
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

// EnablePaymentMode : ""
func (s *Service) EnablePaymentMode(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePaymentMode(ctx, UniqueID)
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

//DisablePaymentMode : ""
func (s *Service) DisablePaymentMode(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePaymentMode(ctx, UniqueID)
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

//DeletePaymentMode : ""
func (s *Service) DeletePaymentMode(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePaymentMode(ctx, UniqueID)
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

// FilterPaymentMode : ""
func (s *Service) FilterPaymentMode(ctx *models.Context, filter *models.PaymentModeFilter, pagination *models.Pagination) ([]models.RefPaymentMode, error) {
	return s.Daos.FilterPaymentMode(ctx, filter, pagination)

}
