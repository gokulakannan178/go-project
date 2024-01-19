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

//SaveKnowlegdeDomain :""
func (d *Daos) SaveKnowlegdeDomain(ctx *models.Context, KnowlegdeDomain *models.KnowledgeDomain) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).InsertOne(ctx.CTX, KnowlegdeDomain)
	if err != nil {
		return err
	}
	KnowlegdeDomain.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleKnowlegdeDomain : ""
func (d *Daos) GetSingleKnowlegdeDomain(ctx *models.Context, code string) (*models.RefKnowledgeDomain, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var KnowlegdeDomains []models.RefKnowledgeDomain
	var KnowlegdeDomain *models.RefKnowledgeDomain
	if err = cursor.All(ctx.CTX, &KnowlegdeDomains); err != nil {
		return nil, err
	}
	if len(KnowlegdeDomains) > 0 {
		KnowlegdeDomain = &KnowlegdeDomains[0]
	}
	return KnowlegdeDomain, nil
}

//UpdateKnowlegdeDomain : ""
func (d *Daos) UpdateKnowlegdeDomain(ctx *models.Context, KnowlegdeDomain *models.KnowledgeDomain) error {
	selector := bson.M{"_id": KnowlegdeDomain.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": KnowlegdeDomain, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterKnowlegdeDomain : ""
func (d *Daos) FilterKnowledgeDomain(ctx *models.Context, KnowledgeDomainfilter *models.KnowledgeDomainFilter, pagination *models.Pagination) ([]models.RefKnowledgeDomain, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN, bson.M{
		"kdId": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.PROJECTKNOWLEDGEDOMAINSTATUSACTIVE}},
			{"$eq": []string{"$knowledgeDomain", "$$kdId"}},
		}}}},
	}, "ref.projects", "ref.projects")...)
	query := []bson.M{}
	if KnowledgeDomainfilter != nil {

		if len(KnowledgeDomainfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": KnowledgeDomainfilter.ActiveStatus}})
		}
		if len(KnowledgeDomainfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": KnowledgeDomainfilter.Status}})
		}
		//Regex
		if KnowledgeDomainfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: KnowledgeDomainfilter.Regex.Name, Options: "xi"}})
		}
		if KnowledgeDomainfilter.OmitProjectKD.Is {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$ne": KnowledgeDomainfilter.OmitProjectKD.Project}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if KnowledgeDomainfilter != nil {
		if KnowledgeDomainfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{KnowledgeDomainfilter.SortBy: KnowledgeDomainfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("KD pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("KnowlegdeDomain query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var KnowlegdeDomains []models.RefKnowledgeDomain
	if err = cursor.All(context.TODO(), &KnowlegdeDomains); err != nil {
		return nil, err
	}
	return KnowlegdeDomains, nil
}

//EnableKnowlegdeDomain :""
func (d *Daos) EnableKnowlegdeDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.KNOWLEDGEDOMAINSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableKnowlegdeDomain :""
func (d *Daos) DisableKnowlegdeDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.KNOWLEDGEDOMAINSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteKnowlegdeDomain :""
func (d *Daos) DeleteKnowlegdeDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.KNOWLEDGEDOMAINSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
