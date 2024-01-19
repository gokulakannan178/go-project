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

// SaveDepartment : ""
func (d *Daos) SaveDepartment(ctx *models.Context, department *models.Department) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).InsertOne(ctx.CTX, department)
	if err != nil {
		return err
	}
	department.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateState : ""
func (d *Daos) UpdateDepartment(ctx *models.Context, dept *models.Department) error {
	selector := bson.M{"uniqueId": dept.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dept, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleDepartment : ""
func (d *Daos) GetSingleDepartment(ctx *models.Context, uniqueID string) (*models.RefDepartment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "branch", "uniqueId", "ref.branch", "ref.branch")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var departments []models.RefDepartment
	var department *models.RefDepartment
	if err = cursor.All(ctx.CTX, &departments); err != nil {
		return nil, err
	}
	if len(departments) > 0 {
		department = &departments[0]
	}
	return department, err
}

// EnableDepartment : ""
func (d *Daos) EnableDepartment(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.DEPARTMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableDepartment : ""
func (d *Daos) DisableDepartment(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.DEPARTMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteDepartment :""
func (d *Daos) DeleteDepartment(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEPARTMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDepartment : ""
func (d *Daos) FilterDepartment(ctx *models.Context, filter *models.FilterDepartment, pagination *models.Pagination) ([]models.RefDepartment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		if len(filter.BranchId) > 0 {
			query = append(query, bson.M{"branch": bson.M{"$in": filter.BranchId}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}

		if filter.Regex.Type != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: filter.Regex.Type, Options: "xi"}})
		}
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "branch", "uniqueId", "ref.branch", "ref.branch")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var departmentFilter []models.RefDepartment
	if err = cursor.All(context.TODO(), &departmentFilter); err != nil {
		return nil, err
	}
	return departmentFilter, nil
}

// GetSingleDepartmentActivewithName : ""
func (d *Daos) GetSingleDepartmentActivewithName(ctx *models.Context, uniqueID string) (*models.RefDepartment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.ONBOARDINGPOLICYSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "branch", "uniqueId", "ref.branch", "ref.branch")...)
	d.Shared.BsonToJSONPrintTag("Dept query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var departments []models.RefDepartment
	var department *models.RefDepartment
	if err = cursor.All(ctx.CTX, &departments); err != nil {
		return nil, err
	}
	if len(departments) > 0 {
		department = &departments[0]
	}
	return department, err
}
func (d *Daos) GetSingleDepartmentUniqueCheck(ctx *models.Context, uniqueID string, org string) (*models.RefDepartment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"organisationId": org, "name": primitive.Regex{Pattern: uniqueID, Options: "xi"}, "status": constants.ORGANISATIONCONFIGSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "branch", "uniqueId", "ref.branch", "ref.branch")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var departments []models.RefDepartment
	var department *models.RefDepartment
	if err = cursor.All(ctx.CTX, &departments); err != nil {
		return nil, err
	}
	if len(departments) > 0 {
		department = &departments[0]
	}
	return department, err
}
