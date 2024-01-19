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

//SavePropertyRequiredDocument :""
func (d *Daos) SavePropertyRequiredDocument(ctx *models.Context, propertyRequiredDocument *models.PropertyRequiredDocument) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).InsertOne(ctx.CTX, propertyRequiredDocument)
	return err
}

//GetSinglePropertyRequiredDocument : ""
func (d *Daos) GetSinglePropertyRequiredDocument(ctx *models.Context, code string) (*models.RefPropertyRequiredDocument, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": code}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTLIST, "documentId", "uniqueId", "ref.documentId", "ref.documentId")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyRequiredDocuments []models.RefPropertyRequiredDocument
	var propertyRequiredDocument *models.RefPropertyRequiredDocument
	if err = cursor.All(ctx.CTX, &propertyRequiredDocuments); err != nil {
		return nil, err
	}
	if len(propertyRequiredDocuments) > 0 {
		propertyRequiredDocument = &propertyRequiredDocuments[0]
	}
	return propertyRequiredDocument, nil
}

//UpdatePropertyRequiredDocument : ""
func (d *Daos) UpdatePropertyRequiredDocument(ctx *models.Context, propertyRequiredDocument *models.PropertyRequiredDocument) error {
	selector := bson.M{"uniqueId": propertyRequiredDocument.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyRequiredDocument, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPropertyRequiredDocument : ""
func (d *Daos) FilterPropertyRequiredDocument(ctx *models.Context, propertyRequiredDocumentfilter *models.PropertyRequiredDocumentFilter, pagination *models.Pagination) ([]models.RefPropertyRequiredDocument, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyRequiredDocumentfilter != nil {
		if len(propertyRequiredDocumentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyRequiredDocumentfilter.Status}})
		}
		if len(propertyRequiredDocumentfilter.For) > 0 {
			query = append(query, bson.M{"for": bson.M{"$in": propertyRequiredDocumentfilter.For}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyRequiredDocumentfilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyRequiredDocumentfilter.SortBy: propertyRequiredDocumentfilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTLIST, "documentId", "uniqueId", "ref.documentId", "ref.documentId")...)
	if propertyRequiredDocumentfilter != nil {
		if propertyRequiredDocumentfilter.IsPropertyDcumentsUploaded == true {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYDOCUMENT, "documentId", "documentId", "ref.document", "ref.document")...)
		}
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("propertyRequiredDocument query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyRequiredDocuments []models.RefPropertyRequiredDocument
	if err = cursor.All(context.TODO(), &propertyRequiredDocuments); err != nil {
		return nil, err
	}
	return propertyRequiredDocuments, nil
}

//EnablePropertyRequiredDocument :""
func (d *Daos) EnablePropertyRequiredDocument(ctx *models.Context, code string) error {
	query := bson.M{"uniqueId": code}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYREQUIREDDOCUMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyRequiredDocument :""
func (d *Daos) DisablePropertyRequiredDocument(ctx *models.Context, code string) error {
	query := bson.M{"uniqueId": code}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYREQUIREDDOCUMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyRequiredDocument :""
func (d *Daos) DeletePropertyRequiredDocument(ctx *models.Context, code string) error {
	query := bson.M{"uniqueId": code}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYREQUIREDDOCUMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYREQUIREDDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
