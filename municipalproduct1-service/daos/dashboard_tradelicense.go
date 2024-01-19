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

// SaveTradeLicenseDashboard : ""
func (d *Daos) SaveTradeLicenseDashboard(ctx *models.Context, tradeLicense *models.TradeLicenseDashboard) error {
	d.Shared.BsonToJSONPrint(tradeLicense)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).InsertOne(ctx.CTX, tradeLicense)
	return err
}

// GetSingleDashBoardProperty  : ""
func (d *Daos) GetSingleTradeLicenseDashboard(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseDashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefTradeLicenseDashboard
	var tower *models.RefTradeLicenseDashboard
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateDashBoardProperty: ""
func (d *Daos) UpdateTradeLicenseDashboard(ctx *models.Context, tradeLicense *models.TradeLicenseDashboard) error {
	selector := bson.M{"uniqueId": tradeLicense.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": tradeLicense}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableDashBoardProperty : ""
func (d *Daos) EnableTradeLicenseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDTRADELICENSESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDashBoardProperty : ""
func (d *Daos) DisableTradeLicenseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDTRADELICENSESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDashBoardProperty : ""
func (d *Daos) DeleteTradeLicenseDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDTRADELICENSESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterTradeLicenseDashboard(ctx *models.Context, filter *models.TradeLicenseDashboardFilter, pagination *models.Pagination) ([]models.RefTradeLicenseDashboard, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tradeLicense []models.RefTradeLicenseDashboard
	if err = cursor.All(context.TODO(), &tradeLicense); err != nil {
		return nil, err
	}
	return tradeLicense, nil
}

// FilterShopRentQuery: ""
func (d *Daos) FilterTradeLicenseQuery(ctx *models.Context, filter *models.DashboardTradeLicenseDemandAndCollectionFilter) []bson.M {
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
	return query
}

// DashboardTradeLicenseDemandAndCollection : ""
func (d *Daos) DashboardTradeLicenseDemandAndCollection(ctx *models.Context, filter *models.DashboardTradeLicenseDemandAndCollectionFilter) (*models.DashboardTradeLicenseDemandAndCollection, error) {
	mainpipeline := []bson.M{}
	query := []bson.M{}
	query = d.FilterTradeLicenseQuery(ctx, filter)
	if len(query) > 0 {
		mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainpipeline = append(mainpipeline, bson.M{"$group": bson.M{"_id": nil,
		"totalDemandArrear":      bson.M{"$sum": "$demand.arrear.total"},
		"totalDemandCurrent":     bson.M{"$sum": "$demand.current.total"},
		"totalDemandTax":         bson.M{"$sum": "$demand.total.total"},
		"totalCollectionArrear":  bson.M{"$sum": "$collection.arrear.total"},
		"totalCollectionCurrent": bson.M{"$sum": "$collection.current.total"},
		"totalCollectionTax":     bson.M{"$sum": "$collection.total.total"},
	},
	})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property Query =>", mainpipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashboardTradeLicenseDemandAndCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashboardTradeLicenseDemandAndCollection{}, nil

}

// DashBoardStatusWiseMobileTowerCollectionAndChart : ""
func (d *Daos) DashBoardStatusWiseTradeLicenseCollectionAndChart(ctx *models.Context, filter *models.DashboardTradeLicenseDemandAndCollectionFilter) (*models.DashBoardStatusWiseTradeLicenseCollection, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"active": []bson.M{
			bson.M{"$match": bson.M{"status": "Active"}},
			bson.M{"$group": bson.M{"_id": nil, "active": bson.M{"$sum": 1}}}},
		"pending": []bson.M{
			bson.M{"$match": bson.M{"status": "Pending"}},
			bson.M{"$group": bson.M{"_id": nil, "pending": bson.M{"$sum": 1}}}},
		"expired": []bson.M{
			bson.M{"$match": bson.M{"status": "Expired"}},
			bson.M{"$group": bson.M{"_id": nil, "expired": bson.M{"$sum": 1}}}},
		"disabled": []bson.M{
			bson.M{"$match": bson.M{"status": "Disabled"}},
			bson.M{"$group": bson.M{"_id": nil, "disabled": bson.M{"$sum": 1}}}},
		"rejected": []bson.M{
			bson.M{"$match": bson.M{"status": "Rejected"}},
			bson.M{"$group": bson.M{"_id": nil, "rejected": bson.M{"$sum": 1}}}},
		"today": []bson.M{
			bson.M{"$match": bson.M{
				"status":     bson.M{"$in": []string{"Active", "Pending", "Disabled"}},
				"created.on": bson.M{"$gte": filter.TodayRange.From, "$lte": filter.TodayRange.To},
			}},
			bson.M{"$group": bson.M{"_id": nil, "today": bson.M{"$sum": 1}}}},
		"yesterday": []bson.M{
			bson.M{"$match": bson.M{
				"status":     bson.M{"$in": []string{"Active", "Pending", "Disabled"}},
				"created.on": bson.M{"$gte": filter.YesterdayRange.From, "$lte": filter.YesterdayRange.To},
			}},
			bson.M{"$group": bson.M{"_id": nil, "yesterday": bson.M{"$sum": 1}}}},
	},
	},
		bson.M{"$addFields": bson.M{"active": bson.M{"$arrayElemAt": []interface{}{"$active", 0}}}},
		bson.M{"$addFields": bson.M{"pending": bson.M{"$arrayElemAt": []interface{}{"$pending", 0}}}},
		bson.M{"$addFields": bson.M{"expired": bson.M{"$arrayElemAt": []interface{}{"$expired", 0}}}},
		bson.M{"$addFields": bson.M{"disabled": bson.M{"$arrayElemAt": []interface{}{"$disabled", 0}}}},
		bson.M{"$addFields": bson.M{"rejected": bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}}}},
		bson.M{"$addFields": bson.M{"today": bson.M{"$arrayElemAt": []interface{}{"$today", 0}}}},
		bson.M{"$addFields": bson.M{"yesterday": bson.M{"$arrayElemAt": []interface{}{"$yesterday", 0}}}},

		bson.M{"$addFields": bson.M{"active": "$active.active", "pending": "$pending.pending", "expired": "$expired.expired",
			"disabled": "$disabled.disabled", "rejected": "$rejected.rejected", "today": "$today.today", "yesterday": "$yesterday.yesterday"}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("tradelicense Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashBoardStatusWiseTradeLicenseCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashBoardStatusWiseTradeLicenseCollection{}, nil

}

// DashBoardStatusWiseShopRentCollectionAndChart : ""
func (d *Daos) UserwiseTradelicenseReport(ctx *models.Context, filter *models.UserFilter) ([]models.UserwiseTradeLicense, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.Type}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	var sd, ed time.Time
	t := time.Now()
	if filter != nil {
		if filter.DateRange == nil {
			sd = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
			ed = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		}
		if filter.DateRange.From != nil {
			sd = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
			// var ed time.Time
			if filter.DateRange.To != nil {
				ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
			} else {
				ed = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
			}
		}
	}
	fmt.Println("sd ===>", sd)
	fmt.Println("ed ===>", ed)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"as":   "tradelicensepayments",
		"from": "tradelicensepayments",
		"let":  bson.M{"userId": "$userName"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$eq": []string{"$details.collector.by", "$$userId"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			}}}},
			bson.M{"$facet": bson.M{
				"cash": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "Cash"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
				"cheque": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "Cheque"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
				"netbanking": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "NB"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
				"dd": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "DD"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
			}},
			bson.M{"$addFields": bson.M{"cash": bson.M{"$arrayElemAt": []interface{}{"$cash", 0}}}},
			bson.M{"$addFields": bson.M{"cheque": bson.M{"$arrayElemAt": []interface{}{"$cheque", 0}}}},
			bson.M{"$addFields": bson.M{"netbanking": bson.M{"$arrayElemAt": []interface{}{"$netbanking", 0}}}},
			bson.M{"$addFields": bson.M{"dd": bson.M{"$arrayElemAt": []interface{}{"$dd", 0}}}},
		}}},
	)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"tradelicensepayments": bson.M{"$arrayElemAt": []interface{}{"$tradelicensepayments", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("shoprent Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.UserwiseTradeLicense
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	return ddac, nil

}
