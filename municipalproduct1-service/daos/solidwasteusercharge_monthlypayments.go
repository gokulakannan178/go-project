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
func (d *Daos) SaveSolidWasteUserChargeMonthlyPayment(ctx *models.Context, mtp *models.SolidWasteUserChargeMonthlyPayments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).InsertOne(ctx.CTX, mtp)
	d.Shared.BsonToJSONPrintTag("Solid Waste User Charge payment resp - ", res)
	return err
}

//SaveSolidWasteUserChargePaymentFY :""
func (d *Daos) SaveSolidWasteUserChargeMonthlyPaymentFYs(ctx *models.Context, mtpfy []models.SolidWasteChargeMonthlyPaymentsfY) error {
	var insertData []interface{}
	for _, v := range mtpfy {
		insertData = append(insertData, v)
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSFY).InsertMany(ctx.CTX, insertData)
	d.Shared.BsonToJSONPrintTag("Solid Waste User Charge payment resp - ", res)
	return err
}

// SaveSolidWasteUserChargePaymentBasic :""
func (d *Daos) SaveSolidWasteUserChargePaymentBasic(ctx *models.Context, mtpBasic *models.SolidWasteChargeMonthlyPaymentsBasics) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC).InsertOne(ctx.CTX, mtpBasic)
	d.Shared.BsonToJSONPrintTag("Solid Waste User Charge payment resp - ", res)
	return err
}

// GetSingleSolidWasteUserChargePayment : ""
func (d *Daos) GetSingleSolidWasteUserChargePayment(ctx *models.Context, tnxID string) (*models.RefSolidWasteChargeMonthlyPayments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tnxId": tnxID,
		},
	})
	mainPipeline = append(mainPipeline, d.RefQueryForSolidWasteUserChargePayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)

	d.Shared.BsonToJSONPrintTag("Get Single Solid Waste User Charge Payment Query  - ", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rmtps []models.RefSolidWasteChargeMonthlyPayments
	var rmtp *models.RefSolidWasteChargeMonthlyPayments
	if err = cursor.All(ctx.CTX, &rmtps); err != nil {
		return nil, err
	}
	if len(rmtps) > 0 {
		rmtp = &rmtps[0]
	}
	return rmtp, nil
}

func (d *Daos) RefQueryForSolidWasteUserChargePayment(ctx *models.Context) []bson.M {
	var mainPipeline []bson.M
	//Look for FInancial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSFY,
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
			"from": constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC,
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "basics.solidWasteUserCharge.categoryId", "uniqueId", "basics.solidWasteUserCharge.ref.category", "basics.solidWasteUserCharge.ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY, "basics.solidWasteUserCharge.subCategoryId", "uniqueId", "basics.solidWasteUserCharge.ref.subCategory", "basics.solidWasteUserCharge.ref.subCategory")...)
	return mainPipeline
}

// MakeSolidWasteUserChargePayment : ""
func (d *Daos) MakeSolidWasteUserChargePayment(ctx *models.Context, mmtpr *models.MakeSolidWasteUserChargePaymentReq) error {
	query := bson.M{
		"tnxId": mmtpr.TnxID,
	}
	paymentData := bson.M{
		"$set": bson.M{
			"reciptNo":       d.GetUniqueID(ctx, constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSRECEIPT),
			"details":        mmtpr.Details,
			"status":         mmtpr.Status,
			"creator":        mmtpr.Creator,
			"completionDate": mmtpr.CompletionDate,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid Waste User Charge payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid Waste User Charge payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid Waste User Charge payment basic resp - ", res)
	return nil
}

// GetSolidWasteUserChargeIDsWithOwnerNames : ""
func (d *Daos) GetSolidWasteUserChargeIDsWithOwnerNames(ctx *models.Context, solidWasteUserChargeOwnerName string, solidWasteUserChargeOwnerNameMobileNo string) ([]string, error) {
	mainPipeline := []bson.M{}
	query := bson.M{}
	if solidWasteUserChargeOwnerName == "" && solidWasteUserChargeOwnerNameMobileNo == "" {
		return []string{}, nil
	}
	if solidWasteUserChargeOwnerName != "" {
		query["solidWasteUserCharge.ownerName"] = primitive.Regex{Pattern: solidWasteUserChargeOwnerName, Options: "sxi"}

	}
	if solidWasteUserChargeOwnerNameMobileNo != "" {
		query["solidWasteUserCharge.mobileNo"] = primitive.Regex{Pattern: solidWasteUserChargeOwnerNameMobileNo, Options: "sxi"}

	}
	mainPipeline = append(mainPipeline, bson.M{"$match": query})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": "$solidWasteUserChargeId"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": nil, "solidWasteUserChargeIds": bson.M{"$push": "$_id"}},
	})
	var solidWasteUserChargeIDsWithOwnerNames []models.SolidWasteUserChargeIDsWithOwnerNames
	// var data models.PropertyIDsWithOwnerNames
	d.Shared.BsonToJSONPrintTag("solid waste user charge query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &solidWasteUserChargeIDsWithOwnerNames); err != nil {
		return nil, err
	}
	if len(solidWasteUserChargeIDsWithOwnerNames) > 0 {
		return solidWasteUserChargeIDsWithOwnerNames[0].SolidWasteUserChargeIDs, nil
	}
	return []string{}, nil
}

