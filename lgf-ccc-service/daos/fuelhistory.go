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

//SaveFuelHistory :""
func (d *Daos) SaveFuelHistory(ctx *models.Context, fuelhistory *models.FuelHistory) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).InsertOne(ctx.CTX, fuelhistory)
	return err
}

//GetSingleFuelHistory : ""
func (d *Daos) GetSingleFuelHistory(ctx *models.Context, UniqueID string) (*models.RefFuelHistory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var fuelhistorys []models.RefFuelHistory
	var fuelhistory *models.RefFuelHistory
	if err = cursor.All(ctx.CTX, &fuelhistorys); err != nil {
		return nil, err
	}
	if len(fuelhistorys) > 0 {
		fuelhistory = &fuelhistorys[0]
	}
	return fuelhistory, nil
}
func (d *Daos) GetSingleFuelHistoryWithDriverID(ctx *models.Context, UniqueID string) (*models.RefFuelHistory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"driver.uniqueId": UniqueID, "status": constants.VECHILESTATUSASSIGN}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var fuelhistorys []models.RefFuelHistory
	var fuelhistory *models.RefFuelHistory
	if err = cursor.All(ctx.CTX, &fuelhistorys); err != nil {
		return nil, err
	}
	if len(fuelhistorys) > 0 {
		fuelhistory = &fuelhistorys[0]
	}
	return fuelhistory, nil
}

//UpdateFuelHistory : ""
func (d *Daos) UpdateFuelHistory(ctx *models.Context, fuelhistory *models.FuelHistory) error {
	selector := bson.M{"uniqueId": fuelhistory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": fuelhistory} //update model(user)
	_, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableFuelHistory :""
func (d *Daos) EnableFuelHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FUELHISTORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFuelHistory :""
func (d *Daos) DisableFuelHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FUELHISTORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFuelHistory :""
func (d *Daos) DeleteFuelHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FUELHISTORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FuelHistoryFilter : ""
func (d *Daos) FuelHistoryFilter(ctx *models.Context, filter *models.FuelHistoryFilter, pagination *models.Pagination) ([]models.FuelHistory, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFUELHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var fuelhistoryFilter []models.FuelHistory
	if err := cursor.All(context.TODO(), &fuelhistoryFilter); err != nil {
		return nil, err
	}
	return fuelhistoryFilter, nil
}
