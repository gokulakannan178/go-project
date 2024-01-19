package daos

import (
	"hrms-services/constants"
	"hrms-services/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveUserLocation :""
func (d *Daos) SaveUserLocation(ctx *models.Context, userLocation *models.UserLocation) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATION).InsertOne(ctx.CTX, userLocation)
	if err != nil {
		return err
	}
	userLocation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
