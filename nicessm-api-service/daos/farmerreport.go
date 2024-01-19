package daos

import (
	"context"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FilterFarmerReport : ""
func (d *Daos) FilterFarmerReport(ctx *models.Context, filter *models.FarmerReportFilter, pagination *models.Pagination) ([]models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTFARMER, bson.M{
		"farmerId": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.PROJECTFARMERSTATUSACTIVE}},
			{"$eq": []string{"$farmer", "$$farmerId"}},
		}}}},
	}, "ref.projects", "ref.projects")...)
	query := []bson.M{}
	if filter != nil {

		if !filter.Village.ID.IsZero() {
			query = append(query, bson.M{"village": bson.M{"$" + filter.Village.Condition: filter.Village.ID}})
		}
		if !filter.State.ID.IsZero() {
			query = append(query, bson.M{"state": bson.M{"$" + filter.State.Condition: filter.State.ID}})
		}
		if !filter.District.ID.IsZero() {
			query = append(query, bson.M{"district": bson.M{"$" + filter.District.Condition: filter.District.ID}})
		}
		if !filter.GramPanchayat.ID.IsZero() {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$" + filter.GramPanchayat.Condition: filter.GramPanchayat.ID}})
		}
		if !filter.Block.ID.IsZero() {
			query = append(query, bson.M{"block": bson.M{"$" + filter.Block.Condition: filter.Block.ID}})
		}
		if filter.Education.ID != "" {
			query = append(query, bson.M{"education": bson.M{"$" + filter.Education.Condition: filter.Education.ID}})
		}
		if !filter.Asset.ID.IsZero() {
			query = append(query, bson.M{"assert": bson.M{"$" + filter.Asset.Condition: filter.Asset.ID}})
		}
		if filter.YearlyIncome.ID != "" {
			query = append(query, bson.M{"yearlyIncome": bson.M{"$" + filter.YearlyIncome.Condition: filter.YearlyIncome.ID}})
		}
		if filter.CreatedDate.ID != nil {
			query = append(query, bson.M{"createdDate": bson.M{"$" + filter.CreatedDate.Condition: filter.CreatedDate.ID}})
		}
		if filter.VoiceSmsStatus.ID != constants.REPORTFALSE {
			query = append(query, bson.M{"voiceSmsStatus": bson.M{"$" + filter.VoiceSmsStatus.Condition: filter.VoiceSmsStatus.ID}})
		}
		if filter.Gender.ID != "" {
			query = append(query, bson.M{"gender": bson.M{"$" + filter.Gender.Condition: filter.Gender.ID}})
		}
		if filter.SmsStatus.ID != constants.REPORTFALSE {
			query = append(query, bson.M{"smsStatus": bson.M{"$" + filter.SmsStatus.Condition: filter.SmsStatus.ID}})
		}
		if !filter.Organisation.ID.IsZero() {
			query = append(query, bson.M{"farmerOrg": bson.M{"$" + filter.Organisation.Condition: filter.Organisation.ID}})
		}
		if !filter.Project.ID.IsZero() {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$" + filter.Project.Condition: filter.Project.ID}})
		}
		if filter.Age.ID != 0 {
			//	var findage time.Time
			findage := time.Now().Year() - filter.Age.ID
			t := time.Date(findage, time.December, 31, 0, 0, 0, 0, time.Local)
			t2 := time.Date(findage, time.January, 1, 0, 0, 0, 0, time.Local)
			//t2.AddDate()
			//	query = append(query, bson.M{"dateOfBirth": bson.M{"$" + filter.Age.Condition: t.Year()}})
			query = append(query, bson.M{"dateOfBirth": bson.M{"$gte": t2, "$lte": t}})
		}

		//Regex
		// if Queryfilter.Regex.Query != "" {
		// 	query = append(query, bson.M{"query": primitive.Regex{Pattern: Queryfilter.Regex.Query, Options: "xi"}})
		// }
		if filter.CreatedFrom.Date != nil {
			var sd, ed time.Time
			var sdcondition, edcondition string = "gte", "lte"
			sd = time.Date(filter.CreatedFrom.Date.Year(), filter.CreatedFrom.Date.Month(), filter.CreatedFrom.Date.Day(), 0, 0, 0, 0, filter.CreatedFrom.Date.Location())
			ed = time.Date(filter.CreatedFrom.Date.Year(), filter.CreatedFrom.Date.Month(), filter.CreatedFrom.Date.Day(), 23, 59, 59, 0, filter.CreatedFrom.Date.Location())
			sdcondition = filter.CreatedFrom.Condition

			if filter.CreatedTo.Date != nil {
				ed = time.Date(filter.CreatedTo.Date.Year(), filter.CreatedTo.Date.Month(), filter.CreatedTo.Date.Day(), 23, 59, 59, 0, filter.CreatedTo.Date.Location())
				edcondition = filter.CreatedTo.Condition
			}
			query = append(query, bson.M{"date": bson.M{"$" + sdcondition: sd, "$" + edcondition: ed}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$" + filter.Condition: query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("Farmer pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "farmerOrg", "_id", "ref.farmerOrg", "ref.farmerOrg")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "soilType", "_id", "ref.soilType", "ref.soilType")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSET, "assert", "_id", "ref.assert", "ref.assert")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &Farmers); err != nil {
		return nil, err
	}
	return Farmers, nil
}

