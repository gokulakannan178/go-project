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

//SaveEmployeePayslip :""
func (d *Daos) SaveEmployeePayslip(ctx *models.Context, employeePayslip *models.EmployeePayslip) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).InsertOne(ctx.CTX, employeePayslip)
	if err != nil {
		return err
	}
	employeePayslip.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleEmployeePayslip : ""
func (d *Daos) GetSingleEmployeePayslip(ctx *models.Context, uniqueID string) (*models.RefEmployeePayslip, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeePayslipCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeePayslips []models.RefEmployeePayslip
	var EmployeePayslip *models.RefEmployeePayslip
	if err = cursor.All(ctx.CTX, &EmployeePayslips); err != nil {
		return nil, err
	}
	if len(EmployeePayslips) > 0 {
		EmployeePayslip = &EmployeePayslips[0]
	}
	return EmployeePayslip, nil
}

//UpdateEmployeePayslip : ""
func (d *Daos) UpdateEmployeePayslip(ctx *models.Context, employeePayslip *models.EmployeePayslip) error {
	selector := bson.M{"uniqueId": employeePayslip.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeePayslip}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableEmployeePayslip :""
func (d *Daos) EnableEmployeePayslip(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYSLIPSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEmployeePayslip :""
func (d *Daos) DisableEmployeePayslip(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYSLIPSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEmployeePayslip :""
func (d *Daos) DeleteEmployeePayslip(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYSLIPSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterEmployeePayslip : ""
func (d *Daos) FilterEmployeePayslip(ctx *models.Context, filter *models.FilterEmployeePayslip, pagination *models.Pagination) ([]models.RefEmployeePayslip, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": filter.EmployeeId}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	// if filter != nil {
	// 	if EmployeePayslipfilter.SortField != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{EmployeePayslipfilter.SortField: EmployeePayslipfilter.SortOrder}})
	// 	}
	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeePayslipCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeePayslip query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEPAYSLIP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeePayslips []models.RefEmployeePayslip
	if err = cursor.All(context.TODO(), &EmployeePayslips); err != nil {
		return nil, err
	}
	return EmployeePayslips, nil
}
