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

// SaveMobileTowerDeleteRequest :""
func (d *Daos) SaveMobileTowerDeleteRequest(ctx *models.Context, request *models.MobileTowerDeleteRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERDELETEREQUEST).InsertOne(ctx.CTX, request)
	return err
}

// GetSingleropertyDeleteRequest : ""
func (d *Daos) GetSingleMobileTowerDeleteRequest(ctx *models.Context, UniqueID string) (*models.RefMobileTowerDeleteRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.requestedUserType", "ref.requestedUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.actionUserType", "ref.actionUserType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERDELETEREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefMobileTowerDeleteRequest
	var request *models.RefMobileTowerDeleteRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// AcceptMobileTowerDeleteRequestUpdate : ""
func (d *Daos) AcceptMobileTowerDeleteRequestUpdate(ctx *models.Context, accept *models.AcceptMobileTowerDeleteRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERDELETEREQUESTSTATUSCOMPLETED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERDELETEREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectMobileTowerDeleteRequest : ""
func (d *Daos) RejectMobileTowerDeleteRequestUpdate(ctx *models.Context, reject *models.RejectMobileTowerDeleteRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERDELETEREQUESTSTATUSREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERDELETEREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMobileTowerDeleteRequest : ""
func (d *Daos) FilterMobileTowerDeleteRequest(ctx *models.Context, filter *models.MobileTowerDeleteRequestFilter, pagination *models.Pagination) ([]models.RefMobileTowerDeleteRequest, error) {
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
	fmt.Println("sortBy====>", filter.SortBy)
	fmt.Println("sortOrder====>", filter.SortOrder)
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERDELETEREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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

	// Lookup

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.requestedUserType", "ref.requestedUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.actionUserType", "ref.actionUserType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property delete request query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERDELETEREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefMobileTowerDeleteRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}

// UpdateMobileTowerStatusDeletedV2 : ""

func (d *Daos) UpdateMobileTowerStatusDeletedV2(ctx *models.Context, UniqueID string) error {
	fmt.Println("UniqueId ============>", UniqueID)
	selector := bson.M{"uniqueId": UniqueID}
	updateInterface := bson.M{"$set": bson.M{"status": constants.SHOPRENTSTATUSDELETED}}

	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
