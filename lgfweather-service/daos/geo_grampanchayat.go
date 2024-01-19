package daos

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveGramPanchayat :""
func (d *Daos) SaveGramPanchayat(ctx *models.Context, grampanchayat *models.GramPanchayat) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).InsertOne(ctx.CTX, grampanchayat)
	return err
}

//GetSingleGramPanchayat : ""
func (d *Daos) GetSingleGramPanchayat(ctx *models.Context, uniqueID string) (*models.RefGramPanchayat, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var grampanchayats []models.RefGramPanchayat
	var grampanchayat *models.RefGramPanchayat
	if err = cursor.All(ctx.CTX, &grampanchayats); err != nil {
		return nil, err
	}
	if len(grampanchayats) > 0 {
		grampanchayat = &grampanchayats[0]
	}
	return grampanchayat, nil
}

//UpdateGramPanchayat : ""
func (d *Daos) UpdateGramPanchayat(ctx *models.Context, grampanchayat *models.GramPanchayat) error {
	selector := bson.M{"uniqueId": grampanchayat.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": grampanchayat, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterGramPanchayat : ""
func (d *Daos) FilterGramPanchayat(ctx *models.Context, grampanchayatfilter *models.GramPanchayatFilter, pagination *models.Pagination) ([]models.RefGramPanchayat, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if grampanchayatfilter != nil {
		if len(grampanchayatfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": grampanchayatfilter.UniqueID}})
		}
		if len(grampanchayatfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": grampanchayatfilter.Status}})
		}
		if len(grampanchayatfilter.BlockCodes) > 0 {
			query = append(query, bson.M{"blockCode": bson.M{"$in": grampanchayatfilter.BlockCodes}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("grampanchayat query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var grampanchayats []models.RefGramPanchayat
	if err = cursor.All(context.TODO(), &grampanchayats); err != nil {
		return nil, err
	}
	return grampanchayats, nil
}

//EnableGramPanchayat :""
func (d *Daos) EnableGramPanchayat(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableGramPanchayat :""
func (d *Daos) DisableGramPanchayat(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteGramPanchayat :""
func (d *Daos) DeleteGramPanchayat(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetActiveGramPanchayat(ctx *models.Context) ([]models.GramPanchayat, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": constants.GRAMPANCHAYATSTATUSACTIVE})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("activestate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.GramPanchayat
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}
