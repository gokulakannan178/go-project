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

//SaveEmployeeHistory :""
func (d *Daos) SaveEmployeeHistory(ctx *models.Context, employeeHistory *models.EmployeeHistory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).InsertOne(ctx.CTX, employeeHistory)
	if err != nil {
		return err
	}
	employeeHistory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleEmployeeHistory : ""
func (d *Daos) GetSingleEmployeeHistory(ctx *models.Context, uniqueID string) (*models.RefEmployeeHistory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeHistoryCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeHistorys []models.RefEmployeeHistory
	var employeeHistory *models.RefEmployeeHistory
	if err = cursor.All(ctx.CTX, &employeeHistorys); err != nil {
		return nil, err
	}
	if len(employeeHistorys) > 0 {
		employeeHistory = &employeeHistorys[0]
	}
	return employeeHistory, nil
}

//UpdateEmployeeHistory : ""
func (d *Daos) UpdateEmployeeHistory(ctx *models.Context, employeeHistory *models.EmployeeHistory) error {
	selector := bson.M{"uniqueId": employeeHistory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeHistory}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableEmployeeHistory :""
func (d *Daos) EnableEmployeeHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEHISTORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEmployeeHistory :""
func (d *Daos) DisableEmployeeHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEHISTORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEmployeeHistory :""
func (d *Daos) DeleteEmployeeHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEHISTORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterEmployeeHistory : ""
func (d *Daos) FilterEmployeeHistory(ctx *models.Context, filter *models.FilterEmployeeHistory, pagination *models.Pagination) ([]models.RefEmployeeHistory, error) {
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

	// if filter != nil {
	// 	if EmployeeHistoryfilter.SortField != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{EmployeeHistoryfilter.SortField: EmployeeHistoryfilter.SortOrder}})
	// 	}
	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeHistoryCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeHistory query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeHistorys []models.RefEmployeeHistory
	if err = cursor.All(context.TODO(), &EmployeeHistorys); err != nil {
		return nil, err
	}
	return EmployeeHistorys, nil
}
