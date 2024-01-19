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

//SaveVehicleInsurance :""
func (d *Daos) SaveVehicleInsurance(ctx *models.Context, vehicleinsurance *models.VehicleInsurance) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).InsertOne(ctx.CTX, vehicleinsurance)
	return err
}

//GetSingleVehicleInsurance : ""
func (d *Daos) GetSingleVehicleInsurance(ctx *models.Context, UniqueID string) (*models.RefVehicleInsurance, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehicleinsurances []models.RefVehicleInsurance
	var vehicleinsurance *models.RefVehicleInsurance
	if err = cursor.All(ctx.CTX, &vehicleinsurances); err != nil {
		return nil, err
	}
	if len(vehicleinsurances) > 0 {
		vehicleinsurance = &vehicleinsurances[0]
	}
	return vehicleinsurance, nil
}
func (d *Daos) GetSingleVehicleInsuranceWithDriverID(ctx *models.Context, UniqueID string) (*models.RefVehicleInsurance, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"driver.uniqueId": UniqueID, "status": constants.VECHILESTATUSASSIGN}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var VehicleInsurances []models.RefVehicleInsurance
	var VehicleInsurance *models.RefVehicleInsurance
	if err = cursor.All(ctx.CTX, &VehicleInsurances); err != nil {
		return nil, err
	}
	if len(VehicleInsurances) > 0 {
		VehicleInsurance = &VehicleInsurances[0]
	}
	return VehicleInsurance, nil
}

//UpdateVehicleInsurance : ""
func (d *Daos) UpdateVehicleInsurance(ctx *models.Context, vehicleinsurance *models.VehicleInsurance) error {
	selector := bson.M{"uniqueId": vehicleinsurance.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": vehicleinsurance} //update model(user)
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableVehicleInsurance :""
func (d *Daos) EnableVehicleInsurance(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLEINSURANCESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableVehicleInsurance :""
func (d *Daos) DisableVehicleInsurance(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLEINSURANCESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVehicleInsurance :""
func (d *Daos) DeleteVehicleInsurance(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VEHICLEINSURANCESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// VehicleInsuranceFilter : ""
func (d *Daos) VehicleInsuranceFilter(ctx *models.Context, filter *models.VehicleInsuranceFilter, pagination *models.Pagination) ([]models.VehicleInsurance, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": filter.UniqueID}})
		}
		// if len(userfilter.Type) > 0 {
		// 	query = append(query, bson.M{"type": bson.M{"$in": userfilter.Type}})
		// }
		// if len(userfilter.OmitID) > 0 {
		// 	query = append(query, bson.M{"userName": bson.M{"$nin": userfilter.OmitID}})
		// }
		// if len(userfilter.OrganisationID) > 0 {
		// 	query = append(query, bson.M{"organisationId": bson.M{"$in": userfilter.OrganisationID}})
		// }

		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		// if userfilter.Regex.Contact != "" {
		// 	query = append(query, bson.M{"mobile": primitive.Regex{Pattern: userfilter.Regex.Contact, Options: "xi"}})
		// }
		// if userfilter.Regex.UserName != "" {
		// 	query = append(query, bson.M{"userName": primitive.Regex{Pattern: userfilter.Regex.UserName, Options: "xi"}})
		// }
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
	// //Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVEHICLEINSURANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vehicleinsuranceFilter []models.VehicleInsurance
	if err := cursor.All(context.TODO(), &vehicleinsuranceFilter); err != nil {
		return nil, err
	}
	return vehicleinsuranceFilter, nil
}
