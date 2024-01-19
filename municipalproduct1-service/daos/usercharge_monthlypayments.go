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

// SavePaymentGateway :""
func (d *Daos) SaveUserChargeMonthlyPayment(ctx *models.Context, mtp *models.UserChargeMonthlyPayments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).InsertOne(ctx.CTX, mtp)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment resp - ", res)
	return err
}

// SaveUserChargePaymentFY :""
func (d *Daos) SaveUserMonthlyChargePaymentFYs(ctx *models.Context, mtpfy []models.UserChargetMonthlyPaymentsfY) error {
	var insertData []interface{}
	for _, v := range mtpfy {
		insertData = append(insertData, v)
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTSFY).InsertMany(ctx.CTX, insertData)
	d.Shared.BsonToJSONPrintTag("Shop Rentpayment Fys resp - ", res)
	return err
}

// SaveUserChargePaymentBasic : ""
func (d *Daos) SaveUserChargePaymentBasic(ctx *models.Context, mtbasics *models.UserChargePaymentsBasics) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTBASICS).InsertOne(ctx.CTX, mtbasics)
	d.Shared.BsonToJSONPrintTag("Shop RentpaymentBasics  resp - ", res)
	return err
}
func (d *Daos) GetSingleUserChargeMonthlyPayment(ctx *models.Context, tnxID string) (*models.RefUserChargeMonthlyPayments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tnxId": tnxID,
		},
	})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGEPAYMENTBASICS, "txnId", "txnId", "basics", "basics")...)

	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// 	"from":         "userchargepaymentbasics",
	// 	"as":           "basics",
	// 	"localField":   "txnId",
	// 	"foreignField": "txnId",
	// }})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"basics": bson.M{"$arrayElemAt": []interface{}{"$basics", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"as": "fys",
		"from": "userchargepaymentfy",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{
				"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
				}}},
			},
			bson.M{"$sort": bson.M{"month.fyOrder": 1}},
			bson.M{"$group": bson.M{
				"_id":    "$fy.uniqueId",
				"fy":     bson.M{"$first": bson.M{"name": "$fy.name", "fyId": "$fy.uniqueId"}},
				"months": bson.M{"$push": "$month"},
			}},
		}}})
	//	mainPipeline = append(mainPipeline, d.RefQueryForUserChargePayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "propertyId", "uniqueId", "ref.property", "ref.property")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "propertyId", "propertyId", "ref.propertyowner", "ref.propertyowner")...)

	d.Shared.BsonToJSONPrintTag("Get Single Shop Rent Payment Query  - ", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rmtps []models.RefUserChargeMonthlyPayments
	var rmtp *models.RefUserChargeMonthlyPayments
	if err = cursor.All(ctx.CTX, &rmtps); err != nil {
		return nil, err
	}
	if len(rmtps) > 0 {
		rmtp = &rmtps[0]
	}
	return rmtp, nil
}

// FilterUserChargePayment : ""
func (d *Daos) FilterUserChargeMonthlyPayment(ctx *models.Context, filter *models.UserChargeMonthlyPaymentsFilter, pagination *models.Pagination) ([]models.RefUserChargeMonthlyPayments, error) {
	var mainPipeline, query []bson.M
	if filter != nil {
		if len(filter.UserChargeID) > 0 {
			query = append(query, bson.M{"UserChargeId": bson.M{"$in": filter.UserChargeID}})
		}

		if len(filter.MadeAT) > 0 {
			query = append(query, bson.M{"details.madeAt.at": bson.M{"$in": filter.MadeAT}})
		}
		if len(filter.Collector) > 0 {
			query = append(query, bson.M{"details.collector.by": bson.M{"$in": filter.Collector}})
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
		if filter.SearchBox.UserChargeID != "" {
			query = append(query, bson.M{"UserChargeId": primitive.Regex{Pattern: filter.SearchBox.UserChargeID, Options: "xi"}})

		}

		// if filter.SearchBox.OwnerName != "" {

		// 	UserChargeIds, err := d.GetUserChargeIDsWithOwnerNames(ctx, filter.SearchBox.OwnerName, filter.SearchBox.OwnerMobile)
		// 	if err != nil {
		// 		log.Println("ERR IN GETING - Shop Rent IDs WithOwner Names " + err.Error())
		// 	} else {
		// 		if len(UserChargeIds) > 0 {
		// 			fmt.Println("got Shop Rent Ids - ", UserChargeIds)
		// 			query = append(query, bson.M{"UserChargeId": bson.M{"$in": UserChargeIds}})
		// 		}
		// 	}
		// }

		// if filter.SearchBox.OwnerMobile != "" {

		// 	UserChargeIds, err := d.GetUserChargeIDsWithMobileNos(ctx, filter.SearchBox.OwnerName, filter.SearchBox.OwnerMobile)
		// 	if err != nil {
		// 		log.Println("ERR IN GETING - Shop Rent IDs With Mobile Numbers " + err.Error())
		// 	} else {
		// 		if len(UserChargeIds) > 0 {
		// 			fmt.Println("got Shop Rent Ids - ", UserChargeIds)
		// 			query = append(query, bson.M{"UserChargeId": bson.M{"$in": UserChargeIds}})
		// 		}
		// 	}
		// }

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
			query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})
			// query = append(query, bson.M{"$gte": []interface{}{"$completionDate", sd}})
			// query = append(query, bson.M{"$lte": []interface{}{"$completionDate", ed}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGEPAYMENTBASICS, "txnId", "txnId", "basics", "basics")...)

	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// 	"from":         "userchargepaymentbasics",
	// 	"as":           "basics",
	// 	"localField":   "txnId",
	// 	"foreignField": "txnId",
	// }})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"basics": bson.M{"$arrayElemAt": []interface{}{"$basics", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"as": "fys",
		"from": "userchargepaymentfy",
		"let":  bson.M{"txnId": "$txnId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{
				"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$txnId", "$$txnId"}},
				}}}},

			bson.M{"$group": bson.M{
				"_id":    "$fy.uniqueId",
				"fy":     bson.M{"$first": bson.M{"name": "$fy.name", "fyId": "$fy.uniqueId"}},
				"months": bson.M{"$push": "$month"},
			}},
		}}})
	//mainPipeline = append(mainPipeline, d.RefQueryForUserChargePayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "propertyId", "uniqueId", "ref.property", "ref.property")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "propertyId", "propertyId", "ref.propertyowner", "ref.propertyowner")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Filter Mobile Tower Payment =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefUserChargeMonthlyPayments
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil

}

