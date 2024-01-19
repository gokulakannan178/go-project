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

//SaveCircle :""
func (d *Daos) SaveCircle(ctx *models.Context, circle *models.Circle) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).InsertOne(ctx.CTX, circle)
	return err
}

//GetSingleCircle : ""
func (d *Daos) GetSingleCircle(ctx *models.Context, code string) (*models.RefCircle, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"code": code}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "zoneCode", "code", "ref.zone", "ref.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.zone.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.village.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.stateCode", "code", "ref.state", "ref.state")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var circles []models.RefCircle
	var circle *models.RefCircle
	if err = cursor.All(ctx.CTX, &circles); err != nil {
		return nil, err
	}
	if len(circles) > 0 {
		circle = &circles[0]
	}
	return circle, nil
}

//UpdateCircle : ""
func (d *Daos) UpdateCircle(ctx *models.Context, circle *models.Circle) error {
	selector := bson.M{"code": circle.Code}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": circle, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCircle : ""
func (d *Daos) FilterCircle(ctx *models.Context, filter *models.CircleFilter, pagination *models.Pagination) ([]models.RefCircle, error) {
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
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Circle query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var circles []models.RefCircle
	if err = cursor.All(context.TODO(), &circles); err != nil {
		return nil, err
	}
	return circles, nil
}

//EnableCircle :""
func (d *Daos) EnableCircle(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.CIRCLESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCircle :""
func (d *Daos) DisableCircle(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.CIRCLESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCircle :""
func (d *Daos) DeleteCircle(ctx *models.Context, code string) error {
	query := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"status": constants.CIRCLESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
