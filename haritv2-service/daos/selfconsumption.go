package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveSelfConsumption : ""
func (d *Daos) SaveSelfConsumption(ctx *models.Context, selfconsumption *models.SelfConsumption) error {
	d.Shared.BsonToJSONPrint(selfconsumption)
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).InsertOne(ctx.CTX, selfconsumption)
	return err
}

// GetSingleSelfConsumption : ""
func (d *Daos) GetSingleSelfConsumption(ctx *models.Context, UniqueID string) (*models.RefSelfConsumption, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "mobile", "mobile", "ref.chairman", "ref.chairman")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ulbId", "uniqueId", "ref.ulbId", "ref.ulbId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "by", "uniqueId", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERTYPE, "byType", "uniqueId", "ref.userType", "ref.userType")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefSelfConsumption
	var tower *models.RefSelfConsumption
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	fmt.Println(tower)
	return tower, nil

}

// UpdateSelfConsumption : ""
func (d *Daos) UpdateSelfConsumption(ctx *models.Context, business *models.SelfConsumption) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableSelfConsumption : ""
func (d *Daos) EnableSelfConsumption(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SELFCONSUMPTIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableSelfConsumption : ""
func (d *Daos) DisableSelfConsumption(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SELFCONSUMPTIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteSelfConsumption : ""
func (d *Daos) DeleteSelfConsumption(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SELFCONSUMPTIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSelfConsumption : ""
func (d *Daos) FilterSelfConsumption(ctx *models.Context, filter *models.SelfConsumptionFilter, pagination *models.Pagination) ([]models.RefSelfConsumption, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.ULBID) > 0 {
			query = append(query, bson.M{"ulbId": bson.M{"$in": filter.ULBID}})
		}
		if len(filter.CompanyID) > 0 {
			query = append(query, bson.M{"companyId": bson.M{"$in": filter.CompanyID}})
		}
	}
	//daterange
	if filter.SelfConsumptionRange != nil {
		//var sd,ed time.Time
		if filter.SelfConsumptionRange.From != nil {
			sd := time.Date(filter.SelfConsumptionRange.From.Year(), filter.SelfConsumptionRange.From.Month(), filter.SelfConsumptionRange.From.Day(), 0, 0, 0, 0, filter.SelfConsumptionRange.From.Location())
			ed := time.Date(filter.SelfConsumptionRange.From.Year(), filter.SelfConsumptionRange.From.Month(), filter.SelfConsumptionRange.From.Day(), 23, 59, 59, 0, filter.SelfConsumptionRange.From.Location())
			if filter.SelfConsumptionRange.To != nil {
				ed = time.Date(filter.SelfConsumptionRange.To.Year(), filter.SelfConsumptionRange.To.Month(), filter.SelfConsumptionRange.To.Day(), 23, 59, 59, 0, filter.SelfConsumptionRange.To.Location())
			}
			query = append(query, bson.M{"selfConsumptionRange": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	// //Regex
	if filter.Regex.Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
	}
	// if filter.Regex.Email != "" {
	// 	query = append(query, bson.M{"email": primitive.Regex{Pattern: filter.Regex.Email, Options: "xi"}})
	// }
	// if filter.Regex.Mobile != "" {
	// 	query = append(query, bson.M{"primaryContact.ph": primitive.Regex{Pattern: filter.Regex.Mobile, Options: "xi"}})
	// }
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "companyId", "uniqueId", "ref.ulbId", "ref.ulbId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "by", "userName", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSERTYPE, "byType", "uniqueId", "ref.userType", "ref.userType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var selfconsumption []models.RefSelfConsumption
	if err = cursor.All(context.TODO(), &selfconsumption); err != nil {
		return nil, err
	}
	return selfconsumption, nil
}
func (d *Daos) GetSingleSelfConsumptionwithmobileno(ctx *models.Context, mobile string) (*models.RefSelfConsumption, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"primaryContact.ph": mobile}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "mobile", "mobile", "ref.chairman", "ref.chairman")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSELFCONSUMPTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefSelfConsumption
	var tower *models.RefSelfConsumption
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	fmt.Println(tower)
	return tower, nil

}
