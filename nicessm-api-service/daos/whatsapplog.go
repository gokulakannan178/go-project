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

//SaveWhatsappLog :""
func (d *Daos) SaveWhatsappLog(ctx *models.Context, whatsapplog *models.WhatsappLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).InsertOne(ctx.CTX, whatsapplog)
	if err != nil {
		return err
	}
	whatsapplog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleWhatsappLog : ""
func (d *Daos) GetSingleWhatsappLog(ctx *models.Context, UniqueID string) (*models.RefWhatsappLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var whatsapplogs []models.RefWhatsappLog
	var whatsapplog *models.RefWhatsappLog
	if err = cursor.All(ctx.CTX, &whatsapplogs); err != nil {
		return nil, err
	}
	if len(whatsapplogs) > 0 {
		whatsapplog = &whatsapplogs[0]
	}
	return whatsapplog, nil
}

//UpdateWhatsappLog : ""
func (d *Daos) UpdateWhatsappLog(ctx *models.Context, whatsapplog *models.WhatsappLog) error {

	selector := bson.M{"_id": whatsapplog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": whatsapplog}
	_, err := ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterWhatsappLog : ""
func (d *Daos) FilterWhatsappLog(ctx *models.Context, whatsapplogfilter *models.WhatsappLogFilter, pagination *models.Pagination) ([]models.RefWhatsappLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if whatsapplogfilter != nil {

		if len(whatsapplogfilter.IsJob) > 0 {
			query = append(query, bson.M{"isJob": bson.M{"$in": whatsapplogfilter.IsJob}})
		}
		if len(whatsapplogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": whatsapplogfilter.Status}})
		}
		if len(whatsapplogfilter.No) > 0 {
			query = append(query, bson.M{"to.no": bson.M{"$in": whatsapplogfilter.No}})
		}
		if len(whatsapplogfilter.UserName) > 0 {
			query = append(query, bson.M{"to.userName": bson.M{"$in": whatsapplogfilter.UserName}})
		}
		if len(whatsapplogfilter.UserType) > 0 {
			query = append(query, bson.M{"to.userType": bson.M{"$in": whatsapplogfilter.UserType}})
		}
		if len(whatsapplogfilter.Name) > 0 {
			query = append(query, bson.M{"to.name": bson.M{"$in": whatsapplogfilter.Name}})
		}
		//Regex
		if whatsapplogfilter.Regex.SentFor != "" {
			query = append(query, bson.M{"sentFor": primitive.Regex{Pattern: whatsapplogfilter.Regex.SentFor, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if whatsapplogfilter != nil {
		if whatsapplogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{whatsapplogfilter.SortBy: whatsapplogfilter.SortOrder}})

		}

	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("WhatsappLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var whatsapplogs []models.RefWhatsappLog
	if err = cursor.All(context.TODO(), &whatsapplogs); err != nil {
		return nil, err
	}
	return whatsapplogs, nil
}

//EnableWhatsappLog :""
func (d *Daos) EnableWhatsappLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WHATSAPPLOGSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableWhatsappLog :""
func (d *Daos) DisableWhatsappLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WHATSAPPLOGSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWhatsappLog :""
func (d *Daos) DeleteWhatsappLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WHATSAPPLOGSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWHATSAPPLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
