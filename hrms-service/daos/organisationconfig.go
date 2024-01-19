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

func (d *Daos) SaveOrganisationConfig(ctx *models.Context, organisationConfig *models.OrganisationConfig) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).InsertOne(ctx.CTX, organisationConfig)
	if err != nil {
		return err
	}
	organisationConfig.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleOrganisationConfig(ctx *models.Context, UniqueID string) (*models.RefOrganisationConfig, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisationConfigs []models.RefOrganisationConfig
	var organisationConfig *models.RefOrganisationConfig
	if err = cursor.All(ctx.CTX, &organisationConfigs); err != nil {
		return nil, err
	}
	if len(organisationConfigs) > 0 {
		organisationConfig = &organisationConfigs[0]
	}
	return organisationConfig, nil
}

func (d *Daos) UpdateOrganisationConfig(ctx *models.Context, organisationConfig *models.OrganisationConfig) error {

	selector := bson.M{"_id": organisationConfig.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": organisationConfig}
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterOrganisationConfig(ctx *models.Context, organisationConfigfilter *models.OrganisationConfigFilter, pagination *models.Pagination) ([]models.RefOrganisationConfig, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if organisationConfigfilter != nil {

		if len(organisationConfigfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": organisationConfigfilter.Status}})
		}
		//Regex
		if organisationConfigfilter.Searchbox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: organisationConfigfilter.Searchbox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if organisationConfigfilter != nil {
		if organisationConfigfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{organisationConfigfilter.SortBy: organisationConfigfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("language query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisationConfig []models.RefOrganisationConfig
	if err = cursor.All(context.TODO(), &organisationConfig); err != nil {
		return nil, err
	}
	return organisationConfig, nil
}

func (d *Daos) EnableOrganisationConfig(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONCONFIGSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableOrganisationConfig(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONCONFIGSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteOrganisationConfig(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONCONFIGSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) SetdefaultOrganisationConfig(ctx *models.Context, UniqueID string) error {
	filter := bson.M{
		"isdefault": bson.M{
			"$eq": true,
		},
	}
	updatemany := bson.M{"$set": bson.M{"isdefault": false}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).UpdateMany(ctx.CTX, filter, updatemany)
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isdefault": constants.ORGANISATIONCONFIGSTATUSTRUE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) GetactiveOrganisationConfig(ctx *models.Context, IsDefault bool) (*models.OrganisationConfig, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isdefault": IsDefault}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATIONCONFIG, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATIONCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisations []models.OrganisationConfig
	var organisation *models.OrganisationConfig
	if err = cursor.All(ctx.CTX, &organisations); err != nil {
		return nil, err
	}
	if len(organisations) > 0 {
		organisation = &organisations[0]
	}
	return organisation, nil
}
