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

// SaveSolidWasteUserChargeSubCategory : ""
func (d *Daos) SaveSolidWasteUserChargeSubCategory(ctx *models.Context, solidwasteuserchargesubcategory *models.SolidWasteUserChargeSubCategory) error {
	d.Shared.BsonToJSONPrint(solidwasteuserchargesubcategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).InsertOne(ctx.CTX, solidwasteuserchargesubcategory)
	return err
}

// GetSingleSolidWasteUserChargeSubCategory : ""
func (d *Daos) GetSingleSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserChargeSubCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteuserchargesubcategorys []models.RefSolidWasteUserChargeSubCategory
	var solidwasteuserchargesubcategory *models.RefSolidWasteUserChargeSubCategory
	if err = cursor.All(ctx.CTX, &solidwasteuserchargesubcategorys); err != nil {
		return nil, err
	}
	if len(solidwasteuserchargesubcategorys) > 0 {
		solidwasteuserchargesubcategory = &solidwasteuserchargesubcategorys[0]
	}
	return solidwasteuserchargesubcategory, nil
}

// UpdateSolidWasteUserChargeSubCategory : ""
func (d *Daos) UpdateSolidWasteUserChargeSubCategory(ctx *models.Context, solidwasteuserchargesubcategory *models.SolidWasteUserChargeSubCategory) error {
	selector := bson.M{"uniqueId": solidwasteuserchargesubcategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": solidwasteuserchargesubcategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableSolidWasteUserChargeSubCategory : ""
func (d *Daos) EnableSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGESUBCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableSolidWasteUserChargeSubCategory : ""
func (d *Daos) DisableSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGESUBCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteSolidWasteUserChargeSubCategory : ""
func (d *Daos) DeleteSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGESUBCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSolidWasteUserChargeSubCategory : ""
func (d *Daos) FilterSolidWasteUserChargeSubCategory(ctx *models.Context, filter *models.SolidWasteUserChargeSubCategoryFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserChargeSubCategory, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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

	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteuserchargesubcategory []models.RefSolidWasteUserChargeSubCategory
	if err = cursor.All(context.TODO(), &solidwasteuserchargesubcategory); err != nil {
		return nil, err
	}
	return solidwasteuserchargesubcategory, nil
}
