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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveProperty :""
func (d *Daos) SaveProperty(ctx *models.Context, property *models.Property, collectionName string) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).InsertOne(ctx.SC, property)
	return err
}

//SaveProperty :""
func (d *Daos) SavePropertyV2(ctx *models.Context, db *mongo.Database, sc *mongo.SessionContext, property *models.Property) error {
	_, err := db.Collection(constants.COLLECTIONPROPERTY).InsertOne(ctx.CTX, property)
	return err
}

//SaveProperty :""
func (d *Daos) SavePropertyV3(db *mongo.Database, sc mongo.SessionContext, property *models.Property) error {
	_, err := db.Collection(constants.COLLECTIONPROPERTY).InsertOne(sc, property)
	return err
}

// PropertyFloorsLookup :
func (d *Daos) PropertyFloorsLookup(from string, localField string, foreignField string, as string, addField string) []bson.M {
	var Lookups []bson.M
	var pipeline []bson.M
	pipeline = []bson.M{
		bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
			bson.M{"$eq": []string{"$status", constants.PROPERTYFLOORSTATUSACTIVE}},
		}}}},
	}
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "usageType", "uniqueId", "ref.usageType", "ref.usageType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "constructionType", "uniqueId", "ref.constructionType", "ref.constructionType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONOCCUMANCYTYPE, "occupancyType", "uniqueId", "ref.occupancyType", "ref.occupancyType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR, "nonResUsageType", "uniqueId", "ref.nonResUsageType", "ref.nonResUsageType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONFLOORRATABLEAREA, "ratableAreaType", "uniqueId", "ref.floorRatableArea", "ref.floorRatableArea")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONFLOORTYPE, "no", "uniqueId", "ref.floorNo", "ref.floorNo")...)

	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from":     from,
			"as":       as,
			"let":      bson.M{"propertyId": "$uniqueId"},
			"pipeline": pipeline,
		},
		})
	return Lookups
}

// PropertyFloorsLookup :
func (d *Daos) PropertyFloorsLookupV2(from string, localField string, foreignField string, as string, addField string) []bson.M {
	var Lookups []bson.M
	var pipeline []bson.M
	pipeline = []bson.M{
		bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
			bson.M{"$eq": []string{"$status", constants.PROPERTYFLOORSTATUSACTIVE}},
		}}}},
	}
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "usageType", "uniqueId", "ref.usageType", "ref.usageType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "constructionType", "uniqueId", "ref.constructionType", "ref.constructionType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONOCCUMANCYTYPE, "occupancyType", "uniqueId", "ref.occupancyType", "ref.occupancyType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR, "nonResUsageType", "uniqueId", "ref.nonResUsageType", "ref.nonResUsageType")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONFLOORRATABLEAREA, "ratableAreaType", "uniqueId", "ref.floorRatableArea", "ref.floorRatableArea")...)
	pipeline = append(pipeline, d.CommonLookup(constants.COLLECTIONFLOORTYPE, "no", "uniqueId", "ref.floorNo", "ref.floorNo")...)

	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from":     from,
			"as":       as,
			"let":      bson.M{"propertyId": "$" + localField},
			"pipeline": pipeline,
		},
		})
	return Lookups
}

// PropertyOwnersLookup : ""
func (d *Daos) PropertyOwnersLookup(from string, localField string, foreignField string, as string, addField string) []bson.M {
	var Lookups []bson.M
	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from": from,
			"as":   as,
			"let":  bson.M{"propertyId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
					bson.M{"$eq": []string{"$status", constants.PROPERTYOWNERSTATUSACTIVE}},
				}}}},
			},
		},
		})
	return Lookups
}

// PropertyOwnersLookup : ""
func (d *Daos) PropertyOwnersLookupV2(from string, localField string, foreignField string, as string, addField string) []bson.M {
	var Lookups []bson.M
	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from": from,
			"as":   as,
			"let":  bson.M{"propertyId": "$" + localField},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
					bson.M{"$eq": []string{"$status", constants.PROPERTYOWNERSTATUSACTIVE}},
				}}}},
			},
		},
		})
	return Lookups
}

// EstimatedPropertyFloorsLookup : ""
func (d *Daos) EstimatedPropertyFloorsLookup(from string, localField string, foreignField string, as string, addField string) []bson.M {
	var Lookups []bson.M
	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from": from,
			"as":   as,
			"let":  bson.M{"propertyId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$estimatedId", "$$propertyId"}},
					bson.M{"$eq": []string{"$status", constants.ESTIMATEDPROPERTYFLOORSTATUSACTIVE}},
				}}}},
			},
		},
		})
	return Lookups
}

