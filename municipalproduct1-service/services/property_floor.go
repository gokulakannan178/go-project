package services

import (
	"errors"
	"log"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateFloorPropertyID : ""
func (s *Service) UpdateFloorPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
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
			err = s.Daos.UpdateFloorPropertyID(ctx, uniqueIds)
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
