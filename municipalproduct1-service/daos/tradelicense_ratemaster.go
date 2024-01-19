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
)

// SaveTradeLicenseRateMaster : ""
func (d *Daos) SaveTradeLicenseRateMaster(ctx *models.Context, block *models.TradeLicenseRateMaster) error {
	d.Shared.BsonToJSONPrint(block)
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).InsertOne(ctx.CTX, block)
	return err
}

// GetSingleTradeLicenseRateMaster : ""
func (d *Daos) GetSingleTradeLicenseRateMaster(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseRateMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	// LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "ref.tradeLicensecategoryType", "ref.tradeLicenseCategoryType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefTradeLicenseRateMaster
	var tower *models.RefTradeLicenseRateMaster
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateTradeLicenseRateMaster : ""
func (d *Daos) UpdateTradeLicenseRateMaster(ctx *models.Context, rateMaster *models.TradeLicenseRateMaster) error {
	selector := bson.M{"uniqueId": rateMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": rateMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableTradeLicenseRateMaster : ""
func (d *Daos) EnableTradeLicenseRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSERATEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableTradeLicenseRateMaster : ""
func (d *Daos) DisableTradeLicenseRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSERATEMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteTradeLicenseRateMaster : ""
func (d *Daos) DeleteTradeLicenseRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSERATEMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTradeLicenseRateMaster : ""
func (d *Daos) FilterTradeLicenseRateMaster(ctx *models.Context, filter *models.TradeLicenseRateMasterFilter, pagination *models.Pagination) ([]models.RefTradeLicenseRateMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.TLCTID) > 0 {
			query = append(query, bson.M{"tlctId": bson.M{"$in": filter.TLCTID}})
		}
		if len(filter.TLBTID) > 0 {
			query = append(query, bson.M{"tlbtId": bson.M{"$in": filter.TLBTID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	// LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "ref.tradeLicensecategoryType", "ref.tradeLicenseCategoryType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSERATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rateMaster []models.RefTradeLicenseRateMaster
	if err = cursor.All(context.TODO(), &rateMaster); err != nil {
		return nil, err
	}
	return rateMaster, nil
}
