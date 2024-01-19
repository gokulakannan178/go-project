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

//SaveDemoUser :""
func (d *Daos) SaveDemoUser(ctx *models.Context, demoUser *models.DemoUser) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).InsertOne(ctx.CTX, demoUser)
	if err != nil {
		return err
	}
	demoUser.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleDemoUser : ""
func (d *Daos) GetSingleDemoUser(ctx *models.Context, UniqueID string) (*models.RefDemoUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("DemoUser query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var demoUsers []models.RefDemoUser
	var demoUser *models.RefDemoUser
	if err = cursor.All(ctx.CTX, &demoUsers); err != nil {
		return nil, err
	}
	if len(demoUsers) > 0 {
		demoUser = &demoUsers[0]
	}
	return demoUser, nil
}

//UpdateDemoUser : ""
func (d *Daos) UpdateDemoUser(ctx *models.Context, demoUser *models.DemoUser) error {
	selector := bson.M{"uniqueId": demoUser.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": demoUser}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableDemoUser :""
func (d *Daos) EnableDemoUser(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEMOUSERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDemoUser :""
func (d *Daos) DisableDemoUser(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEMOUSERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDemoUser :""
func (d *Daos) DeleteDemoUser(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEMOUSERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDemoUser : ""
func (d *Daos) FilterDemoUser(ctx *models.Context, filter *models.DemoUserFilter, pagination *models.Pagination) ([]models.RefDemoUser, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		// if len(filter.OrganisationId) > 0 {
		// 	query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		// }
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("DemoUser query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEMOUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var demoUsers []models.RefDemoUser
	if err = cursor.All(context.TODO(), &demoUsers); err != nil {
		return nil, err
	}
	return demoUsers, nil
}
