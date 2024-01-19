package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveAreaAssignLog(ctx *models.Context, areaassignlog *models.AreaAssignLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	areaassignlog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONAREAASSIGNLOG)
	areaassignlog.Status = constants.AREAASSIGNLOGSTATUSACTIVE
	//user.Password = "#nature32" //Default Password
	t := time.Now()
	areaassignlog.StartDate = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	areaassignlog.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveAreaAssignLog(ctx, areaassignlog)
		if dberr != nil {

			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

//GetSingleAreaAssignLog :""
func (s *Service) GetSingleAreaAssignLog(ctx *models.Context, UniqueID string) (*models.RefAreaAssignLog, error) {
	user, err := s.Daos.GetSingleAreaAssignLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//UpdateAreaAssignLog : ""
func (s *Service) UpdateAreaAssignLog(ctx *models.Context, areaassignlog *models.AreaAssignLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAreaAssignLog(ctx, areaassignlog)
		if err != nil {
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())

			//return errors.New("Transaction Aborted - " + err.Error())
		}
		return err
	}
	return nil
}

//EnableAreaAssignLog : ""
func (s *Service) EnableAreaAssignLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableAreaAssignLog(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}

		return err
	}
	return nil
}

//DisableAreaAssignLog : ""
func (s *Service) DisableAreaAssignLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableAreaAssignLog(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

//DeleteAreaAssignLog : ""
func (s *Service) DeleteAreaAssignLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteAreaAssignLog(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

//AreaAssignLogFilter:""
func (s *Service) AreaAssignLogFilter(ctx *models.Context, filter *models.AreaAssignLogFilter, pagination *models.Pagination) (user []models.AreaAssignLog, err error) {
	return s.Daos.AreaAssignLogFilter(ctx, filter, pagination)

}
