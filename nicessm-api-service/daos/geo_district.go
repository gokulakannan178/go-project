package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveDistrict :""
func (d *Daos) SaveDistrict(ctx *models.Context, district *models.District) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).InsertOne(ctx.CTX, district)
	if err != nil {
		return err
	}
	district.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDistrict : ""
func (d *Daos) GetSingleDistrict(ctx *models.Context, code string) (*models.RefDistrict, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSOILTYPE, "soilTypes", "_id", "ref.soilTypes", "ref.soilTypes")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAGROECOLOGICALZONE, "agroEcologicalZones", "_id", "ref.agroEcologicalZones", "ref.agroEcologicalZones")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.RefDistrict
	var district *models.RefDistrict
	if err = cursor.All(ctx.CTX, &districts); err != nil {
		return nil, err
	}
	if len(districts) > 0 {
		district = &districts[0]
	}
	return district, nil
}

//UpdateDistrict : ""
func (d *Daos) UpdateDistrict(ctx *models.Context, district *models.District) error {

	selector := bson.M{"_id": district.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": district, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDistrict : ""
func (d *Daos) FilterDistrict(ctx *models.Context, districtfilter *models.DistrictFilter, pagination *models.Pagination) ([]models.RefDistrict, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if districtfilter != nil {
		if len(districtfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": districtfilter.ActiveStatus}})
		}
		if len(districtfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": districtfilter.State}})
		}
		if len(districtfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": districtfilter.Status}})
		}
		//Regex
		if districtfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: districtfilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSOILTYPE, "soilTypes", "_id", "ref.soilTypes", "ref.soilTypes")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAGROECOLOGICALZONE, "agroEcologicalZones", "_id", "ref.agroEcologicalZones", "ref.agroEcologicalZones")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("district query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.RefDistrict
	if err = cursor.All(context.TODO(), &districts); err != nil {
		return nil, err
	}
	return districts, nil
}

//EnableDistrict :""
func (d *Daos) EnableDistrict(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDistrict :""
func (d *Daos) DisableDistrict(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrict :""
func (d *Daos) DeleteDistrict(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleDistrictWithName : ""
func (d *Daos) GetSingleDistrictWithName(ctx *models.Context, Name string, StateID primitive.ObjectID) ([]models.RefDistrict, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
	}
	query = append(query, bson.M{"state": StateID})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("district query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.RefDistrict
	if err = cursor.All(context.TODO(), &districts); err != nil {
		return nil, err
	}
	return districts, nil
}
func (d *Daos) GetSingleDistrictWithUniqueId(ctx *models.Context, UniqueID string) (*models.RefDistrict, error) {
	mainPipeline := []bson.M{}

	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("district query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.RefDistrict
	if err = cursor.All(context.TODO(), &districts); err != nil {
		return nil, err
	}
	if len(districts) > 0 {
		return &districts[0], nil
	}

	return nil, errors.New("distric not found")
}

//GetSingleStateWithName : ""
func (d *Daos) GetSingleDistrictWithNameV2(ctx *models.Context, Name string, StateID primitive.ObjectID, isRegex bool) (*models.RefDistrict, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Name != "" {
		if isRegex {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
		} else {
			query = append(query, bson.M{"name": Name})

		}
	}
	query = append(query, bson.M{"state": StateID})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("district query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.RefDistrict
	if err = cursor.All(context.TODO(), &districts); err != nil {
		return nil, err
	}
	if len(districts) > 0 {
		return &districts[0], nil
	}
	return nil, errors.New("district not available")
}
func (d *Daos) GetActiveDistrict(ctx *models.Context) ([]models.District, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": constants.DISTRICTSTATUSACTIVE})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("activeDistrict query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Districts []models.District
	if err = cursor.All(context.TODO(), &Districts); err != nil {
		return nil, err
	}
	return Districts, nil
}
func (d *Daos) GetDistrictWeatherDataWithSeverityType(ctx *models.Context, filter *models.StateWeatherAlertMasterFilterv2, pagination *models.Pagination) ([]models.GetDistrictLeveWeatherDataAlert, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active"}})
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"state": 1}})

	query := []bson.M{}
	query = append(query, bson.M{"$eq": []string{"$severityType", "$$weathercode"}})
	query = append(query, bson.M{"$eq": []string{"$district", "$$districtcode"}})
	if !filter.ParameterId.IsZero() {
		query = append(query, bson.M{"$eq": []interface{}{"$parameterid", filter.ParameterId}})

	}
	if !filter.Month.IsZero() {
		query = append(query, bson.M{"$eq": []interface{}{"$month", filter.Month}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONWEATHERALERTTYPE,
		"as":   "severitytype",
		"let":  bson.M{"districtcode": "$_id"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}}}}}},
			//{"$project": bson.M{"name": 1}},
			{
				"$lookup": bson.M{
					"from": constants.COLLECTIONDISTRICTWEATHERALERTMASTER,
					"as":   "weatherdata",
					"let":  bson.M{"weathercode": "$_id"},
					"pipeline": []bson.M{
						{"$match": bson.M{"$expr": bson.M{"$and": query}}},
						//	{"$project": bson.M{"min": 1, "max": 1}},
					}}},
			bson.M{"$addFields": bson.M{"weatherdata": bson.M{"$arrayElemAt": []interface{}{"$weatherdata", 0}}}},
			bson.M{"$project": bson.M{
				"k":   "$name",
				"v":   "$$ROOT",
				"_id": 0,
			}},
		}},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"severitytype": bson.M{"$arrayToObject": "$severitytype"}}})
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		d.Shared.BsonToJSONPrintTag("state pagenation query =>", paginationPipeline)
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("DistrictsWeatherdataAlert query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.GetDistrictLeveWeatherDataAlert
	if err = cursor.All(context.TODO(), &districts); err != nil {
		return nil, err
	}
	return districts, nil
}
