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

// SaveLetterUpload : ""
func (d *Daos) SaveLetterUpload(ctx *models.Context, lu *models.LetterUpload) error {
	d.Shared.BsonToJSONPrint(lu)
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).InsertOne(ctx.CTX, lu)
	return err
}

// GetSinglePropertyPenalty : ""
func (d *Daos) GetSingleLetterUpload(ctx *models.Context, UniqueID string) (*models.RefLetterUpload, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("letter upload getsingle query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var lus []models.RefLetterUpload
	var lu *models.RefLetterUpload
	if err = cursor.All(ctx.CTX, &lus); err != nil {
		return nil, err
	}
	if len(lus) > 0 {
		lu = &lus[0]
	}
	return lu, nil
}

// UpdatePropertyPenalty : ""
func (d *Daos) UpdateLetterUpload(ctx *models.Context, lu *models.LetterUpload) error {
	selector := bson.M{"uniqueId": lu.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": lu, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableLetterUpload : ""
func (d *Daos) EnableLetterUpload(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERUPLOADSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableLetterUpload : ""
func (d *Daos) DisableLetterUpload(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERUPLOADSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteLetterUpload : ""
func (d *Daos) DeleteLetterUpload(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERUPLOADRSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLetterUpload : ""
func (d *Daos) FilterLetterUpload(ctx *models.Context, filter *models.LetterUploadFilter, pagination *models.Pagination) ([]models.RefLetterUpload, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.No != "" {
			query = append(query, bson.M{"no": primitive.Regex{Pattern: filter.Regex.No, Options: "xi"}})
		}
		if filter.Regex.From != "" {
			query = append(query, bson.M{"from": primitive.Regex{Pattern: filter.Regex.From, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLETTERUPLOAD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var letteruploads []models.RefLetterUpload
	if err = cursor.All(context.TODO(), &letteruploads); err != nil {
		return nil, err
	}
	return letteruploads, nil
}
