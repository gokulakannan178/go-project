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

// SaveUserChargeRateMaster : ""
func (d *Daos) SaveUserChargeRateMaster(ctx *models.Context, userchargeratemaster *models.UserChargeRateMaster) error {
	d.Shared.BsonToJSONPrint(userchargeratemaster)
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).InsertOne(ctx.CTX, userchargeratemaster)
	return err
}

// GetSingleUserChargeRateMaster : ""
func (d *Daos) GetSingleUserChargeRateMaster(ctx *models.Context, UniqueID string) (*models.RefUserChargeRateMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserChargeRateMaster
	var tower *models.RefUserChargeRateMaster
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserChargeRateMaster : ""
func (d *Daos) UpdateUserChargeRateMaster(ctx *models.Context, business *models.UserChargeRateMaster) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserChargeRateMaster : ""
func (d *Daos) EnableUserChargeRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGERATEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserChargeRateMaster : ""
func (d *Daos) DisableUserChargeRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGERATEMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserChargeRateMaster : ""
func (d *Daos) DeleteUserChargeRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERCHARGERATEMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserChargeRateMaster : ""
func (d *Daos) FilterUserChargeRateMaster(ctx *models.Context, filter *models.UserChargeRateMasterFilter, pagination *models.Pagination) ([]models.RefUserChargeRateMaster, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var userchargeratemaster []models.RefUserChargeRateMaster
	if err = cursor.All(context.TODO(), &userchargeratemaster); err != nil {
		return nil, err
	}
	return userchargeratemaster, nil
}

func (d *Daos) GetSingleUserChargeRateMasterWithCategoryId(ctx *models.Context, UniqueID string) (*models.RefUserChargeRateMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"categoryId": UniqueID, "status": "Active"}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"doe": 1}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERCHARGERATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserChargeRateMaster
	var tower *models.RefUserChargeRateMaster
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}