func (d *Daos) MakeUserChargePayment(ctx *models.Context, mmtpr *models.MakeUserChargePaymentReq) error {
	query := bson.M{
		"tnxId": mmtpr.TnxID,
	}
	paymentData := bson.M{
		"$set": bson.M{
			"reciptNo":                  d.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGEPAYMENTSRECEIPT),
			"details":                   mmtpr.Details,
			"status":                    mmtpr.Status,
			"creator":                   mmtpr.Creator,
			"completionDate":            mmtpr.CompletionDate,
			"collectionReceived.status": "Pending",
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTBASICS).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

// BouncePayment : ""
func (d *Daos) UserChargeBouncePayment(ctx *models.Context, bp *models.BouncePayment) error {
	selector := bson.M{"tnxId": bp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks,
		"bouncedInfo": bson.M{"bouncedActionDate": bp.ActionDate, "bouncedDate": bp.Date, "remark": bp.Remarks, "by": bp.By, "byType": bp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTBASICS).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// VerifyPayment : ""
func (d *Daos) UserChargVerifyPayment(ctx *models.Context, vp *models.VerifyPayment) error {
	selector := bson.M{"tnxId": vp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTCOMPLETED, "remark": ""}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTCOMPLETED, "remark": vp.Remarks,
		"verifiedInfo": bson.M{"actionDate": vp.ActionDate, "actualDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTBASICS).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// NotVerifiedPayment : ""
func (d *Daos) UserChargNotVerifiedPayment(ctx *models.Context, vp *models.NotVerifiedPayment) error {
	selector := bson.M{"tnxId": vp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTNOTVERIFIED, "remark": vp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTNOTVERIFIED, "remark": vp.Remarks,
		"notVerifiedInfo": bson.M{"actionDate": vp.ActionDate, "actualDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTBASICS).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// RejectPayment : ""
func (d *Daos) UserChargRejectPayment(ctx *models.Context, rp *models.RejectPayment) error {
	selector := bson.M{"tnxId": rp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENREJECTED, "rejectedRemark": rp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENREJECTED, "remark": rp.Remarks,
		"rejectedInfo": bson.M{"actionDate": rp.ActionDate, "actualDate": rp.Date, "remark": rp.Remarks, "by": rp.By, "byType": rp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTBASICS).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

func (d *Daos) GetSingleUserChargePaymentWithTxtID(ctx *models.Context, txtID string) (*models.UserChargePayments, error) {
	mainPipeline := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})

	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{

	// 	"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
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

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.UserChargePayments
	var pp *models.UserChargePayments
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

//GetSinglePropertyPaymentDemandBasicWithTxtID :""
func (d *Daos) GetSingleUserChargePaymentBasicWithTxtID(ctx *models.Context, txtID string) (*models.UserChargePaymentsBasics, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTBASICS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdbs []models.UserChargePaymentsBasics
	var ppdb *models.UserChargePaymentsBasics
	if err = cursor.All(ctx.CTX, &ppdbs); err != nil {
		return nil, err
	}
	if len(ppdbs) > 0 {
		ppdb = &ppdbs[0]
	}
	return ppdb, nil
}

func (d *Daos) GetUserChargePaymentFycWithTxtID(ctx *models.Context, txtID string) ([]models.UserChargePaymentsfY, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"fy.order": -1}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTSFY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdfys []models.UserChargePaymentsfY
	if err = cursor.All(ctx.CTX, &ppdfys); err != nil {
		return nil, err
	}

	return ppdfys, nil
}

//DateRangeWiseTradeLisencePaymentReport : ""
func (d *Daos) DateRangeWiseUserchargePaymentReport(ctx *models.Context, filter *models.DateWiseUserchargeReportFilter) (*models.RefDateWiseTradeLicensePaymentReport, error) {
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
				"_id":               "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrearTotal"},
				"currentCollection": bson.M{"$sum": "$demand.currentTotal"},
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
				"_id":               "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrearTotal"},
				"currentCollection": bson.M{"$sum": "$demand.currentTotal"},
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
				"_id":               "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrearTotal"},
				"currentCollection": bson.M{"$sum": "$demand.currentTotal"},
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
				"_id":               "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrearTotal"},
				"currentCollection": bson.M{"$sum": "$demand.currentTotal"},
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
				"_id":               "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrearTotal"},
				"currentCollection": bson.M{"$sum": "$demand.currentTotal"},
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
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

func (d *Daos) UpdateUserchargePayeenamewithTxnId(ctx *models.Context, TnxID string, name string) error {
	query := bson.M{"tnxId": TnxID}
	update := bson.M{"$set": bson.M{"details.payeeName": name}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
