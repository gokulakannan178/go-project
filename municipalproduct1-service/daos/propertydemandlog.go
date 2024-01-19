package daos

import (
	"errors"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePropertyDemandLog :""
func (d *Daos) SavePropertyDemandLog(ctx *models.Context, propertyDemandLog *models.PropertyDemandLog) error {
	opts := options.Update().SetUpsert(true)
	query := bson.M{"propertyId": propertyDemandLog.PropertyID}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDEMANDLOG).UpdateOne(ctx.CTX, query, bson.M{"$set": propertyDemandLog}, opts)
	return err
}

//GetSinglePropertyDemandLog
func (d *Daos) GetSinglePropertyDemandLog(ctx *models.Context, propertyID string) (*models.PropertyDemandLog, error) {
	propertyDemandLog := new(models.PropertyDemandLog)
	query := bson.M{"propertyId": propertyID}
	d.Shared.BsonToJSONPrint(query)
	err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDEMANDLOG).FindOne(ctx.CTX, query).Decode(&propertyDemandLog)
	if err != nil {
		return nil, errors.New("GetSinglePropertyDemandLog find one - " + err.Error())
	}
	return propertyDemandLog, nil
}

// UpdatePropertyDemandLogPropertyID :""
func (d *Daos) UpdatePropertyDemandLogPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDEMANDLOG).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
