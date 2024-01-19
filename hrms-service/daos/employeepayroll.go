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

//SaveEmployeePayroll : ""
func (d *Daos) SaveEmployeePayroll(ctx *models.Context, employeePayroll *models.EmployeePayroll) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).InsertOne(ctx.CTX, employeePayroll)
	if err != nil {
		return err
	}
	employeePayroll.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeePayroll : ""
func (d *Daos) UpdateEmployeePayroll(ctx *models.Context, employeePayroll *models.EmployeePayroll) error {
	selector := bson.M{"uniqueId": employeePayroll.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeePayroll}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeePayroll : ""
func (d *Daos) GetSingleEmployeePayroll(ctx *models.Context, uniqueID string) (*models.RefEmployeePayroll, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeePayrolls []models.RefEmployeePayroll
	var EmployeePayroll *models.RefEmployeePayroll
	if err = cursor.All(ctx.CTX, &EmployeePayrolls); err != nil {
		return nil, err
	}
	if len(EmployeePayrolls) > 0 {
		EmployeePayroll = &EmployeePayrolls[0]
	}
	return EmployeePayroll, err
}

// EnableEmployeePayroll : ""
func (d *Daos) EnableEmployeePayroll(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYROLLSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeePayroll : ""
func (d *Daos) DisableEmployeePayroll(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYROLLSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeePayroll :""
func (d *Daos) DeleteEmployeePayroll(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYROLLSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeePayroll : ""
func (d *Daos) FilterEmployeePayroll(ctx *models.Context, employeePayroll *models.FilterEmployeePayroll, pagination *models.Pagination) ([]models.RefEmployeePayroll, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeePayroll != nil {
		if len(employeePayroll.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeePayroll.Status}})
		}
		if len(employeePayroll.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeePayroll.OrganisationId}})
		}
		if len(employeePayroll.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeePayroll.EmployeeId}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeePayroll != nil {
		if employeePayroll.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeePayroll.SortBy: employeePayroll.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeePayrollFilter []models.RefEmployeePayroll
	if err = cursor.All(context.TODO(), &employeePayrollFilter); err != nil {
		return nil, err
	}
	return employeePayrollFilter, nil
}
func (d *Daos) GetSingleEmployeePayrollWithEmployee(ctx *models.Context, employee string) (*models.RefEmployeePayroll, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employee, "status": "Active"}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeePayrolls []models.RefEmployeePayroll
	var EmployeePayroll *models.RefEmployeePayroll
	if err = cursor.All(ctx.CTX, &EmployeePayrolls); err != nil {
		return nil, err
	}
	if len(EmployeePayrolls) > 0 {
		EmployeePayroll = &EmployeePayrolls[0]
	}
	return EmployeePayroll, err
}
func (d *Daos) ArchivedEmployeePayroll(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update2 := models.Updated{}
	update2.On = &t
	update2.By = constants.SYSTEM
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYROLLSTATUSARCHIVED, "updated": update2, "endDate": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYROLL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
