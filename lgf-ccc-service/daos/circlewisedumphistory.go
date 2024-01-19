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

// SaveCircleWiseDumpHistory : ""
func (d *Daos) SaveCircleWiseDumpHistory(ctx *models.Context, CircleWiseDumpHistory *models.CircleWiseDumpHistory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).InsertOne(ctx.CTX, CircleWiseDumpHistory)
	if err != nil {
		return err
	}
	CircleWiseDumpHistory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateState : ""
func (d *Daos) UpdateCircleWiseDumpHistory(ctx *models.Context, dept *models.CircleWiseDumpHistory) error {
	selector := bson.M{"uniqueId": dept.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dept, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleCircleWiseDumpHistory : ""
func (d *Daos) GetSingleCircleWiseDumpHistory(ctx *models.Context, uniqueID string) (*models.RefCircleWiseDumpHistory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CircleWiseDumpHistorys []models.RefCircleWiseDumpHistory
	var CircleWiseDumpHistory *models.RefCircleWiseDumpHistory
	if err = cursor.All(ctx.CTX, &CircleWiseDumpHistorys); err != nil {
		return nil, err
	}
	if len(CircleWiseDumpHistorys) > 0 {
		CircleWiseDumpHistory = &CircleWiseDumpHistorys[0]
	}
	return CircleWiseDumpHistory, err
}

// EnableCircleWiseDumpHistory : ""
func (d *Daos) EnableCircleWiseDumpHistory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CIRCLEWISEDUMPHISTORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableCircleWiseDumpHistory : ""
func (d *Daos) DisableCircleWiseDumpHistory(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CIRCLEWISEDUMPHISTORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteCircleWiseDumpHistory :""
func (d *Daos) DeleteCircleWiseDumpHistory(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CIRCLEWISEDUMPHISTORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCircleWiseDumpHistory : ""
func (d *Daos) FilterCircleWiseDumpHistory(ctx *models.Context, filter *models.FilterCircleWiseDumpHistory, pagination *models.Pagination) ([]models.RefCircleWiseDumpHistory, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CircleWiseDumpHistoryFilter []models.RefCircleWiseDumpHistory
	if err = cursor.All(context.TODO(), &CircleWiseDumpHistoryFilter); err != nil {
		return nil, err
	}
	return CircleWiseDumpHistoryFilter, nil
}
