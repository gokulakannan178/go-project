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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveDashBoardProperty : ""
func (d *Daos) SaveDashBoardProperty(ctx *models.Context, property *models.PropertyDashBoard) error {
	d.Shared.BsonToJSONPrint(property)
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).InsertOne(ctx.CTX, property)
	return err
}

// GetSingleDashBoardProperty  : ""
func (d *Daos) GetSingleDashBoardProperty(ctx *models.Context, UniqueID string) (*models.RefPropertyDashBoard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefPropertyDashBoard
	var tower *models.RefPropertyDashBoard
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateDashBoardProperty: ""
func (d *Daos) UpdateDashBoardProperty(ctx *models.Context, property *models.PropertyDashBoard) error {
	selector := bson.M{"uniqueId": property.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": property}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableDashBoardProperty : ""
func (d *Daos) EnableDashBoardProperty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDPROPERTYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDashBoardProperty : ""
func (d *Daos) DisableDashBoardProperty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDPROPERTYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDashBoardProperty : ""
func (d *Daos) DeleteDashBoardProperty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DASHBOARDPROPERTYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDashBoardProperty : ""
func (d *Daos) FilterDashBoardProperty(ctx *models.Context, filter *models.PropertyDashBoardFilter, pagination *models.Pagination) ([]models.RefPropertyDashBoard, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDASHBOARDPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var prop []models.RefPropertyDashBoard
	if err = cursor.All(context.TODO(), &prop); err != nil {
		return nil, err
	}
	return prop, nil
}

func (d *Daos) PropertyOverallDemandCron(ctx *models.Context, propertyfilter *models.PropertyFilter) error {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)

	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, propertyfilter)
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$ref.demand.total.totalTax"},
		"current": bson.M{"$sum": "$ref.demand.current.totalTax"},
		"arrear":  bson.M{"$sum": "$ref.demand.arrear.totalTax"},
		"survey":  bson.M{"$sum": 1},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.overall.demand.total":   "$total",
		"property.overall.demand.current": "$current",
		"property.overall.demand.arrear":  "$arrear",
		"property.overall.survey":         "$survey",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("demand =>", demand.Property.Overall)
	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.overall.demand.total":   demand.Property.Overall.Demand.Total,
			"property.overall.demand.current": demand.Property.Overall.Demand.Current,
			"property.overall.demand.arrear":  demand.Property.Overall.Demand.Arrear,
			"property.overall.survey":         demand.Property.Overall.Survey,

			"isDefault": true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (d *Daos) PropertyTodayDemandCron(ctx *models.Context, propertyfilter *models.PropertyFilter) error {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)

	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, propertyfilter)
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$ref.demand.total.totalTax"},
		"current": bson.M{"$sum": "$ref.demand.current.totalTax"},
		"arrear":  bson.M{"$sum": "$ref.demand.arrear.totalTax"},
		"survey":  bson.M{"$sum": 1},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.today.demand.total":   "$total",
		"property.today.demand.current": "$current",
		"property.today.demand.arrear":  "$arrear",
		"property.today.survey":         "$survey",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("demand =>", demand.Property.Today)

	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.today.demand.total":   demand.Property.Today.Demand.Total,
			"property.today.demand.current": demand.Property.Today.Demand.Current,
			"property.today.demand.arrear":  demand.Property.Today.Demand.Arrear,
			"property.today.survey":         demand.Property.Today.Survey,
			"isDefault":                     true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (d *Daos) PropertyOverallCollectionCron(ctx *models.Context, filter *models.PropertyPaymentFilter) error {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$details.amount"},
		"current": bson.M{"$sum": "$demand.current"},
		"arrear":  bson.M{"$sum": "$demand.arrear"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.overall.collection.total":   "$total",
		"property.overall.collection.current": "$current",
		"property.overall.collection.arrear":  "$arrear",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("demand =>", demand.Property.Overall)

	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.overall.collection.total":   demand.Property.Overall.Collection.Total,
			"property.overall.collection.current": demand.Property.Overall.Collection.Current,
			"property.overall.collection.arrear":  demand.Property.Overall.Collection.Arrear,
			"isDefault":                           true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (d *Daos) PropertyTodayCollectionCron(ctx *models.Context, filter *models.PropertyPaymentFilter) error {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$details.amount"},
		"current": bson.M{"$sum": "$demand.current"},
		"arrear":  bson.M{"$sum": "$demand.arrear"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.today.collection.total":   "$total",
		"property.today.collection.current": "$current",
		"property.today.collection.arrear":  "$arrear",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("collection =>", demand.Property.Today)

	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.today.collection.total":   demand.Property.Today.Collection.Total,
			"property.today.collection.current": demand.Property.Today.Collection.Current,
			"property.today.collection.arrear":  demand.Property.Today.Collection.Arrear,
			"isDefault":                         true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

// GetOverAllPropertyDashBoard  : ""
func (d *Daos) GetOverAllPropertyDashBoard(ctx *models.Context) (*models.OverallDashBoard, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.OverallDashBoard
	var tower *models.OverallDashBoard
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// PropertyMonthDemandCron : ""
func (d *Daos) PropertyMonthDemandCron(ctx *models.Context, propertyfilter *models.PropertyFilter) error {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)

	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, propertyfilter)
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$ref.demand.total.totalTax"},
		"current": bson.M{"$sum": "$ref.demand.current.totalTax"},
		"arrear":  bson.M{"$sum": "$ref.demand.arrear.totalTax"},
		"survey":  bson.M{"$sum": 1},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.month.demand.total":   "$total",
		"property.month.demand.current": "$current",
		"property.month.demand.arrear":  "$arrear",
		"property.month.survey":         "$survey",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("demand =>", demand.Property.Month)

	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.month.demand.total":   demand.Property.Month.Demand.Total,
			"property.month.demand.current": demand.Property.Month.Demand.Current,
			"property.month.demand.arrear":  demand.Property.Month.Demand.Arrear,
			"property.month.survey":         demand.Property.Month.Survey,
			"isDefault":                     true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

// PropertyMonthCollectionCron : ""
func (d *Daos) PropertyMonthCollectionCron(ctx *models.Context, filter *models.PropertyPaymentFilter) error {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$details.amount"},
		"current": bson.M{"$sum": "$demand.current"},
		"arrear":  bson.M{"$sum": "$demand.arrear"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.month.collection.total":   "$total",
		"property.month.collection.current": "$current",
		"property.month.collection.arrear":  "$arrear",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("collection =>", demand.Property.Month)

	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.month.collection.total":   demand.Property.Month.Collection.Total,
			"property.month.collection.current": demand.Property.Month.Collection.Current,
			"property.month.collection.arrear":  demand.Property.Month.Collection.Arrear,
			"isDefault":                         true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

// PropertyYearDemandCron : ""
func (d *Daos) PropertyYearDemandCron(ctx *models.Context, propertyfilter *models.PropertyFilter) error {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)

	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, propertyfilter)
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$ref.demand.total.totalTax"},
		"current": bson.M{"$sum": "$ref.demand.current.totalTax"},
		"arrear":  bson.M{"$sum": "$ref.demand.arrear.totalTax"},
		"survey":  bson.M{"$sum": 1},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.year.demand.total":   "$total",
		"property.year.demand.current": "$current",
		"property.year.demand.arrear":  "$arrear",
		"property.year.survey":         "$survey",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("demand =>", demand.Property.Year.Survey)

	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.year.demand.total":   demand.Property.Year.Demand.Total,
			"property.year.demand.current": demand.Property.Year.Demand.Current,
			"property.year.demand.arrear":  demand.Property.Year.Demand.Arrear,
			"property.year.survey":         demand.Property.Year.Survey,
			"isDefault":                    true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

// PropertyYearCollectionCron : ""
func (d *Daos) PropertyYearCollectionCron(ctx *models.Context, filter *models.PropertyPaymentFilter) error {
	mainPipeline := []bson.M{}
	// mainPipeline = append(mainPipeline, d.FilterPropertyPaymentQuery(ctx, filter)...)
	query := []bson.M{}

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"total":   bson.M{"$sum": "$details.amount"},
		"current": bson.M{"$sum": "$demand.current"},
		"arrear":  bson.M{"$sum": "$demand.arrear"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"_id": nil,
		"property.year.collection.total":   "$total",
		"property.year.collection.current": "$current",
		"property.year.collection.arrear":  "$arrear",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var demands []models.OverallDashBoard
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return err
	}
	if len(demands) < 1 {
		return errors.New("No demands")
	}
	demand := demands[0]
	opts := options.Update().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("collection =>", demand.Property.Year)

	res, err := ctx.DB.Collection(constants.COLLECTIONOVERALLDASHBOARD).UpdateOne(ctx.CTX, bson.M{"isDefault": true}, bson.M{
		"$set": bson.M{
			"property.year.collection.total":   demand.Property.Year.Collection.Total,
			"property.year.collection.current": demand.Property.Year.Collection.Current,
			"property.year.collection.arrear":  demand.Property.Year.Collection.Arrear,
			"isDefault":                        true,
		},
	}, opts)
	fmt.Println("successfully updated", res.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}
