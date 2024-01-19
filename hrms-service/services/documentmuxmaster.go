package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDocumentMuxMaster :""
func (s *Service) SaveDocumentMuxMaster(ctx *models.Context, documentMuxMaster *models.DocumentMuxMaster) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	documentMuxMaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDOCUMENTMUXMASTER)
	documentMuxMaster.Status = constants.DOCUMENTMUXMASTERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DocumentMuxMaster.created")
	documentMuxMaster.Created = created
	log.Println("b4 DocumentMuxMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDocumentMuxMaster(ctx, documentMuxMaster)
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

//UpdateDocumentMuxMaster : ""
func (s *Service) UpdateDocumentMuxMaster(ctx *models.Context, documentMuxMaster *models.DocumentMuxMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDocumentMuxMaster(ctx, documentMuxMaster)
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

//EnableDocumentMuxMaster : ""
func (s *Service) EnableDocumentMuxMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDocumentMuxMaster(ctx, UniqueID)
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

//DisableDocumentMuxMaster : ""
func (s *Service) DisableDocumentMuxMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDocumentMuxMaster(ctx, UniqueID)
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

//DeleteDocumentMuxMaster : ""
func (s *Service) DeleteDocumentMuxMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDocumentMuxMaster(ctx, UniqueID)
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

//GetSingleDocumentMuxMaster :""
func (s *Service) GetSingleDocumentMuxMaster(ctx *models.Context, UniqueID string) (*models.RefDocumentMuxMaster, error) {
	documentMuxMaster, err := s.Daos.GetSingleDocumentMuxMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return documentMuxMaster, nil
}

//FilterDocumentMuxMaster :""
func (s *Service) FilterDocumentMuxMaster(ctx *models.Context, filter *models.FilterDocumentMuxMaster, pagination *models.Pagination) ([]models.RefDocumentMuxMaster, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDocumentMuxMaster(ctx, filter, pagination)

}
