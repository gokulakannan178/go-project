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

//SaveOrderPayments :""
func (d *Daos) SaveOrderPayment(ctx *models.Context, payment *models.OrderPayment) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).InsertOne(ctx.CTX, payment)
	if err != nil {
		return err
	}
	payment.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleOrderPayments : ""
func (d *Daos) GetSingleOrderPayment(ctx *models.Context, UniqueID string) (*models.RefOrderPayment, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefOrderPayment
	var payment *models.RefOrderPayment
	if err = cursor.All(ctx.CTX, &payments); err != nil {
		return nil, err
	}
	if len(payments) > 0 {
		payment = &payments[0]
	}
	return payment, nil
}

//UpdateOrderPayments : ""
func (d *Daos) UpdateOrderPayment(ctx *models.Context, payment *models.OrderPayment) error {

	selector := bson.M{"_id": payment.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payment}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOrderPayments : ""
func (d *Daos) FilterOrderPayment(ctx *models.Context, filter *models.OrderPaymentFilter, pagination *models.Pagination) ([]models.RefOrderPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		//Regex
		if filter.SearchBox.PayeeName != "" {
			query = append(query, bson.M{"payeeName": primitive.Regex{Pattern: filter.SearchBox.PayeeName, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("OrderPayments query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payment []models.RefOrderPayment
	if err = cursor.All(context.TODO(), &payment); err != nil {
		return nil, err
	}
	return payment, nil
}

//EnableOrderPayments :""
func (d *Daos) EnableOrderPayment(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORDERPAYMENTSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOrderPayments :""
func (d *Daos) DisableOrderPayment(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORDERPAYMENTSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOrderPayments :""
func (d *Daos) DeleteOrderPayment(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORDERPAYMENTSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) GetpaymentbyOrderID(ctx *models.Context, OrderID string) ([]models.RefOrderPayment, error) {
	mainPipeline := []bson.M{}
	id, err := primitive.ObjectIDFromHex(OrderID)
	if err != nil {
		return nil, err
	}
	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"orderId": id}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("orderpayment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefOrderPayment
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}

	return payments, nil
}
func (d *Daos) UpdateOrderPaymentStatus(ctx *models.Context, UniqueID string, status string, pendingamount float64) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"payment.status": status, "payment.pendingAmount": pendingamount, "status": status}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORDER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
