package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveScenario : ""
func (d *Daos) SaveScenario(ctx *models.Context, Scenario *models.Scenario) error {
	d.Shared.BsonToJSONPrint(Scenario)
	_, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).InsertOne(ctx.CTX, Scenario)
	return err
}

// GetSingleScenario : ""
func (d *Daos) GetSingleScenario(ctx *models.Context, UniqueID string) (*models.RefScenario, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("scenario query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Scenarios []models.RefScenario
	var Scenario *models.RefScenario
	if err = cursor.All(ctx.CTX, &Scenarios); err != nil {
		return nil, err
	}
	if len(Scenarios) > 0 {
		Scenario = &Scenarios[0]
	}
	return Scenario, nil
}

// UpdateScenario : ""
func (d *Daos) UpdateScenario(ctx *models.Context, Scenario *models.Scenario) error {
	selector := bson.M{"uniqueId": Scenario.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": Scenario}
	_, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableScenario : ""
func (d *Daos) EnableScenario(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SCENARIOSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableScenario : ""
func (d *Daos) DisableScenario(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SCENARIOSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteScenario : ""
func (d *Daos) DeleteScenario(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SCENARIOSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterScenario : ""
func (d *Daos) FilterScenario(ctx *models.Context, filter *models.ScenarioFilter, pagination *models.Pagination) ([]models.RefScenario, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		// //Regex Using searchBox Struct
		if filter.SearchText.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchText.Name, Options: "xi"}})
		}
		// if filter.SearchText.UniqueID != "" {
		// 	query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		// }
		// if filter.DateRange != nil {
		// 	//var sd,ed time.Time
		// 	if filter.DateRange.From != nil {
		// 		sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
		// 		ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
		// 		if filter.DateRange.To != nil {
		// 			ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
		// 		}
		// 		query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

		// 	}
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("scenario query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSCENARIO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Scenario []models.RefScenario
	if err = cursor.All(context.TODO(), &Scenario); err != nil {
		return nil, err
	}
	return Scenario, nil
}
