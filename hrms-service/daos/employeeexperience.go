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

//SaveEmployeeExperience : ""
func (d *Daos) SaveEmployeeExperience(ctx *models.Context, employeeExperience *models.EmployeeExperience) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).InsertOne(ctx.CTX, employeeExperience)
	if err != nil {
		return err
	}
	employeeExperience.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeExperience : ""
func (d *Daos) UpdateEmployeeExperience(ctx *models.Context, employeeExperience *models.EmployeeExperience) error {
	selector := bson.M{"uniqueId": employeeExperience.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeExperience}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeExperience : ""
func (d *Daos) GetSingleEmployeeExperience(ctx *models.Context, uniqueID string) (*models.RefEmployeeExperience, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeExperiences []models.RefEmployeeExperience
	var EmployeeExperience *models.RefEmployeeExperience
	if err = cursor.All(ctx.CTX, &EmployeeExperiences); err != nil {
		return nil, err
	}
	if len(EmployeeExperiences) > 0 {
		EmployeeExperience = &EmployeeExperiences[0]
	}
	return EmployeeExperience, err
}

// EnableEmployeeExperience : ""
func (d *Daos) EnableEmployeeExperience(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEXPERIENCESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeExperience : ""
func (d *Daos) DisableEmployeeExperience(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEXPERIENCESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeExperience :""
func (d *Daos) DeleteEmployeeExperience(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEEXPERIENCESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeExperience : ""
func (d *Daos) FilterEmployeeExperience(ctx *models.Context, employeeExperience *models.FilterEmployeeExperience, pagination *models.Pagination) ([]models.RefEmployeeExperience, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeExperience != nil {
		if len(employeeExperience.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeExperience.Status}})
		}
		if len(employeeExperience.EmployeeID) > 0 {
			query = append(query, bson.M{"employeeID": bson.M{"$in": employeeExperience.EmployeeID}})
		}
		if len(employeeExperience.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeExperience.OrganisationID}})
		}
		//Regex
		if employeeExperience.Regex.CompanyName != "" {
			query = append(query, bson.M{"companyName": primitive.Regex{Pattern: employeeExperience.Regex.CompanyName, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeeExperience != nil {
		if employeeExperience.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeExperience.SortBy: employeeExperience.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEEXPERIENCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeExperienceFilter []models.RefEmployeeExperience
	if err = cursor.All(context.TODO(), &employeeExperienceFilter); err != nil {
		return nil, err
	}
	return employeeExperienceFilter, nil
}
