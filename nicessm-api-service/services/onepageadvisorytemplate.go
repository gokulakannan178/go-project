package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveOnePageAdvisoryTemplate :""
func (s *Service) SaveOnePageAdvisoryTemplate(ctx *models.Context, OnePageAdvisoryTemplate *models.OnePageAdvisoryTemplate) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	OnePageAdvisoryTemplate.Status = constants.ONEPAGEADVISORYTEMPLATESTATUSACTIVE
	OnePageAdvisoryTemplate.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnePageAdvisoryTemplate.created")
	OnePageAdvisoryTemplate.Created = &created
	log.Println("b4 OnePageAdvisoryTemplate.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOnePageAdvisoryTemplate(ctx, OnePageAdvisoryTemplate)
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

//UpdateOnePageAdvisoryTemplate : ""
func (s *Service) UpdateOnePageAdvisoryTemplate(ctx *models.Context, OnePageAdvisoryTemplate *models.OnePageAdvisoryTemplate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOnePageAdvisoryTemplate(ctx, OnePageAdvisoryTemplate)
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

//EnableOnePageAdvisoryTemplate : ""
func (s *Service) EnableOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableOnePageAdvisoryTemplate(ctx, UniqueID)
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

//DisableOnePageAdvisoryTemplate : ""
func (s *Service) DisableOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableOnePageAdvisoryTemplate(ctx, UniqueID)
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

//DeleteOnePageAdvisoryTemplate : ""
func (s *Service) DeleteOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOnePageAdvisoryTemplate(ctx, UniqueID)
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

//GetSingleOnePageAdvisoryTemplate :""
func (s *Service) GetSingleOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) (*models.RefOnePageAdvisoryTemplate, error) {
	OnePageAdvisoryTemplate, err := s.Daos.GetSingleOnePageAdvisoryTemplate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return OnePageAdvisoryTemplate, nil
}

//FilterOnePageAdvisoryTemplate :""
func (s *Service) FilterOnePageAdvisoryTemplate(ctx *models.Context, OnePageAdvisoryTemplatefilter *models.OnePageAdvisoryTemplateFilter, pagination *models.Pagination) (OnePageAdvisoryTemplate []models.RefOnePageAdvisoryTemplate, err error) {
	return s.Daos.FilterOnePageAdvisoryTemplate(ctx, OnePageAdvisoryTemplatefilter, pagination)
}
