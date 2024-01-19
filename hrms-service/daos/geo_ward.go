package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveWard :""
func (d *Daos) SaveWard(ctx *models.Context, ward *models.Ward) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONWARD).InsertOne(ctx.CTX, ward)
	return err
}

//GetSingleWard : ""
func (d *Daos) GetSingleWard(ctx *models.Context, code string) (*models.RefWard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"code": code}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "zoneCode", "code", "ref.zone", "ref.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.zone.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.village.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var wards []models.RefWard
	var ward *models.RefWard
	if err = cursor.All(ctx.CTX, &wards); err != nil {
		return nil, err
	}
	if len(wards) > 0 {
		ward = &wards[0]
	}
	return ward, nil
}

//UpdateWard : ""
func (d *Daos) UpdateWard(ctx *models.Context, ward *models.Ward) error {
	selector := bson.M{"code": ward.Code}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ward, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARD).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterWard : ""
func (d *Daos) FilterWard(ctx *models.Context, wardfilter *models.WardFilter, pagination *models.Pagination) ([]models.RefWard, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if wardfilter != nil {
		if len(wardfilter.Codes) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": wardfilter.Codes}})
		}
		if len(wardfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": wardfilter.Status}})
		}
		if len(wardfilter.ZoneCodes) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": wardfilter.ZoneCodes}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "zoneCode", "code", "ref.zone", "ref.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.zone.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.village.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("ward query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var wards []models.RefWard
	if err = cursor.All(context.TODO(), &wards); err != nil {
		return nil, err
	}
	return wards, nil
}

//EnableWard :""
func (d *Daos) EnableWard(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.WARDSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARD).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableWard :""
func (d *Daos) DisableWard(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.WARDSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARD).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWard :""
func (d *Daos) DeleteWard(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.WARDSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARD).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
