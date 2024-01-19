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

//SaveResidentialType :""
func (d *Daos) SaveResidentialType(ctx *models.Context, residentialType *models.ResidentialType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).InsertOne(ctx.CTX, residentialType)
	return err
}

//GetSingleResidentialType : ""
func (d *Daos) GetSingleResidentialType(ctx *models.Context, UniqueID string) (*models.RefResidentialType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var residentialTypes []models.RefResidentialType
	var residentialType *models.RefResidentialType
	if err = cursor.All(ctx.CTX, &residentialTypes); err != nil {
		return nil, err
	}
	if len(residentialTypes) > 0 {
		residentialType = &residentialTypes[0]
	}
	return residentialType, nil
}

//UpdateResidentialType : ""
func (d *Daos) UpdateResidentialType(ctx *models.Context, residentialType *models.ResidentialType) error {
	selector := bson.M{"uniqueId": residentialType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": residentialType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterResidentialType : ""
func (d *Daos) FilterResidentialType(ctx *models.Context, residentialTypefilter *models.ResidentialTypeFilter, pagination *models.Pagination) ([]models.RefResidentialType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if residentialTypefilter != nil {

		if len(residentialTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": residentialTypefilter.Status}})
		}
		if len(residentialTypefilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": residentialTypefilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("residentialType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var residentialTypes []models.RefResidentialType
	if err = cursor.All(context.TODO(), &residentialTypes); err != nil {
		return nil, err
	}
	return residentialTypes, nil
}

//EnableResidentialType :""
func (d *Daos) EnableResidentialType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.RESIDENTIALTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableResidentialType :""
func (d *Daos) DisableResidentialType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.RESIDENTIALTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteResidentialType :""
func (d *Daos) DeleteResidentialType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.RESIDENTIALTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRESIDENTIALTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
