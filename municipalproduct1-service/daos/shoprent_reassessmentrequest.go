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
func (d *Daos) SaveShoprentReassessmentRequestUpdate(ctx *models.Context, request *models.ShoprentReassessmentRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTREASSESSMENTREQUEST).InsertOne(ctx.CTX, request)
	return err
}

//GetSingleShoprentReassessmentRequest : ""
func (d *Daos) GetSingleShoprentReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefShoprentReassessmentRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// Lookup
	// Lookups
	// mainPipeline = append(mainPipeline, d.PropertyFloorsLookupV2(constants.COLLECTIONPROPERTYFLOOR, "propertyId", "propertyId", "previous.ref.floors", "previous.ref.floors")...)
	// mainPipeline = append(mainPipeline, d.PropertyOwnersLookupV2(constants.COLLECTIONPROPERTYOWNER, "propertyId", "propertyId", "previous.ref.propertyOwner", "previous.ref.propertyOwner")...)
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefShoprentReassessmentRequest
	var request *models.RefShoprentReassessmentRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// RejectShoprentReassessmentRequestUpdate : ""
func (d *Daos) AcceptShoprentReassessmentRequestUpdate(ctx *models.Context, accept *models.AcceptShoprentReassessmentRequestUpdate) error {
	t := time.Now()
	fmt.Println("shoprent accept daos")

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTREASSESSMENTREQUESTSTATUSCOMPLETED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectShoprentReassessmentRequestUpdate : ""
func (d *Daos) RejectShoprentReassessmentRequestUpdate(ctx *models.Context, reject *models.RejectShoprentReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTREASSESSMENTREQUESTSTATUSREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterShoprentReassessmentRequest : ""
func (d *Daos) FilterShoprentReassessmentRequest(ctx *models.Context, filter *models.ShoprentReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefShoprentReassessmentRequest, error) {
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
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTREASSESSMENTREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefShoprentReassessmentRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}

func (d *Daos) BasicShoprentReassessmentRequestUpdateToPayments(ctx *models.Context, request *models.RefShoprentReassessmentRequest) error {
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
	shoprentPaymentFindQuery := bson.M{
		"status":         constants.SHOPRENTPAYMENTSTATUSCOMPLETED,
		"shopRentId":     request.ShoprentID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("BasicReassessmentRequestUpdateToPayments query =>", shoprentPaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Find(ctx.CTX, shoprentPaymentFindQuery, nil)
	if err != nil {
		return err
	}
	var shoprentPayments []models.RefShopRentPayments
	if err = cursor.All(context.TODO(), &shoprentPayments); err != nil {
		return err
	}

	if len(shoprentPayments) > 0 {
		fmt.Println("shoprentPayments is > 0")
		tnxIDs := []string{}
		for _, v := range shoprentPayments {
			tnxIDs = append(tnxIDs, v.TnxID)
			fmt.Println("shoprentPayment tnxId is", v.TnxID)

		}
		if len(tnxIDs) > 0 {
			shoprentPaymentQuery := bson.M{
				"tnxId": bson.M{"$in": tnxIDs},
			}
			d.Shared.BsonToJSONPrintTag("shoprentPayment query =>", shoprentPaymentQuery)
			paymentUpdate := bson.M{
				"$set": bson.M{"address": request.New.Address},
			}
			d.Shared.BsonToJSONPrintTag("paymentUpdate query =>", paymentUpdate)

			paymentUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).UpdateMany(ctx.CTX, shoprentPaymentQuery, paymentUpdate)
			if err != nil {
				return errors.New("Error in updating shoprent payments " + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Update Res =>", paymentUpdateRes)
			basicpaymentUpdate := bson.M{
				"$set": bson.M{
					"shopRent.address":           request.New.Address,
					"shopRent.ownerName":         request.New.OwnerName,
					"shopRent.mobileNo":          request.New.MobileNo,
					"shopRent.shopCategoryId":    request.New.ShopCategoryID,
					"shopRent.shopSubCategoryId": request.New.ShopSubCategoryID,
					"shopRent.sqft":              request.New.Sqft,
					"shopRent.dateFrom":          request.New.DateFrom,
					"shopRent.rentAmount":        request.New.RentAmount,
					"shopRent.guardianName":      request.New.GuardianName,
				},
			}

			d.Shared.BsonToJSONPrintTag("basicpaymentUpdate =>", basicpaymentUpdate)

			paymentBasicUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).UpdateMany(ctx.CTX, shoprentPaymentQuery, basicpaymentUpdate)
			if err != nil {
				return errors.New("Error in updating shoprent payments basics" + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Basic Update Res =>", paymentBasicUpdateRes)

		}
	}

	return nil
}
