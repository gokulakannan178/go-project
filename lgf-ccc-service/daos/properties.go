package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PropertiesRequest : ""
func (d *Daos) SaveProperties(ctx *models.Context, properties *models.Properties) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).InsertOne(ctx.CTX, properties)
	if err != nil {
		return err
	}
	return nil
}

// GetSingleProperties : ""
func (d *Daos) GetSingleProperties(ctx *models.Context, uniqueID string) (*models.RefProperties, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.zone", "ref.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadtype", "uniqueId", "ref.roadtype", "ref.roadtype")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.state", "ref.state")...)

	d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.RefProperties
	var property *models.RefProperties
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return nil, err
	}
	if len(properties) > 0 {
		property = &properties[0]
	}
	return property, err
}
func (d *Daos) GetSinglePropertiesWithHoldingNumber(ctx *models.Context, holdingNumber string) (*models.RefProperties, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"holdingNumber": holdingNumber}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.RefProperties
	var property *models.RefProperties
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return nil, err
	}
	if len(properties) > 0 {
		property = &properties[0]
	}
	return property, err
}

//UpdateProperties : ""
func (d *Daos) UpdateProperties(ctx *models.Context, properties *models.Properties) error {
	selector := bson.M{"uniqueId": properties.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": properties}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableProperties : ""
func (d *Daos) EnableProperties(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PROPERTIESSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableProperties : ""
func (d *Daos) DisableProperties(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PROPERTIESSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteProperties :""
func (d *Daos) DeleteProperties(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTIESSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// // EnableProperties : ""
// func (d *Daos) InProgressProperties(ctx *models.Context, uniqueID string) error {
// 	selector := bson.M{"uniqueId": uniqueID}
// 	data := bson.M{"$set": bson.M{"status": constants.PropertiesSTATUSINPROGRESS}}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, selector, data)
// 	return err
// }

// // DisableProperties : ""
// func (d *Daos) PendingProperties(ctx *models.Context, uniqueID string) error {
// 	selector := bson.M{"uniqueId": uniqueID}
// 	data := bson.M{"$set": bson.M{"status": constants.PropertiesSTATUSPENDING}}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, selector, data)
// 	return err
// }

// //DeleteProperties :""
// func (d *Daos) InitProperties(ctx *models.Context, UniqueID string) error {
// 	query := bson.M{"uniqueId": UniqueID}
// 	update := bson.M{"$set": bson.M{"status": constants.PropertiesSTATUSINIT}}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, query, update)
// 	if err != nil {
// 		return errors.New("Not Changed" + err.Error())
// 	}
// 	return err
// }

// //DeleteProperties :""
// func (d *Daos) CompletedProperties(ctx *models.Context, Properties *models.Properties) error {
// 	selector := bson.M{"uniqueId": Properties.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": Properties}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err

// }

// FilterProperties : ""
func (d *Daos) FilterProperties(ctx *models.Context, property *models.FilterProperties, pagination *models.Pagination) ([]models.RefProperties, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if property != nil {
		if len(property.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": property.Status}})
		}
		if len(property.GCID) > 0 {
			query = append(query, bson.M{"gcUser.id": bson.M{"$in": property.GCID}})
		}
		if len(property.ManagerID) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": property.ManagerID}})
		}
		if len(property.CitizenID) > 0 {
			query = append(query, bson.M{"citizen.id": bson.M{"$in": property.CitizenID}})
		}
		if len(property.RegisterID) > 0 {
			query = append(query, bson.M{"registerBy.id": bson.M{"$in": property.RegisterID}})
		}
		if len(property.Mobile) > 0 {
			query = append(query, bson.M{"mobile": bson.M{"$in": property.Mobile}})
		}
		if len(property.Pincode) > 0 {
			query = append(query, bson.M{"pincode": bson.M{"$in": property.Pincode}})
		}
		if len(property.PropertyType) > 0 {
			query = append(query, bson.M{"propertyType": bson.M{"$in": property.PropertyType}})
		}
		if len(property.OwnerType) > 0 {
			query = append(query, bson.M{"ownerType": bson.M{"$in": property.OwnerType}})
		}
		if len(property.WardCode) > 0 {
			query = append(query, bson.M{"address.wardCode": bson.M{"$in": property.WardCode}})
		}
		if len(property.StateCode) > 0 {
			query = append(query, bson.M{"address.stateCode": bson.M{"$in": property.StateCode}})
		}
		if len(property.DistrictCode) > 0 {
			query = append(query, bson.M{"address.districtCode": bson.M{"$in": property.DistrictCode}})
		}
		if len(property.ZoneCode) > 0 {
			query = append(query, bson.M{"address.zoneCode": bson.M{"$in": property.ZoneCode}})
		}
		if len(property.VillageCode) > 0 {
			query = append(query, bson.M{"address.villageCode": bson.M{"$in": property.VillageCode}})
		}
		//Regex
		if property.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: property.Regex.Name, Options: "xi"}})
		}
		if property.Regex.Citizen.Name != "" {
			query = append(query, bson.M{"citizen.name": primitive.Regex{Pattern: property.Regex.Citizen.Name, Options: "xi"}})
		}
		if property.Regex.HoldingNumber != "" {
			query = append(query, bson.M{"holdingNumber": primitive.Regex{Pattern: property.Regex.HoldingNumber, Options: "xi"}})
		}
		if property.Regex.NfcID != "" {
			query = append(query, bson.M{"nfcId": primitive.Regex{Pattern: property.Regex.NfcID, Options: "xi"}})
		}
		if property.Regex.PropertyType != "" {
			query = append(query, bson.M{"propertyType": primitive.Regex{Pattern: property.Regex.PropertyType, Options: "xi"}})
		}
		if property.Regex.HouseUID != "" {
			query = append(query, bson.M{"houseUid": primitive.Regex{Pattern: property.Regex.HouseUID, Options: "xi"}})
		}
	}
	if property.DateRange.From != nil {
		t := *property.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if property.DateRange.To != nil {
			t2 := *property.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"registerDate": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if property != nil {
		if property.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{property.SortBy: property.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.zone", "ref.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadtype", "uniqueId", "ref.roadtype", "ref.roadtype")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.RefProperties
	if err = cursor.All(context.TODO(), &properties); err != nil {
		return nil, err
	}
	return properties, nil
}

// func (d *Daos) GetDetailProperties(ctx *models.Context, uniqueID string) (*models.RefProperties, error) {
// 	mainPipeline := []bson.M{}
// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": constants.PropertiesSTATUSCOMPLETED}})
// 	//LookUp
// 	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

// 	//d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)

// 	//Aggregation
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var Propertiess []models.RefProperties
// 	var Properties *models.RefProperties
// 	if err = cursor.All(ctx.CTX, &Propertiess); err != nil {
// 		return nil, err
// 	}
// 	if len(Propertiess) > 0 {
// 		Properties = &Propertiess[0]
// 	}
// 	return Properties, err
// }
// func (d *Daos) AssignProperties(ctx *models.Context, Properties *models.Properties) error {
// 	selector := bson.M{"uniqueId": Properties.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": bson.M{"gcUser": Properties.GCUser, "assignDate": &t, "status": constants.PropertiesSTATUSPENDING}}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

func (d *Daos) GetSinglePropertyWithCondition(ctx *models.Context, key, value string) (*models.RefProperties, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{key: value}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.RefProperties
	var property *models.RefProperties
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return nil, err
	}
	if len(properties) > 0 {
		property = &properties[0]
	}
	return property, nil
}

func (d *Daos) CheckResgisterMobileNumber(ctx *models.Context, mobile string, uniqueId string) (*models.RefProperties, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": mobile, "uniqueid": bson.M{"$nin": uniqueId}}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.RefProperties
	var property *models.RefProperties
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return nil, err
	}
	if len(properties) > 0 {
		property = &properties[0]
	}
	return property, err
}

func (d *Daos) GetSinglePropertiesWithHouseUID(ctx *models.Context, houseUid string) (*models.RefProperties, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"houseUid": houseUid, "status": constants.PROPERTIESSTATUSACTIVE}})
	// //LookUp
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.RefProperties
	var property *models.RefProperties
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return nil, err
	}
	if len(properties) > 0 {
		property = &properties[0]
	}
	return property, err
}

