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

// SaveMonthlyTarget : ""
func (d *Daos) SaveMonthlyTarget(ctx *models.Context, block *models.MonthlyTarget) error {
	d.Shared.BsonToJSONPrint(block)
	_, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).InsertOne(ctx.CTX, block)
	return err
}

// GetSingleMonthlyTarget  : ""
func (d *Daos) GetSingleMonthlyTarget(ctx *models.Context, UniqueID string) (*models.RefMonthlyTarget, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefMonthlyTarget
	var tower *models.RefMonthlyTarget
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateMonthlyTarget: ""
func (d *Daos) UpdateMonthlyTarget(ctx *models.Context, crop *models.MonthlyTarget) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMonthlyTarget : ""
func (d *Daos) EnableMonthlyTarget(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MONTHLYTARGETSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableMonthlyTarget : ""
func (d *Daos) DisableMonthlyTarget(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MONTHLYTARGETSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteMonthlyTarget : ""
func (d *Daos) DeleteMonthlyTarget(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MONTHLYTARGETSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterMonthlyTarget(ctx *models.Context, filter *models.MonthlyTargetFilter, pagination *models.Pagination) ([]models.RefMonthlyTarget, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMONTHLYTARGET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var prop []models.RefMonthlyTarget
	if err = cursor.All(context.TODO(), &prop); err != nil {
		return nil, err
	}
	return prop, nil
}
