package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SavePropertyWallet :""
func (d *Daos) SavePropertyWallet(ctx *models.Context, propertyWallet *models.PropertyWallet) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).InsertOne(ctx.CTX, propertyWallet)
	return err
}

//GetSinglePropertyWallet : ""
func (d *Daos) GetSinglePropertyWallet(ctx *models.Context, UniqueID string) (*models.RefPropertyWallet, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyWallets []models.RefPropertyWallet
	var propertyWallet *models.RefPropertyWallet
	if err = cursor.All(ctx.CTX, &propertyWallets); err != nil {
		return nil, err
	}
	if len(propertyWallets) > 0 {
		propertyWallet = &propertyWallets[0]
	}
	return propertyWallet, nil
}

//UpdatePropertyWallet : ""
func (d *Daos) UpdatePropertyWallet(ctx *models.Context, propertyWallet *models.PropertyWallet) error {
	selector := bson.M{"uniqueId": propertyWallet.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyWallet, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnablePropertyWallet :""
func (d *Daos) EnablePropertyWallet(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYWALLETSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyWallet :""
func (d *Daos) DisablePropertyWallet(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYWALLETSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyWallet :""
func (d *Daos) DeletePropertyWallet(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYWALLETSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterPropertyWallet : ""
func (d *Daos) FilterPropertyWallet(ctx *models.Context, filter *models.PropertyWalletFilter, pagination *models.Pagination) ([]models.RefPropertyWallet, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if filter.SearchText.MobileNo != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.SearchText.MobileNo, Options: "xi"}})

		}
		if filter.SearchText.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.SearchText.OwnerName, Options: "xi"}})

		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})

		}
		if filter.Amount != nil {
			//var sd,ed time.Time
			if filter.Amount.From != nil {
				sd := time.Date(filter.Amount.From.Year(), filter.Amount.From.Month(), filter.Amount.From.Day(), 0, 0, 0, 0, filter.Amount.From.Location())
				ed := time.Date(filter.Amount.From.Year(), filter.Amount.From.Month(), filter.Amount.From.Day(), 23, 59, 59, 0, filter.Amount.From.Location())
				if filter.Amount.To != nil {
					ed = time.Date(filter.Amount.To.Year(), filter.Amount.To.Month(), filter.Amount.To.Day(), 23, 59, 59, 0, filter.Amount.To.Location())
				}
				query = append(query, bson.M{"amount": bson.M{"$gte": sd, "$lte": ed}})

			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("propertyWallet query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyWallets []models.RefPropertyWallet
	if err = cursor.All(context.TODO(), &propertyWallets); err != nil {
		return nil, err
	}
	return propertyWallets, nil
}
