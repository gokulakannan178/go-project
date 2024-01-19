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

//SavePaymentGateway :""
func (d *Daos) SaveShopRentMonthlyPayment(ctx *models.Context, mtp *models.ShopRentMonthlyPayments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).InsertOne(ctx.CTX, mtp)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

//SaveShopRentPaymentFY :""
func (d *Daos) SaveShopMonthlyRentPaymentFYs(ctx *models.Context, mtpfy []models.ShopRenttMonthlyPaymentsfY) error {
	var insertData []interface{}
	for _, v := range mtpfy {
		insertData = append(insertData, v)
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSFY).InsertMany(ctx.CTX, insertData)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

func (d *Daos) GetSingleShopRentMonthlyPayment(ctx *models.Context, tnxID string) (*models.RefShopRentMonthlyPayments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tnxId": tnxID,
		},
	})
	mainPipeline = append(mainPipeline, d.RefQueryForShopRentPayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)

	d.Shared.BsonToJSONPrintTag("Get Single Shop Rent Payment Query  - ", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rmtps []models.RefShopRentMonthlyPayments
	var rmtp *models.RefShopRentMonthlyPayments
	if err = cursor.All(ctx.CTX, &rmtps); err != nil {
		return nil, err
	}
	if len(rmtps) > 0 {
		rmtp = &rmtps[0]
	}
	return rmtp, nil
}

// FilterShopRentPayment : ""
func (d *Daos) FilterShopRentMonthlyPayment(ctx *models.Context, filter *models.ShopRentMonthlyPaymentsFilter, pagination *models.Pagination) ([]models.RefShopRentMonthlyPayments, error) {
	var mainPipeline, query []bson.M
	if filter != nil {
		if len(filter.ShopRentID) > 0 {
			query = append(query, bson.M{"shopRentId": bson.M{"$in": filter.ShopRentID}})
		}

		if len(filter.MadeAT) > 0 {
			query = append(query, bson.M{"details.madeAt.at": bson.M{"$in": filter.MadeAT}})
		}

		if len(filter.MOP) > 0 {
			query = append(query, bson.M{"details.mop.mode": bson.M{"$in": filter.MOP}})
		}
		if len(filter.Scenario) > 0 {
			query = append(query, bson.M{"scenario": bson.M{"$in": filter.Scenario}})
		}
		if filter.SearchBox.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.SearchBox.OwnerName, Options: "xi"}})

		}
		if filter.SearchBox.ShopRentID != "" {
			query = append(query, bson.M{"shopRentId": primitive.Regex{Pattern: filter.SearchBox.ShopRentID, Options: "xi"}})

		}

		// if filter.SearchBox.OwnerName != "" {

		// 	shopRentIds, err := d.GetShopRentIDsWithOwnerNames(ctx, filter.SearchBox.OwnerName, filter.SearchBox.OwnerMobile)
		// 	if err != nil {
		// 		log.Println("ERR IN GETING - Shop Rent IDs WithOwner Names " + err.Error())
		// 	} else {
		// 		if len(shopRentIds) > 0 {
		// 			fmt.Println("got Shop Rent Ids - ", shopRentIds)
		// 			query = append(query, bson.M{"shopRentId": bson.M{"$in": shopRentIds}})
		// 		}
		// 	}
		// }

		if filter.SearchBox.OwnerMobile != "" {

			shopRentIds, err := d.GetShopRentIDsWithMobileNos(ctx, filter.SearchBox.OwnerName, filter.SearchBox.OwnerMobile)
			if err != nil {
				log.Println("ERR IN GETING - Shop Rent IDs With Mobile Numbers " + err.Error())
			} else {
				if len(shopRentIds) > 0 {
					fmt.Println("got Shop Rent Ids - ", shopRentIds)
					query = append(query, bson.M{"shopRentId": bson.M{"$in": shopRentIds}})
				}
			}
		}

		if len(filter.ReceiptNO) > 0 {
			query = append(query, bson.M{"reciptNo": bson.M{"$in": filter.ReceiptNO}})
		}
		if len(filter.FY) > 0 {
			query = append(query, bson.M{"financialYear.uniqueId": bson.M{"$in": filter.FY}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
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
		if filter.CompletionDate.From != nil {
			sd := time.Date(filter.CompletionDate.From.Year(), filter.CompletionDate.From.Month(), filter.CompletionDate.From.Day(), 0, 0, 0, 0, filter.CompletionDate.From.Location())
			ed := time.Date(filter.CompletionDate.From.Year(), filter.CompletionDate.To.Month(), filter.CompletionDate.To.Day(), 23, 59, 59, 0, filter.CompletionDate.To.Location())
			if filter.CompletionDate.To != nil {
				ed = time.Date(filter.CompletionDate.To.Year(), filter.CompletionDate.To.Month(), filter.CompletionDate.To.Day(), 23, 59, 59, 0, filter.CompletionDate.To.Location())
			}
			query = append(query, bson.M{"$gte": []interface{}{"$completionDate", sd}})
			query = append(query, bson.M{"$lte": []interface{}{"$completionDate", ed}})

		}

	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.RefQueryForShopRentPayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.id", "userName", "ref.collector", "ref.collector")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Filter Mobile Tower Payment =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefShopRentMonthlyPayments
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil

}
