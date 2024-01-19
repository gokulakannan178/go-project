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

//SaveNonResidentialUsageFactor :""
func (d *Daos) SaveNonResidentialUsageFactor(ctx *models.Context, nonResidentialUsageFactor *models.NonResidentialUsageFactor) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).InsertOne(ctx.CTX, nonResidentialUsageFactor)
	return err
}

//GetSingleNonResidentialUsageFactor : ""
func (d *Daos) GetSingleNonResidentialUsageFactor(ctx *models.Context, UniqueID string) (*models.RefNonResidentialUsageFactor, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var nonResidentialUsageFactors []models.RefNonResidentialUsageFactor
	var nonResidentialUsageFactor *models.RefNonResidentialUsageFactor
	if err = cursor.All(ctx.CTX, &nonResidentialUsageFactors); err != nil {
		return nil, err
	}
	if len(nonResidentialUsageFactors) > 0 {
		nonResidentialUsageFactor = &nonResidentialUsageFactors[0]
	}
	return nonResidentialUsageFactor, nil
}

//UpdateNonResidentialUsageFactor : ""
func (d *Daos) UpdateNonResidentialUsageFactor(ctx *models.Context, nonResidentialUsageFactor *models.NonResidentialUsageFactor) error {
	selector := bson.M{"uniqueId": nonResidentialUsageFactor.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": nonResidentialUsageFactor, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterNonResidentialUsageFactor : ""
func (d *Daos) FilterNonResidentialUsageFactor(ctx *models.Context, nonResidentialUsageFactorfilter *models.NonResidentialUsageFactorFilter, pagination *models.Pagination) ([]models.RefNonResidentialUsageFactor, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if nonResidentialUsageFactorfilter != nil {

		if len(nonResidentialUsageFactorfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": nonResidentialUsageFactorfilter.Status}})
		}
		if len(nonResidentialUsageFactorfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": nonResidentialUsageFactorfilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("nonResidentialUsageFactor query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var nonResidentialUsageFactors []models.RefNonResidentialUsageFactor
	if err = cursor.All(context.TODO(), &nonResidentialUsageFactors); err != nil {
		return nil, err
	}
	return nonResidentialUsageFactors, nil
}

//EnableNonResidentialUsageFactor :""
func (d *Daos) EnableNonResidentialUsageFactor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NONRESIDENTIALUSAGEFACTORSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableNonResidentialUsageFactor :""
func (d *Daos) DisableNonResidentialUsageFactor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NONRESIDENTIALUSAGEFACTORSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteNonResidentialUsageFactor :""
func (d *Daos) DeleteNonResidentialUsageFactor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NONRESIDENTIALUSAGEFACTORSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
