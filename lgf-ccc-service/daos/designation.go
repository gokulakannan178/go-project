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

// SaveDesignation : ""
func (d *Daos) SaveDesignation(ctx *models.Context, designation *models.Designation) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).InsertOne(ctx.CTX, designation)
	if err != nil {
		return err
	}
	designation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateDesignation : ""
func (d *Daos) UpdateDesignation(ctx *models.Context, designation *models.Designation) error {
	selector := bson.M{"uniqueId": designation.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": designation}
	_, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleDesignation : ""
func (d *Daos) GetSingleDesignation(ctx *models.Context, uniqueID string) (*models.RefDesignation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var designations []models.RefDesignation
	var designation *models.RefDesignation
	if err = cursor.All(ctx.CTX, &designations); err != nil {
		return nil, err
	}
	if len(designations) > 0 {
		designation = &designations[0]
	}
	return designation, err
}

// EnableDesignation : ""
func (d *Daos) EnableDesignation(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.DESIGNATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableDesignation : ""
func (d *Daos) DisableDesignation(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.DESIGNATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteDesignation :""
func (d *Daos) DeleteDesignation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DESIGNATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDesignation : ""
func (d *Daos) FilterDesignation(ctx *models.Context, filter *models.FilterDesignation, pagination *models.Pagination) ([]models.RefDesignation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDESIGNATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var designationFilter []models.RefDesignation
	if err = cursor.All(context.TODO(), &designationFilter); err != nil {
		return nil, err
	}
	return designationFilter, nil
}
