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
)

// SaveLetterUpload : ""
func (d *Daos) SaveLetterGenerate(ctx *models.Context, lg *models.LetterGenerate) error {
	d.Shared.BsonToJSONPrint(lg)
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).InsertOne(ctx.CTX, lg)
	return err
}

// GetSingleLetterGenerate : ""
func (d *Daos) GetSingleLetterGenerate(ctx *models.Context, UniqueID string) (*models.RefLetterGenerate, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("letter generate getsingle query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var lgs []models.RefLetterGenerate
	var lg *models.RefLetterGenerate
	if err = cursor.All(ctx.CTX, &lgs); err != nil {
		return nil, err
	}
	if len(lgs) > 0 {
		lg = &lgs[0]
	}
	return lg, nil
}

// UpdateLetterGenerate : ""
func (d *Daos) UpdateLetterGenerate(ctx *models.Context, lg *models.LetterGenerate) error {
	selector := bson.M{"uniqueId": lg.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": lg, "$push": bson.M{"update": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableLetterUpload : ""
func (d *Daos) EnableLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {

	query := bson.M{"uniqueId": lg.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERGENERATESTATUSACTIVE,
		"activated.action": lg.Action,
		"activated.on":     lg.On,
		"activated.by":     lg.By,
		"activated.bytype": lg.ByType}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) ApprovedLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {

	query := bson.M{"uniqueId": lg.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERGENERATESTATUSAPPROVED,
		"activated.action": lg.Action,
		"activated.on":     lg.On,
		"activated.by":     lg.By,
		"activated.bytype": lg.ByType}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableLetterGenerate : ""
func (d *Daos) DisableLetterGenerate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERGENERATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteLetterGenerate : ""
func (d *Daos) DeleteLetterGenerate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERGENERATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// BlockedLetterGenerate : ""
func (d *Daos) BlockedLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {
	query := bson.M{"uniqueId": lg.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERGENERATESTATUSBLOCKED,
		"blocked.action": lg.Action,
		"blocked.on":     lg.On,
		"blocked.by":     lg.By,
		"blocked.bytype": lg.ByType}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// SubmittedLetterGenerate: ""
func (d *Daos) SubmittedLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {
	query := bson.M{"uniqueId": lg.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LETTERGENERATESTATUSSUBMITTED,
		"submitted.action": lg.Action,
		"submitted.on":     lg.On,
		"submitted.by":     lg.By,
		"submitted.bytype": lg.ByType}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) UploadLetterGenerate(ctx *models.Context, lg *models.LetterGenerate) error {
	query := bson.M{"uniqueId": lg.UniqueID}
	update := bson.M{"$set": lg}
	_, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLetterGenerate : ""
func (d *Daos) FilterLetterGenerate(ctx *models.Context, filter *models.LetterGenerateFilter, pagination *models.Pagination) ([]models.RefLetterGenerate, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("letter generate filter query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLETTERGENERATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var letterGenerates []models.RefLetterGenerate
	if err = cursor.All(context.TODO(), &letterGenerates); err != nil {
		return nil, err
	}
	return letterGenerates, nil
}
