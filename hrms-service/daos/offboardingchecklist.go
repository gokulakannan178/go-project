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

// SaveOffboardingCheckList : ""
func (d *Daos) SaveOffboardingCheckList(ctx *models.Context, offboardingchecklist *models.OffboardingCheckList) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).InsertOne(ctx.CTX, offboardingchecklist)
	if err != nil {
		return err
	}
	offboardingchecklist.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// GetSingleOffboardingCheckList : ""
func (d *Daos) GetSingleOffboardingCheckList(ctx *models.Context, uniqueID string) (*models.RefOffboardingCheckList, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER, "offboardingchecklistmasterId", "uniqueId", "ref.offboardingchecklistmasterId", "ref.offboardingchecklistmasterId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGPOLICY, "offboardingpolicyId", "uniqueId", "ref.offboardingpolicyId", "ref.offboardingpolicyId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offboardingchecklists []models.RefOffboardingCheckList
	var offboardingchecklist *models.RefOffboardingCheckList
	if err = cursor.All(ctx.CTX, &offboardingchecklists); err != nil {
		return nil, err
	}
	if len(offboardingchecklists) > 0 {
		offboardingchecklist = &offboardingchecklists[0]
	}
	return offboardingchecklist, err
}

//UpdateOffboardingCheckList : ""
func (d *Daos) UpdateOffboardingCheckList(ctx *models.Context, offboardingchecklist *models.OffboardingCheckList) error {
	selector := bson.M{"uniqueId": offboardingchecklist.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": offboardingchecklist}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// OffboardingCheckListRemoveNotPresentValue : ""
func (d *Daos) OffboardingCheckListRemoveNotPresentValue(ctx *models.Context, offboardingpolicyId string, arrayValue []string) error {
	selector := bson.M{"offboardingpolicyId": offboardingpolicyId, "offboardingchecklistmasterId": bson.M{"$nin": arrayValue}}
	d.Shared.BsonToJSONPrintTag("selector query in offboarding checklist =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGCHECKLISTSTATUSDELETED}}
	d.Shared.BsonToJSONPrintTag("data query in offboarding checklist =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).UpdateMany(ctx.CTX, selector, data)
	return err
}

// OffboardingCheckListUpsert : ""
func (d *Daos) OffboardingCheckListUpsert(ctx *models.Context, offboardingpolicyId string, arrayValue []string, name string) error {
	fmt.Println("arrayValue", arrayValue)
	for _, v := range arrayValue {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"offboardingpolicyId": offboardingpolicyId, "offboardingchecklistmasterId": v}
		fmt.Println("updateQuery===>", updateQuery)
		offboardingCheckList := new(models.OffboardingCheckList)
		offboardingCheckList.Status = constants.OFFBOARDINGCHECKLISTSTATUSACTIVE
		offboardingCheckList.Name = name
		offboardingCheckList.UniqueID = d.GetUniqueID(ctx, constants.COLLECTIONOFFBOARDINGCHECKLIST)
		fmt.Println("present added =======>", offboardingCheckList.UniqueID)
		updateData := bson.M{"$set": offboardingCheckList}
		if _, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

// EnableOffboardingCheckList : ""
func (d *Daos) EnableOffboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGCHECKLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableOffboardingCheckList : ""
func (d *Daos) DisableOffboardingCheckList(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGCHECKLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteOffboardingCheckList :""
func (d *Daos) DeleteOffboardingCheckList(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGCHECKLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOffboardingCheckList : ""
func (d *Daos) FilterOffboardingCheckList(ctx *models.Context, offboardingchecklist *models.FilterOffboardingCheckList, pagination *models.Pagination) ([]models.RefOffboardingCheckList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if offboardingchecklist != nil {
		if len(offboardingchecklist.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": offboardingchecklist.Status}})
		}
		if len(offboardingchecklist.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": offboardingchecklist.OrganisationID}})
		}
		//Regex
		if offboardingchecklist.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: offboardingchecklist.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if offboardingchecklist != nil {
		if offboardingchecklist.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{offboardingchecklist.SortBy: offboardingchecklist.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER, "offboardingchecklistmasterId", "uniqueId", "ref.offboardingchecklistmasterId", "ref.offboardingchecklistmasterId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOFFBOARDINGPOLICY, "offboardingpolicyId", "uniqueId", "ref.offboardingpolicyId", "ref.offboardingpolicyId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offboardingchecklistFilter []models.RefOffboardingCheckList
	if err = cursor.All(context.TODO(), &offboardingchecklistFilter); err != nil {
		return nil, err
	}
	return offboardingchecklistFilter, nil
}
