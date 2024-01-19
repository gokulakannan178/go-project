package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveVehicle :""
func (d *Daos) SaveVehicle(ctx *models.Context, vehicle *models.Vehicle) error {
	_, err := ctx.DB.Collection(constants.COLLVEHICLE).InsertOne(ctx.CTX, vehicle)
	return err
}

//GetSingleVehicle : ""
func (d *Daos) GetSingleVehicle(ctx *models.Context, UniqueID string) (*models.RefVehicle, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"regNo": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLVEHICLE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehicles []models.RefVehicle
	var vehicle *models.RefVehicle
	if err = cursor.All(ctx.CTX, &vehicles); err != nil {
		return nil, err
	}
	if len(vehicles) > 0 {
		vehicle = &vehicles[0]
	}
	return vehicle, nil
}

//UpdateVehicle : ""
func (d *Daos) UpdateVehicle(ctx *models.Context, vehicle *models.Vehicle) error {
	selector := bson.M{"regNo": vehicle.RegNo}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": vehicle, "$push": bson.M{"updatedLog": update}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterVehicle : ""
func (d *Daos) FilterVehicle(ctx *models.Context, vehiclefilter *models.VehicleFilter, pagination *models.Pagination) ([]models.RefVehicle, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if vehiclefilter != nil {

		if len(vehiclefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": vehiclefilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLVEHICLE).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("vehicle query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLVEHICLE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehicles []models.RefVehicle
	if err = cursor.All(context.TODO(), &vehicles); err != nil {
		return nil, err
	}
	return vehicles, nil
}

//EnableVehicle :""
func (d *Daos) EnableVehicle(ctx *models.Context, UniqueID string) error {
	query := bson.M{"regNo": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableVehicle :""
func (d *Daos) DisableVehicle(ctx *models.Context, UniqueID string) error {
	query := bson.M{"regNo": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVehicle :""
func (d *Daos) DeleteVehicle(ctx *models.Context, UniqueID string) error {
	query := bson.M{"regNo": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
