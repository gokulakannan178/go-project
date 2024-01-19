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

// SaveLeaseDayWise : ""
func (d *Daos) SaveLeaseDayWise(ctx *models.Context, lease *models.LeaseDashboardDayWise) error {
	d.Shared.BsonToJSONPrint(lease)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).InsertOne(ctx.CTX, lease)
	return err
}

// GetSingleLeaseDayWise : ""
func (d *Daos) GetSingleLeaseDayWise(ctx *models.Context, UniqueID string) (*models.RefLeaseDashboardDayWise, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefLeaseDashboardDayWise
	var tower *models.RefLeaseDashboardDayWise
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateLeaseDayWise : ""
func (d *Daos) UpdateLeaseDayWise(ctx *models.Context, lease *models.LeaseDashboardDayWise) error {
	selector := bson.M{"uniqueId": lease.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": lease}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableLeaseDayWise : ""
func (d *Daos) EnableLeaseDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDLEASEDAYWISESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableLeaseDayWise : ""
func (d *Daos) DisableLeaseDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDLEASEDAYWISESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteLeaseDayWise : ""
func (d *Daos) DeleteLeaseDayWise(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDLEASEDAYWISESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLeaseDayWise : ""
func (d *Daos) FilterLeaseDayWise(ctx *models.Context, filter *models.LeaseDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefLeaseDashboardDayWise, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASEDAYWISE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refLease []models.RefLeaseDashboardDayWise
	if err = cursor.All(context.TODO(), &refLease); err != nil {
		return nil, err
	}
	return refLease, nil
}
