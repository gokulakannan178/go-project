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

//SavePaymentGateway :""
func (d *Daos) SavePaymentGateway(ctx *models.Context, payment *models.PaymentGateway) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).InsertOne(ctx.CTX, payment)
	return err
}

//GetSinglePaymentGateway : ""
func (d *Daos) GetSinglePaymentGateway(ctx *models.Context, UniqueID string) (*models.PaymentGateway, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.PaymentGateway
	var user *models.PaymentGateway
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//GetSinglePaymentGateway : ""
func (d *Daos) GetDefaultPaymentGateway(ctx *models.Context) (*models.PaymentGateway, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active"}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.PaymentGateway
	var user *models.PaymentGateway
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//UpdatePaymentGateway : ""
func (d *Daos) UpdatePaymentGateway(ctx *models.Context, paymentgateway *models.PaymentGateway) error {
	selector := bson.M{"uniqueId": paymentgateway.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": paymentgateway}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPaymentGateway : ""
func (d *Daos) FilterPaymentGateway(ctx *models.Context, paymentgatewayfilter *models.PaymentGatewayFilter, pagination *models.Pagination) ([]models.PaymentGatewayFilter, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if paymentgatewayfilter != nil {
		if len(paymentgatewayfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": paymentgatewayfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.PaymentGatewayFilter
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

//EnableUser :""
func (d *Daos) EnablePaymentGateway(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYMENTGATEWAYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePaymentGateway :""
func (d *Daos) DisablePaymentGateway(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYMENTGATEWAYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePaymentGateway :""
func (d *Daos) DeletePaymentGateway(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYMENTGATEWAYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTGATEWAY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
