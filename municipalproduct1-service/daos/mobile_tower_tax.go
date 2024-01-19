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
)

// SaveMobileTowerTax : ""
func (d *Daos) SaveMobileTowerTax(ctx *models.Context, mobile *models.MobileTowerTax) error {
	d.Shared.BsonToJSONPrint(mobile)
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERTAX).InsertOne(ctx.CTX, mobile)
	return err
}

// GetSingleMobileTowerTax : ""
func (d *Daos) GetSingleMobileTowerTax(ctx *models.Context, UniqueID string) (*models.MobileTowerTax, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("mobile tower query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.MobileTowerTax
	var tower *models.MobileTowerTax
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateMobileTowerTax : ""
func (d *Daos) UpdateMobileTowerTax(ctx *models.Context, mobile *models.MobileTowerTax) error {
	selector := bson.M{"uniqueId": mobile.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": mobile}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERTAX).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMobileTowerTax : ""
func (d *Daos) EnableMobileTowerTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERTAXSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableMobileTowerTax : ""
func (d *Daos) DisableMobileTowerTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERTAXSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteMobileTowerTax : ""
func (d *Daos) DeleteMobileTowerTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERTAXSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMobileTowerTax : ""
func (d *Daos) FilterMobileTowerTax(ctx *models.Context, filter *models.MobileTowerTaxFilter, pagination *models.Pagination) ([]models.RefMobileTowerTax, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERTAX).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("mobile tower query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefMobileTowerTax
	if err = cursor.All(context.TODO(), &towers); err != nil {
		return nil, err
	}
	return towers, nil
}
