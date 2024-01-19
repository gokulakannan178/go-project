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
func (d *Daos) SaveReassessmentRequestUpdate(ctx *models.Context, request *models.ReassessmentRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).InsertOne(ctx.CTX, request)
	return err
}

//GetSingleReassessmentRequest : ""
func (d *Daos) GetSingleReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefReassessmentRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	tempress := new(models.RefReassessmentRequest)
	if err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).FindOne(ctx.CTX, bson.M{"uniqueId": UniqueID}).Decode(&tempress); err != nil {
		return nil, errors.New("Error in find one resaaesment - " + err.Error())

	}
	if len(tempress.New.Ref.ReassessmentFloors) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$unwind": "$new.ref.reassessmentFloors"})
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFLOORTYPE, "new.ref.reassessmentFloors.no", "uniqueId", "new.ref.reassessmentFloors.ref.floorNo", "new.ref.reassessmentFloors.ref.floorNo")...)
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "new.ref.reassessmentFloors.usageType", "uniqueId", "new.ref.reassessmentFloors.ref.usageType", "new.ref.reassessmentFloors.ref.usageType")...)
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "new.ref.reassessmentFloors.constructionType", "uniqueId", "new.ref.reassessmentFloors.ref.constructionType", "new.ref.reassessmentFloors.ref.constructionType")...)
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOCCUMANCYTYPE, "new.ref.reassessmentFloors.occupancyType", "uniqueId", "new.ref.reassessmentFloors.ref.occupancyType", "new.ref.reassessmentFloors.ref.occupancyType")...)
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR, "new.ref.reassessmentFloors.nonResUsageType", "uniqueId", "new.ref.reassessmentFloors.ref.nonresidentialusagefactors", "new.ref.reassessmentFloors.ref.nonresidentialusagefactors")...)
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFLOORRATABLEAREA, "new.ref.reassessmentFloors.ratableAreaType", "uniqueId", "new.ref.reassessmentFloors.ref.floorRatableArea", "new.ref.reassessmentFloors.ref.floorRatableArea")...)

		mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
			"_id":               "$_id",
			"data":              bson.M{"$first": "$$ROOT"},
			"floors":            bson.M{"$push": "$new.ref.reassessmentFloors"},
			"constructiontypes": bson.M{"$push": "$new.ref.reassessmentFloors"},
			"usagetypes":        bson.M{"$push": "$new.ref.reassessmentFloors"},
			"occupancyTypes":    bson.M{"$push": "$new.ref.reassessmentFloors"},
			"nonResUsageType":   bson.M{"$push": "$new.ref.reassessmentFloors"},
			"ratableAreaType":   bson.M{"$push": "$new.ref.reassessmentFloors"},
		}})
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data.new.ref.reassessmentFloors": "$floors"}})
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data.new.ref.reassessmentFloors": "$usagetypes"}})
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data.new.ref.reassessmentFloors": "$occupancyTypes"}})
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data.new.ref.reassessmentFloors": "$nonResUsageType"}})
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data.new.ref.reassessmentFloors": "$ratableAreaType"}})
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data.new.ref.reassessmentFloors": "$constructiontypes"}})
		mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"floors": 0}})
		mainPipeline = append(mainPipeline, bson.M{"$replaceRoot": bson.M{"newRoot": "$data"}})
	}

	// Lookup
	// Previous Address
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookupV2(constants.COLLECTIONPROPERTYFLOOR, "previous.uniqueId", "propertyId", "previous.ref.floors", "previous.ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookupV2(constants.COLLECTIONPROPERTYOWNER, "previous.uniqueId", "propertyId", "previous.ref.propertyOwner", "previous.ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "previous.ref.address.state", "previous.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "previous.ref.address.district", "previous.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "previous.ref.address.village", "previous.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "previous.ref.address.zone", "previous.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "previous.ref.address.ward", "previous.ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "previous.propertyTypeId", "uniqueId", "previous.ref.propertyType", "previous.ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "previous.yoa", "uniqueId", "previous.ref.yoa", "previous.ref.yoa")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "previous.municipalityId", "uniqueId", "previous.ref.municipalType", "previous.ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "previous.uniqueId", "propertyId", "previous.ref.demand", "previous.ref.demand")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "previous.roadTypeId", "uniqueId", "previous.ref.roadType", "previous.ref.roadType")...)

	//New Address
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookupV2(constants.COLLECTIONPROPERTYFLOOR, "new.uniqueId", "propertyId", "new.ref.floors", "new.ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookupV2(constants.COLLECTIONPROPERTYOWNER, "new.uniqueId", "propertyId", "new.ref.propertyOwner", "new.ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "new.ref.address.state", "new.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "new.ref.address.district", "new.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "new.ref.address.village", "new.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "new.ref.address.zone", "new.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "new.ref.address.ward", "new.ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "new.propertyTypeId", "uniqueId", "new.ref.propertyType", "new.ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "new.yoa", "uniqueId", "new.ref.yoa", "new.ref.yoa")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "new.municipalityId", "uniqueId", "new.ref.municipalType", "new.ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "new.uniqueId", "propertyId", "new.ref.demand", "new.ref.demand")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "new.roadTypeId", "uniqueId", "new.ref.roadType", "new.ref.roadType")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.requestedUserType", "ref.requestedUserType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.byType", "type", "ref.requestedUserType", "ref.requestedUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.actionUserType", "ref.actionUserType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.byType", "type", "ref.actionUserType", "ref.actionUserType")...)
	d.Shared.BsonToJSONPrintTag("reassessment request query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var requests []models.RefReassessmentRequest
	var request *models.RefReassessmentRequest
	if err = cursor.All(ctx.CTX, &requests); err != nil {
		return nil, err
	}
	if len(requests) > 0 {
		request = &requests[0]
	}
	return request, nil
}

