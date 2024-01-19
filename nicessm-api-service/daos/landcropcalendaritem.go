package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveLandCropCalendar :""
func (d *Daos) SaveLandCropCalendar(ctx *models.Context, landCropCalendar *models.LandCropCalendar) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).InsertOne(ctx.CTX, landCropCalendar)
	if err != nil {
		return err
	}
	landCropCalendar.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateLandCropCalendar : ""
func (d *Daos) UpdateLandCropCalendar(ctx *models.Context, landCropCalendar *models.LandCropCalendar) error {

	selector := bson.M{"_id": landCropCalendar.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": landCropCalendar}
	_, err := ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableLandCropCalendar :""
func (d *Daos) EnableLandCropCalendar(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LANDCROPCALENDARITEMSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLandCropCalendar :""
func (d *Daos) DisableLandCropCalendar(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LANDCROPCALENDARITEMSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLandCropCalendar :""
func (d *Daos) DeleteLandCropCalendar(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LANDCROPCALENDARITEMSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleLandCropCalendar : ""
func (d *Daos) GetSingleLandCropCalendar(ctx *models.Context, UniqueID string) (*models.RefLandCropCalendar, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "stage", "_id", "ref.stage", "ref.stage")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var landCropCalendars []models.RefLandCropCalendar
	var landCropCalendar *models.RefLandCropCalendar
	if err = cursor.All(ctx.CTX, &landCropCalendars); err != nil {
		return nil, err
	}
	if len(landCropCalendars) > 0 {
		landCropCalendar = &landCropCalendars[0]
	}
	return landCropCalendar, nil
}

//FilterLandCropCalendar : ""
func (d *Daos) FilterLandCropCalendar(ctx *models.Context, landCropCalendarfilter *models.LandCropCalendarFilter, pagination *models.Pagination) ([]models.RefLandCropCalendar, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if landCropCalendarfilter != nil {
		if len(landCropCalendarfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": landCropCalendarfilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "stage", "_id", "ref.stage", "ref.stage")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("LandCropCalendar query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANDCROPCALENDARITEM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var landCropCalendars []models.RefLandCropCalendar
	if err = cursor.All(context.TODO(), &landCropCalendars); err != nil {
		return nil, err
	}
	return landCropCalendars, nil
}
