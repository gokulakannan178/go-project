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

//FilterUserReport : ""
func (d *Daos) FilterUserReport(ctx *models.Context, filter *models.UserReportFilter, pagination *models.Pagination) ([]models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTUSER, bson.M{
		"userId": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.PROJECTUSERSTATUSACTIVE}},
			{"$eq": []string{"$user", "$$userId"}},
		}}}},
	}, "ref.projects", "ref.projects")...)
	query := []bson.M{}
	if filter != nil {

		if !filter.Village.ID.IsZero() {
			query = append(query, bson.M{"villageCode": bson.M{"$" + filter.Village.Condition: filter.Village.ID}})
		}
		if !filter.State.ID.IsZero() {
			query = append(query, bson.M{"stateCode": bson.M{"$" + filter.State.Condition: filter.State.ID}})
		}
		if !filter.District.ID.IsZero() {
			query = append(query, bson.M{"districtCode": bson.M{"$" + filter.District.Condition: filter.District.ID}})
		}
		if !filter.GramPanchayat.ID.IsZero() {
			query = append(query, bson.M{"grampanchayatCode": bson.M{"$" + filter.GramPanchayat.Condition: filter.GramPanchayat.ID}})
		}
		if !filter.Block.ID.IsZero() {
			query = append(query, bson.M{"blockCode": bson.M{"$" + filter.Block.Condition: filter.Block.ID}})
		}
		if !filter.KnowledgeDomain.ID.IsZero() {
			query = append(query, bson.M{"knowledgeDomains": bson.M{"$" + filter.KnowledgeDomain.Condition: filter.KnowledgeDomain.ID}})
		}
		if !filter.SubDomain.ID.IsZero() {
			query = append(query, bson.M{"subDomains": bson.M{"$" + filter.SubDomain.Condition: filter.SubDomain.ID}})
		}
		if !filter.Language.ID.IsZero() {
			query = append(query, bson.M{"languageExpertise": bson.M{"$" + filter.Language.Condition: filter.Language.ID}})
		}
		if filter.Role.ID != "" {
			query = append(query, bson.M{"type": bson.M{"$" + filter.Role.Condition: filter.Role.ID}})
		}
		if filter.UserName.ID != "" {
			query = append(query, bson.M{"userName": bson.M{"$" + filter.UserName.Condition: filter.UserName.ID}})
		}
		if filter.Experience.ID != 0 {
			query = append(query, bson.M{"experience": bson.M{"$" + filter.Experience.Condition: filter.Experience.ID}})
		}
		if filter.Gender.ID != "" {
			query = append(query, bson.M{"gender": bson.M{"$" + filter.Gender.Condition: filter.Gender.ID}})
		}
		if filter.AccessLevel.ID != "" {
			query = append(query, bson.M{"accessPrivilege.accessLevel": bson.M{"$" + filter.AccessLevel.Condition: filter.AccessLevel.ID}})
		}
		if !filter.Organisation.ID.IsZero() {
			query = append(query, bson.M{"userOrg": bson.M{"$" + filter.Organisation.Condition: filter.Organisation.ID}})
		}
		if !filter.Project.ID.IsZero() {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$" + filter.Project.Condition: filter.Project.ID}})
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
		d.Shared.BsonToJSONPrintTag("user pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "blockCode", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "grampanchayatCode", "_id", "ref.grampanchayat", "ref.grampanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "districtCode", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "villageCode", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "userOrg", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISTRICT, "accessPrivilege.districts", "_id", "ref.accessDistricts", "ref.accessDistricts")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "accessPrivilege.states", "_id", "ref.accessStates", "ref.accessStates")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "accessPrivilege.villages", "_id", "ref.accessVillages", "ref.accessVillages")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONBLOCK, "accessPrivilege.blocks", "_id", "ref.accessBlocks", "ref.accessBlocks")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONGRAMPANCHAYAT, "accessPrivilege.grampanchayats", "_id", "ref.accessGrampanchayats", "ref.accessGrampanchayats")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomains", "_id", "ref.knowledgeDomains", "ref.knowledgeDomains")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSUBDOMAIN, "subDomains", "_id", "ref.subDomains", "ref.subDomains")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Query query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Users []models.RefUser
	if err = cursor.All(context.TODO(), &Users); err != nil {
		return nil, err
	}
	return Users, nil
}

