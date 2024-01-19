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

// SavePropertyPartPayment : ""
func (d *Daos) SavePropertyPartPayment(ctx *models.Context, partPayment *models.PropertyPartPayment) error {
	d.Shared.BsonToJSONPrint(partPayment)
	status, err := d.GetMopPaymentStatus(ctx, partPayment.Details.MOP.Mode)
	if err != nil {
		return err
	}
	partPayment.Status = status
	if status == constants.PROPERTYPAYMENTCOMPLETED {
		t := time.Now()
		partPayment.CompletionDate = &t
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).InsertOne(ctx.CTX, partPayment)
	return err
}

func (d *Daos) GetSinglePropertyPartPaymentPayment(ctx *models.Context, tnxID string) (*models.RefPropertyPartPayments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"uniqueId": tnxID,
		},
	})

	d.Shared.BsonToJSONPrintTag("Get Single Property Part Payment Query  - ", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rppps []models.RefPropertyPartPayments
	var rppp *models.RefPropertyPartPayments
	if err = cursor.All(ctx.CTX, &rppps); err != nil {
		return nil, err
	}
	if len(rppps) > 0 {
		rppp = &rppps[0]
	}
	return rppp, nil
}

// VerifyPropertyPartPayment : ""
func (d *Daos) VerifyPropertyPartPayment(ctx *models.Context, action *models.MakePropertyPartPaymentsAction) (string, error) {
	payment, err := d.GetSinglePropertyPartPaymentPayment(ctx, action.UniqueID)
	if err != nil {
		return "", err
	}
	query := bson.M{"uniqueId": action.UniqueID}
	t := time.Now()
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo":   action.PropertyPartPaymentsAction,
			"status":         constants.PROPERTYPARTPAYMENTSTATUSCOMPLETED,
			"completionDate": &t,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.PROPERTYPARTPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.PROPERTYPARTPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}

	d.Shared.BsonToJSONPrintTag("Property part payment basic resp - ", res)
	return payment.PropertyPartPaymentID, nil
}

