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
)

// SaveTradeLicense : ""
func (d *Daos) SaveTradeLicense(ctx *models.Context, tradeLicense *models.TradeLicense) error {
	d.Shared.BsonToJSONPrint(tradeLicense)
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).InsertOne(ctx.CTX, tradeLicense)
	return err
}

// GetSingleTradeLicense : ""
func (d *Daos) GetSingleTradeLicense(ctx *models.Context, UniqueID string) (*models.RefTradeLicense, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

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

// UpdateTradeLicense : ""
func (d *Daos) UpdateTradeLicense(ctx *models.Context, tradeLicense *models.TradeLicense) error {
	selector := bson.M{"uniqueId": tradeLicense.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": tradeLicense}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableTradeLicense : ""
func (d *Daos) EnableTradeLicense(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableTradeLicense : ""
func (d *Daos) DisableTradeLicense(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteTradeLicense : ""
func (d *Daos) DeleteTradeLicense(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//RejectedTradeLicense :""
func (d *Daos) RejectedTradeLicense(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSESTATUSREJECTED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTradeLicense : ""
func (d *Daos) FilterTradeLicense(ctx *models.Context, filter *models.TradeLicenseFilter, pagination *models.Pagination) ([]models.RefTradeLicense, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}

		if len(filter.TLBTID) > 0 {
			query = append(query, bson.M{"tlbtId": bson.M{"$in": filter.TLBTID}})
		}
		if len(filter.TLCTID) > 0 {
			query = append(query, bson.M{"tlctId": bson.M{"$in": filter.TLCTID}})
		}
		if len(filter.CreatedBy) > 0 {
			query = append(query, bson.M{"created.by": bson.M{"$in": filter.CreatedBy}})
		}
		if len(filter.ApprovedBy) > 0 {
			query = append(query, bson.M{"approved.by": bson.M{"$in": filter.ApprovedBy}})
		}
		if len(filter.NotApprovedBy) > 0 {
			query = append(query, bson.M{"notApproved.by": bson.M{"$in": filter.NotApprovedBy}})
		}
		if len(filter.VerifiedBy) > 0 {
			query = append(query, bson.M{"verified.by": bson.M{"$in": filter.VerifiedBy}})
		}
		if filter.SearchText.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: filter.SearchText.OwnerName, Options: "xi"}})

		}
		if filter.SearchText.MobileNo != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.SearchText.MobileNo, Options: "xi"}})

		}
		if filter.SearchText.GuardianName != "" {
			query = append(query, bson.M{"guardianName": primitive.Regex{Pattern: filter.SearchText.GuardianName, Options: "xi"}})

		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		}
		if filter.SearchText.LisenceNo != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.LisenceNo, Options: "xi"}})
		}
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

	if filter.IsExpired {
		t := time.Now()
		ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		query = append(query, bson.M{"licenseExpiryDate": bson.M{"$lte": ed}})

	}

	if filter.LicenseExpiryDate != nil {
		//var sd,ed time.Time
		if filter.LicenseExpiryDate.From != nil {
			sd := time.Date(filter.LicenseExpiryDate.From.Year(), filter.LicenseExpiryDate.From.Month(), filter.LicenseExpiryDate.From.Day(), 0, 0, 0, 0, filter.LicenseExpiryDate.From.Location())
			ed := time.Date(filter.LicenseExpiryDate.From.Year(), filter.LicenseExpiryDate.From.Month(), filter.LicenseExpiryDate.From.Day(), 23, 59, 59, 0, filter.LicenseExpiryDate.From.Location())
			if filter.LicenseExpiryDate.To != nil {
				ed = time.Date(filter.LicenseExpiryDate.To.Year(), filter.LicenseExpiryDate.To.Month(), filter.LicenseExpiryDate.To.Day(), 23, 59, 59, 0, filter.LicenseExpiryDate.To.Location())
			}
			query = append(query, bson.M{"licenseExpiryDate": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	if filter.LicenseDate != nil {
		//var sd,ed time.Time
		if filter.LicenseDate.From != nil {
			sd := time.Date(filter.LicenseDate.From.Year(), filter.LicenseDate.From.Month(), filter.LicenseDate.From.Day(), 0, 0, 0, 0, filter.LicenseDate.From.Location())
			ed := time.Date(filter.LicenseDate.From.Year(), filter.LicenseDate.From.Month(), filter.LicenseDate.From.Day(), 23, 59, 59, 0, filter.LicenseDate.From.Location())
			if filter.LicenseDate.To != nil {
				ed = time.Date(filter.LicenseDate.To.Year(), filter.LicenseDate.To.Month(), filter.LicenseDate.To.Day(), 23, 59, 59, 0, filter.LicenseDate.To.Location())
			}
			query = append(query, bson.M{"licenseDate": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	if filter.CreatedDateRange != nil {
		//var sd,ed time.Time
		if filter.CreatedDateRange.From != nil {
			sd := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 0, 0, 0, 0, filter.CreatedDateRange.From.Location())
			ed := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 23, 59, 59, 0, filter.CreatedDateRange.From.Location())
			if filter.CreatedDateRange.To != nil {
				ed = time.Date(filter.CreatedDateRange.To.Year(), filter.CreatedDateRange.To.Month(), filter.CreatedDateRange.To.Day(), 23, 59, 59, 0, filter.CreatedDateRange.To.Location())
			}
			query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

		}
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "ref.tradeLicenseCategoryType", "ref.tradeLicenseCategoryType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refTradeLicense []models.RefTradeLicense
	if err = cursor.All(context.TODO(), &refTradeLicense); err != nil {
		return nil, err
	}
	return refTradeLicense, nil
}

// VerifyTradeLicensePayment : ""
func (d *Daos) VerifyTradeLicensePayment(ctx *models.Context, action *models.MakeTradeLicensePaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.TradeLicensePaymentsAction,
			"status":       constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment basic resp - ", res)
	return nil
}

// NotVerifyTradeLicensePayment : ""
func (d *Daos) NotVerifyTradeLicensePayment(ctx *models.Context, action *models.MakeTradeLicensePaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.TradeLicensePaymentsAction,
			"status":       constants.TRADELICENSEPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment basic resp - ", res)
	return nil
}

// RejectTradeLicensePayment : ""
func (d *Daos) RejectTradeLicensePayment(ctx *models.Context, action *models.MakeTradeLicensePaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.TradeLicensePaymentsAction,
			"status":       constants.TRADELICENSEPAYMENRSTATUSREJECTED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.TRADELICENSEPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Trade license payment basic resp - ", res)
	return nil
}

func (d *Daos) UpdateTradeLicenseCalc(ctx *models.Context, demand *models.TradeLicenseDemand) error {
	selector := bson.M{"uniqueId": demand.TradeLicense.UniqueID}
	data := bson.M{"$set": bson.M{
		"demand":             demand.TradeLicense.Demand,
		"collection":         demand.TradeLicense.Collections,
		"pendingCollections": demand.TradeLicense.PendingCollections,
		"outstanding":        demand.TradeLicense.OutStanding,
	}}
	d.Shared.BsonToJSONPrintTag("UpdateTradeLicenseCalc selector - ", selector)

	res, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		return err
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower demand update resp - ", res)

	return nil
}

// ApproveTradeLicense : ""
func (d *Daos) ApproveTradeLicense(ctx *models.Context, accept *models.ApproveTradeLicense) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSESTATUSACTIVE, "esign": accept.ESign,
		"approved": models.Action{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// NotApproveTradeLicense : ""
func (d *Daos) NotApproveTradeLicense(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	t := time.Now()

	query := bson.M{"uniqueId": notApprove.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSESTATUSNOTAPPROVED,
		"notApproved": models.Updated{
			On:      &t,
			By:      notApprove.UserName,
			ByType:  notApprove.UserType,
			Remarks: notApprove.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetTradeLicenseSAFDashboard(ctx *models.Context, filter *models.GetTradeLicenseSAFDashboardFilter) (*models.TradeLicenseSAFDashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"init": []bson.M{
			bson.M{"$match": bson.M{"status": "Init"}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
		"pending": []bson.M{
			bson.M{"$match": bson.M{"status": "Pending"}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
		"active": []bson.M{
			bson.M{"$match": bson.M{"status": "Active"}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
		"rejected": []bson.M{
			bson.M{"$match": bson.M{"status": "NotApproved"}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
		"expired": []bson.M{
			bson.M{"$match": bson.M{
				"status":            "Active",
				"licenseExpiryDate": bson.M{"$lte": time.Now()},
			}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
	}},
		bson.M{
			"$addFields": bson.M{
				"init":     bson.M{"$arrayElemAt": []interface{}{"$init", 0}},
				"pending":  bson.M{"$arrayElemAt": []interface{}{"$pending", 0}},
				"active":   bson.M{"$arrayElemAt": []interface{}{"$active", 0}},
				"rejected": bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}},
				"expired":  bson.M{"$arrayElemAt": []interface{}{"$expired", 0}},
			},
		},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetTradeLicenseSAFDashboard query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var results []models.TradeLicenseSAFDashboard
	var result *models.TradeLicenseSAFDashboard
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	if len(results) > 0 {
		result = &results[0]
	}
	return result, nil
}

// ApproveTradeLicense : ""
func (d *Daos) VerifyTradeLicense(ctx *models.Context, accept *models.ApproveTradeLicense) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSESTATUSVERIFIED,
		"verified": models.Action{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
