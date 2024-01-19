package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveAreaAssignLog :""
func (d *Daos) SaveAreaAssignLog(ctx *models.Context, areaassignlog *models.AreaAssignLog) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).InsertOne(ctx.CTX, areaassignlog)
	return err
}

//GetSingleAreaAssignLog : ""
func (d *Daos) GetSingleAreaAssignLog(ctx *models.Context, UniqueID string) (*models.RefAreaAssignLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var areaassignlogs []models.RefAreaAssignLog
	var areaassignlog *models.RefAreaAssignLog
	if err = cursor.All(ctx.CTX, &areaassignlogs); err != nil {
		return nil, err
	}
	if len(areaassignlogs) > 0 {
		areaassignlog = &areaassignlogs[0]
	}
	return areaassignlog, nil
}
func (d *Daos) GetSingleAreaAssignLogWithUserId(ctx *models.Context, UniqueID string) (*models.RefAreaAssignLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"user.uniqueId": UniqueID, "status": constants.VECHILESTATUSASSIGN}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("AreaAssignLogwithdriverId =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var areaAssignLogs []models.RefAreaAssignLog
	var areaAssignLog *models.RefAreaAssignLog
	if err = cursor.All(ctx.CTX, &areaAssignLogs); err != nil {
		return nil, err
	}
	if len(areaAssignLogs) > 0 {
		areaAssignLog = &areaAssignLogs[0]
	}
	return areaAssignLog, nil
}

//UpdateAreaAssignLog : ""
func (d *Daos) UpdateAreaAssignLog(ctx *models.Context, areaassignlog *models.AreaAssignLog) error {
	selector := bson.M{"uniqueId": areaassignlog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": areaassignlog} //update model(user)
	_, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAreaAssignLog :""
func (d *Daos) EnableAreaAssignLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AREAASSIGNLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAreaAssignLog :""
func (d *Daos) DisableAreaAssignLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AREAASSIGNLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAreaAssignLog :""
func (d *Daos) DeleteAreaAssignLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AREAASSIGNLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) RevokeAreaAssignLog(ctx *models.Context, uniqueId string) error {
	selector := bson.M{"uniqueId": uniqueId}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	updateInterface := bson.M{"$set": bson.M{"status": constants.VECHILESTATUSREVOKE, "endDate": &t}} //update model(user)
	_, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// AreaAssignLogFilter : ""
func (d *Daos) AreaAssignLogFilter(ctx *models.Context, filter *models.AreaAssignLogFilter, pagination *models.Pagination) ([]models.AreaAssignLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.UserId) > 0 {
			query = append(query, bson.M{"user.uniqueId": bson.M{"$in": filter.UserId}})
		}
		// if len(userfilter.Type) > 0 {
		// 	query = append(query, bson.M{"type": bson.M{"$in": userfilter.Type}})
		// }
		// if len(userfilter.OmitID) > 0 {
		// 	query = append(query, bson.M{"userName": bson.M{"$nin": userfilter.OmitID}})
		// }
		// if len(userfilter.OrganisationID) > 0 {
		// 	query = append(query, bson.M{"organisationId": bson.M{"$in": userfilter.OrganisationID}})
		// }

		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		// if userfilter.Regex.Contact != "" {
		// 	query = append(query, bson.M{"mobile": primitive.Regex{Pattern: userfilter.Regex.Contact, Options: "xi"}})
		// }
		// if userfilter.Regex.UserName != "" {
		// 	query = append(query, bson.M{"userName": primitive.Regex{Pattern: userfilter.Regex.UserName, Options: "xi"}})
		// }
	}
	if filter.DateRange.From != nil {
		t := *filter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if filter.DateRange.To != nil {
			t2 := *filter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	// //Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	//d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var areaassignlog []models.AreaAssignLog
	if err := cursor.All(context.TODO(), &areaassignlog); err != nil {
		return nil, err
	}
	return areaassignlog, nil
}

func (d *Daos) GetSingleAreaAssignLogWithWard(ctx *models.Context, UniqueID string) (*models.RefAreaAssignLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"area.wardCode": UniqueID, "status": constants.VECHILESTATUSASSIGN}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("AreaAssignLogwithdriverId =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var areaAssignLogs []models.RefAreaAssignLog
	var areaAssignLog *models.RefAreaAssignLog
	if err = cursor.All(ctx.CTX, &areaAssignLogs); err != nil {
		return nil, err
	}
	if len(areaAssignLogs) > 0 {
		areaAssignLog = &areaAssignLogs[0]
	}
	return areaAssignLog, nil
}

func (d *Daos) GetSingleAreaAssignLogWithZone(ctx *models.Context, UniqueID string) (*models.RefAreaAssignLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"area.zoneCode": UniqueID, "status": constants.VECHILESTATUSASSIGN}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("AreaAssignLogwithdriverId =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAREAASSIGNLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var areaAssignLogs []models.RefAreaAssignLog
	var areaAssignLog *models.RefAreaAssignLog
	if err = cursor.All(ctx.CTX, &areaAssignLogs); err != nil {
		return nil, err
	}
	if len(areaAssignLogs) > 0 {
		areaAssignLog = &areaAssignLogs[0]
	}
	return areaAssignLog, nil
}
