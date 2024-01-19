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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveDocumentPolicyDocuments :""
func (d *Daos) SaveDocumentPolicyDocuments(ctx *models.Context, documentpolicydocuments *models.DocumentPolicyDocuments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).InsertOne(ctx.CTX, documentpolicydocuments)
	if err != nil {
		return err
	}
	documentpolicydocuments.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDocumentPolicyDocuments : ""
func (d *Daos) GetSingleDocumentPolicyDocuments(ctx *models.Context, uniqueID string) (*models.RefDocumentPolicyDocuments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDocumentMuxMasterCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentpolicydocuments []models.RefDocumentPolicyDocuments
	var documentpolicydocument *models.RefDocumentPolicyDocuments
	if err = cursor.All(ctx.CTX, &documentpolicydocuments); err != nil {
		return nil, err
	}
	if len(documentpolicydocuments) > 0 {
		documentpolicydocument = &documentpolicydocuments[0]
	}
	return documentpolicydocument, nil
}

// DocumentPolicyDocumentsRemoveNotPresentValue : ""
func (d *Daos) DocumentPolicyDocumentsRemoveNotPresentValue(ctx *models.Context, documentpolicyId string, arrayValue []string) error {
	selector := bson.M{"documentPolicyID": documentpolicyId, "documentMasterID": bson.M{"$nin": arrayValue}}
	//	d.Shared.BsonToJSONPrintTag("selector query in DocumentpolicyDocuments =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.DOCUMENTPOLICYDOCUMENTSSTATUSDELETED}}
	//d.Shared.BsonToJSONPrintTag("data query in DocumentpolicyDocuments =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).UpdateMany(ctx.CTX, selector, data)
	return err
}

// DocumentPolicyDocumentsUpsert : ""
func (d *Daos) DocumentPolicyDocumentsUpsert(ctx *models.Context, documentpolicyId string, arrayValue []string, name string) error {
	//fmt.Println("arrayValue", arrayValue)
	for _, v := range arrayValue {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"documentPolicyID": documentpolicyId, "documentMasterID": v}
		fmt.Println("updateQuery===>", updateQuery)
		documentPolicyDocuments := new(models.DocumentPolicyDocuments)
		documentPolicyDocuments.Status = constants.DOCUMENTPOLICYDOCUMENTSSTATUSACTIVE
		documentPolicyDocuments.Name = name
		documentPolicyDocuments.UniqueID = d.GetUniqueID(ctx, constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS)
		//fmt.Println("present added =======>", documentPolicyDocuments.UniqueID)
		updateData := bson.M{"$set": documentPolicyDocuments}
		if _, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//UpdateDocumentPolicyDocuments : ""
func (d *Daos) UpdateDocumentPolicyDocuments(ctx *models.Context, documentpolicydocuments *models.DocumentPolicyDocuments) error {
	selector := bson.M{"uniqueId": documentpolicydocuments.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": documentpolicydocuments}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableDocumentPolicyDocuments :""
func (d *Daos) EnableDocumentPolicyDocuments(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTPOLICYDOCUMENTSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDocumentPolicyDocuments :""
func (d *Daos) DisableDocumentPolicyDocuments(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTPOLICYDOCUMENTSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDocumentPolicyDocuments :""
func (d *Daos) DeleteDocumentPolicyDocuments(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DOCUMENTPOLICYDOCUMENTSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDocumentPolicyDocuments : ""
func (d *Daos) FilterDocumentPolicyDocuments(ctx *models.Context, documentpolicydocumentsFilter *models.FilterDocumentPolicyDocuments, pagination *models.Pagination) ([]models.RefDocumentPolicyDocuments, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if documentpolicydocumentsFilter != nil {

		if len(documentpolicydocumentsFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": documentpolicydocumentsFilter.Status}})
		}
		//Regex
		if documentpolicydocumentsFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: documentpolicydocumentsFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if documentpolicydocumentsFilter != nil {
		if documentpolicydocumentsFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{documentpolicydocumentsFilter.SortBy: documentpolicydocumentsFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("DocumentPolicyDocuments query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var documentpolicydocuments []models.RefDocumentPolicyDocuments
	if err = cursor.All(context.TODO(), &documentpolicydocuments); err != nil {
		return nil, err
	}
	return documentpolicydocuments, nil
}
