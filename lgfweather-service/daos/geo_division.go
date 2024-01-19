package daos

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveDivision :""
func (d *Daos) SaveDivision(ctx *models.Context, division *models.Division) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).InsertOne(ctx.CTX, division)
	return err
}

//GetSingleDivision : ""
func (d *Daos) GetSingleDivision(ctx *models.Context, uniqueID string) (*models.RefDivision, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "code", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Divisions []models.RefDivision
	var Division *models.RefDivision
	if err = cursor.All(ctx.CTX, &Divisions); err != nil {
		return nil, err
	}
	if len(Divisions) > 0 {
		Division = &Divisions[0]
	}
	return Division, nil
}

//UpdateDivision : ""
func (d *Daos) UpdateDivision(ctx *models.Context, division *models.Division) error {
	selector := bson.M{"uniqueId": division.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": division, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDivision : ""
func (d *Daos) FilterDivision(ctx *models.Context, divisionfilter *models.DivisionFilter, pagination *models.Pagination) ([]models.RefDivision, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if divisionfilter != nil {
		if len(divisionfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": divisionfilter.UniqueID}})
		}
		if len(divisionfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": divisionfilter.Status}})
		}
		if len(divisionfilter.StateCode) > 0 {
			query = append(query, bson.M{"stateCode": bson.M{"$in": divisionfilter.StateCode}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "code", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Division query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Divisions []models.RefDivision
	if err = cursor.All(context.TODO(), &Divisions); err != nil {
		return nil, err
	}
	return Divisions, nil
}

//EnableDivision :""
func (d *Daos) EnableDivision(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DIVISIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDivision :""
func (d *Daos) DisableDivision(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DIVISIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDivision :""
func (d *Daos) DeleteDivision(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DIVISIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetActiveDivision(ctx *models.Context) ([]models.Division, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": constants.DIVISIONSTATUSACTIVE})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("activestate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDIVISION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var divisions []models.Division
	if err = cursor.All(context.TODO(), &divisions); err != nil {
		return nil, err
	}
	return divisions, nil
}
