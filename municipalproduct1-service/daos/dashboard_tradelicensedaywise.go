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

// SaveTradeLicenseDashboardDayWise : ""
func (d *Daos) SaveTradeLicenseDashboardDayWise(ctx *models.Context, tradeLicense *models.TradeLicenseDashboardDayWise) error {
	d.Shared.BsonToJSONPrint(tradeLicense)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).InsertOne(ctx.CTX, tradeLicense)
	return err
}

// GetSingleDashBoardProperty  : ""
func (d *Daos) GetSingleTradeLicenseDashboardDayWise(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseDashboardDayWise, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefTradeLicenseDashboardDayWise
	var tower *models.RefTradeLicenseDashboardDayWise
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateDashBoardProperty: ""
func (d *Daos) UpdateTradeLicenseDashboardDayWise(ctx *models.Context, tradeLicense *models.TradeLicenseDashboardDayWise) error {
	selector := bson.M{"uniqueId": tradeLicense.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": tradeLicense}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableDashBoardProperty : ""
func (d *Daos) EnableTradeLicenseDashboardDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDTRADELICENSEDAYWISESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDashBoardProperty : ""
func (d *Daos) DisableTradeLicenseDashboardDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDTRADELICENSEDAYWISESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDashBoardProperty : ""
func (d *Daos) DeleteTradeLicenseDashboardDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDTRADELICENSEDAYWISESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterTradeLicenseDashboardDayWise(ctx *models.Context, filter *models.TradeLicenseDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefTradeLicenseDashboardDayWise, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tradeLicense []models.RefTradeLicenseDashboardDayWise
	if err = cursor.All(context.TODO(), &tradeLicense); err != nil {
		return nil, err
	}
	return tradeLicense, nil
}
