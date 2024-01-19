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

// SavePropertyFixedArvLog : ""
func (d *Daos) SavePropertyFixedArvLog(ctx *models.Context, propertyfixedarvlog *models.PropertyFixedArvLog) error {
	d.Shared.BsonToJSONPrint(propertyfixedarvlog)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).InsertOne(ctx.CTX, propertyfixedarvlog)
	return err
}

// GetSinglePropertyFixedArvLog : ""
func (d *Daos) GetSinglePropertyFixedArvLog(ctx *models.Context, UniqueID string) (*models.RefPropertyFixedArvLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYFIXEDARVLOGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYFIXEDARVLOGSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyfixedarvlogs []models.RefPropertyFixedArvLog
	var propertyfixedarvlog *models.RefPropertyFixedArvLog
	if err = cursor.All(ctx.CTX, &propertyfixedarvlogs); err != nil {
		return nil, err
	}
	if len(propertyfixedarvlogs) > 0 {
		propertyfixedarvlog = &propertyfixedarvlogs[0]
	}
	return propertyfixedarvlog, nil
}

// UpdatePropertyFixedArvLog : ""
func (d *Daos) UpdatePropertyFixedArvLog(ctx *models.Context, propertyfixedarvlog *models.PropertyFixedArvLog) error {
	selector := bson.M{"uniqueId": propertyfixedarvlog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": propertyfixedarvlog}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyFixedArvLog : ""
func (d *Daos) EnablePropertyFixedArvLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyFixedArvLog : ""
func (d *Daos) DisablePropertyFixedArvLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyFixedArvLog : ""
func (d *Daos) DeletePropertyFixedArvLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyFixedArvLog : ""
func (d *Daos) FilterPropertyFixedArvLog(ctx *models.Context, filter *models.PropertyFixedArvLogFilter, pagination *models.Pagination) ([]models.RefPropertyFixedArvLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueIDs}})
		}
		if len(filter.FyID) > 0 {
			query = append(query, bson.M{"fyid": bson.M{"$in": filter.FyID}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if len(filter.ARV) > 0 {
			query = append(query, bson.M{"arv": bson.M{"$in": filter.ARV}})
		}
		if len(filter.CreatedBy) > 0 {
			query = append(query, bson.M{"Created.by": bson.M{"$in": filter.CreatedBy}})
		}
		//regex

		if filter.Regex.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.Regex.UniqueID, Options: "xi"}})
		}
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
	}
	if filter.CreatedDate != nil {
		//var sd,ed time.Time
		if filter.CreatedDate.From != nil {
			sd := time.Date(filter.CreatedDate.From.Year(), filter.CreatedDate.From.Month(), filter.CreatedDate.From.Day(), 0, 0, 0, 0, filter.CreatedDate.From.Location())
			ed := time.Date(filter.CreatedDate.From.Year(), filter.CreatedDate.To.Month(), filter.CreatedDate.To.Day(), 23, 59, 59, 0, filter.CreatedDate.To.Location())
			if filter.CreatedDate.To != nil {
				ed = time.Date(filter.CreatedDate.To.Year(), filter.CreatedDate.To.Month(), filter.CreatedDate.To.Day(), 23, 59, 59, 0, filter.CreatedDate.To.Location())
			}
			query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYFIXEDARVLOGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYFIXEDARVLOGSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARVLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyfixedarvlog []models.RefPropertyFixedArvLog
	if err = cursor.All(context.TODO(), &propertyfixedarvlog); err != nil {
		return nil, err
	}
	return propertyfixedarvlog, nil
}
