package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveOrder :""
func (d *Daos) SaveOrder(ctx *models.Context, order *models.Order) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONORDER).InsertOne(ctx.CTX, order)
	return err
}

//UpsertOrder :""
func (d *Daos) UpsertOrder(ctx *models.Context, order *models.Order) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": order.UniqueID}
	updateData := bson.M{"$set": order}
	if _, err := ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil
}

//GetSingleProduct : ""
func (d *Daos) GetSingleOrder(ctx *models.Context, uniqueID string) (*models.RefOrder, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})

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

// UpdateOrderStatus : ""
func (d *Daos) UpdateOrderStatus(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ORDERSTATUSACTIVE, "transportStatus": constants.ORDERSTATUSDELIVERED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// UpdateOrderStatus : ""
func (d *Daos) UpdateOrderPaymentStatus(ctx *models.Context, UniqueID string, status string, pendingamount float64) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"payment.status": status, "payment.pendingAmount": pendingamount, "status": status}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterOrder : ""
func (d *Daos) FilterOrder(ctx *models.Context, filter *models.OrderFilter, pagination *models.Pagination) ([]models.RefOrder, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.CompanyID) > 0 {
			query = append(query, bson.M{"from.id": bson.M{"$in": filter.CompanyID}})
		}
		if len(filter.CustomerID) > 0 {
			query = append(query, bson.M{"to.id": bson.M{"$in": filter.CustomerID}})
		}
		if len(filter.CustomerType) > 0 {
			query = append(query, bson.M{"to.type": bson.M{"$in": filter.CustomerType}})
		}
		if len(filter.CompanyType) > 0 {
			query = append(query, bson.M{"from.type": bson.M{"$in": filter.CompanyType}})
		}
		if len(filter.TransportID) > 0 {
			query = append(query, bson.M{"transport.companyId": bson.M{"$in": filter.TransportID}})
		}
		if len(filter.TransportStatus) > 0 {
			query = append(query, bson.M{"transport.status": bson.M{"$in": filter.TransportStatus}})
		}
		if len(filter.TransportType) > 0 {
			query = append(query, bson.M{"transport.type": bson.M{"$in": filter.TransportType}})
		}
		if len(filter.PaymentStatus) > 0 {
			query = append(query, bson.M{"paymentStatus": bson.M{"$in": filter.PaymentStatus}})
		}
		if len(filter.DriverID) > 0 {
			query = append(query, bson.M{"transport.driverId": bson.M{"$in": filter.DriverID}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
	d.Shared.BsonToJSONPrintTag("product query =>", mainPipeline)
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
