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

//SaveOwnership :""
func (d *Daos) SaveOwnership(ctx *models.Context, ownership *models.Ownership) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).InsertOne(ctx.CTX, ownership)
	return err
}

//GetSingleOwnership : ""
func (d *Daos) GetSingleOwnership(ctx *models.Context, UniqueID string) (*models.RefOwnership, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ownerships []models.RefOwnership
	var ownership *models.RefOwnership
	if err = cursor.All(ctx.CTX, &ownerships); err != nil {
		return nil, err
	}
	if len(ownerships) > 0 {
		ownership = &ownerships[0]
	}
	return ownership, nil
}

//UpdateOwnership : ""
func (d *Daos) UpdateOwnership(ctx *models.Context, ownership *models.Ownership) error {
	selector := bson.M{"uniqueId": ownership.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ownership, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOwnership : ""
func (d *Daos) FilterOwnership(ctx *models.Context, ownershipfilter *models.OwnershipFilter, pagination *models.Pagination) ([]models.RefOwnership, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if ownershipfilter != nil {

		if len(ownershipfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": ownershipfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("ownership query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ownerships []models.RefOwnership
	if err = cursor.All(context.TODO(), &ownerships); err != nil {
		return nil, err
	}
	return ownerships, nil
}

//EnableOwnership :""
func (d *Daos) EnableOwnership(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOWNERSHIPSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOwnership :""
func (d *Daos) DisableOwnership(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOWNERSHIPSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOwnership :""
func (d *Daos) DeleteOwnership(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOWNERSHIPSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOWNERSHIP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
