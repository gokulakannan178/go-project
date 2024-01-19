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

//SaveOffenceVideo :""
func (d *Daos) SaveOffenceVideo(ctx *models.Context, offenceVideo *models.OffenceVideo) error {
	_, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).InsertOne(ctx.CTX, offenceVideo)
	return err
}

//GetSingleOffenceVideo : ""
func (d *Daos) GetSingleOffenceVideo(ctx *models.Context, UniqueID string) (*models.RefOffenceVideo, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offenceVideos []models.RefOffenceVideo
	var offenceVideo *models.RefOffenceVideo
	if err = cursor.All(ctx.CTX, &offenceVideos); err != nil {
		return nil, err
	}
	if len(offenceVideos) > 0 {
		offenceVideo = &offenceVideos[0]
	}
	return offenceVideo, nil
}

//UpdateOffenceVideo : ""
func (d *Daos) UpdateOffenceVideo(ctx *models.Context, offenceVideo *models.OffenceVideo) error {
	selector := bson.M{"uniqueId": offenceVideo.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": offenceVideo, "$push": bson.M{"updatedLog": update}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOffenceVideo : ""
func (d *Daos) FilterOffenceVideo(ctx *models.Context, offenceVideofilter *models.OffenceVideoFilter, pagination *models.Pagination) ([]models.RefOffenceVideo, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if offenceVideofilter != nil {
		if len(offenceVideofilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": offenceVideofilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("offenceVideo query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offenceVideos []models.RefOffenceVideo
	if err = cursor.All(context.TODO(), &offenceVideos); err != nil {
		return nil, err
	}
	return offenceVideos, nil
}

//EnableOffenceVideo :""
func (d *Daos) EnableOffenceVideo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFENCEVIDEOSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOffenceVideo :""
func (d *Daos) DisableOffenceVideo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFENCEVIDEOSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOffenceVideo :""
func (d *Daos) DeleteOffenceVideo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFENCEVIDEOSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLOFFENCEVIDEO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
