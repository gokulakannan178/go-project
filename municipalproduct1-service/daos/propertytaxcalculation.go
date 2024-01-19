package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//CalculationQuery : ""
func (d *Daos) CalculationQuery(propertyID string) []bson.M {
	var MainPipeLine []bson.M
	matchQuery := bson.M{"$match": bson.M{"uniqueId": propertyID}}
	var pipelineQuery []bson.M
	pipelineQuery = append(pipelineQuery, matchQuery)
	lookUpQuery1 := bson.M{"$lookup": bson.M{"from": constants.COLLECTIONPROPERTYCONFIGURATION, "as": "propertyConfig", "pipeline": pipelineQuery}}
	addFieldQuery1 := bson.M{"$addFields": bson.M{"propertyConfig": bson.M{"$arrayElemAt": []interface{}{"$propertyConfig", 0}}}}
	addFieldQuery2 := bson.M{"$addFields": bson.M{"percentAreaBuildup": bson.M{"$multiply": []interface{}{bson.M{"$divide": []string{"$buildUpArea", "$areaOfPlot"}}, 100}},
		"taxableVacantLand": bson.M{"$subtract": []interface{}{"$areaOfPlot", bson.M{"$multiply": []string{"$buildUpArea", "$propertyConfig.taxableVacantLandConfig"}}}}}}
	lookUpQuery2 := bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONFINANCIALYEAR,
		"as":   "fys",
		"let":  bson.M{"tempDOA": "$doa", "tempMT": "$municipalTypeId", "tempRT": "$roadTypeId", "taxablevland": "$taxableVacantLand", "percentAreaBuiltUp": "$percentAreaBuildup", "propertyConfig": "$propertyConfig"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []interface{}{
			bson.M{"$lt": []string{"$$tempDOA", "$to"}},
		}}}},
			{"$lookup": bson.M{
				"from": constants.COLLECTIONVACANTLANDRATE,
				"as":   "vlr",
				"let":  bson.M{"tempFrom": "$from", "tempTo": "$to", "mt": "$$tempMT", "rt": "$$tempRT", "doa": "$$tempDOA"},
				"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []interface{}{
					bson.M{"$eq": []string{"$municipalityTypeId", "$$mt"}},
					bson.M{"$eq": []string{"$roadTypeId", "$$rt"}},
					bson.M{"$lte": []string{"$doe", "$$tempTo"}},
					bson.M{"$gte": []string{"$doe", "$$doa"}},
				}}}},
					{"$sort": bson.M{"doe": -1}},
					{"$limit": 1},
				},
			}},
			{"$addFields": bson.M{"vlr": bson.M{"$arrayElemAt": []interface{}{"$vlr", 0}}}},
			{"$addFields": bson.M{"vacantLandTax": bson.M{"$cond": bson.M{"if": bson.M{"$lt": []string{"$$percentAreaBuiltUp", "$$propertyConfig.vacantLandRatePercentage"}}, "then": bson.M{"$multiply": []string{"$$taxablevland", "$vlr.rate"}}, "else": 0}}}},
		},
	},
	}
	MainPipeLine = append(MainPipeLine, matchQuery, lookUpQuery1, addFieldQuery1, addFieldQuery2, lookUpQuery2)
	return MainPipeLine
}

//GetPropertyTaxCalculation : ""
func (d *Daos) GetPropertyTaxCalculation(ctx *models.Context, UniqueID string) ([]models.PropertyTaxCalculation, error) {
	mainPipeline := d.CalculationQuery(UniqueID)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertytaxCalculation []models.PropertyTaxCalculation

	if err = cursor.All(ctx.CTX, &propertytaxCalculation); err != nil {
		return nil, err
	}
	return propertytaxCalculation, nil
}
