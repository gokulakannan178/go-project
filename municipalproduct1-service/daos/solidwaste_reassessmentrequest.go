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
)

// SaveSolidWasteReassessmentRequestUpdate :""
func (d *Daos) SaveSolidWasteReassessmentRequestUpdate(ctx *models.Context, request *models.SolidWasteReassessmentRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEREASSESSMENTREQUEST).InsertOne(ctx.CTX, request)
	return err
}

// GetSingleSolidWasteReassessmentRequest : ""
func (d *Daos) GetSingleSolidWasteReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefSolidWasteReassessmentRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "uniqueId", "ref.requestedUserType", "ref.requestedUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "action.byType", "uniqueId", "ref.actionUserType", "ref.actionUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "new.ref.address.state", "new.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "new.ref.address.district", "new.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "new.ref.address.village", "new.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "new.ref.address.zone", "new.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "new.ref.address.ward", "new.ref.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefSolidWasteReassessmentRequest
	var request *models.RefSolidWasteReassessmentRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// RejectSolidWasteReassessmentRequestUpdate : ""
func (d *Daos) AcceptSolidWasteReassessmentRequestUpdate(ctx *models.Context, accept *models.AcceptSolidWasteReassessmentRequestUpdate) error {
	t := time.Now()
	fmt.Println("shoprent accept daos")

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEREASSESSMENTREQUESTSTATUSCOMPLETED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectSolidWasteReassessmentRequestUpdate : ""
func (d *Daos) RejectSolidWasteReassessmentRequestUpdate(ctx *models.Context, reject *models.RejectSolidWasteReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEREASSESSMENTREQUESTSTATUSREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// BasicSolidWasteReassessmentRequestUpdateToPayments : ""
func (d *Daos) BasicSolidWasteReassessmentRequestUpdateToPayments(ctx *models.Context, request *models.RefSolidWasteReassessmentRequest) error {
	//get current Financial year

	cfy, err := d.GetCurrentFinancialYear(ctx)
	if err != nil {
		return errors.New("Error in getting current financial year " + err.Error())
	}
	if cfy == nil {
		return errors.New("current financial year is nil")
	}
	sd := time.Date(cfy.From.Year(), cfy.From.Month(), cfy.From.Day(), 0, 0, 0, 0, cfy.From.Location())
	ed := time.Date(cfy.To.Year(), cfy.To.Month(), cfy.To.Day(), 23, 59, 59, 0, cfy.To.Location())
	solidWastePaymentFindQuery := bson.M{
		"status":                 constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSCOMPLETED,
		"solidWasteUserChargeId": request.SolidWasteID,
		"completionDate":         bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("BasicReassessmentRequestUpdateToPayments query =>", solidWastePaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).Find(ctx.CTX, solidWastePaymentFindQuery, nil)
	if err != nil {
		return err
	}
	var solidWastePayments []models.RefShopRentPayments
	if err = cursor.All(context.TODO(), &solidWastePayments); err != nil {
		return err
	}

	if len(solidWastePayments) > 0 {
		fmt.Println("shoprentPayments is > 0")
		tnxIDs := []string{}
		for _, v := range solidWastePayments {
			tnxIDs = append(tnxIDs, v.TnxID)
			fmt.Println("solidWastePayment tnxId is", v.TnxID)

		}
		if len(tnxIDs) > 0 {
			solidwastePaymentQuery := bson.M{
				"tnxId": bson.M{"$in": tnxIDs},
			}
			d.Shared.BsonToJSONPrintTag("shoprentPayment query =>", solidwastePaymentQuery)
			paymentUpdate := bson.M{
				"$set": bson.M{"address": request.New.Address},
			}
			d.Shared.BsonToJSONPrintTag("paymentUpdate query =>", paymentUpdate)

			paymentUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).UpdateMany(ctx.CTX, solidwastePaymentQuery, paymentUpdate)
			if err != nil {
				return errors.New("Error in updating solidwaste payments " + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Update Res =>", paymentUpdateRes)
			basicpaymentUpdate := bson.M{
				"$set": bson.M{
					"solidWasteUserCharge.address":           request.New.Address,
					"solidWasteUserCharge.ownerName":         request.New.OwnerName,
					"solidWasteUserCharge.mobileNo":          request.New.MobileNo,
					"solidWasteUserCharge.shopCategoryId":    request.New.CategoryID,
					"solidWasteUserCharge.shopSubCategoryId": request.New.SubCategoryID,
					"solidWasteUserCharge.dateFrom":          request.New.DateFrom,
				},
			}

			d.Shared.BsonToJSONPrintTag("basicpaymentUpdate =>", basicpaymentUpdate)

			paymentBasicUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC).UpdateMany(ctx.CTX, solidwastePaymentQuery, basicpaymentUpdate)
			if err != nil {
				return errors.New("Error in updating solidwaste payments basics" + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Basic Update Res =>", paymentBasicUpdateRes)

		}
	}

	return nil
}

// FilterSolidWasteReassessmentRequest : ""
func (d *Daos) FilterSolidWasteReassessmentRequest(ctx *models.Context, filter *models.SolidWasteReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefSolidWasteReassessmentRequest, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEREASSESSMENTREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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
	// Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "uniqueId", "ref.requestedUserType", "ref.requestedUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "action.byType", "uniqueId", "ref.actionUserType", "ref.actionUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "new.ref.address.state", "new.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "new.ref.address.district", "new.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "new.ref.address.village", "new.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "new.ref.address.zone", "new.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "new.ref.address.ward", "new.ref.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefSolidWasteReassessmentRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}
