package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveLanguageTranslation :""
func (s *Service) SaveLanguageTranslation(ctx *models.Context, LanguageTranslation *models.LanguageTranslations) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//LanguageTranslation.ActiveStatus = true
	LanguageTranslation.Status = constants.LANGAUAGESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 organisation.created")
	LanguageTranslation.Created = &created
	log.Println("b4 organisation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveLanguageTranslation(ctx, LanguageTranslation)
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

//UpdateLanguageTranslation: ""
func (s *Service) UpdateLanguageTranslation(ctx *models.Context, languagetranslation *models.LanguageTranslations) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateLanguageTranslation(ctx, languagetranslation)
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

//EnableLanguageTranslation : ""
func (s *Service) EnableLanguageTranslation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableLanguageTranslation(ctx, UniqueID)
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

//DisableLanguageTranslation : ""
func (s *Service) DisableLanguageTranslation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableLanguageTranslation(ctx, UniqueID)
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

//DeleteLanguageTranslation : ""
func (s *Service) DeleteLanguageTranslation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteLanguageTranslation(ctx, UniqueID)
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

//GetSingleLanguageTranslation :""
func (s *Service) GetSingleLanguageTranslation(ctx *models.Context, UniqueID string) (*models.RefLanguageTranslation, error) {
	languagetranslation, err := s.Daos.GetSingleLanguageTranslation(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return languagetranslation, nil
}

//FilterLanguageTranslation :""
func (s *Service) FilterLanguageTranslation(ctx *models.Context, organisationfilter *models.LanguageTranslationFilter, pagination *models.Pagination) (organisation []models.RefLanguageTranslation, err error) {
	return s.Daos.FilterLanguageTranslation(ctx, organisationfilter, pagination)
}
func (s *Service) GetSingleLanguageTranslationWithType(ctx *models.Context, UniqueID string) (*models.RefLanguageTranslation, error) {
	languagetranslation, err := s.Daos.GetSingleLanguageTranslationWithType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return languagetranslation, nil
}
