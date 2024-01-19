package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SavePayment : ""
func (d *Daos) SavePayment(ctx *models.Context, payment *models.Payment) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).InsertOne(ctx.CTX, payment)
	if err != nil {
		return err
	}
	payment.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdatePayment : ""
func (d *Daos) UpdatePayment(ctx *models.Context, payment *models.Payment) error {
	selector := bson.M{"uniqueId": payment.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payment}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSinglePayment : ""
func (d *Daos) GetSinglePayment(ctx *models.Context, uniqueID string) (*models.RefPayment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBANKINFORMATION, "bankId", "uniqueId", "ref.bankId", "ref.bankId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Payments []models.RefPayment
	var Payment *models.RefPayment
	if err = cursor.All(ctx.CTX, &Payments); err != nil {
		return nil, err
	}
	if len(Payments) > 0 {
		Payment = &Payments[0]
	}
	return Payment, err
}

// EnablePayment : ""
func (d *Daos) EnablePayment(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisablePayment : ""
func (d *Daos) DisablePayment(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeletePayment :""
func (d *Daos) DeletePayment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPayment : ""
func (d *Daos) FilterPayment(ctx *models.Context, Payment *models.FilterPayment, pagination *models.Pagination) ([]models.RefPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Payment != nil {
		if len(Payment.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Payment.Status}})
		}
		if len(Payment.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": Payment.OrganisationId}})
		}
		//Regex
		if Payment.Regex.Name != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: Payment.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBANKINFORMATION, "bankId", "uniqueId", "ref.bankId", "ref.bankId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var paymentFilter []models.RefPayment
	if err = cursor.All(context.TODO(), &paymentFilter); err != nil {
		return nil, err
	}
	return paymentFilter, nil
}

//PaymentPending : ""
func (d *Daos) PaymentPending(ctx *models.Context, payment *models.Payment) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).InsertOne(ctx.CTX, payment)
	if err != nil {
		return err
	}
	payment.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// PaymentAccept : ""
func (d *Daos) PaymentAccept(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYMENTSTATUSACCEPT}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// PaymentReject : ""
func (d *Daos) PaymentReject(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYMENTSTATUSREJECT}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// PaymentReceived : ""
func (d *Daos) PaymentReceived(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYMENTSTATUSRECEVIED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}
