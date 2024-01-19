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

// SaveUserWardAccess : ""
func (d *Daos) SaveUserWardAccess(ctx *models.Context, block *models.UserWardAccess) error {
	d.Shared.BsonToJSONPrint(block)
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).InsertOne(ctx.CTX, block)
	return err
}

// GetSingleUserWardAccess  : ""
func (d *Daos) GetSingleUserWardAccess(ctx *models.Context, UniqueID string) (*models.RefUserWardAccess, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserWardAccess
	var tower *models.RefUserWardAccess
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserWardAccess: ""
func (d *Daos) UpdateUserWardAccess(ctx *models.Context, crop *models.UserWardAccess) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserWardAccess : ""
func (d *Daos) EnableUserWardAccess(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERWARDACCESSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserWardAccess : ""
func (d *Daos) DisableUserWardAccess(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERWARDACCESSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserWardAccess : ""
func (d *Daos) DeleteUserWardAccess(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERWARDACCESSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterUserWardAccess(ctx *models.Context, filter *models.UserWardAccessFilter, pagination *models.Pagination) ([]models.RefUserWardAccess, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERWARDACCESS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var prop []models.RefUserWardAccess
	if err = cursor.All(context.TODO(), &prop); err != nil {
		return nil, err
	}
	return prop, nil
}
