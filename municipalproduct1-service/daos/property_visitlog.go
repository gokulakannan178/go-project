package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SavePropertyVisitLog : ""
func (d *Daos) SavePropertyVisitLog(ctx *models.Context, visitlog *models.PropertyVisitLog) error {
	d.Shared.BsonToJSONPrint(visitlog)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).InsertOne(ctx.CTX, visitlog)
	return err
}

// GetSinglePropertyVisitLog : ""
func (d *Daos) GetSinglePropertyVisitLog(ctx *models.Context, UniqueID string) (*models.RefPropertyVisitLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("letter upload getsingle query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var visitlogs []models.RefPropertyVisitLog
	var visitlog *models.RefPropertyVisitLog
	if err = cursor.All(ctx.CTX, &visitlogs); err != nil {
		return nil, err
	}
	if len(visitlogs) > 0 {
		visitlog = &visitlogs[0]
	}
	return visitlog, nil
}

// UpdatePropertyVisitLog : ""
func (d *Daos) UpdatePropertyVisitLog(ctx *models.Context, visitlog *models.PropertyVisitLog) error {
	selector := bson.M{"uniqueId": visitlog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": visitlog, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyVisitLog : ""
func (d *Daos) EnablePropertyVisitLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVISITLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyVisitLog : ""
func (d *Daos) DisablePropertyVisitLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVISITLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyVisitLog : ""
func (d *Daos) DeletePropertyVisitLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVISITLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyVisitLog : ""
func (d *Daos) FilterPropertyVisitLog(ctx *models.Context, filter *models.PropertyVisitLogFilter, pagination *models.Pagination) ([]models.RefPropertyVisitLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if len(filter.UserId) > 0 {
			query = append(query, bson.M{"userId": bson.M{"$in": filter.UserId}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"userType": bson.M{"$in": filter.UserType}})
		}
		if filter.NextDateRange != nil {
			if filter.NextDateRange.From != nil {
				sd := time.Date(filter.NextDateRange.From.Year(), filter.NextDateRange.From.Month(), filter.NextDateRange.From.Day(), 0, 0, 0, 0, filter.NextDateRange.From.Location())
				ed := time.Date(filter.NextDateRange.From.Year(), filter.NextDateRange.From.Month(), filter.NextDateRange.From.Day(), 23, 59, 59, 0, filter.NextDateRange.From.Location())
				if filter.NextDateRange.To != nil {
					ed = time.Date(filter.NextDateRange.To.Year(), filter.NextDateRange.To.Month(), filter.NextDateRange.To.Day(), 23, 59, 59, 0, filter.NextDateRange.To.Location())
				}
				query = append(query, bson.M{"nextVisitDate": bson.M{"$gte": sd, "$lte": ed}})
			}
		}
		if (filter.Address) != nil {
			if len(filter.Address.Country) > 0 {
				query = append(query, bson.M{"address.country": bson.M{"$in": filter.Address.Country}})
			}
			if len(filter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": filter.Address.StateCode}})
			}
			if len(filter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": filter.Address.DistrictCode}})
			}

			if len(filter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": filter.Address.VillageCode}})
			}
			if len(filter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": bson.M{"$in": filter.Address.ZoneCode}})
			}
			if len(filter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": bson.M{"$in": filter.Address.WardCode}})
			}
		}

		//Regex Using searchBox Struct
		if filter.SearchBox.MobileNo != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.SearchBox.MobileNo, Options: "xi"}})
		}
		if filter.SearchBox.Ownername != "" {
			query = append(query, bson.M{"ownername": primitive.Regex{Pattern: filter.SearchBox.Ownername, Options: "xi"}})
		}
		if filter.SearchBox.PropertyID != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.SearchBox.PropertyID, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).CountDocuments(ctx.CTX, func() bson.M {
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

	//LookUps

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "propertyId", "uniqueId", "ref.property", "ref.property")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE, "remarkId", "uniqueId", "ref.remark", "ref.remark")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var visitlogs []models.RefPropertyVisitLog
	if err = cursor.All(context.TODO(), &visitlogs); err != nil {
		return nil, err
	}
	return visitlogs, nil
}

// UpdatePropertyVisitLogPropertyID :""
func (d *Daos) UpdatePropertyVisitLogPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYVISITLOG).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
