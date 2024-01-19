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

//SaveFarmerCropCalendar :""
func (d *Daos) SaveFarmerCropCalendar(ctx *models.Context, farmerCropCalendar *models.FarmerCropCalendar) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).InsertOne(ctx.CTX, farmerCropCalendar)
	if err != nil {
		return err
	}
	farmerCropCalendar.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateFarmerCropCalendar : ""
func (d *Daos) UpdateFarmerCropCalendar(ctx *models.Context, farmerCropCalendar *models.FarmerCropCalendar) error {

	selector := bson.M{"_id": farmerCropCalendar.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": farmerCropCalendar}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableFarmerCropCalendar :""
func (d *Daos) EnableFarmerCropCalendar(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCROPCALENDARSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmerCropCalendar :""
func (d *Daos) DisableFarmerCropCalendar(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCROPCALENDARSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmerCropCalendar :""
func (d *Daos) DeleteFarmerCropCalendar(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCROPCALENDARSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleFarmerCropCalendar : ""
func (d *Daos) GetSingleFarmerCropCalendar(ctx *models.Context, UniqueID string) (*models.RefFarmerCropCalendar, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "crop", "_id", "ref.crop", "ref.crop")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "interCrop", "_id", "ref.interCrop", "ref.interCrop")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "season", "_id", "ref.season", "ref.season")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "veriety", "_id", "ref.veriety", "ref.veriety")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmerCropCalendars []models.RefFarmerCropCalendar
	var farmerCropCalendar *models.RefFarmerCropCalendar
	if err = cursor.All(ctx.CTX, &farmerCropCalendars); err != nil {
		return nil, err
	}
	if len(farmerCropCalendars) > 0 {
		farmerCropCalendar = &farmerCropCalendars[0]
	}
	return farmerCropCalendar, nil
}

//FilterFarmerCropCalendar : ""
func (d *Daos) FilterFarmerCropCalendar(ctx *models.Context, FarmerCropCalendarfilter *models.FarmerCropCalendarFilter, pagination *models.Pagination) ([]models.RefFarmerCropCalendar, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if FarmerCropCalendarfilter != nil {
		if len(FarmerCropCalendarfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": FarmerCropCalendarfilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "crop", "_id", "ref.crop", "ref.crop")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "season", "_id", "ref.season", "ref.season")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "veriety", "_id", "ref.veriety", "ref.veriety")...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("FarmerCropCalendar query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROPCALENDAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmerCropCalendars []models.RefFarmerCropCalendar
	if err = cursor.All(context.TODO(), &farmerCropCalendars); err != nil {
		return nil, err
	}
	return farmerCropCalendars, nil
}
