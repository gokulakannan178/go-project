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

// SaveEmployeeLeave : ""
func (d *Daos) SaveEmployeeLeave(ctx *models.Context, employeeLeave *models.EmployeeLeave) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).InsertOne(ctx.CTX, employeeLeave)
	if err != nil {
		return err
	}
	//employeeLeave.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeLeave : ""
func (d *Daos) UpdateEmployeeLeave(ctx *models.Context, employeeLeave *models.EmployeeLeave) error {
	selector := bson.M{"uniqueId": employeeLeave.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeLeave}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeLeave : ""
func (d *Daos) GetSingleEmployeeLeave(ctx *models.Context, uniqueID string) (*models.RefEmployeeLeave, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeLeaves []models.RefEmployeeLeave
	var EmployeeLeave *models.RefEmployeeLeave
	if err = cursor.All(ctx.CTX, &EmployeeLeaves); err != nil {
		return nil, err
	}
	if len(EmployeeLeaves) > 0 {
		EmployeeLeave = &EmployeeLeaves[0]
	}
	return EmployeeLeave, err
}

// EnableEmployeeLeave : ""
func (d *Daos) EnableEmployeeLeave(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEELEAVESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeLeave : ""
func (d *Daos) DisableEmployeeLeave(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEELEAVESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeLeave :""
func (d *Daos) DeleteEmployeeLeave(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEELEAVESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeLeave : ""
func (d *Daos) FilterEmployeeLeave(ctx *models.Context, employeeLeave *models.FilterEmployeeLeave, pagination *models.Pagination) ([]models.RefEmployeeLeave, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeLeave != nil {
		if len(employeeLeave.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeLeave.Status}})
		}
		if len(employeeLeave.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeLeave.OrganisationId}})
		}
		if len(employeeLeave.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeLeave.EmployeeId}})
		}
		if len(employeeLeave.LeaveType) > 0 {
			query = append(query, bson.M{"leaveType": bson.M{"$in": employeeLeave.LeaveType}})
		}
		//Regex
		if employeeLeave.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeeLeave.Regex.Name, Options: "xi"}})
		}
		if employeeLeave.Regex.LeaveType != "" {
			query = append(query, bson.M{"leaveType": primitive.Regex{Pattern: employeeLeave.Regex.LeaveType, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeeLeave != nil {
		if employeeLeave.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeLeave.SortBy: employeeLeave.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rf []models.RefEmployeeLeave
	if err = cursor.All(context.TODO(), &rf); err != nil {
		return nil, err
	}
	return rf, nil
}

// GetEmployeeLeaveCount : ""
func (d *Daos) GetEmployeeLeaveCount(ctx *models.Context, employeeleavecount *models.EmployeeLeaveCount) ([]models.RefEmployeeLeaveCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$facet": bson.M{
			"totalLeave": []bson.M{
				bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"employeeId": employeeleavecount.EmployeeId}, bson.M{"leaveType": employeeleavecount.LeaveType}, bson.M{"organisationId": employeeleavecount.OrganisationId}}}},
				bson.M{"$count": "totalLeave"}},
		},
	},

		bson.M{"$addFields": bson.M{"totalLeave": bson.M{"$arrayElemAt": []interface{}{"$totalLeave", 0}}}},
		bson.M{"$addFields": bson.M{"totalLeave": "$totalLeave.totalLeave"}})

	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rf []models.RefEmployeeLeaveCount
	if err = cursor.All(context.TODO(), &rf); err != nil {
		return nil, err
	}
	return rf, nil
}
func (d *Daos) UpdateEmployeeLeaveFromTimeOff(ctx *models.Context, employeeLeave *models.UpdateEmployeeLeave) error {
	selector := bson.M{"employeeId": employeeLeave.EmployeeId, "leaveType": employeeLeave.LeaveType}
	data := bson.M{"$inc": bson.M{"value": -employeeLeave.Value}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).UpdateOne(ctx.CTX, selector, data)
	return err
}
func (d *Daos) RevertEmployeeLeaveFromTimeOff(ctx *models.Context, employeeLeave *models.UpdateEmployeeLeave) error {
	selector := bson.M{"employeeId": employeeLeave.EmployeeId, "leaveType": employeeLeave.LeaveType}
	data := bson.M{"$inc": bson.M{"value": employeeLeave.Value}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).UpdateOne(ctx.CTX, selector, data)
	return err
}
func (d *Daos) GetSingleEmployeeLeaveWithEmployeeId(ctx *models.Context, uniqueID string) (*models.RefEmployeeLeave, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": uniqueID}})
	//LookUp
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeLeaves []models.RefEmployeeLeave
	var EmployeeLeave *models.RefEmployeeLeave
	if err = cursor.All(ctx.CTX, &EmployeeLeaves); err != nil {
		return nil, err
	}
	if len(EmployeeLeaves) > 0 {
		EmployeeLeave = &EmployeeLeaves[0]
	}
	if EmployeeLeave == nil {
		return nil, errors.New("Employee leave Not Found")
	}
	return EmployeeLeave, err
}
func (d *Daos) EmployeeLeaveList(ctx *models.Context, filter *models.FilterEmployeeLeave) ([]models.EmployeeLeaveListV2, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.EmployeeId}})
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

	// if filter != nil {
	// 	if filter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	// 	}
	// }
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONLEAVEPOLICY,
			"as":   "employeeLeave",
			"let":  bson.M{"leavePolicyID": "$leavePolicyID", "employeeid": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
					"$eq": []string{"$uniqueId", "$$leavePolicyID"},
					//  {"$eq":["$employeeID","$$employeeid"]},
				}}}},
				{"$unwind": bson.M{"path": "$-"}},
				{
					"$lookup": bson.M{
						"from": constants.COLLECTIONLEAVEMASTER,
						"as":   "leave",
						"let":  bson.M{"leavemasterId": "$-", "leavePolicyID": "$$leavePolicyID"},
						"pipeline": []bson.M{
							{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
								"$eq": []string{"$uniqueId", "$$leavemasterId"},
								//  {"$eq":["$employeeID","$$employeeid"]},
							}}}},
							{"$project": bson.M{"name": 1, "uniqueId": 1, "leaveType": 1, "value": 1}},
							{
								"$lookup": bson.M{
									"from": constants.COLLECTIONEMPLOYEELEAVE,
									"as":   "value",
									"let":  bson.M{"leavemasterId": "$uniqueId", "leavePolicyID": "$$leavePolicyID"},
									"pipeline": []bson.M{
										{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
											{"$eq": []string{"$leaveType", "$$leavemasterId"}},
											{"$eq": []string{"$employeeId", "$$employeeid"}},
										}}}},
										//	{"$project": bson.M{"name": 1, "uniqueId": 1, "file": 1}},
									},
								},
							},
							{"$addFields": bson.M{"value": bson.M{"$arrayElemAt": []interface{}{"$value.value", 0}}}},
						},
					},
				},
				{"$addFields": bson.M{"leave": bson.M{"$arrayElemAt": []interface{}{"$leave", 0}}}},
				bson.M{"$project": bson.M{
					"k":   "$leave.name",
					"v":   "$$ROOT.leave",
					"_id": 0,
				}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{

		"employeeLeave": bson.M{"$arrayToObject": "$employeeLeave"}}})
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$group": bson.M{
	// 		"_id":{""},
	// 		"employeeLeave": bson.M{"$push": "$$ROOT.employeeLeave"},
	// 	},
	// })
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"employee": "$uniqueId"},
	})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeDocumentsCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeDocuments  =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.EmployeeLeaveListV2
	//var employeeDocuments *models.EmployeeLeaveList
	if err = cursor.All(context.TODO(), &employeeDocumentss); err != nil {
		return nil, err
	}
	// if len(employeeDocumentss) < 0 {
	// 	employeeDocuments = &employeeDocumentss[0]
	// }
	return employeeDocumentss, nil
}
func (d *Daos) EmployeeLeaveListV2(ctx *models.Context, filter *models.FilterEmployeeLeaveList) (*models.EmployeeLeaveList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if filter.EmployeeId != "" {
			query = append(query, bson.M{"uniqueId": filter.EmployeeId})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	query2 := bson.M{"$match": bson.M{}}
	if filter.IsZero == "No" {
		query2["$match"] = bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$gt": []interface{}{"$leave.value", 0}}}}}
	}

	// if filter != nil {
	// 	if filter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	// 	}
	// }
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONLEAVEPOLICY,
			"as":   "employeeLeave",
			"let":  bson.M{"leavePolicyID": "$leavePolicyID", "employeeid": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
					"$eq": []string{"$uniqueId", "$$leavePolicyID"},
					//  {"$eq":["$employeeID","$$employeeid"]},
				}}}},

				{"$unwind": bson.M{"path": "$-"}},
				{
					"$lookup": bson.M{
						"from": constants.COLLECTIONLEAVEMASTER,
						"as":   "leave",
						"let":  bson.M{"leavemasterId": "$-", "leavePolicyID": "$$leavePolicyID"},
						"pipeline": []bson.M{
							{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
								"$eq": []string{"$uniqueId", "$$leavemasterId"},
								//  {"$eq":["$employeeID","$$employeeid"]},

							}}}},

							{"$project": bson.M{"name": 1, "uniqueId": 1, "leaveType": 1, "value": 1}},
							{
								"$lookup": bson.M{
									"from": constants.COLLECTIONEMPLOYEELEAVE,
									"as":   "value",
									"let":  bson.M{"leavemasterId": "$uniqueId", "leavePolicyID": "$$leavePolicyID"},
									"pipeline": []bson.M{
										{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
											{"$eq": []string{"$leaveType", "$$leavemasterId"}},
											{"$eq": []string{"$employeeId", "$$employeeid"}},
										}}}},
										//	{"$project": bson.M{"name": 1, "uniqueId": 1, "file": 1}},
									},
								},
							},
							{"$addFields": bson.M{"value": bson.M{"$arrayElemAt": []interface{}{"$value.value", 0}}}},
						},
					},
				},
				{"$addFields": bson.M{"leave": bson.M{"$arrayElemAt": []interface{}{"$leave", 0}}}},
				query2,
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"employeeLeave": "$employeeLeave.leave"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"employee": "$uniqueId"},
	})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeDocumentsCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeLeave  =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.EmployeeLeaveList
	var employeeDocuments *models.EmployeeLeaveList
	if err = cursor.All(context.TODO(), &employeeDocumentss); err != nil {
		return nil, err
	}
	if len(employeeDocumentss) > 0 {
		employeeDocuments = &employeeDocumentss[0]
	}
	return employeeDocuments, nil
}
func (d *Daos) EmployeeLeaveListV3(ctx *models.Context, filter *models.FilterEmployeeLeaveList) (*models.EmployeeLeaveList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if filter.EmployeeId != "" {
			query = append(query, bson.M{"uniqueId": filter.EmployeeId})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	query2 := bson.M{"$match": bson.M{}}
	if filter.IsZero == "No" {
		query2["$match"] = bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$gt": []interface{}{"$leave.value", 0}}}}}
	}

	// if filter != nil {
	// 	if filter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	// 	}
	// }
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONLEAVEPOLICY,
			"as":   "employeeLeave",
			"let":  bson.M{"leavePolicyID": "$leavePolicyID", "employeeid": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
					"$eq": []string{"$uniqueId", "$$leavePolicyID"},
					//  {"$eq":["$employeeID","$$employeeid"]},
				}}}},

				{"$unwind": bson.M{"path": "$leavemaster"}},
				{
					"$lookup": bson.M{
						"from": constants.COLLECTIONLEAVEMASTER,
						"as":   "leave",
						"let":  bson.M{"leavemasterId": "$leavemaster.uniqueId", "leavePolicyID": "$$leavePolicyID"},
						"pipeline": []bson.M{
							{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
								"$eq": []string{"$uniqueId", "$$leavemasterId"},
								//  {"$eq":["$employeeID","$$employeeid"]},
							}}}},

							{"$project": bson.M{"name": 1, "uniqueId": 1, "leaveType": 1, "value": 1}},
							{
								"$lookup": bson.M{
									"from": constants.COLLECTIONEMPLOYEELEAVE,
									"as":   "value",
									"let":  bson.M{"leavemasterId": "$uniqueId", "leavePolicyID": "$$leavePolicyID"},
									"pipeline": []bson.M{
										{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
											{"$eq": []string{"$leaveType", "$$leavemasterId"}},
											{"$eq": []string{"$employeeId", "$$employeeid"}},
										}}}},
										//	{"$project": bson.M{"name": 1, "uniqueId": 1, "file": 1}},
										{
											"$lookup": bson.M{
												"as":   "request",
												"from": "employeetimeoff",
												"let":  bson.M{"leavemasterId": "$$leavemasterId", "employeeId": "$$employeeid"},
												"pipeline": []bson.M{
													{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
														{"$eq": []string{"$leaveType", "$$leavemasterId"}},
														{"$eq": []string{"$employeeId", "$$employeeid"}},
														{"$eq": []string{"$status", "Request"}},
													}}}},
													{"$group": bson.M{
														"_id":          nil,
														"numberOfDays": bson.M{"$sum": "$numberOfDays"}}},
												},
											},
										},

										{"$addFields": bson.M{"numberOfDays": bson.M{"$arrayElemAt": []interface{}{"$request.numberOfDays", 0}}}},
										//{"$inc":{"value":-"numberOfDays"}}
										{"$project": bson.M{"value": 1, "numberOfDays": 1, "_id": 0}},
										//{"$project": bson.M{"value": bson.M{"$subtract": []string{"$value", "$numberOfDays"}}}},
									},
								},
							},
							{"$addFields": bson.M{"value": bson.M{"$arrayElemAt": []interface{}{"$value", 0}}}},
							{"$addFields": bson.M{"numberOfDays": "$value.numberOfDays"}},
							{"$addFields": bson.M{"value": "$value.value"}},
						},
					},
				},
				{"$addFields": bson.M{"leave": bson.M{"$arrayElemAt": []interface{}{"$leave", 0}}}},
				query2,
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"employeeLeave": "$employeeLeave.leave"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"employee": "$uniqueId"},
	})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeDocumentsCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeLeave  =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.EmployeeLeaveList
	var employeeDocuments *models.EmployeeLeaveList
	if err = cursor.All(context.TODO(), &employeeDocumentss); err != nil {
		return nil, err
	}
	if len(employeeDocumentss) > 0 {
		employeeDocuments = &employeeDocumentss[0]
	}
	return employeeDocuments, nil
}
func (d *Daos) GetAllEmployeeLeaveList(ctx *models.Context, filter *models.FilterEmployeeLeaveList) ([]models.EmployeeLeaveList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if filter.EmployeeId != "" {
			query = append(query, bson.M{"uniqueId": filter.EmployeeId})
		}

	}
	query = append(query, bson.M{"status": bson.M{"$nin": []string{constants.EMPLOYEESTATUSDELETED, constants.EMPLOYEESTATUSREJECT, constants.EMPLOYEESTATUSRELIEVE, constants.EMPLOYEESTATUSOFFBOARD, constants.EMPLOYEESTATUSONBORADING}}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	query2 := bson.M{"$match": bson.M{}}
	if filter.IsZero == "No" {
		query2["$match"] = bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$gt": []interface{}{"$leave.value", 0}}}}}
	}

	// if filter != nil {
	// 	if filter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	// 	}
	// }
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONLEAVEPOLICY,
			"as":   "employeeLeave",
			"let":  bson.M{"leavePolicyID": "$leavePolicyID", "employeeid": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
					"$eq": []string{"$uniqueId", "$$leavePolicyID"},
					//  {"$eq":["$employeeID","$$employeeid"]},
				}}}},

				{"$unwind": bson.M{"path": "$leavemaster"}},
				{
					"$lookup": bson.M{
						"from": constants.COLLECTIONLEAVEMASTER,
						"as":   "leave",
						"let":  bson.M{"leavemasterId": "$leavemaster.uniqueId", "leavePolicyID": "$$leavePolicyID"},
						"pipeline": []bson.M{
							{"$match": bson.M{"$expr": bson.M{"$and": bson.M{
								"$eq": []string{"$uniqueId", "$$leavemasterId"},
								//  {"$eq":["$employeeID","$$employeeid"]},
							}}}},

							{"$project": bson.M{"name": 1, "uniqueId": 1, "leaveType": 1, "value": 1}},
							{
								"$lookup": bson.M{
									"from": constants.COLLECTIONEMPLOYEELEAVE,
									"as":   "value",
									"let":  bson.M{"leavemasterId": "$uniqueId", "leavePolicyID": "$$leavePolicyID"},
									"pipeline": []bson.M{
										{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
											{"$eq": []string{"$leaveType", "$$leavemasterId"}},
											{"$eq": []string{"$employeeId", "$$employeeid"}},
										}}}},
										//	{"$project": bson.M{"name": 1, "uniqueId": 1, "file": 1}},
										{
											"$lookup": bson.M{
												"as":   "request",
												"from": "employeetimeoff",
												"let":  bson.M{"leavemasterId": "$$leavemasterId", "employeeId": "$$employeeid"},
												"pipeline": []bson.M{
													{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
														{"$eq": []string{"$leaveType", "$$leavemasterId"}},
														{"$eq": []string{"$employeeId", "$$employeeid"}},
														{"$eq": []string{"$status", "Request"}},
													}}}},
													{"$group": bson.M{
														"_id":          nil,
														"numberOfDays": bson.M{"$sum": "$numberOfDays"}}},
												},
											},
										},

										{"$addFields": bson.M{"numberOfDays": bson.M{"$arrayElemAt": []interface{}{"$request.numberOfDays", 0}}}},
										//{"$inc":{"value":-"numberOfDays"}}
										{"$project": bson.M{"value": 1, "numberOfDays": 1, "_id": 0}},
										//{"$project": bson.M{"value": bson.M{"$subtract": []string{"$value", "$numberOfDays"}}}},
									},
								},
							},
							{"$addFields": bson.M{"value": bson.M{"$arrayElemAt": []interface{}{"$value", 0}}}},
							{"$addFields": bson.M{"numberOfDays": "$value.numberOfDays"}},
							{"$addFields": bson.M{"value": "$value.value"}},
						},
					},
				},
				{"$addFields": bson.M{"leave": bson.M{"$arrayElemAt": []interface{}{"$leave", 0}}}},
				query2,
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"employeeLeave": "$employeeLeave.leave"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"employee": "$uniqueId"},
	})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmployeeDocumentsCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmployeeLeave  =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeDocumentss []models.EmployeeLeaveList
	// var employeeDocuments *models.EmployeeLeaveList
	if err = cursor.All(context.TODO(), &employeeDocumentss); err != nil {
		return nil, err
	}

	return employeeDocumentss, nil
}
func (d *Daos) UpdateEmployeeLeaveWithEmployeeId(ctx *models.Context, EmployeeId string, LeaveId string, Value int64) error {
	selector := bson.M{"employeeId": EmployeeId, "leaveType": LeaveId}
	data := bson.M{"$set": bson.M{"value": Value}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVE).UpdateOne(ctx.CTX, selector, data)
	return err
}
