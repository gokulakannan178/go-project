package daos

import (
	"context"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//SavePayment :""
func (d *Daos) SavePayment(ctx *models.Context, product *models.Payment) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).InsertOne(ctx.CTX, product)
	return err
}

//FilterPayment : ""
func (d *Daos) FilterPayment(ctx *models.Context, paymentFilter *models.PaymentFilter, pagination *models.Pagination) ([]models.RefPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if paymentFilter != nil {
		if len(paymentFilter.SaleID) > 0 {
			query = append(query, bson.M{"saleId": bson.M{"$in": paymentFilter.SaleID}})
		}
		if len(paymentFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": paymentFilter.Status}})
		}
		if len(paymentFilter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": paymentFilter.Type}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if paymentFilter != nil {
		if paymentFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{paymentFilter.SortBy: paymentFilter.SortOrder}})
		}
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
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCTCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefPayment
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil
}
