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

//SaveDocumentMaster :""
func (d *Daos) SaveDocumentMaster(ctx *models.Context, documentMaster *models.DocumentMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).InsertOne(ctx.CTX, documentMaster)
	if err != nil {
		return err
	}
	documentMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDocumentMaster : ""
func (d *Daos) GetSingleDocumentMaster(ctx *models.Context, uniqueID string) (*models.RefDocumentMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentmasters []models.RefDocumentMaster
	var documentmaster *models.RefDocumentMaster
	if err = cursor.All(ctx.CTX, &documentmasters); err != nil {
		return nil, err
	}
	if len(documentmasters) > 0 {
		documentmaster = &documentmasters[0]
	}
	return documentmaster, nil
}

//GetSingleDocumentMasterWithActive : ""
func (d *Daos) GetSingleDocumentMasterWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefDocumentMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentmasters []models.RefDocumentMaster
	var documentmaster *models.RefDocumentMaster
	if err = cursor.All(ctx.CTX, &documentmasters); err != nil {
		return nil, err
	}
	if len(documentmasters) > 0 {
		documentmaster = &documentmasters[0]
	}
	return documentmaster, nil
}

//UpdateDocumentMaster : ""
func (d *Daos) UpdateDocumentMaster(ctx *models.Context, documentMaster *models.DocumentMaster) error {
	selector := bson.M{"uniqueId": documentMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": documentMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableDocumentMaster :""
func (d *Daos) EnableDocumentMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDocumentMaster :""
func (d *Daos) DisableDocumentMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDocumentMaster :""
func (d *Daos) DeleteDocumentMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDocumentMaster : ""
func (d *Daos) FilterDocumentMaster(ctx *models.Context, documentmasterfilter *models.FilterDocumentMaster, pagination *models.Pagination) ([]models.RefDocumentMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if documentmasterfilter != nil {

		if len(documentmasterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": documentmasterfilter.Status}})
		}
		if len(documentmasterfilter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": documentmasterfilter.OrganisationId}})
		}
		//Regex
		if documentmasterfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: documentmasterfilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if documentmasterfilter != nil {
		if documentmasterfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{documentmasterfilter.SortBy: documentmasterfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Document Master query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentMasters []models.RefDocumentMaster
	if err = cursor.All(context.TODO(), &documentMasters); err != nil {
		return nil, err
	}
	return documentMasters, nil
}
