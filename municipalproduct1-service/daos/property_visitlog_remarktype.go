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

// SavePropertyVisitLogRemarkType : ""
func (d *Daos) SavePropertyVisitLogRemarkType(ctx *models.Context, remark *models.PropertyVisitLogRemarkType) error {
	d.Shared.BsonToJSONPrint(remark)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).InsertOne(ctx.CTX, remark)
	return err
}

// GetSinglePropertyVisitLogRemarkType : ""
func (d *Daos) GetSinglePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) (*models.RefPropertyVisitLogRemarkType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("letter upload getsingle query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var remarks []models.RefPropertyVisitLogRemarkType
	var remark *models.RefPropertyVisitLogRemarkType
	if err = cursor.All(ctx.CTX, &remarks); err != nil {
		return nil, err
	}
	if len(remarks) > 0 {
		remark = &remarks[0]
	}
	return remark, nil
}

// UpdatePropertyVisitLogRemarkType : ""
func (d *Daos) UpdatePropertyVisitLogRemarkType(ctx *models.Context, remark *models.PropertyVisitLogRemarkType) error {
	selector := bson.M{"uniqueId": remark.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": remark, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyVisitLogRemarkType : ""
func (d *Daos) EnablePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVISITLOGREMARKTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyVisitLogRemarkType : ""
func (d *Daos) DisablePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVISITLOGREMARKTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyVisitLogRemarkType : ""
func (d *Daos) DeletePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVISITLOGREMARKTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyVisitLogRemarkType : ""
func (d *Daos) FilterPropertyVisitLogRemarkType(ctx *models.Context, filter *models.PropertyVisitLogRemarkTypeFilter, pagination *models.Pagination) ([]models.RefPropertyVisitLogRemarkType, error) {
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
		mainPipeline = append(mainPipeline, []bson.M{{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var remarks []models.RefPropertyVisitLogRemarkType
	if err = cursor.All(context.TODO(), &remarks); err != nil {
		return nil, err
	}
	return remarks, nil
}
