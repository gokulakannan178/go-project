package daos

import (
	"context"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FilterPropertyOtherDemandPartPayment : ""
func (d *Daos) FilterPropertyOtherDemandPartPayment(ctx *models.Context, filter *models.PropertyOtherDemandPartPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyOtherDemandPartPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.TnxID) > 0 {
			query = append(query, bson.M{"tnxId": bson.M{"$in": filter.TnxID}})
		}
		if len(filter.ReciptNo) > 0 {
			query = append(query, bson.M{"reciptNo": bson.M{"$in": filter.ReciptNo}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if len(filter.PayeeName) > 0 {
			query = append(query, bson.M{"details.payeeName": bson.M{"$in": filter.PayeeName}})
		}
		if len(filter.CollectorID) > 0 {
			query = append(query, bson.M{"details.collector.id": bson.M{"$in": filter.CollectorID}})
		}
		if len(filter.MOP) > 0 {
			query = append(query, bson.M{"details.mop.mode": bson.M{"$in": filter.MOP}})
		}
		if len(filter.MadeAt) > 0 {
			query = append(query, bson.M{"details.madeAt.at": bson.M{"$in": filter.MadeAt}})
		}

		if filter.PaymentDateRange != nil {
			//var sd,ed time.Time
			if filter.PaymentDateRange.From != nil {
				sd := time.Date(filter.PaymentDateRange.From.Year(), filter.PaymentDateRange.From.Month(), filter.PaymentDateRange.From.Day(), 0, 0, 0, 0, filter.PaymentDateRange.From.Location())
				ed := time.Date(filter.PaymentDateRange.From.Year(), filter.PaymentDateRange.From.Month(), filter.PaymentDateRange.From.Day(), 23, 59, 59, 0, filter.PaymentDateRange.From.Location())
				if filter.PaymentDateRange.To != nil {
					ed = time.Date(filter.PaymentDateRange.To.Year(), filter.PaymentDateRange.To.Month(), filter.PaymentDateRange.To.Day(), 23, 59, 59, 0, filter.PaymentDateRange.To.Location())
				}
				query = append(query, bson.M{"paymentDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.CompletionDateRange != nil {
			//var sd,ed time.Time
			if filter.CompletionDateRange.From != nil {
				sd := time.Date(filter.CompletionDateRange.From.Year(), filter.CompletionDateRange.From.Month(), filter.CompletionDateRange.From.Day(), 0, 0, 0, 0, filter.CompletionDateRange.From.Location())
				ed := time.Date(filter.CompletionDateRange.From.Year(), filter.CompletionDateRange.To.Month(), filter.CompletionDateRange.To.Day(), 23, 59, 59, 0, filter.CompletionDateRange.To.Location())
				if filter.CompletionDateRange.To != nil {
					ed = time.Date(filter.CompletionDateRange.To.Year(), filter.CompletionDateRange.To.Month(), filter.CompletionDateRange.To.Day(), 23, 59, 59, 0, filter.CompletionDateRange.To.Location())
				}
				query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		//Regex Using searchBox Struct
		if filter.Regex.ReciptNo != "" {
			query = append(query, bson.M{"reciptNo": primitive.Regex{Pattern: filter.Regex.ReciptNo, Options: "xi"}})
		}
		if filter.Regex.PropertyID != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.Regex.PropertyID, Options: "xi"}})
		}
		if filter.Regex.PayeeName != "" {
			query = append(query, bson.M{"details.payeeName": primitive.Regex{Pattern: filter.Regex.PayeeName, Options: "xi"}})
		}
		if filter.Address != nil {
			if len(filter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": filter.Address.StateCode}})
			}
			if len(filter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": filter.Address.DistrictCode}})
			}
			if len(filter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": filter.Address.VillageCode}})
			}
			if len(filter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": bson.M{"$in": filter.Address.ZoneCode}})
			}
			if len(filter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": bson.M{"$in": filter.Address.WardCode}})
			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPARTPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.id", "userName", "ref.collector", "ref.collector")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("propertyPartPayment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPARTPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var partPayments []models.RefPropertyOtherDemandPartPayment
	if err = cursor.All(context.TODO(), &partPayments); err != nil {
		return nil, err
	}
	return partPayments, nil
}
