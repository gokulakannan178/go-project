package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveOffenceType :""
func (d *Daos) SaveOffenceType(ctx *models.Context, offenceType *models.OffenceType) error {
	_, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).InsertOne(ctx.CTX, offenceType)
	return err
}

//GetSingleOffenceType : ""
func (d *Daos) GetSingleOffenceType(ctx *models.Context, UniqueID string) (*models.RefOffenceType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offenceTypes []models.RefOffenceType
	var offenceType *models.RefOffenceType
	if err = cursor.All(ctx.CTX, &offenceTypes); err != nil {
		return nil, err
	}
	if len(offenceTypes) > 0 {
		offenceType = &offenceTypes[0]
	}
	return offenceType, nil
}

//UpdateOffenceType : ""
func (d *Daos) UpdateOffenceType(ctx *models.Context, offenceType *models.OffenceType) error {
	selector := bson.M{"uniqueId": offenceType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": offenceType, "$push": bson.M{"updatedLog": update}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOffenceType : ""
func (d *Daos) FilterOffenceType(ctx *models.Context, offenceTypefilter *models.OffenceTypeFilter, pagination *models.Pagination) ([]models.RefOffenceType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if offenceTypefilter != nil {
		if len(offenceTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": offenceTypefilter.Status}})
		}
		if len(offenceTypefilter.VehicleType) > 0 {
			query = append(query, bson.M{"vehicleType": bson.M{"$in": offenceTypefilter.VehicleType}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("offenceType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offenceTypes []models.RefOffenceType
	if err = cursor.All(context.TODO(), &offenceTypes); err != nil {
		return nil, err
	}
	return offenceTypes, nil
}

//EnableOffenceType :""
func (d *Daos) EnableOffenceType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFENCETYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOffenceType :""
func (d *Daos) DisableOffenceType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFENCETYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOffenceType :""
func (d *Daos) DeleteOffenceType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFENCETYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
