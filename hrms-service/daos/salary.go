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

//SaveSalary : ""
func (d *Daos) SaveSalary(ctx *models.Context, salary *models.Salary) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSALARY).InsertOne(ctx.CTX, salary)
	if err != nil {
		return err
	}
	salary.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateSalary : ""
func (d *Daos) UpdateSalary(ctx *models.Context, salary *models.Salary) error {
	selector := bson.M{"uniqueId": salary.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": salary}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleSalary : ""
func (d *Daos) GetSingleSalary(ctx *models.Context, uniqueID string) (*models.RefSalary, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Salarys []models.RefSalary
	var Salary *models.RefSalary
	if err = cursor.All(ctx.CTX, &Salarys); err != nil {
		return nil, err
	}
	if len(Salarys) > 0 {
		Salary = &Salarys[0]
	}
	return Salary, err
}

// EnableSalary : ""
func (d *Daos) EnableSalary(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SALARYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableSalary : ""
func (d *Daos) DisableSalary(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SALARYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteSalary :""
func (d *Daos) DeleteSalary(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SALARYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSalary : ""
func (d *Daos) FilterSalary(ctx *models.Context, Salary *models.FilterSalary, pagination *models.Pagination) ([]models.RefSalary, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Salary != nil {
		if len(Salary.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Salary.Status}})
		}
		if len(Salary.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": Salary.OrganisationId}})
		}
		//Regex
		if Salary.Regex.Name != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: Salary.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSALARY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var salaryFilter []models.RefSalary
	if err = cursor.All(context.TODO(), &salaryFilter); err != nil {
		return nil, err
	}
	return salaryFilter, nil
}
