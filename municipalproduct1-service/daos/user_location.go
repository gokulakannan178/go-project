package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
)

//SaveUserLocation :""
func (d *Daos) SaveUserLocation(ctx *models.Context, userLocation *models.UserLocation) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATION).InsertOne(ctx.CTX, userLocation)
	return err
}
