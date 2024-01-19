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

// SaveChecklist : ""
func (d *Daos) SaveChecklist(ctx *models.Context, Checklist *models.Checklist) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).InsertOne(ctx.CTX, Checklist)
	if err != nil {
		return err
	}
	Checklist.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// UpdateChecklist : ""
func (d *Daos) UpdateChecklist(ctx *models.Context, Checklist *models.Checklist) error {
	selector := bson.M{"uniqueId": Checklist.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Checklist}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleChecklist : ""
func (d *Daos) GetSingleChecklist(ctx *models.Context, uniqueID string) (*models.RefChecklist, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVECHILE, "vehicleId", "uniqueId", "ref.vehicle", "ref.vehicle")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Checklists []models.RefChecklist
	var Checklist *models.RefChecklist
	if err = cursor.All(ctx.CTX, &Checklists); err != nil {
		return nil, err
	}
	if len(Checklists) > 0 {
		Checklist = &Checklists[0]
	}
	return Checklist, err
}

// EnableChecklist : ""
func (d *Daos) EnableChecklist(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CHECKLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableChecklist : ""
func (d *Daos) DisableChecklist(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CHECKLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteState :""
func (d *Daos) DeleteChecklist(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CHECKLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterChecklist : ""
func (d *Daos) FilterChecklist(ctx *models.Context, filter *models.FilterChecklist, pagination *models.Pagination) ([]models.RefChecklist, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		if len(filter.CheckUserId) > 0 {
			query = append(query, bson.M{"checkBy.id": bson.M{"$in": filter.CheckUserId}})
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
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVECHILE, "vehicleId", "uniqueId", "ref.vehicle", "ref.vehicle")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ChecklistFilter []models.RefChecklist
	if err = cursor.All(context.TODO(), &ChecklistFilter); err != nil {
		return nil, err
	}
	return ChecklistFilter, nil
}
