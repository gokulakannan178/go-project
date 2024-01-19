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

// SavePropertyPenalty : ""
func (d *Daos) SavePropertyPenalty(ctx *models.Context, penalty *models.PenaltyLogs) error {
	d.Shared.BsonToJSONPrint(penalty)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).InsertOne(ctx.CTX, penalty)
	return err
}

// GetSinglePropertyPenalty : ""
func (d *Daos) GetSinglePropertyPenalty(ctx *models.Context, UniqueID string) (*models.PenaltyLogs, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var penaltys []models.PenaltyLogs
	var penalty *models.PenaltyLogs
	if err = cursor.All(ctx.CTX, &penaltys); err != nil {
		return nil, err
	}
	if len(penaltys) > 0 {
		penalty = &penaltys[0]
	}
	return penalty, nil
}

// UpdatePropertyPenalty : ""
func (d *Daos) UpdatePropertyPenalty(ctx *models.Context, penalty *models.PenaltyLogs) error {
	selector := bson.M{"uniqueId": penalty.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": penalty}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyPenalty : ""
func (d *Daos) EnablePropertyPenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPENALTYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyPenalty : ""
func (d *Daos) DisablePropertyPenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPENALTYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyPenalty : ""
func (d *Daos) DeletePropertyPenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPENALTYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyPenalty : ""
func (d *Daos) FilterPropertyPenalty(ctx *models.Context, filter *models.PropertyPenaltyFilter, pagination *models.Pagination) ([]models.RefPropertyPenalty, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPENALTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var penaltys []models.RefPropertyPenalty
	if err = cursor.All(context.TODO(), &penaltys); err != nil {
		return nil, err
	}
	return penaltys, nil
}
