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

// SaveSolidWasteUserChargeCategory : ""
func (d *Daos) SaveSolidWasteUserChargeCategory(ctx *models.Context, solidwasteuserchargecategory *models.SolidWasteUserChargeCategory) error {
	d.Shared.BsonToJSONPrint(solidwasteuserchargecategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).InsertOne(ctx.CTX, solidwasteuserchargecategory)
	return err
}

// GetSingleSolidWasteUserChargeCategory : ""
func (d *Daos) GetSingleSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserChargeCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteuserchargecategorys []models.RefSolidWasteUserChargeCategory
	var solidwasteuserchargecategory *models.RefSolidWasteUserChargeCategory
	if err = cursor.All(ctx.CTX, &solidwasteuserchargecategorys); err != nil {
		return nil, err
	}
	if len(solidwasteuserchargecategorys) > 0 {
		solidwasteuserchargecategory = &solidwasteuserchargecategorys[0]
	}
	return solidwasteuserchargecategory, nil
}

// UpdateSolidWasteUserChargeCategory : ""
func (d *Daos) UpdateSolidWasteUserChargeCategory(ctx *models.Context, solidwasteuserchargecategory *models.SolidWasteUserChargeCategory) error {
	selector := bson.M{"uniqueId": solidwasteuserchargecategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": solidwasteuserchargecategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableSolidWasteUserChargeCategory : ""
func (d *Daos) EnableSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGERATESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableSolidWasteUserChargeCategory : ""
func (d *Daos) DisableSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGERATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteSolidWasteUserChargeCategory : ""
func (d *Daos) DeleteSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGERATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSolidWasteUserChargeCategory : ""
func (d *Daos) FilterSolidWasteUserChargeCategory(ctx *models.Context, filter *models.SolidWasteUserChargeCategoryFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserChargeCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueIDs}})
		}
		//regex

		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.Regex.UniqueID, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteuserchargecategory []models.RefSolidWasteUserChargeCategory
	if err = cursor.All(context.TODO(), &solidwasteuserchargecategory); err != nil {
		return nil, err
	}
	return solidwasteuserchargecategory, nil
}
