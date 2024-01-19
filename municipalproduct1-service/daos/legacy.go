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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveLEGACYS :""
func (d *Daos) SaveLegacy(ctx *models.Context, legacy *models.RegLegacyProperty) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONLEGACY).InsertOne(ctx.CTX, legacy.LegacyProperty)
	if err != nil {
		return err
	}
	fmt.Println(res)
	var tempLegacyPropertyFys []interface{}
	for _, v := range legacy.LegacyPropertyFy {
		tempLegacyPropertyFys = append(tempLegacyPropertyFys, v)
	}
	fy, err := ctx.DB.Collection(constants.COLLECTIONLEGACYYEAR).InsertMany(ctx.CTX, tempLegacyPropertyFys)
	if err != nil {
		return err
	}
	fmt.Println(fy)
	return nil
}

// GetLegacyForAProperty : ""
func (d *Daos) GetLegacyForAProperty(ctx *models.Context, propertyID string) (*models.RefLegacyPropertyPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"propertyId": propertyID})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Lookup to legacy financial year
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONLEGACYYEAR,
			"as":   "legacyPropertyFy",
			"let":  bson.M{"propertyId": propertyID},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []interface{}{"$propertyId", "$$propertyId"}},
					bson.M{"$eq": []interface{}{"$status", "Active"}},
				}}}},
				bson.M{
					"$lookup": bson.M{
						"from": constants.COLLECTIONFINANCIALYEAR,
						"as":   "ref.fy",
						"let":  bson.M{"fyId": "$fyId"},
						"pipeline": []bson.M{
							bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
								bson.M{"$eq": []interface{}{"$uniqueId", "$$fyId"}},
							}}}}}},
				},
				bson.M{"$addFields": bson.M{"ref.fy": bson.M{"$arrayElemAt": []interface{}{"$ref.fy", 0}}}},
			},
		},
	})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("sinale legacy property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEGACY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rlpps []models.RefLegacyPropertyPayment
	if err = cursor.All(context.TODO(), &rlpps); err != nil {
		return nil, err
	}
	var rlpp models.RefLegacyPropertyPayment
	if len(rlpps) > 0 {
		rlpp = rlpps[0]
		rlpp.IsAvailable = true
	}
	return &rlpp, nil
}

func (d *Daos) UpdateLegacyForAProperty(ctx *models.Context, legacy *models.RegLegacyProperty) error {
	t := time.Now()
	updated := models.Updated{
		On:     &t,
		By:     legacy.Updated,
		ByType: legacy.UpdatedType,
	}
	legacyPropertyUpdateQuery := bson.M{"propertyId": legacy.LegacyProperty.PropertyID}
	legacyPropertyUpdateData := bson.M{"$set": legacy.LegacyProperty, "$push": bson.M{"updated": updated}}
	res, err := ctx.DB.Collection(constants.COLLECTIONLEGACY).UpdateOne(ctx.CTX, legacyPropertyUpdateQuery, legacyPropertyUpdateData)
	if err != nil {
		return err
	}
	fmt.Println(res)
	fyUniqueIDs := []string{}
	for _, v := range legacy.LegacyPropertyFy {
		if v.UniqueID != "" {
			fyUniqueIDs = append(fyUniqueIDs, v.UniqueID)
		}
	}
	// if len(fyUniqueIDs) > 0 {
	legacyFyPropertyDeleteIds := func() bson.M {
		if len(fyUniqueIDs) > 0 {
			return bson.M{"uniqueId": bson.M{"$nin": fyUniqueIDs}}
		}
		return bson.M{}
	}()
	legacyFyPropertyDeleteData := bson.M{"$set": bson.M{"status": constants.LEGACYPROPERTYFYSTATUSDELETED}}
	res, err = ctx.DB.Collection(constants.COLLECTIONLEGACYYEAR).UpdateMany(ctx.CTX, legacyFyPropertyDeleteIds, legacyFyPropertyDeleteData)
	if err != nil {
		return err
	}
	fmt.Println("deleted count ==>", len(fyUniqueIDs))
	d.Shared.BsonToJSONPrintTag("deleted entries result", res)
	// }
	for _, v := range legacy.LegacyPropertyFy {
		opts := options.Update().SetUpsert(true)
		legacyFyPropertyUpdateId := bson.M{"uniqueId": v.UniqueID}
		legacyFyPropertyUpdateData := bson.M{"$set": v}
		res, err := ctx.DB.Collection(constants.COLLECTIONLEGACYYEAR).UpdateOne(ctx.CTX, legacyFyPropertyUpdateId, legacyFyPropertyUpdateData, opts)
		if err != nil {
			return err
		}
		d.Shared.BsonToJSONPrintTag("updated entries result", res)
	}
	return nil
}

