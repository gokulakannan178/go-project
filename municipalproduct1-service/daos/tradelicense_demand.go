package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//CalcTradeLicenseDemand :""
func (d *Daos) CalcTradeLicenseDemand(ctx *models.Context, mainPipeline []bson.M) (*models.TradeLicenseDemand, error) {
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "ref.tradeLicenseCategoryType", "ref.tradeLicenseCategoryType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	d.Shared.BsonToJSONPrintTag("CalcTradeLicenseDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var demands []models.TradeLicenseDemand
	var demand *models.TradeLicenseDemand
	if err = cursor.All(ctx.CTX, &demands); err != nil {
		return nil, err
	}
	if len(demands) > 0 {
		demand = &demands[0]
	}
	return demand, nil
}
