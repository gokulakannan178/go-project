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

// SaveOffboardingCheckListMaster : ""
func (d *Daos) SaveOffboardingCheckListMaster(ctx *models.Context, offboardingchecklistmaster *models.OffboardingCheckListMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).InsertOne(ctx.CTX, offboardingchecklistmaster)
	if err != nil {
		return err
	}
	offboardingchecklistmaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateOffboardingCheckListMaster : ""
func (d *Daos) UpdateOffboardingCheckListMaster(ctx *models.Context, offboardingchecklistmaster *models.OffboardingCheckListMaster) error {
	selector := bson.M{"uniqueId": offboardingchecklistmaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": offboardingchecklistmaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleOffboardingCheckListMaster : ""
func (d *Daos) GetSingleOffboardingCheckListMaster(ctx *models.Context, uniqueID string) (*models.RefOffboardingCheckListMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offboardingchecklistmasters []models.RefOffboardingCheckListMaster
	var offboardingchecklistmaster *models.RefOffboardingCheckListMaster
	if err = cursor.All(ctx.CTX, &offboardingchecklistmasters); err != nil {
		return nil, err
	}
	if len(offboardingchecklistmasters) > 0 {
		offboardingchecklistmaster = &offboardingchecklistmasters[0]
	}
	return offboardingchecklistmaster, err
}

// GetSingleOffboardingCheckListMasterWithActive : ""
func (d *Daos) GetSingleOffboardingCheckListMasterWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefOffboardingCheckListMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offboardingchecklistmasters []models.RefOffboardingCheckListMaster
	var offboardingchecklistmaster *models.RefOffboardingCheckListMaster
	if err = cursor.All(ctx.CTX, &offboardingchecklistmasters); err != nil {
		return nil, err
	}
	if len(offboardingchecklistmasters) > 0 {
		offboardingchecklistmaster = &offboardingchecklistmasters[0]
	}
	return offboardingchecklistmaster, err
}

// EnableOffboardingCheckListMaster : ""
func (d *Daos) EnableOffboardingCheckListMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGCHECKLISTMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableOffboardingCheckListMaster : ""
func (d *Daos) DisableOffboardingCheckListMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGCHECKLISTMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteOffboardingCheckListMaster :""
func (d *Daos) DeleteOffboardingCheckListMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGCHECKLISTMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOffboardingCheckListMaster : ""
func (d *Daos) FilterOffboardingCheckListMaster(ctx *models.Context, offboardingchecklistmaster *models.FilterOffboardingCheckListMaster, pagination *models.Pagination) ([]models.RefOffboardingCheckListMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if offboardingchecklistmaster != nil {
		if len(offboardingchecklistmaster.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": offboardingchecklistmaster.Status}})
		}
		if len(offboardingchecklistmaster.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": offboardingchecklistmaster.OrganisationID}})
		}
		//Regex
		if offboardingchecklistmaster.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: offboardingchecklistmaster.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if offboardingchecklistmaster != nil {
		if offboardingchecklistmaster.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{offboardingchecklistmaster.SortBy: offboardingchecklistmaster.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offboardingchecklistmasterFilter []models.RefOffboardingCheckListMaster
	if err = cursor.All(context.TODO(), &offboardingchecklistmasterFilter); err != nil {
		return nil, err
	}
	return offboardingchecklistmasterFilter, nil
}