// FilterSolidWasteUserChargePayment : ""
func (d *Daos) FilterSolidWasteUserChargePayment(ctx *models.Context, filter *models.SolidWasteUserChargePaymentsFilter, pagination *models.Pagination) ([]models.RefSolidWasteChargeMonthlyPayments, error) {
	var mainPipeline, query []bson.M
	if filter != nil {
		if len(filter.SolidWasteUserChargeID) > 0 {
			query = append(query, bson.M{"solidWasteUserChargeId": bson.M{"$in": filter.SolidWasteUserChargeID}})
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
		if filter.Regex.SolidWasteUserChargeID != "" {
			query = append(query, bson.M{"solidWasteUserChargeId": primitive.Regex{Pattern: filter.Regex.SolidWasteUserChargeID, Options: "xi"}})

		}
		if filter.Regex.ReceiptNO != "" {
			query = append(query, bson.M{"reciptNo": primitive.Regex{Pattern: filter.Regex.ReceiptNO, Options: "xi"}})

		}

		if filter.Regex.OwnerName != "" {

			solidWasteUserChargeIds, err := d.GetSolidWasteUserChargeIDsWithOwnerNames(ctx, filter.Regex.OwnerName, filter.Regex.OwnerMobile)
			if err != nil {
				log.Println("ERR IN GETING - Shop Rent IDs WithOwner Names " + err.Error())
			} else {
				if len(solidWasteUserChargeIds) > 0 {
					fmt.Println("got solid waste user charge Ids - ", solidWasteUserChargeIds)
					query = append(query, bson.M{"solidWasteUserChargeId": bson.M{"$in": solidWasteUserChargeIds}})
				}
			}
		}

		if len(filter.ReceiptNo) > 0 {
			query = append(query, bson.M{"reciptNo": bson.M{"$in": filter.ReceiptNo}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.RefQueryForSolidWasteUserChargePayment(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)

	d.Shared.BsonToJSONPrintTag("solid Waste User Charge Payment Query - ", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefSolidWasteChargeMonthlyPayments
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil

}

// VerifySolidWasteUserChargePayment : ""
func (d *Daos) VerifySolidWasteUserChargePayment(ctx *models.Context, action *models.MakeSolidWasteUserChargePaymentsAction) (string, error) {
	fmt.Println("action.TnxID:", action.TnxID)
	payment, err := d.GetSingleSolidWasteUserChargePayment(ctx, action.TnxID)
	if err != nil {
		return "", err
	}

	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.SolidWasteUserChargePaymentsAction,
			"status":       constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}

	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment basic resp - ", res)
	return payment.SolidWasteUserChargeID, nil
}

// NotVerifySolidWasteUserChargePayment : ""
func (d *Daos) NotVerifySolidWasteUserChargePayment(ctx *models.Context, action *models.MakeSolidWasteUserChargePaymentsAction) (string, error) {
	payment, err := d.GetSingleSolidWasteUserChargePayment(ctx, action.TnxID)
	if err != nil {
		return "", err
	}
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.SolidWasteUserChargePaymentsAction,
			"status":       constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment basic resp - ", res)
	return payment.SolidWasteUserChargeID, nil
}

// RejectSolidWasteUserChargePayment : ""
func (d *Daos) RejectSolidWasteUserChargePayment(ctx *models.Context, action *models.MakeSolidWasteUserChargePaymentsAction) (string, error) {
	payment, err := d.GetSingleSolidWasteUserChargePayment(ctx, action.TnxID)
	if err != nil {
		return "", err
	}
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.SolidWasteUserChargePaymentsAction,
			"status":       constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSREJECTED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Solid waste user charge payment basic resp - ", res)
	return payment.SolidWasteUserChargeID, nil
}
