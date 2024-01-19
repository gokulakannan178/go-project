package daos

import (
	"context"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetShopRentForOverAllDemand : ""
func (d *Daos) GetShopRentForOverAllDemand(ctx *models.Context, status []string) ([]models.ShopRentMin, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$in": status}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"uniqueId": 1}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetShopRentForOverAllDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shoprentMin []models.ShopRentMin
	if err = cursor.All(context.TODO(), &shoprentMin); err != nil {
		return nil, err
	}
	return shoprentMin, nil
}
