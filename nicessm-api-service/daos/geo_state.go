package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveState :""
func (d *Daos) SaveState(ctx *models.Context, state *models.State) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSTATE).InsertOne(ctx.CTX, state)
	if err != nil {
		return err
	}
	state.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleState : ""
func (d *Daos) GetSingleState(ctx *models.Context, code string) (*models.RefState, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.RefState
	var state *models.RefState
	if err = cursor.All(ctx.CTX, &states); err != nil {
		return nil, err
	}
	if len(states) > 0 {
		state = &states[0]
	}
	return state, nil
}

//UpdateState : ""
func (d *Daos) UpdateState(ctx *models.Context, state *models.State) error {
	selector := bson.M{"_id": state.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": state, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterState : ""
func (d *Daos) FilterState(ctx *models.Context, statefilter *models.StateFilter, pagination *models.Pagination) ([]models.RefState, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTSTATE, bson.M{
		"stateId": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.PROJECTSTATESTATUSACTIVE}},
			{"$eq": []string{"$state", "$$stateId"}},
		}}}},
	}, "ref.projects", "ref.projects")...)
	query := []bson.M{}
	if statefilter != nil {

		if len(statefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": statefilter.ActiveStatus}})
		}
		if len(statefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": statefilter.Status}})
		}
		//Regex
		if statefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: statefilter.Regex.Name, Options: "xi"}})
		}
		if statefilter.OmitProjectState.Is {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$ne": statefilter.OmitProjectState.Project}})
		}
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})

	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// 	"from": constants.COLLECTIONPROJECTSTATE,
	// 	"as":   "omitIds",
	// 	"let":  bson.M{"projectId": statefilter.OmitProjectState},
	// 	"pipeline": []bson.M{
	// 		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 			{"$eq": []interface{}{"$project", "$$projectId"}},
	// 		}}}},
	// 		{"$group": bson.M{"_id": nil, "stateIds": bson.M{"$push": "$state"}}},
	// 	},
	// }})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"omitIds": bson.M{"$arrayElemAt": []interface{}{"$omitIds", 0}}}})
	// // mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"omitIds.stateIds": bson.M{"$cond": bson.M{"if": bson.M{"$isArray": "$omitIds.stateIds"}, "then": "$omitIds.stateIds", "else": []interface{}{}}}}})

	// mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": bson.M{"$ne": bson.M{"$cond": bson.M{"if": bson.M{"$isArray": "$omitIds.stateIds"}, "then": "$omitIds.stateIds", "else": []interface{}{}}}}}})

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		d.Shared.BsonToJSONPrintTag("state pagenation query =>", paginationPipeline)
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.RefState
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}

