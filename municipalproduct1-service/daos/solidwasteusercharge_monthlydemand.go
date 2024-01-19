package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//CalcShopRentDemand :""
func (d *Daos) CalcSolidWasteUserChargeDemand(ctx *models.Context, mainPipeline []bson.M) (*models.SolidWasteUserChargeDemand, error) {
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)
	d.Shared.BsonToJSONPrintTag("CalcSolidWasteUserChargeDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGE).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var demands []models.SolidWasteUserChargeDemand
	var demand *models.SolidWasteUserChargeDemand
	if err = cursor.All(ctx.CTX, &demands); err != nil {
		return nil, err
	}
	if len(demands) > 0 {
		demand = &demands[0]
	}
	return demand, nil
}
