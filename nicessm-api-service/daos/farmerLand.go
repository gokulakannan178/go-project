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

//SaveFarmerLand :""
func (d *Daos) SaveFarmerLand(ctx *models.Context, FarmerLand *models.FarmerLand) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMERLAND).InsertOne(ctx.CTX, FarmerLand)
	if err != nil {
		return err
	}
	FarmerLand.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleFarmerLand : ""
func (d *Daos) GetSingleFarmerLand(ctx *models.Context, code string) (*models.RefFarmerLand, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "ref.gramPanchayat.block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.block.district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERLAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var FarmerLands []models.RefFarmerLand
	var FarmerLand *models.RefFarmerLand
	if err = cursor.All(ctx.CTX, &FarmerLands); err != nil {
		return nil, err
	}
	if len(FarmerLands) > 0 {
		FarmerLand = &FarmerLands[0]
	}
	return FarmerLand, nil
}

//UpdateFarmerLand : ""
func (d *Daos) UpdateFarmerLand(ctx *models.Context, FarmerLand *models.FarmerLand) error {
	selector := bson.M{"_id": FarmerLand.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": FarmerLand, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERLAND).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFarmerLand : ""
func (d *Daos) FilterFarmerLand(ctx *models.Context, FarmerLandfilter *models.FarmerLandFilter, pagination *models.Pagination) ([]models.RefFarmerLand, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if FarmerLandfilter != nil {

		if len(FarmerLandfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": FarmerLandfilter.ActiveStatus}})
		}
		if len(FarmerLandfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": FarmerLandfilter.Status}})
		}
		if len(FarmerLandfilter.CultivationPractice) > 0 {
			query = append(query, bson.M{"cultivationPractice": bson.M{"$in": FarmerLandfilter.CultivationPractice}})
		}
		if len(FarmerLandfilter.IrrigationType) > 0 {
			query = append(query, bson.M{"irrigationType": bson.M{"$in": FarmerLandfilter.IrrigationType}})
		}
		if len(FarmerLandfilter.LandPosition) > 0 {
			query = append(query, bson.M{"landPosition": bson.M{"$in": FarmerLandfilter.LandPosition}})
		}
		if len(FarmerLandfilter.LandType) > 0 {
			query = append(query, bson.M{"landType": bson.M{"$in": FarmerLandfilter.LandType}})
		}
		if len(FarmerLandfilter.OwnerShip) > 0 {
			query = append(query, bson.M{"ownerShip": bson.M{"$in": FarmerLandfilter.OwnerShip}})
		}
		if len(FarmerLandfilter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": FarmerLandfilter.Village}})
		}
		if len(FarmerLandfilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": FarmerLandfilter.GramPanchayat}})
		}
		if len(FarmerLandfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": FarmerLandfilter.Block}})
		}
		if len(FarmerLandfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": FarmerLandfilter.District}})
		}
		if len(FarmerLandfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": FarmerLandfilter.State}})
		}
		if len(FarmerLandfilter.Farmer) > 0 {
			query = append(query, bson.M{"farmer": bson.M{"$in": FarmerLandfilter.Farmer}})
		}
		//Regex
		if FarmerLandfilter.Regex.Type != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: FarmerLandfilter.Regex.Type, Options: "xi"}})
		}

		if FarmerLandfilter.Regex.ParcelNumber != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: FarmerLandfilter.Regex.ParcelNumber, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFARMERLAND).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "soilType", "_id", "ref.soilType", "ref.soilType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLANDCROP, "_id", "farmerLand", "ref.landCrop", "ref.landCrop")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "ref.landCrop.crop", "_id", "ref.crop", "ref.crop")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("FarmerLand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERLAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var FarmerLands []models.RefFarmerLand
	if err = cursor.All(context.TODO(), &FarmerLands); err != nil {
		return nil, err
	}
	return FarmerLands, nil
}

//EnableFarmerLand :""
func (d *Daos) EnableFarmerLand(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERLANDSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERLAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmerLand :""
func (d *Daos) DisableFarmerLand(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERLANDSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERLAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmerLand :""
func (d *Daos) DeleteFarmerLand(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERLANDSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERLAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) SumFarmerLandArea(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := []bson.M{}
	query = append(query, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

		{"$eq": []string{"$status", "Active"}},
		{"$eq": []interface{}{"$farmer", id}},
	}}}})
	query = append(query, bson.M{"$group": bson.M{
		"_id":            nil,
		"totalLand":      bson.M{"$sum": "$areaInAcre"},
		"cultivatedArea": bson.M{"$sum": "$cultivatedArea"},
		"vacantArea":     bson.M{"$sum": "$vacantArea"},
	}})
	d.Shared.BsonToJSONPrintTag("Farmerland Sum query =>", query)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERLAND).Aggregate(ctx.CTX, query, nil)
	if err != nil {
		return err
	}
	var lands []models.Land
	var land *models.Land
	if err = cursor.All(context.TODO(), &lands); err != nil {
		return err
	}
	if len(lands) > 0 {
		land = &lands[0]
	} else {
		land = new(models.Land)
	}
	query2 := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"totalLand": land.TotalLand, "cultivatedArea": land.CultivatedArea, "vacantArea": land.VacantArea}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query2, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return nil
}