//FilterDuplicateUserReport : ""
func (d *Daos) FilterDuplicateUserReport(ctx *models.Context, userfilter *models.DuplicateUserFilter, pagination *models.Pagination) ([]models.DuplicateUserReport, error) {

	paginationPipeline := []bson.M{}
	mainPipeline := d.UserReportFilter(ctx, &userfilter.UserFilter)
	//Adding $match from filter

	mainPipeline = append(mainPipeline, d.UserlookupQueryConstration(ctx)...)
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"userOrg": 1}})
	if userfilter != nil {
		mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$" + userfilter.By, "users": bson.M{"$push": "$$ROOT"}}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"count": bson.M{"$size": "$users"}}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"count": bson.M{"$gt": 1}}})
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("user pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.DuplicateUserReport
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
func (d *Daos) UserReportFilter(ctx *models.Context, userfilter *models.UserFilter) []bson.M {
	mainPipeline := []bson.M{}
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROJECTUSER, "_id", "user", "ref.projects", "ref.projects")...)
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTUSER, bson.M{
		"userId": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.PROJECTUSERSTATUSACTIVE}},
			{"$eq": []string{"$user", "$$userId"}},
		}}}},
	}, "ref.projects", "ref.projects")...)
	query := []bson.M{}
	if userfilter != nil {
		if len(userfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": userfilter.Status}})
		}
		if len(userfilter.Manager) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": userfilter.Manager}})
		}
		if len(userfilter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": userfilter.Type}})
		}
		if len(userfilter.OmitID) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$nin": userfilter.OmitID}})
		}
		if len(userfilter.OrganisationID) > 0 {
			query = append(query, bson.M{"userOrg": bson.M{"$in": userfilter.OrganisationID}})
		}
		if len(userfilter.AccessLevel) > 0 {
			query = append(query, bson.M{"accessPrivilege.accessLevel": bson.M{"$in": userfilter.AccessLevel}})
		}
		if len(userfilter.States) > 0 {
			query = append(query, bson.M{"stateCode": bson.M{"$in": userfilter.States}})
		}
		if len(userfilter.Districts) > 0 {
			query = append(query, bson.M{"districtCode": bson.M{"$in": userfilter.Districts}})
		}
		if len(userfilter.Blocks) > 0 {
			query = append(query, bson.M{"blockCode": bson.M{"$in": userfilter.Blocks}})
		}
		if len(userfilter.Villages) > 0 {
			query = append(query, bson.M{"villageCode": bson.M{"$in": userfilter.Villages}})
		}
		if len(userfilter.Grampanchayats) > 0 {
			query = append(query, bson.M{"grampanchayatCode": bson.M{"$in": userfilter.Grampanchayats}})
		}
		//Regex
		if userfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: userfilter.Regex.Name, Options: "xi"}})
		}
		if userfilter.Regex.Contact != "" {
			query = append(query, bson.M{"mobileNumber": primitive.Regex{Pattern: userfilter.Regex.Contact, Options: "xi"}})
		}
		if userfilter.Regex.UserName != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: userfilter.Regex.UserName, Options: "xi"}})
		}
		if userfilter.Regex.FirstName != "" {
			query = append(query, bson.M{"firstName": primitive.Regex{Pattern: userfilter.Regex.FirstName, Options: "xi"}})
		}
		if userfilter.Regex.Lastname != "" {
			query = append(query, bson.M{"lastname": primitive.Regex{Pattern: userfilter.Regex.Lastname, Options: "xi"}})
		}

		if userfilter.OmitProjectUser.Is {
			query = append(query, bson.M{"ref.projects.project": bson.M{"$ne": userfilter.OmitProjectUser.Project}})
		}
		if userfilter.CreatedFrom.StartDate != nil {
			var sd, ed time.Time
			var sdcondition, edcondition string = "gte", "lte"
			sd = time.Date(userfilter.CreatedFrom.StartDate.Year(), userfilter.CreatedFrom.StartDate.Month(), userfilter.CreatedFrom.StartDate.Day(), 0, 0, 0, 0, userfilter.CreatedFrom.StartDate.Location())
			ed = time.Date(userfilter.CreatedFrom.EndDate.Year(), userfilter.CreatedFrom.EndDate.Month(), userfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, userfilter.CreatedFrom.EndDate.Location())

			if userfilter.CreatedFrom.EndDate != nil {
				ed = time.Date(userfilter.CreatedFrom.EndDate.Year(), userfilter.CreatedFrom.EndDate.Month(), userfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, userfilter.CreatedFrom.EndDate.Location())
				//edcondition = userfilter.CreatedTo.Condition
			}
			query = append(query, bson.M{"createdDate": bson.M{"$" + sdcondition: sd, "$" + edcondition: ed}})
		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	return mainPipeline

}
