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
func (d *Daos) SaveMobileTowerPayment(ctx *models.Context, mtp *models.MobileTowerPayments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).InsertOne(ctx.CTX, mtp)
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)
	return err
}

//SaveMobileTowerPaymentFY :""
func (d *Daos) SaveMobileTowerPaymentFYs(ctx *models.Context, mtpfy []models.MobileTowerPaymentsfY) error {
	var insertData []interface{}
	for _, v := range mtpfy {
		insertData = append(insertData, v)
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).InsertMany(ctx.CTX, insertData)
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)
	return err
}

//SaveMobileTowerPaymentBasic :""
func (d *Daos) SaveMobileTowerPaymentBasic(ctx *models.Context, mtpBasic *models.MobileTowerPaymentsBasics) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).InsertOne(ctx.CTX, mtpBasic)
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)
	return err
}

func (d *Daos) GetSingleMobileTowerPayment(ctx *models.Context, tnxID string) (*models.RefMobileTowerPayments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tnxId": tnxID,
		},
	})
	mainPipeline = append(mainPipeline, d.RefQueryForMobileTowerPayment(ctx)...)
	d.Shared.BsonToJSONPrintTag("Get Single Mobile Tower Payment Query  - ", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rmtps []models.RefMobileTowerPayments
	var rmtp *models.RefMobileTowerPayments
	if err = cursor.All(ctx.CTX, &rmtps); err != nil {
		return nil, err
	}
	if len(rmtps) > 0 {
		rmtp = &rmtps[0]
	}
	return rmtp, nil
}

// RefQueryForMobileTowerPayment : ""
func (d *Daos) RefQueryForMobileTowerPayment(ctx *models.Context) []bson.M {
	var mainPipeline []bson.M
	//Look for FInancial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONMOBILETOWERPAYMENTSFY,
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
			"from": constants.COLLECTIONMOBILETOWERPAYMENTSBASIC,
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
	// Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.by", "userName", "ref.collector", "ref.collector")...)
	return mainPipeline
}

func (d *Daos) MakeMobileTowerPayment(ctx *models.Context, mmtpr *models.MakeMobileTowerPaymentReq) error {
	query := bson.M{
		"tnxId": mmtpr.TnxID,
	}
	paymentData := bson.M{
		"$set": bson.M{
			"details":                   mmtpr.Details,
			"status":                    mmtpr.Status,
			"creator":                   mmtpr.Creator,
			"completionDate":            mmtpr.CompletionDate,
			"reciptNo":                  d.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERPAYMENTSRECEIPT),
			"collectionReceived.status": "Pending",
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": mmtpr.Status,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

func (d *Daos) FilterMobileTowerPayment(ctx *models.Context, filter *models.MobileTowerPaymentsFilter, pagination *models.Pagination) ([]models.RefMobileTowerPayments, error) {
	var mainPipeline, query []bson.M
	if filter != nil {
		if len(filter.MobileTowerID) > 0 {
			query = append(query, bson.M{"mobileTowerId": bson.M{"$in": filter.MobileTowerID}})
		}
		if filter.Regex.OwnerName != "" {

			prop, err := d.GetPropertyIDsWithRegex(ctx, "name", filter.Regex.OwnerName)
			if err != nil {
				fmt.Println("Error in geting GetPropertyIDsWithRegex name" + err.Error())
			}
			if len(prop) > 0 {
				filter.PropertyID = append(filter.PropertyID, prop...)
			}
		}
		if filter.Regex.OwnerMobile != "" {

			prop, err := d.GetPropertyIDsWithRegex(ctx, "mobile", filter.Regex.OwnerMobile)
			if err != nil {
				fmt.Println("Error in geting GetPropertyIDsWithRegex mobile" + err.Error())
			}
			if len(prop) > 0 {
				filter.PropertyID = append(filter.PropertyID, prop...)
			}
		}
		if filter.Regex.PropertyID != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.Regex.PropertyID, Options: "xi"}})
		}
		if filter.Regex.MobileTowerID != "" {
			query = append(query, bson.M{"mobileTowerId": primitive.Regex{Pattern: filter.Regex.MobileTowerID, Options: "xi"}})
		}
		if filter.Regex.ReciptNo != "" {
			query = append(query, bson.M{"reciptNo": primitive.Regex{Pattern: filter.Regex.ReciptNo, Options: "xi"}})
		}

		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if len(filter.ReceiptNO) > 0 {
			query = append(query, bson.M{"reciptNo": bson.M{"$in": filter.ReceiptNO}})
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
		if len(filter.Scenario) > 0 {
			query = append(query, bson.M{"scenario": bson.M{"$in": filter.Scenario}})
		}
		if len(filter.MOP) > 0 {
			query = append(query, bson.M{"details.mop.mode": bson.M{"$in": filter.MOP}})
		}
		if len(filter.Collector) > 0 {
			query = append(query, bson.M{"details.collector.by": bson.M{"$in": filter.Collector}})
		}
		if len(filter.CollectorType) > 0 {
			query = append(query, bson.M{"details.collector.bytype": bson.M{"$in": filter.CollectorType}})
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
		// mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.RefQueryForMobileTowerPayment(ctx)...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Filter Mobile Tower Payment =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefMobileTowerPayments
	if err = cursor.All(context.TODO(), &payments); err != nil {
		return nil, err
	}
	return payments, nil

}

func (d *Daos) VerifyMobileTowerPayment(ctx *models.Context, action *models.MakeMobileTowerPaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.MobileTowerPaymentsAction,
			"status":       constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

func (d *Daos) NotVerifyMobileTowerPayment(ctx *models.Context, action *models.MakeMobileTowerPaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.MobileTowerPaymentsAction,
			"status":       constants.MOBILETOWERPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

func (d *Daos) RejectMobileTowerPayment(ctx *models.Context, action *models.MakeMobileTowerPaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.MobileTowerPaymentsAction,
			"status":       constants.MOBILETOWERPAYMENRSTATUSREJECTED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

//CalcMobileTowerDemand :""
func (d *Daos) CalcMobileTowerPaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefMobileTowerPayments, error) {
	d.Shared.BsonToJSONPrintTag("CalcMobileTowerDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}

	var mtps []models.RefMobileTowerPayments
	if err = cursor.All(context.TODO(), &mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

//CalcMobileTowerPendingPaymens :""
func (d *Daos) CalcMobileTowerPendingPaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefMobileTowerPayments, error) {
	d.Shared.BsonToJSONPrintTag("CalcMobileTowerDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var mtps []models.RefMobileTowerPayments
	if err = cursor.All(context.TODO(), &mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

// EnableMobileTowerPayments : ""
func (d *Daos) EnableMobileTowerPayments(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERSTATUSACTIVE}}
	d.Shared.BsonToJSONPrintTag("query query =>", query)
	d.Shared.BsonToJSONPrintTag("update =>", update)

	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// EnableMobileTowerPayments : ""
func (d *Daos) PendingMobileTowerPayments(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	d.Shared.BsonToJSONPrintTag("query query =>", query)
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERSTATUSPENDING}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
