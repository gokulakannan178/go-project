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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveStoredCalc :""
func (d *Daos) SaveStoredCalc(ctx *models.Context, storedcalc *models.StoredCalc) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).InsertOne(ctx.CTX, storedcalc)
	return err
}

//SaveStoredCalcDemand :""
func (d *Daos) SaveStoredCalcDemandWithUpsert(ctx *models.Context, storedcalc *models.StoredCalculationDemand) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"propertyId": storedcalc.PropertyID}
	updateData := bson.M{"$set": storedcalc}
	if _, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALCDEMAND).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in upserting - " + err.Error())
	}
	return nil
}

//SaveManyStoredCalDemandFy
func (d *Daos) SaveManyStoredCalDemandFyWithUpsert(ctx *models.Context, storedcalc []models.StoredCalculationDemandfy) error {
	for _, v := range storedcalc {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"propertyId": v.PropertyID, "fyId": v.FyId}
		updateData := bson.M{"$set": v}
		if _, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALCDEMANDFY).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//GetSingleStoredCalc : ""
func (d *Daos) GetSingleStoredCalc(ctx *models.Context, code string) (*models.RefStoredCalc, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"code": code}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var storedcalcs []models.RefStoredCalc
	var storedcalc *models.RefStoredCalc
	if err = cursor.All(ctx.CTX, &storedcalcs); err != nil {
		return nil, err
	}
	if len(storedcalcs) > 0 {
		storedcalc = &storedcalcs[0]
	}
	return storedcalc, nil
}

//UpdateStoredCalc : ""
func (d *Daos) UpdateStoredCalc(ctx *models.Context, storedcalc *models.StoredCalc) error {
	selector := bson.M{"uniqueId": storedcalc.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": storedcalc, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterStoredCalc : ""
func (d *Daos) FilterStoredCalc(ctx *models.Context, storedcalcfilter *models.StoredCalcFilter, pagination *models.Pagination) ([]models.RefStoredCalc, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if storedcalcfilter != nil {
		if len(storedcalcfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": storedcalcfilter.UniqueID}})
		}
		if len(storedcalcfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": storedcalcfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("storedcalc query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var storedcalcs []models.RefStoredCalc
	if err = cursor.All(context.TODO(), &storedcalcs); err != nil {
		return nil, err
	}
	return storedcalcs, nil
}

//EnableStoredCalc :""
func (d *Daos) EnableStoredCalc(ctx *models.Context, code string) error {
	query := bson.M{"uniqueId": code}
	update := bson.M{"$set": bson.M{"status": constants.STOREDCALCSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableStoredCalc :""
func (d *Daos) DisableStoredCalc(ctx *models.Context, code string) error {
	query := bson.M{"uniqueId": code}
	update := bson.M{"$set": bson.M{"status": constants.STOREDCALCSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteStoredCalc :""
func (d *Daos) DeleteStoredCalc(ctx *models.Context, code string) error {
	query := bson.M{"uniqueId": code}
	update := bson.M{"$set": bson.M{"status": constants.STOREDCALCSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTOREDCALC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
