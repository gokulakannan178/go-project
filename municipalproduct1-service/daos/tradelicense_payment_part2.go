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

// SaveTradeLicensePaymentsPart2 : ""
func (d *Daos) SaveTradeLicensePaymentsPart2(ctx *models.Context, tlpp2 *models.TradeLicensePaymentsPart2) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).InsertOne(ctx.CTX, tlpp2)
	d.Shared.BsonToJSONPrintTag("trade license payment resp - ", res)
	return err
}

// SaveTradeLicensePaymentsBasicsPart2 : ""
func (d *Daos) SaveTradeLicensePaymentsBasicsPart2(ctx *models.Context, tlpbp2 *models.TradeLicensePaymentsBasicsPart2) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASICPART2).InsertOne(ctx.CTX, tlpbp2)
	d.Shared.BsonToJSONPrintTag("trade license payment basic resp - ", res)
	return err
}

// SaveTradeLicensePaymentsFYsPart2 : ""
func (d *Daos) SaveTradeLicensePaymentsFYsPart2(ctx *models.Context, tlpfyp2 []models.TradeLicensePaymentsfY) error {
	var insertData []interface{}
	for _, v := range tlpfyp2 {
		insertData = append(insertData, v)
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2).InsertMany(ctx.CTX, insertData)
	d.Shared.BsonToJSONPrintTag("trade license payment fys resp - ", res)
	return err
}

// GetSingleTradeLicensePaymentPart2 : ""
func (d *Daos) GetSingleTradeLicensePaymentPart2(ctx *models.Context, tnxID string) (*models.RefTradeLicensePayments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tnxId": tnxID,
		},
	})
	mainPipeline = append(mainPipeline, d.RefQueryForTradeLicensePaymentPart2(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)

	d.Shared.BsonToJSONPrintTag("Get Single Trade Rent Payment Part2 Query  - ", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).Aggregate(ctx.CTX, mainPipeline, nil)
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

// RefQueryForTradeLicensePaymentPart2 : ""
func (d *Daos) RefQueryForTradeLicensePaymentPart2(ctx *models.Context) []bson.M {
	var mainPipeline []bson.M
	//Look for FInancial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2,
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
			"from": constants.COLLECTIONTRADELICENSEPAYMENTSBASICPART2,
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
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "basics.tradeLicense.ref.tradeLicenseBusinessType", "basics.shopRent.ref.tradeLicense")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "basics.tradeLicense.ref.tradeLicenseCategoryType", "basics.tradeLicense.ref.tradeLicenseCategoryType")...)
	return mainPipeline
}

// MakeTradeLicensePaymentPart2 : ""
func (d *Daos) MakeTradeLicensePaymentPart2(ctx *models.Context, mtlprp2 *models.MakeTradeLicensePaymentReqPart2) error {
	query := bson.M{
		"tnxId": mtlprp2.TnxID,
	}
	paymentData := bson.M{
		"$set": bson.M{
			"reciptNo":       d.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSEPAYMENTSRECEIPTPART2),
			"details":        mtlprp2.Details,
			"status":         mtlprp2.Status,
			"creator":        mtlprp2.Creator,
			"completionDate": mtlprp2.CompletionDate,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating trade license payment part2 - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("trade license payment part2 resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": mtlprp2.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("trade license payment fys part2 resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": mtlprp2.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASICPART2).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating tradelicense payment part2 - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("trade license payment basic part2 resp - ", res)
	return nil
}

// FilterTradeLicensePaymentPart2 : ""
func (d *Daos) FilterTradeLicensePaymentPart2(ctx *models.Context, filter *models.TradeLicensePaymentsFilterPart2, pagination *models.Pagination) ([]models.RefTradeLicensePaymentsPart2, error) {
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
				log.Println("ERR IN GETING - trade license IDs WithOwner Names " + err.Error())
			} else {
				if len(tradeLicenseIds) > 0 {
					fmt.Println("got trade license Ids - ", tradeLicenseIds)
					query = append(query, bson.M{"tradeLicenseId": bson.M{"$in": tradeLicenseIds}})
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
			query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.RefQueryForTradeLicensePaymentPart2(ctx)...)
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Filter trade License Payment Part2=>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefTradeLicensePaymentsPart2
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil

}

// GetSingleTradeLicenseV2Part2 : ""
func (d *Daos) GetSingleTradeLicenseV2Part2(ctx *models.Context, UniqueID string) (*models.RefTradeLicense, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSEPAYMENTSPART2,
			"as":   "ref.payments",
			"let":  bson.M{"uniqueId": "$uniqueId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$tradeLicenseId", "$$uniqueId"}},
					{"$eq": []string{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
				}}},
			},
				bson.M{
					"$lookup": bson.M{
						"from": constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2,
						"as":   "fys",
						"let":  bson.M{"tnxId": "$tnxId"},
						"pipeline": []bson.M{bson.M{
							"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$tnxId", "$$tnxId"}}}}}, // {"$eq": []string{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},

						},
							bson.M{
								"$lookup": bson.M{
									"from": constants.COLLECTIONTRADELICENSEPAYMENTSBASICPART2,
									"as":   "basics",
									"let":  bson.M{"tnxId": "$tnxId"},
									"pipeline": []bson.M{bson.M{
										"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$tnxId", "$$tnxId"}}}}}, // {"$eq": []string{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},

									},
									}}},
							bson.M{"$addFields": bson.M{"basics": bson.M{"$arrayElemAt": []interface{}{"$basics", 0}}}},
						}}},
			}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSERATEMASTER,
			"as":   "ref.marketRate",
			"let":  bson.M{"tlctId": "$tlctId", "tlbtId": "$tlbtId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$tlctId", "$$tlctId"}},
					{"$eq": []string{"$tlbtId", "$$tlbtId"}},
					{"$eq": []string{"$status", constants.TRADELICENSERATEMASTERSTATUSACTIVE}},
				}}},
			},
			},
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.marketRate": bson.M{"$arrayElemAt": []interface{}{"$ref.marketRate", 0}}}})

	d.Shared.BsonToJSONPrintTag("trade license v2 part2 query =>", mainPipeline)

	// LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "ref.tradeLicenseCategoryType", "ref.tradeLicenseCategoryType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefTradeLicense
	var tower *models.RefTradeLicense
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateExpiryDatePart2 :""
func (d *Daos) UpdateExpiryDatePart2(ctx *models.Context, uniqueID string, expiryDate *time.Time, fromDate *time.Time, status string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"licenseExpiryDate": expiryDate, "status": status, "licenseDate": fromDate}}
	d.Shared.BsonToJSONPrintTag("query==>", query)
	d.Shared.BsonToJSONPrintTag("update==>", update)
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return nil
}

