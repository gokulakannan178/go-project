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
func (d *Daos) SaveMobileTowerReassessmentRequestUpdate(ctx *models.Context, request *models.MobileTowerReassessmentRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREASSESSMENTREQUEST).InsertOne(ctx.CTX, request)
	return err
}

//GetSingleMobileTowerReassessmentRequest : ""
func (d *Daos) GetSingleMobileTowerReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefMobileTowerReassessmentRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "previous.ref.address.state", "previous.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "previous.ref.address.district", "previous.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "previous.ref.address.village", "previous.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "previous.ref.address.zone", "previous.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "previous.ref.address.ward", "previous.ref.address.ward")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefMobileTowerReassessmentRequest
	var request *models.RefMobileTowerReassessmentRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// RejectMobileTowerReassessmentRequestUpdate : ""
func (d *Daos) AcceptMobileTowerReassessmentRequestUpdate(ctx *models.Context, accept *models.AcceptMobileTowerReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERREASSESSMENTREQUESTSTATUSCOMPLETED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectMobileTowerReassessmentRequestUpdate : ""
func (d *Daos) RejectMobileTowerReassessmentRequestUpdate(ctx *models.Context, reject *models.RejectMobileTowerReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERREASSESSMENTREQUESTSTATUSREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMobileTowerReassessmentRequest : ""
func (d *Daos) FilterMobileTowerReassessmentRequest(ctx *models.Context, filter *models.MobileTowerReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefMobileTowerReassessmentRequest, error) {
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
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREASSESSMENTREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefMobileTowerReassessmentRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}

func (d *Daos) BasicMobileTowerReassessmentRequestUpdateToPayments(ctx *models.Context, request *models.RefMobileTowerReassessmentRequest) error {
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
	mobileTowerPaymentFindQuery := bson.M{
		"status":         constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		"mobileTowerId":  request.MobileTowerID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("BasicReassessmentRequestUpdateToPayments query =>", mobileTowerPaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).Find(ctx.CTX, mobileTowerPaymentFindQuery, nil)
	if err != nil {
		return err
	}
	var mobileTowerPayments []models.RefMobileTowerPayments
	if err = cursor.All(context.TODO(), &mobileTowerPayments); err != nil {
		return err
	}

	if len(mobileTowerPayments) > 0 {
		tnxIDs := []string{}
		for _, v := range mobileTowerPayments {
			tnxIDs = append(tnxIDs, v.TnxID)
		}
		if len(tnxIDs) > 0 {
			mobileTowerPaymentQuery := bson.M{
				"tnxId": bson.M{"$in": tnxIDs},
			}
			paymentUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateMany(ctx.CTX, mobileTowerPaymentQuery,
				bson.M{
					"$set": bson.M{"address": request.New.Address},
				},
			)
			if err != nil {
				return errors.New("Error in updating mobileTower payments " + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Update Res =>", paymentUpdateRes)
			paymentBasicUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, mobileTowerPaymentQuery,
				bson.M{
					"$set": bson.M{
						"mobileTower.address":   request.New.Address,
						"mobileTower.ownerName": request.New.OwnerName,
						"mobileTower.mobileNo":  request.New.MobileNo,
					},
				},
			)
			if err != nil {
				return errors.New("Error in updating mobileTower payments basics" + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Basic Update Res =>", paymentBasicUpdateRes)

		}
	}

	return nil
}
