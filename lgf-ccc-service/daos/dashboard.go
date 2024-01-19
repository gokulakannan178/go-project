package daos

import (
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) GetCollectionCount(ctx *models.Context, uniqueID string) (*models.Dashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$facet": bson.M{
			"total": []bson.M{
				bson.M{"$count": "total"}},

			"active": []bson.M{
				bson.M{"$match": bson.M{"status": "Active"}},
				bson.M{"$count": "active"}},
			"disabled": []bson.M{
				bson.M{"$match": bson.M{"status": "Disabled"}},
				bson.M{"$count": "disabled"}}},
	},

		bson.M{"$addFields": bson.M{"total": bson.M{"$arrayElemAt": []interface{}{"$total", 0}}}},
		bson.M{"$addFields": bson.M{"active": bson.M{"$arrayElemAt": []interface{}{"$active", 0}}}},
		bson.M{"$addFields": bson.M{"disabled": bson.M{"$arrayElemAt": []interface{}{"$disabled", 0}}}},
		bson.M{"$addFields": bson.M{"active": "$active.active", "disabled": "$disabled.disabled", "total": "$total.total"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(uniqueID).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Dashboards []models.Dashboard
	var Dashboard *models.Dashboard
	if err = cursor.All(ctx.CTX, &Dashboards); err != nil {
		return nil, err
	}
	if len(Dashboards) > 0 {
		Dashboard = &Dashboards[0]
	}
	return Dashboard, err
}

func (d *Daos) GetHousevisitedCount(ctx *models.Context, uniqueID string) (*models.HousevisitedCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$facet": bson.M{
			"total": []bson.M{
				bson.M{"$match": bson.M{
					"isStatus": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSCOLLECTED, constants.HOUSEVISITEDSTATUSNOTAVAILABLE}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$eq": []interface{}{"$status", "Active"}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "total": bson.M{"$sum": 1}}}},
			"collected": []bson.M{
				bson.M{"$match": bson.M{
					"isStatus": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSCOLLECTED}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$eq": []interface{}{"$status", "Active"}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "collected": bson.M{"$sum": 1}}}},
			"notCollected": []bson.M{
				bson.M{"$match": bson.M{
					"isStatus": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSNOTAVAILABLE}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$eq": []interface{}{"$status", "Active"}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "notCollected": bson.M{"$sum": 1}}}},
		}},

		bson.M{"$addFields": bson.M{"total": bson.M{"$arrayElemAt": []interface{}{"$total", 0}}}},
		bson.M{"$addFields": bson.M{"collected": bson.M{"$arrayElemAt": []interface{}{"$collected", 0}}}},
		bson.M{"$addFields": bson.M{"notCollected": bson.M{"$arrayElemAt": []interface{}{"$notCollected", 0}}}},
		bson.M{"$addFields": bson.M{"collected": "$collected.collected", "notCollected": "$notCollected.notCollected", "total": "$total.total"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(uniqueID).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Dashboards []models.HousevisitedCount
	var Dashboard *models.HousevisitedCount
	if err = cursor.All(ctx.CTX, &Dashboards); err != nil {
		return nil, err
	}
	if len(Dashboards) > 0 {
		Dashboard = &Dashboards[0]
	}
	return Dashboard, err
}
func (d *Daos) GetvehicleCount(ctx *models.Context, uniqueID string) (*models.Dashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$facet": bson.M{
			"total": []bson.M{
				bson.M{"$match": bson.M{
					"status": bson.M{"$in": []string{constants.VECHILESTATUSACTIVE, constants.VECHILESTATUSDISABLED}}},
				},

				bson.M{"$group": bson.M{"_id": nil, "total": bson.M{"$sum": 1}}}},

			"Active": []bson.M{
				bson.M{"$match": bson.M{"status": constants.VECHILESTATUSACTIVE}},
				bson.M{"$count": "Active"}},
			"InActive": []bson.M{
				bson.M{"$match": bson.M{"status": constants.VECHILESTATUSDISABLED}},
				bson.M{"$count": "InActive"}}},
	},

		bson.M{"$addFields": bson.M{"total": bson.M{"$arrayElemAt": []interface{}{"$total", 0}}}},
		bson.M{"$addFields": bson.M{"Active": bson.M{"$arrayElemAt": []interface{}{"$Active", 0}}}},
		bson.M{"$addFields": bson.M{"InActive": bson.M{"$arrayElemAt": []interface{}{"$InActive", 0}}}},
		bson.M{"$addFields": bson.M{"Active": "$Active.Active", "disabled": "$InActive.InActive", "total": "$total.total"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(uniqueID).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Dashboards []models.Dashboard
	var Dashboard *models.Dashboard
	if err = cursor.All(ctx.CTX, &Dashboards); err != nil {
		return nil, err
	}
	if len(Dashboards) > 0 {
		Dashboard = &Dashboards[0]
	}
	return Dashboard, err
}
func (d *Daos) GetDumbSiteCount(ctx *models.Context, uniqueID string) (*models.DumbSiteCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$facet": bson.M{
			"total": []bson.M{
				bson.M{"$count": "total"}},

			"totalVehicle": []bson.M{
				bson.M{"$match": bson.M{"status": constants.DUMPHISTORYSTATUSACTIVE}},
				bson.M{"$count": "totalVehicle"}},
			"quantity": []bson.M{
				bson.M{"$match": bson.M{
					"status": bson.M{"$in": []string{constants.DUMPHISTORYSTATUSACTIVE}}},
				},

				bson.M{"$group": bson.M{"_id": nil, "quantity": bson.M{"$sum": "$quantity"}}}},
		}},

		bson.M{"$addFields": bson.M{"total": bson.M{"$arrayElemAt": []interface{}{"$total", 0}}}},
		bson.M{"$addFields": bson.M{"totalVehicle": bson.M{"$arrayElemAt": []interface{}{"$totalVehicle", 0}}}},
		bson.M{"$addFields": bson.M{"quantity": bson.M{"$arrayElemAt": []interface{}{"$quantity", 0}}}},
		bson.M{"$addFields": bson.M{"totalVehicle": "$totalVehicle.totalVehicle", "quantity": "$quantity.quantity", "total": "$total.total"}})

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(uniqueID).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Dashboards []models.DumbSiteCount
	var Dashboard *models.DumbSiteCount
	if err = cursor.All(ctx.CTX, &Dashboards); err != nil {
		return nil, err
	}
	if len(Dashboards) > 0 {
		Dashboard = &Dashboards[0]
	}
	return Dashboard, err
}

func (d *Daos) GetUsertypeCount(ctx *models.Context, uniqueID string) (*models.UserTypeCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{

		"admin": []bson.M{
			bson.M{"$match": bson.M{
				"status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}},
			},
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$type", "Admin"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "admin": bson.M{"$sum": 1}}}},
		"projectManager": []bson.M{
			bson.M{"$match": bson.M{
				"status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}},
			},
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$type", "ProjectManager"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "projectManager": bson.M{"$sum": 1}}}},

		"garbageCollector": []bson.M{
			bson.M{"$match": bson.M{
				"status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}},
			},
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$type", "GarbageCollector"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "garbageCollector": bson.M{"$sum": 1}}}},

		"dumbSiteUser": []bson.M{
			bson.M{"$match": bson.M{
				"status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}},
			},
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$type", "DumbSiteUser"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "dumbSiteUser": bson.M{"$sum": 1}}}},
		"citizen": []bson.M{
			bson.M{"$match": bson.M{
				"status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}},
			},
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$type", "Citizen"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "citizen": bson.M{"$sum": 1}}}},

		"driver": []bson.M{
			bson.M{"$match": bson.M{
				"status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}},
			},
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$type", "Driver"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "driver": bson.M{"$sum": 1}}}},
	}},
		bson.M{"$addFields": bson.M{
			"admin":            bson.M{"$arrayElemAt": []interface{}{"$admin.admin", 0}},
			"projectManager":   bson.M{"$arrayElemAt": []interface{}{"$projectManager.projectManager", 0}},
			"garbageCollector": bson.M{"$arrayElemAt": []interface{}{"$garbageCollector.garbageCollector", 0}},
			"dumbSiteUser":     bson.M{"$arrayElemAt": []interface{}{"$dumbSiteUser.dumbSiteUser", 0}},
			"citizen":          bson.M{"$arrayElemAt": []interface{}{"$citizen.citizen", 0}},
			"driver":           bson.M{"$arrayElemAt": []interface{}{"$driver.driver", 0}},
		}},
	)

	d.Shared.BsonToJSONPrintTag("Project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(uniqueID).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserTypeCount []models.UserTypeCount
	var UserTypeCounts *models.UserTypeCount
	if err = cursor.All(ctx.CTX, &UserTypeCount); err != nil {
		return nil, err
	}
	if len(UserTypeCount) > 0 {
		UserTypeCounts = &UserTypeCount[0]
	}

	return UserTypeCounts, nil
}

func (d *Daos) GetPropertyCount(ctx *models.Context, filter *models.FilterProperties) (*models.PropertyCount, error) {
	mainPipeline := []bson.M{}
	//var sd, ed *time.Time
	var sd, ed time.Time

	//	query := []bson.M{}
	if filter != nil {

		if filter.DateRange.From != nil {

			if filter.DateRange.From != nil {
				sd = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				ed = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				}
				//query = append(query, bson.M{"entryTime": bson.M{"$gte": sd, "$lte": ed}})

			}
		}

		// if len(query) > 0 {
		// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
		// }
	}
	startmonth := d.Shared.BeginningOfMonth(*filter.DateRange.From)
	sm := &startmonth
	endmonth := d.Shared.EndOfMonth(*filter.DateRange.From)
	em := &endmonth
	fmt.Println("sm======>", sm)
	fmt.Println("em======>", em)
	fmt.Println("ed======>", ed)
	startWeek := d.Shared.StartDayOfWeek(*filter.DateRange.From)
	sw := &startWeek
	endweek := d.Shared.EndDayOfWeek(*filter.DateRange.From)
	ew := &endweek
	fmt.Println("sw======>", sw)
	fmt.Println("ew======>", ew)
	fmt.Println("ilter.DateRange.From.Year(),======>", filter.DateRange.From.Year())
	endyear := time.Date(filter.DateRange.From.Year(), time.December, 31, 23, 59, 59, 0, filter.DateRange.From.Location())
	startyear := time.Date(filter.DateRange.From.Year(), time.January, 1, 0, 0, 0, 0, filter.DateRange.From.Location())

	//startyear := time.Date(filter.DateRange.From.Year(), 0, 0, 0, 0, 0, 0, filter.DateRange.From.Location())
	sy := &startyear
	// eyear := time.Date(filter.DateRange.To.Year(), 0, 0, 23, 59, 59, 0, filter.DateRange.To.Location())
	ey := &endyear
	fmt.Println("sy======>", startyear)
	fmt.Println("eyear======>", endyear)
	mainPipeline = append(mainPipeline,

		bson.M{"$facet": bson.M{
			"totalProperty": []bson.M{
				bson.M{"$match": bson.M{
					"status": bson.M{"$in": []string{constants.PROPERTIESSTATUSACTIVE, constants.PROPERTIESSTATUSDISABLED}}},
				},

				bson.M{"$group": bson.M{"_id": nil, "totalProperty": bson.M{"$sum": 1}}}},
			"today": []bson.M{
				bson.M{"$match": bson.M{
					"status": bson.M{"$in": []string{constants.PROPERTIESSTATUSACTIVE, constants.PROPERTIESSTATUSDISABLED}}},
				},
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$gte": []interface{}{"$registerDate", sd}},
					bson.M{"$lte": []interface{}{"$registerDate", ed}},
				}}}},
				bson.M{"$group": bson.M{"_id": nil, "today": bson.M{"$sum": 1}}}},

			"thisWeek": []bson.M{
				bson.M{"$match": bson.M{
					"status": bson.M{"$in": []string{constants.PROPERTIESSTATUSACTIVE, constants.PROPERTIESSTATUSDISABLED}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$gte": []interface{}{"$registerDate", sw}},
							bson.M{"$lte": []interface{}{"$registerDate", ew}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "thisWeek": bson.M{"$sum": 1}}}},
			"thisMonth": []bson.M{
				bson.M{"$match": bson.M{
					"status": bson.M{"$in": []string{constants.PROPERTIESSTATUSACTIVE, constants.PROPERTIESSTATUSDISABLED}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$gte": []interface{}{"$registerDate", sm}},
							bson.M{"$lte": []interface{}{"$registerDate", em}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "thisMonth": bson.M{"$sum": 1}}}},
			"thisYear": []bson.M{
				bson.M{"$match": bson.M{
					"status": bson.M{"$in": []string{constants.PROPERTIESSTATUSACTIVE, constants.PROPERTIESSTATUSDISABLED}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$gte": []interface{}{"$registerDate", sy}},
							bson.M{"$lte": []interface{}{"$registerDate", ey}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "thisYear": bson.M{"$sum": 1}}}},
		}},
		// 	bson.M{"$addFields": bson.M{
		// 		"userId":       "$$userId",
		// 		"date":         sd,
		// 		"totalAmount":  bson.M{"$arrayElemAt": []interface{}{"$totalAmount.totalAmount", 0}},
		// 		"closed":       bson.M{"$arrayElemAt": []interface{}{"$closed.closed", 0}},
		// 		"opened":       bson.M{"$arrayElemAt": []interface{}{"$opened.Opened", 0}},
		// 		"totalVehicle": bson.M{"$arrayElemAt": []interface{}{"$totalVehicle.totalVehicle", 0}},
		// 		"cash":         bson.M{"$arrayElemAt": []interface{}{"$cash.cash", 0}},
		// 		"upiId":        bson.M{"$arrayElemAt": []interface{}{"$upiId.upiId", 0}},
		// 		"card":         bson.M{"$arrayElemAt": []interface{}{"$card.card", 0}},
		// 		"paytm":        bson.M{"$arrayElemAt": []interface{}{"$paytm.paytm", 0}},
		// 	}},
		// )
		// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})
		// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		// 	"userId":       "$report.userId",
		// 	"closed":       "$report.closed",
		// 	"opened":       "$report.opened",
		// 	"totalAmount":  "$report.totalAmount",
		// 	"totalVehicle": "$report.totalVehicle",
		// 	"cash":         "$report.cash",
		// 	"upiId":        "$report.upiId",
		// 	"card":         "$report.card",
		// 	"paytm":        "$report.paytm",
		// 	"date":         "$report.date",
		//}}
		bson.M{"$addFields": bson.M{
			"totalProperty": bson.M{"$arrayElemAt": []interface{}{"$totalProperty.totalProperty", 0}},
			"today":         bson.M{"$arrayElemAt": []interface{}{"$today.today", 0}},
			"thisWeek":      bson.M{"$arrayElemAt": []interface{}{"$thisWeek.thisWeek", 0}},
			"thisMonth":     bson.M{"$arrayElemAt": []interface{}{"$thisMonth.thisMonth", 0}},
			"thisYear":      bson.M{"$arrayElemAt": []interface{}{"$thisYear.thisYear", 0}},
		}},
	)

	d.Shared.BsonToJSONPrintTag("Project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTIES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserTypeCount []models.PropertyCount
	var UserTypeCounts *models.PropertyCount
	if err = cursor.All(ctx.CTX, &UserTypeCount); err != nil {
		return nil, err
	}
	if len(UserTypeCount) > 0 {
		UserTypeCounts = &UserTypeCount[0]
	}

	return UserTypeCounts, nil
}

func (d *Daos) GetGarbaggeCount(ctx *models.Context, filter *models.FilterHouseVisited) (*models.PropertyCount, error) {
	mainPipeline := []bson.M{}
	//var sd, ed *time.Time
	var sd, ed time.Time

	//	query := []bson.M{}
	if filter != nil {

		if filter.DateRange.From != nil {

			if filter.DateRange.From != nil {
				sd = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				ed = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				}
				//query = append(query, bson.M{"entryTime": bson.M{"$gte": sd, "$lte": ed}})

			}
		}

		// if len(query) > 0 {
		// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
		// }
	}
	startmonth := d.Shared.BeginningOfMonth(*filter.DateRange.From)
	sm := &startmonth
	endmonth := d.Shared.EndOfMonth(*filter.DateRange.From)
	em := &endmonth
	startWeek := d.Shared.StartDayOfWeek(*filter.DateRange.From)
	sw := &startWeek
	endweek := d.Shared.EndDayOfWeek(*filter.DateRange.From)
	ew := &endweek
	fmt.Println("sm======>", sm)
	fmt.Println("em======>", em)
	fmt.Println("ed======>", ed)

	fmt.Println("ilter.DateRange.From.Year(),======>", filter.DateRange.From.Year())

	endyear := time.Date(filter.DateRange.From.Year(), time.December, 31, 23, 59, 59, 0, filter.DateRange.From.Location())
	startyear := time.Date(filter.DateRange.From.Year(), time.January, 1, 0, 0, 0, 0, filter.DateRange.From.Location())

	//startyear := time.Date(filter.DateRange.From.Year(), 0, 0, 0, 0, 0, 0, filter.DateRange.From.Location())
	sy := &startyear
	// eyear := time.Date(filter.DateRange.To.Year(), 0, 0, 23, 59, 59, 0, filter.DateRange.To.Location())
	ey := &endyear
	fmt.Println("sy======>", startyear)
	fmt.Println("eyear======>", endyear)
	mainPipeline = append(mainPipeline,

		bson.M{"$facet": bson.M{
			// "totalProperty": []bson.M{
			// 	bson.M{"$match": bson.M{
			// 		"status": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSCOLLECTED, constants.PROPERTIESSTATUSDISABLED}}},
			// 	},

			// 	bson.M{"$group": bson.M{"_id": nil, "totalProperty": bson.M{"$sum": 1}}}},
			"today": []bson.M{
				bson.M{"$match": bson.M{
					"isStatus": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSCOLLECTED, constants.HOUSEVISITEDSTATUSNOTAVAILABLE}}},
				},
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$gte": []interface{}{"$date", sd}},
					bson.M{"$lte": []interface{}{"$date", ed}},
				}}}},
				bson.M{"$group": bson.M{"_id": nil, "today": bson.M{"$sum": 1}}}},

			"thisWeek": []bson.M{
				bson.M{"$match": bson.M{
					"isStatus": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSCOLLECTED, constants.HOUSEVISITEDSTATUSNOTAVAILABLE}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$gte": []interface{}{"$date", sw}},
							bson.M{"$lte": []interface{}{"$date", ew}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "thisWeek": bson.M{"$sum": 1}}}},
			"thisMonth": []bson.M{
				bson.M{"$match": bson.M{
					"isStatus": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSCOLLECTED, constants.HOUSEVISITEDSTATUSNOTAVAILABLE}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$gte": []interface{}{"$date", sm}},
							bson.M{"$lte": []interface{}{"$date", em}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "thisMonth": bson.M{"$sum": 1}}}},
			"thisYear": []bson.M{
				bson.M{"$match": bson.M{
					"isStatus": bson.M{"$in": []string{constants.HOUSEVISITEDSTATUSCOLLECTED, constants.HOUSEVISITEDSTATUSNOTAVAILABLE}}},
				},
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$gte": []interface{}{"$date", sy}},
							bson.M{"$lte": []interface{}{"$date", ey}},
						}},
				}},
				bson.M{"$group": bson.M{"_id": nil, "thisYear": bson.M{"$sum": 1}}}},
		}},

		bson.M{"$addFields": bson.M{
			//"totalProperty": bson.M{"$arrayElemAt": []interface{}{"$totalProperty.totalProperty", 0}},
			"today":     bson.M{"$arrayElemAt": []interface{}{"$today.today", 0}},
			"thisWeek":  bson.M{"$arrayElemAt": []interface{}{"$thisWeek.thisWeek", 0}},
			"thisMonth": bson.M{"$arrayElemAt": []interface{}{"$thisMonth.thisMonth", 0}},
			"thisYear":  bson.M{"$arrayElemAt": []interface{}{"$thisYear.thisYear", 0}},
		}},
	)

	d.Shared.BsonToJSONPrintTag("Project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var UserTypeCount []models.PropertyCount
	var UserTypeCounts *models.PropertyCount
	if err = cursor.All(ctx.CTX, &UserTypeCount); err != nil {
		return nil, err
	}
	if len(UserTypeCount) > 0 {
		UserTypeCounts = &UserTypeCount[0]
	}

	return UserTypeCounts, nil
}
