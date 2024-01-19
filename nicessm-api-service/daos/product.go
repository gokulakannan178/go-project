package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveProduct :""
func (d *Daos) SaveProduct(ctx *models.Context, product *models.Product) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).InsertOne(ctx.CTX, product)
	if err != nil {
		return err
	}
	product.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleProduct : ""
func (d *Daos) GetSingleProduct(ctx *models.Context, UniqueID string) (*models.RefProduct, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var productS []models.RefProduct
	var product *models.RefProduct
	if err = cursor.All(ctx.CTX, &productS); err != nil {
		return nil, err
	}
	if len(productS) > 0 {
		product = &productS[0]
	}
	return product, nil
}

//UpdateProduct : ""
func (d *Daos) UpdateProduct(ctx *models.Context, product *models.Product) error {

	selector := bson.M{"_id": product.ID}
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

//FilterProduct : ""
func (d *Daos) FilterProduct(ctx *models.Context, filter *models.ProductFilter, pagination *models.Pagination) ([]models.RefProduct, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": filter.ActiveStatus}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		//Regex
		if filter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchBox.Name, Options: "xi"}})
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Product query =>", mainPipeline)
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

//EnableProduct :""
func (d *Daos) EnableProduct(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProduct :""
func (d *Daos) DisableProduct(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProduct :""
func (d *Daos) DeleteProduct(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPRODUCT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) UpdateProductQuantity(ctx *models.Context, ProductID primitive.ObjectID, Quantity float64) {

	selector := bson.M{"_id": ProductID}
	update := bson.M{"$inc": bson.M{"quantity": (Quantity * -1)}}
	res := ctx.DB.Collection(constants.COLLECTIONPRODUCT).FindOneAndUpdate(ctx.CTX, selector, update)
	log.Println(res)

}
