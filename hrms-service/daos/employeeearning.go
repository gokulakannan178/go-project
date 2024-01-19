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

//SaveEmployeeEarning : ""
func (d *Daos) SaveEmployeeEarning(ctx *models.Context, employeeEarning *models.EmployeeEarning) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).InsertOne(ctx.CTX, employeeEarning)
	if err != nil {
		return err
	}
	employeeEarning.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeEarning : ""
func (d *Daos) UpdateEmployeeEarning(ctx *models.Context, employeeEarning *models.EmployeeEarning) error {
	selector := bson.M{"uniqueId": employeeEarning.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeEarning}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeEarning : ""
func (d *Daos) GetSingleEmployeeEarning(ctx *models.Context, uniqueID string) (*models.RefEmployeeEarning, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEEEARNINGMASTER, "earningId", "uniqueId", "ref.earningId", "ref.earningId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeEarnings []models.RefEmployeeEarning
	var EmployeeEarning *models.RefEmployeeEarning
	if err = cursor.All(ctx.CTX, &EmployeeEarnings); err != nil {
		return nil, err
	}
	if len(EmployeeEarnings) > 0 {
		EmployeeEarning = &EmployeeEarnings[0]
	}
	return EmployeeEarning, err
}

// EnableEmployeeEarning : ""
func (d *Daos) EnableEmployeeEarning(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEARNINGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeEarning : ""
func (d *Daos) DisableEmployeeEarning(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEARNINGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeEarning :""
func (d *Daos) DeleteEmployeeEarning(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEARNINGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeEarning : ""
func (d *Daos) FilterEmployeeEarning(ctx *models.Context, employeeEarning *models.FilterEmployeeEarning, pagination *models.Pagination) ([]models.RefEmployeeEarning, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeEarning != nil {
		if len(employeeEarning.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeEarning.Status}})
		}
		if len(employeeEarning.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeEarning.OrganisationId}})
		}
		if len(employeeEarning.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeEarning.EmployeeId}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeeEarning != nil {
		if employeeEarning.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeEarning.SortBy: employeeEarning.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEEEARNINGMASTER, "earningId", "uniqueId", "ref.earningId", "ref.earningId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeEarningFilter []models.RefEmployeeEarning
	if err = cursor.All(context.TODO(), &employeeEarningFilter); err != nil {
		return nil, err
	}
	return employeeEarningFilter, nil
}
func (d *Daos) GetSingleEmployeeEaringWithEmployee(ctx *models.Context, employee string, earning string) (*models.RefEmployeeEarning, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employee, "earningId": earning, "status": "Active"}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("GetSingleEmployeeEaringWithEmployee query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeePayrolls []models.RefEmployeeEarning
	var EmployeePayroll *models.RefEmployeeEarning
	if err = cursor.All(ctx.CTX, &EmployeePayrolls); err != nil {
		return nil, err
	}
	if len(EmployeePayrolls) > 0 {
		EmployeePayroll = &EmployeePayrolls[0]
	}
	return EmployeePayroll, err
}
func (d *Daos) ArchivedEmployeeEarning(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update2 := models.Updated{}
	update2.On = &t
	update2.By = constants.SYSTEM
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYROLLSTATUSARCHIVED, "updated": update2, "endDate": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEARNING).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
