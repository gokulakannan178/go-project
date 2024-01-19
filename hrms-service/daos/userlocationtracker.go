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

//SaveUserLocationTracker : ""
func (d *Daos) SaveUserLocationTracker(ctx *models.Context, userLocationTracker *models.UserLocationTracker) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).InsertOne(ctx.CTX, userLocationTracker)
	if err != nil {
		return err
	}
	userLocationTracker.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateUserLocationTracker : ""
func (d *Daos) UpdateUserLocationTracker(ctx *models.Context, userLocationTracker *models.UserLocationTracker) error {
	selector := bson.M{"uniqueId": userLocationTracker.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": userLocationTracker}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleUserLocationTracker : ""
func (d *Daos) GetSingleUserLocationTracker(ctx *models.Context, uniqueID string) (*models.RefUserLocationTracker, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserLocationTrackers []models.RefUserLocationTracker
	var UserLocationTracker *models.RefUserLocationTracker
	if err = cursor.All(ctx.CTX, &UserLocationTrackers); err != nil {
		return nil, err
	}
	if len(UserLocationTrackers) > 0 {
		UserLocationTracker = &UserLocationTrackers[0]
	}
	return UserLocationTracker, err
}

// EnableUserLocationTracker : ""
func (d *Daos) EnableUserLocationTracker(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.USERLOCATIONTRACKERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableUserLocationTracker : ""
func (d *Daos) DisableUserLocationTracker(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.USERLOCATIONTRACKERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteUserLocationTracker :""
func (d *Daos) DeleteUserLocationTracker(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERLOCATIONTRACKERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserLocationTracker : ""
func (d *Daos) FilterUserLocationTracker(ctx *models.Context, userLocationTrackerFilter *models.FilterUserLocationTracker, pagination *models.Pagination) ([]models.RefUserLocationTracker, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if userLocationTrackerFilter != nil {
		if len(userLocationTrackerFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": userLocationTrackerFilter.Status}})
		}
		if len(userLocationTrackerFilter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": userLocationTrackerFilter.OrganisationId}})
		}
		//Regex
		if userLocationTrackerFilter.Regex.UserName != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: userLocationTrackerFilter.Regex.UserName, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if userLocationTrackerFilter != nil {
		if userLocationTrackerFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{userLocationTrackerFilter.SortBy: userLocationTrackerFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserLocationTrackerFilter []models.RefUserLocationTracker
	if err = cursor.All(context.TODO(), &UserLocationTrackerFilter); err != nil {
		return nil, err
	}
	return UserLocationTrackerFilter, nil
}
