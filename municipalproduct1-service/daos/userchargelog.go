package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveUserChargeLog : ""
func (d *Daos) SaveUserChargeLog(ctx *models.Context, UserChargeLog *models.UserChargeLog) error {
	d.Shared.BsonToJSONPrint(UserChargeLog)
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).InsertOne(ctx.CTX, UserChargeLog)
	return err
}

// GetSingleUserChargeLog : ""
func (d *Daos) GetSingleUserChargeLog(ctx *models.Context, UniqueID string) (*models.RefUserChargeLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserChargeLog
	var tower *models.RefUserChargeLog
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserChargeLog : ""
func (d *Daos) UpdateUserChargeLog(ctx *models.Context, business *models.UserChargeLog) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserChargeLog : ""
func (d *Daos) EnableUserChargeLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGELOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserChargeLog : ""
func (d *Daos) DisableUserChargeLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGELOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserChargeLog : ""
func (d *Daos) DeleteUserChargeLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGELOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserChargeLog : ""
func (d *Daos) FilterUserChargeLog(ctx *models.Context, filter *models.UserChargeLogFilter, pagination *models.Pagination) ([]models.RefUserChargeLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).CountDocuments(ctx.CTX, func() bson.M {
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
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserChargeLog []models.RefUserChargeLog
	if err = cursor.All(context.TODO(), &UserChargeLog); err != nil {
		return nil, err
	}
	return UserChargeLog, nil
}
