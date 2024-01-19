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

//SaveDocumentMuxMaster :""
func (d *Daos) SaveDocumentMuxMaster(ctx *models.Context, documentMuxMaster *models.DocumentMuxMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).InsertOne(ctx.CTX, documentMuxMaster)
	if err != nil {
		return err
	}
	documentMuxMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDocumentMuxMaster : ""
func (d *Daos) GetSingleDocumentMuxMaster(ctx *models.Context, uniqueID string) (*models.RefDocumentMuxMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDocumentMuxMasterCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentMuxMasters []models.RefDocumentMuxMaster
	var documentMuxMaster *models.RefDocumentMuxMaster
	if err = cursor.All(ctx.CTX, &documentMuxMasters); err != nil {
		return nil, err
	}
	if len(documentMuxMasters) > 0 {
		documentMuxMaster = &documentMuxMasters[0]
	}
	return documentMuxMaster, nil
}

//UpdateDocumentMuxMaster : ""
func (d *Daos) UpdateDocumentMuxMaster(ctx *models.Context, documentMuxMaster *models.DocumentMuxMaster) error {
	selector := bson.M{"uniqueId": documentMuxMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": documentMuxMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableDocumentMuxMaster :""
func (d *Daos) EnableDocumentMuxMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTMUXMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDocumentMuxMaster :""
func (d *Daos) DisableDocumentMuxMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTMUXMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDocumentMuxMaster :""
func (d *Daos) DeleteDocumentMuxMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTMUXMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDocumentMuxMaster : ""
func (d *Daos) FilterDocumentMuxMaster(ctx *models.Context, filter *models.FilterDocumentMuxMaster, pagination *models.Pagination) ([]models.RefDocumentMuxMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDocumentMuxMasterCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentMuxMaster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTMUXMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentMuxMasters []models.RefDocumentMuxMaster
	if err = cursor.All(context.TODO(), &documentMuxMasters); err != nil {
		return nil, err
	}
	return documentMuxMasters, nil
}
