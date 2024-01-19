package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveOrganisation :""
func (d *Daos) SaveOrganisation(ctx *models.Context, organisation *models.Organisation) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).InsertOne(ctx.CTX, organisation)
	if err != nil {
		return err
	}
	organisation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleOrganisation : ""
func (d *Daos) GetSingleOrganisation(ctx *models.Context, UniqueID string) (*models.RefOrganisation, error) {
	fmt.Println("organisation====>", UniqueID)
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	d.Shared.BsonToJSONPrintTag("organisation query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisations []models.RefOrganisation
	var organisation *models.RefOrganisation
	if err = cursor.All(ctx.CTX, &organisations); err != nil {
		return nil, err
	}
	if len(organisations) > 0 {
		organisation = &organisations[0]
	}
	return organisation, nil
}

//UpdateOrganisation : ""
func (d *Daos) UpdateOrganisation(ctx *models.Context, organisation *models.Organisation) error {

	selector := bson.M{"_id": organisation.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": organisation}
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOrganisation : ""
func (d *Daos) FilterOrganisation(ctx *models.Context, organisationfilter *models.OrganisationFilter, pagination *models.Pagination) ([]models.RefOrganisation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if organisationfilter != nil {

		if len(organisationfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": organisationfilter.ActiveStatus}})
		}
		if len(organisationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": organisationfilter.Status}})
		}
		//Regex
		if organisationfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: organisationfilter.Regex.Name, Options: "xi"}})
		}
		if len(organisationfilter.ID) > 0 {
			query = append(query, bson.M{"_id": bson.M{"$in": organisationfilter.ID}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if organisationfilter != nil {
		if organisationfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{organisationfilter.SortBy: organisationfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("organisation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisations []models.RefOrganisation
	if err = cursor.All(context.TODO(), &organisations); err != nil {
		return nil, err
	}
	return organisations, nil
}

//EnableOrganisation :""
func (d *Daos) EnableOrganisation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOrganisation :""
func (d *Daos) DisableOrganisation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOrganisation :""
func (d *Daos) DeleteOrganisation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleOrganisationWithName : ""
func (d *Daos) GetSingleOrganisationWithName(ctx *models.Context, Name string) ([]models.RefOrganisation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("organisation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var orgs []models.RefOrganisation
	if err = cursor.All(context.TODO(), &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}

//GetSingleOrganisationWithName : ""
func (d *Daos) GetSingleOrganisationWithNameV2(ctx *models.Context, Name string, isRegex bool) (*models.RefOrganisation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Name != "" {
		if isRegex {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
		} else {
			query = append(query, bson.M{"name": Name})

		}
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("organisation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var orgs []models.RefOrganisation
	if err = cursor.All(context.TODO(), &orgs); err != nil {
		return nil, err
	}
	if len(orgs) > 0 {
		return &orgs[0], nil
	}
	return nil, errors.New("org not available")
}
func (d *Daos) GetSingleOrganisationWithUniqueID(ctx *models.Context, UniqueID string) (*models.RefOrganisation, error) {
	mainPipeline := []bson.M{}

	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("organisation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisations []models.RefOrganisation
	if err = cursor.All(context.TODO(), &organisations); err != nil {
		return nil, err
	}
	if len(organisations) > 0 {
		return &organisations[0], nil
	}

	return nil, errors.New("organisation not found")
}
func (d *Daos) OrganistationProjectDetails(ctx *models.Context, statefilter *models.StateFilter) ([]models.OrganistationProjectDetails, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active"}},

		bson.M{"$project": bson.M{"id": 1, "name": 1, "uniqueId": 1}},
		//bson.M{"$group": bson.M{"_id": "$_id", "state": bson.M{"$push": "$name"}, "stateCode": bson.M{"$push": "$uniqueId"}}}
	)
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROJECT,
		"as":   "projects",
		"let":  bson.M{"orgid": "$_id"},
		"pipeline": []bson.M{{
			"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$organisation", "$$orgid"}},
			}}}},
			bson.M{"$project": bson.M{"id": 1, "name": 1, "state": 1, "uniqueId": 1}},
			bson.M{"$sort": bson.M{"name": 1}},
		}},
	})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("OrganistationProjectDetails query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var orgs []models.OrganistationProjectDetails
	if err = cursor.All(context.TODO(), &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}
