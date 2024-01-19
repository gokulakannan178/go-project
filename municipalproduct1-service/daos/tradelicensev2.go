package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetSingleTradeLicenseV2 : ""
func (d *Daos) GetSingleTradeLicenseV2(ctx *models.Context, UniqueID string) (*models.RefTradeLicense, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
			"as":   "ref.payments",
			"let":  bson.M{"uniqueId": "$uniqueId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$tradeLicenseId", "$$uniqueId"}},
					{"$eq": []string{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
				}}},
			},
				bson.M{"$sort": bson.M{"financialYear.to": -1}},
				bson.M{
					"$lookup": bson.M{
						"from": constants.COLLECTIONTRADELICENSEPAYMENTSFY,
						"as":   "fys",
						"let":  bson.M{"tnxId": "$tnxId"},
						"pipeline": []bson.M{bson.M{
							"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$tnxId", "$$tnxId"}}}}}, // {"$eq": []string{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},

						},
							bson.M{
								"$lookup": bson.M{
									"from": constants.COLLECTIONTRADELICENSEPAYMENTSBASIC,
									"as":   "basics",
									"let":  bson.M{"tnxId": "$tnxId"},
									"pipeline": []bson.M{bson.M{
										"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$tnxId", "$$tnxId"}}}}}, // {"$eq": []string{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},

									},
									}}},
							bson.M{"$addFields": bson.M{"basics": bson.M{"$arrayElemAt": []interface{}{"$basics", 0}}}},
						}}},
			}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSERATEMASTER,
			"as":   "ref.marketRate",
			"let":  bson.M{"tlctId": "$tlctId", "tlbtId": "$tlbtId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$tlctId", "$$tlctId"}},
					{"$eq": []string{"$tlbtId", "$$tlbtId"}},
					{"$eq": []string{"$status", constants.TRADELICENSERATEMASTERSTATUSACTIVE}},
				}}},
			},
			},
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.marketRate": bson.M{"$arrayElemAt": []interface{}{"$ref.marketRate", 0}}}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	// LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "ref.tradeLicenseCategoryType", "ref.tradeLicenseCategoryType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefTradeLicense
	var tower *models.RefTradeLicense
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}