//GetSingleProperty : ""
func (d *Daos) GetSingleProperty(ctx *models.Context, UniqueID string) (*models.RefProperty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	//Lookups
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "ownerId", "uniqueId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "yoa", "uniqueId", "ref.yoa", "ref.yoa")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGECATEGORY, "userCharge.categoryId", "uniqueId", "ref.userchargecategory", "ref.userchargecategory")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGERATEMASTER, "userCharge.categoryId", "categoryId", "ref.UserChargeCategory", "ref.UserChargeCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userCharge.createdBy.by", "userName", "ref.userChargeCreator", "ref.userChargeCreator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userchargeverifiedInfo.by", "userName", "ref.userchargeactivator", "ref.userchargeactivator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userchargerejectedInfo.by", "userName", "ref.userchargerejector", "ref.userchargerejector")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONUSERCHARGERATEMASTER,
			"as":   "ref.userchargeratemaster",
			"let":  bson.M{"categoryId": "$userCharge.categoryId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$categoryId", "$$categoryId"}},
					{"$eq": []string{"$status", "Active"}},
				}}},
				},
				{"$sort": bson.M{"doe": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.userchargeratemaster": bson.M{"$arrayElemAt": []interface{}{"$ref.userchargeratemaster", 0}}}})
	//Aggregation
	//d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	var property *models.RefProperty
	if err = cursor.All(ctx.CTX, &propertys); err != nil {
		return nil, err
	}
	if len(propertys) > 0 {
		property = &propertys[0]
	}
	return property, nil
}

//UpdateProperty : ""
func (d *Daos) UpdateProperty(ctx *models.Context, property *models.Property) error {
	selector := bson.M{"uniqueId": property.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": property, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPropertyQuery : ""
func (d *Daos) FilterPropertyQuery(ctx *models.Context, propertyfilter *models.PropertyFilter) []bson.M {
	query := []bson.M{}
	if propertyfilter != nil {

		if len(propertyfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyfilter.Status}})
		}
		if len(propertyfilter.UserChargeStatus) > 0 {
			query = append(query, bson.M{"userCharge.status": bson.M{"$in": propertyfilter.UserChargeStatus}})
		}
		if len(propertyfilter.Type) > 0 {
			query = append(query, bson.M{"propertyTypeId": bson.M{"$in": propertyfilter.Type}})
		}
		if len(propertyfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": propertyfilter.UniqueID}})
		}
		if len(propertyfilter.IsMatched) > 0 {
			query = append(query, bson.M{"isMatched": bson.M{"$in": propertyfilter.IsMatched}})
		}

		if len(propertyfilter.OldHoldingNumber) > 0 {
			query = append(query, bson.M{"oldHoldingNumber": bson.M{"$in": propertyfilter.OldHoldingNumber}})
		}
		if len(propertyfilter.IsUserCharge) > 0 {
			query = append(query, bson.M{"userCharge.isUserCharge": bson.M{"$in": propertyfilter.IsUserCharge}})
		}
		if propertyfilter.IsGeoTagged == "Yes" {
			query = append(query, bson.M{"address.isGeoTagged": bson.M{"$eq": "Yes"}})

		}
		if propertyfilter.IsGeoTagged == "No" {
			query = append(query, bson.M{"address.isGeoTagged": bson.M{"$ne": "Yes"}})

		}
		if propertyfilter.Address != nil {
			if len(propertyfilter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": propertyfilter.Address.StateCode}})
			}
			if len(propertyfilter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": propertyfilter.Address.DistrictCode}})
			}
			if len(propertyfilter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": propertyfilter.Address.VillageCode}})
			}
			if len(propertyfilter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": bson.M{"$in": propertyfilter.Address.ZoneCode}})
			}
			if len(propertyfilter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": bson.M{"$in": propertyfilter.Address.WardCode}})
			}
			if propertyfilter.IsLocation {
				query = append(query, bson.M{"locSize": bson.M{"$eq": 2}})
			}
		}

		//Regex
		if propertyfilter.Regex.PropertyNo != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: propertyfilter.Regex.PropertyNo, Options: "xi"}})
		}
		if propertyfilter.Regex.ApplicationNo != "" {
			query = append(query, bson.M{"applicationNo": primitive.Regex{Pattern: propertyfilter.Regex.ApplicationNo, Options: "xi"}})
		}
		if propertyfilter.AppliedRange != nil {
			if propertyfilter.AppliedRange.From != nil {
				sd := time.Date(propertyfilter.AppliedRange.From.Year(), propertyfilter.AppliedRange.From.Month(), propertyfilter.AppliedRange.From.Day(), 0, 0, 0, 0, propertyfilter.AppliedRange.From.Location())
				var ed time.Time
				if propertyfilter.AppliedRange.To != nil {
					ed = time.Date(propertyfilter.AppliedRange.To.Year(), propertyfilter.AppliedRange.To.Month(), propertyfilter.AppliedRange.To.Day(), 23, 59, 59, 0, propertyfilter.AppliedRange.To.Location())
				} else {
					ed = time.Date(propertyfilter.AppliedRange.From.Year(), propertyfilter.AppliedRange.From.Month(), propertyfilter.AppliedRange.From.Day(), 23, 59, 59, 0, propertyfilter.AppliedRange.From.Location())
				}
				query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if propertyfilter.Regex.OwnerName != "" {

			propertyIds, err := d.GetPropertyIDsWithOwnerNames(ctx, propertyfilter.Regex.OwnerName)
			if err != nil {
				log.Println("ERR IN GETING - Property IDs WithOwner Names " + err.Error())
			} else {
				if len(propertyIds) > 0 {
					fmt.Println("got Property Ids - ", propertyIds)
					query = append(query, bson.M{"uniqueId": bson.M{"$in": propertyIds}})
				}
			}
		}
		if propertyfilter.Regex.Mobile != "" {

			propertyIds, err := d.GetPropertyIDsWithMobileNos(ctx, propertyfilter.Regex.Mobile)
			if err != nil {
				log.Println("ERR IN GETING - Property IDs With Owner Mobile No " + err.Error())
			} else {
				if len(propertyIds) > 0 {
					fmt.Println("got Property Ids - ", propertyIds)
					query = append(query, bson.M{"uniqueId": bson.M{"$in": propertyIds}})
				}
			}
		}

	}
	return query
}

