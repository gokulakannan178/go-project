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

// SaveUserchargePaymentModeChange :""
func (d *Daos) SaveUserchargePaymentModeChange(ctx *models.Context, request *models.UserchargePaymentModeChangeRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).InsertOne(ctx.CTX, request)
	return err
}

// GetSingleUserchargePaymentModeChange : ""
func (d *Daos) GetSingleUserchargePaymentModeChange(ctx *models.Context, UniqueID string) (*models.RefUserchargePaymentModeChangeRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requested.by", "userName", "ref.requestedBy", "ref.requestedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approved.by", "userName", "ref.approvedBy", "ref.approvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejected.by", "userName", "ref.rejectedBy", "ref.rejectedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requested.byType", "uniqueId", "ref.requestedByType", "ref.requestedByType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefUserchargePaymentModeChangeRequest
	var request *models.RefUserchargePaymentModeChangeRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// RejectUserchargePaymentModeChange : ""
func (d *Daos) AcceptUserchargePaymentModeChange(ctx *models.Context, accept *models.AcceptUserchargePaymentModeChangeRequest) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGEPAYMENTMODECHANGESTATUSCOMPLETED,
		"approved": models.Action{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectUserchargePaymentModeChange : ""
func (d *Daos) RejectUserchargePaymentModeChange(ctx *models.Context, reject *models.RejectUserchargePaymentModeChangeRequest) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGEPAYMENTMODECHANGESTATUSREJECTED,
		"rejected": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserchargePaymentModeChange : ""
func (d *Daos) FilterUserchargePaymentModeChange(ctx *models.Context, filter *models.UserchargePaymentModeChangeRequestFilter, pagination *models.Pagination) ([]models.RefUserchargePaymentModeChangeRequest, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requested.by", "userName", "ref.requestedBy", "ref.requestedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approved.by", "userName", "ref.approvedBy", "ref.approvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejected.by", "userName", "ref.rejectedBy", "ref.rejectedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requested.byType", "uniqueId", "ref.requestedByType", "ref.requestedByType")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefUserchargePaymentModeChangeRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}
func (d *Daos) GetSingleUserchargePaymentModeWithTxtID(ctx *models.Context, txtID string) (*models.UserchargePaymentModeChangeRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})

	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{

	// 	"from": constants.COLLECTIONUserchargePayment,
	// 	"as":   "ref.partPayments",
	// 	"let":  bson.M{"tnxId": "$tnxId"},
	// 	"pipeline": []bson.M{
	// 		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 			{"$eq": []string{"$tnxId", "$$tnxId"}},
	// 			{"$eq": []string{"$status", constants.UserchargePaymentCOMPLETED}},
	// 		}}}},
	// 	},
	// }})

	d.Shared.BsonToJSONPrintTag("property payment query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.UserchargePaymentModeChangeRequest
	var pp *models.UserchargePaymentModeChangeRequest
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

// UpdateUserchargePaymentModeChangePropertyID :""
func (d *Daos) UpdateUserchargePaymentModeChangePropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) UpdateUserchargePayments(ctx *models.Context, propertyPayment *models.UserChargePayments) error {
	selector := bson.M{"tnxId": propertyPayment.TnxID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": propertyPayment}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEPAYMENTS).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
