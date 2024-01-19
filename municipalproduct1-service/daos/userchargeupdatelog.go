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

// SaveUserChargeUpdateLog : ""
func (d *Daos) SaveUserChargeUpdateLog(ctx *models.Context, UserChargeUpdateLog *models.UserChargeUpdateLog) error {
	d.Shared.BsonToJSONPrint(UserChargeUpdateLog)
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).InsertOne(ctx.CTX, UserChargeUpdateLog)
	return err
}

// GetSingleUserChargeUpdateLog : ""
func (d *Daos) GetSingleUserChargeUpdateLog(ctx *models.Context, UniqueID string) (*models.RefUserChargeUpdateLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserChargeUpdateLog
	var tower *models.RefUserChargeUpdateLog
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserChargeUpdateLog : ""
func (d *Daos) UpdateUserChargeUpdateLog(ctx *models.Context, business *models.UserChargeUpdateLog) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserChargeUpdateLog : ""
func (d *Daos) EnableUserChargeUpdateLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGEUPDATELOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserChargeUpdateLog : ""
func (d *Daos) DisableUserChargeUpdateLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGEUPDATELOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserChargeUpdateLog : ""
func (d *Daos) DeleteUserChargeUpdateLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGEUPDATELOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserChargeUpdateLog : ""
func (d *Daos) FilterUserChargeUpdateLog(ctx *models.Context, filter *models.UserChargeUpdateLogFilter, pagination *models.Pagination) ([]models.RefUserChargeUpdateLog, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGEUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserChargeUpdateLog []models.RefUserChargeUpdateLog
	if err = cursor.All(context.TODO(), &UserChargeUpdateLog); err != nil {
		return nil, err
	}
	return UserChargeUpdateLog, nil
}
