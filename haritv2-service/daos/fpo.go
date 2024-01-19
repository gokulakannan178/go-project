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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveFPO : ""
func (d *Daos) SaveFPO(ctx *models.Context, fpo *models.FPO) error {
	d.Shared.BsonToJSONPrint(fpo)
	_, err := ctx.DB.Collection(constants.COLLECTIONFPO).InsertOne(ctx.CTX, fpo)
	return err
}

// GetSingleFPO : ""
func (d *Daos) GetSingleFPO(ctx *models.Context, UniqueID string) (*models.RefFPO, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "mobile", "mobile", "ref.chairman", "ref.chairman")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefFPO
	var tower *models.RefFPO
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	fmt.Println(tower)
	return tower, nil

}

// UpdateFPO : ""
func (d *Daos) UpdateFPO(ctx *models.Context, business *models.FPO) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPO).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableFPO : ""
func (d *Daos) EnableFPO(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FPOSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableFPO : ""
func (d *Daos) DisableFPO(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FPOSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteFPO : ""
func (d *Daos) DeleteFPO(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FPOSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterFPO : ""
func (d *Daos) FilterFPO(ctx *models.Context, filter *models.FPOFilter, pagination *models.Pagination) ([]models.RefFPO, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}

	//Regex
	if filter.Regex.Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
	}
	if filter.Regex.Email != "" {
		query = append(query, bson.M{"email": primitive.Regex{Pattern: filter.Regex.Email, Options: "xi"}})
	}
	if filter.Regex.Mobile != "" {
		query = append(query, bson.M{"mobile": primitive.Regex{Pattern: filter.Regex.Mobile, Options: "xi"}})
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFPO).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var fpo []models.RefFPO
	if err = cursor.All(context.TODO(), &fpo); err != nil {
		return nil, err
	}
	return fpo, nil
}

// FPOMasterReport : ""
func (d *Daos) FPOMasterReport(ctx *models.Context, filter *models.FPOReportFilter, pagination *models.Pagination) ([]models.FPOReport, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if filter.Address != nil {
			if filter.Address.StateCode != "" {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": []string{filter.Address.StateCode}}})
			}
			if filter.Address.DistrictCode != "" {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": []string{filter.Address.DistrictCode}}})
			}
			if filter.Address.BlockCode != "" {
				query = append(query, bson.M{"address.blockCode": bson.M{"$in": []string{filter.Address.BlockCode}}})
			}
			if filter.Address.GramPanchayatCode != "" {
				query = append(query, bson.M{"address.gramPanchayatCode": bson.M{"$in": []string{filter.Address.GramPanchayatCode}}})
			}
			if filter.Address.VillageCode != "" {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": []string{filter.Address.VillageCode}}})
			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFPO).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "sale",
		"as":   "compostPurchasedTillDate",
		"let":  bson.M{"fpoId": "$uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{
				"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$customer.type", "FPO"}},
					{"$eq": []string{"$customer.id", "$$fpoId"}},
					{"$eq": []string{"$transport.status", "Delivered"}},
				}},
			}},
			{"$addFields": bson.M{"quantity": bson.M{"$sum": "$items.quantity"}}},
			// bson.M{"$group": bson.M{
			// 	"_id": nil, "totalPurchase": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$totalAmount"}, "quantity": bson.M{"$sum": "$quantity"},
			// }},

			{"$group": bson.M{
				"_id": bson.M{"companyType": "$company.type", "companyId": "$company.id"}, "totalPurchase": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$totalAmount"}, "quantity": bson.M{"$sum": "$quantity"},
			}},

			{"$group": bson.M{
				"_id": bson.M{"companyType": "$_id.companyType"}, "companyCount": bson.M{"$sum": 1}, "totalPurchase": bson.M{"$sum": "$totalPurchase"}, "amount": bson.M{"$sum": "$amount"}, "quantity": bson.M{"$sum": "$quantity"},
			}},

			{"$group": bson.M{
				"_id": nil, "totalPurchase": bson.M{"$sum": "$totalPurchase"}, "amount": bson.M{"$sum": "$amount"}, "quantity": bson.M{"$sum": "$quantity"},
				"ulb":      bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "ULB"}}, "then": "$companyCount", "else": 0}}},
				"fpo":      bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "FPO"}}, "then": "$companyCount", "else": 0}}},
				"customer": bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "Customer"}}, "then": "$companyCount", "else": 0}}},
				"self":     bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "Self"}}, "then": "$companyCount", "else": 0}}},
			}},
		},
	}})
	if filter.Date != nil {
		sd := d.Shared.BeginningOfMonth(*filter.Date)
		ed := d.Shared.EndOfMonth(*filter.Date)
		mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
			"from": "sale",
			"as":   "compostPurchasedCurrMonth",
			"let":  bson.M{"fpoId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{
					"$expr": bson.M{"$and": []bson.M{
						{"$eq": []string{"$customer.type", "FPO"}},
						{"$eq": []string{"$customer.id", "$$fpoId"}},
						{"$eq": []string{"$transport.status", "Delivered"}},
						{"$gte": []interface{}{"$createdOn.on", sd}},
						{"$lte": []interface{}{"$createdOn.on", ed}},
					}},
				}},
				{"$addFields": bson.M{"quantity": bson.M{"$sum": "$items.quantity"}}},
				// bson.M{"$group": bson.M{
				// 	"_id": nil, "totalPurchase": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$totalAmount"}, "quantity": bson.M{"$sum": "$quantity"},
				// }},
				{"$group": bson.M{
					"_id": bson.M{"companyType": "$company.type", "companyId": "$company.id"}, "totalPurchase": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$totalAmount"}, "quantity": bson.M{"$sum": "$quantity"},
				}},

				{"$group": bson.M{
					"_id": bson.M{"companyType": "$_id.companyType"}, "companyCount": bson.M{"$sum": 1}, "totalPurchase": bson.M{"$sum": "$totalPurchase"}, "amount": bson.M{"$sum": "$amount"}, "quantity": bson.M{"$sum": "$quantity"},
				}},

				{"$group": bson.M{
					"_id": nil, "totalPurchase": bson.M{"$sum": "$totalPurchase"}, "amount": bson.M{"$sum": "$amount"}, "quantity": bson.M{"$sum": "$quantity"},
					"ulb":      bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "ULB"}}, "then": "$companyCount", "else": 0}}},
					"fpo":      bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "FPO"}}, "then": "$companyCount", "else": 0}}},
					"customer": bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "Customer"}}, "then": "$companyCount", "else": 0}}},
					"self":     bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "Self"}}, "then": "$companyCount", "else": 0}}},
				}},
			},
		}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "orders",
		"as":   "pendingOrders",
		"let":  bson.M{"fpoId": "$uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{
				"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$customer.type", "FPO"}},
					bson.M{"$eq": []string{"$customer.id", "$$fpoId"}},
					bson.M{"$eq": []string{"$status", "Placed"}},
				},
				}}},
			{"$group": bson.M{
				"_id": bson.M{"companyType": "$company.type", "companyId": "$company.id"}, "companyNames": bson.M{"$first": "$company.name"}, "totalPurchase": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$totalAmount"}, "quantity": bson.M{"$sum": "$quantity"},
			}},
			{"$group": bson.M{
				"_id": bson.M{"companyType": "$_id.companyType"}, "companyCount": bson.M{"$sum": 1}, "totalPurchase": bson.M{"$sum": "$totalPurchase"}, "amount": bson.M{"$sum": "$amount"}, "quantity": bson.M{"$sum": "$quantity"},
				"companyNames": bson.M{"$push": "$companyNames"},
			}},
			{"$group": bson.M{
				"_id": nil, "totalPurchase": bson.M{"$sum": "$totalPurchase"}, "amount": bson.M{"$sum": "$amount"}, "quantity": bson.M{"$sum": "$quantity"},
				"ulb": bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "ULB"}}, "then": "$companyCount", "else": 0}}},

				"fpo":      bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "FPO"}}, "then": "$companyCount", "else": 0}}},
				"customer": bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "Customer"}}, "then": "$companyCount", "else": 0}}},
				"self":     bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.companyType", "Self"}}, "then": "$companyCount", "else": 0}}},
				//             "ulbNames":{"$push":{"$cond":{"if":{"$eq":["$_id.companyType","ULB"]},"then":"$companyNames","else":0}}},
				//             "fpoNames":{"$push":{"$cond":{"if":{"$eq":["$_id.companyType","FPO"]},"then":"$companyNames","else":0}}},

			}},
		}},
	},
	)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"compostPurchasedTillDate":  bson.M{"$arrayElemAt": []interface{}{"$compostPurchasedTillDate", 0}},
		"compostPurchasedCurrMonth": bson.M{"$arrayElemAt": []interface{}{"$compostPurchasedCurrMonth", 0}},
		"pendingOrders":             bson.M{"$arrayElemAt": []interface{}{"$pendingOrders", 0}},
	}})
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "ref.address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "ref.address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "ref.address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)

	// Aggregation
	d.Shared.BsonToJSONPrintTag("mainpipeline query =>", mainPipeline)
	var data []models.FPOReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil

}

