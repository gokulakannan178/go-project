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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePayroll :""
func (d *Daos) SavePayroll(ctx *models.Context, payroll *models.Payroll) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).InsertOne(ctx.CTX, payroll)
	if err != nil {
		return err
	}
	payroll.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSinglePayroll : ""
func (d *Daos) GetSinglePayroll(ctx *models.Context, uniqueID string) (*models.RefPayroll, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrolls []models.RefPayroll
	var payroll *models.RefPayroll
	if err = cursor.All(ctx.CTX, &payrolls); err != nil {
		return nil, err
	}
	if len(payrolls) > 0 {
		payroll = &payrolls[0]
	}
	return payroll, nil
}
func (d *Daos) SavePayrollWithUpsert(ctx *models.Context, payroll *models.Payroll) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"employeeId": payroll.EmployeeId}
	updateData := bson.M{"$set": payroll}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

//UpdatePayroll : ""
func (d *Daos) UpdatePayroll(ctx *models.Context, payroll *models.Payroll) error {
	selector := bson.M{"uniqueId": payroll.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payroll}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnablePayroll :""
func (d *Daos) EnablePayroll(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePayroll :""
func (d *Daos) DisablePayroll(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePayroll :""
func (d *Daos) DeletePayroll(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterPayroll : ""
func (d *Daos) FilterPayroll(ctx *models.Context, filter *models.FilterPayroll, pagination *models.Pagination) ([]models.RefPayroll, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
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

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employeeId", "ref.employeeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "ref.employeeId.designationId", "uniqueId", "ref.designationId", "ref.designationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Payroll query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrolls []models.RefPayroll
	if err = cursor.All(context.TODO(), &payrolls); err != nil {
		return nil, err
	}
	return payrolls, nil
}
func (d *Daos) GetSinglePayrollWithDays(ctx *models.Context, uniqueID int64) (*models.RefPayroll, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"payroll": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employeeId", "ref.employeeId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrolls []models.RefPayroll
	var payroll *models.RefPayroll
	if err = cursor.All(ctx.CTX, &payrolls); err != nil {
		return nil, err
	}
	if len(payrolls) > 0 {
		payroll = &payrolls[0]
	}
	return payroll, nil
}
func (d *Daos) GetSinglePayrollWithEmployee(ctx *models.Context, uniqueID string) (*models.RefPayroll, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employeeId", "ref.employeeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBANKINFORMATION, "employeeId", "employeeId", "ref.bank", "ref.bank")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "ref.employeeId.designationId", "uniqueId", "ref.designationId", "ref.designationId")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	d.Shared.BsonToJSONPrintTag("Payroll query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrolls []models.RefPayroll
	var payroll *models.RefPayroll
	if err = cursor.All(ctx.CTX, &payrolls); err != nil {
		return nil, err
	}
	if len(payrolls) > 0 {
		payroll = &payrolls[0]
	}
	return payroll, nil
}
