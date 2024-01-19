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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveGrade :""
func (d *Daos) SaveGrade(ctx *models.Context, grade *models.Grade) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONGRADE).InsertOne(ctx.CTX, grade)
	if err != nil {
		return err
	}
	grade.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleGrade : ""
func (d *Daos) GetSingleGrade(ctx *models.Context, uniqueID string) (*models.RefGrade, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONGRADE, "uniqueId", "GradeID", "ref.gradeAssetsId", "ref.gradeAssetsId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRADE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var grades []models.RefGrade
	var grade *models.RefGrade
	if err = cursor.All(ctx.CTX, &grades); err != nil {
		return nil, err
	}
	if len(grades) > 0 {
		grade = &grades[0]
	}
	return grade, nil
}

//UpdateGrade : ""
func (d *Daos) UpdateGrade(ctx *models.Context, grade *models.Grade) error {
	selector := bson.M{"uniqueId": grade.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": grade}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRADE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableGrade :""
func (d *Daos) EnableGrade(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GRADESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRADE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableGrade :""
func (d *Daos) DisableGrade(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GRADESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRADE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteGrade :""
func (d *Daos) DeleteGrade(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.GRADESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRADE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterGrade : ""
func (d *Daos) FilterGrade(ctx *models.Context, gradeFilter *models.GradeFilter, pagination *models.Pagination) ([]models.RefGrade, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if gradeFilter != nil {

		if len(gradeFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": gradeFilter.Status}})
		}
		if len(gradeFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": gradeFilter.OrganisationID}})
		}
		//Regex
		if gradeFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: gradeFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if gradeFilter != nil {
		if gradeFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{gradeFilter.SortBy: gradeFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONGRADE).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRADE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gradesFilter []models.RefGrade
	if err = cursor.All(context.TODO(), &gradeFilter); err != nil {
		return nil, err
	}
	return gradesFilter, nil
}
