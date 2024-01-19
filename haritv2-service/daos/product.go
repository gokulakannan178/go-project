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

//SaveProduct :""
func (d *Daos) SaveProduct(ctx *models.Context, product *models.Product) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).InsertOne(ctx.CTX, product)
	return err
}

//GetSingleProduct : ""
func (d *Daos) GetSingleProduct(ctx *models.Context, uniqueID string) (*models.RefProduct, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCTCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var products []models.RefProduct
	var product *models.RefProduct
	if err = cursor.All(ctx.CTX, &products); err != nil {
		return nil, err
	}
	if len(products) > 0 {
		product = &products[0]
	}
	return product, nil
}

//UpdateProduct : ""
func (d *Daos) UpdateProduct(ctx *models.Context, product *models.Product) error {
	selector := bson.M{"uniqueId": product.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": product}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProduct :""
func (d *Daos) EnableProduct(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProduct :""
func (d *Daos) DisableProduct(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProduct :""
func (d *Daos) DeleteProduct(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterProduct : ""
func (d *Daos) FilterProduct(ctx *models.Context, productfilter *models.ProductFilter, pagination *models.Pagination) ([]models.RefProduct, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if productfilter != nil {
		if len(productfilter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": productfilter.Name}})
		}
		if len(productfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": productfilter.Status}})
		}
		if len(productfilter.CategoryID) > 0 {
			query = append(query, bson.M{"categoryId": bson.M{"$in": productfilter.CategoryID}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if productfilter != nil {
		if productfilter.SortField != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{productfilter.SortField: productfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCTCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("product query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var products []models.RefProduct
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

//GetDefaultProduct : ""
func (d *Daos) GetDefaultProduct(ctx *models.Context) (*models.RefProduct, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var products []models.RefProduct
	var product *models.RefProduct
	if err = cursor.All(ctx.CTX, &products); err != nil {
		return nil, err
	}
	if len(products) > 0 {
		product = &products[0]
	}
	return product, nil
}
