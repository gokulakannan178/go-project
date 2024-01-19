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

// SaveUserChargeDayWiseDashboard : ""
func (d *Daos) SaveUserChargeDayWiseDashboard(ctx *models.Context, userCharge *models.UserChargeDashboardDayWise) error {
	d.Shared.BsonToJSONPrint(userCharge)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).InsertOne(ctx.CTX, userCharge)
	return err
}

// GetSingleDashBoardProperty  : ""
func (d *Daos) GetSingleUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) (*models.RefUserChargeDashboardDayWise, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserChargeDashboardDayWise
	var tower *models.RefUserChargeDashboardDayWise
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserChargeDayWiseDashboard: ""
func (d *Daos) UpdateUserChargeDayWiseDashboard(ctx *models.Context, userCharge *models.UserChargeDashboardDayWise) error {
	selector := bson.M{"uniqueId": userCharge.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": userCharge}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserChargeDayWiseDashboard : ""
func (d *Daos) EnableUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDUSERCHARGEDAYWISESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserChargeDayWiseDashboard : ""
func (d *Daos) DisableUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDUSERCHARGEDAYWISESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserChargeDayWiseDashboard : ""
func (d *Daos) DeleteUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDUSERCHARGEDAYWISESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserChargeDayWiseDashboard : ""
func (d *Daos) FilterUserChargeDayWiseDashboard(ctx *models.Context, filter *models.UserChargeDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefUserChargeDashboardDayWise, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var userCharge []models.RefUserChargeDashboardDayWise
	if err = cursor.All(context.TODO(), &userCharge); err != nil {
		return nil, err
	}
	return userCharge, nil
}
