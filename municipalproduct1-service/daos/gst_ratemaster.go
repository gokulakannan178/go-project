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

// SaveGSTRateMaster : ""
func (d *Daos) SaveGSTRateMaster(ctx *models.Context, gst *models.GSTRateMaster) error {
	d.Shared.BsonToJSONPrint(gst)
	_, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).InsertOne(ctx.CTX, gst)
	return err
}

// GetSingleGSTRateMaster : ""
func (d *Daos) GetSingleGSTRateMaster(ctx *models.Context, UniqueID string) (*models.GSTRateMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gsts []models.GSTRateMaster
	var gst *models.GSTRateMaster
	if err = cursor.All(ctx.CTX, &gsts); err != nil {
		return nil, err
	}
	if len(gsts) > 0 {
		gst = &gsts[0]
	}
	return gst, nil
}

// UpdateGSTRateMaster : ""
func (d *Daos) UpdateGSTRateMaster(ctx *models.Context, gst *models.GSTRateMaster) error {
	selector := bson.M{"uniqueId": gst.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": gst}
	_, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableGSTRateMaster : ""
func (d *Daos) EnableGSTRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GSTRATEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableGSTRateMaster : ""
func (d *Daos) DisableGSTRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GSTRATEMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteGSTRateMaster : ""
func (d *Daos) DeleteGSTRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GSTRATEMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterGSTRateMaster : ""
func (d *Daos) FilterGSTRateMaster(ctx *models.Context, gstfilter *models.GSTRateMasterFilter, pagination *models.Pagination) ([]models.GSTRateMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if gstfilter != nil {
		if len(gstfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": gstfilter.Status}})
		}
		if len(gstfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": gstfilter.UniqueID}})
		}
		if len(gstfilter.Name) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": gstfilter.Name}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGSTRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gsts []models.GSTRateMaster
	if err = cursor.All(context.TODO(), &gsts); err != nil {
		return nil, err
	}
	return gsts, nil
}
