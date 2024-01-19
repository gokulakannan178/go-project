package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveDisease :""
func (d *Daos) SaveOnePageAdvisoryTemplate(ctx *models.Context, OnePageAdvisoryTemplate *models.OnePageAdvisoryTemplate) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).InsertOne(ctx.CTX, OnePageAdvisoryTemplate)
	if err != nil {
		return err
	}
	OnePageAdvisoryTemplate.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) (*models.RefOnePageAdvisoryTemplate, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OnePageAdvisoryTemplates []models.RefOnePageAdvisoryTemplate
	var OnePageAdvisoryTemplate *models.RefOnePageAdvisoryTemplate
	if err = cursor.All(ctx.CTX, &OnePageAdvisoryTemplates); err != nil {
		return nil, err
	}
	if len(OnePageAdvisoryTemplates) > 0 {
		OnePageAdvisoryTemplate = &OnePageAdvisoryTemplates[0]
	}
	return OnePageAdvisoryTemplate, nil
}

//UpdateOnePageAdvisoryTemplate : ""
func (d *Daos) UpdateOnePageAdvisoryTemplate(ctx *models.Context, OnePageAdvisoryTemplate *models.OnePageAdvisoryTemplate) error {

	selector := bson.M{"_id": OnePageAdvisoryTemplate.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": OnePageAdvisoryTemplate}
	_, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOnePageAdvisoryTemplate : ""
func (d *Daos) FilterOnePageAdvisoryTemplate(ctx *models.Context, OnePageAdvisoryTemplatefilter *models.OnePageAdvisoryTemplateFilter, pagination *models.Pagination) ([]models.RefOnePageAdvisoryTemplate, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if OnePageAdvisoryTemplatefilter != nil {

		if len(OnePageAdvisoryTemplatefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": OnePageAdvisoryTemplatefilter.ActiveStatus}})
		}
		if len(OnePageAdvisoryTemplatefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": OnePageAdvisoryTemplatefilter.Status}})
		}
		//Regex
		if OnePageAdvisoryTemplatefilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: OnePageAdvisoryTemplatefilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if OnePageAdvisoryTemplatefilter != nil {
		if OnePageAdvisoryTemplatefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{OnePageAdvisoryTemplatefilter.SortBy: OnePageAdvisoryTemplatefilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("OnePageAdvisoryTemplate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var OnePageAdvisoryTemplates []models.RefOnePageAdvisoryTemplate
	if err = cursor.All(context.TODO(), &OnePageAdvisoryTemplates); err != nil {
		return nil, err
	}
	return OnePageAdvisoryTemplates, nil
}

//EnableOnePageAdvisoryTemplate :""
func (d *Daos) EnableOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ONEPAGEADVISORYTEMPLATESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ONEPAGEADVISORYTEMPLATESTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOnePageAdvisoryTemplate :""
func (d *Daos) DeleteOnePageAdvisoryTemplate(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ONEPAGEADVISORYTEMPLATESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONONEPAGEADVISORYTEMPLATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
