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

// SaveEmployeeEmployeeEmployeeOffboardingCheckList : ""
func (d *Daos) SaveEmployeeOffboardingCheckList(ctx *models.Context, employeeoffboardingchecklist *models.EmployeeOffboardingCheckList) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).InsertOne(ctx.CTX, employeeoffboardingchecklist)
	if err != nil {
		return err
	}
	employeeoffboardingchecklist.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeOffboardingCheckList : ""
func (d *Daos) UpdateEmployeeOffboardingCheckList(ctx *models.Context, employeeoffboardingchecklist *models.EmployeeOffboardingCheckList) error {
	selector := bson.M{"uniqueId": employeeoffboardingchecklist.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeoffboardingchecklist}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeOffboardingCheckList : ""
func (d *Daos) GetSingleEmployeeOffboardingCheckList(ctx *models.Context, uniqueID string) (*models.RefEmployeeOffboardingCheckList, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGPOLICY, "offboardingpolicyId", "uniqueId", "ref.offboardingpolicyId", "ref.offboardingpolicyId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER, "offboardingchecklistmasterId", "uniqueId", "ref.offboardingchecklistmasterId", "ref.offboardingchecklistmasterId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeOffboardingCheckLists []models.RefEmployeeOffboardingCheckList
	var EmployeeOffboardingCheckList *models.RefEmployeeOffboardingCheckList
	if err = cursor.All(ctx.CTX, &EmployeeOffboardingCheckLists); err != nil {
		return nil, err
	}
	if len(EmployeeOffboardingCheckLists) > 0 {
		EmployeeOffboardingCheckList = &EmployeeOffboardingCheckLists[0]
	}
	return EmployeeOffboardingCheckList, err
}

// EmployeeOffboardingCheckListFinal : ""
func (d *Daos) EmployeeOffboardingCheckListFinal(ctx *models.Context, EmployeeID string, PolicyId string) (*models.RefEmployeeOffboardingCheckListv2, error) {
	fmt.Println("PolicyId", PolicyId)
	mainPipeline := []bson.M{}
	//LookUp
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": EmployeeID}},

		bson.M{"$lookup": bson.M{
			"from": constants.COLLECTIONOFFBOARDINGPOLICY,
			"as":   "offboardingpolicy",
			"let":  bson.M{"employeeId": "$uniqueId", "policyId": "$offboardingpolicyId"},
			"pipeline": []bson.M{

				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

					bson.M{"$eq": []interface{}{"$status", "Active"}},

					bson.M{"$eq": []interface{}{"$uniqueId", "$$policyId"}},
				}}}},

				bson.M{"$lookup": bson.M{
					"from": constants.COLLECTIONOFFBOARDINGCHECKLIST,
					"as":   "checklist",
					"let":  bson.M{"employeeId": "$$employeeId", "policyId": "$$policyId"},

					"pipeline": []bson.M{

						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

							bson.M{"$eq": []interface{}{"$status", "Active"}},

							bson.M{"$eq": []interface{}{"$offboardingpolicyId", "$$policyId"}},
						}}}},

						//for policy name

						bson.M{"$lookup": bson.M{
							"from":         constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER,
							"as":           "ref.checklistName",
							"localField":   "offboardingchecklistmasterId",
							"foreignField": "uniqueId"}},
						//for policy is checked
						bson.M{"$lookup": bson.M{
							"from": constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST,
							"as":   "ref.isChecked",
							"let":  bson.M{"employeeId": "$$employeeId", "policyId": "$$policyId", "checklistId": "$offboardingchecklistmasterId"},
							"pipeline": []bson.M{
								bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
									bson.M{"$eq": []interface{}{"$status", "Active"}},
									bson.M{"$eq": []interface{}{"$employeeId", "$$employeeId"}},
									bson.M{"$eq": []interface{}{"$offboardingchecklistmasterId", "$$checklistId"}}}}}},
							}}},
						bson.M{"$addFields": bson.M{
							"ref.isChecked":     bson.M{"$arrayElemAt": []interface{}{"$ref.isChecked", 0}},
							"ref.checklistName": bson.M{"$arrayElemAt": []interface{}{"$ref.checklistName", 0}}}},
					},
				}},
			},
		}},

		bson.M{"$addFields": bson.M{"offboardingpolicy": bson.M{"$arrayElemAt": []interface{}{"$offboardingpolicy", 0}}}})

	//fmt.Println("mainPipeline", mainPipeline)
	d.Shared.BsonToJSONPrintTag("mainPipeline query =>", mainPipeline)

	fmt.Println("i am here in employee onboarding checklist")
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeOffboardingCheckLists []models.RefEmployeeOffboardingCheckListv2
	var EmployeeOffboardingCheckList *models.RefEmployeeOffboardingCheckListv2
	if err = cursor.All(ctx.CTX, &EmployeeOffboardingCheckLists); err != nil {
		return nil, err
	}
	if len(EmployeeOffboardingCheckLists) > 0 {
		EmployeeOffboardingCheckList = &EmployeeOffboardingCheckLists[0]
	}
	return EmployeeOffboardingCheckList, err
}

// EnableEmployeeOffboardingCheckList : ""
func (d *Daos) EnableEmployeeOffboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEOFFBOARDINGCHECKLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeOffboardingCheckList : ""
func (d *Daos) DisableEmployeeOffboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEOFFBOARDINGCHECKLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeOffboardingCheckList :""
func (d *Daos) DeleteEmployeeOffboardingCheckList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEOFFBOARDINGCHECKLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeOffboardingCheckList : ""
func (d *Daos) FilterEmployeeOffboardingCheckList(ctx *models.Context, employeeoffboardingchecklist *models.FilterEmployeeOffboardingCheckList, pagination *models.Pagination) ([]models.RefEmployeeOffboardingCheckList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeoffboardingchecklist != nil {
		if len(employeeoffboardingchecklist.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeoffboardingchecklist.Status}})
		}
		if len(employeeoffboardingchecklist.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeoffboardingchecklist.OrganisationId}})
		}
		if len(employeeoffboardingchecklist.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeoffboardingchecklist.EmployeeId}})
		}
		if len(employeeoffboardingchecklist.OffboardingpolicyId) > 0 {
			query = append(query, bson.M{"offboardingpolicyId": bson.M{"$eq": employeeoffboardingchecklist.OffboardingpolicyId}})
		}
		if len(employeeoffboardingchecklist.OffboardingchecklistmasterId) > 0 {
			query = append(query, bson.M{"offboardingchecklistId": bson.M{"$eq": employeeoffboardingchecklist.OffboardingchecklistmasterId}})
		}
		//Regex
		// if EmployeeOffboardingCheckList.Regex.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: EmployeeOffboardingCheckList.Regex.Name, Options: "xi"}})
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGPOLICY, "offboardingpolicyId", "uniqueId", "ref.offboardingpolicyId", "ref.offboardingpolicyId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER, "offboardingchecklistmasterId", "uniqueId", "ref.offboardingchecklistmasterId", "ref.offboardingchecklistmasterId")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rf []models.RefEmployeeOffboardingCheckList
	if err = cursor.All(context.TODO(), &rf); err != nil {
		return nil, err
	}
	return rf, nil
}
