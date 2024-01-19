package services

import (
	"errors"
	"haritv2-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

//UpdateUlbLocation : ""
func (s *Service) UpdateUlbLocation(ctx *models.Context, ulbloc *models.UpdateLocation) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUlbLocation(ctx, ulbloc)
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
