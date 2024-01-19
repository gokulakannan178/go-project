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
func (d *Daos) SaveTradeLicensePayment(ctx *models.Context, mtp *models.TradeLicensePayments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).InsertOne(ctx.CTX, mtp)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

//SaveTradeLicensePaymentFY :""
func (d *Daos) SaveTradeLicensePaymentFYs(ctx *models.Context, mtpfy []models.TradeLicensePaymentsfY) error {
	var insertData []interface{}
	for _, v := range mtpfy {
		insertData = append(insertData, v)
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFY).InsertMany(ctx.CTX, insertData)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

//SaveTradeLicensePaymentBasic :""
func (d *Daos) SaveTradeLicensePaymentBasic(ctx *models.Context, mtpBasic *models.TradeLicensePaymentsBasics) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).InsertOne(ctx.CTX, mtpBasic)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

func (d *Daos) GetSingleTradeLicensePayment(ctx *models.Context, tnxID string) (*models.RefTradeLicensePayments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tnxId": tnxID,
		},
	})
	mainPipeline = append(mainPipeline, d.RefQueryForTradeLicensePayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)

	d.Shared.BsonToJSONPrintTag("Get Single Trade License Payment Query  - ", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rmtps []models.RefTradeLicensePayments
	var rmtp *models.RefTradeLicensePayments
	if err = cursor.All(ctx.CTX, &rmtps); err != nil {
		return nil, err
	}
	if len(rmtps) > 0 {
		rmtp = &rmtps[0]
	}
	return rmtp, nil
}

func (d *Daos) RefQueryForTradeLicensePayment(ctx *models.Context) []bson.M {
	var mainPipeline []bson.M
	//Look for FInancial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSEPAYMENTSFY,
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
			"from": constants.COLLECTIONTRADELICENSEPAYMENTSBASIC,
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "basics.shopRent.tlbtId", "uniqueId", "basics.shopRent.ref.tradeLicenseBusinessType", "basics.shopRent.ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "basics.shopRent.tlctId", "uniqueId", "basics.shopRent.ref.tradeLicenseCategoryType", "basics.shopRent.ref.tradeLicenseCategoryType")...)
	return mainPipeline
}

