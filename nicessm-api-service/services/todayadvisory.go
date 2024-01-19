package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) SaveTodayAdvisory(ctx *models.Context, content *models.Content) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	content.Status = constants.MARKETSTATUSACTIVE
	content.ActiveStatus = true
	//market.District = primitive.ObjectID
	t := time.Now()
	//created := models.Created{}
	//created.On = &t
	//created.By = constants.SYSTEM
	log.Println("b4 organisation.created")
	content.DateCreated = &t
	//content.DateReviewed = &t
	log.Println("b4 organisation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTodayAdvisory(ctx, content)
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

func (s *Service) UpdateTodayAdvisory(ctx *models.Context, content *models.Content) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTodayAdvisory(ctx, content)
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

func (s *Service) EnableTodayAdvisory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTodayAdvisory(ctx, UniqueID)
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

func (s *Service) DisableTodayAdvisory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTodayAdvisory(ctx, UniqueID)
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

func (s *Service) DeleteTodayAdvisory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTodayAdvisory(ctx, UniqueID)
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

func (s *Service) GetSingleTodayAdvisory(ctx *models.Context, UniqueID string) (*models.RefContent, error) {
	content, err := s.Daos.GetSingleTodayAdvisory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *Service) FilterTodayAdvisory(ctx *models.Context, todayadvisoryfilter *models.ContentFilter, pagination *models.Pagination) (content []models.RefContent, err error) {
	return s.Daos.FilterTodayAdvisory(ctx, todayadvisoryfilter, pagination)
}
func (s *Service) GetTodayAdvisory(ctx *models.Context) ([]models.RefContent, error) {
	content, err := s.Daos.GetTodayAdvisory(ctx)
	if err != nil {
		return nil, err
	}
	return content, nil
}
