package services

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveAdvertisement :""
func (s *Service) SaveAdvertisement(ctx *models.Context, advertisement *models.Advertisement) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	advertisement.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONADVERTISEMENT)
	advertisement.Status = constants.ADVERTISEMENTOWNERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 advertisement.created")
	advertisement.Created = &created
	log.Println("b4 advertisement.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAdvertisement(ctx, advertisement)
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

//UpdateAdvertisement : ""
func (s *Service) UpdateAdvertisement(ctx *models.Context, advertisement *models.Advertisement) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAdvertisement(ctx, advertisement)
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

//EnableAdvertisement : ""
func (s *Service) EnableAdvertisement(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAdvertisement(ctx, UniqueID)
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

//DisableAdvertisement : ""
func (s *Service) DisableAdvertisement(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAdvertisement(ctx, UniqueID)
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

//DeleteAdvertisement : ""
func (s *Service) DeleteAdvertisement(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAdvertisement(ctx, UniqueID)
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

//GetSingleAdvertisement :""
func (s *Service) GetSingleAdvertisement(ctx *models.Context, UniqueID string) (*models.RefAdvertisement, error) {
	advertisement, err := s.Daos.GetSingleAdvertisement(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return advertisement, nil
}

//FilterAdvertisement :""
func (s *Service) FilterAdvertisement(ctx *models.Context, advertisementfilter *models.AdvertisementFilter, pagination *models.Pagination) (advertisement []models.RefAdvertisement, err error) {
	return s.Daos.FilterAdvertisement(ctx, advertisementfilter, pagination)
}
