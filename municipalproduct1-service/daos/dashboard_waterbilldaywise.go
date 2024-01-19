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

// SaveWaterBillDayWiseDashboard : ""
func (d *Daos) SaveWaterBillDayWiseDashboard(ctx *models.Context, waterBill *models.WaterBillDashboardDayWise) error {
	d.Shared.BsonToJSONPrint(waterBill)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).InsertOne(ctx.CTX, waterBill)
	return err
}

// GetSingleDashBoardProperty  : ""
func (d *Daos) GetSingleWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) (*models.RefWaterBillDashboardDayWise, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefWaterBillDashboardDayWise
	var tower *models.RefWaterBillDashboardDayWise
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateDashBoardProperty: ""
func (d *Daos) UpdateWaterBillDayWiseDashboard(ctx *models.Context, waterBill *models.WaterBillDashboardDayWise) error {
	selector := bson.M{"uniqueId": waterBill.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": waterBill}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableDashBoardProperty : ""
func (d *Daos) EnableWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDWATERBILLDAYWISESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDashBoardProperty : ""
func (d *Daos) DisableWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDWATERBILLDAYWISESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDashBoardProperty : ""
func (d *Daos) DeleteWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDWATERBILLDAYWISESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterWaterBillDayWiseDashboard(ctx *models.Context, filter *models.WaterBillDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefWaterBillDashboardDayWise, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILLDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var waterBill []models.RefWaterBillDashboardDayWise
	if err = cursor.All(context.TODO(), &waterBill); err != nil {
		return nil, err
	}
	return waterBill, nil
}
