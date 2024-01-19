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

//SaveMunicipalType :""
func (d *Daos) SaveMunicipalType(ctx *models.Context, municipalType *models.MunicipalType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).InsertOne(ctx.CTX, municipalType)
	return err
}

//GetSingleMunicipalType : ""
func (d *Daos) GetSingleMunicipalType(ctx *models.Context, UniqueID string) (*models.RefMunicipalType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var municipalTypes []models.RefMunicipalType
	var municipalType *models.RefMunicipalType
	if err = cursor.All(ctx.CTX, &municipalTypes); err != nil {
		return nil, err
	}
	if len(municipalTypes) > 0 {
		municipalType = &municipalTypes[0]
	}
	return municipalType, nil
}

//UpdateMunicipalType : ""
func (d *Daos) UpdateMunicipalType(ctx *models.Context, municipalType *models.MunicipalType) error {
	selector := bson.M{"uniqueId": municipalType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": municipalType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterMunicipalType : ""
func (d *Daos) FilterMunicipalType(ctx *models.Context, municipalTypefilter *models.MunicipalTypeFilter, pagination *models.Pagination) ([]models.RefMunicipalType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if municipalTypefilter != nil {

		if len(municipalTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": municipalTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("municipalType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var municipalTypes []models.RefMunicipalType
	if err = cursor.All(context.TODO(), &municipalTypes); err != nil {
		return nil, err
	}
	return municipalTypes, nil
}

//EnableMunicipalType :""
func (d *Daos) EnableMunicipalType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MUNICIPALTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableMunicipalType :""
func (d *Daos) DisableMunicipalType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MUNICIPALTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteMunicipalType :""
func (d *Daos) DeleteMunicipalType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MUNICIPALTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSelectableMunicipalType : ""
func (d *Daos) GetSelectableMunicipalType(ctx *models.Context) (*models.MunicipalType, error) {
	mt := new(models.MunicipalType)
	query2 := bson.M{"isSelectable": true}
	err := ctx.DB.Collection(constants.COLLECTIONMUNICIPALTYPES).FindOne(ctx.CTX, query2).Decode(&mt)
	return mt, err
}
