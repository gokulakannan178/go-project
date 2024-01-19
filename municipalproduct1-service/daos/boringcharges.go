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

// SaveBoringCharges : ""
func (d *Daos) SaveBoringCharges(ctx *models.Context, boringCharges *models.BoringCharges) error {
	d.Shared.BsonToJSONPrint(boringCharges)
	_, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).InsertOne(ctx.CTX, boringCharges)
	return err
}

// GetSingleBoringCharges : ""
func (d *Daos) GetSingleBoringCharges(ctx *models.Context, UniqueID string) (*models.RefBoringCharges, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var boringcharges []models.RefBoringCharges
	var boringcharge *models.RefBoringCharges
	if err = cursor.All(ctx.CTX, &boringcharges); err != nil {
		return nil, err
	}
	if len(boringcharges) > 0 {
		boringcharge = &boringcharges[0]
	}
	return boringcharge, nil
}

// UpdateBoringCharges : ""
func (d *Daos) UpdateBoringCharges(ctx *models.Context, boringCharges *models.BoringCharges) error {
	selector := bson.M{"uniqueId": boringCharges.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": boringCharges}
	_, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableBoringCharges : ""
func (d *Daos) EnableBoringCharges(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BORINGCHARGESSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableBoringCharges: ""
func (d *Daos) DisableBoringCharges(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BORINGCHARGESSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteBoringCharges: ""
func (d *Daos) DeleteBoringCharges(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BORINGCHARGESSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterBoringCharges : ""
func (d *Daos) FilterBoringCharges(ctx *models.Context, filter *models.BoringChargesFilter, pagination *models.Pagination) ([]models.RefBoringCharges, error) {
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

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"requestor.on": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBORINGCHARGES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var boringcharges []models.RefBoringCharges
	if err = cursor.All(context.TODO(), &boringcharges); err != nil {
		return nil, err
	}
	return boringcharges, nil
}
