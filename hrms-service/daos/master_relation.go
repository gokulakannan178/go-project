package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveRelation :""
func (d *Daos) SaveRelation(ctx *models.Context, relation *models.Relation) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONRELATION).InsertOne(ctx.CTX, relation)
	return err
}

//GetSingleRelation : ""
func (d *Daos) GetSingleRelation(ctx *models.Context, UniqueID string) (*models.RefRelation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONRELATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var relations []models.RefRelation
	var relation *models.RefRelation
	if err = cursor.All(ctx.CTX, &relations); err != nil {
		return nil, err
	}
	if len(relations) > 0 {
		relation = &relations[0]
	}
	return relation, nil
}

//UpdateRelation : ""
func (d *Daos) UpdateRelation(ctx *models.Context, relation *models.Relation) error {
	selector := bson.M{"uniqueId": relation.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": relation, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRELATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterRelation : ""
func (d *Daos) FilterRelation(ctx *models.Context, relationfilter *models.RelationFilter, pagination *models.Pagination) ([]models.RefRelation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if relationfilter != nil {

		if len(relationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": relationfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONRELATION).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("relation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONRELATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var relations []models.RefRelation
	if err = cursor.All(context.TODO(), &relations); err != nil {
		return nil, err
	}
	return relations, nil
}

//EnableRelation :""
func (d *Daos) EnableRelation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.RELATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRELATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableRelation :""
func (d *Daos) DisableRelation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.RELATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRELATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteRelation :""
func (d *Daos) DeleteRelation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.RELATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONRELATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
