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
)

// SaveExpenseSubcategory : ""
func (d *Daos) SaveExpenseSubcategory(ctx *models.Context, expensesubcategory *models.ExpenseSubcategory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).InsertOne(ctx.CTX, expensesubcategory)
	if err != nil {
		return err
	}
	expensesubcategory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateExpenseSubcategory : ""
func (d *Daos) UpdateExpenseSubcategory(ctx *models.Context, expensesubcategory *models.ExpenseSubcategory) error {
	selector := bson.M{"uniqueId": expensesubcategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": expensesubcategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleExpenseSubcategory : ""
func (d *Daos) GetSingleExpenseSubcategory(ctx *models.Context, uniqueID string) (*models.RefExpenseSubcategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensesubcategorys []models.RefExpenseSubcategory
	var expensesubcategory *models.RefExpenseSubcategory
	if err = cursor.All(ctx.CTX, &expensesubcategorys); err != nil {
		return nil, err
	}
	if len(expensesubcategorys) > 0 {
		expensesubcategory = &expensesubcategorys[0]
	}
	return expensesubcategory, err
}

// GetSingleExpenseSubcategoryWithActive : ""
func (d *Daos) GetSingleExpenseSubcategoryWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefExpenseSubcategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensesubcategorys []models.RefExpenseSubcategory
	var expensesubcategory *models.RefExpenseSubcategory
	if err = cursor.All(ctx.CTX, &expensesubcategorys); err != nil {
		return nil, err
	}
	if len(expensesubcategorys) > 0 {
		expensesubcategory = &expensesubcategorys[0]
	}
	return expensesubcategory, err
}

// EnableExpenseSubcategory : ""
func (d *Daos) EnableExpenseSubcategory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EXPENSESUBCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableExpenseSubcategory : ""
func (d *Daos) DisableExpenseSubcategory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EXPENSESUBCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteExpenseSubcategory :""
func (d *Daos) DeleteExpenseSubcategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EXPENSESUBCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterExpenseSubcategory : ""
func (d *Daos) FilterExpenseSubcategory(ctx *models.Context, expensesubcategory *models.FilterExpenseSubcategory, pagination *models.Pagination) ([]models.RefExpenseSubcategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if expensesubcategory != nil {
		if len(expensesubcategory.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": expensesubcategory.Status}})
		}
		if len(expensesubcategory.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": expensesubcategory.OrganisationID}})
		}
		//Regex
		if expensesubcategory.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: expensesubcategory.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if expensesubcategory != nil {
		if expensesubcategory.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{expensesubcategory.SortBy: expensesubcategory.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSESUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensesubcategoryFilter []models.RefExpenseSubcategory
	if err = cursor.All(context.TODO(), &expensesubcategoryFilter); err != nil {
		return nil, err
	}
	return expensesubcategoryFilter, nil
}
