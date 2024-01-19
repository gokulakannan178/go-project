package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePropertyRequiredDocument :""
func (s *Service) SavePropertyRequiredDocument(ctx *models.Context, propertyRequiredDocument *models.PropertyRequiredDocument) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	propertyRequiredDocument.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYREQUIREDDOCUMENT)
	propertyRequiredDocument.Status = constants.PROPERTYREQUIREDDOCUMENTSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 propertyRequiredDocument.created")
	propertyRequiredDocument.Created = created
	log.Println("b4 propertyRequiredDocument.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePropertyRequiredDocument(ctx, propertyRequiredDocument)
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

//UpdatePropertyRequiredDocument : ""
func (s *Service) UpdatePropertyRequiredDocument(ctx *models.Context, propertyRequiredDocument *models.PropertyRequiredDocument) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyRequiredDocument(ctx, propertyRequiredDocument)
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

//EnablePropertyRequiredDocument : ""
func (s *Service) EnablePropertyRequiredDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyRequiredDocument(ctx, UniqueID)
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

//DisablePropertyRequiredDocument : ""
func (s *Service) DisablePropertyRequiredDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyRequiredDocument(ctx, UniqueID)
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

//DeletePropertyRequiredDocument : ""
func (s *Service) DeletePropertyRequiredDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyRequiredDocument(ctx, UniqueID)
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

//GetSinglePropertyRequiredDocument :""
func (s *Service) GetSinglePropertyRequiredDocument(ctx *models.Context, UniqueID string) (*models.RefPropertyRequiredDocument, error) {
	propertyRequiredDocument, err := s.Daos.GetSinglePropertyRequiredDocument(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertyRequiredDocument, nil
}

//FilterPropertyRequiredDocument :""
func (s *Service) FilterPropertyRequiredDocument(ctx *models.Context, propertyRequiredDocumentfilter *models.PropertyRequiredDocumentFilter, pagination *models.Pagination) (propertyRequiredDocument []models.RefPropertyRequiredDocument, err error) {
	return s.Daos.FilterPropertyRequiredDocument(ctx, propertyRequiredDocumentfilter, pagination)
}
