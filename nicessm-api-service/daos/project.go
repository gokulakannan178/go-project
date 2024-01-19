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

//SaveProject :""
func (d *Daos) SaveProject(ctx *models.Context, project *models.Project) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).InsertOne(ctx.CTX, project)
	if err != nil {
		return err
	}
	project.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateProject : ""
func (d *Daos) UpdateProject(ctx *models.Context, project *models.Project) error {

	selector := bson.M{"_id": project.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": project}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProject :""
func (d *Daos) EnableProject(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSACTIVE, "activeStatus": constants.PROJECTSTATUSTRUE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProject :""
func (d *Daos) DisableProject(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSDISABLED, "activeStatus": constants.PROJECTSTATUSFALSE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProject :""
func (d *Daos) DeleteProject(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProject : ""
func (d *Daos) GetSingleProject(ctx *models.Context, UniqueID string) (*models.RefProject, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROJECTSTATE,
		"as":   "ref.states",
		"let":  bson.M{"projectId": "$_id"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$project", "$$projectId"}},
			}}}},
			{"$lookup": bson.M{
				"from": constants.COLLECTIONSTATE,
				"as":   "ref.state",
				"let":  bson.M{"stateId": "$state"},
				"pipeline": []bson.M{
					{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						{"$eq": []string{"$status", "Active"}},
						{"$eq": []string{"$_id", "$$stateId"}},
					}}}},
				},
			}},
			{
				"$addFields": bson.M{"ref.state": bson.M{"$arrayElemAt": []interface{}{"$ref.state", 0}}},
			},
			{"$group": bson.M{"_id": nil, "stateIds": bson.M{"$push": "$state"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.states": bson.M{"$arrayElemAt": []interface{}{"$ref.states", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.stateIDs": "$ref.states.stateIds"}})

	//partner ids

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROJECTPARTNER,
		"as":   "ref.projectPartner",
		"let":  bson.M{"projectId": "$_id"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$project", "$$projectId"}},
			}}}},

			{"$group": bson.M{"_id": nil, "projectPartnerIDs": bson.M{"$push": "$partner"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.projectPartner": bson.M{"$arrayElemAt": []interface{}{"$ref.projectPartner", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.projectPartnerIDs": "$ref.projectPartner.projectPartnerIDs"}})

	//KD ids

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN,
		"as":   "ref.projectkd",
		"let":  bson.M{"projectId": "$_id"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$project", "$$projectId"}},
			}}}},

			{"$group": bson.M{"_id": nil, "knowledgeDomaiIDs": bson.M{"$push": "$knowledgeDomain"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.projectkd": bson.M{"$arrayElemAt": []interface{}{"$ref.projectkd", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.knowledgeDomaiIDs": "$ref.projectkd.knowledgeDomaiIDs"}})

	// lookup
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROJECTSTATE, "_id", "project", "ref.states", "ref.states")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "ref.states.state", "_id", "ref.states", "ref.states")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "stateId", "_id", "ref.states", "ref.states")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN, "knowledgeDomainId", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROJECTPARTNER, "partnerId", "_id", "ref.partner", "ref.partner")...)
	d.Shared.BsonToJSONPrintTag("project aggurecation query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Projects []models.RefProject
	var Project *models.RefProject
	if err = cursor.All(ctx.CTX, &Projects); err != nil {
		return nil, err
	}
	if len(Projects) > 0 {
		Project = &Projects[0]
	}
	return Project, nil
}