func (d *Daos) MakeTradeLicensePayment(ctx *models.Context, mmtpr *models.MakeTradeLicensePaymentReq) error {
	query := bson.M{
		"tnxId": mmtpr.TnxID,
	}
	paymentData := bson.M{
		"$set": bson.M{
			"reciptNo":                  d.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSEPAYMENTSRECEIPT),
			"details":                   mmtpr.Details,
			"status":                    mmtpr.Status,
			"creator":                   mmtpr.Creator,
			"completionDate":            mmtpr.CompletionDate,
			"collectionReceived.status": "Pending",
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

// GetTradeLicenseIDsWithOwnerNames : ""
func (d *Daos) GetTradeLicenseIDsWithOwnerNames(ctx *models.Context, shopRentOwnerName string, shopRentOwnerMobileNo string) ([]string, error) {
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
	var shopRentIDsWithOwnerNames []models.TradeLicenseIDsWithOwnerNames
	// var data models.PropertyIDsWithOwnerNames
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &shopRentIDsWithOwnerNames); err != nil {
		return nil, err
	}
	if len(shopRentIDsWithOwnerNames) > 0 {
		return shopRentIDsWithOwnerNames[0].TradeLicenseIDs, nil
	}
	return []string{}, nil
}

// FilterTradeLicensePayment : ""
func (d *Daos) FilterTradeLicensePayment(ctx *models.Context, filter *models.TradeLicensePaymentsFilter, pagination *models.Pagination) ([]models.RefTradeLicensePayments, error) {
	var mainPipeline, query []bson.M
	if filter != nil {
		if len(filter.TradeLicenseID) > 0 {
			query = append(query, bson.M{"tradeLicenseId": bson.M{"$in": filter.TradeLicenseID}})
		}
		if len(filter.MadeAT) > 0 {
			query = append(query, bson.M{"details.madeAt.at": bson.M{"$in": filter.MadeAT}})
		}

		if len(filter.MOP) > 0 {
			query = append(query, bson.M{"details.mop.mode": bson.M{"$in": filter.MOP}})
		}
		if filter.Regex.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.Regex.OwnerName, Options: "xi"}})

		}
		if filter.Regex.TradeLicenseID != "" {
			query = append(query, bson.M{"tradeLicenseId": primitive.Regex{Pattern: filter.Regex.TradeLicenseID, Options: "xi"}})

		}

		if filter.Regex.OwnerName != "" {

			tradeLicenseIds, err := d.GetTradeLicenseIDsWithOwnerNames(ctx, filter.Regex.OwnerName, filter.Regex.OwnerMobile)
			if err != nil {
				log.Println("ERR IN GETING - Shop Rent IDs WithOwner Names " + err.Error())
			} else {
				if len(tradeLicenseIds) > 0 {
					fmt.Println("got Shop Rent Ids - ", tradeLicenseIds)
					query = append(query, bson.M{"shopRentId": bson.M{"$in": tradeLicenseIds}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.RefQueryForTradeLicensePayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.id", "userName", "ref.collector", "ref.collector")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Filter trade License Payment =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefTradeLicensePayments
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil

}

func (d *Daos) GetPayedFinancialYearsOfTradeLicense(ctx *models.Context, traseLisenceID string) ([]string, error) {
	var mainPipeline []bson.M
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{"tradeLicenseId": traseLisenceID,
			"status": bson.M{"$in": []string{constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSEPAYMENTSFY,
			"as":   "completedFys",
			"let":  bson.M{"id": "$tradeLicenseId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{
					"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []interface{}{"$tradeLicenseId", "$$id"}},
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
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

//CalcTradeLicenseDemand :""
func (d *Daos) CalcTradeLicensePaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefTradeLicensePayments, error) {
	d.Shared.BsonToJSONPrintTag("CalcTradeLicenseDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}

	var mtps []models.RefTradeLicensePayments
	if err = cursor.All(context.TODO(), &mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

//CalcTradeLicensePendingPaymens :""
func (d *Daos) CalcTradeLicensePendingPaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefTradeLicensePayments, error) {
	d.Shared.BsonToJSONPrintTag("CalcTradeLicenseDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var mtps []models.RefTradeLicensePayments
	if err = cursor.All(context.TODO(), &mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

//TradeLicenseBouncePayment
func (d *Daos) TradeLicenseBouncePayment(ctx *models.Context, bp *models.BouncePayment) error {
	selector := bson.M{"tnxId": bp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks,
		"bouncedInfo": bson.M{"bouncedActionDate": bp.ActionDate, "bouncedDate": bp.Date, "remark": bp.Remarks, "by": bp.By, "byType": bp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

//
func (d *Daos) GetSingleTradeLicensePaymentWithTxtID(ctx *models.Context, txtID string) (*models.TradeLicensePayments, error) {
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

	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.TradeLicensePayments
	var pp *models.TradeLicensePayments
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

//DateRangeWiseTradeLisencePaymentReport : ""
func (d *Daos) DateRangeWiseTradeLisencePaymentReport(ctx *models.Context, filter *models.DateWiseTradeLicenseReportFilter) (*models.RefDateWiseTradeLicensePaymentReport, error) {
	mainPipeline := []bson.M{}
	// query := []bson.M{}

	resFY, err := d.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("financial year is nil")
	}
	fmt.Println("resFY.Name======>", resFY.Name)
	fmt.Println("resFY.From======>", resFY.From)
	fmt.Println("resFY.To======>", resFY.To)

	// getting the start and end of the month
	var monthsd, monthed *time.Time
	monthsdt := d.Shared.BeginningOfMonth(*filter.Date)
	monthsd = &monthsdt
	monthedt := d.Shared.EndOfMonth(*filter.Date)
	monthed = &monthedt
	t := filter.Date
	monthData := t.Month()
	monthName := monthData.String()

	// getting the start and end of the week
	var weeksd, weeked *time.Time
	weeksdt := d.Shared.StartDayOfWeek(*filter.Date)
	weekedt := d.Shared.EndDayOfWeek(*filter.Date)
	weeksd = &weeksdt
	weeked = &weekedt
	weekStart := weeksd.Format("2006-01-02")
	weekEnd := weeked.Format("2006-01-02")
	fmt.Println("start date of the week======>", weeksd)
	fmt.Println("start date of the week======>", weeked)

	// getting the start and end time of a day
	day := t.Format("2006-01-02")

	var daysd, dayed time.Time
	if filter != nil {
		daysd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, daysd.Location())
		dayed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 0, dayed.Location())
	}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"overall": []bson.M{
			{
				"$match": bson.M{"$and": []bson.M{
					bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
				}}},
			{"$group": bson.M{
				"_id":               "$tradeLicenseId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear.total"},
				"currentCollection": bson.M{"$sum": "$demand.current.total"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			{"$group": bson.M{
				"_id":               nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			}},
		},
		"year": []bson.M{
			{
				"$match": bson.M{"$and": []bson.M{
					bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
					bson.M{"completionDate": bson.M{"$gte": resFY.From,
						"$lte": resFY.To}},
				}}},
			{"$group": bson.M{
				"_id":               "$tradeLicenseId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear.total"},
				"currentCollection": bson.M{"$sum": "$demand.current.total"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			{"$group": bson.M{
				"_id":               nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			}},
		},
		"month": []bson.M{
			{
				"$match": bson.M{"$and": []bson.M{
					bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
					bson.M{"completionDate": bson.M{"$gte": monthsd,
						"$lte": monthed}},
				}}},
			{"$group": bson.M{
				"_id":               "$tradeLicenseId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear.total"},
				"currentCollection": bson.M{"$sum": "$demand.current.total"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			{"$group": bson.M{
				"_id":               nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			}},
		},
		"week": []bson.M{
			{
				"$match": bson.M{"$and": []bson.M{
					bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
					bson.M{"completionDate": bson.M{"$gte": weeksd,
						"$lte": weeked}},
				}}},
			{"$group": bson.M{
				"_id":               "$tradeLicenseId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear.total"},
				"currentCollection": bson.M{"$sum": "$demand.current.total"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			{"$group": bson.M{
				"_id":               nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			}},
		},
		"day": []bson.M{
			{
				"$match": bson.M{"$and": []bson.M{
					bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
					bson.M{"completionDate": bson.M{"$gte": daysd,
						"$lte": dayed}},
				}}},
			{"$group": bson.M{
				"_id":               "$tradeLicenseId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear.total"},
				"currentCollection": bson.M{"$sum": "$demand.current.total"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			{"$group": bson.M{
				"_id":               nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			}},
		},
	}},
		bson.M{"$addFields": bson.M{"overall": bson.M{"$arrayElemAt": []interface{}{"$overall", 0}}}},
		bson.M{"$addFields": bson.M{"year": bson.M{"$arrayElemAt": []interface{}{"$year", 0}}}},
		bson.M{"$addFields": bson.M{"month": bson.M{"$arrayElemAt": []interface{}{"$month", 0}}}},
		bson.M{"$addFields": bson.M{"week": bson.M{"$arrayElemAt": []interface{}{"$week", 0}}}},
		bson.M{"$addFields": bson.M{"day": bson.M{"$arrayElemAt": []interface{}{"$day", 0}}}},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("RefDateWiseTradeLicensePaymentReport query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var reports []models.RefDateWiseTradeLicensePaymentReport
	var report *models.RefDateWiseTradeLicensePaymentReport
	if err = cursor.All(ctx.CTX, &reports); err != nil {
		return nil, err
	}
	if len(reports) > 0 {

		report = &reports[0]
		fmt.Println(report)
		report.Year.FyName = resFY.Name
		report.Month.FyMonth = monthName
		report.Week.FyWeek = weekStart + " " + "to" + " " + weekEnd
		report.Day.FyDay = day

	}

	return report, nil
}

//GetSinglePropertyPaymentDemandBasicWithTxtID :""
func (d *Daos) GetSingleTradelicencePaymentDemandBasicWithTxtID(ctx *models.Context, txtID string) (*models.TradeLicensePaymentsBasics, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdbs []models.TradeLicensePaymentsBasics
	var ppdb *models.TradeLicensePaymentsBasics
	if err = cursor.All(ctx.CTX, &ppdbs); err != nil {
		return nil, err
	}
	if len(ppdbs) > 0 {
		ppdb = &ppdbs[0]
	}
	return ppdb, nil
}

//GetPropertyPaymentDemandFycWithTxtID :""
func (d *Daos) GetTradelicencePaymentDemandFycWithTxtID(ctx *models.Context, txtID string) ([]models.TradeLicensePaymentsfY, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"fy.order": -1}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdfys []models.TradeLicensePaymentsfY
	if err = cursor.All(ctx.CTX, &ppdfys); err != nil {
		return nil, err
	}

	return ppdfys, nil
}

func (d *Daos) UpdateTradelicensePayeenamewithTxnId(ctx *models.Context, TnxID string, name string) error {
	query := bson.M{"tnxId": TnxID}
	update := bson.M{"$set": bson.M{"details.payeeName": name}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
