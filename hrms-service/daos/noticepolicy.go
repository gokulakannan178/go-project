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

// SaveNoticePolicy : ""
func (d *Daos) SaveNoticePolicy(ctx *models.Context, noticepolicy *models.NoticePolicy) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).InsertOne(ctx.CTX, noticepolicy)
	if err != nil {
		return err
	}
	noticepolicy.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateNoticePolicy : ""
func (d *Daos) UpdateNoticePolicy(ctx *models.Context, noticepolicy *models.NoticePolicy) error {
	selector := bson.M{"uniqueId": noticepolicy.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": noticepolicy}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleNoticePolicy : ""
func (d *Daos) GetSingleNoticePolicy(ctx *models.Context, uniqueID string) (*models.RefNoticePolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var noticepolicys []models.RefNoticePolicy
	var noticepolicy *models.RefNoticePolicy
	if err = cursor.All(ctx.CTX, &noticepolicys); err != nil {
		return nil, err
	}
	if len(noticepolicys) > 0 {
		noticepolicy = &noticepolicys[0]
	}
	return noticepolicy, err
}

// EnableNoticePolicy : ""
func (d *Daos) EnableNoticePolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NOTICEPOLICYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableNoticePolicy : ""
func (d *Daos) DisableNoticePolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NOTICEPOLICYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteNoticePolicy :""
func (d *Daos) DeleteNoticePolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NOTICEPOLICYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterNoticePolicy : ""
func (d *Daos) FilterNoticePolicy(ctx *models.Context, noticepolicy *models.FilterNoticePolicy, pagination *models.Pagination) ([]models.RefNoticePolicy, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if noticepolicy != nil {
		if len(noticepolicy.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": noticepolicy.Status}})
		}
		if len(noticepolicy.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": noticepolicy.OrganisationId}})
		}
		//Regex
		if noticepolicy.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: noticepolicy.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if noticepolicy != nil {
		if noticepolicy.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{noticepolicy.SortBy: noticepolicy.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rf []models.RefNoticePolicy
	if err = cursor.All(context.TODO(), &rf); err != nil {
		return nil, err
	}
	return rf, nil
}
func (d *Daos) GetSingleNoticePolicyActiveWithName(ctx *models.Context, uniqueID string) (*models.RefNoticePolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.ONBOARDINGPOLICYSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("noticepolicy query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTICEPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var noticepolicys []models.RefNoticePolicy
	var noticepolicy *models.RefNoticePolicy
	if err = cursor.All(ctx.CTX, &noticepolicys); err != nil {
		return nil, err
	}
	if len(noticepolicys) > 0 {
		noticepolicy = &noticepolicys[0]
	}
	return noticepolicy, err
}
