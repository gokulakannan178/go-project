package daos

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveDistrict :""
func (d *Daos) SaveDistrict(ctx *models.Context, district *models.District) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).InsertOne(ctx.CTX, district)
	return err
}

//GetSingleDistrict : ""
func (d *Daos) GetSingleDistrict(ctx *models.Context, code string) (*models.RefDistrict, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"code": code}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDIVISION, "divisionCode", "code", "ref.division", "ref.division")...)
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
func (d *Daos) GetSingleDistrictWithName(ctx *models.Context, code string) (*models.RefDistrict, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": code}})
	//Lookups
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDIVISION, "divisionCode", "code", "ref.division", "ref.division")...)
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
	selector := bson.M{"code": district.Code}
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
		if len(districtfilter.Codes) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": districtfilter.Codes}})
		}
		if len(districtfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": districtfilter.Status}})
		}
		if len(districtfilter.DivisionCode) > 0 {
			query = append(query, bson.M{"divisionCode": bson.M{"$in": districtfilter.DivisionCode}})
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDIVISION, "divisionCode", "code", "ref.division", "ref.division")...)
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
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDistrict :""
func (d *Daos) DisableDistrict(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrict :""
func (d *Daos) DeleteDistrict(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
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
func (d *Daos) FilterDistrictByState(ctx *models.Context, filter *models.DistrictFilter, pagination *models.Pagination) ([]models.DistrictName, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.State) > 0 {
			query = append(query, bson.M{"imdStateName": bson.M{"$in": filter.State}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
		"_id":          nil,
		"stateName":    bson.M{"$push": "$imdStateName"},
		"districtName": bson.M{"$push": "$imdDistrictName"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"stateName": bson.M{"$arrayElemAt": []interface{}{"$stateName", 0}}}})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"districtName": bson.M{"$arrayElemAt": []interface{}{"$districtName", 0}}}})
	// mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"stateName": 1, "districtName": 1, "_id": 0}})

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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.DistrictName
	if err = cursor.All(context.TODO(), &districts); err != nil {
		return nil, err
	}
	return districts, nil
}