// RejectReassessmentRequestUpdate : ""
func (d *Daos) AcceptReassessmentRequestUpdate(ctx *models.Context, accept *models.AcceptReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.REASSESSMENTREQUESTSTATUSCOMPLETED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectReassessmentRequestUpdate : ""
func (d *Daos) RejectReassessmentRequestUpdate(ctx *models.Context, reject *models.RejectReassessmentRequestUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.REASSESSMENTREQUESTSTATUSREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterReassessmentRequest : ""
func (d *Daos) FilterReassessmentRequest(ctx *models.Context, filter *models.ReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefReassessmentRequest, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if filter.Regex.PropertyNo != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.Regex.PropertyNo, Options: "xi"}})
		}
		if filter.AppliedRange != nil {
			if filter.AppliedRange.From != nil {
				sd := time.Date(filter.AppliedRange.From.Year(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Day(), 0, 0, 0, 0, filter.AppliedRange.From.Location())
				var ed time.Time
				if filter.AppliedRange.To != nil {
					ed = time.Date(filter.AppliedRange.To.Year(), filter.AppliedRange.To.Month(), filter.AppliedRange.To.Day(), 23, 59, 59, 0, filter.AppliedRange.To.Location())
				} else {
					ed = time.Date(filter.AppliedRange.From.Year(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Day(), 23, 59, 59, 0, filter.AppliedRange.From.Location())
				}
				query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.Address != nil {
			if len(filter.Address.StateCode) > 0 {
				query = append(query, bson.M{"previous.address.stateCode": bson.M{"$in": filter.Address.StateCode}})
			}
			if len(filter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"previous.address.districtCode": bson.M{"$in": filter.Address.DistrictCode}})
			}
			if len(filter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"previous.address.villageCode": bson.M{"$in": filter.Address.VillageCode}})
			}
			if len(filter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"previous.address.zoneCode": bson.M{"$in": filter.Address.ZoneCode}})
			}
			if len(filter.Address.WardCode) > 0 {
				query = append(query, bson.M{"previous.address.wardCode": bson.M{"$in": filter.Address.WardCode}})
			}
			// if filter.IsLocation {
			// 	query = append(query, bson.M{"locSize": bson.M{"$eq": 2}})
			// }
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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
	// Previous Address
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookupV2(constants.COLLECTIONPROPERTYFLOOR, "previous.uniqueId", "propertyId", "previous.ref.floors", "previous.ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookupV2(constants.COLLECTIONPROPERTYOWNER, "previous.uniqueId", "propertyId", "previous.ref.propertyOwner", "previous.ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "previous.ref.address.state", "previous.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "previous.ref.address.district", "previous.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "previous.ref.address.village", "previous.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "previous.ref.address.zone", "previous.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "previous.ref.address.ward", "previous.ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "previous.propertyTypeId", "uniqueId", "previous.ref.propertyType", "previous.ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "previous.yoa", "uniqueId", "previous.ref.yoa", "previous.ref.yoa")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "previous.municipalityId", "uniqueId", "previous.ref.municipalType", "previous.ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "previous.uniqueId", "propertyId", "previous.ref.demand", "previous.ref.demand")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "previous.roadTypeId", "uniqueId", "previous.ref.roadType", "previous.ref.roadType")...)

	//New Address
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookupV2(constants.COLLECTIONPROPERTYFLOOR, "new.uniqueId", "propertyId", "new.ref.floors", "new.ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookupV2(constants.COLLECTIONPROPERTYOWNER, "new.uniqueId", "propertyId", "new.ref.propertyOwner", "new.ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "new.ref.address.state", "new.ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "new.ref.address.district", "new.ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "new.ref.address.village", "new.ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "new.ref.address.zone", "new.ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "new.ref.address.ward", "new.ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "new.propertyTypeId", "uniqueId", "new.ref.propertyType", "new.ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "new.yoa", "uniqueId", "new.ref.yoa", "new.ref.yoa")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "new.municipalityId", "uniqueId", "new.ref.municipalType", "new.ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "new.uniqueId", "propertyId", "new.ref.demand", "new.ref.demand")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "new.roadTypeId", "uniqueId", "new.ref.roadType", "new.ref.roadType")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedUser", "ref.requestedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.requestedUserType", "ref.requestedUserType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.byType", "type", "ref.requestedUserType", "ref.requestedUserType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionUser", "ref.actionUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requester.byType", "name", "ref.actionUserType", "ref.actionUserType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.byType", "type", "ref.actionUserType", "ref.actionUserType")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var request []models.RefReassessmentRequest
	if err = cursor.All(context.TODO(), &request); err != nil {
		return nil, err
	}
	return request, nil
}

// UpdatePropertyReassessmentRequestPropertyID :""
func (d *Daos) UpdatePropertyReassessmentRequestPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONREASSESSMENTREQUEST).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
