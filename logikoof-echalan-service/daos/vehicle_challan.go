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

//SaveVehicleChallan :""
func (d *Daos) SaveVehicleChallan(ctx *models.Context, vehicleChallan *models.VehicleChallan) error {
	_, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).InsertOne(ctx.CTX, vehicleChallan)
	return err
}

//GetSingleVehicleChallan : ""
func (d *Daos) GetSingleVehicleChallan(ctx *models.Context, UniqueID string) (*models.RefVehicleChallan, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLOFFENCETYPE, "offenceType", "uniqueId", "ref.offenceType", "ref.offenceType")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehicleChallans []models.RefVehicleChallan
	var vehicleChallan *models.RefVehicleChallan
	if err = cursor.All(ctx.CTX, &vehicleChallans); err != nil {
		return nil, err
	}
	if len(vehicleChallans) > 0 {
		vehicleChallan = &vehicleChallans[0]
	}
	return vehicleChallan, nil
}

//UpdateVehicleChallan : ""
func (d *Daos) UpdateVehicleChallan(ctx *models.Context, vehicleChallan *models.VehicleChallan) error {
	selector := bson.M{"uniqueId": vehicleChallan.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": vehicleChallan, "$push": bson.M{"updatedLog": update}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterVehicleChallan : ""
func (d *Daos) FilterVehicleChallan(ctx *models.Context, vehicleChallanfilter *models.VehicleChallanFilter, pagination *models.Pagination) ([]models.RefVehicleChallan, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if vehicleChallanfilter != nil {

		if len(vehicleChallanfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": vehicleChallanfilter.Status}})
		}
		if len(vehicleChallanfilter.VehicleType) > 0 {
			query = append(query, bson.M{"vehicle.type": bson.M{"$in": vehicleChallanfilter.VehicleType}})
		}
		if len(vehicleChallanfilter.Mobile) > 0 {
			query = append(query, bson.M{"vehicle.mobile": bson.M{"$in": vehicleChallanfilter.Mobile}})
		}
		if len(vehicleChallanfilter.RegNo) > 0 {
			query = append(query, bson.M{"vehicle.regNo": bson.M{"$in": vehicleChallanfilter.RegNo}})
		}
		// if vehicleChallanfilter.IsOffenceVideo {
		// 	query = append(query, bson.M{"videos": bson.M{"$exists": true, "$ne": []bson.M{}, "$not": bson.M{"$size": 0}}})
		// }
		if len(vehicleChallanfilter.OffenceType) > 0 {
			query = append(query, bson.M{"offenceType": bson.M{"$in": vehicleChallanfilter.OffenceType}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLOFFENCETYPE, "offenceType", "uniqueId", "ref.offenceType", "ref.offenceType")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("vehicleChallan query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehicleChallans []models.RefVehicleChallan
	if err = cursor.All(context.TODO(), &vehicleChallans); err != nil {
		return nil, err
	}
	return vehicleChallans, nil
}

//EnableVehicleChallan :""
func (d *Daos) EnableVehicleChallan(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLECHALLANSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableVehicleChallan :""
func (d *Daos) DisableVehicleChallan(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLECHALLANSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVehicleChallan :""
func (d *Daos) DeleteVehicleChallan(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLECHALLANSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
