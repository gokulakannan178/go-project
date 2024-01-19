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

//SaveVehicleType :""
func (d *Daos) SaveVehicleType(ctx *models.Context, VehicleType *models.VehicleType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).InsertOne(ctx.CTX, VehicleType)
	return err
}

//GetSingleVehicleType : ""
func (d *Daos) GetSingleVehicleType(ctx *models.Context, UniqueID string) (*models.RefVechileType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefVechileType
	var user *models.RefVechileType
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//UpdateVehicleType : ""
func (d *Daos) UpdateVehicleType(ctx *models.Context, user *models.VehicleType) error {
	selector := bson.M{"uniqueId": user.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": user} //update model(user)
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableVehicleType :""
func (d *Daos) EnableVehicleType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLETYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableVehicleType :""
func (d *Daos) DisableVehicleType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLETYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVehicleType :""
func (d *Daos) DeleteVehicleType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLETYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// VehicleTypeFilter : ""
func (d *Daos) VehicleTypeFilter(ctx *models.Context, VehicleTypefilter *models.VehicleTypeFilter, pagination *models.Pagination) ([]models.VehicleType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if VehicleTypefilter != nil {
		if len(VehicleTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": VehicleTypefilter.Status}})
		}
		if len(VehicleTypefilter.UniqueID) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": VehicleTypefilter.UniqueID}})
		}
		if len(VehicleTypefilter.VehicleType) > 0 {
			query = append(query, bson.M{"vehicleType": bson.M{"$in": VehicleTypefilter.VehicleType}})
		}

		//Regex
		if VehicleTypefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: VehicleTypefilter.Regex.Name, Options: "xi"}})
		}
		if VehicleTypefilter.Regex.FuelType != "" {
			query = append(query, bson.M{"fuelType": primitive.Regex{Pattern: VehicleTypefilter.Regex.FuelType, Options: "xi"}})
		}
	}
	if VehicleTypefilter.DateRange.From != nil {
		t := *VehicleTypefilter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if VehicleTypefilter.DateRange.To != nil {
			t2 := *VehicleTypefilter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	// //Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	//d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLETYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var VehicleTypeFilter []models.VehicleType
	if err := cursor.All(context.TODO(), &VehicleTypeFilter); err != nil {
		return nil, err
	}
	return VehicleTypeFilter, nil
}
