package daos

import (
	"bpms-service/constants"
	"bpms-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//GetSinglePlanCRFRuleByScenario : ""
func (d *Daos) GetSinglePlanCRFRuleByScenario(ctx *models.Context, scenario string) (*models.PlanCRFRuleEngine, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"scenario": scenario}})
	d.Shared.BsonToJSONPrintTag("Plan query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANCRFRULEENGINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.PlanCRFRuleEngine
	var re *models.PlanCRFRuleEngine
	if err = cursor.All(ctx.CTX, &res); err != nil {
		return nil, err
	}
	if len(res) > 0 {
		re = &res[0]
	}
	return re, nil
}
