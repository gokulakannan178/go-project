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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *Daos) SaveHelperBeat(ctx *models.Context, helperbeat *models.HelperBeat) error {

	_, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).InsertOne(ctx.CTX, helperbeat)

	// HelperBeat.UniqueID = res.InsertedID.(primitive.ObjectID)
	return err
}
func (d *Daos) SaveHelperBeatWithUpsert(ctx *models.Context, helperbeat *models.HelperBeat) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"beatId": helperbeat.BeatID, "employee.id": helperbeat.Employee.Id}
	updateData := bson.M{"$set": helperbeat}
	if _, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

func (d *Daos) GetSingleHelperBeat(ctx *models.Context, UniqueID string) (*models.RefHelperBeat, error) {
	// id, err := primitive.ObjectIDFromHex(UniqueID)
	// if err != nil {
	// 	return nil, err
	// }
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "employee.id", "userName", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "beatId", "uniqueId", "ref.beatmaster", "ref.beatmaster")...)

	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "ticketId", "uniqueId", "ref.ticketId", "ref.ticketId")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "projectId", "uniqueId", "ref.project", "ref.project")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var helperbeats []models.RefHelperBeat
	var helperbeat *models.RefHelperBeat
	if err = cursor.All(ctx.CTX, &helperbeats); err != nil {
		return nil, err
	}
	if len(helperbeats) > 0 {
		helperbeat = &helperbeats[0]
	}
	return helperbeat, nil
}

func (d *Daos) UpdateHelperBeat(ctx *models.Context, helperbeat *models.HelperBeat) error {

	selector := bson.M{"uniqueId": helperbeat.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": helperbeat}
	_, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterHelperBeat(ctx *models.Context, filter *models.FilterHelperBeat, pagination *models.Pagination) ([]models.RefHelperBeat, error) {
	mainPipeline := []bson.M{}
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTICKET, "ticketId", "uniqueId", "ref.ticket", "ref.ticket")...)

	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"ref.ticket.status": bson.M{"$in": filter.Status}})
		}
		if len(filter.TicketId) > 0 {
			query = append(query, bson.M{"ticketId": bson.M{"$in": filter.TicketId}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"employee.id": bson.M{"$in": filter.EmployeeId}})
		}
		if len(filter.BeatID) > 0 {
			query = append(query, bson.M{"beatId": bson.M{"$in": filter.BeatID}})
		}
		if len(filter.ProjectId) > 0 {
			query = append(query, bson.M{"projectId": bson.M{"$in": filter.ProjectId}})
		}
		if len(filter.IsStatus) > 0 {
			query = append(query, bson.M{"ref.ticket.isStatus": bson.M{"$in": filter.IsStatus}})
		}

		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.TicketId != "" {
			query = append(query, bson.M{"ticketId": primitive.Regex{Pattern: filter.Regex.TicketId, Options: "xi"}})
		}

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
	// if pagination != nil {
	// 	mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
	// 	//Getting Total count
	// 	totalCount, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).CountDocuments(ctx.CTX, func() bson.M {
	// 		if query != nil {
	// 			if len(query) > 0 {
	// 				return bson.M{"$and": query}
	// 			}
	// 		}
	// 		return bson.M{}
	// 	}())
	// 	if err != nil {
	// 		log.Println("Error in getting pagination")
	// 	}
	// 	fmt.Println("count", totalCount)
	// 	pagination.Count = int(totalCount)
	// 	d.Shared.PaginationData(pagination)
	// }
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("Commodity Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "employee.id", "userName", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "beatId", "uniqueId", "ref.beatmaster", "ref.beatmaster")...)
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "projectId", "uniqueId", "ref.project", "ref.project")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Vechile query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var HelperBeat []models.RefHelperBeat
	if err = cursor.All(context.TODO(), &HelperBeat); err != nil {
		return nil, err
	}
	return HelperBeat, nil
}

func (d *Daos) EnableHelperBeat(ctx *models.Context, UniqueID string) error {
	// id, err := primitive.ObjectIDFromHex(UniqueID)
	// if err != nil {
	// 	return err
	// }
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HELPERBEATSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableHelperBeat(ctx *models.Context, UniqueID string) error {
	// id, err := primitive.ObjectIDFromHex(UniqueID)
	// if err != nil {
	// 	return err
	// }
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HELPERBEATSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteHelperBeat(ctx *models.Context, UniqueID string) error {
	// id, err := primitive.ObjectIDFromHex(UniqueID)
	// if err != nil {
	// 	return err
	// }
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HELPERBEATSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) RemoveInActiveHelperInBeat(ctx *models.Context, TicketId string, EmployeeId []string) error {
	selector := bson.M{"beatId": TicketId, "employee.id": bson.M{"$nin": EmployeeId}}
	d.Shared.BsonToJSONPrintTag("selector query in onboarding checklist =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.HELPERBEATSTATUSDELETED}}
	d.Shared.BsonToJSONPrintTag("data query in onboarding checklist =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONHELPERBEAT).UpdateMany(ctx.CTX, selector, data)
	return err
}
