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

//SavePropertyPayment :""
func (d *Daos) SavePropertyPayment(ctx *models.Context, propertyPayment *models.PropertyPayment) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).InsertOne(ctx.SC, propertyPayment)
	return err
}

//GetSinglePropertyPaymentWithTxtID :""
func (d *Daos) GetSinglePropertyPaymentWithTxtID(ctx *models.Context, txtID string) (*models.PropertyPayment, error) {
	mainPipeline := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{

		"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
		"as":   "ref.partPayments",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$tnxId", "$$tnxId"}},
				{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			}}}},
		},
	}})

	d.Shared.BsonToJSONPrintTag("property payment query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.PropertyPayment
	var pp *models.PropertyPayment
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

//GetSinglePropertyPaymentReceiptNo :""
func (d *Daos) GetSinglePropertyPaymentReceiptNo(ctx *models.Context, receiptNo string) (*models.PropertyPayment, error) {
	mainPipeline := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"reciptNo": receiptNo}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{

		"from": constants.COLLECTIONPROPERTYPARTPAYMENT,
		"as":   "ref.partPayments",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$tnxId", "$$tnxId"}},
				{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			}}}},
		},
	}})

	d.Shared.BsonToJSONPrintTag("property payment query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.PropertyPayment
	var pp *models.PropertyPayment
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

//CompletePropertyPaymentWithTxtID : ""
func (d *Daos) CompletePropertyPaymentWithTxtID(ctx *models.Context, payment *models.PropertyMakePayment, isPartPayment bool) error {

	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}

	if payment != nil {
		if payment.Details.MOP.Mode == "NB" && payment.Details.MOP.PropertyPaymentCardRNet.Vendor == "HDFC_CC_AVENUE" {
			status = constants.PROPERTYPAYMENTCOMPLETED
		}
	}

	selector := bson.M{"tnxId": payment.TnxID}
	update := bson.M{"$set": bson.M{
		"status":                    status,
		"details":                   payment.Details,
		"completionDate":            time.Now(),
		"reciptURL":                 payment.ReciptURL,
		"collectionReceived.status": "Pending",
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

//CompletePropertyPaymentWithTxtID : ""
func (d *Daos) CompletePropertyPaymentWithTxtIDForceFulForPartPayment(ctx *models.Context, tnxID string) error {
	selector := bson.M{"tnxId": tnxID}
	update := bson.M{"$set": bson.M{
		"status":         constants.PROPERTYPAYMENTCOMPLETED,
		"completionDate": time.Now(),
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

//UpdatePendingAmount : ""
func (d *Daos) UpdatePendingAmount(ctx *models.Context, tnxID string, pendingAmount float64) error {
	selector := bson.M{"tnxId": tnxID}
	update := bson.M{"$set": bson.M{

		"pendingAmount": pendingAmount,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

//UpdatePendingAmount : ""
func (d *Daos) UpdateDetailsAmountForCompletedPartPayment(ctx *models.Context, tnxID string, amount float64) error {
	selector := bson.M{"tnxId": tnxID}
	update := bson.M{"$set": bson.M{

		"details.amount":   amount,
		"details.mop.mode": "PartPayment",
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

//CompletePropertyPaymentWithTxtID : ""
func (d *Daos) PartiallyCompletedPropertyPaymentWithTxtIDForceFulForPartPayment(ctx *models.Context, TnxID string) error {
	selector := bson.M{"tnxId": TnxID}
	update := bson.M{"$set": bson.M{
		"status": constants.PROPERTYPAYMENTPARTIALLYCOMPLETED,
		"type":   constants.PROPERTYPAYMENTTYPEPARTPAYMENT,
		//	"details": propertyPayment,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

// CompletePropertyPaymentWithTxtIDPart : ""
func (d *Daos) CompletePropertyPaymentWithTxtIDPart(ctx *models.Context, payment *models.PropertyMakePayment) error {
	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}
	selector := bson.M{"tnxId": payment.TnxID}
	update := bson.M{"$set": bson.M{
		"status": status,
		"completionDate": func() *time.Time {
			if status == constants.PROPERTYPAYMENTCOMPLETED {
				t := time.Now()
				return &t
			}
			return nil
		}(),
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

// RecordPartPayment : ""
func (d *Daos) RecordPartPayment(ctx *models.Context, payment *models.PropertyMakePayment) error {
	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}
	data := bson.M{"$set": bson.M{
		"tnxId":   payment.TnxID,
		"status":  status,
		"details": payment.Details,
		"completionDate": func() *time.Time {
			if payment != nil {
				if payment.Details.MOP.Mode == constants.MOPCASH {
					t := time.Now()
					return &t
				}
			}
			return nil
		}(),
		"paymentDate": time.Now(),
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPARTPAYMENT).InsertOne(ctx.CTX, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

//SavePropertyPaymentDemandBasic :""
func (d *Daos) SavePropertyPaymentDemandBasic(ctx *models.Context, propertyPaymentDemandBasic *models.PropertyPaymentDemandBasic) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).InsertOne(ctx.SC, propertyPaymentDemandBasic)
	return err
}

//GetSinglePropertyPaymentDemandBasicWithTxtID :""
func (d *Daos) GetSinglePropertyPaymentDemandBasicWithTxtID(ctx *models.Context, txtID string) (*models.PropertyPaymentDemandBasic, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdbs []models.PropertyPaymentDemandBasic
	var ppdb *models.PropertyPaymentDemandBasic
	if err = cursor.All(ctx.CTX, &ppdbs); err != nil {
		return nil, err
	}
	if len(ppdbs) > 0 {
		ppdb = &ppdbs[0]
	}
	return ppdb, nil
}

//CompleteSinglePropertyPaymentDemandBasicWithTxtID : ""
func (d *Daos) CompleteSinglePropertyPaymentDemandBasicWithTxtID(ctx *models.Context, payment *models.PropertyMakePayment) error {
	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}
	if payment != nil {
		if payment.Details.MOP.Mode == "NB" && payment.Details.MOP.PropertyPaymentCardRNet.Vendor == "HDFC_CC_AVENUE" {
			status = constants.PROPERTYPAYMENTCOMPLETED
		}
	}

	selector := bson.M{"tnxId": payment.TnxID}
	update := bson.M{"$set": bson.M{
		"status": status,
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

//SaveManyPropertyPaymentDemandFy :""
func (d *Daos) SaveManyPropertyPaymentDemandFy(ctx *models.Context, ppdfys []models.PropertyPaymentDemandFy) error {
	var data []interface{}
	for _, v := range ppdfys {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).InsertMany(ctx.SC, data)
	return err
}

//GetPropertyPaymentDemandFycWithTxtID :""
func (d *Daos) GetPropertyPaymentDemandFycWithTxtID(ctx *models.Context, txtID string) ([]models.PropertyPaymentDemandFy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"fy.order": -1}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdfys []models.PropertyPaymentDemandFy
	if err = cursor.All(ctx.CTX, &ppdfys); err != nil {
		return nil, err
	}

	return ppdfys, nil
}

//CompletePropertyPaymentDemandFycWithTxtID : ""
func (d *Daos) CompletePropertyPaymentDemandFycWithTxtID(ctx *models.Context, payment *models.PropertyMakePayment) error {
	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}
	if payment != nil {
		if payment.Details.MOP.Mode == "NB" && payment.Details.MOP.PropertyPaymentCardRNet.Vendor == "HDFC_CC_AVENUE" {
			status = constants.PROPERTYPAYMENTCOMPLETED
		}
	}

	selector := bson.M{"tnxId": payment.TnxID}
	update := bson.M{"$set": bson.M{
		"status": status,
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, selector, update, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

//GetAllPaymentsForProperty : ""
func (d *Daos) GetAllPaymentsForProperty(ctx *models.Context, id string) ([]models.RefPropertyPayment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"propertyId": id, "status": constants.PROPERTYPAYMENTCOMPLETED}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENTBASIC,
		"as":   "basic",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
				bson.M{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			}}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"basic": bson.M{"$arrayElemAt": []interface{}{"$basic", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "fys",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$tnxId", "$$tnxId"}},
				bson.M{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			}}}},
			bson.M{"$sort": bson.M{"fy.order": -1}},
		},
	}})
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefPropertyPayment
	if err = cursor.All(ctx.CTX, &payments); err != nil {
		return nil, err
	}
	return payments, nil
}

//DashboardTotalCollection : ""
func (d *Daos) DashboardTotalCollection(ctx *models.Context, filter *models.DashboardTotalCollectionFilter) ([]models.DashboardTotalCollectionRef, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "currentYear",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"status", "Completed"}},
				bson.M{"$eq": []interface{}{"$fy.isCurrent", true}},
				bson.M{"$eq": []interface{}{"$tnxId", "$$tnxId"}},
			}}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "arriearYears",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"status", "Completed"}},
				bson.M{"$eq": []interface{}{"$fy.isCurrent", false}},
				bson.M{"$eq": []interface{}{"$tnxId", "$$tnxId"}},
			}}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"currentYear": bson.M{"$arrayElemAt": []interface{}{"$currentYear", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.DashboardTotalCollectionRef
	if err = cursor.All(ctx.CTX, &payments); err != nil {
		return nil, err
	}
	return payments, nil

}

//FilterPropertyPayment : ""
func (d *Daos) FilterPropertyPayment(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$propertyId", "data": bson.M{"$push": "$$ROOT"}}})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data": bson.M{"$arrayElemAt": []interface{}{"$data", 0}}}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYPAYMENTBASIC, "tnxId", "tnxId", "basic", "basic")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROPERTYPAYMENTFY, "tnxId", "tnxId", "fys", "fys")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.id", "userName", "ref.collector", "ref.collector")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "collectionReceived.by", "userName", "ref.collectionReceivedBy", "ref.collectionReceivedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "collectionReceived.byType", "uniqueId", "ref.collectionReceivedByType", "ref.collectionReceivedByType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejectedInfo.by", "uniqueId", "ref.rejectedBy", "ref.rejectedBy")...)

	// mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "basic.property.uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.RefPropertyPayment
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *Daos) FilterPropertyPaymentQuery(ctx *models.Context, filter *models.PropertyPaymentFilter) []bson.M {
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.CollectionReceivedStatus) > 0 {
			query = append(query, bson.M{"collectionReceived.status": bson.M{"$in": filter.CollectionReceivedStatus}})
		}
		if filter.AvoidPartPayment {
			query = append(query, bson.M{"type": bson.M{"$nin": []string{constants.PROPERTYPAYMENTTYPEPARTPAYMENT}}})
		}
		if len(filter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.Type}})
		}
		if len(filter.PropertyIds) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyIds}})
		}
		if len(filter.MadeAt) > 0 {
			query = append(query, bson.M{"details.madeAt.at": bson.M{"$in": filter.MadeAt}})
		}
		if len(filter.MOP) > 0 {
			query = append(query, bson.M{"details.mop.mode": bson.M{"$in": filter.MOP}})
		}
		if len(filter.Collector) > 0 {
			query = append(query, bson.M{"details.collector.id": bson.M{"$in": filter.Collector}})
		}
		if len(filter.CollectorByType) > 0 {
			query = append(query, bson.M{"details.collector.type": bson.M{"$in": filter.CollectorByType}})
		}
		if len(filter.CollectionReceivedBy) > 0 {
			query = append(query, bson.M{"collectionReceived.by": bson.M{"$in": filter.CollectionReceivedBy}})
		}

		if len(filter.ReceiptNo) > 0 {
			query = append(query, bson.M{"reciptNo": bson.M{"$in": filter.ReceiptNo}})
		}
		if filter.DateRange != nil {
			//var sd,ed time.Time
			if filter.DateRange.From != nil {
				sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				}
				query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.BouncedDateRange != nil {
			//var sd,ed time.Time
			if filter.BouncedDateRange.From != nil {
				sd := time.Date(filter.BouncedDateRange.From.Year(), filter.BouncedDateRange.From.Month(), filter.BouncedDateRange.From.Day(), 0, 0, 0, 0, filter.BouncedDateRange.From.Location())
				ed := time.Date(filter.BouncedDateRange.From.Year(), filter.BouncedDateRange.To.Month(), filter.BouncedDateRange.To.Day(), 23, 59, 59, 0, filter.BouncedDateRange.To.Location())
				if filter.BouncedDateRange.To != nil {
					ed = time.Date(filter.BouncedDateRange.To.Year(), filter.BouncedDateRange.To.Month(), filter.BouncedDateRange.To.Day(), 23, 59, 59, 0, filter.BouncedDateRange.To.Location())
				}
				query = append(query, bson.M{"notVerifiedInfo.notVerifiedDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.CollectionReceivedDateRange != nil {
			//var sd,ed time.Time
			if filter.CollectionReceivedDateRange.From != nil {
				sd := time.Date(filter.CollectionReceivedDateRange.From.Year(), filter.CollectionReceivedDateRange.From.Month(), filter.CollectionReceivedDateRange.From.Day(), 0, 0, 0, 0, filter.CollectionReceivedDateRange.From.Location())
				ed := time.Date(filter.CollectionReceivedDateRange.From.Year(), filter.CollectionReceivedDateRange.To.Month(), filter.CollectionReceivedDateRange.To.Day(), 23, 59, 59, 0, filter.CollectionReceivedDateRange.To.Location())
				if filter.CollectionReceivedDateRange.To != nil {
					ed = time.Date(filter.CollectionReceivedDateRange.To.Year(), filter.CollectionReceivedDateRange.To.Month(), filter.CollectionReceivedDateRange.To.Day(), 23, 59, 59, 0, filter.CollectionReceivedDateRange.To.Location())
				}
				query = append(query, bson.M{"collectionReceived.date": bson.M{"$gte": sd, "$lte": ed}})

			}
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
		if filter.SearchText.OwnerName != "" {

			propertyIds, err := d.GetPropertyIDsWithOwnerNames(ctx, filter.SearchText.OwnerName)
			if err != nil {
				log.Println("ERR IN GETING - Property IDs WithOwner Names " + err.Error())
			} else {
				if len(propertyIds) > 0 {
					fmt.Println("got Property Ids - ", propertyIds)
					query = append(query, bson.M{"uniqueId": bson.M{"$in": propertyIds}})
				}
			}
		}
		if filter.SearchText.HoldingNo != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.SearchText.HoldingNo, Options: "xi"}})
		}
		if filter.SearchText.ReceiptNo != "" {
			query = append(query, bson.M{"reciptNo": primitive.Regex{Pattern: filter.SearchText.ReceiptNo, Options: "xi"}})
		}
		// if filter.SearchText.OwnerMobile != "" {

		// 	propertyIds, err := d.GetPropertyIDsWithOwnerNames(ctx, filter.SearchText.OwnerMobile)
		// 	if err != nil {
		// 		log.Println("ERR IN GETING - Property IDs WithOwner Names " + err.Error())
		// 	} else {
		// 		if len(propertyIds) > 0 {
		// 			fmt.Println("got Property Ids - ", propertyIds)
		// 			query = append(query, bson.M{"uniqueId": bson.M{"$in": propertyIds}})
		// 		}
		// 	}
		// }
	}

	return query
}

