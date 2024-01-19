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

// SaveWardWiseDumpHistory : ""
func (d *Daos) SaveWardWiseDumpHistory(ctx *models.Context, WardWiseDumpHistory *models.WardWiseDumpHistory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).InsertOne(ctx.CTX, WardWiseDumpHistory)
	if err != nil {
		return err
	}
	WardWiseDumpHistory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateState : ""
func (d *Daos) UpdateWardWiseDumpHistory(ctx *models.Context, dept *models.WardWiseDumpHistory) error {
	selector := bson.M{"uniqueId": dept.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dept, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleWardWiseDumpHistory : ""
func (d *Daos) GetSingleWardWiseDumpHistory(ctx *models.Context, uniqueID string) (*models.RefWardWiseDumpHistory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WardWiseDumpHistorys []models.RefWardWiseDumpHistory
	var WardWiseDumpHistory *models.RefWardWiseDumpHistory
	if err = cursor.All(ctx.CTX, &WardWiseDumpHistorys); err != nil {
		return nil, err
	}
	if len(WardWiseDumpHistorys) > 0 {
		WardWiseDumpHistory = &WardWiseDumpHistorys[0]
	}
	return WardWiseDumpHistory, err
}

// EnableWardWiseDumpHistory : ""
func (d *Daos) EnableWardWiseDumpHistory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.WARDWISEDUMPHISTORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableWardWiseDumpHistory : ""
func (d *Daos) DisableWardWiseDumpHistory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.WARDWISEDUMPHISTORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteWardWiseDumpHistory :""
func (d *Daos) DeleteWardWiseDumpHistory(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WARDWISEDUMPHISTORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterWardWiseDumpHistory : ""
func (d *Daos) FilterWardWiseDumpHistory(ctx *models.Context, filter *models.FilterWardWiseDumpHistory, pagination *models.Pagination) ([]models.RefWardWiseDumpHistory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		if len(filter.BranchId) > 0 {
			query = append(query, bson.M{"branch": bson.M{"$in": filter.BranchId}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}

		if filter.Regex.Type != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: filter.Regex.Type, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WardWiseDumpHistoryFilter []models.RefWardWiseDumpHistory
	if err = cursor.All(context.TODO(), &WardWiseDumpHistoryFilter); err != nil {
		return nil, err
	}
	return WardWiseDumpHistoryFilter, nil
}
