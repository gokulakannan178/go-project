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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveOverallPropertyDemand : ""
func (d *Daos) SaveOverallPropertyDemand(ctx *models.Context, opd *models.OverallPropertyDemand) error {
	d.Shared.BsonToJSONPrint(opd)
	_, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).InsertOne(ctx.CTX, opd)
	return err
}

// GetSingleOverallPropertyDemand  : ""
func (d *Daos) GetSingleOverallPropertyDemand(ctx *models.Context, PropertyID string) (*models.RefOverallPropertyDemand, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"propertyId": PropertyID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefOverallPropertyDemand
	var tower *models.RefOverallPropertyDemand
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateOverallPropertyDemand: ""
func (d *Daos) UpdateOverallPropertyDemand(ctx *models.Context, opd *models.OverallPropertyDemand) error {
	opts := options.Update().SetUpsert(true)
	selector := bson.M{"propertyId": opd.PropertyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": opd}
	_, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).UpdateOne(ctx.CTX, selector, data, opts)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableOverallPropertyDemand : ""
func (d *Daos) EnableOverallPropertyDemand(ctx *models.Context, PropertyID string) error {
	query := bson.M{"propertyId": PropertyID}
	update := bson.M{"$set": bson.M{"status": constants.OVERALLPROPERTYDEMANDSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableOverallPropertyDemand : ""
func (d *Daos) DisableOverallPropertyDemand(ctx *models.Context, PropertyID string) error {
	query := bson.M{"propertyId": PropertyID}
	update := bson.M{"$set": bson.M{"status": constants.OVERALLPROPERTYDEMANDSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteOverallPropertyDemand : ""
func (d *Daos) DeleteOverallPropertyDemand(ctx *models.Context, PropertyID string) error {
	query := bson.M{"propertyId": PropertyID}
	update := bson.M{"$set": bson.M{"status": constants.OVERALLPROPERTYDEMANDSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterOverallPropertyDemand(ctx *models.Context, filter *models.OverallPropertyDemandFilter, pagination *models.Pagination) ([]models.RefOverallPropertyDemand, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if filter.TotalTax != nil {
			if filter.TotalTax.From != 0 {
				query = append(query, bson.M{"totalTax": bson.M{"$gte": filter.TotalTax.From}})
				if filter.TotalTax.To != 0 {
					query = append(query, bson.M{"totalTax": bson.M{"$gte": filter.TotalTax.From, "$lte": filter.TotalTax.To}})
				}
			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var prop []models.RefOverallPropertyDemand
	if err = cursor.All(context.TODO(), &prop); err != nil {
		return nil, err
	}
	return prop, nil
}

// UpdateOverAllPropertyDemandPropertyID :""
func (d *Daos) UpdateOverAllPropertyDemandPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOVERALLPROPERTYDEMAND).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
