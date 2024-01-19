package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveProductConfiguration : ""
func (d *Daos) SaveProductConfiguration(ctx *models.Context, pc *models.ProductConfiguration) error {
	d.Shared.BsonToJSONPrint(pc)

	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).InsertOne(ctx.CTX, pc)
	return err
}

// UpsertProductConfiguration : ""
func (d *Daos) UpsertProductConfiguration(ctx *models.Context, pc *models.ProductConfiguration) error {
	d.Shared.BsonToJSONPrint(pc)
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"isDefault": true}
	updateData := bson.M{"$set": pc}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil

}

func (d *Daos) GetSingleDefaultProductConfiguration(ctx *models.Context) (*models.RefProductConfiguration, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pcs []models.RefProductConfiguration
	var pc *models.RefProductConfiguration
	if err = cursor.All(ctx.CTX, &pcs); err != nil {
		return nil, err
	}
	if len(pcs) > 0 {
		pc = &pcs[0]
	}
	return pc, err
}

//GetSingleProductConfiguration : ""
func (d *Daos) GetSingleProductConfiguration(ctx *models.Context, uniqueID string) (*models.RefProductConfiguration, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pcs []models.RefProductConfiguration
	var pc *models.RefProductConfiguration
	if err = cursor.All(ctx.CTX, &pcs); err != nil {
		return nil, err
	}
	if len(pcs) > 0 {
		pc = &pcs[0]
	}
	return pc, err
}

//GetProductLogo : ""
func (d *Daos) GetProductLogo(ctx *models.Context, uniqueID string) (*models.Logo, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ls []models.Logo
	var l *models.Logo
	if err = cursor.All(ctx.CTX, &ls); err != nil {
		return nil, err
	}
	if len(ls) > 0 {
		l = &ls[0]
	}
	return l, err
}

//GetWatermarkLogo : ""
func (d *Daos) GetWatermarkLogo(ctx *models.Context, uniqueID string) (*models.WatermarkLogo, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var wls []models.WatermarkLogo
	var wl *models.WatermarkLogo
	if err = cursor.All(ctx.CTX, &wls); err != nil {
		return nil, err
	}
	if len(wls) > 0 {
		wl = &wls[0]
	}
	return wl, err
}

// FilterProductConfiguration : ""
func (d *Daos) FilterProductConfiguration(ctx *models.Context, filter *models.ProductConfigurationFilter, pagination *models.Pagination) ([]models.RefProductConfiguration, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	// if filter != nil {
	// 	if len(filter.Status) > 0 {
	// 		query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
	// 	}
	// }
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIGURATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pmTarget []models.RefProductConfiguration
	if err = cursor.All(context.TODO(), &pmTarget); err != nil {
		return nil, err
	}
	return pmTarget, nil
}
