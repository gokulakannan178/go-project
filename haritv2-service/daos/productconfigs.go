package daos

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) SaveProductConfig(ctx *models.Context, user *models.ProductConfig) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).InsertOne(ctx.CTX, user)
	return err
}

func (d *Daos) EnableProductConfig(ctx *models.Context, UniqueID string) error {
	filter := bson.M{
		"status": bson.M{
			"$eq": "Active",
		},
	}
	updatemany := bson.M{"$set": bson.M{"status": "Disable"}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateMany(ctx.CTX, filter, updatemany)
	query := bson.M{"UniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": "Active"}}
	_, id := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateOne(ctx.CTX, query, update)

	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return id
}
func (d *Daos) GetactiveProductConfig(ctx *models.Context, Status string) (*models.ProductConfig, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": Status}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCTCONFIG, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var products []models.ProductConfig
	var product *models.ProductConfig
	if err = cursor.All(ctx.CTX, &products); err != nil {
		return nil, err
	}
	if len(products) > 0 {
		product = &products[0]
	}
	return product, nil
}
