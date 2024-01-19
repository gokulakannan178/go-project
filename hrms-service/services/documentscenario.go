package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDocumentScenario :""
func (s *Service) SaveDocumentScenario(ctx *models.Context, documentScenario *models.DocumentScenario) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	documentScenario.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDOCUMENTSCENARIO)
	documentScenario.Status = constants.DOCUMENTSCENARIOSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DocumentScenario.created")
	documentScenario.Created = created
	log.Println("b4 DocumentScenario.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDocumentScenario(ctx, documentScenario)
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

//UpdateDocumentScenario : ""
func (s *Service) UpdateDocumentScenario(ctx *models.Context, documentScenario *models.DocumentScenario) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDocumentScenario(ctx, documentScenario)
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

//EnableDocumentScenario : ""
func (s *Service) EnableDocumentScenario(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDocumentScenario(ctx, UniqueID)
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

//DisableDocumentScenario : ""
func (s *Service) DisableDocumentScenario(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDocumentScenario(ctx, UniqueID)
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

//DeleteDocumentScenario : ""
func (s *Service) DeleteDocumentScenario(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDocumentScenario(ctx, UniqueID)
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

//GetSingleDocumentScenario :""
func (s *Service) GetSingleDocumentScenario(ctx *models.Context, UniqueID string) (*models.RefDocumentScenario, error) {
	documentScenario, err := s.Daos.GetSingleDocumentScenario(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return documentScenario, nil
}

//FilterDocumentScenario :""
func (s *Service) FilterDocumentScenario(ctx *models.Context, filter *models.FilterDocumentScenario, pagination *models.Pagination) ([]models.RefDocumentScenario, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDocumentScenario(ctx, filter, pagination)

}