//FilterProperty : ""
func (d *Daos) FilterProperty(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) ([]models.RefProperty, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyfilter != nil {
		if propertyfilter.IsLocation {
			mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"address.location.coordinates": bson.M{"$exists": true}}})
			mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"locSize": bson.M{"$size": "$address.location.coordinates"}}})
		}
	}

	query = d.FilterPropertyQuery(ctx, propertyfilter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyfilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.SortBy: propertyfilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	if propertyfilter.DemandSortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.DemandSortBy: propertyfilter.DemandSortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"summary.toPay.total.totalTax": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYWALLET, "uniqueId", "propertyId", "ref.wallet", "ref.wallet")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "log.by.id", "userName", "ref.activator", "ref.activator")...)

	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGERATEMASTER, "userCharge.categoryId", "categoryId", "ref.UserChargeCategory", "ref.UserChargeCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGECATEGORY, "userCharge.categoryId", "uniqueId", "ref.userchargecategory", "ref.userchargecategory")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGERATEMASTER, "userCharge.categoryId", "categoryId", "ref.UserChargeCategory", "ref.UserChargeCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userchargeverifiedInfo.by", "userName", "ref.userchargeactivator", "ref.userchargeactivator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userchargerejectedInfo.by", "userName", "ref.userchargerejector", "ref.userchargerejector")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userCharge.createdBy.by", "userName", "ref.userChargeCreator", "ref.userChargeCreator")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONUSERCHARGERATEMASTER,
			"as":   "ref.userchargeratemaster",
			"let":  bson.M{"categoryId": "$userCharge.categoryId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$categoryId", "$$categoryId"}},
					{"$eq": []string{"$status", "Active"}},
				}}},
				},
				{"$sort": bson.M{"doe": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.userchargeratemaster": bson.M{"$arrayElemAt": []interface{}{"$ref.userchargeratemaster", 0}}}})

	if propertyfilter.OmitZeroDemand {
		query = append(query, bson.M{"ref.demand.total.totalTax": bson.M{"$gte": 0}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}

// PropertyWiseDemandandCollectionJSON : ""
func (d *Daos) PropertyWiseDemandandCollectionJSON(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) ([]models.RefProperty, error) {
	resFYs, err := d.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("error in getting current financial year" + err.Error())
	}
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyfilter != nil {
		if propertyfilter.IsLocation {
			mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"address.location.coordinates": bson.M{"$exists": true}}})
			mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"locSize": bson.M{"$size": "$address.location.coordinates"}}})
		}
	}

	query = d.FilterPropertyQuery(ctx, propertyfilter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyfilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.SortBy: propertyfilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	if propertyfilter != nil {
		if !propertyfilter.RemoveLookup.PropertyOwner {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
		}
		if !propertyfilter.RemoveLookup.State {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
		}
		if !propertyfilter.RemoveLookup.District {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
		}
		if !propertyfilter.RemoveLookup.Village {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
		}
		if !propertyfilter.RemoveLookup.Zone {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
		}
		if !propertyfilter.RemoveLookup.Ward {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
		}
		if !propertyfilter.RemoveLookup.PropertyFloor {
			mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
		}
		if !propertyfilter.RemoveLookup.PropertyType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
		}
		if !propertyfilter.RemoveLookup.RoadType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
		}
		if !propertyfilter.RemoveLookup.MunicipalType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
		}
		if !propertyfilter.RemoveLookup.Wallet {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYWALLET, "uniqueId", "propertyId", "ref.wallet", "ref.wallet")...)
		}
		if !propertyfilter.RemoveLookup.User {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.user", "ref.user")...)
		}
		if !propertyfilter.RemoveLookup.Activator {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "log.by.id", "userName", "ref.activator", "ref.activator")...)
		}
	}

	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)
	if propertyfilter.OmitZeroDemand {
		query = append(query, bson.M{"ref.demand.total.totalTax": bson.M{"$gte": 0}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENTFY,
		"as":   "ref.collections",
		"let":  bson.M{"varUniqueId": "$uniqueId"},
		"pipeline": []bson.M{bson.M{
			"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$eq": []string{"$propertyId", "$$varUniqueId"}},
			}}},
		},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"arrearTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"currentTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"arrearPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"currentPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"arrearRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"currentRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"arrearAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"totalTax":    bson.M{"$sum": "$fy.totalTax"},
				"otherDemand": bson.M{"$sum": "$fy.otherDemand"},
			}},
		},
	},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.collections": bson.M{"$arrayElemAt": []interface{}{"$ref.collections", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "ref.propertyPayments",
		"let":  bson.M{"varUniqueId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$eq": []string{"$$varUniqueId", "$propertyId"}},
			}}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"rebate":  bson.M{"$sum": "$demand.rebate"},
				"formFee": bson.M{"$sum": "$demand.formFee"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.propertyPayments": bson.M{
				"$arrayElemAt": []interface{}{"$ref.propertyPayments", 0},
			},
		},
	})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}

// ZoneAndWardWiseReport : ""
func (d *Daos) ZoneAndWardWiseReport(ctx *models.Context, filter *models.ZoneAndWardWiseReportFilter, pagination *models.Pagination) ([]models.ZoneAndWardWiseReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if filter.Address != nil {
			if len(filter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"villageCode": bson.M{"$in": filter.Address.VillageCode}})
			}
			if len(filter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"code": bson.M{"$in": filter.Address.ZoneCode}})
			}
		}
	}
	// Adding $match to filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONWARD,
			"as":   "wards",
			"let":  bson.M{"zoneId": "$code"},
			"pipeline": []bson.M{
				bson.M{
					"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$status", constants.WARDSTATUSACTIVE}},
						bson.M{"$eq": []string{"$zoneCode", "$$zoneId"}},
					}}},
				},
				bson.M{"$lookup": bson.M{
					"from": constants.COLLECTIONPROPERTYPAYMENT,
					"as":   "payments",
					"let":  bson.M{"zoneId": "$$zoneId", "wardId": "$code"},
					"pipeline": []bson.M{
						bson.M{"$match": func() bson.M {
							var propertyPaymentQuery []bson.M
							propertyPaymentQuery = append(propertyPaymentQuery, bson.M{"$eq": []string{"$address.wardCode", "$$wardId"}})
							propertyPaymentQuery = append(propertyPaymentQuery, bson.M{"$eq": []string{"$address.zoneCode", "$$zoneId"}})
							propertyPaymentQuery = append(propertyPaymentQuery, bson.M{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}})
							if filter != nil {
								if filter.DateRange != nil {
									if filter.DateRange.From != nil {
										sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
										ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
										if filter.DateRange.To != nil {
											ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
										}
										propertyPaymentQuery = append(propertyPaymentQuery, bson.M{"$gte": []interface{}{"$completionDate", sd}})
										propertyPaymentQuery = append(propertyPaymentQuery, bson.M{"$lte": []interface{}{"$completionDate", ed}})

									}
								}
							}
							if len(propertyPaymentQuery) > 0 {
								return bson.M{"$expr": bson.M{"$and": propertyPaymentQuery}}
							}
							return bson.M{}

						}(),
						},
						bson.M{"$group": bson.M{
							"_id": "$propertyId",
							"TC":  bson.M{"$sum": "$details.amount"},
						}},
						bson.M{"$group": bson.M{
							"_id":              nil,
							"totalCollections": bson.M{"$sum": "$TC"},
							"totalProperties":  bson.M{"$sum": 1},
						}},
					}},
				},
				bson.M{"$addFields": bson.M{"payments": bson.M{"$arrayElemAt": []interface{}{"$payments", 0}}}},
			},
		}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONZONE).CountDocuments(ctx.CTX, func() bson.M {
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

	// Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONZONE).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var report []models.ZoneAndWardWiseReport
	if err := cursor.All(ctx.CTX, &report); err != nil {
		return nil, err
	}

	return report, nil

}

// WardwiseDemandandCollection : ""
func (d *Daos) WardwiseDemandandCollection(ctx *models.Context, propertyfilter *models.PropertyWardwiseDemandFilter, pagination *models.Pagination) ([]models.WardwiseDemandandCollection, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyfilter != nil {
		if len(propertyfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyfilter.Status}})
		}
		if len(propertyfilter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": propertyfilter.Zone}})
		}
		if len(propertyfilter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": propertyfilter.Ward}})
		}

	}

	// Adding $match to filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyfilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.SortBy: propertyfilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{

			"from": "properties",
			"as":   "properties",
			"let":  bson.M{"code": "$code"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$address.wardCode", "$$code"}},
					bson.M{"$in": []interface{}{"$status", []string{"Active", "Init"}}},
				}}}},
				bson.M{"$group": bson.M{
					"_id":                    "$address.wardCode",
					"totalDemandArrear":      bson.M{"$sum": "$demand.arrear"},
					"totalDemandCurrent":     bson.M{"$sum": "$demand.current"},
					"totalDemandTax":         bson.M{"$sum": "$demand.totalTax"},
					"totalCollectionArrear":  bson.M{"$sum": "$collection.arrear"},
					"totalCollectionCurrent": bson.M{"$sum": "$collection.current"},
					"totalCollectionTax":     bson.M{"$sum": "$collection.totalTax"},
				}}}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"properties": bson.M{"$arrayElemAt": []interface{}{"$properties", 0}}}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
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

	// Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var wardwiseDemandandCollection []models.WardwiseDemandandCollection
	if err := cursor.All(ctx.CTX, &wardwiseDemandandCollection); err != nil {
		return nil, err
	}

	return wardwiseDemandandCollection, nil

}

