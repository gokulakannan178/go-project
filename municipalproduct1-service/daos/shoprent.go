package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveShopRent :""
func (d *Daos) SaveShopRent(ctx *models.Context, shoprent *models.ShopRent) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).InsertOne(ctx.CTX, shoprent)
	return err
}

//GetSingleShopRent : ""
func (d *Daos) GetSingleShopRent(ctx *models.Context, UniqueID string) (*models.RefShopRent, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Lookups
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shoprents []models.RefShopRent
	var shoprent *models.RefShopRent
	if err = cursor.All(ctx.CTX, &shoprents); err != nil {
		return nil, err
	}
	if len(shoprents) > 0 {
		shoprent = &shoprents[0]
	}
	return shoprent, nil
}

//UpdateShopRent : ""
func (d *Daos) UpdateShopRent(ctx *models.Context, shoprent *models.ShopRent) error {
	selector := bson.M{"uniqueId": shoprent.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": shoprent, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// FilterShopRent : ""
func (d *Daos) FilterShopRent(ctx *models.Context, filter *models.ShopRentFilter, pagination *models.Pagination) ([]models.RefShopRent, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}

		if len(filter.ShopCategoryID) > 0 {
			query = append(query, bson.M{"shopCategoryId": bson.M{"$in": filter.ShopCategoryID}})
		}
		if len(filter.ShopSubCategoryID) > 0 {
			query = append(query, bson.M{"shopSubCategoryId": bson.M{"$in": filter.ShopSubCategoryID}})
		}
		if filter.SearchText.MobileNo != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.SearchText.MobileNo, Options: "xi"}})
		}
		if filter.SearchText.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.SearchText.OwnerName, Options: "xi"}})
		}

		if filter.SearchText.GuardianName != "" {
			query = append(query, bson.M{"guardianName": primitive.Regex{Pattern: filter.SearchText.GuardianName, Options: "xi"}})

		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		}
		if filter.Address != nil {
			if len(filter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": filter.Address.StateCode}})
			}
			if len(filter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": filter.Address.DistrictCode}})
			}
			if len(filter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": filter.Address.VillageCode}})
			}
			if len(filter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": bson.M{"$in": filter.Address.ZoneCode}})
			}
			if len(filter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": bson.M{"$in": filter.Address.WardCode}})
			}
		}

		if filter.FromDateRange != nil {
			//var sd,ed time.Time
			if filter.FromDateRange.From != nil {
				sd := time.Date(filter.FromDateRange.From.Year(), filter.FromDateRange.From.Month(), filter.FromDateRange.From.Day(), 0, 0, 0, 0, filter.FromDateRange.From.Location())
				ed := time.Date(filter.FromDateRange.From.Year(), filter.FromDateRange.From.Month(), filter.FromDateRange.From.Day(), 23, 59, 59, 0, filter.FromDateRange.From.Location())
				if filter.FromDateRange.To != nil {
					ed = time.Date(filter.FromDateRange.To.Year(), filter.FromDateRange.To.Month(), filter.FromDateRange.To.Day(), 23, 59, 59, 0, filter.FromDateRange.To.Location())
				}
				query = append(query, bson.M{"dateFrom": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.ToDateRange != nil {
			//var sd,ed time.Time
			if filter.ToDateRange.From != nil {
				sd := time.Date(filter.ToDateRange.From.Year(), filter.ToDateRange.From.Month(), filter.ToDateRange.From.Day(), 0, 0, 0, 0, filter.ToDateRange.From.Location())
				ed := time.Date(filter.ToDateRange.From.Year(), filter.ToDateRange.From.Month(), filter.ToDateRange.From.Day(), 23, 59, 59, 0, filter.ToDateRange.From.Location())
				if filter.ToDateRange.To != nil {
					ed = time.Date(filter.ToDateRange.To.Year(), filter.ToDateRange.To.Month(), filter.ToDateRange.To.Day(), 23, 59, 59, 0, filter.ToDateRange.To.Location())
				}
				query = append(query, bson.M{"dateTo": bson.M{"$gte": sd, "$lte": ed}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).CountDocuments(ctx.CTX, func() bson.M {
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

	// LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refShopRent []models.RefShopRent
	if err = cursor.All(context.TODO(), &refShopRent); err != nil {
		return nil, err
	}
	return refShopRent, nil
}

//EnableShopRent :""
func (d *Daos) EnableShopRent(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//RejectedShopRent :""
func (d *Daos) RejectedShopRent(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSTATUSREJECTED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableShopRent :""
func (d *Daos) DisableShopRent(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteShopRent :""
func (d *Daos) DeleteShopRent(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// VerifyShopRentPayment : ""
func (d *Daos) VerifyShopRentPayment(ctx *models.Context, action *models.MakeShopRentPaymentsAction) (string, error) {
	fmt.Println("action.TnxID:", action.TnxID)
	payment, err := d.GetSingleShopRentPayment(ctx, action.TnxID)
	if err != nil {
		return "", err
	}

	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.ShopRentPaymentsAction,
			"status":       constants.SHOPRENTPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.SHOPRENTPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.SHOPRENTPAYMENTSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}

	d.Shared.BsonToJSONPrintTag("Shop rent payment basic resp - ", res)
	return payment.ShopRentID, nil
}

// NotVerifyShopRentPayment : ""
func (d *Daos) NotVerifyShopRentPayment(ctx *models.Context, action *models.MakeShopRentPaymentsAction) (string, error) {
	payment, err := d.GetSingleShopRentPayment(ctx, action.TnxID)
	if err != nil {
		return "", err
	}
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.ShopRentPaymentsAction,
			"status":       constants.SHOPRENTPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.SHOPRENTPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.SHOPRENTPAYMENTSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment basic resp - ", res)
	return payment.ShopRentID, nil
}

// RejectShopRentPayment : ""
func (d *Daos) RejectShopRentPayment(ctx *models.Context, action *models.MakeShopRentPaymentsAction) (string, error) {
	payment, err := d.GetSingleShopRentPayment(ctx, action.TnxID)
	if err != nil {
		return "", err
	}
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.ShopRentPaymentsAction,
			"status":       constants.SHOPRENTPAYMENTSTATUSREJECTED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.SHOPRENTPAYMENTSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.SHOPRENTPAYMENTSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return "", errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Shop rent payment basic resp - ", res)
	return payment.ShopRentID, nil
}

//UpdateShopRentCalc : ""
/*
* Commented demand - by solomon 17-Nov-2021
* Reason - because demand must be updated overall only during adding and during reassessent
 */
func (d *Daos) UpdateShopRentCalc(ctx *models.Context, demand *models.ShopRentDemand) error {
	selector := bson.M{"uniqueId": demand.ShopRent.UniqueID}
	data := bson.M{"$set": bson.M{
		//"demand":             demand.ShopRent.Demand,
		"collection":         demand.ShopRent.Collections,
		"pendingCollections": demand.ShopRent.PendingCollections,
		"outstanding":        demand.ShopRent.OutStanding,
	}}

	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower demand update resp - ", res)

	return nil
}

func (d *Daos) UpdateOverallShopRentDemand(ctx *models.Context, shoprentID string, demand *models.ShopRentTotalDemand) error {
	selector := bson.M{"uniqueId": shoprentID}
	data := bson.M{"$set": bson.M{
		"demand": demand,
	}}

	res, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower demand update resp - ", res)

	return nil
}

// GetAllShopRentIds : ""
func (d *Daos) GetAllShopRentIds(ctx *models.Context) ([]string, error) {
	t := []struct{ UniqueID string }{}
	query := bson.M{"status": bson.M{"$in": []string{"Active", "Init"}}}
	projection := bson.M{"_id": 0, "uniqueId": 1}
	opts := options.Find()
	opts.SetProjection(projection)
	// opts.SetSkip(0)
	// opts.SetLimit(500)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Find(ctx.CTX, query, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx.CTX, &t)
	if err != nil {
		return nil, err
	}
	data := []string{}
	for _, v := range t {
		data = append(data, v.UniqueID)
	}
	return data, nil
}
