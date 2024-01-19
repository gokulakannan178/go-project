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

//SaveCart :""
func (d *Daos) SaveCart(ctx *models.Context, cart *models.Cart) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).InsertOne(ctx.CTX, cart)
	return err
}

//GetSingleCart : ""
func (d *Daos) GetSingleCart(ctx *models.Context, uniqueID string) (*models.RefCart, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var carts []models.RefCart
	var cart *models.RefCart
	if err = cursor.All(ctx.CTX, &carts); err != nil {
		return nil, err
	}
	if len(carts) > 0 {
		cart = &carts[0]
	}
	return cart, nil
}

//UpdateCart : ""
func (d *Daos) UpdateCart(ctx *models.Context, cart *models.Cart) error {
	selector := bson.M{"uniqueId": cart.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": cart, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCart : ""
func (d *Daos) FilterCart(ctx *models.Context, cartfilter *models.CartFilter, pagination *models.Pagination) ([]models.RefCart, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if cartfilter != nil {
		if len(cartfilter.ULBID) > 0 {
			query = append(query, bson.M{"ulbId": bson.M{"$in": cartfilter.ULBID}})
		}
		if len(cartfilter.FPOID) > 0 {
			query = append(query, bson.M{"fpoId": bson.M{"$in": cartfilter.FPOID}})
		}
		if len(cartfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": cartfilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ulbId", "uniqueId", "ref.ulb", "ref.ulb")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "ref.ulb.uniqueId", "companyId", "ref.ulb.ref.inventory", "ref.ulb.ref.inventory")...)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.calculatedPrice": bson.M{"$multiply": []interface{}{"$ref.ulb.ref.inventory.price", "$quantity"}}}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCART).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("cart query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var carts []models.RefCart
	if err = cursor.All(context.TODO(), &carts); err != nil {
		return nil, err
	}
	return carts, nil
}

//EnableCart :""
func (d *Daos) EnableCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CARTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCart :""
func (d *Daos) DisableCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CARTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCart :""
func (d *Daos) DeleteCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CARTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
