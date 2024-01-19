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

// SaveExpenseCategoryList : ""
func (d *Daos) SaveExpenseCategoryList(ctx *models.Context, expensecategorylist *models.ExpenseCategoryList) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).InsertOne(ctx.CTX, expensecategorylist)
	if err != nil {
		return err
	}
	expensecategorylist.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveExpenseCategoryListUpdert(ctx *models.Context, expensecategorylist *models.ExpenseCategoryList) error {
	//fmt.Println("arrayValue", arrayValue)
	//	for _, v := range expensecategorylist.ExpenseCategoryListmasterID {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"categoryId": expensecategorylist.CategoryId, "subcategoryId": expensecategorylist.SubcategoryId}
	fmt.Println("updateQuery===>", updateQuery)

	//fmt.Println("present added =======>", AssetPolicyAssets.UniqueID)
	updateData := bson.M{"$set": expensecategorylist}
	if _, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	//s}
	return nil
}

// GetSingleExpenseCategoryList : ""
func (d *Daos) GetSingleExpenseCategoryList(ctx *models.Context, uniqueID string) (*models.RefExpenseCategoryList, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEXPENSESUBCATEGORY, "subcategoryId", "uniqueId", "ref.expenseSubcategory", "ref.expenseSubcategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEXPENSECATEGORY, "categoryId", "uniqueId", "ref.expensecategory", "ref.expensecategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensecategorylists []models.RefExpenseCategoryList
	var expensecategorylist *models.RefExpenseCategoryList
	if err = cursor.All(ctx.CTX, &expensecategorylists); err != nil {
		return nil, err
	}
	if len(expensecategorylists) > 0 {
		expensecategorylist = &expensecategorylists[0]
	}
	return expensecategorylist, err
}

//UpdateExpenseCategoryList : ""
func (d *Daos) UpdateExpenseCategoryList(ctx *models.Context, expensecategorylist *models.ExpenseCategoryList) error {
	selector := bson.M{"uniqueId": expensecategorylist.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": expensecategorylist}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// ExpenseCategoryListRemoveNotPresentValue : ""
func (d *Daos) ExpenseCategoryListRemoveNotPresentValue(ctx *models.Context, categoryId string, arrayValue []string) error {
	selector := bson.M{"categoryId": categoryId, "subcategoryId": bson.M{"$nin": arrayValue}}
	d.Shared.BsonToJSONPrintTag("selector query in onboarding checklist =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.EXPENSECATEGORYLISTSTATUSDELETED}}
	d.Shared.BsonToJSONPrintTag("data query in onboarding checklist =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).UpdateMany(ctx.CTX, selector, data)
	return err
}

// ExpenseCategoryListUpsert : ""
func (d *Daos) ExpenseCategoryListUpsert(ctx *models.Context, CategoryId string, arrayValue []string, name string) error {
	fmt.Println("arrayValue", arrayValue)
	for _, v := range arrayValue {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"categoryId": CategoryId, "subcategoryId": v}
		fmt.Println("updateQuery===>", updateQuery)
		expenseCategoryList := new(models.ExpenseCategoryList)
		expenseCategoryList.Status = constants.EXPENSECATEGORYLISTSTATUSACTIVE
		expenseCategoryList.Name = name
		expenseCategoryList.UniqueID = d.GetUniqueID(ctx, constants.COLLECTIONEXPENSECATEGORYLIST)
		fmt.Println("present added =======>", expenseCategoryList.UniqueID)
		updateData := bson.M{"$set": expenseCategoryList}
		if _, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

// EnableExpenseCategoryList : ""
func (d *Daos) EnableExpenseCategoryList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EXPENSECATEGORYLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableExpenseCategoryList : ""
func (d *Daos) DisableExpenseCategoryList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EXPENSECATEGORYLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteExpenseCategoryList :""
func (d *Daos) DeleteExpenseCategoryList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EXPENSECATEGORYLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterExpenseCategoryList : ""
func (d *Daos) FilterExpenseCategoryList(ctx *models.Context, expensecategorylist *models.FilterExpenseCategoryList, pagination *models.Pagination) ([]models.RefExpenseCategoryList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if expensecategorylist != nil {
		if len(expensecategorylist.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": expensecategorylist.Status}})
		}
		if len(expensecategorylist.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": expensecategorylist.OrganisationID}})
		}
		//Regex
		if expensecategorylist.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: expensecategorylist.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if expensecategorylist != nil {
		if expensecategorylist.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{expensecategorylist.SortBy: expensecategorylist.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEXPENSESUBCATEGORY, "subcategoryId", "uniqueId", "ref.expenseSubcategory", "ref.expenseSubcategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEXPENSECATEGORY, "categoryId", "uniqueId", "ref.expensecategory", "ref.expensecategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEXPENSECATEGORYLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var expensecategorylistFilter []models.RefExpenseCategoryList
	if err = cursor.All(context.TODO(), &expensecategorylistFilter); err != nil {
		return nil, err
	}
	return expensecategorylistFilter, nil
}