//EnableProperty :""
func (d *Daos) EnableProperty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProperty :""
func (d *Daos) DisableProperty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProperty :""
func (d *Daos) DeleteProperty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//PropertyTimelineUpdate : ""
func (d *Daos) PropertyTimelineUpdate(ctx *models.Context, ID string, data interface{}, timeline models.PropertyTimeline) error {
	selector := bson.M{"uniqueId": ID}
	updateData := bson.M{"$set": data, "$push": bson.M{"log": timeline}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, updateData)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//GetPropertyDemandCalc : ""
func (d *Daos) GetPropertyDemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter, collectionName string) (*models.PropertyDemand, error) {
	var rootCollectionName string
	if collectionName == constants.COLLECTIONESTIMATEDPROPERTYDEMAND {
		rootCollectionName = constants.COLLECTIONESTIMATEDPROPERTYDEMAND
	} else {
		rootCollectionName = constants.COLLECTIONPROPERTY
	}

	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"uniqueId": filter.PropertyID})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	additionalQuery, err := d.GetPropertyDemandCalcQueryV2(ctx, filter, collectionName)
	if err != nil {
		return nil, errors.New("Additional query failed - " + err.Error())
	}
	mainPipeline = append(mainPipeline, additionalQuery...)
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROPERTYWALLET, "uniqueId", "propertyId", "ref.wallet", "ref.wallet")...)
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	var cursor *mongo.Cursor

	cursor, err = ctx.DB.Collection(rootCollectionName).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var propertyDemands []models.PropertyDemand
	var propertyDemand *models.PropertyDemand
	if err = cursor.All(ctx.CTX, &propertyDemands); err != nil {
		return nil, err
	}
	if len(propertyDemands) > 0 {
		propertyDemand = &propertyDemands[0]
	}

	if propertyDemand == nil {
		fmt.Println("Aggregation Failed")
	}
	return propertyDemand, nil
}

//GetMultiplePropertyDemandCalc : ""
func (d *Daos) GetMultiplePropertyDemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter, pagination *models.Pagination) ([]models.PropertyDemand, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.PropertyIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.PropertyIDs}})
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
	}

	// query = append(query, bson.M{"uniqueId": filter.PropertyID})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	additionalQuery, err := d.GetPropertyDemandCalcQuery(ctx, filter)
	if err != nil {
		return nil, errors.New("Additional query failed - " + err.Error())
	}
	mainPipeline = append(mainPipeline, additionalQuery...)
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"totalTax": -1}})
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyDemands []models.PropertyDemand
	if err = cursor.All(ctx.CTX, &propertyDemands); err != nil {
		return nil, err
	}
	return propertyDemands, nil
}

