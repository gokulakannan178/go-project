package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTaskMessage : ""
func (s *Service) SaveTaskMessage(ctx *models.Context, taskMessage *models.TaskMessage) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	taskMessage.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTASKMESSAGE)
	taskMessage.Status = constants.TASKMESSAGESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTaskMessage(ctx, taskMessage)
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

// FilterTaskMessage : ""
func (s *Service) FilterTaskMessage(ctx *models.Context, ftm *models.FilterTaskMessage, pagination *models.Pagination) (taskMessage []models.RefTaskMessage, err error) {
	return s.Daos.FilterTaskMessage(ctx, ftm, pagination)
}
