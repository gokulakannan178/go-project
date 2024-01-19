package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//CalcUserChargeDemand :""
func (d *Daos) CalcUserChargeMonthlyDemand(ctx *models.Context, mainPipeline []bson.M) (*models.UserChargeMonthlyDemand, error) {
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUserChargeSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.UserChargeShopCategory", "ref.UserChargeShopCategory")...)
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUserChargeSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.UserChargeShopSubCategory", "ref.UserChargeShopSubCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONUserChargeRATEMASTER, "uniqueId", "UserChargeId", "ref.UserChargeRateMaster", "ref.UserChargeRateMaster")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUserChargeRATEMASTER, "uniqueId", "UserChargeId", "ref.UserChargeRateMaster", "ref.UserChargeRateMaster")...)
	d.Shared.BsonToJSONPrintTag("CalcUserChargeDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var demands []models.UserChargeMonthlyDemand
	var demand *models.UserChargeMonthlyDemand
	if err = cursor.All(ctx.CTX, &demands); err != nil {
		return nil, err
	}
	if len(demands) > 0 {
		demand = &demands[0]
	}
	return demand, nil
}
