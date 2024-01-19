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

// SaveUserChargeCategory : ""
func (d *Daos) SaveUserChargeCategory(ctx *models.Context, userchargecategory *models.UserChargeCategory) error {
	d.Shared.BsonToJSONPrint(userchargecategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).InsertOne(ctx.CTX, userchargecategory)
	return err
}

// GetSingleUserChargeCategory : ""
func (d *Daos) GetSingleUserChargeCategory(ctx *models.Context, UniqueID string) (*models.RefUserChargeCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserChargeCategory
	var tower *models.RefUserChargeCategory
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserChargeCategory : ""
func (d *Daos) UpdateUserChargeCategory(ctx *models.Context, business *models.UserChargeCategory) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserChargeCategory : ""
func (d *Daos) EnableUserChargeCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGECATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserChargeCategory : ""
func (d *Daos) DisableUserChargeCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGECATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserChargeCategory : ""
func (d *Daos) DeleteUserChargeCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGECATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserChargeCategory : ""
func (d *Daos) FilterUserChargeCategory(ctx *models.Context, filter *models.UserChargeCategoryFilter, pagination *models.Pagination) ([]models.RefUserChargeCategory, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var userchargecategory []models.RefUserChargeCategory
	if err = cursor.All(context.TODO(), &userchargecategory); err != nil {
		return nil, err
	}
	return userchargecategory, nil
}
