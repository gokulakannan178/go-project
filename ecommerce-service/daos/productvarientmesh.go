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

// SaveProductVariantMesh : ""
func (d *Daos) SaveProductVariantMesh(ctx *models.Context, product *models.ProductVariantMesh) error {
	d.Shared.BsonToJSONPrint(product)
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).InsertOne(ctx.CTX, product)
	return err
}

// GetSingleProductVariantMesh : ""
func (d *Daos) GetSingleProductVariantMesh(ctx *models.Context, UniqueID string) (*models.RefProductVariantMesh, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefProductVariantMesh
	var tower *models.RefProductVariantMesh
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateProductVariantMesh : ""
func (d *Daos) UpdateProductVariantMesh(ctx *models.Context, crop *models.ProductVariantMesh) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableProductVariantMesh : ""
func (d *Daos) EnableProductVariantMesh(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTVARIANTMESHSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableProductVariantMesh : ""
func (d *Daos) DisableProductVariantMesh(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTVARIANTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteProductVariantMesh : ""
func (d *Daos) DeleteProductVariantMesh(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTVARIANTMESHSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterProductVariantMesh : ""
func (d *Daos) FilterProductVariantMesh(ctx *models.Context, filter *models.ProductVariantMeshFilter, pagination *models.Pagination) ([]models.RefProductVariantMesh, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTVARIANTMESH).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var product []models.RefProductVariantMesh
	if err = cursor.All(context.TODO(), &product); err != nil {
		return nil, err
	}
	return product, nil
}
