package services

import (
	"errors"
	"fmt"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDocumentPolicy :""
func (s *Service) SaveDocumentPolicy(ctx *models.Context, documentPolicy *models.DocumentPolicy) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	documentPolicy.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDOCUMENTPOLICY)
	documentPolicy.Status = constants.DOCUMENTPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DocumentPolicy.created")
	documentPolicy.Created = created
	log.Println("b4 DocumentPolicy.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDocumentPolicy(ctx, documentPolicy)
		if dberr != nil {
			return dberr
		}

		// refDocumentPolicy, err := s.Daos.GetSingleDocumentPolicy(ctx, documentPolicy.UniqueID)
		// if err != nil {
		// 	return err
		// }
		// fmt.Println("refDocumentPolicy", refDocumentPolicy.DocumentMasterId)

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		for _, v := range documentPolicy.DocumentMasterId {
			documentPolicyDocuments := new(models.DocumentPolicyDocuments)
			documentmaster, err := s.Daos.GetSingleDocumentMasterWithActive(ctx, v, constants.DOCUMENTMASTERSTATUSACTIVE)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if documentmaster != nil {
				documentPolicyDocuments.DocumentPolicyID = documentPolicy.UniqueID

				documentPolicyDocuments.DocumentMasterID = v
				documentPolicyDocuments.Name = documentPolicy.Name
			} else {
				return errors.New("documentmaster not founded")
			}
			err = s.SaveDocumentPolicyDocumentsWithoutTransaction(ctx, documentPolicyDocuments)
			if err != nil {
				return err
			}
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

//GetSingleDocumentPolicy :""
func (s *Service) GetSingleDocumentPolicy(ctx *models.Context, UniqueID string) (*models.RefDocumentPolicy, error) {
	documentpolicy, err := s.Daos.GetSingleDocumentPolicy(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return documentpolicy, nil
}

//UpdateDocumentPolicy : ""
func (s *Service) UpdateDocumentPolicy(ctx *models.Context, documentPolicy *models.DocumentPolicy) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.DocumentPolicyDocumentsRemoveNotPresentValue(ctx, documentPolicy.UniqueID, documentPolicy.DocumentMasterId)
		if err != nil {
			return err
		}
		err = s.Daos.DocumentPolicyDocumentsUpsert(ctx, documentPolicy.UniqueID, documentPolicy.DocumentMasterId, documentPolicy.Name)
		if err != nil {
			return err
		}

		fmt.Println("error==>", err)

		err = s.Daos.UpdateDocumentPolicy(ctx, documentPolicy)
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

//EnableDocumentPolicy : ""
func (s *Service) EnableDocumentPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDocumentPolicy(ctx, UniqueID)
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

//DisableDocumentPolicy : ""
func (s *Service) DisableDocumentPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDocumentPolicy(ctx, UniqueID)
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

//DeleteDocumentPolicy : ""
func (s *Service) DeleteDocumentPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDocumentPolicy(ctx, UniqueID)
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

//FilterDocumentPolicy :""
func (s *Service) FilterDocumentPolicy(ctx *models.Context, documentpolicyFilter *models.FilterDocumentPolicy, pagination *models.Pagination) ([]models.RefDocumentPolicy, error) {
	err := s.DocumentPolicyDataAccess(ctx, documentpolicyFilter)
	if err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDocumentPolicy(ctx, documentpolicyFilter, pagination)

}
func (s *Service) DocumentPolicyDataAccess(ctx *models.Context, filter *models.FilterDocumentPolicy) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationID = append(filter.OrganisationID, v.UniqueID)
				}
			}

		}

	}
	return err
}
