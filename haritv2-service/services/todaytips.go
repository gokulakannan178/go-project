package services

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveTodayTips :""
func (s *Service) SaveTodayTips(ctx *models.Context, TodayTips *models.TodayTips) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	TodayTips.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTODAYTIPS)
	TodayTips.Status = constants.TODAYTIPSSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 TodayTips.created")
	TodayTips.Created = &created
	log.Println("b4 TodayTips.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTodayTips(ctx, TodayTips)
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

//UpdateTodayTips : ""
func (s *Service) UpdateTodayTips(ctx *models.Context, TodayTips *models.TodayTips) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTodayTips(ctx, TodayTips)
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

//EnableTodayTips : ""
func (s *Service) EnableTodayTips(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTodayTips(ctx, UniqueID)
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

//DisableTodayTips : ""
func (s *Service) DisableTodayTips(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTodayTips(ctx, UniqueID)
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

//DeleteTodayTips : ""
func (s *Service) DeleteTodayTips(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTodayTips(ctx, UniqueID)
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

//GetSingleTodayTips :""
func (s *Service) GetSingleTodayTips(ctx *models.Context, UniqueID string) (*models.RefTodayTips, error) {
	TodayTips, err := s.Daos.GetSingleTodayTips(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return TodayTips, nil
}

//FilterTodayTips :""
func (s *Service) FilterTodayTips(ctx *models.Context, filter *models.TodayTipsFilter, pagination *models.Pagination) ([]models.RefTodayTips, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterTodayTips(ctx, filter, pagination)

}

//GetSingleTodayTips :""
func (s *Service) GetTodayTips(ctx *models.Context) (*models.RefTodayTips, error) {
	TodayTips, err := s.Daos.GetTodayTips(ctx)
	if err != nil {
		return nil, err
	}
	return TodayTips, nil
}
