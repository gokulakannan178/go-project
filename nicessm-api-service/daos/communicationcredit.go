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

//SaveCommunicationCredit :""
func (d *Daos) SaveCommunicationCredit(ctx *models.Context, communicationcredit *models.CommunicationCredit) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).InsertOne(ctx.CTX, communicationcredit)
	if err != nil {
		return err
	}
	communicationcredit.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleCommunicationCredit : ""
func (d *Daos) GetSingleCommunicationCredit(ctx *models.Context, UniqueID string) (*models.RefCommunicationCredit, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var communicationcredits []models.RefCommunicationCredit
	var communicationcredit *models.RefCommunicationCredit
	if err = cursor.All(ctx.CTX, &communicationcredits); err != nil {
		return nil, err
	}
	if len(communicationcredits) > 0 {
		communicationcredit = &communicationcredits[0]
	}
	return communicationcredit, nil
}

//UpdateCommunicationCredit : ""
func (d *Daos) UpdateCommunicationCredit(ctx *models.Context, communicationcredit *models.CommunicationCredit) error {

	selector := bson.M{"_id": communicationcredit.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": communicationcredit}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCommunicationCredit : ""
func (d *Daos) FilterCommunicationCredit(ctx *models.Context, communicationcreditfilter *models.CommunicationCreditFilter, pagination *models.Pagination) ([]models.RefCommunicationCredit, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if communicationcreditfilter != nil {

		if len(communicationcreditfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": communicationcreditfilter.Status}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("CommunicationCredit query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var communicationcredits []models.RefCommunicationCredit
	if err = cursor.All(context.TODO(), &communicationcredits); err != nil {
		return nil, err
	}
	return communicationcredits, nil
}

//EnableCommunicationCredit :""
func (d *Daos) EnableCommunicationCredit(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMUNICATIONCREDITSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommunicationCredit :""
func (d *Daos) DisableCommunicationCredit(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMUNICATIONCREDITSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommunicationCredit :""
func (d *Daos) DeleteCommunicationCredit(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMUNICATIONCREDITSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) UpdateCommunicationCreditWithBalance(ctx *models.Context, uniqueID string, communicationcredit float64) error {

	selector := bson.M{"uniqueID": uniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$inc": bson.M{"balanceCredit": communicationcredit}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetSingleCommunicationCreditWithUniqueId(ctx *models.Context, UniqueID string) (*models.RefCommunicationCredit, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueID": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDIT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var communicationcredits []models.RefCommunicationCredit
	var communicationcredit *models.RefCommunicationCredit
	if err = cursor.All(ctx.CTX, &communicationcredits); err != nil {
		return nil, err
	}
	if len(communicationcredits) > 0 {
		communicationcredit = &communicationcredits[0]
	}
	return communicationcredit, nil
}
