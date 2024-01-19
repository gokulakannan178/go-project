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

// SaveMobileTowerRegistrationTax : ""
func (d *Daos) SaveMobileTowerRegistrationTax(ctx *models.Context, mtrt *models.MobileTowerRegistrationTax) error {
	d.Shared.BsonToJSONPrint(mtrt)
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).InsertOne(ctx.CTX, mtrt)
	return err
}

// GetSingleMobileTowerRegistrationTax : ""
func (d *Daos) GetSingleMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) (*models.RefMobileTowerRegistrationTax, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefMobileTowerRegistrationTax
	var tower *models.RefMobileTowerRegistrationTax
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// GetSingleDefaultMobileTowerRegistrationTax : ""
func (d *Daos) GetSingleDefaultMobileTowerRegistrationTax(ctx *models.Context) (*models.RefMobileTowerRegistrationTax, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefMobileTowerRegistrationTax
	var tower *models.RefMobileTowerRegistrationTax
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateMobileTowerRegistrationTax : ""
func (d *Daos) UpdateMobileTowerRegistrationTax(ctx *models.Context, business *models.MobileTowerRegistrationTax) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMobileTowerRegistrationTax : ""
func (d *Daos) EnableMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERREGISTRATIONTAXSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableMobileTowerRegistrationTax : ""
func (d *Daos) DisableMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERREGISTRATIONTAXSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteMobileTowerRegistrationTax : ""
func (d *Daos) DeleteMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERREGISTRATIONTAXSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMobileTowerRegistrationTax : ""
func (d *Daos) FilterMobileTowerRegistrationTax(ctx *models.Context, filter *models.MobileTowerRegistrationTaxFilter, pagination *models.Pagination) ([]models.RefMobileTowerRegistrationTax, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var mtrt []models.RefMobileTowerRegistrationTax
	if err = cursor.All(context.TODO(), &mtrt); err != nil {
		return nil, err
	}
	return mtrt, nil
}
