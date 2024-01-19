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
)

// SaveMobileTower : ""
func (d *Daos) SaveMobileTower(ctx *models.Context, mobileTower *models.PropertyMobileTower) error {
	d.Shared.BsonToJSONPrint(mobileTower)
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).InsertOne(ctx.CTX, mobileTower)
	return err
}

// GetSingleMobileTower : ""
func (d *Daos) GetSingleMobileTower(ctx *models.Context, UniqueID string) (*models.RefPropertyMobileTower, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefPropertyMobileTower
	var tower *models.RefPropertyMobileTower
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateMobileTower : ""
func (d *Daos) UpdateMobileTower(ctx *models.Context, mobileTower *models.PropertyMobileTower) error {
	selector := bson.M{"uniqueId": mobileTower.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": mobileTower}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMobileTower : ""
func (d *Daos) EnableMobileTower(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMOBILETOWERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableMobileTower : ""
func (d *Daos) DisableMobileTower(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMOBILETOWERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteMobileTower : ""
func (d *Daos) DeleteMobileTower(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYMOBILETOWERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMobileTower : ""
func (d *Daos) FilterMobileTower(ctx *models.Context, filter *models.PropertyMobileTowerFilter, pagination *models.Pagination) ([]models.RefPropertyMobileTower, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refMobileTower []models.RefPropertyMobileTower
	if err = cursor.All(context.TODO(), &refMobileTower); err != nil {
		return nil, err
	}
	return refMobileTower, nil
}

// FilterShopRentQuery: ""
func (d *Daos) FilterMobileTowerQuery(ctx *models.Context, filter *models.DashboardMobileTowerDemandAndCollectionFilter) []bson.M {
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	return query
}

// DashboardMobileTowerDemandAndCollection : ""
func (d *Daos) DashboardMobileTowerDemandAndCollection(ctx *models.Context, filter *models.DashboardMobileTowerDemandAndCollectionFilter) (*models.DashboardMobileTowerDemandAndCollection, error) {
	mainpipeline := []bson.M{}
	query := []bson.M{}
	query = d.FilterMobileTowerQuery(ctx, filter)
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashboardMobileTowerDemandAndCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashboardMobileTowerDemandAndCollection{}, nil

}

// DashBoardStatusWiseMobileTowerCollectionAndChart : ""
func (d *Daos) DashBoardStatusWiseMobileTowerCollectionAndChart(ctx *models.Context, filter *models.DashboardMobileTowerDemandAndCollectionFilter) (*models.DashBoardStatusWiseMobileTowerCollection, error) {

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
		bson.M{"$addFields": bson.M{"rejected": bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}}}},
		bson.M{"$addFields": bson.M{"disabled": bson.M{"$arrayElemAt": []interface{}{"$disabled", 0}}}},
		bson.M{"$addFields": bson.M{"today": bson.M{"$arrayElemAt": []interface{}{"$today", 0}}}},
		bson.M{"$addFields": bson.M{"yesterday": bson.M{"$arrayElemAt": []interface{}{"$yesterday", 0}}}},
		bson.M{"$addFields": bson.M{"active": "$active.active", "pending": "$pending.pending", "rejected": "$rejected.rejected",
			"disabled": "$disabled.disabled", "today": "$today.today", "yesterday": "$yesterday.yesterday"}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("mobiletower Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.DashBoardStatusWiseMobileTowerCollection
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	if len(ddac) > 0 {
		return &ddac[0], nil
	}
	return &models.DashBoardStatusWiseMobileTowerCollection{}, nil

}
