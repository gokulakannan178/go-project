package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveWhatsappLog :""
func (s *Service) SaveWhatsappLog(ctx *models.Context, whatsapplog *models.WhatsappLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	whatsapplog.Status = constants.WHATSAPPLOGSTATUSACTIVE
	whatsapplog.IsJob = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 WhatsappLog.created")
	whatsapplog.Created = &created
	log.Println("b4 WhatsappLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWhatsappLog(ctx, whatsapplog)
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

//UpdateOrganisation : ""
func (s *Service) UpdateWhatsappLog(ctx *models.Context, whatsapplog *models.WhatsappLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateWhatsappLog(ctx, whatsapplog)
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

//EnableOrganisation : ""
func (s *Service) EnableWhatsappLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableWhatsappLog(ctx, UniqueID)
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

//DisableWhatsappLog : ""
func (s *Service) DisableWhatsappLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableWhatsappLog(ctx, UniqueID)
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

//DeleteWhatsappLog : ""
func (s *Service) DeleteWhatsappLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWhatsappLog(ctx, UniqueID)
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

//GetSingleWhatsappLog :""
func (s *Service) GetSingleWhatsappLog(ctx *models.Context, UniqueID string) (*models.RefWhatsappLog, error) {
	whatsapplog, err := s.Daos.GetSingleWhatsappLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return whatsapplog, nil
}

//FilterWhatsappLog :""
func (s *Service) FilterWhatsappLog(ctx *models.Context, whatsapplogfilter *models.WhatsappLogFilter, pagination *models.Pagination) (whatsapplog []models.RefWhatsappLog, err error) {
	return s.Daos.FilterWhatsappLog(ctx, whatsapplogfilter, pagination)
}
