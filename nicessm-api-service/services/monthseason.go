package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveMonthSeason :""
func (s *Service) SaveMonthSeason(ctx *models.Context, MonthSeason *models.MonthSeason) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	MonthSeason.Status = constants.MONTHSEASONSTATUSACTIVE
	MonthSeason.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 MonthSeason.created")
	MonthSeason.Created = &created
	log.Println("b4 MonthSeason.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMonthSeason(ctx, MonthSeason)
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

//UpdateMonthSeason : ""
func (s *Service) UpdateMonthSeason(ctx *models.Context, MonthSeason *models.MonthSeason) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateMonthSeason(ctx, MonthSeason)
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

//EnableMonthSeason : ""
func (s *Service) EnableMonthSeason(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableMonthSeason(ctx, UniqueID)
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

//DisableMonthSeason : ""
func (s *Service) DisableMonthSeason(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableMonthSeason(ctx, UniqueID)
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

//DeleteMonthSeason : ""
func (s *Service) DeleteMonthSeason(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMonthSeason(ctx, UniqueID)
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

//GetSingleMonthSeason :""
func (s *Service) GetSingleMonthSeason(ctx *models.Context, UniqueID string) (*models.RefMonthSeason, error) {
	MonthSeason, err := s.Daos.GetSingleMonthSeason(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return MonthSeason, nil
}

//FilterMonthSeason :""
func (s *Service) FilterMonthSeason(ctx *models.Context, MonthSeasonfilter *models.MonthSeasonFilter, pagination *models.Pagination) (MonthSeason []models.RefMonthSeason, err error) {
	return s.Daos.FilterMonthSeason(ctx, MonthSeasonfilter, pagination)
}
