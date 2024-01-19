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
func (d *Daos) SaveUserLoginLog(ctx *models.Context, UserLoginLog *models.UserLoginLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).InsertOne(ctx.CTX, UserLoginLog)
	if err != nil {
		return err
	}
	UserLoginLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleUserLoginLog(ctx *models.Context, UniqueID string) (*models.RefUserLoginLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserLoginLogs []models.RefUserLoginLog
	var UserLoginLog *models.RefUserLoginLog
	if err = cursor.All(ctx.CTX, &UserLoginLogs); err != nil {
		return nil, err
	}
	if len(UserLoginLogs) > 0 {
		UserLoginLog = &UserLoginLogs[0]
	}
	return UserLoginLog, nil
}

//UpdateUserLoginLog : ""
func (d *Daos) UpdateUserLoginLog(ctx *models.Context, UserLoginLog *models.UserLoginLog) error {

	selector := bson.M{"_id": UserLoginLog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": UserLoginLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterUserLoginLog : ""
func (d *Daos) FilterUserLoginLog(ctx *models.Context, UserLoginLogfilter *models.UserLoginLogFilter, pagination *models.Pagination) ([]models.RefUserLoginLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if UserLoginLogfilter != nil {

		if len(UserLoginLogfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": UserLoginLogfilter.ActiveStatus}})
		}
		if len(UserLoginLogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": UserLoginLogfilter.Status}})
		}
		//Regex
		if UserLoginLogfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: UserLoginLogfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if UserLoginLogfilter != nil {
		if UserLoginLogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{UserLoginLogfilter.SortBy: UserLoginLogfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("UserLoginLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserLoginLogs []models.RefUserLoginLog
	if err = cursor.All(context.TODO(), &UserLoginLogs); err != nil {
		return nil, err
	}
	return UserLoginLogs, nil
}

//EnableUserLoginLog :""
func (d *Daos) EnableUserLoginLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERLOGINLOGSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableUserLoginLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERLOGINLOGSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteUserLoginLog :""
func (d *Daos) DeleteUserLoginLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERLOGINLOGSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) UpdateUserLogout(ctx *models.Context, UserLoginLog *models.UserLoginLog) error {

	selector := bson.M{"userId": UserLoginLog.UserId}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"logOutTime": t, "status": constants.USERLOGINLOGSTATUSLOGOUT}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) LogoutUserLoginLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	t := time.Now()
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERLOGINLOGSTATUSLOGOUT, "logOutTime": t}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleUserLoginLogWithStatus(ctx *models.Context, UniqueID string) (*models.RefUserLoginLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userId": id, "status": constants.USERLOGINLOGSTATUSLOGIN}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("UserLoginLog query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOGINLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserLoginLogs []models.RefUserLoginLog
	var UserLoginLog *models.RefUserLoginLog
	if err = cursor.All(ctx.CTX, &UserLoginLogs); err != nil {
		return nil, err
	}
	if len(UserLoginLogs) > 0 {
		UserLoginLog = &UserLoginLogs[0]
	}
	return UserLoginLog, nil
}
