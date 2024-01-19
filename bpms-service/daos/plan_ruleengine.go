package daos

import (
	"bpms-service/constants"
	"bpms-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//GetSinglePlanRuleByScenario : ""
func (d *Daos) GetSinglePlanRuleByScenario(ctx *models.Context, scenario string) (*models.PlanRuleEngine, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"scenario": scenario}})
	d.Shared.BsonToJSONPrintTag("Plan query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANRULEENGINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.PlanRuleEngine
	var re *models.PlanRuleEngine
	if err = cursor.All(ctx.CTX, &res); err != nil {
		return nil, err
	}
	if len(res) > 0 {
		re = &res[0]
	}
	return re, nil
}

//GetSinglePlanRuleByScenarioAndFromStatus : ""
func (d *Daos) GetSinglePlanRuleByScenarioAndFromStatus(ctx *models.Context, scenario string, fromStatus string) (*models.PlanRuleEngine, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"scenario": scenario, "from": fromStatus}})
	d.Shared.BsonToJSONPrintTag("Plan query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANRULEENGINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.PlanRuleEngine
	var re *models.PlanRuleEngine
	if err = cursor.All(ctx.CTX, &res); err != nil {
		return nil, err
	}
	if len(res) > 0 {
		re = &res[0]
	}
	return re, nil
}
