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

//SaveFarmer :""
func (d *Daos) SaveFarmer(ctx *models.Context, Farmer *models.Farmer) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMER).InsertOne(ctx.CTX, Farmer)
	if err != nil {
		return err
	}
	Farmer.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveManyFarmers(ctx *models.Context, farmers []interface{}) ([]interface{}, error) {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMER).InsertMany(ctx.CTX, farmers)
	if err != nil {
		return nil, err
	}
	return res.InsertedIDs, nil
}

//GetSingleFarmer : ""
func (d *Daos) GetSingleFarmer(ctx *models.Context, code string) (*models.RefFarmer, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "farmerOrg", "_id", "ref.farmerOrg", "ref.farmerOrg")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "soilType", "_id", "ref.soilType", "ref.soilType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSET, "assert", "_id", "ref.assert", "ref.assert")...)
	mainPipeline = append(mainPipeline, d.FarmerlookupQueryConstration(ctx)...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.RefFarmer
	var Farmer *models.RefFarmer
	if err = cursor.All(ctx.CTX, &Farmers); err != nil {
		return nil, err
	}
	if len(Farmers) > 0 {
		Farmer = &Farmers[0]
	}
	return Farmer, nil
}

//GetSingleFarmer : ""
func (d *Daos) GetSingleFarmerWithFarmerId(ctx *models.Context, code string) ([]models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"farmerID": code}})

	//Lookups
	//mainPipeline = append(mainPipeline, d.FarmerlookupQueryConstration(ctx)...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	//Aggregation
	var projects []models.RefFarmer
	if err = cursor.All(context.TODO(), &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

//UpdateFarmer : ""
func (d *Daos) UpdateFarmer(ctx *models.Context, Farmer *models.Farmer) error {
	selector := bson.M{"_id": Farmer.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Farmer, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//UpdateCastFarmer : ""
func (d *Daos) UpdateCastFarmer(ctx *models.Context, Farmer string, cast string) error {
	selector := bson.M{"farmerID": Farmer}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"cast": cast}}
	d.Shared.BsonToJSONPrintTag("farmer cast update selector=>", selector)
	d.Shared.BsonToJSONPrintTag("farmer cast update query=>", updateInterface)

	_, err := ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFarmer : ""
func (d *Daos) FilterFarmer(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) ([]models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	productconfig, err := d.GetactiveProductConfig(ctx, true)
	if err != nil {
		return nil, err
	}
	if !productconfig.IsSingleProject {
		Farmerfilter.RemoveLookup.Projects = true
	}
	mainPipeline, err = d.FilterFarmerQuery(ctx, Farmerfilter)
	if err != nil {
		return nil, err
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		// paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		paginationPipeline = append(paginationPipeline, bson.M{
			"$count": "count",
		})
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("farmer pagenation query =>", paginationPipeline)
		// c := d.Redis.GetValue("FARMERCOUNT")
		// fmt.Println("Farmer Count=====>", c)
		// // var ok bool

		// countStr, _ := c.(string)
		// fmt.Println("countStr====>", countStr)
		// //	fmt.Println("ok====>", ok)
		// if countStr == "" {
		// 	fmt.Println("CHECKING COUNT=======================================")
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
		err = d.Redis.SetValue("FARMERCOUNT", totalCount, 900)
		if err != nil {
			log.Println("redis error -", err.Error())
		}
		pagination.Count = int(totalCount)
		// } else {
		// 	//pagination.Count
		// 	i, err := strconv.Atoi(countStr)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	pagination.Count = i
		//	}

		d.Shared.PaginationData(pagination)
	}
	//	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	//Lookups
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)
	if !Farmerfilter.RemoveLookup.GramPanchayat {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)

	}
	if !Farmerfilter.RemoveLookup.Village {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	}
	if !Farmerfilter.RemoveLookup.Block {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	}
	if !Farmerfilter.RemoveLookup.District {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	}
	if !Farmerfilter.RemoveLookup.State {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	}
	if !Farmerfilter.RemoveLookup.FarmerOrg {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "farmerOrg", "_id", "ref.farmerOrg", "ref.farmerOrg")...)
	}
	if !Farmerfilter.RemoveLookup.SoilType {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "soilType", "_id", "ref.soilType", "ref.soilType")...)
	}
	if !Farmerfilter.RemoveLookup.Assert {
		mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSET, "assert", "_id", "ref.assert", "ref.assert")...)
	}

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

