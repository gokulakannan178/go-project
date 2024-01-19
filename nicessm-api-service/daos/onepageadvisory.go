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

//SaveDisease :""
func (d *Daos) SaveOnePageAdvisory(ctx *models.Context, OnePageAdvisory *models.OnePageAdvisory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).InsertOne(ctx.CTX, OnePageAdvisory)
	if err != nil {
		return err
	}
	OnePageAdvisory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleOnePageAdvisory(ctx *models.Context, UniqueID string) (*models.RefOnePageAdvisory, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OnePageAdvisorys []models.RefOnePageAdvisory
	var OnePageAdvisory *models.RefOnePageAdvisory
	if err = cursor.All(ctx.CTX, &OnePageAdvisorys); err != nil {
		return nil, err
	}
	if len(OnePageAdvisorys) > 0 {
		OnePageAdvisory = &OnePageAdvisorys[0]
	}
	return OnePageAdvisory, nil
}

//UpdateOnePageAdvisory : ""
func (d *Daos) UpdateOnePageAdvisory(ctx *models.Context, OnePageAdvisory *models.OnePageAdvisory) error {

	selector := bson.M{"_id": OnePageAdvisory.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": OnePageAdvisory}
	_, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOnePageAdvisory : ""
func (d *Daos) FilterOnePageAdvisory(ctx *models.Context, OnePageAdvisoryfilter *models.OnePageAdvisoryFilter, pagination *models.Pagination) ([]models.RefOnePageAdvisory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if OnePageAdvisoryfilter != nil {

		if len(OnePageAdvisoryfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": OnePageAdvisoryfilter.ActiveStatus}})
		}
		if len(OnePageAdvisoryfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": OnePageAdvisoryfilter.Status}})
		}
		//Regex
		if OnePageAdvisoryfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: OnePageAdvisoryfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if OnePageAdvisoryfilter != nil {
		if OnePageAdvisoryfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{OnePageAdvisoryfilter.SortBy: OnePageAdvisoryfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("OnePageAdvisory query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OnePageAdvisorys []models.RefOnePageAdvisory
	if err = cursor.All(context.TODO(), &OnePageAdvisorys); err != nil {
		return nil, err
	}
	return OnePageAdvisorys, nil
}

//EnableOnePageAdvisory :""
func (d *Daos) EnableOnePageAdvisory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ONEPAGEADVISORYSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableOnePageAdvisory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ONEPAGEADVISORYSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOnePageAdvisory :""
func (d *Daos) DeleteOnePageAdvisory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ONEPAGEADVISORYSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
