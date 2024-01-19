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

//SaveEmployeeJob :""
func (d *Daos) SaveEmployeeJob(ctx *models.Context, employeeJob *models.EmployeeJob) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).InsertOne(ctx.CTX, employeeJob)
	if err != nil {
		return err
	}
	employeeJob.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleEmployeeJob : ""
func (d *Daos) GetSingleEmployeeJob(ctx *models.Context, uniqueID string) (*models.RefEmployeeJob, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeJobCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeJobs []models.RefEmployeeJob
	var EmployeeJob *models.RefEmployeeJob
	if err = cursor.All(ctx.CTX, &EmployeeJobs); err != nil {
		return nil, err
	}
	if len(EmployeeJobs) > 0 {
		EmployeeJob = &EmployeeJobs[0]
	}
	return EmployeeJob, nil
}

//UpdateEmployeeJob : ""
func (d *Daos) UpdateEmployeeJob(ctx *models.Context, employeeJob *models.EmployeeJob) error {
	selector := bson.M{"uniqueId": employeeJob.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeJob}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableEmployeeJob :""
func (d *Daos) EnableEmployeeJob(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEJOBSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEmployeeJob :""
func (d *Daos) DisableEmployeeJob(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEJOBSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEmployeeJob :""
func (d *Daos) DeleteEmployeeJob(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEJOBSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterEmployeeJob : ""
func (d *Daos) FilterEmployeeJob(ctx *models.Context, filter *models.FilterEmployeeJob, pagination *models.Pagination) ([]models.RefEmployeeJob, error) {
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
	// 	if EmployeeJobfilter.SortField != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{EmployeeJobfilter.SortField: EmployeeJobfilter.SortOrder}})
	// 	}
	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeJobCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeJob query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEJOB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeJobs []models.RefEmployeeJob
	if err = cursor.All(context.TODO(), &EmployeeJobs); err != nil {
		return nil, err
	}
	return EmployeeJobs, nil
}
