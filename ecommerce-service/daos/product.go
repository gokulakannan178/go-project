package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveProduct : ""
func (d *Daos) SaveProduct(ctx *models.Context, block *models.Product) error {
	d.Shared.BsonToJSONPrint(block)
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).InsertOne(ctx.CTX, block)
	return err
}

// GetSingleProduct : ""
func (d *Daos) GetSingleProduct(ctx *models.Context, UniqueID string) (*models.RefProduct, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONINVENTORY,
			"as":   "ref.inventory",
			"let":  bson.M{"productId": "$uniqueId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []interface{}{"$productId", "$$productId"}},
					bson.M{"$eq": []interface{}{"$status", "Active"}},
				}}},
			},
			},
		},
	},
	)
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.inventory": bson.M{"$arrayElemAt": []interface{}{"$ref.inventory", 0}}}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefProduct
	var tower *models.RefProduct
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateProduct : ""
func (d *Daos) UpdateProduct(ctx *models.Context, crop *models.Product) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableProduct : ""
func (d *Daos) EnableProduct(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableProduct : ""
func (d *Daos) DisableProduct(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteProduct : ""
func (d *Daos) DeleteProduct(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterProduct : ""
func (d *Daos) FilterProduct(ctx *models.Context, filter *models.ProductFilter, pagination *models.Pagination) ([]models.RefProduct, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.CategoryID) > 0 {
			query = append(query, bson.M{"categoryId": bson.M{"$in": filter.CategoryID}})
		}
		if len(filter.SubCategoryID) > 0 {
			query = append(query, bson.M{"subCategoryId": bson.M{"$in": filter.SubCategoryID}})
		}

		if filter.SearchText.ProductName != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchText.ProductName, Options: "xi"}})

		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})

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
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONINVENTORY,
			"as":   "ref.inventory",
			"let":  bson.M{"productId": "$uniqueId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []interface{}{"$productId", "$$productId"}},
					bson.M{"$eq": []interface{}{"$status", "Active"}},
				}}},
			},
			},
		},
	},
	)
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.inventory": bson.M{"$arrayElemAt": []interface{}{"$ref.inventory", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockCrop []models.RefProduct
	if err = cursor.All(context.TODO(), &blockCrop); err != nil {
		return nil, err
	}
	return blockCrop, nil
}
