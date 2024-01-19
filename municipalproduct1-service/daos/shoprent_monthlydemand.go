package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//CalcShopRentDemand :""
func (d *Daos) CalcShopRentMonthlyDemand(ctx *models.Context, mainPipeline []bson.M) (*models.ShopRentMonthlyDemand, error) {
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSHOPRENTRATEMASTER, "uniqueId", "shopRentId", "ref.shopRentRateMaster", "ref.shopRentRateMaster")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTRATEMASTER, "uniqueId", "shopRentId", "ref.shopRentRateMaster", "ref.shopRentRateMaster")...)
	d.Shared.BsonToJSONPrintTag("CalcShopRentDemand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var demands []models.ShopRentMonthlyDemand
	var demand *models.ShopRentMonthlyDemand
	if err = cursor.All(ctx.CTX, &demands); err != nil {
		return nil, err
	}
	if len(demands) > 0 {
		demand = &demands[0]
	}
	return demand, nil
}
