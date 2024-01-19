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

// SaveVehicleTrip : ""
func (d *Daos) SaveVehicleTrip(ctx *models.Context, VehicleTrip *models.VehicleTrip) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).InsertOne(ctx.CTX, VehicleTrip)
	if err != nil {
		return err
	}
	VehicleTrip.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateState : ""
func (d *Daos) UpdateVehicleTrip(ctx *models.Context, VehicleTrip *models.VehicleTrip) error {
	selector := bson.M{"uniqueId": VehicleTrip.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": VehicleTrip, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleVehicleTrip : ""
func (d *Daos) GetSingleVehicleTrip(ctx *models.Context, uniqueID string) (*models.RefVehicleTrip, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVECHILE, "vehicle.id", "uniqueId", "ref.vehicle", "ref.vehicle")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var VehicleTrips []models.RefVehicleTrip
	var VehicleTrip *models.RefVehicleTrip
	if err = cursor.All(ctx.CTX, &VehicleTrips); err != nil {
		return nil, err
	}
	if len(VehicleTrips) > 0 {
		VehicleTrip = &VehicleTrips[0]
	}
	return VehicleTrip, err
}

// EnableVehicleTrip : ""
func (d *Daos) EnableVehicleTrip(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.VEHICLETRIPSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableVehicleTrip : ""
func (d *Daos) DisableVehicleTrip(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.VEHICLETRIPSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteVehicleTrip :""
func (d *Daos) DeleteVehicleTrip(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLETRIPSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterVehicleTrip : ""
func (d *Daos) FilterVehicleTrip(ctx *models.Context, filter *models.FilterVehicleTrip, pagination *models.Pagination) ([]models.RefVehicleTrip, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVECHILE, "vehicle.id", "uniqueId", "ref.vehicle", "ref.vehicle")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETRIP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var VehicleTrip []models.RefVehicleTrip
	if err = cursor.All(context.TODO(), &VehicleTrip); err != nil {
		return nil, err
	}
	return VehicleTrip, nil
}
