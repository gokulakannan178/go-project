package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveBillclaimConfig :""
func (d *Daos) SaveBillclaimConfig(ctx *models.Context, billclaimconfig *models.BillclaimConfig) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).InsertOne(ctx.CTX, billclaimconfig)
	if err != nil {
		return err
	}
	billclaimconfig.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleBillclaimConfig : ""
func (d *Daos) GetSingleBillclaimConfig(ctx *models.Context, uniqueID string) (*models.RefBillclaimConfig, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONBILLCLAIMCONFIG, "uniqueId", "BillclaimConfigID", "ref.billclaimconfigAssetsId", "ref.billclaimconfigAssetsId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billclaimconfigs []models.RefBillclaimConfig
	var billclaimconfig *models.RefBillclaimConfig
	if err = cursor.All(ctx.CTX, &billclaimconfigs); err != nil {
		return nil, err
	}
	if len(billclaimconfigs) > 0 {
		billclaimconfig = &billclaimconfigs[0]
	}
	return billclaimconfig, nil
}

//UpdateBillclaimConfig : ""
func (d *Daos) UpdateBillclaimConfig(ctx *models.Context, billclaimconfig *models.BillclaimConfig) error {
	selector := bson.M{"uniqueId": billclaimconfig.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": billclaimconfig}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableBillclaimConfig :""
func (d *Daos) EnableBillclaimConfig(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMCONFIGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBillclaimConfig :""
func (d *Daos) DisableBillclaimConfig(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMCONFIGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBillclaimConfig :""
func (d *Daos) DeleteBillclaimConfig(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMCONFIGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterBillclaimConfig : ""
func (d *Daos) FilterBillclaimConfig(ctx *models.Context, billclaimconfigFilter *models.BillclaimConfigFilter, pagination *models.Pagination) ([]models.RefBillclaimConfig, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if billclaimconfigFilter != nil {

		if len(billclaimconfigFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": billclaimconfigFilter.Status}})
		}
		if len(billclaimconfigFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": billclaimconfigFilter.OrganisationID}})
		}
		//Regex
		if billclaimconfigFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: billclaimconfigFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if billclaimconfigFilter != nil {
		if billclaimconfigFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{billclaimconfigFilter.SortBy: billclaimconfigFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billclaimconfigsFilter []models.RefBillclaimConfig
	if err = cursor.All(context.TODO(), &billclaimconfigFilter); err != nil {
		return nil, err
	}
	return billclaimconfigsFilter, nil
}
func (d *Daos) GetSingleBillclaimApprovalLevels(ctx *models.Context, Grade string, Amount float64) (*models.RefBillclaimConfig, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"grade": Grade})
	query = append(query, bson.M{"minAmount": bson.M{"$lte": Amount}})
	query = append(query, bson.M{"maxAmount": bson.M{"$gte": Amount}})

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("GetSingleBillclaimApprovalLevels query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billclaimconfigs []models.RefBillclaimConfig
	var billclaimconfig *models.RefBillclaimConfig
	if err = cursor.All(ctx.CTX, &billclaimconfigs); err != nil {
		return nil, err
	}
	if len(billclaimconfigs) > 0 {
		billclaimconfig = &billclaimconfigs[0]
	}
	return billclaimconfig, nil
}

//GetSingleBillclaimConfigWithLevel : ""
func (d *Daos) GetSingleBillclaimConfigWithLevel(ctx *models.Context, grade string, level int64) (*models.RefBillclaimConfig, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"grade": grade, "level": level, "status": constants.BILLCLAIMCONFIGSTATUSACTIVE}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billclaimconfigs []models.RefBillclaimConfig
	var billclaimconfig *models.RefBillclaimConfig
	if err = cursor.All(ctx.CTX, &billclaimconfigs); err != nil {
		return nil, err
	}
	if len(billclaimconfigs) > 0 {
		billclaimconfig = &billclaimconfigs[0]
	}
	return billclaimconfig, nil
}
