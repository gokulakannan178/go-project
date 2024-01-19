package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveLiveVideo :""
func (d *Daos) SaveLiveVideo(ctx *models.Context, liveVideo *models.LiveVideo) error {
	_, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).InsertOne(ctx.CTX, liveVideo)
	return err
}

//GetSingleLiveVideo : ""
func (d *Daos) GetSingleLiveVideo(ctx *models.Context, UniqueID string) (*models.RefLiveVideo, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var liveVideos []models.RefLiveVideo
	var liveVideo *models.RefLiveVideo
	if err = cursor.All(ctx.CTX, &liveVideos); err != nil {
		return nil, err
	}
	if len(liveVideos) > 0 {
		liveVideo = &liveVideos[0]
	}
	return liveVideo, nil
}

//UpdateLiveVideo : ""
func (d *Daos) UpdateLiveVideo(ctx *models.Context, liveVideo *models.LiveVideo) error {
	selector := bson.M{"uniqueId": liveVideo.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": liveVideo, "$push": bson.M{"updatedLog": update}}
	_, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterLiveVideo : ""
func (d *Daos) FilterLiveVideo(ctx *models.Context, liveVideofilter *models.LiveVideoFilter, pagination *models.Pagination) ([]models.RefLiveVideo, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if liveVideofilter != nil {
		if len(liveVideofilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": liveVideofilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("liveVideo query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var liveVideos []models.RefLiveVideo
	if err = cursor.All(context.TODO(), &liveVideos); err != nil {
		return nil, err
	}
	return liveVideos, nil
}

//EnableLiveVideo :""
func (d *Daos) EnableLiveVideo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LIVEVIDEOSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLiveVideo :""
func (d *Daos) DisableLiveVideo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LIVEVIDEOSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLiveVideo :""
func (d *Daos) DeleteLiveVideo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LIVEVIDEOSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLLIVEVIDEO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
