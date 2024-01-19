package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) DashboardDemandAndCollection(ctx *models.Context, filter *models.DashboardDemandAndCollectionFilter) (*models.DashboardDemandAndCollection, error) {
	mainpipeline := []bson.M{}
	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, &filter.PropertyFilter)
	if len(query) > 0 {
		mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainpipeline = append(mainpipeline, bson.M{"$group": bson.M{"_id": nil,
		"totalDemandArrear":      bson.M{"$sum": "$demand.arrear"},
		"totalDemandCurrent":     bson.M{"$sum": "$demand.current"},
		"totalDemandTax":         bson.M{"$sum": "$demand.totalTax"},
		"totalCollectionArrear":  bson.M{"$sum": "$collection.arrear"},
		"totalCollectionCurrent": bson.M{"$sum": "$collection.current"},
		"totalCollectionTax":     bson.M{"$sum": "$collection.totalTax"},
	},
	})
	mainpipeline = append(mainpipeline)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property Query =>", mainpipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashboardDemandAndCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashboardDemandAndCollection{}, nil

}

func (d *Daos) DashboardDemandAndCollectionV2(ctx *models.Context, filter *models.DashboardDemandAndCollectionFilter) (*models.DashboardDemandAndCollection, error) {
	mainpipeline := []bson.M{}
	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, &filter.PropertyFilter)
	if len(query) > 0 {
		mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainpipeline = append(mainpipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.overalldemand", "ref.overalldemand")...)

	mainpipeline = append(mainpipeline, bson.M{"$group": bson.M{"_id": nil,

		"totalDemandCurrent":     bson.M{"$sum": "$ref.overalldemand.current.total"},
		"totalDemandArrear":      bson.M{"$sum": "$ref.overalldemand.arrear.total"},
		"totalDemandTax":         bson.M{"$sum": "$ref.overalldemand.total.totalTax"},
		"totalCollectionArrear":  bson.M{"$sum": "$collection.arrear"},
		"totalCollectionCurrent": bson.M{"$sum": "$collection.current"},
		"totalCollectionTax":     bson.M{"$sum": "$collection.totalTax"},
	},
	})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property Query =>", mainpipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashboardDemandAndCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashboardDemandAndCollection{}, nil

}
