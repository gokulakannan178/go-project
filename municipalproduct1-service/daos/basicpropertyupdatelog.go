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

//SaveBasicPropertyUpdateLog :""
func (d *Daos) SaveBasicPropertyUpdateLog(ctx *models.Context, bpul *models.BasicPropertyUpdateLog) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).InsertOne(ctx.CTX, bpul)
	return err
}

// AcceptBasicPropertyUpdate : ""
func (d *Daos) AcceptBasicPropertyUpdate(ctx *models.Context, accept *models.AcceptBasicPropertyUpdate) error {

	t := time.Now()
	query := bson.M{"uniqueId": accept.UniqueID}
	var bpul *models.BasicPropertyUpdateLog
	err := ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).FindOne(ctx.CTX, query).Decode(&bpul)
	if err != nil {
		return errors.New("Not able to find the request" + err.Error())
	}
	if bpul == nil {
		return errors.New("Request in nil")
	}
	bpu := new(models.BasicPropertyUpdate)
	bpu.PropertyID = bpul.PropertyID
	bpu.Address = bpul.New.Address
	bpu.Owner = bpul.New.Owner
	bpu.UserName = bpul.Requester.By
	bpu.UserType = bpul.Requester.ByType
	bpu.Proof = bpul.Proof
	bpu.Remarks = bpul.Requester.Remarks
	err = d.BasicUpdateProperty(ctx, bpu)
	if err != nil {
		return errors.New("Error in upddating Property" + err.Error())
	}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYBASICUPDATELOGACCEPTED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectBasicPropertyUpdate : ""
func (d *Daos) RejectBasicPropertyUpdate(ctx *models.Context, reject *models.RejectBasicPropertyUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYBASICUPDATELOGREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterBasicPropertyUpdate : ""
func (d *Daos) FilterBasicPropertyUpdate(ctx *models.Context, filter *models.FilterBasicPropertyUpdate, pagination *models.Pagination) ([]models.RefBasicPropertyUpdateLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if len(filter.UniqueId) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueId}})
		}
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"userType": bson.M{"$in": filter.UserType}})
		}
		if len(filter.Approver) > 0 {
			query = append(query, bson.M{"action.by": bson.M{"$in": filter.ApproverType}})
		}
		if len(filter.ApproverType) > 0 {
			query = append(query, bson.M{"action.byType": bson.M{"$in": filter.ApproverType}})
		}
		if filter.SearchText.PropertyID != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.SearchText.PropertyID, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedBy", "ref.requestedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "uniqueId", "ref.requestedByType", "ref.requestedByType")...)

	//Old Address Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "ref.previous.address.state", "ref.previous.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "ref.previous.address.district", "ref.previous.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "ref.previous.address.village", "ref.previous.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "ref.previous.address.zone", "ref.previous.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "ref.previous.address.ward", "ref.previous.address.ward")...)

	//New Address
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "ref.new.address.state", "ref.new.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "ref.address.district", "ref.new.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "ref.address.village", "ref.new.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "ref.address.zone", "ref.new.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "ref.new.address.ward", "ref.new.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefBasicPropertyUpdateLog
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}

func (d *Daos) GetSingleBasicPropertyUpdateLog(ctx *models.Context, uniqueID string) (*models.RefBasicPropertyUpdateLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})

	//Old Address Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "ref.previous.address.state", "ref.previous.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "ref.previous.address.district", "ref.previous.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "ref.previous.address.village", "ref.previous.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "ref.previous.address.zone", "ref.previous.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "ref.previous.address.ward", "ref.previous.address.ward")...)

	//New Address
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "ref.new.address.state", "ref.new.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "ref.address.district", "ref.new.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "ref.address.village", "ref.new.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "ref.address.zone", "ref.new.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "ref.new.address.ward", "ref.new.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rbpuls []models.RefBasicPropertyUpdateLog
	var rbpul *models.RefBasicPropertyUpdateLog
	if err = cursor.All(ctx.CTX, &rbpuls); err != nil {
		return nil, err
	}
	if len(rbpuls) > 0 {
		rbpul = &rbpuls[0]
	}
	return rbpul, nil

}

