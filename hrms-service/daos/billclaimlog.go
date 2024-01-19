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

//SaveBillClaimLog :""
func (d *Daos) SaveBillClaimLog(ctx *models.Context, billClaimLog *models.BillClaimLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).InsertOne(ctx.CTX, billClaimLog)
	if err != nil {
		return err
	}
	billClaimLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleBillClaimLog : ""
func (d *Daos) GetSingleBillClaimLog(ctx *models.Context, uniqueID string) (*models.RefBillClaimLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBillClaimLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billClaimLogs []models.RefBillClaimLog
	var billClaimLog *models.RefBillClaimLog
	if err = cursor.All(ctx.CTX, &billClaimLogs); err != nil {
		return nil, err
	}
	if len(billClaimLogs) > 0 {
		billClaimLog = &billClaimLogs[0]
	}
	return billClaimLog, nil
}

//UpdateBillClaimLog : ""
func (d *Daos) UpdateBillClaimLog(ctx *models.Context, billClaimLog *models.BillClaimLog) error {
	selector := bson.M{"uniqueId": billClaimLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": billClaimLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableBillClaimLog :""
func (d *Daos) EnableBillClaimLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBillClaimLog :""
func (d *Daos) DisableBillClaimLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBillClaimLog :""
func (d *Daos) DeleteBillClaimLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterBillClaimLog : ""
func (d *Daos) FilterBillClaimLog(ctx *models.Context, filter *models.FilterBillClaimLog, pagination *models.Pagination) ([]models.RefBillClaimLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": filter.EmployeeId}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBillClaimLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("BillClaimLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billClaimLogs []models.RefBillClaimLog
	if err = cursor.All(context.TODO(), &billClaimLogs); err != nil {
		return nil, err
	}
	return billClaimLogs, nil
}
