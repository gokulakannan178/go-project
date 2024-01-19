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

// SaveEmployeeEmployeeEmployeeOnboardingCheckList : ""
func (d *Daos) SaveEmployeeOnboardingCheckList(ctx *models.Context, employeeonboardingchecklist *models.EmployeeOnboardingCheckList) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).InsertOne(ctx.CTX, employeeonboardingchecklist)
	if err != nil {
		return err
	}
	employeeonboardingchecklist.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeOnboardingCheckList : ""
func (d *Daos) UpdateEmployeeOnboardingCheckList(ctx *models.Context, employeeonboardingchecklist *models.EmployeeOnboardingCheckList) error {
	selector := bson.M{"uniqueId": employeeonboardingchecklist.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeonboardingchecklist}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeOnboardingCheckList : ""
func (d *Daos) GetSingleEmployeeOnboardingCheckList(ctx *models.Context, uniqueID string) (*models.RefEmployeeOnboardingCheckList, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGPOLICY, "onboardingpolicyId", "uniqueId", "ref.onboardingpolicyId", "ref.onboardingpolicyId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGCHECKLISTMASTER, "onboardingchecklistmasterId", "uniqueId", "ref.onboardingchecklistmasterId", "ref.onboardingchecklistmasterId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeOnboardingCheckLists []models.RefEmployeeOnboardingCheckList
	var EmployeeOnboardingCheckList *models.RefEmployeeOnboardingCheckList
	if err = cursor.All(ctx.CTX, &EmployeeOnboardingCheckLists); err != nil {
		return nil, err
	}
	if len(EmployeeOnboardingCheckLists) > 0 {
		EmployeeOnboardingCheckList = &EmployeeOnboardingCheckLists[0]
	}
	return EmployeeOnboardingCheckList, err
}

// EmployeeOnboardingCheckListFinal : ""
func (d *Daos) EmployeeOnboardingCheckListFinal(ctx *models.Context, EmployeeID string, PolicyId string) (*models.RefEmployeeOnboardingCheckListv2, error) {
	fmt.Println("PolicyId", PolicyId)
	mainPipeline := []bson.M{}
	//LookUp
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": EmployeeID}},

		bson.M{"$lookup": bson.M{
			"from": constants.COLLECTIONONBOARDINGPOLICY,
			"as":   "onboardingpolicy",
			"let":  bson.M{"employeeId": "$uniqueId", "policyId": "$onboardingpolicyId"},
			"pipeline": []bson.M{

				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

					bson.M{"$eq": []interface{}{"$status", "Active"}},

					bson.M{"$eq": []interface{}{"$uniqueId", "$$policyId"}},
				}}}},

				bson.M{"$lookup": bson.M{
					"from": constants.COLLECTIONONBOARDINGCHECKLIST,
					"as":   "checklist",
					"let":  bson.M{"employeeId": "$$employeeId", "policyId": "$$policyId"},

					"pipeline": []bson.M{

						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

							bson.M{"$eq": []interface{}{"$status", "Active"}},

							bson.M{"$eq": []interface{}{"$onboardingpolicyId", "$$policyId"}},
						}}}},

						//for policy name

						bson.M{"$lookup": bson.M{
							"from":         constants.COLLECTIONONBOARDINGCHECKLISTMASTER,
							"as":           "ref.checklistName",
							"localField":   "onboardingchecklistmasterId",
							"foreignField": "uniqueId"}},
						//for policy is checked
						bson.M{"$lookup": bson.M{
							"from": constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST,
							"as":   "ref.isChecked",
							"let":  bson.M{"employeeId": "$$employeeId", "policyId": "$$policyId", "checklistId": "$onboardingchecklistmasterId"},
							"pipeline": []bson.M{
								bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
									bson.M{"$eq": []interface{}{"$status", "Active"}},
									bson.M{"$eq": []interface{}{"$employeeId", "$$employeeId"}},
									bson.M{"$eq": []interface{}{"$onboardingchecklistmasterId", "$$checklistId"}}}}}},
							}}},
						bson.M{"$addFields": bson.M{
							"ref.isChecked":     bson.M{"$arrayElemAt": []interface{}{"$ref.isChecked", 0}},
							"ref.checklistName": bson.M{"$arrayElemAt": []interface{}{"$ref.checklistName", 0}}}},
					},
				}},
			},
		}},

		bson.M{"$addFields": bson.M{"onboardingpolicy": bson.M{"$arrayElemAt": []interface{}{"$onboardingpolicy", 0}}}})

	//fmt.Println("mainPipeline", mainPipeline)
	d.Shared.BsonToJSONPrintTag("mainPipeline query =>", mainPipeline)

	fmt.Println("i am here in employee onboarding checklist")
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeOnboardingCheckLists []models.RefEmployeeOnboardingCheckListv2
	var EmployeeOnboardingCheckList *models.RefEmployeeOnboardingCheckListv2
	if err = cursor.All(ctx.CTX, &EmployeeOnboardingCheckLists); err != nil {
		return nil, err
	}
	if len(EmployeeOnboardingCheckLists) > 0 {
		EmployeeOnboardingCheckList = &EmployeeOnboardingCheckLists[0]
	}
	return EmployeeOnboardingCheckList, err
}

// EnableEmployeeOnboardingCheckList : ""
func (d *Daos) EnableEmployeeOnboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEONBOARDINGCHECKLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeOnboardingCheckList : ""
func (d *Daos) DisableEmployeeOnboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEONBOARDINGCHECKLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeOnboardingCheckList :""
func (d *Daos) DeleteEmployeeOnboardingCheckList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEONBOARDINGCHECKLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeOnboardingCheckList : ""
func (d *Daos) FilterEmployeeOnboardingCheckList(ctx *models.Context, employeeonboardingchecklist *models.FilterEmployeeOnboardingCheckList, pagination *models.Pagination) ([]models.RefEmployeeOnboardingCheckList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeonboardingchecklist != nil {
		if len(employeeonboardingchecklist.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeonboardingchecklist.Status}})
		}
		if len(employeeonboardingchecklist.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeonboardingchecklist.OrganisationId}})
		}
		if len(employeeonboardingchecklist.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeonboardingchecklist.EmployeeId}})
		}
		if len(employeeonboardingchecklist.OnboardingpolicyId) > 0 {
			query = append(query, bson.M{"onboardingpolicyId": bson.M{"$eq": employeeonboardingchecklist.OnboardingpolicyId}})
		}
		if len(employeeonboardingchecklist.OnboardingchecklistmasterId) > 0 {
			query = append(query, bson.M{"onboardingchecklistmasterId": bson.M{"$eq": employeeonboardingchecklist.OnboardingchecklistmasterId}})
		}
		//Regex
		// if EmployeeOnboardingCheckList.Regex.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: EmployeeOnboardingCheckList.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGPOLICY, "onboardingpolicyId", "uniqueId", "ref.onboardingpolicyId", "ref.onboardingpolicyId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGCHECKLISTMASTER, "onboardingchecklistmasterId", "uniqueId", "ref.onboardingchecklistmasterId", "ref.onboardingchecklistmasterId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEONBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rf []models.RefEmployeeOnboardingCheckList
	if err = cursor.All(context.TODO(), &rf); err != nil {
		return nil, err
	}
	return rf, nil
}
