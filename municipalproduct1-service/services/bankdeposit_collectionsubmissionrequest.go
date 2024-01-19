package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//CollectionSubmissionRequest :""
func (s *Service) CollectionSubmissionRequest(ctx *models.Context, csr *models.CollectionSubmissionRequest) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	csr.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBANKDEPOSITCOLLECTIONSUBMISSIONREQUEST)
	csr.Status = constants.BANKDEPOSITCOLLECTIONSUBMISSIONREQUESTSTATUSACTIVE
	t := time.Now()
	csr.Actioner.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.CollectionSubmissionRequest(ctx, csr)
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

//CollectionSubmissionRequestFilter
func (s *Service) CollectionSubmissionRequestFilter(ctx *models.Context, csr *models.CollectionSubmissionRequestFilter, pagination *models.Pagination) ([]models.RefCollectionSubmissionRequest, error) {
	return s.Daos.CollectionSubmissionRequestFilter(ctx, csr, pagination)
}
