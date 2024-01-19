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

//SaveAVR :""
func (d *Daos) SaveAVR(ctx *models.Context, avr *models.AVR) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONAVR).InsertOne(ctx.CTX, avr)
	return err
}

//GetSingleAVR : ""
func (d *Daos) GetSingleAVR(ctx *models.Context, UniqueID string) (*models.RefAVR, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityTypeId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "constructionTypeId", "uniqueId", "ref.constructionType", "ref.constructionType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "usageTypeId", "uniqueId", "ref.usageType", "ref.usageType")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAVR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var avrs []models.RefAVR
	var avr *models.RefAVR
	if err = cursor.All(ctx.CTX, &avrs); err != nil {
		return nil, err
	}
	if len(avrs) > 0 {
		avr = &avrs[0]
	}
	return avr, nil
}

//UpdateAVR : ""
func (d *Daos) UpdateAVR(ctx *models.Context, avr *models.AVR) error {
	selector := bson.M{"uniqueId": avr.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": avr, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVR).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterAVR : ""
func (d *Daos) FilterAVR(ctx *models.Context, avrfilter *models.AVRFilter, pagination *models.Pagination) ([]models.RefAVR, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if avrfilter != nil {

		if len(avrfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": avrfilter.Status}})
		}
		if len(avrfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": avrfilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAVR).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityTypeId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "constructionTypeId", "uniqueId", "ref.constructionType", "ref.constructionType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "usageTypeId", "uniqueId", "ref.usageType", "ref.usageType")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("avr query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAVR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var avrs []models.RefAVR
	if err = cursor.All(context.TODO(), &avrs); err != nil {
		return nil, err
	}
	return avrs, nil
}

//EnableAVR :""
func (d *Daos) EnableAVR(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AVRSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAVR :""
func (d *Daos) DisableAVR(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AVRSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAVR :""
func (d *Daos) DeleteAVR(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AVRSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
