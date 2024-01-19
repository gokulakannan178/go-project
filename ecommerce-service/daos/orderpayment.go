package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveOrderPayment : ""
func (d *Daos) SaveOrderPayment(ctx *models.Context, OrderPayment *models.OrderPayment) error {
	d.Shared.BsonToJSONPrint(OrderPayment)
	_, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).InsertOne(ctx.CTX, OrderPayment)
	return err
}

// GetSingleOrderPayment : ""
func (d *Daos) GetSingleOrderPayment(ctx *models.Context, UniqueID string) (*models.RefOrderPayment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Lookups
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OrderPayments []models.RefOrderPayment
	var OrderPayment *models.RefOrderPayment
	if err = cursor.All(ctx.CTX, &OrderPayments); err != nil {
		return nil, err
	}
	if len(OrderPayments) > 0 {
		OrderPayment = &OrderPayments[0]
	}
	return OrderPayment, nil
}

// UpdateOrderPayment : ""
func (d *Daos) UpdateOrderPayment(ctx *models.Context, OrderPayment *models.OrderPayment) error {
	selector := bson.M{"uniqueId": OrderPayment.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": OrderPayment}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableOrderPayment : ""
func (d *Daos) EnableOrderPayment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ORDERPAYMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableOrderPayment : ""
func (d *Daos) DisableOrderPayment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ORDERPAYMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteOrderPayment : ""
func (d *Daos) DeleteOrderPayment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ORDERPAYMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOrderPayment : ""
func (d *Daos) FilterOrderPayment(ctx *models.Context, filter *models.OrderPaymentFilter, pagination *models.Pagination) ([]models.RefOrderPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		//Regex Using searchBox Struct
		if filter.SearchText.Payee != "" {
			query = append(query, bson.M{"payee": primitive.Regex{Pattern: filter.SearchText.Payee, Options: "xi"}})
		}

		if filter.DateRange != nil {
			//var sd,ed time.Time
			if filter.DateRange.From != nil {
				sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				}
				query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.AmountRange != nil {
			//var sd,ed time.Time
			query = append(query, bson.M{"amount": bson.M{"$gte": filter.AmountRange.From, "$lte": filter.AmountRange.To}})

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
	d.Shared.BsonToJSONPrintTag("orderpaymet query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OrderPayment []models.RefOrderPayment
	if err = cursor.All(context.TODO(), &OrderPayment); err != nil {
		return nil, err
	}
	return OrderPayment, nil
}

// GetSingleOrderPayment : ""
func (d *Daos) GetCompletedOrderPayment(ctx *models.Context, UniqueID string) ([]models.RefOrderPayment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"orderId": UniqueID, "status": constants.ORDERPAYMENTSTATUSCOMPLETED}})

	d.Shared.BsonToJSONPrintTag("orderpayment query =>", mainPipeline)
	//Lookups
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDERPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OrderPayment []models.RefOrderPayment
	if err = cursor.All(context.TODO(), &OrderPayment); err != nil {
		return nil, err
	}
	return OrderPayment, nil
}
