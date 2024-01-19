package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SavePmAchievement : ""
func (d *Daos) SavePmAchievement(ctx *models.Context, pmAchievement *models.PmAchievement) error {
	d.Shared.BsonToJSONPrint(pmAchievement)
	_, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).InsertOne(ctx.CTX, pmAchievement)
	return err
}

// GetSinglePmAchievement : ""
func (d *Daos) GetSinglePmAchievement(ctx *models.Context, UniqueID string) (*models.RefPmAchievement, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefPmAchievement
	var tower *models.RefPmAchievement
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdatePmAchievement : ""
func (d *Daos) UpdatePmAchievement(ctx *models.Context, business *models.PmAchievement) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}

	_, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePmAchievement : ""
func (d *Daos) EnablePmAchievement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PMACHIEVEMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePmAchievement : ""
func (d *Daos) DisablePmAchievement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PMACHIEVEMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePmAchievement : ""
func (d *Daos) DeletePmAchievement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PMACHIEVEMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPmAchievement : ""
func (d *Daos) FilterPmAchievement(ctx *models.Context, filter *models.PmAchievementFilter, pagination *models.Pagination) ([]models.RefPmAchievement, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPMACHIEVEMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pmAchievement []models.RefPmAchievement
	if err = cursor.All(context.TODO(), &pmAchievement); err != nil {
		return nil, err
	}
	return pmAchievement, nil
}

func PmAchievementMonthWise(ctx *models.Context, filter *models.PmAchievementMonthWiseFilter) (*models.PmAchievementMonthWise, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if filter.FYID != "" {
			query = append(query, bson.M{"uniqueId": filter.FYID})
		}
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "months",
		"as":   "months",
		"let":  bson.M{"fy": "$$ROOT"},
		"pipeline": []bson.M{
			bson.M{"$sort": bson.M{"fyOrder": 1}},
			bson.M{"$addFields": bson.M{
				"startDate": bson.M{
					"$dateFromParts": bson.M{
						"year":  bson.M{"$cond": bson.M{"if": bson.M{"$in": []interface{}{"$month", []int{1, 2, 3}}}, "then": bson.M{"$year": "$$fy.to"}, "else": bson.M{"$year": "$$fy.from"}}},
						"month": "$month", "day": 1,
					},
				},

				"endDate": bson.M{
					"$dateFromParts": bson.M{
						"year":  bson.M{"$cond": bson.M{"if": bson.M{"$in": []interface{}{"$month", []int{1, 2, 3}}}, "then": bson.M{"$year": "$$fy.to"}, "else": bson.M{"$year": "$$fy.from"}}},
						"month": bson.M{"$sum": []interface{}{"$month", 1}}, "day": 1,
					},
				},
			},
			},
			bson.M{"$addFields": bson.M{
				"endDate": bson.M{"$subtract": []interface{}{"$endDate", 1}},
			}},
			bson.M{"$lookup": bson.M{
				"from": "pmachievement",
				"as":   "achivement",
				"let":  bson.M{"fyId": "$$fy.uniqueId", "month": "$month"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []interface{}{"$$fyId", "$fyId"}},
						bson.M{"$eq": []interface{}{"$$month", "$month"}},
						bson.M{"$eq": []interface{}{"$pmId", "1"}},
					}}}},
				},
			}},
			bson.M{
				"$addFields": bson.M{"pmachievement": bson.M{"$arrayElemAt": []interface{}{"$pmachievement", 0}}},
			},
		},
	}})

	return nil, nil
}
