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

// SaveBranch : ""
func (d *Daos) SaveBranch(ctx *models.Context, branch *models.Branch) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).InsertOne(ctx.CTX, branch)
	if err != nil {
		return err
	}
	branch.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// UpdateBranch : ""
func (d *Daos) UpdateBranch(ctx *models.Context, branch *models.Branch) error {
	selector := bson.M{"uniqueId": branch.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": branch}
	_, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleBranch : ""
func (d *Daos) GetSingleBranch(ctx *models.Context, uniqueID string) (*models.RefBranch, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var branchs []models.RefBranch
	var branch *models.RefBranch
	if err = cursor.All(ctx.CTX, &branchs); err != nil {
		return nil, err
	}
	if len(branchs) > 0 {
		branch = &branchs[0]
	}
	return branch, err
}

// EnableBranch : ""
func (d *Daos) EnableBranch(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.BRANCHSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableBranch : ""
func (d *Daos) DisableBranch(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.BRANCHSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteState :""
func (d *Daos) DeleteBranch(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BRANCHSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterBranch : ""
func (d *Daos) FilterBranch(ctx *models.Context, filter *models.FilterBranch, pagination *models.Pagination) ([]models.RefBranch, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBRANCH).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var branchFilter []models.RefBranch
	if err = cursor.All(context.TODO(), &branchFilter); err != nil {
		return nil, err
	}
	return branchFilter, nil
}