//GetPropertyDemandCalcQuery : ""
func (d *Daos) GetPropertyDemandCalcQuery(ctx *models.Context, filter *models.PropertyDemandFilter) ([]bson.M, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "propertyconfiguration",
			"as":   "propertyConfig",
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"uniqueId": "1"}},
			},
		},
	})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"propertyConfig": bson.M{"$arrayElemAt": []interface{}{"$propertyConfig", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONOTHERCHARGES,
			"as":   "otherCharges",
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"uniqueId": "GEN00001"}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"otherCharges": bson.M{"$arrayElemAt": []interface{}{"$otherCharges", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"percentAreaBuildup": bson.M{"$multiply": []interface{}{bson.M{"$divide": []string{"$builtUpArea", "$areaOfPlot"}}, 100}},
		//"taxableVacantLand":  bson.M{"$multiply": []interface{}{"$propertyConfig.taxableVacantLandConfig", bson.M{"$subtract": []string{"$areaOfPlot", "$builtUpArea"}}}}

	}})

	floorPipeline := []bson.M{}
	floorPipeline = append(floorPipeline, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
		bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
		bson.M{"$eq": []string{"$status", constants.PROPERTYFLOORSTATUSACTIVE}},
		bson.M{"$gte": []string{"$$fydateTo", "$dateFrom"}},
		// bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$dateTo", nil}}, "then": true, "else": bson.M{"$cond": bson.M{"if": bson.M{"$gte": []string{"$dateTo", "$$fyfromDate"}}, "then": true, "else": false}}}}}},
		bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$not": []interface{}{"$dateTo"}}, "then": true, "else": bson.M{"$cond": bson.M{"if": bson.M{"$gte": []string{"$dateTo", "$$fyfromDate"}}, "then": true, "else": false}}}}}},
	}}}})
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "usageType", "uniqueId", "ref.usageType", "ref.usageType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "constructionType", "uniqueId", "ref.constructionType", "ref.constructionType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONOCCUMANCYTYPE, "occupancyType", "uniqueId", "ref.occupancyType", "ref.occupancyType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR, "nonResUsageType", "uniqueId", "ref.nonResUsageType", "ref.nonResUsageType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONFLOORRATABLEAREA, "ratableAreaType", "uniqueId", "ref.floorRatableArea", "ref.floorRatableArea")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONFLOORTYPE, "no", "uniqueId", "ref.floorNo", "ref.floorNo")...)
	floorPipeline = append(floorPipeline,
		bson.M{
			"$lookup": bson.M{
				"from": constants.COLLECTIONAVR,
				"as":   "ref.avr",
				"let":  bson.M{"fyfromDate": "$$fyfromDate", "fydateTo": "$$fydateTo", "roadType": "$$roadType", "propertyId": "$$propertyId", "municipalityType": "$$municipalityType", "constructionType": "$constructionType", "usageType": "$usageType"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$municipalityTypeId", "$$municipalityType"}},
						bson.M{"$eq": []string{"$constructionTypeId", "$$constructionType"}},
						bson.M{"$eq": []string{"$roadTypeId", "$$roadType"}},
						bson.M{"$eq": []string{"$usageTypeId", "$$usageType"}},
						//bson.M{"$eq": []string{"$status", "Active"}},
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				},
			},
		})

	floorPipeline = append(floorPipeline, bson.M{
		"$addFields": bson.M{
			"ref.avr": bson.M{"$arrayElemAt": []interface{}{"$ref.avr", 0}},
		},
	})
	/*floorPipeline = append(floorPipeline,
		bson.M{
			"$lookup": bson.M{
				"from":     constants.COLLECTIONPROPERTYTAX,
				"as":       "ref.propertyTax",
				"let":      bson.M{"fyfromDate": "$$fyfromDate", "fydateTo": "$$fydateTo", "roadType": "$$roadType", "propertyId": "$$propertyId", "municipalityType": "$$municipalityType", "constructionType": "$constructionType", "usageType": "$usageType"},
				"pipeline": []bson.M{},
			},
		})
	floorPipeline = append(floorPipeline, bson.M{
		"$addFields": bson.M{
			"ref.propertyTax": bson.M{"$arrayElemAt": []interface{}{"$ref.propertyTax", 0}},
		},
	})*/

	/*
		floorPipeline = append(floorPipeline, bson.M{
			"$addFields": bson.M{
				"arv": bson.M{"$multiply": []interface{}{"$ref.floorRatableArea.rate", "$ref.avr.rate", "$ref.occupancyType.factor"}},
			},
		})

		floorPipeline = append(floorPipeline, bson.M{
			"$addFields": bson.M{
				"aptr": bson.M{"$multiply": []interface{}{"$arv", "$$propertyTax.rate"}},
			},
		})
	*/
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "completedFys",
		"let":  bson.M{"propertyId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
				bson.M{"$eq": []string{"$status", "Completed"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "ids": bson.M{"$push": "$fy.uniqueId"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"completedFys": bson.M{"$arrayElemAt": []interface{}{"$completedFys", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"completedFys.ids": bson.M{"$cond": bson.M{"if": bson.M{"$not": []interface{}{"$completedFys.ids"}}, "then": []interface{}{}, "else": "$completedFys.ids"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "financialyears",
		"as":   "fys",
		"let":  bson.M{"completedFys": "$completedFys.ids", "tempDOA": "$doa", "propertyId": "$uniqueId", "tempMT": "$municipalityId", "tempRT": "$roadTypeId", "taxablevland": "$taxableVacantLand", "percentAreaBuiltUp": "$percentAreaBuildup", "propertyConfig": "$propertyConfig"},
		"pipeline": []bson.M{
			bson.M{"$sort": bson.M{"order": 1}},
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": func() []bson.M {
				fySelectionAndQuery := make([]bson.M, 0)
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$lt": []string{"$$tempDOA", "$to"}})
				if filter != nil {
					if filter.IsOmitPaidFys {
						return fySelectionAndQuery
					}
				}
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$cond": bson.M{"if": bson.M{"$in": []interface{}{"$uniqueId", "$$completedFys"}}, "then": false, "else": true}})
				return fySelectionAndQuery
				//return []bson.M{

				// bson.M{"$ne": []string{"$uniqueId", "$$completedFys.ids"}},

				// bson.M{"$lt": []interface{}{time.Now(), "$to"}},
				//}
			}(),
			}}},
			bson.M{"$match": func() bson.M {
				if len(filter.Fys) > 0 {
					return bson.M{"uniqueId": bson.M{"$in": filter.Fys}}
				}
				return bson.M{}
			}()},
			bson.M{"$lookup": bson.M{
				"from": "vacantlandrates",
				"as":   "vlr",
				"let":  bson.M{"tempFrom": "$from", "tempTo": "$to", "mt": "$$tempMT", "rt": "$$tempRT", "doa": "$$tempDOA"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$municipalityTypeId", "$$mt"}},
						bson.M{"$eq": []string{"$roadTypeId", "$$rt"}},
						//bson.M{"$lte": []string{"$doe", "$$tempTo"}},
						// bson.M{"$gte": []string{"$doe", "$$doa"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
					bson.M{"$limit": 1},
				},
			},
			},
			bson.M{"$addFields": bson.M{"vlr": bson.M{"$arrayElemAt": []interface{}{"$vlr", 0}}}},
			//bson.M{"$addFields": bson.M{"vacantLandTax": bson.M{"$cond": bson.M{"if": bson.M{"$lt": []string{"$$percentAreaBuiltUp", "$$propertyConfig.vacantLandRatePercentage"}}, "then": bson.M{"$multiply": []string{"$$taxablevland", "$vlr.rate"}}, "else": 0}}}},
			bson.M{"$lookup": bson.M{
				"as":   "ref.propertyTax",
				"from": constants.COLLECTIONPROPERTYTAX,
				"let":  bson.M{"fydateTo": "$to", "fyfromDate": "$from"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				}}},
			bson.M{"$addFields": bson.M{"ref.propertyTax": bson.M{"$arrayElemAt": []interface{}{"$ref.propertyTax", 0}}}},

			bson.M{"$lookup": bson.M{
				"as":   "ref.penalty",
				"from": constants.COLLECTIONPENALTY,
				"let":  bson.M{"fydateTo": "$to", "fyfromDate": "$from"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				}}},
			bson.M{"$addFields": bson.M{"ref.penalty": bson.M{"$arrayElemAt": []interface{}{"$ref.penalty", 0}}}},

			bson.M{
				"$lookup": bson.M{
					"from":     "floors",
					"as":       "floors",
					"let":      bson.M{"fyfromDate": "$from", "fydateTo": "$to", "roadType": "$$tempRT", "propertyId": "$$propertyId", "municipalityType": "$$tempMT", "propertyTax": "$ref.propertyTax"},
					"pipeline": floorPipeline,
				},
			},
			bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONLEGACYYEAR,
				"as":   "legacy",
				"let":  bson.M{"fyId": "$uniqueId", "propertyId": "$$propertyId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$status", constants.LEGACYPROPERTYFYSTATUSACTIVE}},
						bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
						bson.M{"$eq": []string{"$fyId", "$$fyId"}},
					}}}},
				},
			},
			},
			bson.M{"$addFields": bson.M{"legacy": bson.M{"$arrayElemAt": []interface{}{"$legacy", 0}}}},
		},
	}})
	return mainPipeline, nil
}

