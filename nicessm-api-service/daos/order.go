package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveOrder :""
func (d *Daos) SaveOrder(ctx *models.Context, order *models.Order) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONORDER).InsertOne(ctx.CTX, order)
	if err != nil {
		return err
	}
	order.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleOrder : ""
func (d *Daos) GetSingleOrder(ctx *models.Context, UniqueID string) (*models.RefOrder, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var orders []models.RefOrder
	var order *models.RefOrder
	if err = cursor.All(ctx.CTX, &orders); err != nil {
		return nil, err
	}
	if len(orders) > 0 {
		order = &orders[0]
	}
	return order, nil
}

//UpdateOrder : ""
func (d *Daos) UpdateOrder(ctx *models.Context, order *models.Order) error {

	selector := bson.M{"_id": order.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": order}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOrder : ""
func (d *Daos) FilterOrder(ctx *models.Context, filter *models.OrderFilter, pagination *models.Pagination) ([]models.RefOrder, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": filter.ActiveStatus}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		//Regex
		if filter.SearchBox.FromName != "" {
			query = append(query, bson.M{"orderFrom.name": primitive.Regex{Pattern: filter.SearchBox.FromName, Options: "xi"}})
		}
		if filter.SearchBox.FromMobile != "" {
			query = append(query, bson.M{"orderFrom.mobile": primitive.Regex{Pattern: filter.SearchBox.FromMobile, Options: "xi"}})
		}
	}
	//daterange
	if filter.DateOredrRange != nil {
		//var sd,ed time.Time
		if filter.DateOredrRange.From != nil {
			sd := time.Date(filter.DateOredrRange.From.Year(), filter.DateOredrRange.From.Month(), filter.DateOredrRange.From.Day(), 0, 0, 0, 0, filter.DateOredrRange.From.Location())
			ed := time.Date(filter.DateOredrRange.From.Year(), filter.DateOredrRange.From.Month(), filter.DateOredrRange.From.Day(), 23, 59, 59, 0, filter.DateOredrRange.From.Location())
			if filter.DateOredrRange.To != nil {
				ed = time.Date(filter.DateOredrRange.To.Year(), filter.DateOredrRange.To.Month(), filter.DateOredrRange.To.Day(), 23, 59, 59, 0, filter.DateOredrRange.To.Location())
			}
			query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONORDER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Order query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var orders []models.RefOrder
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

//EnableOrder :""
func (d *Daos) EnableOrder(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORDERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOrder :""
func (d *Daos) DisableOrder(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORDERSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOrder :""
func (d *Daos) DeleteOrder(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORDERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//SaveOrder :""
func (d *Daos) CreateOrder(ctx *models.Context, order *models.Order) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONORDER).InsertOne(ctx.CTX, order)
	if err != nil {
		return err
	}
	order.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
