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

// SaveMobileTowerDayWise : ""
func (d *Daos) SaveMobileTowerDayWise(ctx *models.Context, mobileTower *models.MobiletowerDashboardDayWise) error {
	d.Shared.BsonToJSONPrint(mobileTower)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).InsertOne(ctx.CTX, mobileTower)
	return err
}

// GetSingleMobileTowerDayWise : ""
func (d *Daos) GetSingleMobileTowerDayWise(ctx *models.Context, UniqueID string) (*models.RefMobileTowerDayWiseDashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefMobileTowerDayWiseDashboard
	var tower *models.RefMobileTowerDayWiseDashboard
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateMobileTowerDayWise : ""
func (d *Daos) UpdateMobileTowerDayWise(ctx *models.Context, mobileTower *models.MobiletowerDashboardDayWise) error {
	selector := bson.M{"uniqueId": mobileTower.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": mobileTower}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMobileTowerDayWise : ""
func (d *Daos) EnableMobileTowerDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDMOBILETOWERDAYWISESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableMobileTowerDayWise : ""
func (d *Daos) DisableMobileTowerDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDMOBILETOWERDAYWISESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteMobileTowerDayWise : ""
func (d *Daos) DeleteMobileTowerDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDMOBILETOWERDAYWISESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMobileTowerDayWise : ""
func (d *Daos) FilterMobileTowerDayWise(ctx *models.Context, filter *models.MobileTowerDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefMobileTowerDayWiseDashboard, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refMobileTower []models.RefMobileTowerDayWiseDashboard
	if err = cursor.All(context.TODO(), &refMobileTower); err != nil {
		return nil, err
	}
	return refMobileTower, nil
}
