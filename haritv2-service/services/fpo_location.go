package services

import (
	"errors"
	"haritv2-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

//FPOUpdateLocation : ""
func (s *Service) FPOUpdateLocation(ctx *models.Context, fpoloc *models.FPOUpdateLocation) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.FPOUpdateLocation(ctx, fpoloc)
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