//DemandCalc : ""
func (d *Daos) DemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter) ([]models.PropertyDemand, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if filter.PropertyID != "" {
			query = append(query, bson.M{"uniqueId": filter.PropertyID})
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
	}

	mainPipeline = append(mainPipeline, bson.M{"$match": func() bson.M {
		if len(query) > 0 {
			return bson.M{"$and": query}
		}
		return bson.M{}
	}()})
	additionalQuery, err := d.GetPropertyDemandCalcQuery(ctx, filter)
	if err != nil {
		return nil, errors.New("Additional query failed - " + err.Error())
	}
	mainPipeline = append(mainPipeline, additionalQuery...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyDemands []models.PropertyDemand
	if err = cursor.All(ctx.CTX, &propertyDemands); err != nil {
		return nil, err
	}
	return propertyDemands, nil

}

// UpdatePropertyGISTagging : ""
func (d *Daos) UpdatePropertyGISTagging(ctx *models.Context, UniqueID string, gis *models.PropertyGISTagging) error {
	selector := bson.M{"uniqueId": UniqueID}
	data := bson.M{"$set": bson.M{"gis": gis}}

	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

// BasicUpdateProperty : ""
func (d *Daos) BasicUpdateProperty(ctx *models.Context, bpu *models.BasicPropertyUpdate) error {
	selector1 := bson.M{"uniqueId": bpu.PropertyID}
	data1 := bson.M{"$set": bson.M{"address": bpu.Address}}
	res1, err1 := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector1, data1)
	if err1 != nil {
		return err1
	}
	fmt.Println(res1)
	opts := options.Update().SetUpsert(true)
	selector2 := bson.M{"uniqueId": bpu.Owner.UniqueID, "propertyId": bpu.PropertyID}
	data2 := bson.M{"$set": bpu.Owner}
	res2, err2 := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, selector2, data2, opts)
	if err2 != nil {
		return err2
	}
	fmt.Println(res2)
	return nil
}

//DashboardPropertyStatus : ""
func (d *Daos) DashboardPropertyStatus(ctx *models.Context, filter *models.DashboardPropertyStatusFilter) (*models.DashboardPropertyStatus, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if filter.Address != nil {
			if len(filter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": filter.Address.StateCode})
			}
			if len(filter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": filter.Address.DistrictCode})
			}
			if len(filter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": filter.Address.VillageCode})
			}
			if len(filter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": filter.Address.ZoneCode})
			}
			if len(filter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": filter.Address.WardCode})
			}
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$status", "count": bson.M{"$sum": 1}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "data": bson.M{"$push": bson.M{"k": "$_id", "v": "$$ROOT"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data": bson.M{"$arrayToObject": "$data"}}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dashPropStatuss []models.DashboardPropertyStatus
	var dashPropStatus *models.DashboardPropertyStatus
	if err = cursor.All(ctx.CTX, &dashPropStatuss); err != nil {
		return nil, err
	}
	if len(dashPropStatuss) > 0 {
		dashPropStatus = &dashPropStatuss[0]
	}
	return dashPropStatus, nil
}