//EnableState :""
func (d *Daos) EnableState(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableState :""
func (d *Daos) DisableState(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATESTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteState :""
func (d *Daos) DeleteState(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleStateWithName : ""
func (d *Daos) GetSingleStateWithName(ctx *models.Context, Name string) ([]models.RefState, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.RefState
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}

//GetSingleStateWithName : ""
func (d *Daos) GetSingleStateWithNameV2(ctx *models.Context, Name string, isRegex bool) (*models.RefState, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Name != "" {
		if isRegex {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
		} else {
			query = append(query, bson.M{"name": Name})

		}
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.RefState
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	if len(states) > 0 {
		return &states[0], nil
	}
	return nil, errors.New("state not available")
}

//GetSingleStateWithName : ""
func (d *Daos) GetSingleStateWithUniqueID(ctx *models.Context, UniqueID string) (*models.RefState, error) {
	mainPipeline := []bson.M{}

	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.RefState
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	if len(states) > 0 {
		return &states[0], nil
	}

	return nil, errors.New("state not found")
}
func (d *Daos) GeoDetatilsReport(ctx *models.Context, statefilter *models.StateFilter) ([]models.GeoDetailsReport2, error) {
	mainPipeline := []bson.M{}
	villagepipeline := []bson.M{}
	villagepipeline = append(villagepipeline, d.CommonLookupAdvanced(constants.COLLECTIONVILLAGE, bson.M{
		"gramPanchayatCode": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.STATESTATUSACTIVE}},
			{"$eq": []string{"$gramPanchayat", "$$gramPanchayatCode"}},
		}}}},
		bson.M{"$project": bson.M{"id": 1, "name": 1, "gramPanchayat": 1}},
		{"$group": bson.M{"_id": "$_id", "village": bson.M{"$push": "$name"}, "villageCode": bson.M{"$push": "$_id"}}},
		{
			"$addFields": bson.M{"village": bson.M{"$arrayElemAt": []interface{}{"$village", 0}}},
		},
		{
			"$addFields": bson.M{"villageCode": bson.M{"$arrayElemAt": []interface{}{"$villageCode", 0}}},
		},
	}, "villages", "villages")...)
	GramPanchayatPipeline := []bson.M{}
	GramPanchayatPipeline = append(GramPanchayatPipeline, d.CommonLookupAdvanced(constants.COLLECTIONGRAMPANCHAYAT, bson.M{
		"blockcode": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.STATESTATUSACTIVE}},
			{"$eq": []string{"$block", "$$blockcode"}},
		}}}},
		bson.M{"$project": bson.M{"id": 1, "name": 1, "block": 1}},
		{"$group": bson.M{"_id": "$_id", "gramPanchayat": bson.M{"$push": "$name"}, "gramPanchayatCode": bson.M{"$push": "$_id"}}},
		{
			"$addFields": bson.M{"grampanchayat": bson.M{"$arrayElemAt": []interface{}{"$grampanchayat", 0}}},
		},
		{
			"$addFields": bson.M{"grampanchayatCode": bson.M{"$arrayElemAt": []interface{}{"$grampanchayatCode", 0}}},
		},
		villagepipeline[0],
	}, "gramPanchayats", "gramPanchayats")...)
	blockPipeline := []bson.M{}
	blockPipeline = append(blockPipeline, d.CommonLookupAdvanced(constants.COLLECTIONBLOCK, bson.M{
		"districcode": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.STATESTATUSACTIVE}},
			{"$eq": []string{"$district", "$$districcode"}},
		}}}},
		bson.M{"$project": bson.M{"id": 1, "name": 1, "district": 1}},
		{"$group": bson.M{"_id": "$_id", "block": bson.M{"$push": "$name"}, "blockCode": bson.M{"$push": "$_id"}}},
		{
			"$addFields": bson.M{"block": bson.M{"$arrayElemAt": []interface{}{"$block", 0}}},
		},
		{
			"$addFields": bson.M{"blockCode": bson.M{"$arrayElemAt": []interface{}{"$blockCode", 0}}},
		},
		GramPanchayatPipeline[0],
	}, "blocks", "blocks")...)

	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{"status": "Active"}},

		bson.M{"$project": bson.M{"id": 1, "name": 1}},
		bson.M{"$group": bson.M{"_id": "$_id", "state": bson.M{"$push": "$name"}, "stateCode": bson.M{"$push": "$_id"}}},
	)
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"state": bson.M{"$arrayElemAt": []interface{}{"$state", 0}}},
		//	"$addFields": bson.M{"stateCode": bson.M{"$arrayElemAt": []interface{}{"$stateCode", 0}}},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"stateCode": bson.M{"$arrayElemAt": []interface{}{"$stateCode", 0}}},
		//	"$addFields": bson.M{"stateCode": bson.M{"$arrayElemAt": []interface{}{"$stateCode", 0}}},
	})
	mainPipeline = append(mainPipeline, d.CommonLookupAdvanced(constants.COLLECTIONDISTRICT, bson.M{
		"statecode": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.STATESTATUSACTIVE}},
			{"$eq": []string{"$state", "$$statecode"}},
		}}}},
		bson.M{"$project": bson.M{"id": 1, "name": 1, "state": 1}},
		{"$group": bson.M{"_id": "$_id", "district": bson.M{"$push": "$name"}, "districtCode": bson.M{"$push": "$id"}}},
		{
			"$addFields": bson.M{"distric": bson.M{"$arrayElemAt": []interface{}{"$distric", 0}}},
		},
		{
			"$addFields": bson.M{"districCode": bson.M{"$arrayElemAt": []interface{}{"$districCode", 0}}},
		},

		blockPipeline[0],
	}, "districts", "districts")...)

	//Aggregation

	//var emptyData *models.GeoDetailsReport2
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.GeoDetailsReport2
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}
func (d *Daos) GeoDetatilsReportV2(ctx *models.Context, statefilter *models.StateFilter) ([]models.GeoDetailsReport2, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active"}},

		bson.M{"$project": bson.M{"id": 1, "name": 1, "uniqueId": 1}},
		//bson.M{"$group": bson.M{"_id": "$_id", "state": bson.M{"$push": "$name"}, "stateCode": bson.M{"$push": "$uniqueId"}}}
	)
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"state": 1}})

	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$addFields": bson.M{"state": bson.M{"$arrayElemAt": []interface{}{"$state", 0}}},
	// 	//	"$addFields": bson.M{"stateCode": bson.M{"$arrayElemAt": []interface{}{"$stateCode", 0}}},
	// })
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$addFields": bson.M{"stateCode": bson.M{"$arrayElemAt": []interface{}{"$stateCode", 0}}},
	// 	//	"$addFields": bson.M{"stateCode": bson.M{"$arrayElemAt": []interface{}{"$stateCode", 0}}},
	// })

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONDISTRICT,
		"as":   "districts",
		"let":  bson.M{"stateCode": "$_id"},
		"pipeline": []bson.M{

			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$state", "$$stateCode"}},
			}}}},
			bson.M{"$project": bson.M{"id": 1, "name": 1, "state": 1, "uniqueId": 1}},
			bson.M{"$sort": bson.M{"name": 1}},
			// {"$group": bson.M{"_id": "$_id", "distric": bson.M{"$push": "$name"}, "districCode": bson.M{"$push": "$uniqueId"}}},
			// {
			// 	"$addFields": bson.M{"distric": bson.M{"$arrayElemAt": []interface{}{"$distric", 0}}},
			// },
			// {
			// 	"$addFields": bson.M{"districCode": bson.M{"$arrayElemAt": []interface{}{"$districCode", 0}}},
			// },
			{
				"$lookup": bson.M{
					"from": constants.COLLECTIONBLOCK,
					"as":   "blocks",
					"let":  bson.M{"districtcode": "$_id"},
					"pipeline": []bson.M{

						{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							{"$eq": []string{"$status", "Active"}},
							{"$eq": []string{"$district", "$$districtcode"}},
						}}}},
						bson.M{"$project": bson.M{"id": 1, "name": 1, "district": 1, "uniqueId": 1}},
						bson.M{"$sort": bson.M{"name": 1}},
						// {"$group": bson.M{"_id": "$_id", "block": bson.M{"$push": "$name"}, "blockCode": bson.M{"$push": "$uniqueId"}}},
						// {
						// 	"$addFields": bson.M{"block": bson.M{"$arrayElemAt": []interface{}{"$block", 0}}},
						// },
						// {
						// 	"$addFields": bson.M{"blockCode": bson.M{"$arrayElemAt": []interface{}{"$blockCode", 0}}},
						// },
						{
							"$lookup": bson.M{
								"from": constants.COLLECTIONGRAMPANCHAYAT,
								"as":   "gramPanchayats",
								"let":  bson.M{"blockcode": "$_id"},
								"pipeline": []bson.M{
									{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
										{"$eq": []string{"$status", "Active"}},
										{"$eq": []string{"$block", "$$blockcode"}},
									}}}},
									bson.M{"$project": bson.M{"id": 1, "name": 1, "block": 1, "uniqueId": 1}},
									bson.M{"$sort": bson.M{"name": 1}},
									// {"$group": bson.M{"_id": "$_id", "grampanchayat": bson.M{"$push": "$name"}, "grampanchayatCode": bson.M{"$push": "$uniqueId"}}},
									// {
									// 	"$addFields": bson.M{"grampanchayat": bson.M{"$arrayElemAt": []interface{}{"$grampanchayat", 0}}},
									// },
									// {
									// 	"$addFields": bson.M{"grampanchayatCode": bson.M{"$arrayElemAt": []interface{}{"$grampanchayatCode", 0}}},
									// },
									{
										"$lookup": bson.M{
											"from": constants.COLLECTIONVILLAGE,
											"as":   "villages",
											"let":  bson.M{"gramPanchayatCode": "$_id"},
											"pipeline": []bson.M{
												{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
													{"$eq": []string{"$status", "Active"}},
													{"$eq": []string{"$gramPanchayat", "$$gramPanchayatCode"}},
												}}}},
												bson.M{"$project": bson.M{"id": 1, "name": 1, "gramPanchayat": 1, "uniqueId": 1}},
												bson.M{"$sort": bson.M{"name": 1}},
												// {"$group": bson.M{"_id": "$_id", "village": bson.M{"$push": "$name"}, "villageCode": bson.M{"$push": "$uniqueId"}}},
												// {
												// 	"$addFields": bson.M{"village": bson.M{"$arrayElemAt": []interface{}{"$village", 0}}},
												// },
												// {
												// 	"$addFields": bson.M{"villageCode": bson.M{"$arrayElemAt": []interface{}{"$villageCode", 0}}},
												// },
											},
										}},
								},
							}},
					},
				}},
		}},
	})
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$addFields": bson.M{
	// 		"villages":       bson.M{"$arrayElemAt": []interface{}{"$villages", 0}},
	// 		"gramPanchayats": bson.M{"$arrayElemAt": []interface{}{"$gramPanchayats", 0}},
	// 		"blocks":         bson.M{"$arrayElemAt": []interface{}{"$blocks", 0}},
	// 		"districts":      bson.M{"$arrayElemAt": []interface{}{"$districts", 0}},
	// 	},
	// },
	// )
	//Aggregation
	d.Shared.BsonToJSONPrintTag("geodetails query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.GeoDetailsReport2
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}
func (d *Daos) GetActiveState(ctx *models.Context) ([]models.State, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": constants.STATESTATUSACTIVE})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("activestate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.State
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}
func (d *Daos) GetWeatherDataWithSeverityType(ctx *models.Context, filter *models.StateWeatherAlertMasterFilterv2, pagination *models.Pagination) ([]models.GetStateLeveWeatherDataAlert, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active"}})
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"state": 1}})

	query := []bson.M{}
	query = append(query, bson.M{"$eq": []string{"$severityType", "$$weathercode"}})
	query = append(query, bson.M{"$eq": []string{"$state", "$$statecode"}})
	if !filter.ParameterId.IsZero() {
		query = append(query, bson.M{"$eq": []interface{}{"$parameterid", filter.ParameterId}})

	}
	if !filter.Month.IsZero() {
		query = append(query, bson.M{"$eq": []interface{}{"$month", filter.Month}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONWEATHERALERTTYPE,
		"as":   "severitytype",
		"let":  bson.M{"statecode": "$_id"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}}}}}},
			//{"$project": bson.M{"name": 1}},
			{
				"$lookup": bson.M{
					"from": constants.COLLECTIONSTATEWEATHERALERTMASTER,
					"as":   "weatherdata",
					"let":  bson.M{"weathercode": "$_id"},
					"pipeline": []bson.M{
						{"$match": bson.M{"$expr": bson.M{"$and": query}}},
						//	{"$project": bson.M{"min": 1, "max": 1}},
					}}},
			bson.M{"$addFields": bson.M{"weatherdata": bson.M{"$arrayElemAt": []interface{}{"$weatherdata", 0}}}},
			bson.M{"$project": bson.M{
				"k":   "$name",
				"v":   "$$ROOT",
				"_id": 0,
			}},
		}},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{

		"severitytype": bson.M{"$arrayToObject": "$severitytype"}}})
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		d.Shared.BsonToJSONPrintTag("state pagenation query =>", paginationPipeline)
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("StateWeatherdataAlert query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.GetStateLeveWeatherDataAlert
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}
