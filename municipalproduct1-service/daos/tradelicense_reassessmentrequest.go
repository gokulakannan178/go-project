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

//SaveBasicPropertyUpdateLog :""
func (d *Daos) SaveTradeLicenseReassessmentRequestUpdate(ctx *models.Context, request *models.TradeLicenseReassessmentRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREASSESSMENTREQUEST).InsertOne(ctx.CTX, request)
	return err
}

//GetSingleTradeLicenseReassessmentRequest : ""
func (d *Daos) GetSingleTradeLicenseReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseReassessmentRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// Lookup

	// user Lookup
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefTradeLicenseReassessmentRequest
	var request *models.RefTradeLicenseReassessmentRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// RejectTradeLicenseReassessmentRequestUpdate : ""
func (d *Daos) AcceptTradeLicenseReassessmentRequestUpdate(ctx *models.Context, accept *models.AcceptTradeLicenseReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEREASSESSMENTREQUESTSTATUSCOMPLETED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectTradeLicenseReassessmentRequestUpdate : ""
func (d *Daos) RejectTradeLicenseReassessmentRequestUpdate(ctx *models.Context, reject *models.RejectTradeLicenseReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEREASSESSMENTREQUESTSTATUSREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTradeLicenseReassessmentRequest : ""
func (d *Daos) FilterTradeLicenseReassessmentRequest(ctx *models.Context, filter *models.TradeLicenseReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefTradeLicenseReassessmentRequest, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREASSESSMENTREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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
	// user Lookup

	// user Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userName", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "userType", "uniqueId", "ref.requestedUserType", "ref.requestedUserType")...)
	// Action Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "userType", "uniqueId", "ref.actionUserType", "ref.actionUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "new.ref.address.state", "new.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "new.ref.address.district", "new.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "new.ref.address.village", "new.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "new.ref.address.zone", "new.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "new.ref.address.ward", "new.ref.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefTradeLicenseReassessmentRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}

func (d *Daos) BasicTradeLicenseReassessmentRequestUpdateToPayments(ctx *models.Context, request *models.RefTradeLicenseReassessmentRequest) error {
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
	tradeLicensePaymentFindQuery := bson.M{
		"status":         constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		"tradeLicenseId": request.TradeLicenseID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("BasicReassessmentRequestUpdateToPayments query =>", tradeLicensePaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Find(ctx.CTX, tradeLicensePaymentFindQuery, nil)
	if err != nil {
		return err
	}
	var tradeLicensePayments []models.RefTradeLicensePayments
	if err = cursor.All(context.TODO(), &tradeLicensePayments); err != nil {
		return err
	}

	if len(tradeLicensePayments) > 0 {
		tnxIDs := []string{}
		for _, v := range tradeLicensePayments {
			tnxIDs = append(tnxIDs, v.TnxID)
		}
		if len(tnxIDs) > 0 {
			tradeLicensePaymentQuery := bson.M{
				"tnxId": bson.M{"$in": tnxIDs},
			}
			paymentUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).UpdateMany(ctx.CTX, tradeLicensePaymentQuery,
				bson.M{
					"$set": bson.M{"address": request.New.Address},
				},
			)
			if err != nil {
				return errors.New("Error in updating tradeLicense payments " + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Update Res =>", paymentUpdateRes)
			paymentBasicUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).UpdateMany(ctx.CTX, tradeLicensePaymentQuery,
				bson.M{
					"$set": bson.M{
						"tradeLicense.address":   request.New.Address,
						"tradeLicense.ownerName": request.New.OwnerName,
						"tradeLicense.mobileNo":  request.New.MobileNo,
					},
				},
			)
			if err != nil {
				return errors.New("Error in updating tradeLicense payments basics" + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Basic Update Res =>", paymentBasicUpdateRes)

		}
	}

	return nil
}
