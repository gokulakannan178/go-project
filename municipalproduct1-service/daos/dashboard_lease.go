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

// SaveLease : ""
func (d *Daos) SaveLease(ctx *models.Context, lease *models.LeaseDashboard) error {
	d.Shared.BsonToJSONPrint(lease)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).InsertOne(ctx.CTX, lease)
	return err
}

// GetSingleLease : ""
func (d *Daos) GetSingleLease(ctx *models.Context, UniqueID string) (*models.RefLeaseDashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefLeaseDashboard
	var tower *models.RefLeaseDashboard
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateLease : ""
func (d *Daos) UpdateLease(ctx *models.Context, lease *models.LeaseDashboard) error {
	selector := bson.M{"uniqueId": lease.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": lease}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableLease : ""
func (d *Daos) EnableLease(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDLEASESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableLease : ""
func (d *Daos) DisableLease(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDLEASESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteLease : ""
func (d *Daos) DeleteLease(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDLEASESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLease : ""
func (d *Daos) FilterLease(ctx *models.Context, filter *models.LeaseDashboardFilter, pagination *models.Pagination) ([]models.RefLeaseDashboard, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDLEASE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refLease []models.RefLeaseDashboard
	if err = cursor.All(context.TODO(), &refLease); err != nil {
		return nil, err
	}
	return refLease, nil
}