// FPOMasterReport : ""
func (d *Daos) FPOMonthReport(ctx *models.Context, filter *models.FPOMothWiseeportFilter) ([]models.FPOMothWiseeport, error) {
	var sd time.Time
	sd = time.Date(filter.Year, filter.Month, 01, 0, 0, 0, 0, sd.Location())
	ed := d.Shared.EndOfMonth(sd)
	mainPipeline := []bson.M{}
	query := []bson.M{}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query, "company.type": "FPO"}})
	}

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONSALE,
			"as":   "sales",
			"let":  bson.M{"uniqueId": "$uniqueId", "months": filter.Month, "year": filter.Year},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{{"$eq": []string{"$company.id", "$$uniqueId"}},
					{"$eq": []string{"$company.type", "FPO"}},
					{"$gte": []interface{}{"$createdOn.on", sd}},
					{"$lte": []interface{}{"$createdOn.on", ed}},
				}}}},
				{"$group": bson.M{"_id": nil, "totalsaleAmount": bson.M{"$sum": "$totalAmount"}, "NoOfCustomers": bson.M{"$push": "$customer.id"},
					"totalCustomers": bson.M{"$sum": 1}}},
			}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"sales": bson.M{"$arrayElemAt": []interface{}{"$sales", 0}}}})
	// Lookup

	// Aggregation
	d.Shared.BsonToJSONPrintTag("FPOmonthwisereport query =>", mainPipeline)
	var data []models.FPOMothWiseeport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil

}

func (d *Daos) FBONearBy(ctx *models.Context, fbonb *models.FBONearBy, pagination *models.Pagination) ([]models.RefFPO, error) {
	coordinater := []float64{fbonb.Longitude, fbonb.Latitude}

	mainPipeline := []bson.M{}
	query := []bson.M{{"$geoNear": bson.M{"near": bson.M{"type": "Point", "coordinates": coordinater},
		"maxDistance":   fbonb.KM * 1000,
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

		Count, err := GetCountForAggregation(ctx, query, constants.COLLECTIONFPO)
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFPOINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var fpos []models.RefFPO
	if err = cursor.All(context.TODO(), &fpos); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return fpos, err
}
