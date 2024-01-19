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

// SaveShopRentDashboard : ""
func (d *Daos) SaveShopRentDashboard(ctx *models.Context, shopRent *models.ShopRentDashboard) error {
	d.Shared.BsonToJSONPrint(shopRent)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).InsertOne(ctx.CTX, shopRent)
	return err
}

// GetSingleDashBoardProperty  : ""
func (d *Daos) GetSingleShopRentDashboard(ctx *models.Context, UniqueID string) (*models.RefShopRentDashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefShopRentDashboard
	var tower *models.RefShopRentDashboard
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateDashBoardProperty: ""
func (d *Daos) UpdateShopRentDashboard(ctx *models.Context, shopRent *models.ShopRentDashboard) error {
	selector := bson.M{"uniqueId": shopRent.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": shopRent}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableDashBoardProperty : ""
func (d *Daos) EnableShopRentDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDSHOPRENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDashBoardProperty : ""
func (d *Daos) DisableShopRentDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDSHOPRENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDashBoardProperty : ""
func (d *Daos) DeleteShopRentDashboard(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDSHOPRENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterShopRentDashboard(ctx *models.Context, filter *models.ShopRentDashboardFilter, pagination *models.Pagination) ([]models.RefShopRentDashboard, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDSHOPRENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopRent []models.RefShopRentDashboard
	if err = cursor.All(context.TODO(), &shopRent); err != nil {
		return nil, err
	}
	return shopRent, nil
}

// FilterShopRentQuery: ""
func (d *Daos) FilterShopRentQuery(ctx *models.Context, filter *models.DashboardShopRentDemandAndCollectionFilter) []bson.M {
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
	return query
}

// DashboardShopRentDemandAndCollection : ""
func (d *Daos) DashboardShopRentDemandAndCollection(ctx *models.Context, filter *models.DashboardShopRentDemandAndCollectionFilter) (*models.DashboardShopRentDemandAndCollection, error) {
	mainpipeline := []bson.M{}
	query := []bson.M{}
	query = d.FilterShopRentQuery(ctx, filter)
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashboardShopRentDemandAndCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashboardShopRentDemandAndCollection{}, nil

}

// DashBoardStatusWiseShopRentCollectionAndChart : ""
func (d *Daos) DashBoardStatusWiseShopRentCollectionAndChart(ctx *models.Context, filter *models.DashboardShopRentDemandAndCollectionFilter) (*models.DashBoardStatusWiseShopRentCollection, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"active": []bson.M{
			bson.M{"$match": bson.M{"status": "Active"}},
			bson.M{"$group": bson.M{"_id": nil, "active": bson.M{"$sum": 1}}}},
		"pending": []bson.M{
			bson.M{"$match": bson.M{"status": "Pending"}},
			bson.M{"$group": bson.M{"_id": nil, "pending": bson.M{"$sum": 1}}}},
		"disabled": []bson.M{
			bson.M{"$match": bson.M{"status": "Disabled"}},
			bson.M{"$group": bson.M{"_id": nil, "disabled": bson.M{"$sum": 1}}}},
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
		bson.M{"$addFields": bson.M{"disabled": bson.M{"$arrayElemAt": []interface{}{"$disabled", 0}}}},
		bson.M{"$addFields": bson.M{"today": bson.M{"$arrayElemAt": []interface{}{"$today", 0}}}},
		bson.M{"$addFields": bson.M{"yesterday": bson.M{"$arrayElemAt": []interface{}{"$yesterday", 0}}}},
		bson.M{"$addFields": bson.M{"active": "$active.active", "pending": "$pending.pending", "disabled": "$disabled.disabled",
			"rejected": "$rejected.rejected", "today": "$today.today", "yesterday": "$yesterday.yesterday"}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("shoprent Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashBoardStatusWiseShopRentCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashBoardStatusWiseShopRentCollection{}, nil

}
