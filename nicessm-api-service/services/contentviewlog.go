package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveContentViewLog :""
func (s *Service) SaveContentViewLog(ctx *models.Context, ContentViewLog *models.ContentViewLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	ContentViewLog.Status = constants.CONTENTVIEWLOGSTATUSACTIVE
	ContentViewLog.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	ContentViewLog.Date = &t
	ContentViewLog.UniqueId = fmt.Sprintf("%v_%v_%v", t.Day(), t.Month(), t.Year())
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ContentViewLog.created")
	ContentViewLog.Created = &created
	log.Println("b4 ContentViewLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveContentViewLog(ctx, ContentViewLog)
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

//UpdateContentViewLog : ""
func (s *Service) UpdateContentViewLog(ctx *models.Context, ContentViewLog *models.ContentViewLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateContentViewLog(ctx, ContentViewLog)
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

//EnableContentViewLog : ""
func (s *Service) EnableContentViewLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableContentViewLog(ctx, UniqueID)
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

//DisableContentViewLog : ""
func (s *Service) DisableContentViewLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableContentViewLog(ctx, UniqueID)
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

//DeleteContentViewLog : ""
func (s *Service) DeleteContentViewLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteContentViewLog(ctx, UniqueID)
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

//GetSingleContentViewLog :""
func (s *Service) GetSingleContentViewLog(ctx *models.Context, UniqueID string) (*models.RefContentViewLog, error) {
	ContentViewLog, err := s.Daos.GetSingleContentViewLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ContentViewLog, nil
}

//FilterContentViewLog :""
func (s *Service) FilterContentViewLog(ctx *models.Context, ContentViewLogfilter *models.ContentViewLogFilter, pagination *models.Pagination) (ContentViewLog []models.RefContentViewLog, err error) {
	return s.Daos.FilterContentViewLog(ctx, ContentViewLogfilter, pagination)
}
func (s *Service) DayWiseContentViewChart(ctx *models.Context, contentfilter *models.FilterDaywiseViewChart) (content *models.DayWiseContentViewChartReport, err error) {
	if contentfilter.CreatedFrom.StartDate == nil {
		t := time.Now()
		contentfilter.CreatedFrom.StartDate = &t
	}
	sd := s.Shared.BeginningOfMonth(*contentfilter.CreatedFrom.StartDate)
	ed := s.Shared.EndOfMonth(sd)
	fmt.Println(sd.Month(), " & ", ed.Month())
	contentfilter.CreatedFrom.StartDate = &sd
	contentfilter.CreatedFrom.EndDate = &ed
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.DayWiseContentViewChart(ctx, contentfilter)
}
