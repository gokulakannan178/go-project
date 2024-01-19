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

// SaveMajorUpdate : ""
func (d *Daos) SaveMajorUpdate(ctx *models.Context, pmTarget *models.MajorUpdate) error {
	d.Shared.BsonToJSONPrint(pmTarget)
	_, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).InsertOne(ctx.CTX, pmTarget)
	return err
}

// GetSingleMajorUpdate : ""
func (d *Daos) GetSingleMajorUpdate(ctx *models.Context, UniqueID string) (*models.RefMajorUpdate, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefMajorUpdate
	var tower *models.RefMajorUpdate
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateMajorUpdate : ""
func (d *Daos) UpdateMajorUpdate(ctx *models.Context, majorUpdate *models.MajorUpdate) error {
	selector := bson.M{"uniqueId": majorUpdate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": majorUpdate}
	_, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMajorUpdate : ""
func (d *Daos) EnableMajorUpdate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MAJORUPDATESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableMajorUpdate : ""
func (d *Daos) DisableMajorUpdate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MAJORUPDATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteMajorUpdate : ""
func (d *Daos) DeleteMajorUpdate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MAJORUPDATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMajorUpdate : ""
func (d *Daos) FilterMajorUpdate(ctx *models.Context, filter *models.MajorUpdateFilter, pagination *models.Pagination) ([]models.RefMajorUpdate, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMAJORUPDATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pmTarget []models.RefMajorUpdate
	if err = cursor.All(context.TODO(), &pmTarget); err != nil {
		return nil, err
	}
	return pmTarget, nil
}
