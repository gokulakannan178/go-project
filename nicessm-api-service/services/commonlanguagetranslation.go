package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveCommonLanguageTranslations :""
func (s *Service) SaveCommonLanguageTranslations(ctx *models.Context, CommonLanguageTranslations *models.CommonLanguageTranslationss) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//CommonLanguageTranslations.ActiveStatus = true
	CommonLanguageTranslations.Status = constants.LANGAUAGESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 organisation.created")
	CommonLanguageTranslations.Created = &created
	log.Println("b4 organisation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveCommonLanguageTranslations(ctx, CommonLanguageTranslations)
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

//UpdateCommonLanguageTranslations: ""
func (s *Service) UpdateCommonLanguageTranslations(ctx *models.Context, commonlanguagetranslations *models.CommonLanguageTranslationss) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCommonLanguageTranslations(ctx, commonlanguagetranslations)
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

//EnableCommonLanguageTranslations : ""
func (s *Service) EnableCommonLanguageTranslations(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableCommonLanguageTranslations(ctx, UniqueID)
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

//DisableCommonLanguageTranslations : ""
func (s *Service) DisableCommonLanguageTranslations(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableCommonLanguageTranslations(ctx, UniqueID)
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

//DeleteCommonLanguageTranslations : ""
func (s *Service) DeleteCommonLanguageTranslations(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCommonLanguageTranslations(ctx, UniqueID)
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

//GetSingleCommonLanguageTranslations :""
func (s *Service) GetSingleCommonLanguageTranslations(ctx *models.Context, UniqueID string) (*models.RefCommonLanguageTranslations, error) {
	commonlanguagetranslations, err := s.Daos.GetSingleCommonLanguageTranslations(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return commonlanguagetranslations, nil
}

//FilterCommonLanguageTranslations :""
func (s *Service) FilterCommonLanguageTranslations(ctx *models.Context, organisationfilter *models.CommonLanguageTranslationsFilter, pagination *models.Pagination) (organisation []models.RefCommonLanguageTranslations, err error) {
	return s.Daos.FilterCommonLanguageTranslations(ctx, organisationfilter, pagination)
}
func (s *Service) GetSingleCommonLanguageTranslationsWithType(ctx *models.Context, UniqueID string) (*models.RefCommonLanguageTranslations, error) {
	commonlanguagetranslations, err := s.Daos.GetSingleCommonLanguageTranslationsWithType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return commonlanguagetranslations, nil
}
