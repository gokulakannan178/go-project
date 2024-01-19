package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveRole : ""
func (d *Daos) SaveRole(ctx *models.Context, role *models.Role) error {
	d.Shared.BsonToJSONPrint(role)
	_, err := ctx.DB.Collection(constants.COLLECTIONROLE).InsertOne(ctx.CTX, role)
	return err
}

// GetSingleRole : ""
func (d *Daos) GetSingleRole(ctx *models.Context, UniqueID string) (*models.RefRole, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("getsinglerole query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROLE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefRole
	var tower *models.RefRole
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateRole : ""
func (d *Daos) UpdateRole(ctx *models.Context, business *models.Role) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONROLE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableRole : ""
func (d *Daos) EnableRole(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROLESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableRole : ""
func (d *Daos) DisableRole(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROLESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteRole : ""
func (d *Daos) DeleteRole(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROLESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROLE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterRole : ""
func (d *Daos) FilterRole(ctx *models.Context, filter *models.RoleFilter, pagination *models.Pagination) ([]models.RefRole, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.CompanyID) > 0 {
			query = append(query, bson.M{"company.id": bson.M{"$in": filter.CompanyID}})
		}
		if len(filter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": filter.Name}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONROLE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("filterrole query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROLE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var role []models.RefRole
	if err = cursor.All(context.TODO(), &role); err != nil {
		return nil, err
	}
	return role, nil
}
