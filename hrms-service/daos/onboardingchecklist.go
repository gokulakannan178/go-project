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

// SaveOnboardingCheckList : ""
func (d *Daos) SaveOnboardingCheckList(ctx *models.Context, onboardingchecklist *models.OnboardingCheckList) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).InsertOne(ctx.CTX, onboardingchecklist)
	if err != nil {
		return err
	}
	onboardingchecklist.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveOnboardingCheckListUpdert(ctx *models.Context, onboardingchecklist *models.OnboardingCheckList) error {
	//fmt.Println("arrayValue", arrayValue)
	//	for _, v := range onboardingchecklist.OnboardingchecklistmasterID {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"onboardingpolicyId": onboardingchecklist.OnboardingpolicyID, "onboardingchecklistmasterId": onboardingchecklist.OnboardingchecklistmasterID}
	fmt.Println("updateQuery===>", updateQuery)

	//fmt.Println("present added =======>", AssetPolicyAssets.UniqueID)
	updateData := bson.M{"$set": onboardingchecklist}
	if _, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	//s}
	return nil
}

// GetSingleOnboardingCheckList : ""
func (d *Daos) GetSingleOnboardingCheckList(ctx *models.Context, uniqueID string) (*models.RefOnboardingCheckList, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGCHECKLISTMASTER, "onboardingchecklistmasterId", "uniqueId", "ref.onboardingchecklistmasterId", "ref.onboardingchecklistmasterId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGPOLICY, "onboardingpolicyId", "uniqueId", "ref.onboardingpolicyId", "ref.onboardingpolicyId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingchecklists []models.RefOnboardingCheckList
	var onboardingchecklist *models.RefOnboardingCheckList
	if err = cursor.All(ctx.CTX, &onboardingchecklists); err != nil {
		return nil, err
	}
	if len(onboardingchecklists) > 0 {
		onboardingchecklist = &onboardingchecklists[0]
	}
	return onboardingchecklist, err
}

//UpdateOnboardingCheckList : ""
func (d *Daos) UpdateOnboardingCheckList(ctx *models.Context, onboardingchecklist *models.OnboardingCheckList) error {
	selector := bson.M{"uniqueId": onboardingchecklist.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": onboardingchecklist}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// OnboardingCheckListRemoveNotPresentValue : ""
func (d *Daos) OnboardingCheckListRemoveNotPresentValue(ctx *models.Context, onboardingpolicyId string, arrayValue []string) error {
	selector := bson.M{"onboardingpolicyId": onboardingpolicyId, "onboardingchecklistmasterId": bson.M{"$nin": arrayValue}}
	d.Shared.BsonToJSONPrintTag("selector query in onboarding checklist =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.ONBOARDINGCHECKLISTSTATUSDELETED}}
	d.Shared.BsonToJSONPrintTag("data query in onboarding checklist =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).UpdateMany(ctx.CTX, selector, data)
	return err
}

// OnboardingCheckListUpsert : ""
func (d *Daos) OnboardingCheckListUpsert(ctx *models.Context, onboardingpolicyId string, arrayValue []string, name string) error {
	fmt.Println("arrayValue", arrayValue)
	for _, v := range arrayValue {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"onboardingpolicyId": onboardingpolicyId, "onboardingchecklistmasterId": v}
		fmt.Println("updateQuery===>", updateQuery)
		onboardingCheckList := new(models.OnboardingCheckList)
		onboardingCheckList.Status = constants.ONBOARDINGCHECKLISTSTATUSACTIVE
		onboardingCheckList.Name = name
		onboardingCheckList.UniqueID = d.GetUniqueID(ctx, constants.COLLECTIONONBOARDINGCHECKLIST)
		fmt.Println("present added =======>", onboardingCheckList.UniqueID)
		updateData := bson.M{"$set": onboardingCheckList}
		if _, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

// EnableOnboardingCheckList : ""
func (d *Daos) EnableOnboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ONBOARDINGCHECKLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableOnboardingCheckList : ""
func (d *Daos) DisableOnboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ONBOARDINGCHECKLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteOnboardingCheckList :""
func (d *Daos) DeleteOnboardingCheckList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ONBOARDINGCHECKLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOnboardingCheckList : ""
func (d *Daos) FilterOnboardingCheckList(ctx *models.Context, onboardingchecklist *models.FilterOnboardingCheckList, pagination *models.Pagination) ([]models.RefOnboardingCheckList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if onboardingchecklist != nil {
		if len(onboardingchecklist.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": onboardingchecklist.Status}})
		}
		if len(onboardingchecklist.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": onboardingchecklist.OrganisationID}})
		}
		//Regex
		if onboardingchecklist.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: onboardingchecklist.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if onboardingchecklist != nil {
		if onboardingchecklist.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{onboardingchecklist.SortBy: onboardingchecklist.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGCHECKLISTMASTER, "onboardingchecklistmasterId", "uniqueId", "ref.onboardingchecklistmasterId", "ref.onboardingchecklistmasterId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONONBOARDINGPOLICY, "onboardingpolicyId", "uniqueId", "ref.onboardingpolicyId", "ref.onboardingpolicyId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingchecklistFilter []models.RefOnboardingCheckList
	if err = cursor.All(context.TODO(), &onboardingchecklistFilter); err != nil {
		return nil, err
	}
	return onboardingchecklistFilter, nil
}
