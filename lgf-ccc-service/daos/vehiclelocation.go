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

// SaveVehicleLocation : ""
func (d *Daos) SaveVehicleLocation(ctx *models.Context, vehiclelocation *models.VehicleLocation) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).InsertOne(ctx.CTX, vehiclelocation)
	if err != nil {
		return err
	}
	vehiclelocation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateState : ""
func (d *Daos) UpdateVehicleLocation(ctx *models.Context, vehiclelocation *models.VehicleLocation) error {
	selector := bson.M{"uniqueId": vehiclelocation.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": vehiclelocation, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleVehicleLocation : ""
func (d *Daos) GetSingleVehicleLocation(ctx *models.Context, uniqueID string) (*models.RefVehicleLocation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehiclelocations []models.RefVehicleLocation
	var vehiclelocation *models.RefVehicleLocation
	if err = cursor.All(ctx.CTX, &vehiclelocations); err != nil {
		return nil, err
	}
	if len(vehiclelocations) > 0 {
		vehiclelocation = &vehiclelocations[0]
	}
	return vehiclelocation, err
}

// EnableVehicleLocation : ""
func (d *Daos) EnableVehicleLocation(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.VEHICLELOCATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableVehicleLocation : ""
func (d *Daos) DisableVehicleLocation(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.VEHICLELOCATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteVehicleLocation :""
func (d *Daos) DeleteVehicleLocation(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLELOCATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterVehicleLocation : ""
func (d *Daos) FilterVehicleLocation(ctx *models.Context, filter *models.FilterVehicleLocation, pagination *models.Pagination) ([]models.RefVehicleLocation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.VehicleId) > 0 {
			query = append(query, bson.M{"vehicleId": bson.M{"$in": filter.VehicleId}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLELOCATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehiclelocation []models.RefVehicleLocation
	if err = cursor.All(context.TODO(), &vehiclelocation); err != nil {
		return nil, err
	}
	return vehiclelocation, nil
}
