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

// SaveSolidWasteUserChargeRate : ""
func (d *Daos) SaveSolidWasteUserChargeRate(ctx *models.Context, solidwasteuserchargerate *models.SolidWasteUserChargeRate) error {
	d.Shared.BsonToJSONPrint(solidwasteuserchargerate)
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).InsertOne(ctx.CTX, solidwasteuserchargerate)
	return err
}

// GetSingleSolidWasteUserChargeRate : ""
func (d *Daos) GetSingleSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserChargeRate, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteuserchargerates []models.RefSolidWasteUserChargeRate
	var solidwasteuserchargerate *models.RefSolidWasteUserChargeRate
	if err = cursor.All(ctx.CTX, &solidwasteuserchargerates); err != nil {
		return nil, err
	}
	if len(solidwasteuserchargerates) > 0 {
		solidwasteuserchargerate = &solidwasteuserchargerates[0]
	}
	return solidwasteuserchargerate, nil
}

// UpdateSolidWasteUserChargeRate : ""
func (d *Daos) UpdateSolidWasteUserChargeRate(ctx *models.Context, solidwasteuserchargerate *models.SolidWasteUserChargeRate) error {
	selector := bson.M{"uniqueId": solidwasteuserchargerate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": solidwasteuserchargerate}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableSolidWasteUserChargeRate : ""
func (d *Daos) EnableSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGERATESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableSolidWasteUserChargeRate : ""
func (d *Daos) DisableSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGERATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteSolidWasteUserChargeRate : ""
func (d *Daos) DeleteSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGERATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSolidWasteUserChargeRate : ""
func (d *Daos) FilterSolidWasteUserChargeRate(ctx *models.Context, filter *models.SolidWasteUserChargeRateFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserChargeRate, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		// if len(filter.Status) > 0 {
		// 	query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		// }
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
		// if len(filter.Description) > 0 {
		// 	query = append(query, bson.M{"description": bson.M{"$in": filter.Description}})
		// }
		// if len(filter.AssignedBy) > 0 {
		// 	query = append(query, bson.M{"assignedBy": bson.M{"$in": filter.AssignedBy}})
		// }
		// if len(filter.Description) > 0 {
		// 	query = append(query, bson.M{"description": bson.M{"$in": filter.Description}})
		// }
		// if filter.FromDateRange != nil {
		// 	//var sd,ed time.Time
		// 	if filter.FromDateRange.From != nil {
		// 		sd := time.Date(filter.FromDateRange.From.Year(), filter.FromDateRange.From.Month(), filter.FromDateRange.From.Day(), 0, 0, 0, 0, filter.FromDateRange.From.Location())
		// 		ed := time.Date(filter.FromDateRange.From.Year(), filter.FromDateRange.From.Month(), filter.FromDateRange.From.Day(), 23, 59, 59, 0, filter.FromDateRange.From.Location())
		// 		if filter.FromDateRange.To != nil {
		// 			ed = time.Date(filter.FromDateRange.To.Year(), filter.FromDateRange.To.Month(), filter.FromDateRange.To.Day(), 23, 59, 59, 0, filter.FromDateRange.To.Location())
		// 		}
		// 		query = append(query, bson.M{"dateFrom": bson.M{"$gte": sd, "$lte": ed}})

		// 	}
		// }
		// if filter.ToDateRange != nil {
		// 	//var sd,ed time.Time
		// 	if filter.ToDateRange.From != nil {
		// 		sd := time.Date(filter.ToDateRange.From.Year(), filter.ToDateRange.From.Month(), filter.ToDateRange.From.Day(), 0, 0, 0, 0, filter.ToDateRange.From.Location())
		// 		ed := time.Date(filter.ToDateRange.From.Year(), filter.ToDateRange.From.Month(), filter.ToDateRange.From.Day(), 23, 59, 59, 0, filter.ToDateRange.From.Location())
		// 		if filter.ToDateRange.To != nil {
		// 			ed = time.Date(filter.ToDateRange.To.Year(), filter.ToDateRange.To.Month(), filter.ToDateRange.To.Day(), 23, 59, 59, 0, filter.ToDateRange.To.Location())
		// 		}
		// 		query = append(query, bson.M{"dateTo": bson.M{"$gte": sd, "$lte": ed}})

		// 	}
		// }
		// if filter.CreatedDateRange != nil {
		// 	//var sd,ed time.Time
		// 	if filter.CreatedDateRange.From != nil {
		// 		sd := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 0, 0, 0, 0, filter.CreatedDateRange.From.Location())
		// 		ed := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 23, 59, 59, 0, filter.CreatedDateRange.From.Location())
		// 		if filter.CreatedDateRange.To != nil {
		// 			ed = time.Date(filter.CreatedDateRange.To.Year(), filter.CreatedDateRange.To.Month(), filter.CreatedDateRange.To.Day(), 23, 59, 59, 0, filter.CreatedDateRange.To.Location())
		// 		}
		// 		query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

		// 	}
		// }
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGERATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteuserchargerate []models.RefSolidWasteUserChargeRate
	if err = cursor.All(context.TODO(), &solidwasteuserchargerate); err != nil {
		return nil, err
	}
	return solidwasteuserchargerate, nil
}
