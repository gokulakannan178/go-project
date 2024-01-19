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

//SaveEmployeeFamilyMembers : ""
func (d *Daos) SaveEmployeeFamilyMembers(ctx *models.Context, employeeFamilyMembers *models.EmployeeFamilyMembers) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).InsertOne(ctx.CTX, employeeFamilyMembers)
	if err != nil {
		return err
	}
	employeeFamilyMembers.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeFamilyMembers : ""
func (d *Daos) UpdateEmployeeFamilyMembers(ctx *models.Context, employeeFamilyMembers *models.EmployeeFamilyMembers) error {
	selector := bson.M{"uniqueId": employeeFamilyMembers.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeFamilyMembers}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeFamilyMembers : ""
func (d *Daos) GetSingleEmployeeFamilyMembers(ctx *models.Context, uniqueID string) (*models.RefEmployeeFamilyMembers, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeFamilyMemberss []models.RefEmployeeFamilyMembers
	var EmployeeFamilyMembers *models.RefEmployeeFamilyMembers
	if err = cursor.All(ctx.CTX, &EmployeeFamilyMemberss); err != nil {
		return nil, err
	}
	if len(EmployeeFamilyMemberss) > 0 {
		EmployeeFamilyMembers = &EmployeeFamilyMemberss[0]
	}
	return EmployeeFamilyMembers, err
}

// EnableEmployeeFamilyMembers : ""
func (d *Daos) EnableEmployeeFamilyMembers(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEFAMILYMEMBERSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeFamilyMembers : ""
func (d *Daos) DisableEmployeeFamilyMembers(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEFAMILYMEMBERSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeFamilyMembers :""
func (d *Daos) DeleteEmployeeFamilyMembers(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEFAMILYMEMBERSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeFamilyMembers : ""
func (d *Daos) FilterEmployeeFamilyMembers(ctx *models.Context, employeeFamilyMembers *models.FilterEmployeeFamilyMembers, pagination *models.Pagination) ([]models.RefEmployeeFamilyMembers, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeFamilyMembers != nil {
		if len(employeeFamilyMembers.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeFamilyMembers.Status}})
		}
		if len(employeeFamilyMembers.EmployeeID) > 0 {
			query = append(query, bson.M{"employeeID": bson.M{"$in": employeeFamilyMembers.EmployeeID}})
		}
		if len(employeeFamilyMembers.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeFamilyMembers.OrganisationID}})
		}
		//Regex
		if employeeFamilyMembers.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeeFamilyMembers.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if employeeFamilyMembers != nil {
		if employeeFamilyMembers.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeFamilyMembers.SortBy: employeeFamilyMembers.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEFAMILYMEMBERS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeFamilyMembersFilter []models.RefEmployeeFamilyMembers
	if err = cursor.All(context.TODO(), &employeeFamilyMembersFilter); err != nil {
		return nil, err
	}
	return employeeFamilyMembersFilter, nil
}
