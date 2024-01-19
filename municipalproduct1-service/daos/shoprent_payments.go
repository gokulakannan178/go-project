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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SavePaymentGateway :""
func (d *Daos) SaveShopRentPayment(ctx *models.Context, mtp *models.ShopRentPayments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).InsertOne(ctx.CTX, mtp)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

//SaveShopRentPaymentFY :""
func (d *Daos) SaveShopRentPaymentFYs(ctx *models.Context, mtpfy []models.ShopRentPaymentsfY) error {
	var insertData []interface{}
	for _, v := range mtpfy {
		insertData = append(insertData, v)
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSFY).InsertMany(ctx.CTX, insertData)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

//SaveShopRentPaymentBasic :""
func (d *Daos) SaveShopRentPaymentBasic(ctx *models.Context, mtpBasic *models.ShopRentPaymentsBasics) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).InsertOne(ctx.CTX, mtpBasic)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

func (d *Daos) GetSingleShopRentPayment(ctx *models.Context, tnxID string) (*models.RefShopRentPayments, error) {
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
	var rmtps []models.RefShopRentPayments
	var rmtp *models.RefShopRentPayments
	if err = cursor.All(ctx.CTX, &rmtps); err != nil {
		return nil, err
	}
	if len(rmtps) > 0 {
		rmtp = &rmtps[0]
	}
	return rmtp, nil
}

func (d *Daos) RefQueryForShopRentPayment(ctx *models.Context) []bson.M {
	var mainPipeline []bson.M
	//Look for FInancial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONSHOPRENTPAYMENTSFY,
			"as":   "fys",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
				}}}},
			},
		},
	})
	//Look for  basic
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONSHOPRENTPAYMENTSBASIC,
			"as":   "basics",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
				}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"basics": bson.M{"$arrayElemAt": []interface{}{"$basics", 0}}},
	})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "basics.shopRent.shopCategoryId", "uniqueId", "basics.shopRent.ref.shopRentShopCategory", "basics.shopRent.ref.shopRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "basics.shopRent.shopSubCategoryId", "uniqueId", "basics.shopRent.ref.shopRentShopSubCategory", "basics.shopRent.ref.shopRentShopSubCategory")...)
	return mainPipeline
}

func (d *Daos) MakeShopRentPayment(ctx *models.Context, mmtpr *models.MakeShopRentPaymentReq) error {
	query := bson.M{
		"tnxId": mmtpr.TnxID,
	}
	paymentData := bson.M{
		"$set": bson.M{
			"reciptNo":                  d.GetUniqueID(ctx, constants.COLLECTIONSHOPRENTPAYMENTSRECEIPT),
			"details":                   mmtpr.Details,
			"status":                    mmtpr.Status,
			"creator":                   mmtpr.Creator,
			"completionDate":            mmtpr.CompletionDate,
			"collectionReceived.status": "Pending",
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

// GetShopRentIDsWithOwnerNames : ""
func (d *Daos) GetShopRentIDsWithOwnerNames(ctx *models.Context, shopRentOwnerName string, shopRentOwnerMobileNo string) ([]string, error) {
	mainPipeline := []bson.M{}
	query := bson.M{}
	if shopRentOwnerName == "" && shopRentOwnerMobileNo == "" {
		return []string{}, nil
	}
	if shopRentOwnerName != "" {
		query["shopRent.ownerName"] = primitive.Regex{Pattern: shopRentOwnerName, Options: "sxi"}

	}
	if shopRentOwnerMobileNo != "" {
		query["shopRent.mobileNo"] = primitive.Regex{Pattern: shopRentOwnerMobileNo, Options: "sxi"}

	}
	mainPipeline = append(mainPipeline, bson.M{"$match": query})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": "$shopRentId"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": nil, "shopRentIds": bson.M{"$push": "$_id"}},
	})
	var shopRentIDsWithOwnerNames []models.ShopRentIDsWithOwnerNames
	// var data models.PropertyIDsWithOwnerNames
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &shopRentIDsWithOwnerNames); err != nil {
		return nil, err
	}
	if len(shopRentIDsWithOwnerNames) > 0 {
		return shopRentIDsWithOwnerNames[0].ShopRentIDs, nil
	}
	return []string{}, nil
}

// GetShopRentIDsWithOwnerNames : ""
func (d *Daos) GetShopRentIDsWithMobileNos(ctx *models.Context, shopRentOwnerName string, shopRentOwnerMobileNo string) ([]string, error) {
	mainPipeline := []bson.M{}
	query := bson.M{}
	if shopRentOwnerName == "" && shopRentOwnerMobileNo == "" {
		return []string{}, nil
	}
	if shopRentOwnerName != "" {
		query["shopRent.ownerName"] = primitive.Regex{Pattern: shopRentOwnerName, Options: "sxi"}

	}
	if shopRentOwnerMobileNo != "" {
		query["shopRent.mobileNo"] = primitive.Regex{Pattern: shopRentOwnerMobileNo, Options: "sxi"}

	}
	mainPipeline = append(mainPipeline, bson.M{"$match": query})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": "$shopRentId"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": nil, "shopRentIds": bson.M{"$push": "$_id"}},
	})
	var shopRentIDsWithMobileNos []models.ShopRentIDsWithMobileNos
	// var data models.PropertyIDsWithOwnerNames
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &shopRentIDsWithMobileNos); err != nil {
		return nil, err
	}
	if len(shopRentIDsWithMobileNos) > 0 {
		return shopRentIDsWithMobileNos[0].ShopRentIDs, nil
	}
	return []string{}, nil
}

