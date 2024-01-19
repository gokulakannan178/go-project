package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveTradeLicenseRebate :""
func (s *Service) SaveTradeLicenseRebate(ctx *models.Context, tlRebate *models.TradeLicenseRebate) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	tlRebate.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSEREBATE)
	tlRebate.Status = constants.TRADELICENSEREBATESTATUSACTIVE
	t := time.Now()
	created := new(models.Created)
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 tlRebate.created")
	tlRebate.Created = created
	log.Println("b4 tlRebate.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTradeLicenseRebate(ctx, tlRebate)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//UpdateTradeLicenseRebate : ""
func (s *Service) UpdateTradeLicenseRebate(ctx *models.Context, tlRebate *models.TradeLicenseRebate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTradeLicenseRebate(ctx, tlRebate)
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

//EnableTradeLicenseRebate : ""
func (s *Service) EnableTradeLicenseRebate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTradeLicenseRebate(ctx, UniqueID)
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

//DisableTradeLicenseRebate : ""
func (s *Service) DisableTradeLicenseRebate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTradeLicenseRebate(ctx, UniqueID)
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

//DeleteTradeLicenseRebate : ""
func (s *Service) DeleteTradeLicenseRebate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTradeLicenseRebate(ctx, UniqueID)
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

//GetSingleTradeLicenseRebate :""
func (s *Service) GetSingleTradeLicenseRebate(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseRebate, error) {
	tlRebate, err := s.Daos.GetSingleTradeLicenseRebate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tlRebate, nil
}

//FilterTradeLicenseRebate :""
func (s *Service) FilterTradeLicenseRebate(ctx *models.Context, tlRebatefilter *models.TradeLicenseRebateFilter, pagination *models.Pagination) (tlRebate []models.RefTradeLicenseRebate, err error) {
	return s.Daos.FilterTradeLicenseRebate(ctx, tlRebatefilter, pagination)
}
