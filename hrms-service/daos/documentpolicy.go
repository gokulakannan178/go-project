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

//SaveDocumentPolicy :""
func (d *Daos) SaveDocumentPolicy(ctx *models.Context, documentPolicy *models.DocumentPolicy) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).InsertOne(ctx.CTX, documentPolicy)
	if err != nil {
		return err
	}
	documentPolicy.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDocumentPolicy : ""
func (d *Daos) GetSingleDocumentPolicy(ctx *models.Context, uniqueID string) (*models.RefDocumentPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS, "uniqueId", "documentPolicyID", "ref.documentPolicyDocumentsId", "ref.documentPolicyDocumentsId")...)
	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONDOCUMENTMASTER, "documentMasterID", "uniqueId", "documentMaster", "documentMaster")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS,
			"as":   "ref.documentMaster",
			"let":  bson.M{"documentPolicyID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$$documentPolicyID", "$documentPolicyID"}},
					{"$eq": []string{"$status", constants.DOCUMENTMASTERSTATUSACTIVE}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"documentMaster": bson.M{"$arrayElemAt": []interface{}{"$documentMaster", 0}}}},
				{"$project": bson.M{"documentMaster": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.documentMaster": "$ref.documentMaster.documentMaster"}})

	d.Shared.BsonToJSONPrintTag("DocumentType query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentPolicys []models.RefDocumentPolicy
	var documentPolicy *models.RefDocumentPolicy
	if err = cursor.All(ctx.CTX, &documentPolicys); err != nil {
		return nil, err
	}
	if len(documentPolicys) > 0 {
		documentPolicy = &documentPolicys[0]
	}
	return documentPolicy, nil
}

//UpdateDocumentPolicy : ""
func (d *Daos) UpdateDocumentPolicy(ctx *models.Context, documentPolicy *models.DocumentPolicy) error {
	selector := bson.M{"uniqueId": documentPolicy.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": documentPolicy}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableDocumentPolicy :""
func (d *Daos) EnableDocumentPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTPOLICYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDocumentPolicy :""
func (d *Daos) DisableDocumentPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTPOLICYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDocumentPolicy :""
func (d *Daos) DeleteDocumentPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTPOLICYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDocumentPolicy : ""
func (d *Daos) FilterDocumentPolicy(ctx *models.Context, documentpolicyFilter *models.FilterDocumentPolicy, pagination *models.Pagination) ([]models.RefDocumentPolicy, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if documentpolicyFilter != nil {

		if len(documentpolicyFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": documentpolicyFilter.Status}})
		}
		if len(documentpolicyFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": documentpolicyFilter.OrganisationID}})
		}
		//Regex
		if documentpolicyFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: documentpolicyFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if documentpolicyFilter != nil {
		if documentpolicyFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{documentpolicyFilter.SortBy: documentpolicyFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONDOCUMENTMASTER, "documentMasterID", "uniqueId", "documentMaster", "documentMaster")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS,
			"as":   "ref.documentMaster",
			"let":  bson.M{"documentPolicyID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$$documentPolicyID", "$documentPolicyID"}},
					{"$eq": []string{"$status", constants.DOCUMENTMASTERSTATUSACTIVE}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"documentMaster": bson.M{"$arrayElemAt": []interface{}{"$documentMaster", 0}}}},
				{"$project": bson.M{"documentMaster": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.documentMaster": "$ref.documentMaster.documentMaster"}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentPolicys []models.RefDocumentPolicy
	if err = cursor.All(context.TODO(), &documentPolicys); err != nil {
		return nil, err
	}
	return documentPolicys, nil
}
func (d *Daos) GetSingleDocumentPolicyWithActiveName(ctx *models.Context, uniqueID string) (*models.RefDocumentPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.ONBOARDINGPOLICYSTATUSACTIVE}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS, "uniqueId", "documentPolicyID", "ref.documentPolicyDocumentsId", "ref.documentPolicyDocumentsId")...)

	d.Shared.BsonToJSONPrintTag("DocumentType query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentPolicys []models.RefDocumentPolicy
	var documentPolicy *models.RefDocumentPolicy
	if err = cursor.All(ctx.CTX, &documentPolicys); err != nil {
		return nil, err
	}
	fmt.Println("Lengtha==>", len(documentPolicys))
	if len(documentPolicys) > 0 {
		documentPolicy = &documentPolicys[0]
	}
	return documentPolicy, nil
}
