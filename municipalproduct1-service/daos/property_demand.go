package daos

import (
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetPropertyDemandCalcQueryV2 : ""
func (d *Daos) GetPropertyDemandCalcQueryV2(ctx *models.Context, filter *models.PropertyDemandFilter, collectionName string) ([]bson.M, error) {
	var floorName string
	if collectionName == constants.COLLECTIONESTIMATEDPROPERTYDEMAND {
		floorName = constants.COLLECTIONESTIMATEDPROPERTYFLOOR
	} else {
		floorName = constants.COLLECTIONPROPERTYFLOOR
	}
	t := time.Now()
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "propertyconfiguration",
			"as":   "propertyConfig",
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"uniqueId": "1"}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"propertyConfig": bson.M{"$arrayElemAt": []interface{}{"$propertyConfig", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "penalcharges",
			"as":   "ref.penalCharges",
			"let":  bson.M{"doa": "$doa", "propertyTypeId": "$propertyTypeId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$$propertyTypeId", "$propertyTypeId"}},
					bson.M{"$gte": []string{"$$doa", "$doe"}},
				}}}},
				bson.M{"$sort": bson.M{"doe": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.penalCharges": bson.M{"$arrayElemAt": []interface{}{"$ref.penalCharges", 0}}}})
	// lookup for property other demand
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYOTHERDEMAND,
		"as":   "otherDemand",
		"let":  bson.M{"propertyId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
				bson.M{"$eq": []string{"$paymentStatus", constants.PROPERTYOTHERDEMANDPAYMENTSTATUSNOTPAID}},
				bson.M{"$eq": []string{"$status", constants.PROPERTYOTHERDEMANDSTATUSACTIVE}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil,
				"totalAmount": bson.M{"$sum": "$amount"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"otherDemand": bson.M{"$arrayElemAt": []interface{}{"$otherDemand.totalAmount", 0}}}})

	///
	mainPipeline = append(mainPipeline,

		bson.M{"$lookup": bson.M{
			"from": "propertypayments",
			"as":   "alreadyPayedMain",
			"let":  bson.M{"propertyId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
					// bson.M{"$eq": []string{"$fy.uniqueId", "$$fyId"}},
					//bson.M{"$eq": []string{"$status", "Completed"}},
				}}}},
				bson.M{"$match": bson.M{"status": bson.M{"$in": []string{"Completed", "VerificationPending"}}}},
				bson.M{"$group": bson.M{"_id": nil,

					"boreCharge": bson.M{"$sum": "$demand.boreCharge"},
					"formFee":    bson.M{"$sum": "$demand.formFee"},

					// "toBePaid":     bson.M{"$sum": []string{"$fy.tax", "$fy.vacantLandTax", "$fy.compositeTax", "$fy.ecess"}},
				}},
				//bson.M{"$addFields": bson.M{"amount": 40}},
			},
		}})
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{
			"alreadyPayedMain": bson.M{"$arrayElemAt": []interface{}{"$alreadyPayedMain", 0}},
		}})
	////

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONOTHERCHARGES,
			"as":   "otherCharges",
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"uniqueId": "GEN00001"}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"otherCharges": bson.M{"$arrayElemAt": []interface{}{"$otherCharges", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"percentAreaBuildup": bson.M{"$multiply": []interface{}{bson.M{"$divide": []string{"$builtUpArea", "$areaOfPlot"}}, 100}},
		//"taxableVacantLand":  bson.M{"$multiply": []interface{}{"$propertyConfig.taxableVacantLandConfig", bson.M{"$subtract": []string{"$areaOfPlot", "$builtUpArea"}}}}

	}})

	floorPipeline := []bson.M{}
	floorPipeline = append(floorPipeline, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
		bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
		bson.M{"$eq": []string{"$status", constants.PROPERTYFLOORSTATUSACTIVE}},
		bson.M{"$gte": []string{"$$fydateTo", "$dateFrom"}},
		// bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$dateTo", nil}}, "then": true, "else": bson.M{"$cond": bson.M{"if": bson.M{"$gte": []string{"$dateTo", "$$fyfromDate"}}, "then": true, "else": false}}}}}},
		bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$not": []interface{}{"$dateTo"}}, "then": true, "else": bson.M{"$cond": bson.M{"if": bson.M{"$gte": []string{"$dateTo", "$$fyfromDate"}}, "then": true, "else": false}}}}}},
	}}}})
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "usageType", "uniqueId", "ref.usageType", "ref.usageType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "constructionType", "uniqueId", "ref.constructionType", "ref.constructionType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONOCCUMANCYTYPE, "occupancyType", "uniqueId", "ref.occupancyType", "ref.occupancyType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR, "nonResUsageType", "uniqueId", "ref.nonResUsageType", "ref.nonResUsageType")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONFLOORRATABLEAREA, "ratableAreaType", "uniqueId", "ref.floorRatableArea", "ref.floorRatableArea")...)
	floorPipeline = append(floorPipeline, d.CommonLookup(constants.COLLECTIONFLOORTYPE, "no", "uniqueId", "ref.floorNo", "ref.floorNo")...)
	floorPipeline = append(floorPipeline,
		bson.M{
			"$lookup": bson.M{
				"from": constants.COLLECTIONAVR,
				"as":   "ref.avr",
				"let":  bson.M{"address": "$$address", "fyfromDate": "$$fyfromDate", "fydateTo": "$$fydateTo", "roadType": "$$roadType", "propertyId": "$$propertyId", "municipalityType": "$$municipalityType", "constructionType": "$constructionType", "usageType": "$usageType"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$municipalityTypeId", "$$municipalityType"}},
						bson.M{"$eq": []string{"$constructionTypeId", "$$constructionType"}},
						bson.M{"$eq": []string{"$roadTypeId", "$$roadType"}},
						bson.M{"$eq": []string{"$usageTypeId", "$$usageType"}},
						bson.M{"$eq": []string{"$zoneId", "$$address.zoneCode"}},

						//bson.M{"$eq": []string{"$status", "Active"}},
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				},
			},
		})
	floorPipeline = append(floorPipeline,
		bson.M{
			"$lookup": bson.M{
				"from": constants.COLLECTIONCOMPOSITETAXRATEMASTER,
				"as":   "ref.compositeTax",
				"let":  bson.M{"propertyTypeId": "$propertyTypeId", "address": "$$address", "fyfromDate": "$$fyfromDate", "fydateTo": "$$fydateTo", "roadType": "$$roadType", "propertyId": "$$propertyId", "municipalityType": "$$municipalityType", "constructionType": "$constructionType", "usageType": "$usageType", "buildUpArea": "$buildUpArea"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						// bson.M{"$eq": []string{"$municipalityTypeId", "$$municipalityType"}},
						bson.M{"$eq": []string{"$usageType", "$$propertyTypeId"}},
						bson.M{"$gte": []string{"$$buildUpArea", "$minBuildUpArea"}},
						bson.M{"$lte": []string{"$$buildUpArea", "$maxBuildUpArea"}},
						bson.M{"$eq": []string{"$status", "Active"}},
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				},
			},
		})
	floorPipeline = append(floorPipeline,
		bson.M{
			"$lookup": bson.M{
				"from": "vacantlandrates",
				"as":   "ref.vlr",
				"let":  bson.M{"address": "$$address", "fyfromDate": "$$fyfromDate", "fydateTo": "$$fydateTo", "roadType": "$$roadType", "propertyId": "$$propertyId", "municipalityType": "$$municipalityType", "constructionType": "$constructionType", "usageType": "$usageType"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						// bson.M{"$eq": []string{"$constructionTypeId", "$$constructionType"}},
						bson.M{"$eq": []string{"$roadTypeId", "$$roadType"}},
						bson.M{"$eq": []string{"$usageTypeId", "$$usageType"}},
						bson.M{"$eq": []string{"$zoneId", "$$address.zoneCode"}},

						//bson.M{"$eq": []string{"$status", "Active"}},
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				},
			},
		})

	/*
		bson.M{"$lookup": bson.M{
				"from": "vacantlandrates",
				"as":   "vlr",
				"let":  bson.M{"tempFrom": "$from", "tempTo": "$to", "mt": "$$tempMT", "rt": "$$tempRT", "doa": "$$tempDOA"},
				"pipeline": func() []bson.M {

					return []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							bson.M{"$eq": []string{"$municipalityTypeId", "$$mt"}},
							bson.M{"$eq": []string{"$roadTypeId", "$$rt"}},
							//bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
							bson.M{"$lte": []string{"$doe", "$$tempTo"}},
							// bson.M{"$gte": []string{"$doe", "$$doa"}},
						}}}},
						bson.M{"$sort": bson.M{"doe": -1}},
						bson.M{"$limit": 1},
					}
				}(),
			},
	*/

	floorPipeline = append(floorPipeline, bson.M{
		"$addFields": bson.M{
			"ref.avr":          bson.M{"$arrayElemAt": []interface{}{"$ref.avr", 0}},
			"ref.vlr":          bson.M{"$arrayElemAt": []interface{}{"$ref.vlr", 0}},
			"ref.compositeTax": bson.M{"$arrayElemAt": []interface{}{"$ref.compositeTax", 0}},
		},
	})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "completedFys",
		"let":  bson.M{"propertyId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
				bson.M{"$eq": []string{"$status", "Completed"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "ids": bson.M{"$push": "$fy.uniqueId"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"completedFys": bson.M{"$arrayElemAt": []interface{}{"$completedFys", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"completedFys.ids": bson.M{"$cond": bson.M{"if": bson.M{"$not": []interface{}{"$completedFys.ids"}}, "then": []interface{}{}, "else": "$completedFys.ids"}}}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "financialyears",
		"as":   "noteFy",
		"let":  bson.M{"completedFys": "$completedFys.ids", "tempDOA": "$doa", "propertyId": "$uniqueId", "tempMT": "$municipalityId", "tempRT": "$roadTypeId", "taxablevland": "$taxableVacantLand", "percentAreaBuiltUp": "$percentAreaBuildup", "propertyConfig": "$propertyConfig", "tempStatus": "$status", "tempEndDate": "$endDate", "address": "$address", "propertyTypeId": "$propertyTypeId"},
		"pipeline": []bson.M{
			bson.M{"$sort": bson.M{"order": -1}},
			// bson.M{"$match": bson.M{"$expr": bson.M{"$and": func() []bson.M {
			// 	fySelectionAndQuery := make([]bson.M, 0)
			// 	fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$lt": []string{"$$tempDOA", "$to"}})
			// 	// if filter != nil {
			// 	// 	if filter.IsOmitPaidFys {
			// 	// 		return fySelectionAndQuery
			// 	// 	}
			// 	// }
			// 	// fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$cond": bson.M{"if": bson.M{"$in": []interface{}{"$uniqueId", "$$completedFys"}}, "then": false, "else": true}})
			// 	return fySelectionAndQuery
			// }(),
			// }}},
			//

			bson.M{"$match": bson.M{"$expr": bson.M{"$and": func() []bson.M {
				fySelectionAndQuery := make([]bson.M, 0)
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$eq": []string{"Active", "$status"}})
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$eq": []interface{}{false, "$isCurrent"}})
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$lt": []string{"$$tempDOA", "$to"}})
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$gt": []interface{}{
					bson.M{"$ifNull": []interface{}{"$$tempEndDate", &t}},
					"$from",
				}})
				return fySelectionAndQuery
			}(),
			}}},
			bson.M{"$group": bson.M{"_id": nil, "fys": bson.M{"$push": "$name"}}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"noteFy": bson.M{"$arrayElemAt": []interface{}{"$noteFy", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"noteFy": "$noteFy.fys"}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "financialyears",
		"as":   "fys",
		"let":  bson.M{"completedFys": "$completedFys.ids", "tempDOA": "$doa", "propertyId": "$uniqueId", "tempMT": "$municipalityId", "tempRT": "$roadTypeId", "taxablevland": "$taxableVacantLand", "percentAreaBuiltUp": "$percentAreaBuildup", "propertyConfig": "$propertyConfig", "tempStatus": "$status", "tempEndDate": "$endDate", "address": "$address", "propertyTypeId": "$propertyTypeId"},
		"pipeline": []bson.M{
			bson.M{"$sort": bson.M{"order": 1}},
			// bson.M{"$match": bson.M{"$expr": bson.M{"$and": func() []bson.M {
			// 	fySelectionAndQuery := make([]bson.M, 0)
			// 	fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$lt": []string{"$$tempDOA", "$to"}})
			// 	// if filter != nil {
			// 	// 	if filter.IsOmitPaidFys {
			// 	// 		return fySelectionAndQuery
			// 	// 	}
			// 	// }
			// 	// fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$cond": bson.M{"if": bson.M{"$in": []interface{}{"$uniqueId", "$$completedFys"}}, "then": false, "else": true}})
			// 	return fySelectionAndQuery
			// }(),
			// }}},
			//

			bson.M{"$match": bson.M{"$expr": bson.M{"$and": func() []bson.M {
				fySelectionAndQuery := make([]bson.M, 0)
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$eq": []string{"Active", "$status"}})
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$lt": []string{"$$tempDOA", "$to"}})
				fySelectionAndQuery = append(fySelectionAndQuery, bson.M{"$gt": []interface{}{
					bson.M{"$ifNull": []interface{}{"$$tempEndDate", &t}},
					"$from",
				}})
				return fySelectionAndQuery
			}(),
			}}},
			//
			bson.M{"$lookup": bson.M{
				"from": "propertypaymentfys",
				"as":   "alreadyPayed",
				"let":  bson.M{"propertyId": "$$propertyId", "fyId": "$uniqueId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
						bson.M{"$eq": []string{"$fy.uniqueId", "$$fyId"}},
						bson.M{"$eq": []string{"$status", "Completed"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,

						"amount":       bson.M{"$sum": "$fy.totalTax"},
						"fyTax":        bson.M{"$sum": "$fy.tax"},
						"vlTax":        bson.M{"$sum": "$fy.vacantLandTax"},
						"penalty":      bson.M{"$sum": "$fy.penanty"},
						"rebate":       bson.M{"$sum": "$fy.rebate"},
						"compositeTax": bson.M{"$sum": "$fy.compositeTax"},
						"ecess":        bson.M{"$sum": "$fy.ecess"},
						"otherDemand":  bson.M{"$sum": "$fy.otherDemand"},

						// "toBePaid":     bson.M{"$sum": []string{"$fy.tax", "$fy.vacantLandTax", "$fy.compositeTax", "$fy.ecess"}},
					}},
					//bson.M{"$addFields": bson.M{"amount": 40}},
				},
			}},

			bson.M{"$addFields": bson.M{
				"alreadyPayed": bson.M{"$arrayElemAt": []interface{}{"$alreadyPayed", 0}},
			}},
			//bson.M{"$match": bson.M{"fy.totalTax": bson.M{"$gt": 0}}},

			// lookup for property other demand
			bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONPROPERTYOTHERDEMAND,
				"as":   "otherDemand",
				"let":  bson.M{"propertyId": "$$propertyId", "fyId": "$uniqueId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
						bson.M{"$eq": []string{"$fyId", "$$fyId"}},
						bson.M{"$eq": []string{"$paymentStatus", constants.PROPERTYOTHERDEMANDPAYMENTSTATUSNOTPAID}},
						bson.M{"$eq": []string{"$status", constants.PROPERTYOTHERDEMANDSTATUSACTIVE}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"totalAmount": bson.M{"$sum": "$amount"},
					}},
				},
			}},
			bson.M{"$addFields": bson.M{"otherDemand": bson.M{"$arrayElemAt": []interface{}{"$otherDemand.totalAmount", 0}}}},

			// lookup for property other demand - additional penal charge

			bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONPROPERTYOTHERDEMAND,
				"as":   "otherDemandAdditionalPenalty",
				"let":  bson.M{"propertyId": "$$propertyId", "fyId": "$uniqueId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
						bson.M{"$eq": []string{"$fyId", "$$fyId"}},
						bson.M{"$eq": []string{"$oneTimePenalCharges", "Yes"}},
						bson.M{"$eq": []string{"$status", constants.PROPERTYOTHERDEMANDSTATUSACTIVE}},
					}}}},
				},
			}},
			bson.M{"$addFields": bson.M{"otherDemandAdditionalPenalty": bson.M{"$arrayElemAt": []interface{}{"$otherDemandAdditionalPenalty", 0}}}},

			bson.M{"$match": func() bson.M {
				if len(filter.Fys) > 0 {
					return bson.M{"uniqueId": bson.M{"$in": filter.Fys}}
				}
				return bson.M{}
			}()},
			bson.M{"$lookup": bson.M{
				"from": "vacantlandrates",
				"as":   "vlr",
				"let":  bson.M{"tempFrom": "$from", "tempTo": "$to", "mt": "$$tempMT", "rt": "$$tempRT", "doa": "$$tempDOA"},
				"pipeline": func() []bson.M {

					return []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							bson.M{"$eq": []string{"$municipalityTypeId", "$$mt"}},
							bson.M{"$eq": []string{"$roadTypeId", "$$rt"}},
							//bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
							bson.M{"$lte": []string{"$doe", "$$tempTo"}},
							// bson.M{"$gte": []string{"$doe", "$$doa"}},
						}}}},
						bson.M{"$sort": bson.M{"doe": -1}},
						bson.M{"$limit": 1},
					}
				}(),
			},
			},
			bson.M{"$addFields": bson.M{"vlr": bson.M{"$arrayElemAt": []interface{}{"$vlr", 0}}}},
			//bson.M{"$addFields": bson.M{"vacantLandTax": bson.M{"$cond": bson.M{"if": bson.M{"$lt": []string{"$$percentAreaBuiltUp", "$$propertyConfig.vacantLandRatePercentage"}}, "then": bson.M{"$multiply": []string{"$$taxablevland", "$vlr.rate"}}, "else": 0}}}},
			bson.M{"$lookup": bson.M{
				"as":   "ref.propertyTax",
				"from": constants.COLLECTIONPROPERTYTAX,
				"let":  bson.M{"fydateTo": "$to", "fyfromDate": "$from"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				}}},
			bson.M{"$addFields": bson.M{"ref.propertyTax": bson.M{"$arrayElemAt": []interface{}{"$ref.propertyTax", 0}}}},

			bson.M{"$lookup": bson.M{
				"as":   "ref.penalty",
				"from": constants.COLLECTIONPENALTY,
				"let":  bson.M{"fydateTo": "$to", "fyfromDate": "$from"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				}}},
			bson.M{"$addFields": bson.M{"ref.penalty": bson.M{"$arrayElemAt": []interface{}{"$ref.penalty", 0}}}},

			bson.M{
				"$lookup": bson.M{
					"from":     floorName,
					"as":       "floors",
					"let":      bson.M{"propertyTypeId": "$$propertyTypeId", "fyfromDate": "$from", "fydateTo": "$to", "roadType": "$$tempRT", "propertyId": "$$propertyId", "municipalityType": "$$tempMT", "propertyTax": "$ref.propertyTax", "address": "$$address"},
					"pipeline": floorPipeline,
				},
			},
			bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONLEGACYYEAR,
				"as":   "legacy",
				"let":  bson.M{"fyId": "$uniqueId", "propertyId": "$$propertyId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$status", constants.LEGACYPROPERTYFYSTATUSACTIVE}},
						bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
						bson.M{"$eq": []string{"$fyId", "$$fyId"}},
					}}}},
				},
			},
			},
			bson.M{"$addFields": bson.M{"legacy": bson.M{"$arrayElemAt": []interface{}{"$legacy", 0}}}},
			bson.M{
				"$lookup": bson.M{
					"from": constants.COLLECTIONPROPERTYFIXEDARV,
					"as":   "fixedArv",
					"let":  bson.M{"fyId": "$uniqueId", "propertyId": "$$propertyId"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							bson.M{"$eq": []string{"$status", constants.PROPERTYFIXEDARVSTATUSACTIVE}},
							bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
							bson.M{"$eq": []string{"$fyId", "$$fyId"}},
						}}}},
					},
				},
			},
			bson.M{"$addFields": bson.M{"fixedArv": bson.M{"$arrayElemAt": []interface{}{"$fixedArv", 0}}}},
			bson.M{
				"$lookup": bson.M{
					"from": constants.COLLECTIONPROPERTYFIXEDDEMAND,
					"as":   "fixedDemand",
					"let":  bson.M{"fyId": "$uniqueId", "propertyId": "$$propertyId"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							bson.M{"$eq": []string{"$status", constants.PROPERTYFIXEDARVSTATUSACTIVE}},
							bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
							bson.M{"$eq": []string{"$fyId", "$$fyId"}},
						}}}},
					},
				},
			},
			bson.M{"$addFields": bson.M{"fixedDemand": bson.M{"$arrayElemAt": []interface{}{"$fixedDemand", 0}}}},

			bson.M{
				"$lookup": bson.M{
					"from": floorName,
					"as":   "floorBuildupArea",
					"let":  bson.M{"propertyTypeId": "$$propertyTypeId", "fyfromDate": "$from", "fydateTo": "$to", "roadType": "$$tempRT", "propertyId": "$$propertyId", "municipalityType": "$$tempMT", "propertyTax": "$ref.propertyTax", "address": "$$address"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							bson.M{"$eq": []string{"$propertyId", "$$propertyId"}},
							bson.M{"$eq": []string{"$status", constants.PROPERTYFLOORSTATUSACTIVE}},
							bson.M{"$gte": []string{"$$fydateTo", "$dateFrom"}},
							// bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$dateTo", nil}}, "then": true, "else": bson.M{"$cond": bson.M{"if": bson.M{"$gte": []string{"$dateTo", "$$fyfromDate"}}, "then": true, "else": false}}}}}},
							bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$not": []interface{}{"$dateTo"}}, "then": true, "else": bson.M{"$cond": bson.M{"if": bson.M{"$gte": []string{"$dateTo", "$$fyfromDate"}}, "then": true, "else": false}}}}}},
						}}}},
						bson.M{
							"$group": bson.M{"_id": nil, "area": bson.M{"$sum": "$buildUpArea"}},
						},
					},
				},
			},
			bson.M{"$addFields": bson.M{"floorBuildupArea": bson.M{"$arrayElemAt": []interface{}{"$floorBuildupArea", 0}}}},
			bson.M{
				"$lookup": bson.M{
					"from": constants.COLLECTIONCOMPOSITETAXRATEMASTER,
					"as":   "compositeTaxRate",
					"let":  bson.M{"propertyTypeId": "$$propertyTypeId", "floorBuildupArea": "$floorBuildupArea", "fyfromDate": "$from", "fydateTo": "$to", "usageType": "$usageType"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							// bson.M{"$eq": []string{"$municipalityTypeId", "$$municipalityType"}},
							bson.M{"$eq": []string{"$usageType", "$$propertyTypeId"}},
							bson.M{"$gte": []string{"$$floorBuildupArea.area", "$minBuildUpArea"}},
							bson.M{"$lte": []string{"$$floorBuildupArea.area", "$maxBuildUpArea"}},
							bson.M{"$eq": []string{"$status", "Active"}},
							bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
						}}}},
						bson.M{"$sort": bson.M{"doe": -1}},
					},
				},
			},
			bson.M{"$addFields": bson.M{"compositeTaxRate": bson.M{"$arrayElemAt": []interface{}{"$compositeTaxRate", 0}}}},
			bson.M{
				"$lookup": bson.M{
					"from": "panelchargemaster",
					"as":   "panelChargeRatemaster",
					"let":  bson.M{"propertyTypeId": "$$propertyTypeId", "floorBuildupArea": "$floorBuildupArea", "fyfromDate": "$from", "fydateTo": "$to", "usageType": "$usageType"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							// bson.M{"$eq": []string{"$municipalityTypeId", "$$municipalityType"}},
							// bson.M{"$eq": []string{"$usageType", "$$propertyTypeId"}},
							// bson.M{"$gte": []string{"$$floorBuildupArea.area", "$minBuildUpArea"}},
							// bson.M{"$lte": []string{"$$floorBuildupArea.area", "$maxBuildUpArea"}},
							bson.M{"$eq": []string{"$status", "Active"}},
							bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
						}}}},
						bson.M{"$sort": bson.M{"doe": -1}},
					},
				},
			},
			bson.M{"$addFields": bson.M{"panelChargeRatemaster": bson.M{"$arrayElemAt": []interface{}{"$panelChargeRatemaster", 0}}}},
			bson.M{
				"$lookup": bson.M{
					"from": "ecessratemaster",
					"as":   "ecessRateMaster",
					"let":  bson.M{"propertyTypeId": "$$propertyTypeId", "floorBuildupArea": "$floorBuildupArea", "fyfromDate": "$from", "fydateTo": "$to", "usageType": "$usageType"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							// bson.M{"$eq": []string{"$municipalityTypeId", "$$municipalityType"}},
							// bson.M{"$eq": []string{"$usageType", "$$propertyTypeId"}},
							// bson.M{"$gte": []string{"$$floorBuildupArea.area", "$minBuildUpArea"}},
							// bson.M{"$lte": []string{"$$floorBuildupArea.area", "$maxBuildUpArea"}},
							bson.M{"$eq": []string{"$status", "Active"}},
							bson.M{"$lte": []string{"$doe", "$$fydateTo"}},
						}}}},
						bson.M{"$sort": bson.M{"doe": -1}},
					},
				},
			},
			bson.M{"$addFields": bson.M{"ecessRateMaster": bson.M{"$arrayElemAt": []interface{}{"$ecessRateMaster", 0}}}},
		},
	}})
	return mainPipeline, nil
}