// NotVerifyPropertyPartPayment :
func (d *Daos) NotVerifyPropertyPartPayment(ctx *models.Context, action *models.MakePropertyPartPaymentsAction) (string, error) {
	payment, err := d.GetSinglePropertyPartPaymentPayment(ctx, action.UniqueID)
	if err != nil {
		return "", err
	}
	query := bson.M{"uniqueId": action.UniqueID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.PropertyPartPaymentsAction,
			"status":       constants.PROPERTYPARTPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.PROPERTYPARTPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.PROPERTYPARTPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment basic resp - ", res)
	return payment.PropertyID, nil
}

// RejectPropertyPartPayment :
func (d *Daos) RejectPropertyPartPayment(ctx *models.Context, action *models.MakePropertyPartPaymentsAction) (string, error) {
	payment, err := d.GetSinglePropertyPartPaymentPayment(ctx, action.UniqueID)
	if err != nil {
		return "", err
	}
	query := bson.M{"uniqueId": action.UniqueID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.PropertyPartPaymentsAction,
			"status":       constants.PROPERTYPARTPAYMENTSTATUSREJECTED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.PROPERTYPARTPAYMENTSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.PROPERTYPARTPAYMENTSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Property part payment basic resp - ", res)
	return payment.PropertyPartPaymentID, nil
}

//FilterPropertyPartPayment : ""
func (d *Daos) FilterPropertyPartPayment(ctx *models.Context, filter *models.PropertyPartPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyPartPayment, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var partPayments []models.RefPropertyPartPayment
	if err = cursor.All(context.TODO(), &partPayments); err != nil {
		return nil, err
	}
	return partPayments, nil
}

// //CalcPropertyPartDemand :""
// func (d *Daos) CalcPropertyPartPaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefPropertyPartPayments, error) {
// 	d.Shared.BsonToJSONPrintTag("CalcPropertyPartDemand query =>", mainPipeline)
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).Aggregate(ctx.CTX, mainPipeline)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var mtps []models.RefPropertyPartPayments
// 	if err = cursor.All(context.TODO(), &mtps); err != nil {
// 		return nil, err
// 	}
// 	return mtps, nil
// }

// //CalcPropertyPartPendingPaymens :""
// func (d *Daos) CalcPropertyPartPendingPaymens(ctx *models.Context, mainPipeline []bson.M) ([]models.RefPropertyPartPayments, error) {
// 	d.Shared.BsonToJSONPrintTag("CalcPropertyPartDemand query =>", mainPipeline)
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).Aggregate(ctx.CTX, mainPipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var mtps []models.RefPropertyPartPayments
// 	if err = cursor.All(context.TODO(), &mtps); err != nil {
// 		return nil, err
// 	}
// 	return mtps, nil
// }

// func (d *Daos) GetPayedFinancialYearsOfPropertyPartPayment(ctx *models.Context, PartPaymentID string) ([]string, error) {
// 	var mainPipeline []bson.M
// 	mainPipeline = append(mainPipeline, bson.M{
// 		"$match": bson.M{"propertyPartPaymentId": PartPaymentID,
// 			"status": bson.M{"$in": []string{constants.PROPERTYPARTPAYMENTSTATUSCOMPLETED}},
// 		},
// 	})
// 	mainPipeline = append(mainPipeline, bson.M{
// 		"$lookup": bson.M{
// 			"from": constants.COLLECTIONPROPERTYPARTPAYMENTSFY,
// 			"as":   "completedFys",
// 			"let":  bson.M{"id": "$propertyPartPaymentId"},
// 			"pipeline": []bson.M{
// 				bson.M{"$match": bson.M{
// 					"$expr": bson.M{"$and": []bson.M{
// 						bson.M{"$eq": []interface{}{"$propertyPartPaymentId", "$$id"}},
// 						bson.M{"$in": []interface{}{"$status", []string{constants.PROPERTYPARTPAYMENTSTATUSCOMPLETED}}},
// 					}},
// 				}},
// 				bson.M{
// 					"$group": bson.M{"_id": "$fy.uniqueId"},
// 				},
// 				bson.M{
// 					"$group": bson.M{"_id": nil, "data": bson.M{"$push": "$_id"}},
// 				},
// 			},
// 		},
// 	})
// 	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
// 		"completedFys": bson.M{"$arrayElemAt": []interface{}{"$completedFys", 0}},
// 	}})

// 	//Aggregation
// 	d.Shared.BsonToJSONPrintTag("Calculating payed Fys =>", mainPipeline)
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	type sample struct {
// 		CompletedFys struct {
// 			Data []string `json:"data" bson:"data,omitempty"`
// 		} `json:"completedFys" bson:"completedFys,omitempty"`
// 	}
// 	var payments []sample

// 	if err = cursor.All(context.TODO(), &payments); err != nil {
// 		return nil, err
// 	}
// 	if len(payments) > 0 {
// 		return payments[0].CompletedFys.Data, nil
// 	}
// 	return []string{}, nil
// }

// //CalcPropertyPartPaymentDemand :""
// func (d *Daos) CalcPropertyPartPaymentDemand(ctx *models.Context, mainPipeline []bson.M) (*models.PropertyPartPaymentDemand, error) {
// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)
// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
// 	d.Shared.BsonToJSONPrintTag("CalcPropertyPartPaymentDemand query =>", mainPipeline)
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).Aggregate(ctx.CTX, mainPipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var demands []models.PropertyPartPaymentDemand
// 	var demand *models.PropertyPartPaymentDemand
// 	if err = cursor.All(ctx.CTX, &demands); err != nil {
// 		return nil, err
// 	}
// 	if len(demands) > 0 {
// 		demand = &demands[0]
// 	}
// 	return demand, nil
// }

// GetPropertyPaymentsWithPartPayments : ""
func (d *Daos) GetPropertyPaymentsWithPartPayments(ctx *models.Context, tnxId string) (*models.RefMOPPartPayment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": tnxId}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			// "from": "propertypartpayments",
			"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
			"as":   "ref2.completed",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$tnxId", "$$tnxId"}},
					{"$eq": []string{"$status", constants.PROPERTYPARTPAYMENTSTATUSCOMPLETED}},
				}}},
			},
			}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			// "from": "propertypartpayments",
			"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
			"as":   "ref2.chequeBounced",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
					bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPARTPAYMENTSTATUSNOTVERIFIED}},
					bson.M{"$eq": []interface{}{"$details.mop.mode", constants.MOPCHEQUE}},
				}}},
			},
				bson.M{"$lookup": bson.M{"from": constants.COLLECTIONUSER, "localField": "details.collector.id", "foreignField": "userName", "as": "ref.collector"}},
				bson.M{"$addFields": bson.M{"ref.collector": bson.M{"$arrayElemAt": []interface{}{"$ref.collector", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			// "from": "propertypartpayments",
			"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
			"as":   "ref2.chequePending",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
					bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPARTPAYMENTSTATUSPENDING}},
					bson.M{"$eq": []interface{}{"$details.mop.mode", constants.MOPCHEQUE}},
				}}},
			},
				bson.M{"$lookup": bson.M{"from": constants.COLLECTIONUSER, "localField": "details.collector.id", "foreignField": "userName", "as": "ref.collector"}},
				bson.M{"$addFields": bson.M{"ref.collector": bson.M{"$arrayElemAt": []interface{}{"$ref.collector", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			// "from": "propertypartpayments",
			"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
			"as":   "ref2.ddnbBounced",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
					bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPARTPAYMENTSTATUSNOTVERIFIED}},
					bson.M{"$in": []interface{}{"$details.mop.mode", []string{constants.MOPDD, constants.MOPNETBANKING}}},
				}}},
			},
				bson.M{"$lookup": bson.M{"from": constants.COLLECTIONUSER, "localField": "details.collector.id", "foreignField": "userName", "as": "ref.collector"}},
				bson.M{"$addFields": bson.M{"ref.collector": bson.M{"$arrayElemAt": []interface{}{"$ref.collector", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			// "from": "propertypartpayments",
			"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
			"as":   "ref2.ddnbPending",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
					bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPARTPAYMENTSTATUSPENDING}},
					bson.M{"$in": []interface{}{"$details.mop.mode", []string{constants.MOPDD, constants.MOPNETBANKING}}},
				}}},
			},
				bson.M{"$lookup": bson.M{"from": constants.COLLECTIONUSER, "localField": "details.collector.id", "foreignField": "userName", "as": "ref.collector"}},
				bson.M{"$addFields": bson.M{"ref.collector": bson.M{"$arrayElemAt": []interface{}{"$ref.collector", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			// "from": "propertypartpayments",
			"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
			"as":   "ref2.rejected",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
					bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPARTPAYMENTSTATUSREJECTED}},
					// bson.M{"$eq": []interface{}{"$details.mop.mode", constants.MOPCHEQUE}},
				}}},
			},
				bson.M{"$lookup": bson.M{"from": constants.COLLECTIONUSER, "localField": "details.collector.id", "foreignField": "userName", "as": "ref.collector"}},
				bson.M{"$addFields": bson.M{"ref.collector": bson.M{"$arrayElemAt": []interface{}{"$ref.collector", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROPERTYOWNER, "propertyId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Module query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var modules []models.RefMOPPartPayment
	var Module *models.RefMOPPartPayment
	if err = cursor.All(ctx.CTX, &modules); err != nil {
		return nil, err
	}
	if len(modules) > 0 {
		Module = &modules[0]
	}
	return Module, nil
}
