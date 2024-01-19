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

//SaveSector :""
func (d *Daos) SaveSector(ctx *models.Context, sector *models.Sector) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).InsertOne(ctx.CTX, sector)
	return err
}

//GetSingleSector : ""
func (d *Daos) GetSingleSector(ctx *models.Context, code string) (*models.RefSector, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"code": code}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "zoneCode", "code", "ref.zone", "ref.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.zone.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.village.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var sectors []models.RefSector
	var sector *models.RefSector
	if err = cursor.All(ctx.CTX, &sectors); err != nil {
		return nil, err
	}
	if len(sectors) > 0 {
		sector = &sectors[0]
	}
	return sector, nil
}

//UpdateSector : ""
func (d *Daos) UpdateSector(ctx *models.Context, sector *models.Sector) error {
	selector := bson.M{"code": sector.Code}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": sector, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterSector : ""
func (d *Daos) FilterSector(ctx *models.Context, filter *models.SectorFilter, pagination *models.Pagination) ([]models.RefSector, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Codes) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Codes}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.ZoneCodes) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.ZoneCodes}})
		}
		if len(filter.WardCode) > 0 {
			query = append(query, bson.M{"wardCode": bson.M{"$in": filter.WardCode}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Sector query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var sectors []models.RefSector
	if err = cursor.All(context.TODO(), &sectors); err != nil {
		return nil, err
	}
	return sectors, nil
}

//EnableSector :""
func (d *Daos) EnableSector(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.SECTORSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableSector :""
func (d *Daos) DisableSector(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.SECTORSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteSector :""
func (d *Daos) DeleteSector(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.SECTORSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSECTOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
