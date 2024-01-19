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

//SaveEmployeeDeductionMaster : ""
func (d *Daos) SaveEmployeeDeductionMaster(ctx *models.Context, employeeDeductionMaster *models.EmployeeDeductionMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).InsertOne(ctx.CTX, employeeDeductionMaster)
	if err != nil {
		return err
	}
	employeeDeductionMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeDeductionMaster : ""
func (d *Daos) UpdateEmployeeDeductionMaster(ctx *models.Context, employeeDeductionMaster *models.EmployeeDeductionMaster) error {
	selector := bson.M{"uniqueId": employeeDeductionMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeDeductionMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeDeductionMaster : ""
func (d *Daos) GetSingleEmployeeDeductionMaster(ctx *models.Context, uniqueID string) (*models.RefEmployeeDeductionMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeDeductionMasters []models.RefEmployeeDeductionMaster
	var EmployeeDeductionMaster *models.RefEmployeeDeductionMaster
	if err = cursor.All(ctx.CTX, &EmployeeDeductionMasters); err != nil {
		return nil, err
	}
	if len(EmployeeDeductionMasters) > 0 {
		EmployeeDeductionMaster = &EmployeeDeductionMasters[0]
	}
	return EmployeeDeductionMaster, err
}

// EnableEmployeeDeductionMaster : ""
func (d *Daos) EnableEmployeeDeductionMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDEDUCTIONMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeDeductionMaster : ""
func (d *Daos) DisableEmployeeDeductionMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDEDUCTIONMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeDeductionMaster :""
func (d *Daos) DeleteEmployeeDeductionMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDEDUCTIONMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeDeductionMaster : ""
func (d *Daos) FilterEmployeeDeductionMaster(ctx *models.Context, employeeDeductionMaster *models.FilterEmployeeDeductionMaster, pagination *models.Pagination) ([]models.RefEmployeeDeductionMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeDeductionMaster != nil {
		if len(employeeDeductionMaster.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeDeductionMaster.Status}})
		}
		if len(employeeDeductionMaster.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeDeductionMaster.OrganisationId}})
		}
		//Regex
		if employeeDeductionMaster.Regex.Title != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: employeeDeductionMaster.Regex.Title, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDeductionMasterFilter []models.RefEmployeeDeductionMaster
	if err = cursor.All(context.TODO(), &employeeDeductionMasterFilter); err != nil {
		return nil, err
	}
	return employeeDeductionMasterFilter, nil
}