// VerifyPayment : ""
func (d *Daos) VerifyPayment(ctx *models.Context, vp *models.VerifyPayment) error {
	selector := bson.M{"tnxId": vp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTCOMPLETED, "remark": ""}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTCOMPLETED, "remark": "",
		"verifiedInfo": bson.M{"verifiedActionDate": vp.ActionDate, "verifiedDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// NotVerifiedPayment : ""
func (d *Daos) NotVerifiedPayment(ctx *models.Context, vp *models.NotVerifiedPayment) error {
	selector := bson.M{"tnxId": vp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTNOTVERIFIED, "remark": vp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTNOTVERIFIED, "remark": vp.Remarks,
		"notVerifiedInfo": bson.M{"notVerifiedActionDate": vp.ActionDate, "notVerifiedDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// GetMopPaymentStatus : ""
func (d *Daos) GetMopPaymentStatus(ctx *models.Context, mop string) (string, error) {
	res, err := d.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		return "", err
	}

	if res.CompleteChequePayment == "Yes" {
		switch mop {
		case constants.MOPCASH, constants.MOPCHEQUE:
			return constants.PROPERTYPAYMENTCOMPLETED, nil
		case constants.MOPNETBANKING, constants.MOPDD:
			return constants.PROPERTYPAYMENTVERIFICATIONPENDING, nil
		default:
			return "", errors.New("invalid mop")
		}
	} else {
		switch mop {
		case constants.MOPCASH:
			return constants.PROPERTYPAYMENTCOMPLETED, nil
		case constants.MOPCHEQUE, constants.MOPNETBANKING, constants.MOPDD:
			return constants.PROPERTYPAYMENTVERIFICATIONPENDING, nil
		default:
			return "", errors.New("invalid mop")
		}
	}

}

// BouncePayment : ""
func (d *Daos) BouncePayment(ctx *models.Context, bp *models.BouncePayment) error {
	selector := bson.M{"tnxId": bp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTBOUNCED, "remark": bp.Remarks,
		"bouncedInfo": bson.M{"bouncedActionDate": bp.ActionDate, "bouncedDate": bp.Date, "remark": bp.Remarks, "by": bp.By, "byType": bp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// RejectPayment : ""
func (d *Daos) RejectPayment(ctx *models.Context, rp *models.RejectPayment) error {
	selector := bson.M{"tnxId": rp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENREJECTED, "rejectedRemark": rp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENREJECTED, "rejectedRemark": rp.Remarks,
		"rejectedInfo": bson.M{"rejectedActionDate": rp.ActionDate, "rejectedDate": rp.Date, "remark": rp.Remarks, "by": rp.By, "byType": rp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

//RejectPaymentByReceiptNo

func (d *Daos) RejectPaymentByReceiptNo(ctx *models.Context, rp *models.RejectPayment) error {
	selector := bson.M{"tnxId": rp.TnxID}
	addStatus := func() string {
		if rp.Status == "" {
			return constants.PROPERTYPAYMENREJECTED
		}
		return rp.Status
	}()
	update := bson.M{"$set": bson.M{"status": addStatus, "rejectedRemark": rp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": addStatus, "rejectedRemark": rp.Remarks,
		"rejectedInfo": bson.M{"rejectedActionDate": rp.ActionDate, "rejectedDate": rp.Date, "remark": rp.Remarks, "by": rp.By, "byType": rp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// UpdateOverallDashBoard : ""
func (d *Daos) UpdateOverallDashBoard(ctx *models.Context, arrear float64, current float64, total float64, property float64, res *models.RefDateWisePropertyPaymentReport) error {
	d.Shared.BsonToJSONPrint(res)
	fmt.Println("property.overall.collection.arrear", arrear,
		"property.overall.collection.current", current,
		"property.overall.collection.total", total)
	query := bson.M{"isDefault": true}
	update := bson.M{"$set": bson.M{

		"property.overall.collection.arrear":  arrear,
		"property.overall.collection.current": current,
		"property.overall.collection.total":   total,
		// "property.overall.collection.propertyCount": res.Report.Overall.PropertyCount,
		"property.overall.collection.propertyCount": property,

		//year
		"property.year.collection.arrear":        res.Report.Year.ArrearCollection,
		"property.year.collection.current":       res.Report.Year.CurrentCollection,
		"property.year.collection.total":         res.Report.Year.TotalCollection,
		"property.year.collection.propertyCount": res.Report.Year.PropertyCount,

		//month
		"property.month.collection.arrear":        res.Report.Month.ArrearCollection,
		"property.month.collection.current":       res.Report.Month.CurrentCollection,
		"property.month.collection.total":         res.Report.Month.TotalCollection,
		"property.month.collection.propertyCount": res.Report.Month.PropertyCount,

		//month
		"property.today.collection.arrear":        res.Report.Day.ArrearCollection,
		"property.today.collection.current":       res.Report.Day.CurrentCollection,
		"property.today.collection.total":         res.Report.Day.TotalCollection,
		"property.today.collection.propertyCount": res.Report.Day.PropertyCount,

		//week
		"property.week.collection.arrear":        res.Report.Week.ArrearCollection,
		"property.week.collection.current":       res.Report.Week.CurrentCollection,
		"property.week.collection.total":         res.Report.Week.TotalCollection,
		"property.week.collection.propertyCount": res.Report.Week.PropertyCount,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DateRangeWisePropertyPaymentReport : ""
func (d *Daos) DateRangeWisePropertyPaymentReport(ctx *models.Context, filter *models.DateWisePropertyPaymentReportFilter) (*models.RefDateWisePropertyPaymentReport, error) {
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
	//Adding $match from filter
	// if len(query) > 0 {
	// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	// }
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"overall": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{
				bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
			}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear"},
				"currentCollection": bson.M{"$sum": "$demand.current"},
				"totalDemand":       bson.M{"$sum": "$demand.totalTax"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			bson.M{"$group": bson.M{"_id": nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalDemand":       bson.M{"$sum": "$totalDemand"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			},
			},
			bson.M{"$addFields": bson.M{"fyName": "$_id.fyName"}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$group": bson.M{"_id": nil, "overallReport": bson.M{"$push": "$$ROOT"}}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$addFields": bson.M{"overallReport": bson.M{"$arrayElemAt": []interface{}{"$overallReport", 0}}}},
		},

		"year": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{
				bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
				bson.M{"completionDate": bson.M{"$gte": resFY.From,
					"$lte": resFY.To}},
			}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear"},
				"currentCollection": bson.M{"$sum": "$demand.current"},
				"totalDemand":       bson.M{"$sum": "$demand.totalTax"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			bson.M{"$group": bson.M{"_id": nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalDemand":       bson.M{"$sum": "$totalDemand"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			},
			},
			bson.M{"$addFields": bson.M{"fyName": "$_id.fyName"}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$group": bson.M{"_id": nil, "yearWiseReport": bson.M{"$push": "$$ROOT"}}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$addFields": bson.M{"yearWiseReport": bson.M{"$arrayElemAt": []interface{}{"$yearWiseReport", 0}}}},
		},
		"month": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{
				bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
				bson.M{"completionDate": bson.M{"$gte": monthsd,
					"$lte": monthed}},
			}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear"},
				"currentCollection": bson.M{"$sum": "$demand.current"},
				"totalDemand":       bson.M{"$sum": "$demand.totalTax"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			bson.M{"$group": bson.M{"_id": nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalDemand":       bson.M{"$sum": "$totalDemand"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			},
			},
			bson.M{"$addFields": bson.M{"fyMonth": "$_id.fyMonth"}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$group": bson.M{"_id": nil, "monthWiseReport": bson.M{"$push": "$$ROOT"}}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$addFields": bson.M{"monthWiseReport": bson.M{"$arrayElemAt": []interface{}{"$monthWiseReport", 0}}}},
		},
		"week": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{
				bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
				bson.M{"completionDate": bson.M{"$gte": weeksd,
					"$lte": weeked}},
			}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear"},
				"currentCollection": bson.M{"$sum": "$demand.current"},
				"totalDemand":       bson.M{"$sum": "$demand.totalTax"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			bson.M{"$group": bson.M{"_id": nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalDemand":       bson.M{"$sum": "$totalDemand"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			},
			},
			bson.M{"$addFields": bson.M{"fyWeek": "$_id.fyWeek"}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$group": bson.M{"_id": nil, "weekWiseReport": bson.M{"$push": "$$ROOT"}}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$addFields": bson.M{"weekWiseReport": bson.M{"$arrayElemAt": []interface{}{"$weekWiseReport", 0}}}},
		},
		"day": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{
				bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
				bson.M{"completionDate": bson.M{"$gte": daysd,
					"$lte": dayed}},
			}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"arrearCollection":  bson.M{"$sum": "$demand.arrear"},
				"currentCollection": bson.M{"$sum": "$demand.current"},
				"totalDemand":       bson.M{"$sum": "$demand.totalTax"},
				"totalCollection":   bson.M{"$sum": "$details.amount"},
			}},
			bson.M{"$group": bson.M{"_id": nil,
				"propertyCount":     bson.M{"$sum": 1},
				"arrearCollection":  bson.M{"$sum": "$arrearCollection"},
				"currentCollection": bson.M{"$sum": "$currentCollection"},
				"totalDemand":       bson.M{"$sum": "$totalDemand"},
				"totalCollection":   bson.M{"$sum": "$totalCollection"},
			},
			},
			bson.M{"$addFields": bson.M{"fyDay": "$_id.fyDay"}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$group": bson.M{"_id": nil, "dayWiseReport": bson.M{"$push": "$$ROOT"}}},
			bson.M{"$project": bson.M{"_id": 0}},
			bson.M{"$addFields": bson.M{"dayWiseReport": bson.M{"$arrayElemAt": []interface{}{"$dayWiseReport", 0}}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"overall": bson.M{"$arrayElemAt": []interface{}{"$overall.overallReport", 0}}}},
		bson.M{"$addFields": bson.M{"year": bson.M{"$arrayElemAt": []interface{}{"$year.yearWiseReport", 0}}}},
		bson.M{"$addFields": bson.M{"month": bson.M{"$arrayElemAt": []interface{}{"$month.monthWiseReport", 0}}}},
		bson.M{"$addFields": bson.M{"week": bson.M{"$arrayElemAt": []interface{}{"$week.weekWiseReport", 0}}}},
		bson.M{"$addFields": bson.M{"day": bson.M{"$arrayElemAt": []interface{}{"$day.dayWiseReport", 0}}}},
		bson.M{"$group": bson.M{"_id": nil, "report": bson.M{"$push": "$$ROOT"}}},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}},
		bson.M{"$project": bson.M{"_id": 0}},
	)
	// ============================================>

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var reports []models.RefDateWisePropertyPaymentReport
	var report *models.RefDateWisePropertyPaymentReport
	if err = cursor.All(ctx.CTX, &reports); err != nil {
		return nil, err
	}
	if len(reports) > 0 {

		report = &reports[0]
		report.Report.Year.FyName = resFY.Name
		report.Report.Month.FyMonth = monthName
		report.Report.Week.FyWeek = weekStart + " " + "to" + " " + weekEnd
		report.Report.Day.FyDay = day

	}

	return report, nil
}

func (d *Daos) UpdatePropertyPayments(ctx *models.Context, propertyPayment *models.PropertyPayment) error {
	selector := bson.M{"tnxId": propertyPayment.TnxID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": propertyPayment}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) DateWisePropertyPaymentReport(ctx *models.Context, filter *models.DateWisePropertyPaymentReportFilter) ([]models.DateWisePropertyPaymentReport, error) {
	mainPipeline := []bson.M{}
	// query := []bson.M{}

	// getting the start and end of the month
	var monthsd, monthed *time.Time
	monthsdt := d.Shared.BeginningOfMonth(*filter.Date)
	monthsd = &monthsdt
	monthedt := d.Shared.EndOfMonth(*filter.Date)
	monthed = &monthedt
	//t := filter.Date

	var daysd, dayed time.Time
	if filter != nil {
		daysd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, daysd.Location())
		dayed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 0, dayed.Location())
	}
	//Adding $match from filter
	// if len(query) > 0 {
	// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	// }
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": []bson.M{
		bson.M{"status": bson.M{"$in": []string{constants.PROPERTYPAYMENTCOMPLETED}}},
		bson.M{"completionDate": bson.M{"$gte": monthsd,
			"$lte": monthed}},
	}}},
		bson.M{"$group": bson.M{"_id": "$completionDate",
			"propertyCount":     bson.M{"$sum": 1},
			"arrearCollection":  bson.M{"$sum": "$demand.arrear"},
			"arrearPenalty":     bson.M{"$sum": "$demand.penalCharge"},
			"currentCollection": bson.M{"$sum": "$demand.current"},
			"totalDemand":       bson.M{"$sum": "$demand.totalTax"},
			"totalCollection":   bson.M{"$sum": "$details.amount"},
		},
		},
		bson.M{"$addFields": bson.M{"date": "$_id"}},
		bson.M{"$project": bson.M{"_id": 0}},
	)
	// ============================================>

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var report []models.DateWisePropertyPaymentReport
	//var data []models.DateWisePropertyPaymentReport
	if err = cursor.All(context.TODO(), &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (d *Daos) CollectedPropertyPayment(ctx *models.Context, PropertyPayment *models.CollectionReceived) error {
	selector := bson.M{"tnxId": PropertyPayment.TnxID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"collectionReceived": PropertyPayment}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) RejectedPropertyPayment(ctx *models.Context, PropertyPayment *models.CollectionReceived) error {
	selector := bson.M{"tnxId": PropertyPayment.TnxID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"collectionReceived": PropertyPayment}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// PropertyPaymentArrerAndCurrentCollection : ""
func (d *Daos) PropertyPaymentArrerAndCurrentCollection(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.ArrerAndCurrentReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	//mainPipeline = append(mainPipeline, d.FilterPropertyPaymentQuery(ctx, filter)...)
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if filter.StartDate != nil {
			sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
			ed := time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())

			query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "propertypaymentfys",
			"as":   "fys",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{

				{"$match": bson.M{"$expr": bson.M{"$and": bson.M{"$eq": []interface{}{"$tnxId", "$$tnxId"}}}}},
				{"$group": bson.M{
					"_id":                nil,
					"arrearTax":          bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}, "else": 0}}},
					"currentTax":         bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}, "else": 0}}},
					"arrearPenalty":      bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}, "else": 0}}},
					"currentPenalty":     bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}, "else": 0}}},
					"arrearRebate":       bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}, "else": 0}}},
					"currentRebate":      bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}, "else": 0}}},
					"arrearAlreadyPaid":  bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
					"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				}}}},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"fys": bson.M{"$arrayElemAt": []interface{}{"$fys", 0}},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"fys.formfee":        "$demand.formFee",
		"fys.completionDate": "$completionDate",
		"fys.penalty":        "$demand.penalCharge",
	}})
	mainPipeline = append(mainPipeline, bson.M{"$replaceRoot": bson.M{"newRoot": "$fys"}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
		"_id":                bson.M{"$dayOfMonth": "$completionDate"},
		"penalty":            bson.M{"$sum": "$penalty"},
		"arrearTax":          bson.M{"$sum": "$arrearTax"},
		"currentTax":         bson.M{"$sum": "$currentTax"},
		"arrearPenalty":      bson.M{"$sum": "$arrearPenalty"},
		"currentPenalty":     bson.M{"$sum": "$currentPenalty"},
		"arrearRebate":       bson.M{"$sum": "$arrearRebate"},
		"currentRebate":      bson.M{"$sum": "$currentRebate"},
		"arrearAlreadyPaid":  bson.M{"$sum": "$arrearAlreadyPaid"},
		"currentAlreadyPaid": bson.M{"$sum": "$currentAlreadyPaid"},
		"formfee":            bson.M{"$sum": "$formfee"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	// mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "basic.property.uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.ArrerAndCurrentReport
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdatePropertyPaymentBasicPropertyID :""
func (d *Daos) UpdatePropertyPaymentBasicPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"property.uniqueId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// UpdatePropertyPaymentFYPropertyID :""
func (d *Daos) UpdatePropertyPaymentFYPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTFY).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// UpdatePropertyPaymentPropertyID :""
func (d *Daos) UpdatePropertyPaymentPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) GetSinglePropertyPaymentTxtID(ctx *models.Context, txtID string) (*models.PropertyPayment, error) {
	mainPipeline := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})

	d.Shared.BsonToJSONPrintTag("property payment query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.PropertyPayment
	var pp *models.PropertyPayment
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

func (d *Daos) UpdatePropertyPayeenamewithTxnId(ctx *models.Context, TnxID string, name string) error {
	query := bson.M{"tnxId": TnxID}
	update := bson.M{"$set": bson.M{"details.payeeName": name}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) UpdatePropertyPaymentSummary(ctx *models.Context, propertypayment *models.Summary) error {
	selector := bson.M{"tnxId": propertypayment.TnxID}
	update := bson.M{"$set": bson.M{
		"summary": propertypayment,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}