//FilterFarmerBasic : ""
func (d *Daos) FilterFarmerBasic(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) ([]models.RefBasicFarmer, error) {
	mainPipeline, err := d.FilterFarmerQuery(ctx, Farmerfilter)
	if err != nil {
		return nil, err
	}
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		// paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		paginationPipeline = append(paginationPipeline, bson.M{
			"$count": "count",
		})
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
	if !Farmerfilter.RemoveLookup.GramPanchayat {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)

	}
	if !Farmerfilter.RemoveLookup.Village {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	}
	if !Farmerfilter.RemoveLookup.Block {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	}
	if !Farmerfilter.RemoveLookup.District {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	}
	if !Farmerfilter.RemoveLookup.State {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	}
	if !Farmerfilter.RemoveLookup.FarmerOrg {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "farmerOrg", "_id", "ref.farmerOrg", "ref.farmerOrg")...)
	}
	if !Farmerfilter.RemoveLookup.SoilType {
		mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "soilType", "_id", "ref.soilType", "ref.soilType")...)
	}
	if !Farmerfilter.RemoveLookup.Assert {
		mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSET, "assert", "_id", "ref.assert", "ref.assert")...)
	}

	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"_id": 1, "name": 1}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.RefBasicFarmer
	if err = cursor.All(context.TODO(), &Farmers); err != nil {
		return nil, err
	}
	return Farmers, nil
}

func (d *Daos) FarmerlookupQueryConstration(ctx *models.Context) []bson.M {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "farmerOrg", "_id", "ref.farmerOrg", "ref.farmerOrg")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "soilType", "_id", "ref.soilType", "ref.soilType")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSET, "assert", "_id", "ref.assert", "ref.assert")...)

	return mainPipeline

}

