package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveDisease :""
func (d *Daos) SaveUserAcl(ctx *models.Context, useracl *models.UserAcl) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERACL).InsertOne(ctx.CTX, useracl)
	if err != nil {
		return err
	}
	useracl.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveUserAclWithUpsert(ctx *models.Context, useracl *models.UserAcl) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"userType": useracl.UserType}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("useracl query =>", updateQuery)
	updateData := bson.M{"$set": useracl}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERACL).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleDisease : ""
func (d *Daos) GetSingleUserAcl(ctx *models.Context, UniqueID string) (*models.RefUserAcl, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERACL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var useracls []models.RefUserAcl
	var useracl *models.RefUserAcl
	if err = cursor.All(ctx.CTX, &useracls); err != nil {
		return nil, err
	}
	if len(useracls) > 0 {
		useracl = &useracls[0]
	}
	return useracl, nil
}

//UpdateUserAcl : ""
func (d *Daos) UpdateUserAcl(ctx *models.Context, useracl *models.UserAcl) error {

	selector := bson.M{"_id": useracl.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": useracl}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERACL).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterUserAcl : ""
func (d *Daos) FilterUserAcl(ctx *models.Context, useraclfilter *models.UserAclFilter, pagination *models.Pagination) ([]models.RefUserAcl, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if useraclfilter != nil {

		if len(useraclfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": useraclfilter.Status}})
		}
		if len(useraclfilter.UserType) > 0 {
			query = append(query, bson.M{"userType": bson.M{"$in": useraclfilter.UserType}})
		}
		if len(useraclfilter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": useraclfilter.UserName}})
		}
		//Regex
		// if useraclfilter.SearchBox.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: useraclfilter.SearchBox.Name, Options: "xi"}})
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if useraclfilter != nil {
		if useraclfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{useraclfilter.SortBy: useraclfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERACL).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("useracl query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERACL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var useracls []models.RefUserAcl
	if err = cursor.All(context.TODO(), &useracls); err != nil {
		return nil, err
	}
	return useracls, nil
}

//EnableUserAcl :""
func (d *Daos) EnableUserAcl(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERACLSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERACL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableUserAcl(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERACLSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERACL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteUserAcl :""
func (d *Daos) DeleteUserAcl(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERACLSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSERACL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleUserAclWithUserType(ctx *models.Context, UniqueID string) (*models.RefUserAcl, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userType": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERACL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var useracls []models.RefUserAcl
	var useracl *models.RefUserAcl
	if err = cursor.All(ctx.CTX, &useracls); err != nil {
		return nil, err
	}
	if len(useracls) > 0 {
		useracl = &useracls[0]
	}
	return useracl, nil
}
