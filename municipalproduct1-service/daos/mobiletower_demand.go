package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//CalcMobileTowerDemand :""
func (d *Daos) CalcMobileTowerDemand(ctx *models.Context, mainPipeline []bson.M) (*models.MobileTowerDemand, error) {
	d.Shared.BsonToJSONPrintTag("CalcMobileTowerDemand query =>", mainPipeline)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainPipeline)
	if err != nil {
		return nil, err
	}
	var mtds []models.MobileTowerDemand
	var mtd *models.MobileTowerDemand
	if err = cursor.All(ctx.CTX, &mtds); err != nil {
		return nil, err
	}
	if len(mtds) > 0 {
		mtd = &mtds[0]
	}
	return mtd, nil
}

func (d *Daos) UpdateMobileTowerDemand(ctx *models.Context, mobileTowerID string, mttd *models.MobileTowerTotalDemand) error {
	selector := bson.M{"uniqueId": mobileTowerID}
	data := bson.M{"$set": bson.M{
		"demand": mttd,
	}}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower demand update resp - ", res)
	return nil
}

func (d *Daos) UpdateMobileTowerCollection(ctx *models.Context, mobileTowerID string, mttc *models.MobileTowerTotalCollection) error {
	selector := bson.M{"uniqueId": mobileTowerID}
	data := bson.M{"$set": bson.M{
		"collection": mttc,
	}}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower demand update resp - ", res)
	return nil
}

func (d *Daos) UpdateMobileTowerOutStanding(ctx *models.Context, mobileTowerID string, mtto *models.MobileTowerTotalOutStanding) error {
	selector := bson.M{"uniqueId": mobileTowerID}
	data := bson.M{"$set": bson.M{
		"collection": mtto,
	}}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower demand update resp - ", res)
	return nil
}

func (d *Daos) UpdateMobileTowerCalc(ctx *models.Context, mtd *models.MobileTowerDemand) error {
	selector := bson.M{"uniqueId": mtd.PropertyMobileTower.UniqueID}
	data := bson.M{"$set": bson.M{
		"demand":             mtd.PropertyMobileTower.Demand,
		"collection":         mtd.PropertyMobileTower.Collections,
		"pendingCollections": mtd.PropertyMobileTower.PendingCollections,
		"outstanding":        mtd.PropertyMobileTower.OutStanding,
	}}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower demand update resp - ", res)
	return nil
}
