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

//SaveSalaryConfig : ""
func (d *Daos) SaveSalaryConfig(ctx *models.Context, salaryConfig *models.SalaryConfig) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).InsertOne(ctx.CTX, salaryConfig)
	if err != nil {
		return err
	}
	salaryConfig.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateSalaryConfig : ""
func (d *Daos) UpdateSalaryConfig(ctx *models.Context, SalaryConfig *models.SalaryConfig) error {
	selector := bson.M{"uniqueId": SalaryConfig.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": SalaryConfig}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleSalaryConfig : ""
func (d *Daos) GetSingleSalaryConfig(ctx *models.Context, uniqueID string) (*models.RefSalaryConfig, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SalaryConfigs []models.RefSalaryConfig
	var SalaryConfig *models.RefSalaryConfig
	if err = cursor.All(ctx.CTX, &SalaryConfigs); err != nil {
		return nil, err
	}
	if len(SalaryConfigs) > 0 {
		SalaryConfig = &SalaryConfigs[0]
	}
	return SalaryConfig, err
}

// EnableSalaryConfig : ""
func (d *Daos) EnableSalaryConfig(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SALARYCONFIGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableSalaryConfig : ""
func (d *Daos) DisableSalaryConfig(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SALARYCONFIGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteSalaryConfig :""
func (d *Daos) DeleteSalaryConfig(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SALARYCONFIGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSalaryConfig : ""
func (d *Daos) FilterSalaryConfig(ctx *models.Context, salaryConfig *models.FilterSalaryConfig, pagination *models.Pagination) ([]models.RefSalaryConfig, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if salaryConfig != nil {
		if len(salaryConfig.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": salaryConfig.Status}})
		}
		if len(salaryConfig.EmployeeID) > 0 {
			query = append(query, bson.M{"employeeID": bson.M{"$in": salaryConfig.EmployeeID}})
		}
		if len(salaryConfig.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": salaryConfig.OrganisationID}})
		}
		//Regex

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if salaryConfig != nil {
		if salaryConfig.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{salaryConfig.SortBy: salaryConfig.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).CountDocuments(ctx.CTX, func() bson.M {
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
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var salaryConfigFilter []models.RefSalaryConfig
	if err = cursor.All(context.TODO(), &salaryConfigFilter); err != nil {
		return nil, err
	}
	return salaryConfigFilter, nil
}
func (d *Daos) GetSingleSalaryConfigWithEmployeeType(ctx *models.Context, uniqueID string) (*models.RefSalaryConfig, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeType": uniqueID, "status": constants.SALARYCONFIGSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SalaryConfigs []models.RefSalaryConfig
	var SalaryConfig *models.RefSalaryConfig
	if err = cursor.All(ctx.CTX, &SalaryConfigs); err != nil {
		return nil, err
	}
	if len(SalaryConfigs) > 0 {
		SalaryConfig = &SalaryConfigs[0]
	}
	return SalaryConfig, err
}
