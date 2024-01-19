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

//SaveEducationType :""
func (d *Daos) SaveEducationType(ctx *models.Context, EducationType *models.EducationType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).InsertOne(ctx.CTX, EducationType)
	return err
}

//GetSingleEducationType : ""
func (d *Daos) GetSingleEducationType(ctx *models.Context, UniqueID string) (*models.RefEducationType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var educationTypes []models.RefEducationType
	var EducationType *models.RefEducationType
	if err = cursor.All(ctx.CTX, &educationTypes); err != nil {
		return nil, err
	}
	if len(educationTypes) > 0 {
		EducationType = &educationTypes[0]
	}
	return EducationType, nil
}

//UpdateEducationType : ""
func (d *Daos) UpdateEducationType(ctx *models.Context, EducationType *models.EducationType) error {
	selector := bson.M{"uniqueId": EducationType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": EducationType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterEducationType : ""
func (d *Daos) FilterEducationType(ctx *models.Context, educationTypefilter *models.EducationTypeFilter, pagination *models.Pagination) ([]models.RefEducationType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if educationTypefilter != nil {

		if len(educationTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": educationTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("EducationType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var educationTypes []models.RefEducationType
	if err = cursor.All(context.TODO(), &educationTypes); err != nil {
		return nil, err
	}
	return educationTypes, nil
}

//EnableEducationType :""
func (d *Daos) EnableEducationType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EDUCATIONTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEducationType :""
func (d *Daos) DisableEducationType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EDUCATIONTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEducationType :""
func (d *Daos) DeleteEducationType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EDUCATIONTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEDUCATIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
