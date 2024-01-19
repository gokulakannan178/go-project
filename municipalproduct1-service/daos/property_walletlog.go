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

//SavePropertyWalletLog :""
func (d *Daos) SavePropertyWalletLog(ctx *models.Context, propertyWallet *models.PropertyWalletLog) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).InsertOne(ctx.CTX, propertyWallet)
	return err
}

//GetSinglePropertyWalletLog : ""
func (d *Daos) GetSinglePropertyWalletLog(ctx *models.Context, UniqueID string) (*models.RefPropertyWalletLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyWallets []models.RefPropertyWalletLog
	var propertyWallet *models.RefPropertyWalletLog
	if err = cursor.All(ctx.CTX, &propertyWallets); err != nil {
		return nil, err
	}
	if len(propertyWallets) > 0 {
		propertyWallet = &propertyWallets[0]
	}
	return propertyWallet, nil
}

//UpdatePropertyWalletLog : ""
func (d *Daos) UpdatePropertyWalletLog(ctx *models.Context, propertyWallet *models.PropertyWalletLog) error {
	selector := bson.M{"uniqueId": propertyWallet.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyWallet, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnablePropertyWalletLog :""
func (d *Daos) EnablePropertyWalletLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYWALLETSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyWalletLog :""
func (d *Daos) DisablePropertyWalletLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYWALLETSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyWalletLog :""
func (d *Daos) DeletePropertyWalletLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYWALLETSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterPropertyWalletLog : ""
func (d *Daos) FilterPropertyWalletLog(ctx *models.Context, filter *models.PropertyWalletLogFilter, pagination *models.Pagination) ([]models.RefPropertyWalletLog, error) {
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
		if len(filter.Scenario) > 0 {
			query = append(query, bson.M{"scenario": bson.M{"$in": filter.Scenario}})
		}
		if len(filter.CreatedBy) > 0 {
			query = append(query, bson.M{"created.by": bson.M{"$in": filter.CreatedBy}})
		}

		if filter.SearchText.MobileNo != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.SearchText.MobileNo, Options: "xi"}})
		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		}
		if filter.SearchText.WalletID != "" {
			query = append(query, bson.M{"walletId": primitive.Regex{Pattern: filter.SearchText.WalletID, Options: "xi"}})
		}
		if filter.PreTnxAmountDateRange != nil {
			//var sd,ed time.Time
			if filter.PreTnxAmountDateRange.From != nil {
				sd := time.Date(filter.PreTnxAmountDateRange.From.Year(), filter.PreTnxAmountDateRange.From.Month(), filter.PreTnxAmountDateRange.From.Day(), 0, 0, 0, 0, filter.PreTnxAmountDateRange.From.Location())
				ed := time.Date(filter.PreTnxAmountDateRange.From.Year(), filter.PreTnxAmountDateRange.From.Month(), filter.PreTnxAmountDateRange.From.Day(), 23, 59, 59, 0, filter.PreTnxAmountDateRange.From.Location())
				if filter.PreTnxAmountDateRange.To != nil {
					ed = time.Date(filter.PreTnxAmountDateRange.To.Year(), filter.PreTnxAmountDateRange.To.Month(), filter.PreTnxAmountDateRange.To.Day(), 23, 59, 59, 0, filter.PreTnxAmountDateRange.To.Location())
				}
				query = append(query, bson.M{"preTnxAmount": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.PostTnxAmountDateRange != nil {
			//var sd,ed time.Time
			if filter.PostTnxAmountDateRange.From != nil {
				sd := time.Date(filter.PostTnxAmountDateRange.From.Year(), filter.PostTnxAmountDateRange.From.Month(), filter.PostTnxAmountDateRange.From.Day(), 0, 0, 0, 0, filter.PostTnxAmountDateRange.From.Location())
				ed := time.Date(filter.PostTnxAmountDateRange.From.Year(), filter.PostTnxAmountDateRange.From.Month(), filter.PostTnxAmountDateRange.From.Day(), 23, 59, 59, 0, filter.PostTnxAmountDateRange.From.Location())
				if filter.PostTnxAmountDateRange.To != nil {
					ed = time.Date(filter.PostTnxAmountDateRange.To.Year(), filter.PostTnxAmountDateRange.To.Month(), filter.PostTnxAmountDateRange.To.Day(), 23, 59, 59, 0, filter.PostTnxAmountDateRange.To.Location())
				}
				query = append(query, bson.M{"postTnxAmount": bson.M{"$gte": sd, "$lte": ed}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYWALLETLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyWallets []models.RefPropertyWalletLog
	if err = cursor.All(context.TODO(), &propertyWallets); err != nil {
		return nil, err
	}
	return propertyWallets, nil
}