func (d *Daos) GetFinancialYearsForLegacyPayments(ctx *models.Context, propertyId string) ([]models.RefV2LegacyPropertyFy, error) {
	property, err := d.GetSingleProperty(ctx, propertyId)
	if err != nil {
		return nil, errors.New("Error in geting Property - " + err.Error())
	}
	if property == nil {
		return nil, errors.New("Property is nil")
	}
	if property.DOA == nil {
		return nil, errors.New("Property is DOA is not valid")
	}

	mainPipeline := []bson.M{}
	// query := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"order": -1}})
	// mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active", "from": bson.M{"$gte": property.DOA}}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active", "to": bson.M{"$gte": property.DOA}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"from": "legacyyears", "as": "legacyyear", "let": bson.M{"fyId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$fyId", "$$fyId"}},
				bson.M{"$eq": []string{"$status", "Active"}},
				bson.M{"$eq": []string{"$propertyId", propertyId}},
			}}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"legacyyear": bson.M{"$arrayElemAt": []interface{}{"$legacyyear", 0}}}})
	d.Shared.BsonToJSONPrintTag("get legacy fy of property =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rv2lpf []models.RefV2LegacyPropertyFy
	if err = cursor.All(context.TODO(), &rv2lpf); err != nil {
		return nil, err
	}
	return rv2lpf, nil
}

func (d *Daos) GetReqFinancialYearForLegacy(ctx *models.Context, doa *time.Time) ([]models.RefFinancialYear, error) {

	mainPipeline := []bson.M{}
	// query := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"order": -1}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active", "from": bson.M{"$gte": doa}}})

	d.Shared.BsonToJSONPrintTag("get legacy fy of property =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rv2lpf []models.RefFinancialYear
	if err = cursor.All(context.TODO(), &rv2lpf); err != nil {
		return nil, err
	}
	return rv2lpf, nil
}

// FilterLegacy : ""
func (d *Daos) FilterLegacy(ctx *models.Context, filter *models.LegacyPropertyFilter, pagination *models.Pagination) ([]models.RefLegacyPropertyPayment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.PropertyIDs) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyIDs}})
		}
		if filter.DateRange != nil {
			if filter.DateRange.From != nil {
				sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				var ed time.Time
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				} else {
					ed = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
				}
				query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLEGACY).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONLEGACYYEAR,
			"as":   "legacyPropertyFy",
			"let":  bson.M{"propertyId": "$propertyId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []interface{}{"$propertyId", "$$propertyId"}},
					bson.M{"$eq": []interface{}{"$status", "Active"}},
				}}}},
				bson.M{
					"$lookup": bson.M{
						"from": constants.COLLECTIONFINANCIALYEAR,
						"as":   "ref.fy",
						"let":  bson.M{"fyId": "$fyId"},
						"pipeline": []bson.M{
							bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
								bson.M{"$eq": []interface{}{"$uniqueId", "$$fyId"}},
							}}}}}},
				},
				bson.M{"$addFields": bson.M{"ref.fy": bson.M{"$arrayElemAt": []interface{}{"$ref.fy", 0}}}},
			},
		},
	})
	// lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.user", "ref.user")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("legacy filter query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEGACY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.RefLegacyPropertyPayment
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}
