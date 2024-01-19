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

//SaveVacantLandRate :""
func (d *Daos) SaveVacantLandRate(ctx *models.Context, vacantLandRate *models.VacantLandRate) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).InsertOne(ctx.CTX, vacantLandRate)
	return err
}

//GetSingleVacantLandRate : ""
func (d *Daos) GetSingleVacantLandRate(ctx *models.Context, UniqueID string) (*models.RefVacantLandRate, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityTypeId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vacantLandRates []models.RefVacantLandRate
	var vacantLandRate *models.RefVacantLandRate
	if err = cursor.All(ctx.CTX, &vacantLandRates); err != nil {
		return nil, err
	}
	if len(vacantLandRates) > 0 {
		vacantLandRate = &vacantLandRates[0]
	}
	return vacantLandRate, nil
}

//UpdateVacantLandRate : ""
func (d *Daos) UpdateVacantLandRate(ctx *models.Context, vacantLandRate *models.VacantLandRate) error {
	selector := bson.M{"uniqueId": vacantLandRate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": vacantLandRate, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterVacantLandRate : ""
func (d *Daos) FilterVacantLandRate(ctx *models.Context, vacantLandRatefilter *models.VacantLandRateFilter, pagination *models.Pagination) ([]models.RefVacantLandRate, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if vacantLandRatefilter != nil {

		if len(vacantLandRatefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": vacantLandRatefilter.Status}})
		}
		if len(vacantLandRatefilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": vacantLandRatefilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityTypeId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("vacantLandRate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vacantLandRates []models.RefVacantLandRate
	if err = cursor.All(context.TODO(), &vacantLandRates); err != nil {
		return nil, err
	}
	return vacantLandRates, nil
}

//EnableVacantLandRate :""
func (d *Daos) EnableVacantLandRate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVACANTLANDRATESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableVacantLandRate :""
func (d *Daos) DisableVacantLandRate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVACANTLANDRATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVacantLandRate :""
func (d *Daos) DeleteVacantLandRate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYVACANTLANDRATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVACANTLANDRATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
