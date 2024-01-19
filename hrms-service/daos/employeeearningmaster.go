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

//SaveEmployeeEarningMaster : ""
func (d *Daos) SaveEmployeeEarningMaster(ctx *models.Context, employeeEarningMaster *models.EmployeeEarningMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).InsertOne(ctx.CTX, employeeEarningMaster)
	if err != nil {
		return err
	}
	employeeEarningMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeEarningMaster : ""
func (d *Daos) UpdateEmployeeEarningMaster(ctx *models.Context, employeeEarningMaster *models.EmployeeEarningMaster) error {
	selector := bson.M{"uniqueId": employeeEarningMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeEarningMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeEarningMaster : ""
func (d *Daos) GetSingleEmployeeEarningMaster(ctx *models.Context, uniqueID string) (*models.RefEmployeeEarningMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeEarningMasters []models.RefEmployeeEarningMaster
	var EmployeeEarningMaster *models.RefEmployeeEarningMaster
	if err = cursor.All(ctx.CTX, &EmployeeEarningMasters); err != nil {
		return nil, err
	}
	if len(EmployeeEarningMasters) > 0 {
		EmployeeEarningMaster = &EmployeeEarningMasters[0]
	}
	return EmployeeEarningMaster, err
}

// EnableEmployeeEarningMaster : ""
func (d *Daos) EnableEmployeeEarningMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEARNINGMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeEarningMaster : ""
func (d *Daos) DisableEmployeeEarningMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEARNINGMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeEarningMaster :""
func (d *Daos) DeleteEmployeeEarningMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEARNINGMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeEarningMaster : ""
func (d *Daos) FilterEmployeeEarningMaster(ctx *models.Context, employeeEarningMaster *models.FilterEmployeeEarningMaster, pagination *models.Pagination) ([]models.RefEmployeeEarningMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeEarningMaster != nil {
		if len(employeeEarningMaster.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeEarningMaster.Status}})
		}
		if len(employeeEarningMaster.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeEarningMaster.OrganisationId}})
		}
		//Regex
		if employeeEarningMaster.Regex.Title != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: employeeEarningMaster.Regex.Title, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNINGMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeEarningMasterFilter []models.RefEmployeeEarningMaster
	if err = cursor.All(context.TODO(), &employeeEarningMasterFilter); err != nil {
		return nil, err
	}
	return employeeEarningMasterFilter, nil
}
