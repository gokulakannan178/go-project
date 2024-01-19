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

//SaveEmployeeEducation : ""
func (d *Daos) SaveEmployeeEducation(ctx *models.Context, employeeEducation *models.EmployeeEducation) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).InsertOne(ctx.CTX, employeeEducation)
	if err != nil {
		return err
	}
	employeeEducation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeEducation : ""
func (d *Daos) UpdateEmployeeEducation(ctx *models.Context, EmployeeEducation *models.EmployeeEducation) error {
	selector := bson.M{"uniqueId": EmployeeEducation.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": EmployeeEducation}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeEducation : ""
func (d *Daos) GetSingleEmployeeEducation(ctx *models.Context, uniqueID string) (*models.RefEmployeeEducation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeEducations []models.RefEmployeeEducation
	var EmployeeEducation *models.RefEmployeeEducation
	if err = cursor.All(ctx.CTX, &EmployeeEducations); err != nil {
		return nil, err
	}
	if len(EmployeeEducations) > 0 {
		EmployeeEducation = &EmployeeEducations[0]
	}
	return EmployeeEducation, err
}

// EnableEmployeeEducation : ""
func (d *Daos) EnableEmployeeEducation(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEDUCATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeEducation : ""
func (d *Daos) DisableEmployeeEducation(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEDUCATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeEducation :""
func (d *Daos) DeleteEmployeeEducation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEDUCATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeEducation : ""
func (d *Daos) FilterEmployeeEducation(ctx *models.Context, employeeEducation *models.FilterEmployeeEducation, pagination *models.Pagination) ([]models.RefEmployeeEducation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeEducation != nil {
		if len(employeeEducation.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeEducation.Status}})
		}
		if len(employeeEducation.EmployeeID) > 0 {
			query = append(query, bson.M{"employeeID": bson.M{"$in": employeeEducation.EmployeeID}})
		}
		if len(employeeEducation.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeEducation.OrganisationID}})
		}
		//Regex
		if employeeEducation.Regex.InstituteName != "" {
			query = append(query, bson.M{"instituteName": primitive.Regex{Pattern: employeeEducation.Regex.InstituteName, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeeEducation != nil {
		if employeeEducation.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeEducation.SortBy: employeeEducation.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).CountDocuments(ctx.CTX, func() bson.M {
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
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEDUCATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeEducationFilter []models.RefEmployeeEducation
	if err = cursor.All(context.TODO(), &employeeEducationFilter); err != nil {
		return nil, err
	}
	return employeeEducationFilter, nil
}
