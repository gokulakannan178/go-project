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

// SaveOnboardingCheckListMaster : ""
func (d *Daos) SaveOnboardingCheckListMaster(ctx *models.Context, onboardingchecklistmaster *models.OnboardingCheckListMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).InsertOne(ctx.CTX, onboardingchecklistmaster)
	if err != nil {
		return err
	}
	onboardingchecklistmaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateOnboardingCheckListMaster : ""
func (d *Daos) UpdateOnboardingCheckListMaster(ctx *models.Context, onboardingchecklistmaster *models.OnboardingCheckListMaster) error {
	selector := bson.M{"uniqueId": onboardingchecklistmaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": onboardingchecklistmaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleOnboardingCheckListMaster : ""
func (d *Daos) GetSingleOnboardingCheckListMaster(ctx *models.Context, uniqueID string) (*models.RefOnboardingCheckListMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingchecklistmasters []models.RefOnboardingCheckListMaster
	var onboardingchecklistmaster *models.RefOnboardingCheckListMaster
	if err = cursor.All(ctx.CTX, &onboardingchecklistmasters); err != nil {
		return nil, err
	}
	if len(onboardingchecklistmasters) > 0 {
		onboardingchecklistmaster = &onboardingchecklistmasters[0]
	}
	return onboardingchecklistmaster, err
}

// GetSingleOnboardingCheckListMasterWithActive : ""
func (d *Daos) GetSingleOnboardingCheckListMasterWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefOnboardingCheckListMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingchecklistmasters []models.RefOnboardingCheckListMaster
	var onboardingchecklistmaster *models.RefOnboardingCheckListMaster
	if err = cursor.All(ctx.CTX, &onboardingchecklistmasters); err != nil {
		return nil, err
	}
	if len(onboardingchecklistmasters) > 0 {
		onboardingchecklistmaster = &onboardingchecklistmasters[0]
	}
	return onboardingchecklistmaster, err
}

// EnableOnboardingCheckListMaster : ""
func (d *Daos) EnableOnboardingCheckListMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ONBOARDINGCHECKLISTMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableOnboardingCheckListMaster : ""
func (d *Daos) DisableOnboardingCheckListMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ONBOARDINGCHECKLISTMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteOnboardingCheckListMaster :""
func (d *Daos) DeleteOnboardingCheckListMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ONBOARDINGCHECKLISTMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOnboardingCheckListMaster : ""
func (d *Daos) FilterOnboardingCheckListMaster(ctx *models.Context, onboardingchecklistmaster *models.FilterOnboardingCheckListMaster, pagination *models.Pagination) ([]models.RefOnboardingCheckListMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if onboardingchecklistmaster != nil {
		if len(onboardingchecklistmaster.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": onboardingchecklistmaster.Status}})
		}
		if len(onboardingchecklistmaster.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": onboardingchecklistmaster.OrganisationID}})
		}
		//Regex
		if onboardingchecklistmaster.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: onboardingchecklistmaster.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if onboardingchecklistmaster != nil {
		if onboardingchecklistmaster.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{onboardingchecklistmaster.SortBy: onboardingchecklistmaster.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingchecklistmasterFilter []models.RefOnboardingCheckListMaster
	if err = cursor.All(context.TODO(), &onboardingchecklistmasterFilter); err != nil {
		return nil, err
	}
	return onboardingchecklistmasterFilter, nil
}
