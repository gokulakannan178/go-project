package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveRouteMaster : ""
func (d *Daos) SaveRouteMaster(ctx *models.Context, RouteMaster *models.RouteMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).InsertOne(ctx.CTX, RouteMaster)
	if err != nil {
		return err
	}
	RouteMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// UpdateRouteMaster : ""
func (d *Daos) UpdateRouteMaster(ctx *models.Context, RouteMaster *models.RouteMaster) error {
	selector := bson.M{"uniqueId": RouteMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": RouteMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleRouteMaster : ""
func (d *Daos) GetSingleRouteMaster(ctx *models.Context, uniqueID string) (*models.RefRouteMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadtype", "uniqueId", "ref.roadtype", "ref.roadtype")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "area.zoneCode", "code", "ref.zone", "ref.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "area.villageCode", "code", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "area.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "area.stateCode", "code", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "area.wardCode", "code", "ref.ward", "ref.ward")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var RouteMasters []models.RefRouteMaster
	var RouteMaster *models.RefRouteMaster
	if err = cursor.All(ctx.CTX, &RouteMasters); err != nil {
		return nil, err
	}
	if len(RouteMasters) > 0 {
		RouteMaster = &RouteMasters[0]
	}
	return RouteMaster, err
}

// EnableRouteMaster : ""
func (d *Daos) EnableRouteMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ROUTEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableRouteMaster : ""
func (d *Daos) DisableRouteMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ROUTEMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteState :""
func (d *Daos) DeleteRouteMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROUTEMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterRouteMaster : ""
func (d *Daos) FilterRouteMaster(ctx *models.Context, filter *models.FilterRouteMaster, pagination *models.Pagination) ([]models.RefRouteMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.EmployeeId}})
		}
		if len(filter.WardCode) > 0 {
			query = append(query, bson.M{"area.wardCode": bson.M{"$in": filter.WardCode}})
		}
		if len(filter.ZoneCode) > 0 {
			query = append(query, bson.M{"area.zoneCone": bson.M{"$in": filter.ZoneCode}})
		}
		if len(filter.Roadtype) > 0 {
			query = append(query, bson.M{"roadtype": bson.M{"$in": filter.Roadtype}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.ContactNo != "" {
			query = append(query, bson.M{"mobile": primitive.Regex{Pattern: filter.Regex.ContactNo, Options: "xi"}})
		}
		if filter.Regex.Type != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: filter.Regex.Type, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadtype", "uniqueId", "ref.roadtype", "ref.roadtype")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROUTEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var RouteMasterFilter []models.RefRouteMaster
	if err = cursor.All(context.TODO(), &RouteMasterFilter); err != nil {
		return nil, err
	}
	return RouteMasterFilter, nil
}
