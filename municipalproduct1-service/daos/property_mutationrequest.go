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

// SavePropertyMutationRequest :""
func (d *Daos) SavePropertyMutationRequest(ctx *models.Context, request *models.PropertyMutationRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYMUTATIONREQUEST).InsertOne(ctx.CTX, request)
	return err
}

// GetSingleropertyMutationRequest : ""
func (d *Daos) GetSinglePropertyMutationRequest(ctx *models.Context, UniqueID string) (*models.RefPropertyMutationRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.requestedUserType", "ref.requestedUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.actionUserType", "ref.actionUserType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYMUTATIONREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefPropertyMutationRequest
	var request *models.RefPropertyMutationRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// AcceptPropertyMutationRequestUpdate : ""
func (d *Daos) AcceptPropertyMutationRequestUpdate(ctx *models.Context, accept *models.AcceptPropertyMutationRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMUTATIONREQUESTSTATUSCOMPLETED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYMUTATIONREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectPropertyMutationRequest : ""
func (d *Daos) RejectPropertyMutationRequestUpdate(ctx *models.Context, reject *models.RejectPropertyMutationRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMUTATIONREQUESTSTATUSREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYMUTATIONREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyMutationRequest : ""
func (d *Daos) FilterPropertyMutationRequest(ctx *models.Context, filter *models.PropertyMutationRequestFilter, pagination *models.Pagination) ([]models.RefPropertyMutationRequest, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYMUTATIONREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("property mutation request query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYMUTATIONREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefPropertyMutationRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}

// SaveMutatedProperty :""
func (d *Daos) SaveMutatedProperty(ctx *models.Context, mutatedProperty *models.MutatedProperty) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMUTATEDPROPERTY).InsertOne(ctx.CTX, mutatedProperty)
	return err
}

// GetSingleMutatedProperty : ""
func (d *Daos) GetSingleMutatedProperty(ctx *models.Context, UniqueID string) (*models.RefMutatedProperty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "parentId", "uniqueId", "ref.property", "ref.property")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMUTATEDPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefMutatedProperty
	var request *models.RefMutatedProperty
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// FilterMutatedProperty : ""
func (d *Daos) FilterMutatedProperty(ctx *models.Context, filter *models.MutatedPropertyFilter, pagination *models.Pagination) ([]models.RefMutatedProperty, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.ParentID) > 0 {
			query = append(query, bson.M{"parentId": bson.M{"$in": filter.ParentID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMUTATEDPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "parentId", "uniqueId", "ref.property", "ref.property")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property mutation request query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMUTATEDPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var mutatedProperties []models.RefMutatedProperty
	if err = cursor.All(context.TODO(), &mutatedProperties); err != nil {
		return nil, err
	}
	return mutatedProperties, nil
}

// UpdatePropertyEndDate : ""
// func (d *Daos) UpdatePropertyEndDate(ctx *models.Context, UniqueID string) error {
func (d *Daos) UpdatePropertyEndDate(ctx *models.Context, UniqueID string, On *time.Time) error {
	fmt.Println("UniqueId ============>", UniqueID)
	selector := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"endDate": On, "status": constants.PROPERTYSTATUSMUTATED}}
	// updateInterface := bson.M{"$set": bson.M{"endDate": &t, "status": constants.PROPERTYSTATUSMUTATED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// UpdatePropertyMutationRequestPropertyID :""
func (d *Daos) UpdatePropertyMutationRequestPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYMUTATIONREQUEST).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
