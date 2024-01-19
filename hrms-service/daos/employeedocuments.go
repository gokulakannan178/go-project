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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveEmployeeDocuments :""
func (d *Daos) SaveEmployeeDocuments(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).InsertOne(ctx.CTX, employeeDocuments)
	if err != nil {
		return err
	}
	employeeDocuments.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleEmployeeDocuments : ""
func (d *Daos) GetSingleEmployeeDocuments(ctx *models.Context, uniqueID string) (*models.RefEmployeeDocuments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeID", "uniqueId", "ref.employee", "ref.employee")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTPOLICY, "documentPolicyID", "uniqueId", "ref.documentPolicy", "ref.documentPolicy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTMASTER, "documentMasterID", "uniqueId", "ref.documentMaster", "ref.documentMaster")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.RefEmployeeDocuments
	var employeeDocuments *models.RefEmployeeDocuments
	if err = cursor.All(ctx.CTX, &employeeDocumentss); err != nil {
		return nil, err
	}
	if len(employeeDocumentss) > 0 {
		employeeDocuments = &employeeDocumentss[0]
	}
	return employeeDocuments, nil
}

//UpdateEmployeeDocuments : ""
func (d *Daos) UpdateEmployeeDocuments(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	selector := bson.M{"uniqueId": employeeDocuments.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeDocuments}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableEmployeeDocuments :""
func (d *Daos) EnableEmployeeDocuments(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDOCUMENTSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEmployeeDocuments :""
func (d *Daos) DisableEmployeeDocuments(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDOCUMENTSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEmployeeDocuments :""
func (d *Daos) DeleteEmployeeDocuments(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEDOCUMENTSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterEmployeeDocuments : ""
func (d *Daos) FilterEmployeeDocuments(ctx *models.Context, filter *models.FilterEmployeeDocuments, pagination *models.Pagination) ([]models.RefEmployeeDocuments, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationID}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeDocumentsCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeID", "uniqueId", "ref.employee", "ref.employee")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTPOLICY, "documentPolicyID", "uniqueId", "ref.documentPolicy", "ref.documentPolicy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTMASTER, "documentMasterID", "uniqueId", "ref.documentMaster", "ref.documentMaster")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeDocuments query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.RefEmployeeDocuments
	if err = cursor.All(context.TODO(), &employeeDocumentss); err != nil {
		return nil, err
	}
	return employeeDocumentss, nil
}
func (d *Daos) GetSingleEmployeeDocumentsWithDays(ctx *models.Context, uniqueID int64) (*models.RefEmployeeDocuments, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeDocuments": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeDocumentsCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.RefEmployeeDocuments
	var employeeDocuments *models.RefEmployeeDocuments
	if err = cursor.All(ctx.CTX, &employeeDocumentss); err != nil {
		return nil, err
	}
	if len(employeeDocumentss) > 0 {
		employeeDocuments = &employeeDocumentss[0]
	}
	return employeeDocuments, nil
}
func (d *Daos) EmployeeDocumentsList(ctx *models.Context, filter *models.FilterEmployeeDocumentslist) (*models.EmployeeDocumentsList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeID) > 0 {
			query = append(query, bson.M{"uniqueId": filter.EmployeeID})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
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
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONDOCUMENTPOLICYDOCUMENTS,
			"as":   "data",
			"let":  bson.M{"documentPolicyID": "$documentPolicyID", "employeeid": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$documentPolicyID", "$$documentPolicyID"}},
					{"$eq": []string{"$status", constants.DOCUMENTMASTERSTATUSACTIVE}},
				}}}},
				{
					"$lookup": bson.M{
						"from": constants.COLLECTIONDOCUMENTMASTER,
						"as":   "document",
						"let":  bson.M{"documentMasterID": "$documentMasterID", "documentPolicyID": "$$documentPolicyID"},
						"pipeline": []bson.M{
							{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
								{"$eq": []string{"$uniqueId", "$$documentMasterID"}},
								{"$eq": []string{"$status", constants.DOCUMENTMASTERSTATUSACTIVE}},
								//  {"$eq":["$employeeID","$$employeeid"]},
							}}}},
							{"$project": bson.M{"name": 1, "uniqueId": 1, "file": 1}},
							{
								"$lookup": bson.M{
									"from": constants.COLLECTIONEMPLOYEEDOCUMENTS,
									"as":   "file",
									"let":  bson.M{"documentMasterID": "$uniqueId", "documentPolicyID": "$$documentPolicyID"},
									"pipeline": []bson.M{
										{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
											{"$eq": []string{"$documentMasterID", "$$documentMasterID"}},
											{"$eq": []string{"$employeeID", "$$employeeid"}},
											{"$eq": []string{"$documentPolicyID", "$$documentPolicyID"}},
											{"$eq": []string{"$status", constants.DOCUMENTMASTERSTATUSACTIVE}},
										}}}},
										//	{"$project": bson.M{"name": 1, "uniqueId": 1, "file": 1}},
									},
								},
							},
							{"$addFields": bson.M{"file": bson.M{"$arrayElemAt": []interface{}{"$file.uri", 0}}}},
						},
					},
				},
				{"$addFields": bson.M{"document": bson.M{"$arrayElemAt": []interface{}{"$document", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data": "$data.document"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employeeID": "$uniqueId"}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeDocumentsCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeDocuments  =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.EmployeeDocumentsList
	var employeeDocuments *models.EmployeeDocumentsList
	if err = cursor.All(context.TODO(), &employeeDocumentss); err != nil {
		return nil, err
	}
	if len(employeeDocumentss) > 0 {
		employeeDocuments = &employeeDocumentss[0]
	}
	return employeeDocuments, nil
}
func (d *Daos) UpdateEmployeeDocumentsWithUpsert(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"employeeID": employeeDocuments.EmployeeID, "documentMasterID": employeeDocuments.DocumentMasterID, "documentPolicyID": employeeDocuments.DocumentPolicyID}
	updateData := bson.M{"$set": employeeDocuments}
	if _, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}
func (d *Daos) RemoveEmployeeDocuments(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	opts := options.Update().SetUpsert(true)
	employeeDocuments.Uri = ""
	updateQuery := bson.M{"employeeID": employeeDocuments.EmployeeID, "documentMasterID": employeeDocuments.DocumentMasterID, "documentPolicyID": employeeDocuments.DocumentPolicyID}
	updateData := bson.M{"$set": bson.M{"uri": ""}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEDOCUMENTS).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}
