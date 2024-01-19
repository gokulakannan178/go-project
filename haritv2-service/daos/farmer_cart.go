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

//SaveFarmerCart :""
func (d *Daos) SaveFarmerCart(ctx *models.Context, farmerCart *models.FarmerCart) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).InsertOne(ctx.CTX, farmerCart)
	return err
}

//GetSingleFarmerCart : ""
func (d *Daos) GetSingleFarmerCart(ctx *models.Context, uniqueID string) (*models.RefFarmerCart, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ulbId", "uniqueId", "ref.ulb", "ref.ulb")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "ref.ulb.uniqueId", "companyId", "ref.ulb.ref.inventory", "ref.ulb.ref.inventory")...)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.calculatedPrice": bson.M{"$multiply": []interface{}{"$ref.ulb.ref.inventory.price", "$quantity"}}}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmerCarts []models.RefFarmerCart
	var farmerCart *models.RefFarmerCart
	if err = cursor.All(ctx.CTX, &farmerCarts); err != nil {
		return nil, err
	}
	if len(farmerCarts) > 0 {
		farmerCart = &farmerCarts[0]
	}
	return farmerCart, nil
}

//UpdateFarmerCart : ""
func (d *Daos) UpdateFarmerCart(ctx *models.Context, farmerCart *models.FarmerCart) error {
	selector := bson.M{"uniqueId": farmerCart.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": farmerCart, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableFarmerCart :""
func (d *Daos) EnableFarmerCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCARTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmerCart :""
func (d *Daos) DisableFarmerCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCARTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmerCart :""
func (d *Daos) DeleteFarmerCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCARTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterFarmerCart : ""
func (d *Daos) FilterFarmerCart(ctx *models.Context, farmerCartfilter *models.FarmerCartFilter, pagination *models.Pagination) ([]models.RefFarmerCart, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if farmerCartfilter != nil {
		if len(farmerCartfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": farmerCartfilter.UniqueID}})
		}
		if len(farmerCartfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": farmerCartfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ulbId", "uniqueId", "ref.ulb", "ref.ulb")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "ref.ulb.uniqueId", "companyId", "ref.ulb.ref.inventory", "ref.ulb.ref.inventory")...)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.calculatedPrice": bson.M{"$multiply": []interface{}{"$ref.ulb.ref.inventory.price", "$quantity"}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("farmerCart query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmerCarts []models.RefFarmerCart
	if err = cursor.All(context.TODO(), &farmerCarts); err != nil {
		return nil, err
	}
	return farmerCarts, nil
}