//FilterFarmer : ""
func (d *Daos) FilterFarmerReport2(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) ([]models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTFARMER, bson.M{
		"farmerId": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.PROJECTUSERSTATUSACTIVE}},
			{"$eq": []string{"$farmer", "$$farmerId"}},
		}}}},
	}, "ref.projects", "ref.projects")...)
	query := []bson.M{}
	if Farmerfilter != nil {

		if len(Farmerfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Farmerfilter.ActiveStatus}})
		}
		if len(Farmerfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Farmerfilter.Status}})
		}
		if len(Farmerfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": Farmerfilter.State}})
		}
		if len(Farmerfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": Farmerfilter.District}})
		}
		if len(Farmerfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": Farmerfilter.Block}})
		}
		if len(Farmerfilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": Farmerfilter.GramPanchayat}})
		}
		if len(Farmerfilter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": Farmerfilter.Village}})
		}
		if len(Farmerfilter.FarmerOrg) > 0 {
			query = append(query, bson.M{"farmerOrg": bson.M{"$in": Farmerfilter.FarmerOrg}})
		}
		//Regex
		if Farmerfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Farmerfilter.Regex.Name, Options: "xi"}})
		}
		if Farmerfilter.Regex.MobileNumber != "" {
			query = append(query, bson.M{"mobileNumber": primitive.Regex{Pattern: Farmerfilter.Regex.MobileNumber, Options: "xi"}})
		}
		if Farmerfilter.Regex.SpouseName != "" {
			query = append(query, bson.M{"spouseName": primitive.Regex{Pattern: Farmerfilter.Regex.SpouseName, Options: "xi"}})
		}
		if Farmerfilter.OmitProjectFarmer.Is {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$ne": Farmerfilter.OmitProjectFarmer.Project}})
		}
		//daterange
		if Farmerfilter.CreatedDate != nil {
			//var sd,ed time.Time
			if Farmerfilter.CreatedDate.From != nil {
				sd := time.Date(Farmerfilter.CreatedDate.From.Year(), Farmerfilter.CreatedDate.From.Month(), Farmerfilter.CreatedDate.From.Day(), 0, 0, 0, 0, Farmerfilter.CreatedDate.From.Location())
				ed := time.Date(Farmerfilter.CreatedDate.From.Year(), Farmerfilter.CreatedDate.From.Month(), Farmerfilter.CreatedDate.From.Day(), 23, 59, 59, 0, Farmerfilter.CreatedDate.From.Location())
				if Farmerfilter.CreatedDate.To != nil {
					ed = time.Date(Farmerfilter.CreatedDate.To.Year(), Farmerfilter.CreatedDate.To.Month(), Farmerfilter.CreatedDate.To.Day(), 23, 59, 59, 0, Farmerfilter.CreatedDate.To.Location())
				}
				query = append(query, bson.M{"createdDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("farmer pagenation query =>", paginationPipeline)
		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "farmerOrg", "_id", "ref.farmerOrg", "ref.farmerOrg")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "soilType", "_id", "ref.soilType", "ref.soilType")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSET, "assert", "_id", "ref.assert", "ref.assert")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &Farmers); err != nil {
		return nil, err
	}
	return Farmers, nil
} //FilterFarmer : ""
func (d *Daos) FilterDuplicateFarmer(ctx *models.Context, farmerfilter *models.DuplicateFarmerFilter, pagination *models.Pagination) ([]models.DuplicateFarmerReport, error) {
	mainPipeline := []bson.M{}
	// mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTFARMER, bson.M{
	// 	"farmerId": "$_id",
	// }, []bson.M{
	// 	{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 		{"$eq": []string{"$status", constants.PROJECTFARMERSTATUSACTIVE}},
	// 		{"$eq": []string{"$farmer", "$$farmerId"}},
	// 	}}}},
	// }, "ref.projects", "ref.projects")...)
	query := d.FarmerReportFilter(ctx, &farmerfilter.FarmerFilter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"farmerOrg": 1}})
	if farmerfilter != nil {
		mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$" + farmerfilter.By, "farmers": bson.M{"$push": "$$ROOT"}}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"count": bson.M{"$size": "$farmers"}}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"count": bson.M{"$gt": 1}}})

	d.Shared.BsonToJSONPrintTag("DuplicateFarmer query =>", mainPipeline)

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("farmer pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	mainPipeline = append(mainPipeline, d.FarmerlookupQueryConstration(ctx)...)

	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.DuplicateFarmerReport
	if err = cursor.All(context.TODO(), &Farmers); err != nil {
		return nil, err
	}
	return Farmers, nil
}
func (d *Daos) FarmerReportFilter(ctx *models.Context, Farmerfilter *models.FarmerFilter) []bson.M {

	query := []bson.M{}
	if Farmerfilter != nil {

		if len(Farmerfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Farmerfilter.ActiveStatus}})
		}
		if len(Farmerfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Farmerfilter.Status}})
		}
		if len(Farmerfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": Farmerfilter.State}})
		}
		if len(Farmerfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": Farmerfilter.District}})
		}
		if len(Farmerfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": Farmerfilter.Block}})
		}
		if len(Farmerfilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": Farmerfilter.GramPanchayat}})
		}
		if len(Farmerfilter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": Farmerfilter.Village}})
		}
		if len(Farmerfilter.FarmerOrg) > 0 {
			query = append(query, bson.M{"farmerOrg": bson.M{"$in": Farmerfilter.FarmerOrg}})
		}
		//Regex
		if Farmerfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Farmerfilter.Regex.Name, Options: "xi"}})
		}
		if Farmerfilter.Regex.MobileNumber != "" {
			query = append(query, bson.M{"mobileNumber": primitive.Regex{Pattern: Farmerfilter.Regex.MobileNumber, Options: "xi"}})
		}
		if Farmerfilter.Regex.SpouseName != "" {
			query = append(query, bson.M{"spouseName": primitive.Regex{Pattern: Farmerfilter.Regex.SpouseName, Options: "xi"}})
		}
		if Farmerfilter.OmitProjectFarmer.Is {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$ne": Farmerfilter.OmitProjectFarmer.Project}})
		}
		//daterange
		if Farmerfilter.CreatedDate != nil {
			//var sd,ed time.Time
			if Farmerfilter.CreatedDate.From != nil {
				sd := time.Date(Farmerfilter.CreatedDate.From.Year(), Farmerfilter.CreatedDate.From.Month(), Farmerfilter.CreatedDate.From.Day(), 0, 0, 0, 0, Farmerfilter.CreatedDate.From.Location())
				ed := time.Date(Farmerfilter.CreatedDate.From.Year(), Farmerfilter.CreatedDate.From.Month(), Farmerfilter.CreatedDate.From.Day(), 23, 59, 59, 0, Farmerfilter.CreatedDate.From.Location())
				if Farmerfilter.CreatedDate.To != nil {
					ed = time.Date(Farmerfilter.CreatedDate.To.Year(), Farmerfilter.CreatedDate.To.Month(), Farmerfilter.CreatedDate.To.Day(), 23, 59, 59, 0, Farmerfilter.CreatedDate.To.Location())
				}
				query = append(query, bson.M{"createdDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
	}
	return query
}
