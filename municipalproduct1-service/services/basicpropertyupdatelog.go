package services

import (
	"errors"
	"log"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) GetSingleBasicPropertyUpdateLog(ctx *models.Context, UniqueID string) (*models.RefBasicPropertyUpdateLog, error) {
	property, err := s.Daos.GetSingleBasicPropertyUpdateLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return property, nil
}

func (s *Service) BasicPropertyUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbpul *models.RefBasicPropertyUpdateLog) ([]models.RefPropertyPayment, error) {
	return s.Daos.BasicPropertyUpdateGetPaymentsToBeUpdated(ctx, rbpul)
}

// UpdateBasicPropertyUpdateLogPropertyID : ""
func (s *Service) UpdateBasicPropertyUpdateLogPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range uniqueIds.UniqueIDs {
			resProperty, err := s.GetSingleProperty(ctx, v)
			if err != nil {
				return errors.New("Not able to get property - " + err.Error())
			}

			uniqueIds.UniqueID = resProperty.OldUniqueID
			uniqueIds.OldUniqueID = resProperty.OldUniqueID
			uniqueIds.NewUniqueID = resProperty.NewUniqueID
			err = s.Daos.UpdateBasicPropertyUpdateLogPropertyID(ctx, uniqueIds)
			if err != nil {
				return err
			}
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