func (d *Daos) GetPropertiesCountWithWard(ctx *models.Context, WardCode string) (*models.WardWisePropertyCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
		bson.M{"$eq": []string{"$wardCode", WardCode}},
	}}}},
		bson.M{"$group": bson.M{"_id": nil,
			"quantity": bson.M{"$sum": 1}}},
	)
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"quantity": "$_id.quantity"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.WardWisePropertyCount
	var property *models.WardWisePropertyCount
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return nil, err
	}
	if len(properties) > 0 {
		property = &properties[0]
	}
	return property, err
}

func (d *Daos) GetPropertiesCountWithCircle(ctx *models.Context, CircleCode string) (*models.WardWisePropertyCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
		bson.M{"$eq": []string{"$circleCode", CircleCode}},
	}}}},
		bson.M{"$group": bson.M{"_id": nil,
			"quantity": bson.M{"$sum": 1}}},
	)
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"quantity": "$_id.quantity"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var properties []models.WardWisePropertyCount
	var property *models.WardWisePropertyCount
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return nil, err
	}
	if len(properties) > 0 {
		property = &properties[0]
	}
	return property, err
}
func (d *Daos) GetPropertiesCountWithRoadtype(ctx *models.Context, Roadtype string) (int64, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{"roadtype": Roadtype},
	})
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"quantity": "$_id.quantity"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return 0, err
	}
	var properties []models.Properties
	if err = cursor.All(ctx.CTX, &properties); err != nil {
		return 0, err
	}
	var Properties int64
	if len(properties) > 0 {
		Properties = int64(len(properties))
	}
	return Properties, err
}