// VerifyTradeLicensePaymentPart2 : ""
func (d *Daos) VerifyTradeLicensePaymentPart2(ctx *models.Context, action *models.MakeTradeLicensePaymentsActionPart2) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.TradeLicensePaymentsAction,
			"status":       constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating trade license payment part2 - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment part2 resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating trade license payment part2 - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment fys part2 resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASICPART2).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating trade license payment part2 - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment basic part2 resp - ", res)
	return nil
}

// NotVerifyTradeLicensePaymentPart2 : ""
func (d *Daos) NotVerifyTradeLicensePaymentPart2(ctx *models.Context, action *models.MakeTradeLicensePaymentsActionPart2) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.TradeLicensePaymentsAction,
			"status":       constants.TRADELICENSEPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASICPART2).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment basic resp - ", res)
	return nil
}

// RejectTradeLicensePaymentPart2 : ""
func (d *Daos) RejectTradeLicensePaymentPart2(ctx *models.Context, action *models.MakeTradeLicensePaymentsActionPart2) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.TradeLicensePaymentsAction,
			"status":       constants.TRADELICENSEPAYMENRSTATUSREJECTED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASICPART2).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment basic resp - ", res)
	return nil
}

//BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2 : ""
func (d *Daos) BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2(ctx *models.Context, rbtlulp2 *models.RefBasicTradeLicenseUpdateLogV2Part2) ([]models.RefTradeLicensePaymentsPart2, error) {
	//get current Financial year

	cfy, err := d.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	if cfy == nil {
		return nil, errors.New("current financial year is nil")
	}
	sd := time.Date(cfy.From.Year(), cfy.From.Month(), cfy.From.Day(), 0, 0, 0, 0, cfy.From.Location())
	ed := time.Date(cfy.To.Year(), cfy.To.Month(), cfy.To.Day(), 23, 59, 59, 0, cfy.To.Location())
	fmt.Println("sd ===>", sd)
	fmt.Println("ed ===>", ed)
	tradeLicencePaymentFindQuery := bson.M{
		"status":         constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		"tradeLicenseId": rbtlulp2.TradeLicenseID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("tradeLicense payment part2 query =>", tradeLicencePaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSPART2).Find(ctx.CTX, tradeLicencePaymentFindQuery, nil)
	if err != nil {
		return nil, err
	}
	var tradeLicencePaymentsPart2 []models.RefTradeLicensePaymentsPart2
	if err = cursor.All(context.TODO(), &tradeLicencePaymentsPart2); err != nil {
		return nil, err
	}

	return tradeLicencePaymentsPart2, nil
}
