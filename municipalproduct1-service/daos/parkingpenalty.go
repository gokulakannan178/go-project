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

// SaveParkingPenalty : ""
func (d *Daos) SaveParkingPenalty(ctx *models.Context, parkingPenalty *models.ParkingPenalty) error {
	d.Shared.BsonToJSONPrint(parkingPenalty)
	_, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).InsertOne(ctx.CTX, parkingPenalty)
	return err
}

// GetSingleParkingPenalty : ""
func (d *Daos) GetSingleParkingPenalty(ctx *models.Context, UniqueID string) (*models.RefParkingPenalty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefParkingPenalty
	var tower *models.RefParkingPenalty
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateParkingPenalty : ""
func (d *Daos) UpdateParkingPenalty(ctx *models.Context, business *models.ParkingPenalty) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableParkingPenalty : ""
func (d *Daos) EnableParkingPenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PARKINGPENALTYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableParkingPenalty : ""
func (d *Daos) DisableParkingPenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PARKINGPENALTYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteParkingPenalty : ""
func (d *Daos) DeleteParkingPenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PARKINGPENALTYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterParkingPenalty : ""
func (d *Daos) FilterParkingPenalty(ctx *models.Context, filter *models.ParkingPenaltyFilter, pagination *models.Pagination) ([]models.RefParkingPenalty, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPARKINGPENALTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var parkingPenalty []models.RefParkingPenalty
	if err = cursor.All(context.TODO(), &parkingPenalty); err != nil {
		return nil, err
	}
	return parkingPenalty, nil
}