//FilterProject : ""
func (d *Daos) FilterProject(ctx *models.Context, filter *models.ProjectFilter, pagination *models.Pagination) ([]models.RefProject, error) {
	paginationPipeline := []bson.M{}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONPROJECTPARTNER, bson.M{
		"projectId": "$_id",
	}, []bson.M{
		{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []string{"$status", constants.PROJECTPARTNERSTATUSACTIVE}},
			{"$eq": []string{"$partner", "$$projectId"}},
		}}}},
	}, "ref.partners", "ref.partners")...)
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.OmitID) > 0 {
			query = append(query, bson.M{"_id": bson.M{"$nin": filter.OmitID}})
		}
		if len(filter.NationalLevel) > 0 {
			query = append(query, bson.M{"nationalLevel": bson.M{"$in": filter.NationalLevel}})
		}
		if len(filter.Organisation) > 0 {
			query = append(query, bson.M{"organisation": bson.M{"$in": filter.Organisation}})
		}

		if filter.BudgetRange != nil {
			if filter.BudgetRange.From != 0 {
				query = append(query, bson.M{"budget": bson.M{"$gte": filter.BudgetRange.From}})

				if filter.BudgetRange.To != 0 {
					query = append(query, bson.M{"budget": bson.M{"$gte": filter.BudgetRange.From, "$lte": filter.BudgetRange.To}})
				}
			}
		}

		if filter.StartDateRange != nil {
			//var sd,ed time.Time
			if filter.StartDateRange.From != nil {
				sd := time.Date(filter.StartDateRange.From.Year(), filter.StartDateRange.From.Month(), filter.StartDateRange.From.Day(), 0, 0, 0, 0, filter.StartDateRange.From.Location())
				ed := time.Date(filter.StartDateRange.From.Year(), filter.StartDateRange.From.Month(), filter.StartDateRange.From.Day(), 23, 59, 59, 0, filter.StartDateRange.From.Location())
				if filter.StartDateRange.To != nil {
					ed = time.Date(filter.StartDateRange.To.Year(), filter.StartDateRange.To.Month(), filter.StartDateRange.To.Day(), 23, 59, 59, 0, filter.StartDateRange.To.Location())
				}
				query = append(query, bson.M{"startDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.EndDateRange != nil {
			//var sd,ed time.Time
			if filter.EndDateRange.From != nil {
				sd := time.Date(filter.EndDateRange.From.Year(), filter.EndDateRange.From.Month(), filter.EndDateRange.From.Day(), 0, 0, 0, 0, filter.EndDateRange.From.Location())
				ed := time.Date(filter.EndDateRange.From.Year(), filter.EndDateRange.From.Month(), filter.EndDateRange.From.Day(), 23, 59, 59, 0, filter.EndDateRange.From.Location())
				if filter.EndDateRange.To != nil {
					ed = time.Date(filter.EndDateRange.To.Year(), filter.EndDateRange.To.Month(), filter.EndDateRange.To.Day(), 23, 59, 59, 0, filter.EndDateRange.To.Location())
				}
				query = append(query, bson.M{"endDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.Mail != "" {
			query = append(query, bson.M{"mail": primitive.Regex{Pattern: filter.Regex.Mail, Options: "xi"}})
		}
		if filter.OmitProjectPartner.Is {
			// query = append(query, bson.M{"ref.projects.status": constants.PROJECTUSERSTATUSACTIVE})
			query = append(query, bson.M{"ref.partners.project": bson.M{"$ne": filter.OmitProjectPartner.Project}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("project pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination")
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
	// lookup

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "stateId", "_id", "ref.states", "ref.states")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN, "knowledgeDomainId", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONPROJECTPARTNER, "partnerId", "_id", "ref.partner", "ref.partner")...)
	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// 	"from": constants.COLLECTIONPROJECTUSER,
	// 	"as":   "projectuser",
	// 	"let":  bson.M{"user": "$id"},
	// 	"pipeline": []bson.M{
	// 		bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 			bson.M{"$eq": []string{"$status", "Active"}},
	// 		}}}},
	// 		bson.M{"$addFields": bson.M{"projectuser": bson.M{"$arrayElemAt": []interface{}{"$projectuser", 0}}}},
	// 	},
	// }})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Projects []models.RefProject
	if err = cursor.All(context.TODO(), &Projects); err != nil {
		return nil, err
	}
	return Projects, nil
}

//GetSingleProjectWithName : ""
func (d *Daos) GetSingleProjectWithName(ctx *models.Context, Name string, OrganisatinID primitive.ObjectID) ([]models.RefProject, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
	}
	query = append(query, bson.M{"organisation": OrganisatinID})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var projects []models.RefProject
	if err = cursor.All(context.TODO(), &projects); err != nil {
		return nil, err
	}
	return projects, nil
}
func (d *Daos) GetSingleProjectWithUniqueID(ctx *models.Context, UniqueID string) (*models.RefProject, error) {
	mainPipeline := []bson.M{}

	//Adding $match from filter

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var project []models.RefProject
	if err = cursor.All(context.TODO(), &project); err != nil {
		return nil, err
	}
	if len(project) > 0 {
		return &project[0], nil
	}

	return nil, errors.New("project not found")
}
func (d *Daos) ProjectUniquenessCheck(ctx *models.Context, org string, projectname string) (*models.Project, error) {
	orgid, err := primitive.ObjectIDFromHex(org)
	if err != nil {
		return nil, err

	}
	fmt.Println("org---->", org)
	fmt.Println("orgid---->", orgid)
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"organisation": orgid})
	query = append(query, bson.M{"name": primitive.Regex{Pattern: projectname, Options: "i"}})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	d.Shared.BsonToJSONPrintTag("project query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var projects []models.Project
	//var user *models.Project
	if err = cursor.All(ctx.CTX, &projects); err != nil {
		return nil, err
	}
	if len(projects) > 0 {
		return &projects[0], nil
	}
	return nil, errors.New("project not found")

}

func (d *Daos) GetSingleprojectWithName(ctx *models.Context, Name string, organisationID primitive.ObjectID, isRegex bool) (*models.RefProject, error) {
	mainPipeline := []bson.M{}

	query := []bson.M{}
	if Name != "" {
		if isRegex {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Name, Options: "xi"}})
		} else {
			query = append(query, bson.M{"name": Name})

		}
	}
	query = append(query, bson.M{"organisation": organisationID})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var projects []models.RefProject
	if err = cursor.All(context.TODO(), &projects); err != nil {
		return nil, err
	}
	if len(projects) > 0 {
		return &projects[0], nil
	}
	return nil, errors.New("projects not available")
}

func (d *Daos) LegacyUpdateProjectUsers(ctx *models.Context, projectID primitive.ObjectID, farmerIDs []interface{}) error {
	selector := bson.M{
		"_id": projectID,
	}
	update := bson.M{
		"$push": bson.M{"farmers": bson.M{"$each": farmerIDs}},
	}
	// d.Shared.BsonToJSONPrintTag("LegacyUpdateProjectUsers", update)

	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}
