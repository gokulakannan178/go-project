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

//SaveApplicantType :""
func (d *Daos) SaveApplicantType(ctx *models.Context, ApplicantType *models.ApplicantType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).InsertOne(ctx.CTX, ApplicantType)
	return err
}

//GetSingleApplicantType : ""
func (d *Daos) GetSingleApplicantType(ctx *models.Context, UniqueID string) (*models.RefApplicantType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var applicantTypes []models.RefApplicantType
	var ApplicantType *models.RefApplicantType
	if err = cursor.All(ctx.CTX, &applicantTypes); err != nil {
		return nil, err
	}
	if len(applicantTypes) > 0 {
		ApplicantType = &applicantTypes[0]
	}
	return ApplicantType, nil
}

//UpdateApplicantType : ""
func (d *Daos) UpdateApplicantType(ctx *models.Context, ApplicantType *models.ApplicantType) error {
	selector := bson.M{"uniqueId": ApplicantType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ApplicantType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterApplicantType : ""
func (d *Daos) FilterApplicantType(ctx *models.Context, applicantTypefilter *models.ApplicantTypeFilter, pagination *models.Pagination) ([]models.RefApplicantType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if applicantTypefilter != nil {

		if len(applicantTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": applicantTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("ApplicantType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var applicantTypes []models.RefApplicantType
	if err = cursor.All(context.TODO(), &applicantTypes); err != nil {
		return nil, err
	}
	return applicantTypes, nil
}

//EnableApplicantType :""
func (d *Daos) EnableApplicantType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableApplicantType :""
func (d *Daos) DisableApplicantType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteApplicantType :""
func (d *Daos) DeleteApplicantType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
