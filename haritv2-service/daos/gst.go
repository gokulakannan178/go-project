package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveGST :""
func (d *Daos) SaveGST(ctx *models.Context, gst *models.GST) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONGST).InsertOne(ctx.CTX, gst)
	return err
}

//GetSingleGST : ""
func (d *Daos) GetSingleGST(ctx *models.Context, uniqueID string) (*models.RefGST, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gsts []models.RefGST
	var gst *models.RefGST
	if err = cursor.All(ctx.CTX, &gsts); err != nil {
		return nil, err
	}
	if len(gsts) > 0 {
		gst = &gsts[0]
	}
	return gst, nil
}

//UpdateGST : ""
func (d *Daos) UpdateGST(ctx *models.Context, gst *models.GST) error {
	selector := bson.M{"uniqueId": gst.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": gst}
	_, err := ctx.DB.Collection(constants.COLLECTIONGST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableGST :""
func (d *Daos) EnableGST(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GSTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableGST :""
func (d *Daos) DisableGST(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GSTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteGST :""
func (d *Daos) DeleteGST(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GSTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterGST : ""
func (d *Daos) FilterGST(ctx *models.Context, gstfilter *models.GSTFilter, pagination *models.Pagination) ([]models.RefGST, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if gstfilter != nil {
		if len(gstfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": gstfilter.UniqueID}})
		}
		if len(gstfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": gstfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONGST).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("gst query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gsts []models.RefGST
	if err = cursor.All(context.TODO(), &gsts); err != nil {
		return nil, err
	}
	return gsts, nil
}

//GetSingleGST : ""
func (d *Daos) GetDefaultGST(ctx *models.Context) (*models.RefGST, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gsts []models.RefGST
	var gst *models.RefGST
	if err = cursor.All(ctx.CTX, &gsts); err != nil {
		return nil, err
	}
	if len(gsts) > 0 {
		gst = &gsts[0]
	}
	return gst, nil
}
