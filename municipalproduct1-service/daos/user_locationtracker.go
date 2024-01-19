package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveUserLocationTracker : ""
func (d *Daos) SaveUserLocationTracker(ctx *models.Context, tracker *models.UserLocationTracker) error {
	d.Shared.BsonToJSONPrint(tracker)
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).InsertOne(ctx.CTX, tracker)
	return err
}

// GetSingleUserLocationTracker : ""
func (d *Daos) GetSingleUserLocationTracker(ctx *models.Context, UniqueID string) (*models.RefUserLocationTracker, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	// Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userName", "userName", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "userType", "uniqueId", "ref.userType", "ref.userType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserLocationTracker
	var tower *models.RefUserLocationTracker
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserLocationTracker : ""
func (d *Daos) UpdateUserLocationTracker(ctx *models.Context, tracker *models.UserLocationTracker) error {
	selector := bson.M{"uniqueId": tracker.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": tracker}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserLocationTracker : ""
func (d *Daos) EnableUserLocationTracker(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERLOCATIONTRACKERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserLocationTracker : ""
func (d *Daos) DisableUserLocationTracker(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERLOCATIONTRACKERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserLocationTracker : ""
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
func (d *Daos) FilterUserLocationTracker(ctx *models.Context, filter *models.UserLocationTrackerFilter, pagination *models.Pagination) ([]models.RefUserLocationTracker, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"userType": bson.M{"$in": filter.UserType}})
		}
		//var sd,ed time.Time
		if filter.Date.From != nil {
			sd := time.Date(filter.Date.From.Year(), filter.Date.From.Month(), filter.Date.From.Day(), 0, 0, 0, 0, filter.Date.From.Location())
			ed := time.Date(filter.Date.From.Year(), filter.Date.From.Month(), filter.Date.From.Day(), 23, 59, 59, 0, filter.Date.From.Location())
			if filter.Date.To != nil {
				ed = time.Date(filter.Date.To.Year(), filter.Date.To.Month(), filter.Date.To.Day(), 23, 59, 59, 0, filter.Date.To.Location())
			}
			query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
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
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userName", "userName", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "userType", "uniqueId", "ref.userType", "ref.userType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tracker []models.RefUserLocationTracker
	if err = cursor.All(context.TODO(), &tracker); err != nil {
		return nil, err
	}
	return tracker, nil
}

// GetSingleUserLocationTracker : ""
func (d *Daos) UserLocationTrackerCoordinates(ctx *models.Context, coordinates *models.UserLocationTrackerCoordinates) (*models.RefUserLocationTracker, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if coordinates != nil {
		if coordinates.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": bson.M{"$eq": coordinates.UniqueID}})

		}
		if coordinates.StartDate != nil && coordinates.EndDate != nil {
			query = append(query, bson.M{"timeStamp": bson.M{"$gte": coordinates.StartDate, "$lte": coordinates.EndDate}})

		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	//	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"userName": "$userName"}, "data": bson.M{"$push": "$$ROOT"}}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	// Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userName", "userName", "ref.user", "ref.user")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONTRACKER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserLocationTracker
	var tower *models.RefUserLocationTracker
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}
