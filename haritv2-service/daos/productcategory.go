package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveProductCategory :""
func (d *Daos) SaveProductCategory(ctx *models.Context, pc *models.ProductCategory) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).InsertOne(ctx.CTX, pc)
	return err
}

//GetSingleProductCategory : ""
func (d *Daos) GetSingleProductCategory(ctx *models.Context, uniqueID string) (*models.RefProductCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pcs []models.RefProductCategory
	var pc *models.RefProductCategory
	if err = cursor.All(ctx.CTX, &pcs); err != nil {
		return nil, err
	}
	if len(pcs) > 0 {
		pc = &pcs[0]
	}
	return pc, nil
}

//UpdateProductCategory : ""
func (d *Daos) UpdateProductCategory(ctx *models.Context, pc *models.ProductCategory) error {
	selector := bson.M{"uniqueId": pc.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": pc}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProductCategory :""
func (d *Daos) EnableProductCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProductCategory :""
func (d *Daos) DisableProductCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProductCategory :""
func (d *Daos) DeleteProductCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterProductCategory : ""
func (d *Daos) FilterProductCategory(ctx *models.Context, filter *models.ProductCategoryFilter, pagination *models.Pagination) ([]models.RefProductCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": filter.Name}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("pc query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pcs []models.RefProductCategory
	if err = cursor.All(context.TODO(), &pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

//GetDefaultProduct : ""
func (d *Daos) GetDefaultProductCategory(ctx *models.Context) (*models.RefProductCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pcs []models.RefProductCategory
	var pc *models.RefProductCategory
	if err = cursor.All(ctx.CTX, &pcs); err != nil {
		return nil, err
	}
	if len(pcs) > 0 {
		pc = &pcs[0]
	}
	return pc, nil
}
