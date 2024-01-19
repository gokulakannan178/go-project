package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveULB :""
func (d *Daos) SaveULB(ctx *models.Context, ulb *models.ULB) error {
	fmt.Println("adding ulb")
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).InsertOne(ctx.CTX, ulb)
	return err
}

//GetSingleULB : ""
func (d *Daos) GetSingleULB(ctx *models.Context, uniqueID string) (*models.RefULB, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})

	// LookUps
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Get Inventory
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ulbs []models.RefULB
	var ulb *models.RefULB
	if err = cursor.All(ctx.CTX, &ulbs); err != nil {
		return nil, err
	}
	if len(ulbs) > 0 {
		ulb = &ulbs[0]
	}
	return ulb, nil
}

//UpdateULB : ""
func (d *Daos) UpdateULB(ctx *models.Context, ulb *models.ULB) error {
	selector := bson.M{"uniqueId": ulb.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ulb}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//UpdateULB : ""
func (d *Daos) UpdateULBwithMobileno(ctx *models.Context, UniqueID string, MobileNo string) error {

	query := bson.M{"nodalOfficer.mobile": MobileNo, "uniqueId": bson.M{"$ne": UniqueID}}

	fmt.Println("query====>", query)

	count, err := ctx.DB.Collection(constants.COLLECTIONULB).CountDocuments(ctx.CTX, query)
	if err != nil {
		return err
	}
	if count == 0 {
		fmt.Println("success")
		return nil
	} else {
		return errors.New("already mobileno registered")

	}

	return nil
}

func IsPresent(arr []string, a string) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

//FilterULB : ""
func (d *Daos) FilterULB(ctx *models.Context, filter *models.ULBFilter, pagination *models.Pagination) ([]models.RefULB, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.Status) > 0 {

			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.TestCertStatus) > 0 {
			query = append(query, bson.M{"testcert.status": bson.M{"$in": filter.TestCertStatus}})
		}
		if len(filter.Address.DistrictCode) > 0 {
			query = append(query, bson.M{"address.districtCode": bson.M{"$in": filter.Address.DistrictCode}})
		}
		if len(filter.Address.StateCode) > 0 {
			query = append(query, bson.M{"address.stateCode": bson.M{"$in": filter.Address.StateCode}})
		}
		if len(filter.Address.BlockCode) > 0 {
			query = append(query, bson.M{"address.blockCode": bson.M{"$in": filter.Address.BlockCode}})
		}
		if len(filter.Address.VillageCode) > 0 {
			query = append(query, bson.M{"address.villageCode": bson.M{"$in": filter.Address.VillageCode}})
		}
		if len(filter.Address.GramPanchayatCode) > 0 {
			query = append(query, bson.M{"address.gramPanchayatCode": bson.M{"$in": filter.Address.GramPanchayatCode}})
		}
	}

	t := time.Now()
	ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())

	if filter.IsExpDate == "Yes" {
		query = append(query, bson.M{"testcert.expdate": bson.M{"$lte": ed}})

	}
	if filter.IsExpDate == "No" {
		query = append(query, bson.M{"testcert.expdate": bson.M{"$gte": ed}})

	}

	//Regex
	if filter.Regex.Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
	}
	if filter.Regex.NoName != "" {
		query = append(query, bson.M{"nodalOfficer.name": primitive.Regex{Pattern: filter.Regex.NoName, Options: "xi"}})
	}
	if filter.Regex.NoMobile != "" {
		query = append(query, bson.M{"nodalOfficer.mobile": primitive.Regex{Pattern: filter.Regex.NoMobile, Options: "xi"}})
	}
	if filter.Regex.CoName != "" {
		query = append(query, bson.M{"co.name": primitive.Regex{Pattern: filter.Regex.CoName, Options: "xi"}})
	}
	if filter.Regex.CoMobile != "" {
		query = append(query, bson.M{"co.mobile": primitive.Regex{Pattern: filter.Regex.CoMobile, Options: "xi"}})
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONULB).CountDocuments(ctx.CTX, func() bson.M {
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

	// LookUps
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	//get GP
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	//get block
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	//get district
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Get Inventory
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("ulb query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var ulbs []models.RefULB
	if err = cursor.All(ctx.CTX, &ulbs); err != nil {
		return nil, err
	}
	fmt.Println("=====================>", ulbs)
	return ulbs, nil
}

//EnableULB :""
func (d *Daos) EnableULB(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ULBSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableULB :""
func (d *Daos) DisableULB(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ULBSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteULB :""
func (d *Daos) DeleteULB(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ULBSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//AddULBTestCert : ""
func (d *Daos) AddULBTestCert(ctx *models.Context, UniqueID string, ulTbestCert *models.ULBTestCert) error {
	selector := bson.M{"uniqueId": UniqueID}
	data := bson.M{"$set": bson.M{
		"testcert": ulTbestCert,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, data)
	return err
}

//ApplyCert : ""
func (d *Daos) ApplyForTestCert(ctx *models.Context, UniqueID string, ulTbestCert *models.ULBTestCert) error {
	selector := bson.M{"uniqueId": UniqueID}
	data := bson.M{"$set": bson.M{
		"testcert.status":      ulTbestCert.Status,
		"testcert.appliedDate": ulTbestCert.AppliedDate,
		"testcert.appliedDoc":  ulTbestCert.AppliedDoc,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, data)
	return err
}

//ReApplyCert : ""
// the ReApplyForTestCert api will be used for which id have been rejected, they will be applied.
// the api written by balakrishnan.m@logikoof.in, and the written date is june 15,2022
func (d *Daos) ReApplyForTestCert(ctx *models.Context, UniqueID string, ulTbestCert *models.ULBTestCert) error {
	selector := bson.M{"uniqueId": UniqueID}
	data := bson.M{"$set": bson.M{
		"testcert.status":      ulTbestCert.Status,
		"testcert.appliedDate": ulTbestCert.AppliedDate,
		"testcert.appliedDoc":  ulTbestCert.AppliedDoc,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, data)
	return err
}

//AcceptTestCert : ""
func (d *Daos) AcceptTestCert(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	selector := bson.M{"uniqueId": UniqueID}

	ulbTestCert.Status = constants.ULBTESTCERTSTATUSACTIVE
	t := time.Now()
	t2 := t.AddDate(0, 6, 0)
	data := bson.M{"$set": bson.M{
		"testcert.gFCRating": ulbTestCert.GFCRating,
		// "testcert.starRatingDate":       ulbTestCert.StarRatingDate,
		// "testcert.starRatingExpiryDate": ulbTestCert.StarRatingExpiryDate,
		"testcert.status":  constants.ULBTESTCERTSTATUSACTIVE,
		"testcert.doc":     ulbTestCert.Doc,
		"testcert.regDate": t,
		"testcert.expdate": t2,
		"testcert.remarks": ulbTestCert.Remarks,
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, data)
	return err

}

// the RejectTestCert api will be change a testcert status as rejected.
// the api written by balakrishnan.m@logikoof.in, and the written date is june 15,2022
func (d *Daos) RejectTestCert(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	query := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update := bson.M{"$set": bson.M{
		"testcert.status":          constants.ULBTESTCERTSTATUSREJECTED,
		"testcert.remarks":         ulbTestCert.Remarks,
		"testcert.rejectedDate":    t,
		"testcert.rejected.by":     ulbTestCert.Rejected.By,
		"testcert.rejected.byType": ulbTestCert.Rejected.ByType,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//ULBTestCertStatus : ""
func (d *Daos) ULBTestCertStatus(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	selector := bson.M{"uniqueId": UniqueID}
	if ulbTestCert.Status == constants.ULBTESTCERTSTATUSACTIVE {
		ulbTestCert.Status = constants.ULBTESTCERTSTATUSACTIVE
		t := time.Now()
		t2 := t.AddDate(0, int(ulbTestCert.ExpDate.Month()), 0)
		data := bson.M{"$set": bson.M{
			"testcert.gFCRating":            ulbTestCert.GFCRating,
			"testcert.starRatingDate":       ulbTestCert.StarRatingDate,
			"testcert.starRatingExpiryDate": ulbTestCert.StarRatingExpiryDate,
			"testcert.status":               constants.ULBTESTCERTSTATUSACTIVE,
			"testcert.doc":                  ulbTestCert.Doc,
			"testcert.regDate":              t,
			"testcert.expdate":              t2,
			//"testcert.status":               ulbTestCert.Status,
		}}

		_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, data)
		return err
	}
	data := bson.M{"$set": bson.M{
		"testcert.status": ulbTestCert.Status,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, data)
	return err
}

// ULBInventoryUpdateMessasge : ""
func (d *Daos) ULBInventoryUpdateMessage(ctx *models.Context, filter *models.ULBInventoryUpdateMessageFilter) (*models.ULBInventoryUpdateMessageReport, error) {
	var mainpipeline, query []bson.M

	var startMonth, endMonth time.Time
	if filter.Date != nil {
		startMonth = time.Date(filter.Date.Year(), filter.Date.Month(), 1, 0, 0, 0, 0, filter.Date.Location())
		endMonth = time.Date(filter.Date.Year(), filter.Date.Month()+1, 1, 23, 59, 59, 0, filter.Date.Location())
		endMonth = endMonth.AddDate(0, 0, -1)
		query = append(query, bson.M{
			"created.on": bson.M{
				"$gte": startMonth,
				"$lte": endMonth,
			},
		})
	}

	if len(query) > 0 {
		mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainpipeline = append(mainpipeline, bson.M{"$group": bson.M{"_id": "$created.by"}})
	mainpipeline = append(mainpipeline, bson.M{"$group": bson.M{"_id": nil, "ulb": bson.M{"$push": "$_id"}}})
	// lookup
	mainpipeline = append(mainpipeline, bson.M{"$lookup": bson.M{
		"from": "ulb",
		"as":   "ulb",
		"let":  bson.M{"varUlb": "$ulb"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$status", "Active"}},
				//				bson.M{"$ne": []string{"$uniqueId", "$$varUlb"}},
				bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$in": []string{"$uniqueId", "$$varUlb"}}, "then": false, "else": true}}}},
			}}}},
			bson.M{"$match": bson.M{"nodalOfficer.mobile": bson.M{"$ne": ""}}},
			bson.M{"$match": bson.M{"nodalOfficer.mobile": bson.M{"$ne": nil}}},
			bson.M{"$group": bson.M{"_id": nil, "data": bson.M{"$push": bson.M{"ulbName": "$name", "noName": "$nodalOfficer.name", "noMobile": "$nodalOfficer.mobile"}}}},
		},
	}})
	mainpipeline = append(mainpipeline, bson.M{"$addFields": bson.M{"ulb": bson.M{"$arrayElemAt": []interface{}{"$ulb", 0}}}})
	mainpipeline = append(mainpipeline, bson.M{"$addFields": bson.M{"ulb": "$ulb.data"}})

	// Aggregation
	d.Shared.BsonToJSONPrintTag("farmerCart query =>", mainpipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBATCH).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var report []models.ULBInventoryUpdateMessageReport
	if err = cursor.All(context.TODO(), &report); err != nil {
		return nil, err
	}
	if len(report) > 0 {
		return &report[0], nil
	}
	return nil, nil

}

func (d *Daos) GetULBNotUpdatedInventory(ctx *models.Context, filter *models.ULBInventoryUpdateMessageFilter) ([]models.ULBLessData, error) {
	var mainpipeline, query []bson.M
	var startMonth, endMonth time.Time
	if filter != nil {
		if filter.Date == nil {
			return nil, errors.New("plz mention date")
		}
	}
	startMonth = time.Date(filter.Date.Year(), filter.Date.Month(), 1, 0, 0, 0, 0, filter.Date.Location())
	endMonth = time.Date(filter.Date.Year(), filter.Date.Month()+1, 1, 23, 59, 59, 0, filter.Date.Location())
	endMonth = endMonth.AddDate(0, 0, -1)
	query = append(query, bson.M{"status": constants.ULBSTATUSACTIVE})
	if len(query) > 0 {
		mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainpipeline = append(mainpipeline, bson.M{
		"$lookup": bson.M{
			"from": "batch", "as": "batch", "let": bson.M{"ulbId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$companyId", "$$ulbId"}},
					{"$eq": []string{"$productId", "PRODUCT1"}},
					{"$eq": []string{"$pkgType", "PKGTYPE00001"}},
					{"$gte": []interface{}{"$created.on", startMonth}},
					{"$lte": []interface{}{"$created.on", endMonth}},
				}}}},
			},
		}},
	)
	mainpipeline = append(mainpipeline, bson.M{"$addFields": bson.M{"batchCount": bson.M{"$size": "$batch"}}})
	mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"batchCount": bson.M{"$eq": 0}}})
	// Aggregation
	d.Shared.BsonToJSONPrintTag("GetULBNotUpdatedInventory query =>", mainpipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var report []models.ULBLessData
	if err = cursor.All(context.TODO(), &report); err != nil {
		return nil, err
	}
	return report, nil
}

//GetULBNotUpdatedInventoryV2 : "For Custom Month"
func (d *Daos) GetULBNotUpdatedInventoryV2(ctx *models.Context, filter *models.ULBInventoryUpdateMessageFilterV2) ([]models.ULBLessData, error) {
	var mainpipeline, query []bson.M
	var startMonth, endMonth time.Time
	if filter != nil {
		if filter.Date == nil {
			return nil, errors.New("plz mention date")
		}
	}
	startMonth = time.Date(filter.Date.From.Year(), filter.Date.From.Month(), filter.Date.From.Day(), 0, 0, 0, 0, filter.Date.From.Location())
	endMonth = time.Date(filter.Date.To.Year(), filter.Date.To.Month(), filter.Date.To.Day(), 23, 59, 59, 0, filter.Date.To.Location())
	// endMonth = endMonth.AddDate(0, 0, -1)
	query = append(query, bson.M{"status": constants.ULBSTATUSACTIVE})
	if len(query) > 0 {
		mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainpipeline = append(mainpipeline, bson.M{
		"$lookup": bson.M{
			"from": "batch", "as": "batch", "let": bson.M{"ulbId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$companyId", "$$ulbId"}},
					{"$eq": []string{"$productId", "PRODUCT1"}},
					{"$eq": []string{"$pkgType", "PKGTYPE00001"}},
					{"$gte": []interface{}{"$created.on", startMonth}},
					{"$lte": []interface{}{"$created.on", endMonth}},
				}}}},
			},
		}},
	)
	mainpipeline = append(mainpipeline, bson.M{"$addFields": bson.M{"batchCount": bson.M{"$size": "$batch"}}})
	mainpipeline = append(mainpipeline, bson.M{"$match": bson.M{"batchCount": bson.M{"$eq": 0}}})
	// Aggregation
	d.Shared.BsonToJSONPrintTag("GetULBNotUpdatedInventory query =>", mainpipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainpipeline, nil)
	if err != nil {
		return nil, err
	}
	var report []models.ULBLessData
	if err = cursor.All(context.TODO(), &report); err != nil {
		return nil, err
	}
	return report, nil
}
func (d *Daos) GetSingleMobileNoForULB(ctx *models.Context, mobileno string) (*models.RefULB, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"nodalOfficer.mobile": mobileno}})
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ulbs []models.RefULB
	var ulb *models.RefULB
	if err = cursor.All(ctx.CTX, &ulbs); err != nil {
		return nil, err
	}
	if len(ulbs) > 0 {
		ulb = &ulbs[0]
	}
	return ulb, nil
}
