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

// SavePaymentMode : ""
func (d *Daos) SavePaymentMode(ctx *models.Context, PaymentMode *models.PaymentMode) error {
	d.Shared.BsonToJSONPrint(PaymentMode)
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).InsertOne(ctx.CTX, PaymentMode)
	return err
}

// GetSinglePaymentMode : ""
func (d *Daos) GetSinglePaymentMode(ctx *models.Context, UniqueID string) (*models.RefPaymentMode, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var PaymentModes []models.RefPaymentMode
	var PaymentMode *models.RefPaymentMode
	if err = cursor.All(ctx.CTX, &PaymentModes); err != nil {
		return nil, err
	}
	if len(PaymentModes) > 0 {
		PaymentMode = &PaymentModes[0]
	}
	return PaymentMode, nil
}

// UpdatePaymentMode : ""
func (d *Daos) UpdatePaymentMode(ctx *models.Context, PaymentMode *models.PaymentMode) error {
	selector := bson.M{"uniqueId": PaymentMode.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": PaymentMode}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePaymentMode : ""
func (d *Daos) EnablePaymentMode(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYMENTMODESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePaymentMode : ""
func (d *Daos) DisablePaymentMode(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYMENTMODESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePaymentMode : ""
func (d *Daos) DeletePaymentMode(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYMENTMODESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPaymentMode : ""
func (d *Daos) FilterPaymentMode(ctx *models.Context, filter *models.PaymentModeFilter, pagination *models.Pagination) ([]models.RefPaymentMode, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		//Regex Using searchBox Struct
		if filter.SearchText.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchText.Name, Options: "xi"}})
		}
		// if filter.SearchText.UniqueID != "" {
		// 	query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		// }
		// if filter.DateRange != nil {
		// 	//var sd,ed time.Time
		// 	if filter.DateRange.From != nil {
		// 		sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
		// 		ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
		// 		if filter.DateRange.To != nil {
		// 			ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
		// 		}
		// 		query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

		// 	}
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("paymentmode query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENTMODE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var PaymentMode []models.RefPaymentMode
	if err = cursor.All(context.TODO(), &PaymentMode); err != nil {
		return nil, err
	}
	return PaymentMode, nil
}
