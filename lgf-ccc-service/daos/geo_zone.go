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
)

//SaveZone :""
func (d *Daos) SaveZone(ctx *models.Context, zone *models.Zone) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONZONE).InsertOne(ctx.CTX, zone)
	return err
}

//GetSingleZone : ""
func (d *Daos) GetSingleZone(ctx *models.Context, code string) (*models.RefZone, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"code": code}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.ward.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.village.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONZONE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var zones []models.RefZone
	var zone *models.RefZone
	if err = cursor.All(ctx.CTX, &zones); err != nil {
		return nil, err
	}
	if len(zones) > 0 {
		zone = &zones[0]
	}
	return zone, nil
}

//UpdateZone : ""
func (d *Daos) UpdateZone(ctx *models.Context, zone *models.Zone) error {
	selector := bson.M{"code": zone.Code}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": zone, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONZONE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterZone : ""
func (d *Daos) FilterZone(ctx *models.Context, zonefilter *models.ZoneFilter, pagination *models.Pagination) ([]models.RefZone, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if zonefilter != nil {
		if len(zonefilter.Codes) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": zonefilter.Codes}})
		}
		if len(zonefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": zonefilter.Status}})
		}
		if len(zonefilter.VillageCodes) > 0 {
			query = append(query, bson.M{"villageCode": bson.M{"$in": zonefilter.VillageCodes}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONZONE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.ward.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.village.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.stateCode", "code", "ref.state", "ref.state")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "wardCode", "code", "ref.ward", "ref.ward")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("zone query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONZONE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var zones []models.RefZone
	if err = cursor.All(context.TODO(), &zones); err != nil {
		return nil, err
	}
	return zones, nil
}

//EnableZone :""
func (d *Daos) EnableZone(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.ZONESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONZONE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableZone :""
func (d *Daos) DisableZone(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.ZONESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONZONE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteZone :""
func (d *Daos) DeleteZone(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.ZONESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONZONE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
