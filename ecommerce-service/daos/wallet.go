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

//SaveWallet :""
func (d *Daos) SaveWallet(ctx *models.Context, Wallet *models.Wallet) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWALLET).InsertOne(ctx.CTX, Wallet)
	if err != nil {
		return err
	}
	Wallet.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}



//GetSingleWallet : ""
func (d *Daos) GetSingleWallet(ctx *models.Context, UniqueID string) (*models.RefWallet, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWALLET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Wallets []models.RefWallet
	var Wallet *models.RefWallet
	if err = cursor.All(ctx.CTX, &Wallets); err != nil {
		return nil, err
	}
	if len(Wallets) > 0 {
		Wallet = &Wallets[0]
	}
	return Wallet, nil
}

//UpdateWallet : ""
func (d *Daos) UpdateWallet(ctx *models.Context, Wallet *models.Wallet) error {
	selector := bson.M{"uniqueId": Wallet.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Wallet, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterWallet : ""
func (d *Daos) FilterWallet(ctx *models.Context, Walletfilter *models.WalletFilter, pagination *models.Pagination) ([]models.RefWallet, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Walletfilter != nil {
		if len(Walletfilter.Regex.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": Walletfilter.Regex.UniqueID}})
		}
		if len(Walletfilter.Regex.UserType) > 0 {
			query = append(query, bson.M{"userType": bson.M{"$in": Walletfilter.Regex.UniqueID}})
		}
		if len(Walletfilter.Regex.UserID) > 0 {
			query = append(query, bson.M{"userId": bson.M{"$in": Walletfilter.Regex.UserID}})
		}
		if len(Walletfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Walletfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWALLET).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Wallet query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWALLET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Wallets []models.RefWallet
	if err = cursor.All(context.TODO(), &Wallets); err != nil {
		return nil, err
	}
	return Wallets, nil
}

//EnableWallet :""
func (d *Daos) EnableWallet(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WALLETSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableWallet :""
func (d *Daos) DisableWallet(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WALLETSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWallet :""
func (d *Daos) DeleteWallet(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WALLETSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
