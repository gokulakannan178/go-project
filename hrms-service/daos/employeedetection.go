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

//SaveEmployeeDeduction : ""
func (d *Daos) SaveEmployeeDeduction(ctx *models.Context, employeeDeduction *models.EmployeeDeduction) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).InsertOne(ctx.CTX, employeeDeduction)
	if err != nil {
		return err
	}
	employeeDeduction.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeDeduction : ""
func (d *Daos) UpdateEmployeeDeduction(ctx *models.Context, employeeDeduction *models.EmployeeDeduction) error {
	selector := bson.M{"uniqueId": employeeDeduction.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeDeduction}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeDeduction : ""
func (d *Daos) GetSingleEmployeeDeduction(ctx *models.Context, uniqueID string) (*models.RefEmployeeDeduction, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER, "deductionId", "uniqueId", "ref.deductionId", "ref.deductionId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeDeductions []models.RefEmployeeDeduction
	var EmployeeDeduction *models.RefEmployeeDeduction
	if err = cursor.All(ctx.CTX, &EmployeeDeductions); err != nil {
		return nil, err
	}
	if len(EmployeeDeductions) > 0 {
		EmployeeDeduction = &EmployeeDeductions[0]
	}
	return EmployeeDeduction, err
}

// EnableEmployeeDeduction : ""
func (d *Daos) EnableEmployeeDeduction(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDEDUCTIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeDeduction : ""
func (d *Daos) DisableEmployeeDeduction(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDEDUCTIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeDeduction :""
func (d *Daos) DeleteEmployeeDeduction(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDEDUCTIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeDeduction : ""
func (d *Daos) FilterEmployeeDeduction(ctx *models.Context, employeeDeduction *models.FilterEmployeeDeduction, pagination *models.Pagination) ([]models.RefEmployeeDeduction, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeDeduction != nil {
		if len(employeeDeduction.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeDeduction.Status}})
		}
		if len(employeeDeduction.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeDeduction.OrganisationId}})
		}
		if len(employeeDeduction.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeDeduction.EmployeeId}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeeDeduction != nil {
		if employeeDeduction.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeDeduction.SortBy: employeeDeduction.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER, "deductionId", "uniqueId", "ref.deductionId", "ref.deductionId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDeductionFilter []models.RefEmployeeDeduction
	if err = cursor.All(context.TODO(), &employeeDeductionFilter); err != nil {
		return nil, err
	}
	return employeeDeductionFilter, nil
}
func (d *Daos) GetSingleEmployeeDeductionWithEmployee(ctx *models.Context, employee string, deduction string) (*models.RefEmployeeDeduction, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employee, "deductionId": deduction, "status": "Active"}})
	//LookUp
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("GetSingleEmployeeDeductionWithEmployee query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeDeductions []models.RefEmployeeDeduction
	var EmployeeDeduction *models.RefEmployeeDeduction
	if err = cursor.All(ctx.CTX, &EmployeeDeductions); err != nil {
		return nil, err
	}
	if len(EmployeeDeductions) > 0 {
		EmployeeDeduction = &EmployeeDeductions[0]
	}
	return EmployeeDeduction, err
}
func (d *Daos) ArchivedEmployeeDeduction(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update2 := models.Updated{}
	update2.On = &t
	update2.By = constants.SYSTEM
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYROLLSTATUSARCHIVED, "updated": update2, "endDate": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDEDUCTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
