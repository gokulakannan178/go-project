package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveCitizenGraviansLog : ""
func (d *Daos) SaveCitizenGraviansLog(ctx *models.Context, citizengravianslog *models.CitizenGraviansLog) error {
	d.Shared.BsonToJSONPrint(citizengravianslog)
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).InsertOne(ctx.CTX, citizengravianslog)
	return err
}

// GetSingleCitizenGraviansLog : ""
func (d *Daos) GetSingleCitizenGraviansLog(ctx *models.Context, UniqueID string) (*models.RefCitizenGraviansLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var citizengravianslogs []models.RefCitizenGraviansLog
	var citizengravianslog *models.RefCitizenGraviansLog
	if err = cursor.All(ctx.CTX, &citizengravianslogs); err != nil {
		return nil, err
	}
	if len(citizengravianslogs) > 0 {
		citizengravianslog = &citizengravianslogs[0]
	}
	return citizengravianslog, nil
}

// UpdateCitizenGraviansLog : ""
func (d *Daos) UpdateCitizenGraviansLog(ctx *models.Context, citizengravianslog *models.CitizenGraviansLog) error {
	selector := bson.M{"uniqueId": citizengravianslog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": citizengravianslog}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableCitizenGraviansLog : ""
func (d *Daos) EnableCitizenGraviansLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableCitizenGraviansLog : ""
func (d *Daos) DisableCitizenGraviansLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteCitizenGraviansLog : ""
func (d *Daos) DeleteCitizenGraviansLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCitizenGraviansLog : ""
func (d *Daos) FilterCitizenGraviansLog(ctx *models.Context, filter *models.CitizenGraviansLogFilter, pagination *models.Pagination) ([]models.RefCitizenGraviansLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		// if len(filter.UniqueID) > 0 {
		// 	query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		// }
		// if len(filter.Description) > 0 {
		// 	query = append(query, bson.M{"description": bson.M{"$in": filter.Description}})
		// }
		// if len(filter.AssignedBy) > 0 {
		// 	query = append(query, bson.M{"assignedBy": bson.M{"$in": filter.AssignedBy}})
		// }
		// if len(filter.Description) > 0 {
		// 	query = append(query, bson.M{"description": bson.M{"$in": filter.Description}})
		// }
		if filter.FromDateRange != nil {
			//var sd,ed time.Time
			if filter.FromDateRange.From != nil {
				sd := time.Date(filter.FromDateRange.From.Year(), filter.FromDateRange.From.Month(), filter.FromDateRange.From.Day(), 0, 0, 0, 0, filter.FromDateRange.From.Location())
				ed := time.Date(filter.FromDateRange.From.Year(), filter.FromDateRange.From.Month(), filter.FromDateRange.From.Day(), 23, 59, 59, 0, filter.FromDateRange.From.Location())
				if filter.FromDateRange.To != nil {
					ed = time.Date(filter.FromDateRange.To.Year(), filter.FromDateRange.To.Month(), filter.FromDateRange.To.Day(), 23, 59, 59, 0, filter.FromDateRange.To.Location())
				}
				query = append(query, bson.M{"dateFrom": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.ToDateRange != nil {
			//var sd,ed time.Time
			if filter.ToDateRange.From != nil {
				sd := time.Date(filter.ToDateRange.From.Year(), filter.ToDateRange.From.Month(), filter.ToDateRange.From.Day(), 0, 0, 0, 0, filter.ToDateRange.From.Location())
				ed := time.Date(filter.ToDateRange.From.Year(), filter.ToDateRange.From.Month(), filter.ToDateRange.From.Day(), 23, 59, 59, 0, filter.ToDateRange.From.Location())
				if filter.ToDateRange.To != nil {
					ed = time.Date(filter.ToDateRange.To.Year(), filter.ToDateRange.To.Month(), filter.ToDateRange.To.Day(), 23, 59, 59, 0, filter.ToDateRange.To.Location())
				}
				query = append(query, bson.M{"dateTo": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.CreatedDateRange != nil {
			//var sd,ed time.Time
			if filter.CreatedDateRange.From != nil {
				sd := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 0, 0, 0, 0, filter.CreatedDateRange.From.Location())
				ed := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 23, 59, 59, 0, filter.CreatedDateRange.From.Location())
				if filter.CreatedDateRange.To != nil {
					ed = time.Date(filter.CreatedDateRange.To.Year(), filter.CreatedDateRange.To.Month(), filter.CreatedDateRange.To.Day(), 23, 59, 59, 0, filter.CreatedDateRange.To.Location())
				}
				query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANSLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var citizengravianslog []models.RefCitizenGraviansLog
	if err = cursor.All(context.TODO(), &citizengravianslog); err != nil {
		return nil, err
	}
	return citizengravianslog, nil
}
