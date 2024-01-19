package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveProbationary : ""
func (d *Daos) SaveProbationary(ctx *models.Context, probationary *models.Probationary) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).InsertOne(ctx.CTX, probationary)
	if err != nil {
		return err
	}
	probationary.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateProbationary : ""
func (d *Daos) UpdateProbationary(ctx *models.Context, probationary *models.Probationary) error {
	selector := bson.M{"uniqueId": probationary.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": probationary}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleProbationary : ""
func (d *Daos) GetSingleProbationary(ctx *models.Context, uniqueID string) (*models.RefProbationary, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var probationarys []models.RefProbationary
	var probationary *models.RefProbationary
	if err = cursor.All(ctx.CTX, &probationarys); err != nil {
		return nil, err
	}
	if len(probationarys) > 0 {
		probationary = &probationarys[0]
	}
	return probationary, err
}

// EnableProbationary : ""
func (d *Daos) EnableProbationary(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PROBATIONARYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableProbationary : ""
func (d *Daos) DisableProbationary(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PROBATIONARYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteProbationary :""
func (d *Daos) DeleteProbationary(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROBATIONARYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterProbationary : ""
func (d *Daos) FilterProbationary(ctx *models.Context, probationary *models.FilterProbationary, pagination *models.Pagination) ([]models.RefProbationary, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if probationary != nil {
		if len(probationary.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": probationary.Status}})
		}
		if len(probationary.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": probationary.OrganisationId}})
		}
		//Regex
		if probationary.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: probationary.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if probationary != nil {
		if probationary.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{probationary.SortBy: probationary.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var probationaryFilter []models.RefProbationary
	if err = cursor.All(context.TODO(), &probationaryFilter); err != nil {
		return nil, err
	}
	return probationaryFilter, nil
}

// GetSingleProbationaryWithActiveName : ""
func (d *Daos) GetSingleProbationaryWithActiveName(ctx *models.Context, uniqueID string) (*models.RefProbationary, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.ONBOARDINGPOLICYSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("Probationary query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROBATIONARY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var probationarys []models.RefProbationary
	var probationary *models.RefProbationary
	if err = cursor.All(ctx.CTX, &probationarys); err != nil {
		return nil, err
	}
	if len(probationarys) > 0 {
		probationary = &probationarys[0]
	}
	return probationary, err
}