func (d *Daos) BasicPropertyUpdateToPayments(ctx *models.Context, rbpul *models.RefBasicPropertyUpdateLog) error {
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
	propertyPaymentFindQuery := bson.M{
		"status":         constants.PROPERTYPAYMENTCOMPLETED,
		"propertyId":     rbpul.PropertyID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("BasicPropertyUpdateToPayments query =>", propertyPaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Find(ctx.CTX, propertyPaymentFindQuery, nil)
	if err != nil {
		return err
	}
	var propertyPayments []models.RefPropertyPayment
	if err = cursor.All(context.TODO(), &propertyPayments); err != nil {
		return err
	}

	if len(propertyPayments) > 0 {
		tnxIDs := []string{}
		for _, v := range propertyPayments {
			tnxIDs = append(tnxIDs, v.TnxID)
		}
		if len(tnxIDs) > 0 {
			propertyPaymentQuery := bson.M{
				"tnxId": bson.M{"$in": tnxIDs},
			}
			paymentUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateMany(ctx.CTX, propertyPaymentQuery,
				bson.M{
					"$set": bson.M{"address": rbpul.New.Address},
				},
			)
			if err != nil {
				return errors.New("Error in updating property payments " + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Update Res =>", paymentUpdateRes)
			paymentBasicUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateMany(ctx.CTX, propertyPaymentQuery,
				bson.M{
					"$set": bson.M{
						"property.address": rbpul.New.Address,
						"owners":           []models.PropertyOwner{rbpul.New.Owner},
					},
				},
			)
			if err != nil {
				return errors.New("Error in updating property payments basics" + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Basic Update Res =>", paymentBasicUpdateRes)

		}
	}

	return nil
}

//BasicPropertyUpdateGetPaymentsToBeUpdated : ""
func (d *Daos) BasicPropertyUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbpul *models.RefBasicPropertyUpdateLog) ([]models.RefPropertyPayment, error) {
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
	propertyPaymentFindQuery := bson.M{
		"status":         constants.PROPERTYPAYMENTCOMPLETED,
		"propertyId":     rbpul.PropertyID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", propertyPaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Find(ctx.CTX, propertyPaymentFindQuery, nil)
	if err != nil {
		return nil, err
	}
	var propertyPayments []models.RefPropertyPayment
	if err = cursor.All(context.TODO(), &propertyPayments); err != nil {
		return nil, err
	}

	return propertyPayments, nil
}

// BasicReassessmentRequestUpdateToPayments : ""
func (d *Daos) BasicReassessmentRequestUpdateToPayments(ctx *models.Context, request *models.RefReassessmentRequest) error {
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
	propertyPaymentFindQuery := bson.M{
		"status":         constants.PROPERTYPAYMENTCOMPLETED,
		"propertyId":     request.PropertyID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("BasicReassessmentRequestUpdateToPayments query =>", propertyPaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Find(ctx.CTX, propertyPaymentFindQuery, nil)
	if err != nil {
		return err
	}
	var propertyPayments []models.RefPropertyPayment
	if err = cursor.All(context.TODO(), &propertyPayments); err != nil {
		return err
	}

	if len(propertyPayments) > 0 {
		tnxIDs := []string{}
		for _, v := range propertyPayments {
			tnxIDs = append(tnxIDs, v.TnxID)
		}
		if len(tnxIDs) > 0 {
			propertyPaymentQuery := bson.M{
				"tnxId": bson.M{"$in": tnxIDs},
			}
			paymentUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateMany(ctx.CTX, propertyPaymentQuery,
				bson.M{
					"$set": bson.M{"address": request.New.Address},
				},
			)
			if err != nil {
				return errors.New("Error in updating property payments " + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Update Res =>", paymentUpdateRes)
			paymentBasicUpdateRes, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENTBASIC).UpdateMany(ctx.CTX, propertyPaymentQuery,
				bson.M{
					"$set": bson.M{
						"property.address": request.New.Address,
						"owners":           []models.PropertyOwner{request.New.Owner[0]},
					},
				},
			)
			if err != nil {
				return errors.New("Error in updating property payments basics" + err.Error())
			}
			d.Shared.BsonToJSONPrintTag("Payment Basic Update Res =>", paymentBasicUpdateRes)

		}
	}

	return nil
}

// UpdateBasicPropertyUpdateLogPropertyID :""
func (d *Daos) UpdateBasicPropertyUpdateLogPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBASICPROPERTYUPDATELOG).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
