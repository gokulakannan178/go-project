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
)

//SaveAdvertisement :""
func (d *Daos) SaveAdvertisement(ctx *models.Context, advertisement *models.Advertisement) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).InsertOne(ctx.CTX, advertisement)
	return err
}

//GetSingleAdvertisement : ""
func (d *Daos) GetSingleAdvertisement(ctx *models.Context, uniqueID string) (*models.RefAdvertisement, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var advertisements []models.RefAdvertisement
	var advertisement *models.RefAdvertisement
	if err = cursor.All(ctx.CTX, &advertisements); err != nil {
		return nil, err
	}
	if len(advertisements) > 0 {
		advertisement = &advertisements[0]
	}
	return advertisement, nil
}

//UpdateAdvertisement : ""
func (d *Daos) UpdateAdvertisement(ctx *models.Context, popup *models.Advertisement) error {
	selector := bson.M{"uniqueId": popup.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": popup}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAdvertisement:""
func (d *Daos) EnableAdvertisement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ADVERTISEMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAdvertisement :""
func (d *Daos) DisableAdvertisement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ADVERTISEMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePopup Notification :""
func (d *Daos) DeleteAdvertisement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ADVERTISEMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAdvertisement : ""
func (d *Daos) FilterAdvertisement(ctx *models.Context, Advertisementfilter *models.AdvertisementFilter, pagination *models.Pagination) ([]models.RefAdvertisement, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Advertisementfilter != nil {
		if len(Advertisementfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Advertisementfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("pkg query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var advertisements []models.RefAdvertisement
	if err = cursor.All(context.TODO(), &advertisements); err != nil {
		return nil, err
	}
	return advertisements, nil
}