// FilterShopRentPayment : ""
func (d *Daos) FilterShopRentPayment(ctx *models.Context, filter *models.ShopRentPaymentsFilter, pagination *models.Pagination) ([]models.RefShopRentPayments, error) {
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
		if filter.Regex.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.Regex.OwnerName, Options: "xi"}})

		}
		if filter.Regex.ShopRentID != "" {
			query = append(query, bson.M{"shopRentId": primitive.Regex{Pattern: filter.Regex.ShopRentID, Options: "xi"}})

		}
		if filter.Regex.ReceiptNO != "" {
			query = append(query, bson.M{"reciptNo": primitive.Regex{Pattern: filter.Regex.ReceiptNO, Options: "xi"}})

		}

		if filter.Regex.OwnerName != "" {

			shopRentIds, err := d.GetShopRentIDsWithOwnerNames(ctx, filter.Regex.OwnerName, filter.Regex.OwnerMobile)
			if err != nil {
				log.Println("ERR IN GETING - Shop Rent IDs WithOwner Names " + err.Error())
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

		if filter.CompletionDate != nil {
			//var sd,ed time.Time
			if filter.CompletionDate.From != nil {
				sd := time.Date(filter.CompletionDate.From.Year(), filter.CompletionDate.From.Month(), filter.CompletionDate.From.Day(), 0, 0, 0, 0, filter.CompletionDate.From.Location())
				ed := time.Date(filter.CompletionDate.From.Year(), filter.CompletionDate.From.Month(), filter.CompletionDate.From.Day(), 23, 59, 59, 0, filter.CompletionDate.From.Location())
				if filter.CompletionDate.To != nil {
					ed = time.Date(filter.CompletionDate.To.Year(), filter.CompletionDate.To.Month(), filter.CompletionDate.To.Day(), 23, 59, 59, 0, filter.CompletionDate.To.Location())
				}
				query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})

			}
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
	var payments []models.RefShopRentPayments
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil

}

func (d *Daos) GetPayedFinancialYearsOfShoprent(ctx *models.Context, shopID string) ([]string, error) {
	var mainPipeline []bson.M
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{"shopRentId": shopID,
			"status": bson.M{"$in": []string{constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONSHOPRENTPAYMENTSFY,
			"as":   "completedFys",
			"let":  bson.M{"id": "$shopRentId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{
					"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []interface{}{"$shopRentId", "$$id"}},
						bson.M{"$in": []interface{}{"$status", []string{constants.SHOPRENTPAYMENTSTATUSCOMPLETED}}},
					}},
				}},
				bson.M{
					"$group": bson.M{"_id": "$fy.uniqueId"},
				},
				bson.M{
					"$group": bson.M{"_id": nil, "data": bson.M{"$push": "$_id"}},
				},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"completedFys": bson.M{"$arrayElemAt": []interface{}{"$completedFys", 0}},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Calculating payed Fys =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	type sample struct {
		CompletedFys struct {
			Data []string `json:"data" bson:"data,omitempty"`
		} `json:"completedFys" bson:"completedFys,omitempty"`
	}
	var payments []sample

	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	if len(payments) > 0 {
		return payments[0].CompletedFys.Data, nil
	}
	return []string{}, nil
}

//CalcShopRentDemand :""
func (d *Daos) CalcShopRentPaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefShopRentPayments, error) {
	d.Shared.BsonToJSONPrintTag("CalcShopRentDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}

	var mtps []models.RefShopRentPayments
	if err = cursor.All(context.TODO(), &mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

//CalcShopRentPendingPaymens :""
func (d *Daos) CalcShopRentPendingPaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefShopRentPayments, error) {
	d.Shared.BsonToJSONPrintTag("CalcShopRentDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var mtps []models.RefShopRentPayments
	if err = cursor.All(context.TODO(), &mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

func (d *Daos) ShoprentBouncePayment(ctx *models.Context, bp *models.BouncePayment) error {
	selector := bson.M{"tnxId": bp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks,
		"bouncedInfo": bson.M{"bouncedActionDate": bp.ActionDate, "bouncedDate": bp.Date, "remark": bp.Remarks, "by": bp.By, "byType": bp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

//GetSinglePropertyPaymentWithTxtID :""
func (d *Daos) GetSingleShoprentPaymentWithTxtID(ctx *models.Context, txtID string) (*models.ShopRentPayments, error) {
	mainPipeline := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})

	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{

	// 	"from": constants.COLLECTIONs,
	// 	"as":   "ref.partPayments",
	// 	"let":  bson.M{"tnxId": "$tnxId"},
	// 	"pipeline": []bson.M{
	// 		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 			{"$eq": []string{"$tnxId", "$$tnxId"}},
	// 			{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
	// 		}}}},
	// 	},
	// }})

	d.Shared.BsonToJSONPrintTag("property payment query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.ShopRentPayments
	var pp *models.ShopRentPayments
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

func (d *Daos) UpdateShoprentPayeenamewithTxnId(ctx *models.Context, TnxID string, name string) error {
	query := bson.M{"tnxId": TnxID}
	update := bson.M{"$set": bson.M{"details.payeeName": name}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
