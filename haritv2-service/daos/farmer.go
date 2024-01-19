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

//SaveFarmer :""
func (d *Daos) SaveFarmer(ctx *models.Context, farmer *models.Farmer) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).InsertOne(ctx.CTX, farmer)
	return err
}

//GetSingleFarmer : ""
func (d *Daos) GetSingleFarmer(ctx *models.Context, uniqueID string) (*models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.RefFarmer
	var farmer *models.RefFarmer
	if err = cursor.All(ctx.CTX, &farmers); err != nil {
		return nil, err
	}
	if len(farmers) > 0 {
		farmer = &farmers[0]
	}
	return farmer, nil
}

//GetSingleFarmer : ""
func (d *Daos) GetSingleFarmerWithMobileNo(ctx *models.Context, mobileNo string) (*models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNo": mobileNo}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.RefFarmer
	var farmer *models.RefFarmer
	if err = cursor.All(ctx.CTX, &farmers); err != nil {
		return nil, err
	}
	if len(farmers) > 0 {
		farmer = &farmers[0]
	}
	return farmer, nil
}

//UpdateFarmer : ""
func (d *Daos) UpdateFarmer(ctx *models.Context, farmer *models.Farmer) error {
	selector := bson.M{"uniqueId": farmer.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": farmer}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableFarmer :""
func (d *Daos) EnableFarmer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmer :""
func (d *Daos) DisableFarmer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmer :""
func (d *Daos) DeleteFarmer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterFarmer : ""
func (d *Daos) FilterFarmer(ctx *models.Context, farmerfilter *models.FarmerFilter, pagination *models.Pagination) ([]models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if farmerfilter != nil {
		if len(farmerfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": farmerfilter.UniqueID}})
		}
		if len(farmerfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": farmerfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &farmers); err != nil {
		return nil, err
	}
	return farmers, nil
}

//GetSingleUserWithMobileNo : ""
func (d *Daos) GetSingleUserWithMobileNoForFarmer(ctx *models.Context, mobileno string) (*models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNo": mobileno}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.fpo.address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.fpo.address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "ref.fpo.address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "ref.fpo.address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.fpo.address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.RefFarmer
	var farmer *models.RefFarmer
	if err = cursor.All(ctx.CTX, &farmers); err != nil {
		return nil, err
	}
	if len(farmers) > 0 {
		farmer = &farmers[0]
	}
	return farmer, nil
}
