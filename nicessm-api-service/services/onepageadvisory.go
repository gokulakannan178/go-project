package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveOnePageAdvisory :""
func (s *Service) SaveOnePageAdvisory(ctx *models.Context, OnePageAdvisory *models.OnePageAdvisory) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	OnePageAdvisory.Status = constants.ONEPAGEADVISORYSTATUSACTIVE
	OnePageAdvisory.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnePageAdvisory.created")
	OnePageAdvisory.Created = &created
	log.Println("b4 OnePageAdvisory.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOnePageAdvisory(ctx, OnePageAdvisory)
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

//UpdateOnePageAdvisory : ""
func (s *Service) UpdateOnePageAdvisory(ctx *models.Context, OnePageAdvisory *models.OnePageAdvisory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOnePageAdvisory(ctx, OnePageAdvisory)
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

//EnableOnePageAdvisory : ""
func (s *Service) EnableOnePageAdvisory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableOnePageAdvisory(ctx, UniqueID)
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

//DisableOnePageAdvisory : ""
func (s *Service) DisableOnePageAdvisory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableOnePageAdvisory(ctx, UniqueID)
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

//DeleteOnePageAdvisory : ""
func (s *Service) DeleteOnePageAdvisory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOnePageAdvisory(ctx, UniqueID)
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

//GetSingleOnePageAdvisory :""
func (s *Service) GetSingleOnePageAdvisory(ctx *models.Context, UniqueID string) (*models.RefOnePageAdvisory, error) {
	OnePageAdvisory, err := s.Daos.GetSingleOnePageAdvisory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return OnePageAdvisory, nil
}

//FilterOnePageAdvisory :""
func (s *Service) FilterOnePageAdvisory(ctx *models.Context, OnePageAdvisoryfilter *models.OnePageAdvisoryFilter, pagination *models.Pagination) (OnePageAdvisory []models.RefOnePageAdvisory, err error) {
	return s.Daos.FilterOnePageAdvisory(ctx, OnePageAdvisoryfilter, pagination)
}
