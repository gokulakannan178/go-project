package daos

import (
	"context"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ULBNearBy : ""
func (d *Daos) ULBNearBy(ctx *models.Context, ulbnb *models.ULBNearBy, pagination *models.Pagination) ([]models.RefULB, error) {
	coordinater := []float64{ulbnb.Longitude, ulbnb.Latitude}

	mainPipeline := []bson.M{}
	query := []bson.M{{"$geoNear": bson.M{"near": bson.M{"type": "Point", "coordinates": coordinater},
		"maxDistance":   ulbnb.KM * 1000,
		"distanceField": "dist.calculated",
		"spherical":     true,
		"includeLocs":   "dist.location",
	},
	}}
	fmt.Println(query)
	//if d.Shared.GetCmdArg(constants.ENV) != "development" {
	mainPipeline = append(mainPipeline, query...)

	//}
	// if len(ulbnb.CertificateStatus) > 0 {
	//mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"testcert.status": "Active"}})

	// }
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count

		Count, err := GetCountForAggregation(ctx, query, constants.COLLECTIONULB)
		if err != nil {
			log.Println(err)
		}
		pagination.Count = Count
		d.Shared.PaginationData(pagination)
	}
	// LookUps
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Get Inventory
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var ulbs []models.RefULB
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return ulbs, err
}

//UlbInTheState : ""
func (d *Daos) UlbInTheState(ctx *models.Context, stateId string, pagination *models.Pagination) ([]models.RefULB, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{{"address.stateCode": stateId}}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONULB).CountDocuments(ctx.CTX, func() bson.M {
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
	// LookUps
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Get Inventory
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var ulbs []models.RefULB
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return ulbs, err
}

//UlbInTheState : ""
func (d *Daos) UlbInTheStateV2(ctx *models.Context, stateId string, sortBy string, sortorder int, pagination *models.Pagination) ([]models.ULBNearByResponse, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{{"address.stateCode": stateId}}
	//Get Inventory

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	if sortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{sortBy: sortorder}})

	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONULB).CountDocuments(ctx.CTX, func() bson.M {
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
	// LookUps
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)

	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var ulbs []models.ULBNearByResponse
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return ulbs, err
}

//UlbInTheState : ""
func (d *Daos) UlbInTheStateV3(ctx *models.Context, ULBStateIn *models.ULBStateIn, pagination *models.Pagination) ([]models.ULBNearByResponse, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{{"stateID": ULBStateIn.StateID}}
	if len(ULBStateIn.CertificateStatus) > 0 {
		query = append(mainPipeline, bson.M{"ULBStateIn.status": bson.M{"$in": ULBStateIn.CertificateStatus}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONULB).CountDocuments(ctx.CTX, func() bson.M {
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
	// LookUps
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Get Inventory
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var ulbs []models.ULBNearByResponse
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return ulbs, err
}

//UlbCompostInTheState : ""
func (d *Daos) UlbCompostInTheState(ctx *models.Context, stateId string) (*models.CompostInStock, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{{"$match": bson.M{"address.stateCode": stateId}}}
	mainPipeline = append(mainPipeline, query...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.companys", "ref.companys")...)
	group := []bson.M{{"$group": bson.M{"_id": nil, "quantity": bson.M{"$sum": "$ref.companys.quantity"}}}}
	mainPipeline = append(mainPipeline, group...)
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var ulbs []*models.CompostInStock
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return ulbs[0], err
}

func GetCountForAggregation(ctx *models.Context, query []primitive.M, Collection string) (int, error) {
	mainPipeline2 := []bson.M{}

	mainPipeline2 = append(mainPipeline2, query...)

	//Getting Total count
	group := []bson.M{{"$group": bson.M{"_id": nil, "mycount": bson.M{"$sum": 1}}}}
	mainPipeline2 = append(mainPipeline2, group...)
	type countvaule struct {
		Mycount int ` bson:"mycount"`
	}
	var totalCountV []countvaule
	totalCount, err := ctx.DB.Collection(Collection).Aggregate(ctx.CTX, mainPipeline2, nil)
	if err != nil {
		log.Println("Error in geting pagination count")
	}
	if err = totalCount.All(context.TODO(), &totalCountV); err != nil {
		return 0, err
	}
	fmt.Println("count", totalCount)
	if len(totalCountV) > 0 {
		return int(totalCountV[0].Mycount), nil
	}

	return 0, nil

}

//UlbInTheDistrictGps : ""
func (d *Daos) UlbInTheDistrictGps(ctx *models.Context, ulbnb *models.ULBNearBy, pagination *models.Pagination) ([]models.RefULB, error) {
	coordinater := []float64{ulbnb.Longitude, ulbnb.Latitude}
	mainPipeline := []bson.M{}
	query := []bson.M{{"$geoNear": bson.M{"near": bson.M{"type": "Point", "coordinates": coordinater},
		"maxDistance":   ulbnb.KM * 1000,
		"distanceField": "dist.calculated",
		"spherical":     true,
		"includeLocs":   "dist.location",
	},
	}}

	mainPipeline = append(mainPipeline, query...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count

		Count, err := GetCountForAggregation(ctx, query, constants.COLLECTIONULB)
		if err != nil {
			log.Println(err)
		}
		pagination.Count = Count
		d.Shared.PaginationData(pagination)
	}

	// LookUps
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Get Inventory
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var ulbs []models.RefULB
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return ulbs, err
}