//EnableFarmer :""
func (d *Daos) EnableFarmer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmer :""
func (d *Daos) DisableFarmer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmer :""
func (d *Daos) DeleteFarmer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) FarmerUniquenessCheckRegistration(ctx *models.Context, orgID string, param string, value string) (*models.FarmerUniquinessChk, error) {
	id, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return nil, err
	}
	query := bson.M{
		"farmerOrg": id,
		param:       value,
	}
	fmt.Println("query====>", query)

	result := new(models.FarmerUniquinessChk)
	var farmer *models.Farmer
	if err = ctx.DB.Collection(constants.COLLECTIONFARMER).FindOne(ctx.CTX, query).Decode(&farmer); err != nil {
		return nil, err
	}
	fmt.Println("farmerp===>", farmer)
	if farmer == nil {
		result.Success = true
		return result, nil
	}
	if farmer.Status == constants.FARMERSTATUSACTIVE {
		result.Success = false
		result.Message = "farmer alreasy exist"
		return result, nil
	}
	if farmer.Status == constants.FARMERSTATUSDISABLED {
		result.Success = false
		result.Message = "farmer disabled - please contact Administator"
		return result, nil
	}
	if farmer.Status == constants.FARMERSTATUSINIT {
		result.Success = false
		result.Message = "Awaiting Activation - please contact administator"
		return result, nil
	}
	return nil, errors.New("No options Availables")
}
func (d *Daos) FilterFarmerQuery(ctx *models.Context, Farmerfilter *models.FarmerFilter) ([]bson.M, error) {
	mainPipeline := []bson.M{}
	productconfig, err := d.GetactiveProductConfig(ctx, true)
	if err != nil {
		return nil, err
	}
	if !productconfig.IsSingleProject {
		if !Farmerfilter.RemoveLookup.Projects {
			mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTFARMER, bson.M{
				"farmerId": "$_id",
			}, []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$status", constants.PROJECTUSERSTATUSACTIVE}},
					{"$eq": []string{"$farmer", "$$farmerId"}},
				}}}},
				{"$lookup": bson.M{
					"from":         constants.COLLECTIONPROJECT,
					"as":           "ref.project",
					"localField":   "project",
					"foreignField": "_id",
				}},
				{"$addFields": bson.M{"ref.project": bson.M{"$arrayElemAt": []interface{}{"$ref.project", 0}}}},
			}, "ref.projects", "ref.projects")...)
		}
	}

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
		if len(Farmerfilter.NonFarmerOrg) > 0 {
			query = append(query, bson.M{"farmerOrg": bson.M{"$nin": Farmerfilter.NonFarmerOrg}})
		}
		if len(Farmerfilter.NonProject) > 0 {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$nin": Farmerfilter.NonProject}})
			//	query = append(query, bson.M{"ref.project": bson.M{"$elemMatch": bson.M{"project": bson.M{"$nin": Farmerfilter.NonProject}}}})
		}
		//	{ "results": { $elemMatch: { product: { $ne: "xyz" } } } }
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
		if !productconfig.IsSingleProject {
			if len(Farmerfilter.Project) > 0 {
				query = append(query, bson.M{"ref.projects.project": bson.M{"$in": Farmerfilter.Project}})
			}
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
	return mainPipeline, nil
}
func (d *Daos) GetContentDisseminationFarmer(ctx *models.Context, Farmerfilter *models.FarmerFilter) ([]models.DissiminateFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{"_id": 1, "name": 1, "mobileNumber": 1,
			"farmerOrg":     1,
			"ref.projects":  1,
			"farmerID":      1,
			"state":         1,
			"district":      1,
			"block":         1,
			"gramPanchayat": 1,
			"village":       1,
			"status":        1,
		},
	})
	queryPipeline, err := d.FilterFarmerQuery(ctx, Farmerfilter)
	if err != nil {
		return nil, err
	}
	mainPipeline = append(mainPipeline, queryPipeline...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAPPTOKEN, "_id", "userid", "appRegistrationToken", "appRegistrationToken")...)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"appRegistrationToken": "$appRegistrationToken.registrationtoken"}})

	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{"_id": 1, "name": 1, "mobileNumber": 1, "email": 1, "appRegistrationToken": 1, "farmerID": 1},
	})
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})
	d.Shared.BsonToJSONPrintTag("GetContentDisseminationFarmer", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.DissiminateFarmer
	if err = cursor.All(context.TODO(), &farmers); err != nil {
		return nil, err
	}
	return farmers, nil

}
func (d *Daos) GetSingleFarmerWithMobileno(ctx *models.Context, mobileNumber string) (*models.RefFarmer, error) {
	mainPipeline := []bson.M{}

	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNumber": mobileNumber}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("farmermobile query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &farmers); err != nil {
		return nil, err
	}
	if len(farmers) > 0 {
		return &farmers[0], nil
	}

	return nil, errors.New("farmer not found")
}
func (d *Daos) GetSingleFarmerWithMobilenoAndOrg(ctx *models.Context, org string, mobileNumber string) (*models.RefFarmer, error) {

	mainPipeline := []bson.M{}
	orgid, err := primitive.ObjectIDFromHex(org)
	if err != nil {
		return nil, err
	}
	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNumber": mobileNumber, "farmerOrg": orgid}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("farmermobile query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &farmers); err != nil {
		return nil, err
	}
	if len(farmers) > 0 {
		return &farmers[0], nil
	}

	return nil, errors.New("farmer not found")
}

