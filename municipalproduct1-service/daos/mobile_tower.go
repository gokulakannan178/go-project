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

// SavePropertyMobileTower : ""
func (d *Daos) SavePropertyMobileTower(ctx *models.Context, mobileTower *models.PropertyMobileTower) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).InsertOne(ctx.CTX, mobileTower)
	if err != nil {
		return err
	}
	fmt.Println(res)
	// var mobileTowerPropertyFys []interface{}
	// for _, v := range mobileTower.MobileTowerDemandFY {
	// 	mobileTowerPropertyFys = append(mobileTowerPropertyFys, v)
	// }
	// fy, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERYEAR).InsertMany(ctx.CTX, mobileTowerPropertyFys)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(fy)
	return nil
}

// GetSinglePropertyMobileTower : ""
func (d *Daos) GetSinglePropertyMobileTower(ctx *models.Context, UniqueID string) (*models.RefPropertyMobileTower, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefPropertyMobileTower
	var tower *models.RefPropertyMobileTower
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdatePropertyMobileTower : ""
func (d *Daos) UpdatePropertyMobileTower(ctx *models.Context, mobile *models.PropertyMobileTower) error {
	selector := bson.M{"uniqueId": mobile.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": mobile}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyMobileTower : ""
func (d *Daos) EnablePropertyMobileTower(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMOBILETOWERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyMobileTower : ""
func (d *Daos) DisablePropertyMobileTower(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMOBILETOWERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyMobileTower : ""
func (d *Daos) DeletePropertyMobileTower(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMOBILETOWERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//RejectPropertyMobileTower :""
func (d *Daos) RejectPropertyMobileTower(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMOBILETOWERSTATUSREJECTED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyMobileTower : ""
func (d *Daos) FilterPropertyMobileTower(ctx *models.Context, filter *models.PropertyMobileTowerFilter, pagination *models.Pagination) ([]models.RefPropertyMobileTower, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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

		if filter.SearchText.PropertyNo != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.SearchText.PropertyNo, Options: "xi"}})
		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		}
		if filter.SearchText.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.SearchText.OwnerName, Options: "xi"}})
		}
		if filter.SearchText.Mobile != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.SearchText.Mobile, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefPropertyMobileTower
	if err = cursor.All(context.TODO(), &towers); err != nil {
		return nil, err
	}
	return towers, nil
}

// MobileTowerWithMobileNo : ""
func (d *Daos) MobileTowerWithMobileNo(ctx *models.Context, filter *models.MobileTowerWithMobileNoReq, pagination *models.Pagination) ([]models.MobileTowerWithMobileNoRes, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		query = append(query, bson.M{"mobileNo": filter.MobileNo})

		// query = append(query, bson.M{"status": bson.M{"$in": constants.PROPERTYOWNERSTATUSACTIVE}})
		// bson.M{"$expr": bson.M{"$and":
	}
	// Adding $match to filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$expr": bson.M{"$and": query}}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$propertyId"}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "propertyIds": bson.M{"$push": "$_id"}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONPROPERTY,
			"as":   "properties",
			"let":  bson.M{"varPropertyId": "$propertyIds"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", constants.PROPERTYSTATUSACTIVE}},
					bson.M{"$in": []string{"$uniqueId", "$$varPropertyId"}},
				}}}},

				bson.M{"$group": bson.M{"_id": nil, "uniqueId": bson.M{"$push": "$uniqueId"}}},
			},
		},
	})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"properties": bson.M{"$arrayElemAt": []interface{}{"$properties", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONMOBILETOWER,
			"as":   "mobiletowers",
			"let":  bson.M{"varUniqueId": "$properties.uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", constants.MOBILETOWERSTATUSACTIVE}},
					bson.M{"$in": []string{"$propertyId", "$$varUniqueId"}},
				}}}},
			},
		}})

	// Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var report []models.MobileTowerWithMobileNoRes
	if err := cursor.All(ctx.CTX, &report); err != nil {
		return nil, err
	}

	return report, nil

}

// MobileTowerPenaltyUpdate : ""
func (d *Daos) MobileTowerPenaltyUpdate(ctx *models.Context, mobile *models.MobileTowerPenaltyUpdate) error {
	selector := bson.M{"uniqueId": mobile.UniqueID}
	data := bson.M{"$set": bson.M{"mobileTowerPenalty": mobile.MobileTowerPenalty}}

	res, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

// UpdateMobileTowerPropertyID :""
func (d *Daos) UpdateMobileTowerPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
