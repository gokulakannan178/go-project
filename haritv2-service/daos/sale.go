package daos

import (
	"context"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveOrder :""
func (d *Daos) SaveSale(ctx *models.Context, sale *models.Sale) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSALE).InsertOne(ctx.CTX, sale)
	return err
}
func (d *Daos) FilterSale(ctx *models.Context, filter *models.SaleFilter, pagination *models.Pagination) ([]models.RefSale, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.CompanyID) > 0 {
			query = append(query, bson.M{"company.id": bson.M{"$in": filter.CompanyID}})
		}
		if len(filter.CustomerID) > 0 {
			query = append(query, bson.M{"customer.id": bson.M{"$in": filter.CustomerID}})
		}
		if len(filter.CustomerType) > 0 {
			query = append(query, bson.M{"customer.type": bson.M{"$in": filter.CustomerType}})
		}
		if len(filter.CompanyType) > 0 {
			query = append(query, bson.M{"company.type": bson.M{"$in": filter.CompanyType}})
		}
		if len(filter.PaymentStatus) > 0 {
			query = append(query, bson.M{"paymentStatus": bson.M{"$in": filter.PaymentStatus}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

	}
	//daterange
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
	if filter.DateRange != nil {
		//var sd,ed time.Time
		if filter.DateRange.From != nil {
			sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
			ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
			if filter.DateRange.To != nil {
				ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
			}
			query = append(query, bson.M{"createdOn.on": bson.M{"$gte": sd, "$lte": ed}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSALE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("order query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var sales []models.RefSale
	if err = cursor.All(context.TODO(), &sales); err != nil {
		return nil, err
	}
	return sales, nil
}