func (d *Daos) FarmerNearBy(ctx *models.Context, farmernb *models.NearBy, pagination *models.Pagination) ([]models.RefFarmer, error) {
	coordinater := []float64{farmernb.Longitude, farmernb.Latitude}

	mainPipeline := []bson.M{}
	query := []bson.M{{"$geoNear": bson.M{"near": bson.M{"type": "Point", "coordinates": coordinater},
		"maxDistance":   farmernb.KM * 1000,
		"distanceField": "dist.calculated",
		"spherical":     true,
		"includeLocs":   "dist.location",
	},
	}}
	fmt.Println(query)
	mainPipeline = append(mainPipeline, query...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count

		Count, err := GetFarmerCountForAggregation(ctx, query, constants.COLLECTIONFARMER)
		if err != nil {
			log.Println(err)
		}
		pagination.Count = Count
		d.Shared.PaginationData(pagination)
	}
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &farmers); err != nil {
		return nil, err
	}
	return farmers, err
}

func GetFarmerCountForAggregation(ctx *models.Context, query []primitive.M, Collection string) (int, error) {
	mainPipeline2 := []bson.M{}

	mainPipeline2 = append(mainPipeline2, query...)

	//Getting Total count
	group := []bson.M{{"$group": bson.M{"_id": nil, "mycount": bson.M{"$sum": 1}}}}
	mainPipeline2 = append(mainPipeline2, group...)
	type countvaule struct {
		Mycount int ` bson:"mycount"`
	}
	var totalCountV = make([]countvaule, 0)
	totalCount, err := ctx.DB.Collection(Collection).Aggregate(ctx.CTX, mainPipeline2, nil)
	if err != nil {
		log.Println("Error in geting pagination count" + err.Error())
	}
	log.Println(totalCount, totalCountV)
	if err = totalCount.All(ctx.CTX, &totalCountV); err != nil {
		return 0, err
	}
	fmt.Println("count", totalCount)
	if len(totalCountV) > 0 {
		return int(totalCountV[0].Mycount), nil
	}

	return 0, nil

}
func (d *Daos) AddProjectfarmer(ctx *models.Context, farmer *models.FarmerFilter) ([]models.AddProjectFarmer, error) {
	mainPipeline := []bson.M{}
	mainPipeline, err := d.FilterFarmerQuery(ctx, farmer)
	if err != nil {
		return nil, err
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "farmerOrg", "_id", "organisation.name", "organisation")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "organisation",
			"foreignField": "_id",
			"from":         "organisation",
			"localField":   "farmerOrg"}},
		bson.M{"$addFields": bson.M{"organisation": bson.M{"$arrayElemAt": []interface{}{"$organisation.name", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{"_id": 1, "name": 1, "organisation": 1},
	})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Add Project Farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.AddProjectFarmer
	if err = cursor.All(context.TODO(), &Farmers); err != nil {
		return nil, err
	}
	return Farmers, nil
}

//UpdateFarmerProfileImage : ""
func (d *Daos) UpdateFarmerProfileImage(ctx *models.Context, Farmer *models.Farmer) error {
	selector := bson.M{"_id": Farmer.ID}

	update := bson.M{"$set": bson.M{"profileImg": Farmer.ProfileImg}}
	//updateInterface := bson.M{"$set": Farmer, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) FilterFarmerWithLocation(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) ([]models.FarmerLocation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
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
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		// paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		paginationPipeline = append(paginationPipeline, bson.M{
			"$count": "count",
		})
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
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.FarmerLocation
	if err = cursor.All(context.TODO(), &Farmers); err != nil {
		return nil, err
	}
	return Farmers, nil
}
func (d *Daos) GetSingleFarmerWithName(ctx *models.Context, org string, name string) (*models.RefFarmer, error) {

	mainPipeline := []bson.M{}
	orgid, err := primitive.ObjectIDFromHex(org)
	if err != nil {
		return nil, err
	}
	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": name, "farmerOrg": orgid}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("farmermobile query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &farmers); err != nil {
		return nil, err
	}
	if len(farmers) > 0 {
		return &farmers[0], nil
	}

	return nil, errors.New("farmer not found")
}
