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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveEmailLog :""
func (d *Daos) SaveEmailLog(ctx *models.Context, emaillog *models.EmailLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).InsertOne(ctx.CTX, emaillog)
	if err != nil {
		return err
	}
	emaillog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleEmailLog : ""
func (d *Daos) GetSingleEmailLog(ctx *models.Context, uniqueID string) (*models.RefEmailLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var emaillogs []models.RefEmailLog
	var emaillog *models.RefEmailLog
	if err = cursor.All(ctx.CTX, &emaillogs); err != nil {
		return nil, err
	}
	if len(emaillogs) > 0 {
		emaillog = &emaillogs[0]
	}
	return emaillog, nil
}

//UpdateEmailLog : ""
func (d *Daos) UpdateEmailLog(ctx *models.Context, emaillog *models.EmailLog) error {

	selector := bson.M{"uniqueId": emaillog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": emaillog}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableEmailLog :""
func (d *Daos) EnableEmailLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMAILLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEmailLog :""
func (d *Daos) DisableEmailLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMAILLOGSTATUSDISABLE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEmailLog :""
func (d *Daos) DeleteEmailLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMAILLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterEmailLog : ""
func (d *Daos) FilterEmailLog(ctx *models.Context, emaillogfilter *models.EmailLogFilter, pagination *models.Pagination) ([]models.RefEmailLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if emaillogfilter != nil {

		if len(emaillogfilter.IsJob) > 0 {
			query = append(query, bson.M{"isJob": bson.M{"$in": emaillogfilter.IsJob}})
		}
		if len(emaillogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": emaillogfilter.Status}})
		}
		if len(emaillogfilter.Email) > 0 {
			query = append(query, bson.M{"to.email": bson.M{"$in": emaillogfilter.Email}})
		}
		//Regex
		if emaillogfilter.Regex.SentFor != "" {
			query = append(query, bson.M{"sentFor": primitive.Regex{Pattern: emaillogfilter.Regex.SentFor, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if emaillogfilter != nil {
		if emaillogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{emaillogfilter.SortBy: emaillogfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmailLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMAILLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var emaillogs []models.RefEmailLog
	if err = cursor.All(context.TODO(), &emaillogs); err != nil {
		return nil, err
	}
	return emaillogs, nil
}
