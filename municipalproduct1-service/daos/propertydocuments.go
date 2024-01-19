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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePropertyDocument :""
func (d *Daos) SavePropertyDocument(ctx *models.Context, propertyDocument *models.PropertyDocuments) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).InsertOne(ctx.CTX, propertyDocument)
	return err
}

// SavePropertyDocumentv2 :""
func (d *Daos) SavePropertyDocumentv2(ctx *models.Context, db *mongo.Database, sc *mongo.SessionContext, propertyDocument []models.PropertyDocuments) error {
	for _, v := range propertyDocument {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"propertyId": v.PropertyID, "documentId": v.DocumentID}
		updateData := bson.M{"$set": v}
		if _, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//
// func (d *Daos) SavePropertyDocumentv2(ctx *models.Context, db *mongo.Database, sc *mongo.SessionContext, propertyDocument []models.PropertyDocuments) error {
// 	insertdata := []interface{}{}
// 	for _, v := range propertyDocument {
// 		insertdata = append(insertdata, v)
// 	}
// 	result, err := db.Collection(constants.COLLECTIONPROPERTYDOCUMENT).InsertMany(ctx.CTX, insertdata)
// 	fmt.Println("insert result =>", result)
// 	return err
// }

//GetSinglePropertyDocument : ""
func (d *Daos) GetSinglePropertyDocument(ctx *models.Context, UniqueID string) (*models.RefPropertyDocuments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Lookups
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyDocuments []models.RefPropertyDocuments
	var propertyDocument *models.RefPropertyDocuments
	if err = cursor.All(ctx.CTX, &propertyDocuments); err != nil {
		return nil, err
	}
	if len(propertyDocuments) > 0 {
		propertyDocument = &propertyDocuments[0]
	}
	return propertyDocument, nil
}

//UpdatePropertyDocument : ""
func (d *Daos) UpdatePropertyDocument(ctx *models.Context, propertyDocument *models.PropertyDocuments) error {
	selector := bson.M{"uniqueId": propertyDocument.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyDocument, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPropertyDocument : ""
func (d *Daos) FilterPropertyDocument(ctx *models.Context, filter *models.PropertyDocumentsFilter, pagination *models.Pagination) ([]models.RefPropertyDocuments, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("PropertyDocument query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyDocuments []models.RefPropertyDocuments
	if err = cursor.All(context.TODO(), &propertyDocuments); err != nil {
		return nil, err
	}
	return propertyDocuments, nil
}

//EnablePropertyDocument :""
func (d *Daos) EnablePropertyDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYDOCUMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyDocument :""
func (d *Daos) DisablePropertyDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYDOCUMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyDocument :""
func (d *Daos) DeletePropertyDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYDOCUMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RemoveCronLog : ""
func (d *Daos) UpdateMultiplePropertyDocuments(ctx *models.Context, name string) error {
	selector := bson.M{"name": name}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": bson.M{"isCurrentScript": false}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).UpdateMany(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
