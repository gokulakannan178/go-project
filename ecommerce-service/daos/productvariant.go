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
)

// SaveProductVariant : ""
func (d *Daos) SaveProductVariant(ctx *models.Context, block *models.ProductVariant) error {
	d.Shared.BsonToJSONPrint(block)
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).InsertOne(ctx.CTX, block)
	return err
}

// GetSingleProductVariant : ""
func (d *Daos) GetSingleProductVariant(ctx *models.Context, UniqueID string) (*models.RefProductVariant, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCT, "productId", "uniqueId", "ref.product", "ref.product")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "ref.product.categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "ref.product.subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefProductVariant
	var tower *models.RefProductVariant
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateProductVariant : ""
func (d *Daos) UpdateProductVariant(ctx *models.Context, crop *models.ProductVariant) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableProductVariant : ""
func (d *Daos) EnableProductVariant(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTVARIANTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableProductVariant : ""
func (d *Daos) DisableProductVariant(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTVARIANTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteProductVariant : ""
func (d *Daos) DeleteProductVariant(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTVARIANTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterProductVariant : ""
func (d *Daos) FilterProductVariant(ctx *models.Context, filter *models.ProductVariantFilter, pagination *models.Pagination) ([]models.RefProductVariant, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.ProductID) > 0 {
			query = append(query, bson.M{"productId": bson.M{"$in": filter.ProductID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCT, "productId", "uniqueId", "ref.product", "ref.product")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "ref.product.categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "ref.product.subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockCrop []models.RefProductVariant
	if err = cursor.All(context.TODO(), &blockCrop); err != nil {
		return nil, err
	}
	return blockCrop, nil
}

// UpsertProductVarients : ""
// func (d *Daos) UpsertProductVarients(ctx *models.Context, varients *[]models.ProductVariant) error {
// 	d.Shared.BsonToJSONPrint(varients)

// 	opts := options.Update().SetUpsert(true)
// 	for _, v := range *varients {
// 		query := bson.M{"name": v.Name, "productId": v.ProductID, "variantTypeId": v.VariantTypeID}
// 		update := bson.M{"$set": v}
// 		_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).UpdateOne(ctx.CTX, query, update, opts)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
func (d *Daos) ProductVariantRegister(ctx *models.Context, block *models.RegProductVariant) error {
	d.Shared.BsonToJSONPrint(block)
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).InsertOne(ctx.CTX, block)
	return err
}
func (d *Daos) ProductVariantMeshRegister(ctx *models.Context, product *models.RegProductVariant) error {
	var tempProductVarient []interface{}
	for _, v := range product.Mesh {

		tempProductVarient = append(tempProductVarient, v)
	}

	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).InsertMany(ctx.CTX, tempProductVarient)
	return err
}
func (d *Daos) GetMyInventory(ctx *models.Context, filter *models.ProductVariantInventoryFilter, pagination *models.Pagination) ([]models.RefProductVariant, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONINVENTORY, "uniqueId", "productVarientId", "ref.inventory", "ref.inventory")...)
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONINVENTORY,
		"as":   "ref.inventory",
		"let":  bson.M{"varientId": "$uniqueId"},
		"pipeline": []bson.M{

			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$productVarientId", "$$varientId"}},
			}}}}}},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.inventory": bson.M{"$arrayElemAt": []interface{}{"$ref.inventory", 0}}}})
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": "Active"}})
		}
		if len(filter.ProductID) > 0 {
			query = append(query, bson.M{"productId": bson.M{"$in": filter.ProductID}})

		}
		if len(filter.VendorID) > 0 {
			query = append(query, bson.M{"ref.inventory.vendorId": bson.M{"$in": filter.VendorID}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("KD pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCT, "productId", "uniqueId", "ref.product", "ref.product")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "ref.product.categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "ref.product.subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockCrop []models.RefProductVariant
	if err = cursor.All(context.TODO(), &blockCrop); err != nil {
		return nil, err
	}
	return blockCrop, nil
}
