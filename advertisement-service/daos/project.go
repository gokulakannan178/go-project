package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveProject :""
func (d *Daos) SaveProject(ctx *models.Context, project *models.Project) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).InsertOne(ctx.CTX, project)
	return err
}

//GetSingleProject : ""
func (d *Daos) GetSingleProject(ctx *models.Context, UniqueID string) (*models.RefProject, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var projects []models.RefProject
	var project *models.RefProject
	if err = cursor.All(ctx.CTX, &projects); err != nil {
		return nil, err
	}
	if len(projects) > 0 {
		project = &projects[0]
	}
	return project, nil
}

//UpdateProject : ""
func (d *Daos) UpdateProject(ctx *models.Context, project *models.Project) error {
	selector := bson.M{"uniqueId": project.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": project}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterProject : ""
func (d *Daos) FilterProject(ctx *models.Context, projectfilter *models.ProjectFilter, pagination *models.Pagination) ([]models.RefProject, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if projectfilter != nil {

		if len(projectfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": projectfilter.Status}})
		}
		//Regex
		if projectfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: projectfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var projects []models.RefProject
	if err = cursor.All(context.TODO(), &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

//EnableProject :""
func (d *Daos) EnableProject(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTOWNERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProject :""
func (d *Daos) DisableProject(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTOWNERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProject :""
func (d *Daos) DeleteProject(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTOWNERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
