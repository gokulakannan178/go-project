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

// Saveleavemaster : ""
func (d *Daos) SaveLeaveMaster(ctx *models.Context, leavemaster *models.LeaveMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).InsertOne(ctx.CTX, leavemaster)
	if err != nil {
		return err
	}
	leavemaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateLeaveMaster : ""
func (d *Daos) UpdateLeaveMaster(ctx *models.Context, leavemaster *models.LeaveMaster) error {
	selector := bson.M{"uniqueId": leavemaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": leavemaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleLeaveMaster : ""
func (d *Daos) GetSingleLeaveMaster(ctx *models.Context, uniqueID string) (*models.RefLeaveMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var LeaveMasters []models.RefLeaveMaster
	var LeaveMaster *models.RefLeaveMaster
	if err = cursor.All(ctx.CTX, &LeaveMasters); err != nil {
		return nil, err
	}
	if len(LeaveMasters) > 0 {
		LeaveMaster = &LeaveMasters[0]
	}
	return LeaveMaster, err
}

// GetSingleLeaveMasterWithActive : ""
func (d *Daos) GetSingleLeaveMasterWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefLeaveMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var LeaveMasters []models.RefLeaveMaster
	var LeaveMaster *models.RefLeaveMaster
	if err = cursor.All(ctx.CTX, &LeaveMasters); err != nil {
		return nil, err
	}
	if len(LeaveMasters) > 0 {
		LeaveMaster = &LeaveMasters[0]
	}
	return LeaveMaster, err
}

// EnableLeaveMaster : ""
func (d *Daos) EnableLeaveMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.LEAVEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableLeaveMaster : ""
func (d *Daos) DisableLeaveMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.LEAVEMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteLeaveMaster :""
func (d *Daos) DeleteLeaveMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEAVEMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLeaveMaster : ""
func (d *Daos) FilterLeaveMaster(ctx *models.Context, leavemaster *models.FilterLeaveMaster, pagination *models.Pagination) ([]models.RefLeaveMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if leavemaster != nil {
		if len(leavemaster.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": leavemaster.Status}})
		}
		if len(leavemaster.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": leavemaster.OrganisationID}})
		}
		//Regex
		if leavemaster.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: leavemaster.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if leavemaster != nil {
		if leavemaster.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{leavemaster.SortBy: leavemaster.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEAVEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var leaveMaster []models.RefLeaveMaster
	if err = cursor.All(context.TODO(), &leaveMaster); err != nil {
		return nil, err
	}
	return leaveMaster, nil
}
