package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDocumentPolicyDocuments :""
func (s *Service) SaveDocumentPolicyDocuments(ctx *models.Context, documentpolicydocuments *models.DocumentPolicyDocuments) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	documentpolicydocuments.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS)
	documentpolicydocuments.Status = constants.DOCUMENTPOLICYDOCUMENTSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DocumentPolicyDocuments.created")
	documentpolicydocuments.Created = created
	log.Println("b4 DocumentPolicyDocuments.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDocumentPolicyDocuments(ctx, documentpolicydocuments)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

//SaveDocumentPolicyDocuments :""
func (s *Service) SaveDocumentPolicyDocumentsWithoutTransaction(ctx *models.Context, documentpolicydocuments *models.DocumentPolicyDocuments) error {
	log.Println("transaction start")
	//Start Transaction
	documentpolicydocuments.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS)
	documentpolicydocuments.Status = constants.DOCUMENTPOLICYDOCUMENTSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DocumentPolicyDocuments.created")
	documentpolicydocuments.Created = created
	log.Println("b4 DocumentPolicyDocuments.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDocumentPolicyDocuments(ctx, documentpolicydocuments)
		if dberr != nil {
			return dberr
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

//GetSingleDocumentPolicyDocuments :""
func (s *Service) GetSingleDocumentPolicyDocuments(ctx *models.Context, UniqueID string) (*models.RefDocumentPolicyDocuments, error) {
	documentpolicydocuments, err := s.Daos.GetSingleDocumentPolicyDocuments(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return documentpolicydocuments, nil
}

//UpdateDocumentPolicyDocuments : ""
func (s *Service) UpdateDocumentPolicyDocuments(ctx *models.Context, documentpolicydocuments *models.DocumentPolicyDocuments) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDocumentPolicyDocuments(ctx, documentpolicydocuments)
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

//EnableDocumentPolicyDocuments : ""
func (s *Service) EnableDocumentPolicyDocuments(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDocumentPolicyDocuments(ctx, UniqueID)
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

//DisableDocumentPolicyDocuments : ""
func (s *Service) DisableDocumentPolicyDocuments(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDocumentPolicyDocuments(ctx, UniqueID)
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

//DeleteDocumentPolicyDocuments : ""
func (s *Service) DeleteDocumentPolicyDocuments(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDocumentPolicyDocuments(ctx, UniqueID)
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

//FilterDocumentPolicyDocuments :""
func (s *Service) FilterDocumentPolicyDocuments(ctx *models.Context, documentpolicydocumentsFilter *models.FilterDocumentPolicyDocuments, pagination *models.Pagination) ([]models.RefDocumentPolicyDocuments, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDocumentPolicyDocuments(ctx, documentpolicydocumentsFilter, pagination)

}
