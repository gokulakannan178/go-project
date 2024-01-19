package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveAmenities :""
func (d *Daos) SaveAmenities(ctx *models.Context, Amenities *models.Amenities) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).InsertOne(ctx.CTX, Amenities)
	return err
}

//GetSingleAmenities : ""
func (d *Daos) GetSingleAmenities(ctx *models.Context, UniqueID string) (*models.RefAmenities, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var amenitiess []models.RefAmenities
	var Amenities *models.RefAmenities
	if err = cursor.All(ctx.CTX, &amenitiess); err != nil {
		return nil, err
	}
	if len(amenitiess) > 0 {
		Amenities = &amenitiess[0]
	}
	return Amenities, nil
}

//UpdateAmenities : ""
func (d *Daos) UpdateAmenities(ctx *models.Context, Amenities *models.Amenities) error {
	selector := bson.M{"uniqueId": Amenities.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Amenities, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterAmenities : ""
func (d *Daos) FilterAmenities(ctx *models.Context, amenitiesfilter *models.AmenitiesFilter, pagination *models.Pagination) ([]models.RefAmenities, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if amenitiesfilter != nil {

		if len(amenitiesfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": amenitiesfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Amenities query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var amenitiess []models.RefAmenities
	if err = cursor.All(context.TODO(), &amenitiess); err != nil {
		return nil, err
	}
	return amenitiess, nil
}

//EnableAmenities :""
func (d *Daos) EnableAmenities(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AMENITIESSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAmenities :""
func (d *Daos) DisableAmenities(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AMENITIESSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAmenities :""
func (d *Daos) DeleteAmenities(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AMENITIESSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAMENITIES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
