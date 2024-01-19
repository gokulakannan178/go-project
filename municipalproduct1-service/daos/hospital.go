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

// SaveHospital : ""
func (d *Daos) SaveHospital(ctx *models.Context, hospital *models.Hospital) error {
	d.Shared.BsonToJSONPrint(hospital)
	_, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).InsertOne(ctx.CTX, hospital)
	return err
}

// GetSingleHospital : ""
func (d *Daos) GetSingleHospital(ctx *models.Context, UniqueID string) (*models.RefHospital, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONHOSPITALCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONHOSPITALSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var hospitals []models.RefHospital
	var hospital *models.RefHospital
	if err = cursor.All(ctx.CTX, &hospitals); err != nil {
		return nil, err
	}
	if len(hospitals) > 0 {
		hospital = &hospitals[0]
	}
	return hospital, nil
}

// UpdateHospital : ""
func (d *Daos) UpdateHospital(ctx *models.Context, Hospital *models.Hospital) error {
	selector := bson.M{"uniqueId": Hospital.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": Hospital}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableHospital : ""
func (d *Daos) EnableHospital(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HOSPITALSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableHospital : ""
func (d *Daos) DisableHospital(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HOSPITALSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteHospital : ""
func (d *Daos) DeleteHospital(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HOSPITALSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterHospital : ""
func (d *Daos) FilterHospital(ctx *models.Context, filter *models.HospitalFilter, pagination *models.Pagination) ([]models.RefHospital, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		//regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).CountDocuments(ctx.CTX, func() bson.M {
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
	// //Lookup
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONHOSPITALCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONHOSPITALSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOSPITAL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var hospital []models.RefHospital
	if err = cursor.All(context.TODO(), &hospital); err != nil {
		return nil, err
	}
	return hospital, nil
}