//UpdateBoringStatusToProperty : ""
func (d *Daos) UpdateBoringStatusToProperty(ctx *models.Context, propertyID string, boringstatus bool) error {
	selector := bson.M{"uniqueId": propertyID}
	data := bson.M{"$set": bson.M{"isBoringChargePayed": boringstatus}}

	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

//UpdateFormFeeStatusToproperty:""
func (d *Daos) UpdateFormFeeStatusToProperty(ctx *models.Context, propertyID string, formfeestatus bool) error {
	selector := bson.M{"uniqueId": propertyID}
	data := bson.M{"$set": bson.M{"isFormFeePayed": formfeestatus}}

	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

// UpdateDemandToProperty : ""
func (d *Daos) UpdateDemandToProperty(ctx *models.Context, propertyID string, UpdateDemand *models.UpdateDemand) error {
	selector := bson.M{"uniqueId": propertyID}
	data := bson.M{"$set": bson.M{
		"demand": UpdateDemand,
	}}
	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

// UpdateNewDemandToProperty : ""
func (d *Daos) UpdateNewDemandToProperty(ctx *models.Context, propertyID string, UpdateDemand *models.PropertyTaxTotalDemand) error {
	selector := bson.M{"uniqueId": propertyID}
	data := bson.M{"$set": bson.M{
		"ndemand": UpdateDemand,
	}}
	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (d *Daos) GetAllPropertyIds(ctx *models.Context) ([]string, error) {
	t := []struct{ UniqueID string }{}
	query := bson.M{"status": bson.M{"$in": []string{"Active", "Init"}}}
	projection := bson.M{"_id": 0, "uniqueId": 1}
	opts := options.Find()
	opts.SetProjection(projection)
	// opts.SetSkip(0)
	// opts.SetLimit(500)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Find(ctx.CTX, query, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx.CTX, &t)
	if err != nil {
		return nil, err
	}
	data := []string{}
	for _, v := range t {
		data = append(data, v.UniqueID)
	}
	return data, nil
}

func (d *Daos) UpdatePropertyPreviousYrCollection(ctx *models.Context, ppyc *models.PropertyPreviousYrCollection) error {

	query := bson.M{"uniqueId": ppyc.UniqueID}
	update := bson.M{"$set": bson.M{
		"previousCollection.amount":       ppyc.Amount,
		"previousCollection.isCalculated": false,
	}}

	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (d *Daos) UpdatePayedPropertyPreviousYrCollection(ctx *models.Context, uniqueId string) error {

	query := bson.M{"uniqueId": uniqueId}
	update := bson.M{"$set": bson.M{
		"previousCollection.isCalculated": true,
	}}

	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (d *Daos) PropertyOverallDemandReport(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) ([]models.RefProperty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)

	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, propertyfilter)
	if propertyfilter != nil {
		if propertyfilter.OmitZeroDemand {
			query = append(query, bson.M{"ref.demand.total.totalTax": bson.M{"$gt": 0}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// if propertyfilter.SortBy != "" {
	// 	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.SortBy: propertyfilter.SortOrder}})
	// } else {
	// 	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	// }
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("Property Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	//Lookups
	if propertyfilter != nil {
		if !propertyfilter.RemoveLookup.PropertyOwner {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
		}
		if !propertyfilter.RemoveLookup.State {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
		}
		if !propertyfilter.RemoveLookup.District {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
		}
		if !propertyfilter.RemoveLookup.Village {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
		}
		if !propertyfilter.RemoveLookup.Zone {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
		}
		if !propertyfilter.RemoveLookup.Ward {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
		}
		if !propertyfilter.RemoveLookup.PropertyFloor {
			mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
		}
		if !propertyfilter.RemoveLookup.PropertyOwner {
			mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
		}
		if !propertyfilter.RemoveLookup.PropertyType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
		}
		if !propertyfilter.RemoveLookup.RoadType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
		}
		if !propertyfilter.RemoveLookup.MunicipalType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
		}

	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}

//PropertyParkPenaltyEnable :""
func (d *Daos) PropertyParkPenaltyEnable(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"parkPenalty": constants.PropertyParkPenaltyEnable}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//PropertyParkPenaltyDisable :""
func (d *Daos) PropertyParkPenaltyDisable(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"parkPenalty": constants.PropertyParkPenaltyDisable}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// UpdatePropertyLocation : ""
func (d *Daos) UpdatePropertyLocation(ctx *models.Context, property *models.PropertyLocation) error {
	selector := bson.M{"uniqueId": property.PropertyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"address.location": property.Location, "address.isGeoTagged": "Yes"}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// UpdatePropertyPicture : ""
func (d *Daos) UpdatePropertyPicture(ctx *models.Context, property *models.PropertyPicture) error {
	selector := bson.M{"uniqueId": property.PropertyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"picture": property.Picture}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// UpdateStoredPropertyDemand : ""
func (d *Daos) UpdateStoredPropertyDemand(ctx *models.Context, UniqueID string, spd *models.StoredPropertyDemand) error {
	selector := bson.M{"uniqueId": UniqueID}
	data := bson.M{"$set": bson.M{"spd": spd}}

	res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

//GetDemandV3 : ""
func (d *Daos) GetDemandV3(ctx *models.Context, UniqueID string) (*models.StoredPropertyDemand, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"propertyId": UniqueID}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"from": constants.COLLECTIONPROPERTYPAYMENTFY, "as": "collections",
		"let": bson.M{"financialyearId": "$financialyearId", "propertyId": "$propertyId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
				bson.M{"$eq": []string{"$fy.uniqueId", "$$financialyearId"}},
				bson.M{"$eq": []string{"$status", "Completed"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil,
				"vlTax":    bson.M{"$sum": "$fy.vacantLandTax"},
				"flTax":    bson.M{"$sum": "$fy.sumFloorTax"},
				"rebate":   bson.M{"$sum": "$fy.rebate"},
				"penanty":  bson.M{"$sum": "$fy.penanty"},
				"tax":      bson.M{"$sum": "$fy.tax"},
				"totalTax": bson.M{"$sum": "$fy.totalTax"},
			}},
		},
	}},
		bson.M{"$addFields": bson.M{"collections": bson.M{"$arrayElemAt": []interface{}{"$collections", 0}}}},
		bson.M{"$lookup": bson.M{
			"from": "financialyears",
			"as":   "ref.fy",
			"let":  bson.M{"financialyearId": "$financialyearId"},

			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$uniqueId", "$$financialyearId"}},
				}}},
				}},
		}},
		bson.M{"$lookup": bson.M{
			"from": "rebate",
			"as":   "ref.rainWaterHarvestRebate",

			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$uniqueId", "RainWaterHarvesting"}}}}}},
			},
		}},

		bson.M{"$lookup": bson.M{
			"from": "rebate",
			"as":   "ref.earlyPaymentRebate",

			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$uniqueId", "EarlyPayment"}}}}}},
			},
		}},

		bson.M{"$addFields": bson.M{
			"ref.fy":                     bson.M{"$arrayElemAt": []interface{}{"$ref.fy", 0}},
			"ref.earlyPaymentRebate":     bson.M{"$arrayElemAt": []interface{}{"$ref.earlyPaymentRebate", 0}},
			"ref.rainWaterHarvestRebate": bson.M{"$arrayElemAt": []interface{}{"$ref.rainWaterHarvestRebate", 0}},
		}},
		bson.M{"$group": bson.M{"_id": nil, "fys": bson.M{"$push": "$$ROOT"}}},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetDemandV3 query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONOSTOREDPROPERTYDEMANDFYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var demands []models.StoredPropertyDemand
	var demand *models.StoredPropertyDemand
	if err = cursor.All(ctx.CTX, &demands); err != nil {
		return nil, err
	}
	fmt.Println("demanda leen", len(demands))

	if len(demands) > 0 {
		demand = &demands[0]
	}
	return demand, nil
}

