package services

import (
	"errors"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveOffenceType :""
func (s *Service) SaveOffenceType(ctx *models.Context, offenceType *models.OffenceType) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	offenceType.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLOFFENCETYPE)
	offenceType.Status = constants.OFFENCETYPESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 offenceType.created")
	offenceType.Created = created
	log.Println("b4 offenceType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOffenceType(ctx, offenceType)
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

//UpdateOffenceType : ""
func (s *Service) UpdateOffenceType(ctx *models.Context, offenceType *models.OffenceType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOffenceType(ctx, offenceType)
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

//EnableOffenceType : ""
func (s *Service) EnableOffenceType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableOffenceType(ctx, UniqueID)
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

//DisableOffenceType : ""
func (s *Service) DisableOffenceType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableOffenceType(ctx, UniqueID)
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

//DeleteOffenceType : ""
func (s *Service) DeleteOffenceType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOffenceType(ctx, UniqueID)
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

//GetSingleOffenceType :""
func (s *Service) GetSingleOffenceType(ctx *models.Context, UniqueID string) (*models.RefOffenceType, error) {
	offenceType, err := s.Daos.GetSingleOffenceType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return offenceType, nil
}

//FilterOffenceType :""
func (s *Service) FilterOffenceType(ctx *models.Context, offenceTypefilter *models.OffenceTypeFilter, pagination *models.Pagination) (offenceType []models.RefOffenceType, err error) {
	return s.Daos.FilterOffenceType(ctx, offenceTypefilter, pagination)
}
