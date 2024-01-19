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

// SaveExpenseCategory : ""
func (d *Daos) SaveExpenseCategory(ctx *models.Context, expensecategory *models.ExpenseCategory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).InsertOne(ctx.CTX, expensecategory)
	if err != nil {
		return err
	}
	expensecategory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateExpenseCategory : ""
func (d *Daos) UpdateExpenseCategory(ctx *models.Context, ExpenseCategory *models.ExpenseCategory) error {
	selector := bson.M{"uniqueId": ExpenseCategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ExpenseCategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleExpenseCategory : ""
func (d *Daos) GetSingleExpenseCategory(ctx *models.Context, uniqueID string) (*models.RefExpenseCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensecategorys []models.RefExpenseCategory
	var expensecategory *models.RefExpenseCategory
	if err = cursor.All(ctx.CTX, &expensecategorys); err != nil {
		return nil, err
	}
	if len(expensecategorys) > 0 {
		expensecategory = &expensecategorys[0]
	}
	return expensecategory, err
}

// GetSingleExpenseCategory : ""
func (d *Daos) GetSingleActiveExpenseCategoryWithName(ctx *models.Context, uniqueID string) (*models.RefExpenseCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.EXPENSECATEGORYSTATUSACTIVE}})

	d.Shared.BsonToJSONPrintTag("GetSingleActiveExpenseCategoryWithName query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensecategorys []models.RefExpenseCategory
	var expensecategory *models.RefExpenseCategory
	if err = cursor.All(ctx.CTX, &expensecategorys); err != nil {
		return nil, err
	}
	if len(expensecategorys) > 0 {
		expensecategory = &expensecategorys[0]
	}
	return expensecategory, err
}

// GetSingleExpenseCategory : ""
func (d *Daos) GetSingleExpenseCategoryWithActiveStatus(ctx *models.Context, uniqueID string, Status string) (*models.RefExpenseCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONONBOARDINGCHECKLIST, "uniqueId", "expensecategoryId", "ref.onboardingchecklist", "ref.onboardingchecklist")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensecategorys []models.RefExpenseCategory
	var expensecategory *models.RefExpenseCategory
	if err = cursor.All(ctx.CTX, &expensecategorys); err != nil {
		return nil, err
	}
	if len(expensecategorys) > 0 {
		expensecategory = &expensecategorys[0]
	}
	return expensecategory, err
}

// EnableExpenseCategory : ""
func (d *Daos) EnableExpenseCategory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EXPENSECATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableExpenseCategory : ""
func (d *Daos) DisableExpenseCategory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EXPENSECATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteExpenseCategory :""
func (d *Daos) DeleteExpenseCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EXPENSECATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterExpenseCategory : ""
func (d *Daos) FilterExpenseCategory(ctx *models.Context, expenseCategory *models.FilterExpenseCategory, pagination *models.Pagination) ([]models.RefExpenseCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if expenseCategory != nil {
		if len(expenseCategory.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": expenseCategory.Status}})
		}
		if len(expenseCategory.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": expenseCategory.OrganisationID}})
		}
		//Regex
		if expenseCategory.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: expenseCategory.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if expenseCategory != nil {
		if expenseCategory.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{expenseCategory.SortBy: expenseCategory.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensecategoryFilter []models.RefExpenseCategory
	if err = cursor.All(context.TODO(), &expensecategoryFilter); err != nil {
		return nil, err
	}
	return expensecategoryFilter, nil
}
