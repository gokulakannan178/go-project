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

// SaveSolidWasteUserCharge : ""
func (d *Daos) SaveSolidWasteUserCharge(ctx *models.Context, solidwasteusercharge *models.SolidWasteUserCharge) error {
	d.Shared.BsonToJSONPrint(solidwasteusercharge)
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).InsertOne(ctx.CTX, solidwasteusercharge)
	return err
}

// GetSingleSolidWasteUserCharge : ""
func (d *Daos) GetSingleSolidWasteUserCharge(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserCharge, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteusercharges []models.RefSolidWasteUserCharge
	var solidwasteusercharge *models.RefSolidWasteUserCharge
	if err = cursor.All(ctx.CTX, &solidwasteusercharges); err != nil {
		return nil, err
	}
	if len(solidwasteusercharges) > 0 {
		solidwasteusercharge = &solidwasteusercharges[0]
	}
	return solidwasteusercharge, nil
}

// UpdateSolidWasteUserCharge : ""
func (d *Daos) UpdateSolidWasteUserCharge(ctx *models.Context, solidwasteusercharge *models.SolidWasteUserCharge) error {
	selector := bson.M{"uniqueId": solidwasteusercharge.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": solidwasteusercharge}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableSolidWasteUserCharge : ""
func (d *Daos) EnableSolidWasteUserCharge(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableSolidWasteUserCharge : ""
func (d *Daos) DisableSolidWasteUserCharge(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteSolidWasteUserCharge : ""
func (d *Daos) DeleteSolidWasteUserCharge(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SOLIDWASTEUSERCHARGESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSolidWasteUserCharge : ""
func (d *Daos) FilterSolidWasteUserCharge(ctx *models.Context, filter *models.SolidWasteUserChargeFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserCharge, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueIDs}})
		}
		if len(filter.CategoryID) > 0 {
			query = append(query, bson.M{"categoryId": bson.M{"$in": filter.CategoryID}})
		}
		if len(filter.SubCategoryID) > 0 {
			query = append(query, bson.M{"subCategoryId": bson.M{"$in": filter.SubCategoryID}})
		}
		if len(filter.StateCode) > 0 {
			query = append(query, bson.M{"address.stateCode": bson.M{"$in": filter.StateCode}})
		}
		if len(filter.DistrictCode) > 0 {
			query = append(query, bson.M{"address.districtCode": bson.M{"$in": filter.DistrictCode}})
		}
		if len(filter.VillageCode) > 0 {
			query = append(query, bson.M{"address.villageCode": bson.M{"$in": filter.VillageCode}})
		}
		if len(filter.WardCode) > 0 {
			query = append(query, bson.M{"address.wardCode": bson.M{"$in": filter.WardCode}})
		}
		if len(filter.ZoneCode) > 0 {
			query = append(query, bson.M{"address.zoneCode": bson.M{"$in": filter.ZoneCode}})
		}
		//regex

		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.Regex.UniqueID, Options: "xi"}})
		}
		if filter.Regex.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.Regex.OwnerName, Options: "xi"}})
		}
		if filter.Regex.MobileNo != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.Regex.MobileNo, Options: "xi"}})
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
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var solidwasteusercharge []models.RefSolidWasteUserCharge
	if err = cursor.All(context.TODO(), &solidwasteusercharge); err != nil {
		return nil, err
	}
	return solidwasteusercharge, nil
}
