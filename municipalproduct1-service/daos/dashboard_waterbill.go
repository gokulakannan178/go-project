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

// SaveWaterBillDashboard : ""
func (d *Daos) SaveWaterBillDashboard(ctx *models.Context, waterBill *models.WaterBillDashboard) error {
	d.Shared.BsonToJSONPrint(waterBill)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).InsertOne(ctx.CTX, waterBill)
	return err
}

// GetSingleDashBoardProperty  : ""
func (d *Daos) GetSingleWaterBillDashboard(ctx *models.Context, UniqueID string) (*models.RefWaterBillDashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefWaterBillDashboard
	var tower *models.RefWaterBillDashboard
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateDashBoardProperty: ""
func (d *Daos) UpdateWaterBillDashboard(ctx *models.Context, waterBill *models.WaterBillDashboard) error {
	selector := bson.M{"uniqueId": waterBill.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": waterBill}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableDashBoardProperty : ""
func (d *Daos) EnableWaterBillDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDWATERBILLSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDashBoardProperty : ""
func (d *Daos) DisableWaterBillDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDWATERBILLSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDashBoardProperty : ""
func (d *Daos) DeleteWaterBillDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDWATERBILLSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterWaterBillDashboard(ctx *models.Context, filter *models.WaterBillDashboardFilter, pagination *models.Pagination) ([]models.RefWaterBillDashboard, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDWATERBILL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var waterBill []models.RefWaterBillDashboard
	if err = cursor.All(context.TODO(), &waterBill); err != nil {
		return nil, err
	}
	return waterBill, nil
}
