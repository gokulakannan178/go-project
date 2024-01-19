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

// SaveChecklistMaster : ""
func (d *Daos) SaveChecklistMaster(ctx *models.Context, ChecklistMaster *models.ChecklistMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).InsertOne(ctx.CTX, ChecklistMaster)
	if err != nil {
		return err
	}
	ChecklistMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// UpdateChecklistMaster : ""
func (d *Daos) UpdateChecklistMaster(ctx *models.Context, ChecklistMaster *models.ChecklistMaster) error {
	selector := bson.M{"uniqueId": ChecklistMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ChecklistMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleChecklistMaster : ""
func (d *Daos) GetSingleChecklistMaster(ctx *models.Context, uniqueID string) (*models.RefChecklistMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ChecklistMasters []models.RefChecklistMaster
	var ChecklistMaster *models.RefChecklistMaster
	if err = cursor.All(ctx.CTX, &ChecklistMasters); err != nil {
		return nil, err
	}
	if len(ChecklistMasters) > 0 {
		ChecklistMaster = &ChecklistMasters[0]
	}
	return ChecklistMaster, err
}

// EnableChecklistMaster : ""
func (d *Daos) EnableChecklistMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CHECKLISTMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableChecklistMaster : ""
func (d *Daos) DisableChecklistMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CHECKLISTMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteState :""
func (d *Daos) DeleteChecklistMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CHECKLISTMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterChecklistMaster : ""
func (d *Daos) FilterChecklistMaster(ctx *models.Context, filter *models.FilterChecklistMaster, pagination *models.Pagination) ([]models.RefChecklistMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.ContactNo != "" {
			query = append(query, bson.M{"mobile": primitive.Regex{Pattern: filter.Regex.ContactNo, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ChecklistMasterFilter []models.RefChecklistMaster
	if err = cursor.All(context.TODO(), &ChecklistMasterFilter); err != nil {
		return nil, err
	}
	return ChecklistMasterFilter, nil
}
