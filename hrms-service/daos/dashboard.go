package daos

import (
	"hrms-services/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) GetCollectionCount(ctx *models.Context, uniqueID string) (*models.Dashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$facet": bson.M{
			"total": []bson.M{
				bson.M{"$count": "total"}},

			"active": []bson.M{
				bson.M{"$match": bson.M{"status": "Active"}},
				bson.M{"$count": "active"}},
			"disabled": []bson.M{
				bson.M{"$match": bson.M{"status": "Disabled"}},
				bson.M{"$count": "disabled"}}},
	},

		bson.M{"$addFields": bson.M{"total": bson.M{"$arrayElemAt": []interface{}{"$total", 0}}}},
		bson.M{"$addFields": bson.M{"active": bson.M{"$arrayElemAt": []interface{}{"$active", 0}}}},
		bson.M{"$addFields": bson.M{"disabled": bson.M{"$arrayElemAt": []interface{}{"$disabled", 0}}}},
		bson.M{"$addFields": bson.M{"active": "$active.active", "disabled": "$disabled.disabled", "total": "$total.total"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(uniqueID).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Dashboards []models.Dashboard
	var Dashboard *models.Dashboard
	if err = cursor.All(ctx.CTX, &Dashboards); err != nil {
		return nil, err
	}
	if len(Dashboards) > 0 {
		Dashboard = &Dashboards[0]
	}
	return Dashboard, err
}
