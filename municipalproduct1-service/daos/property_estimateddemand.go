package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
)

// SaveEstimatedPropertyDemand :""
func (d *Daos) SaveEstimatedPropertyDemand(ctx *models.Context, property *models.Property, collectionName string) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONESTIMATEDPROPERTYDEMAND).InsertOne(ctx.SC, property)
	return err
}
