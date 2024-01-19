package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveDumpSite :""
func (d *Daos) SaveDumpSite(ctx *models.Context, DumpSite *models.DumpSite) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).InsertOne(ctx.CTX, DumpSite)

	//DumpSite.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

//GetSingleDumpSite : ""
func (d *Daos) GetSingleDumpSite(ctx *models.Context, uniqueID string) (*models.RefDumpSite, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})

	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCIRCLE, "address.circleCode", "code", "ref.circle", "ref.circle")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.sector", "ref.sector")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadtype", "uniqueId", "ref.roadtype", "ref.roadtype")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.state", "ref.state")...)

	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONDumpSiteLOG,
	// 		"as":   "ref.DumpSitelog",
	// 		"let":  bson.M{"DumpSiteId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$status", constants.DumpSiteASSIGNSTATUS}},
	// 				{"$eq": []string{"$DumpSiteId", "$$DumpSiteId"}},
	// 			}}},
	// 			},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.DumpSitelog": bson.M{"$arrayElemAt": []interface{}{"$ref.DumpSitelog", 0}}}})
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DumpSites []models.RefDumpSite
	var DumpSite *models.RefDumpSite
	if err = cursor.All(ctx.CTX, &DumpSites); err != nil {
		return nil, err
	}
	if len(DumpSites) > 0 {
		DumpSite = &DumpSites[0]
	}
	return DumpSite, nil
}

//UpdateDumpSite : ""
func (d *Daos) UpdateDumpSite(ctx *models.Context, DumpSite *models.DumpSite) error {
	selector := bson.M{"uniqueId": DumpSite.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": DumpSite}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// func (d *Daos) UpdateDumpSiteAssign(ctx *models.Context, DumpSite *models.DumpSiteAssign) error {
// 	selector := bson.M{"uniqueId": DumpSite.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": DumpSite}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDumpSite).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

//EnableDumpSite :""
func (d *Daos) EnableDumpSite(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DUMPSITESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDumpSite :""
func (d *Daos) DisableDumpSite(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DUMPSITESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDumpSite :""
func (d *Daos) DeleteDumpSite(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DUMPSITESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDumpSite : ""
func (d *Daos) FilterDumpSite(ctx *models.Context, DumpSiteFilter *models.FilterDumpSite, pagination *models.Pagination) ([]models.RefDumpSite, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DumpSiteFilter != nil {

		if len(DumpSiteFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DumpSiteFilter.Status}})
		}
		//Regex
		if DumpSiteFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: DumpSiteFilter.Regex.Name, Options: "xi"}})
		}
		if DumpSiteFilter.Regex.OwnerName != "" {
			query = append(query, bson.M{"ownerName": primitive.Regex{Pattern: DumpSiteFilter.Regex.OwnerName, Options: "xi"}})
		}
		if DumpSiteFilter.Regex.ManagerName != "" {
			query = append(query, bson.M{"minUser.name": primitive.Regex{Pattern: DumpSiteFilter.Regex.ManagerName, Options: "xi"}})
		}

	}
	if DumpSiteFilter.DateRange.From != nil {
		t := *DumpSiteFilter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if DumpSiteFilter.DateRange.To != nil {
			t2 := *DumpSiteFilter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if DumpSiteFilter != nil {
		if DumpSiteFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{DumpSiteFilter.SortBy: DumpSiteFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCIRCLE, "address.circleCode", "code", "ref.circle", "ref.circle")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.sector", "ref.sector")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadtype", "uniqueId", "ref.roadtype", "ref.roadtype")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.state", "ref.state")...)

	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONDumpSiteLOG,
	// 		"as":   "ref.DumpSitelog",
	// 		"let":  bson.M{"DumpSiteId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$status", constants.DumpSiteASSIGNSTATUS}},
	// 				{"$eq": []string{"$DumpSiteId", "$$DumpSiteId"}},
	// 			}}},
	// 			},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.DumpSitelog": bson.M{"$arrayElemAt": []interface{}{"$ref.DumpSitelog", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DumpSite query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DumpSitesFilter []models.RefDumpSite
	if err = cursor.All(context.TODO(), &DumpSitesFilter); err != nil {
		return nil, err
	}
	return DumpSitesFilter, nil
}

// func (d *Daos) DumpSiteAssign(ctx *models.Context, DumpSite *models.DumpSiteAssign) error {
// 	selector := bson.M{"uniqueId": DumpSite.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": DumpSite}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDumpSite).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// func (d *Daos) RevokeDumpSite(ctx *models.Context, DumpSite *models.DumpSite) error {
// 	selector := bson.M{"employeeId": DumpSite.EmployeeId}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": DumpSite, "status": constants.DumpSiteREVOKESTATUS}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDumpSite).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// GetSingleDumpSiteUsingEmpID : ""
func (d *Daos) GetSingleDumpSiteUsingUniqueId(ctx *models.Context, UniqueID string) (*models.RefDumpSite, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID, "status": bson.M{"$in": []string{constants.DUMPSITESTATUSACTIVE}}}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPSITE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DumpSites []models.RefDumpSite
	var DumpSite *models.RefDumpSite
	if err = cursor.All(ctx.CTX, &DumpSites); err != nil {
		return nil, err
	}
	if len(DumpSites) > 0 {
		DumpSite = &DumpSites[0]
	}
	return DumpSite, nil
}
