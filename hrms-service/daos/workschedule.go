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

//SaveWorkSchedule :""
func (d *Daos) SaveWorkSchedule(ctx *models.Context, workSchedule *models.WorkSchedule) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).InsertOne(ctx.CTX, workSchedule)
	if err != nil {
		return err
	}
	workSchedule.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleWorkSchedule : ""
func (d *Daos) GetSingleWorkSchedule(ctx *models.Context, UniqueID string) (*models.RefWorkSchedule, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("WorkSchedule query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var workSchedules []models.RefWorkSchedule
	var workSchedule *models.RefWorkSchedule
	if err = cursor.All(ctx.CTX, &workSchedules); err != nil {
		return nil, err
	}
	if len(workSchedules) > 0 {
		workSchedule = &workSchedules[0]
	}
	return workSchedule, nil
}

//UpdateWorkSchedule : ""
func (d *Daos) UpdateWorkSchedule(ctx *models.Context, workSchedule *models.WorkSchedule) error {
	selector := bson.M{"uniqueId": workSchedule.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": workSchedule}
	_, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableWorkSchedule :""
func (d *Daos) EnableWorkSchedule(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WORKSCHEDULESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableWorkSchedule :""
func (d *Daos) DisableWorkSchedule(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WORKSCHEDULESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWorkSchedule :""
func (d *Daos) DeleteWorkSchedule(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WORKSCHEDULESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterWorkSchedule : ""
func (d *Daos) FilterWorkSchedule(ctx *models.Context, filter *models.WorkScheduleFilter, pagination *models.Pagination) ([]models.RefWorkSchedule, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationID}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("WorkSchedule query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var workSchedules []models.RefWorkSchedule
	if err = cursor.All(context.TODO(), &workSchedules); err != nil {
		return nil, err
	}
	return workSchedules, nil
}

//GetSingleWorkScheduleWithActive : ""
func (d *Daos) GetSingleWorkScheduleActiveWithName(ctx *models.Context, UniqueID string) (*models.RefWorkSchedule, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": UniqueID, "status": constants.ONBOARDINGPOLICYSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("WorkSchedule query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWORKSCHEDULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var workSchedules []models.RefWorkSchedule
	var workSchedule *models.RefWorkSchedule
	if err = cursor.All(ctx.CTX, &workSchedules); err != nil {
		return nil, err
	}
	if len(workSchedules) > 0 {
		workSchedule = &workSchedules[0]
	}
	return workSchedule, nil
}
