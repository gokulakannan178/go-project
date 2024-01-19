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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveOurService : ""
func (d *Daos) SaveOurService(ctx *models.Context, collection string, OurService *models.OurService) error {
	res, err := ctx.DB.Collection(collection).InsertOne(ctx.CTX, OurService)
	if err != nil {
		return err
	}
	OurService.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateOurService : ""
func (d *Daos) UpdateOurService(ctx *models.Context, scenario string, OurService *models.OurService) error {
	selector := bson.M{"uniqueId": OurService.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": OurService}
	_, err := ctx.DB.Collection(scenario).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleOurService : ""
func (d *Daos) GetSingleOurService(ctx *models.Context, scenario string, uniqueID string) (*models.RefOurService, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(scenario).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OurServices []models.RefOurService
	var OurService *models.RefOurService
	if err = cursor.All(ctx.CTX, &OurServices); err != nil {
		return nil, err
	}
	if len(OurServices) > 0 {
		OurService = &OurServices[0]
	}
	return OurService, err
}

// EnableOurService : ""
func (d *Daos) EnableOurService(ctx *models.Context, scenario string, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OURSERVICESTATUSACTIVE}}
	_, err := ctx.DB.Collection(scenario).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableOurService : ""
func (d *Daos) DisableOurService(ctx *models.Context, scenario string, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OURSERVICESTATUSDISABLED}}
	_, err := ctx.DB.Collection(scenario).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteState :""
func (d *Daos) DeleteOurService(ctx *models.Context, scenario string, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OURSERVICESTATUSDELETED}}
	_, err := ctx.DB.Collection(scenario).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOurService : ""
func (d *Daos) FilterOurService(ctx *models.Context, scenario string, filter *models.FilterOurService, pagination *models.Pagination) ([]models.RefOurService, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.ShowOfficeDirectoryInDashboard) > 0 {
			query = append(query, bson.M{"showOfficeDirectoryInDashboard": bson.M{"$in": filter.ShowOfficeDirectoryInDashboard}})
		}

		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.Title != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: filter.Regex.Title, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(scenario).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(scenario).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OurServiceFilter []models.RefOurService
	if err = cursor.All(context.TODO(), &OurServiceFilter); err != nil {
		return nil, err
	}
	return OurServiceFilter, nil
}
func (d *Daos) UpdateManyDisableOurService(ctx *models.Context, scenario string) error {
	selector := bson.M{}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"status": constants.OURSERVICESTATUSDISABLED}}
	_, err := ctx.DB.Collection(scenario).UpdateMany(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
