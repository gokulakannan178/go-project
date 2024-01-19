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

// SaveDashBoardPropertyDayWiseV2 : ""
func (d *Daos) SaveDashBoardPropertyDayWiseV2(ctx *models.Context, property *models.PropertyDashboardDayWise) error {
	d.Shared.BsonToJSONPrint(property)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).InsertOne(ctx.CTX, property)
	return err
}

// GetSingleDashBoardPropertyDayWiseV2  : ""
func (d *Daos) GetSingleDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) (*models.RefPropertyDashboardDayWise, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefPropertyDashboardDayWise
	var tower *models.RefPropertyDashboardDayWise
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateDashBoardPropertyDayWiseV2: ""
func (d *Daos) UpdateDashBoardPropertyDayWiseV2(ctx *models.Context, property *models.PropertyDashboardDayWise) error {
	selector := bson.M{"uniqueId": property.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": property}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableDashBoardPropertyDayWiseV2 : ""
func (d *Daos) EnableDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDPROPERTYDAYWISESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDashBoardPropertyDayWiseV2 : ""
func (d *Daos) DisableDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDPROPERTYDAYWISESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDashBoardPropertyDayWiseV2 : ""
func (d *Daos) DeleteDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDPROPERTYDAYWISESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardPropertyDayWiseV2 : ""
func (d *Daos) FilterDashBoardPropertyDayWiseV2(ctx *models.Context, filter *models.PropertyDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefPropertyDashboardDayWise, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTYDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var prop []models.RefPropertyDashboardDayWise
	if err = cursor.All(context.TODO(), &prop); err != nil {
		return nil, err
	}
	return prop, nil
}