func (d *Daos) EnableHoldingProperty(ctx *models.Context, property *models.Property) error {
	selector := bson.M{"uniqueId": property.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"reason": property.Reason, "holdingStatus": "Yes"}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) DisableHoldingProperty(ctx *models.Context, propertyID string) error {
	query := bson.M{"uniqueId": propertyID}
	update := bson.M{"$set": bson.M{"holdingStatus": "No"}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// // GetSingleFinancialYearUsingDate : ""
// func (d *Daos) CheckWhetherPenalChargeApplies(ctx *models.Context, Date *time.Time, status string) (*models.RefPenalChargeFYRange, error) {
// 	mainPipeline := []bson.M{}
// 	query := []bson.M{}
// 	if Date != nil {
// 		sd := time.Date(Date.Year(), Date.Month(), Date.Day(), 0, 0, 0, 0, Date.Location())
// 		query = append(query, bson.M{"from": bson.M{"$gte": sd}})
// 		query = append(query, bson.M{"to": bson.M{"$lte": sd}})

// 	}
// 	query = append(query, bson.M{"status": status})
// 	if len(query) > 0 {
// 		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
// 	}
// 	// res, err := d.GetSingleFinancialYearUsingDate(ctx, Date)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// FYID = res.UniqueID
// 	d.Shared.BsonToJSONPrintTag("penal charge using date query =>", mainPipeline)
// 	//Aggregation
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONPENALCHARGEFYRANGE).Aggregate(ctx.CTX, mainPipeline, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var financialYears []models.RefPenalChargeFYRange
// 	var financialYear *models.RefPenalChargeFYRange
// 	details := new(models.RefPenalChargeFYDetails)
// 	if err = cursor.All(ctx.CTX, &financialYears); err != nil {
// 		return nil, err
// 	}
// 	if len(financialYears) > 0 {

// 		financialYear.PenalCharge = "Yes"
// 	}
// 	return financialYear, nil
// }
func (d *Daos) GetAllPropertyDemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter, collectionName string, pagination *models.Pagination) ([]models.PropertyDemand, error) {
	var rootCollectionName string
	if collectionName == constants.COLLECTIONESTIMATEDPROPERTYDEMAND {
		rootCollectionName = constants.COLLECTIONESTIMATEDPROPERTYDEMAND
	} else {
		rootCollectionName = constants.COLLECTIONPROPERTY
	}

	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	additionalQuery, err := d.GetPropertyDemandCalcQueryV2(ctx, filter, collectionName)
	if err != nil {
		return nil, errors.New("Additional query failed - " + err.Error())
	}
	mainPipeline = append(mainPipeline, additionalQuery...)
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROPERTYWALLET, "uniqueId", "propertyId", "ref.wallet", "ref.wallet")...)
	mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	var cursor *mongo.Cursor

	cursor, err = ctx.DB.Collection(rootCollectionName).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var propertyDemands []models.PropertyDemand
	//	var propertyDemand *models.PropertyDemand
	if err = cursor.All(ctx.CTX, &propertyDemands); err != nil {
		return nil, err
	}

	if propertyDemands == nil {
		fmt.Println("Aggregation Failed")
	}
	return propertyDemands, nil
}

// CheckWardWisesOldHoldingNoOfProperty : ""
func (d *Daos) CheckWardWiseOldHoldingNoOfProperty(ctx *models.Context, ward string, oldHoldingNo string) (*models.RefProperty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"address.wardCode": ward, "oldHoldingNumber": oldHoldingNo}})

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	var property *models.RefProperty
	if err = cursor.All(ctx.CTX, &propertys); err != nil {
		return nil, err
	}
	if len(propertys) > 0 {
		property = &propertys[0]
	}
	return property, nil
}

//UpdateProperty : ""
func (d *Daos) UpdatePropertyTotalDemand(ctx *models.Context, property *models.UpdatePropertyTotalDemand) error {
	selector := bson.M{"uniqueId": property.PropertyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"demand.totalTax": property.TotalAmount}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// UpdatePropertyUniqueID :""
func (d *Daos) UpdatePropertyUniqueID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"uniqueId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldUniqueId": uniqueIds.OldUniqueID, "newUniqueId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSinglePropertyUsingOldUniqueID : ""
func (d *Daos) GetSinglePropertyUsingOldUniqueID(ctx *models.Context, UniqueID string) (*models.RefProperty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"oldUniqueId": UniqueID}})

	//Aggregation
	//d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	var property *models.RefProperty
	if err = cursor.All(ctx.CTX, &propertys); err != nil {
		return nil, err
	}
	if len(propertys) > 0 {
		property = &propertys[0]
	}
	return property, nil
}

func (d *Daos) CreateUserChargeForProperty(ctx *models.Context, business *models.PropertyUserCharge) error {
	selector := bson.M{"uniqueId": business.PropertyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{
		"userCharge": business,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) GetPropertyDAOUsingPropertyID(ctx *models.Context, UniqueID string) (*models.RefPropertyFloor, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"propertyId": UniqueID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"dateFrom": 1}})

	//Aggregation
	//d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefPropertyFloor
	var property *models.RefPropertyFloor
	if err = cursor.All(ctx.CTX, &propertys); err != nil {
		return nil, err
	}
	if len(propertys) > 0 {
		property = &propertys[0]
	}
	return property, nil
}

//GetSingleProperty : ""
func (d *Daos) GetSinglePropertyWithUserCharge(ctx *models.Context, UniqueID string) (*models.RefProperty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID, "userCharge.isUserCharge": "Yes"}})

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	var property *models.RefProperty
	if err = cursor.All(ctx.CTX, &propertys); err != nil {
		return nil, err
	}
	if len(propertys) > 0 {
		property = &propertys[0]
	}
	return property, nil
}
