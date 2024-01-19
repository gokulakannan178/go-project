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

//SaveDocumentList :""
func (d *Daos) SaveDocumentList(ctx *models.Context, documentList *models.DocumentList) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).InsertOne(ctx.CTX, documentList)
	return err
}

//GetSingleDocumentList : ""
func (d *Daos) GetSingleDocumentList(ctx *models.Context, UniqueID string) (*models.RefDocumentList, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Lookups
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentLists []models.RefDocumentList
	var documentList *models.RefDocumentList
	if err = cursor.All(ctx.CTX, &documentLists); err != nil {
		return nil, err
	}
	if len(documentLists) > 0 {
		documentList = &documentLists[0]
	}
	return documentList, nil
}

//UpdateDocumentList : ""
func (d *Daos) UpdateDocumentList(ctx *models.Context, documentList *models.DocumentList) error {
	selector := bson.M{"uniqueId": documentList.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": documentList, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDocumentList : ""
func (d *Daos) FilterDocumentList(ctx *models.Context, filter *models.DocumentListFilter, pagination *models.Pagination) ([]models.RefDocumentList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("DocumentList query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentLists []models.RefDocumentList
	if err = cursor.All(context.TODO(), &documentLists); err != nil {
		return nil, err
	}
	return documentLists, nil
}

//EnableDocumentList :""
func (d *Daos) EnableDocumentList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDocumentList :""
func (d *Daos) DisableDocumentList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDocumentList :""
func (d *Daos) DeleteDocumentList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
