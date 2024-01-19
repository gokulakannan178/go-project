package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveProject :""
func (d *Daos) SaveProject(ctx *models.Context, project *models.Project) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).InsertOne(ctx.CTX, project)
	return err
}

//SaveProjectMembers :""
func (d *Daos) SaveProjectMembers(ctx *models.Context, projectmembers []models.ProjectMember) error {
	var data []interface{}
	//array
	for _, v := range projectmembers {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTMEMBER).InsertMany(ctx.CTX, data)
	return err
}

//SaveProjectTeamMember :""
func (d *Daos) SaveProjectTeamMember(ctx *models.Context, project *models.ProjectMember) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTMEMBER).InsertOne(ctx.CTX, project)
	return err
}

//DisableProjectTeamMember :""
func (d *Daos) DisableProjectTeamMember(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTTEAMMEMBERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTMEMBER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProject : ""
func (d *Daos) GetSingleProject(ctx *models.Context, UniqueID string) (*models.Project, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.Project
	var user *models.Project
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//UpdateProject : ""
func (d *Daos) UpdateProject(ctx *models.Context, user *models.Project) error {
	selector := bson.M{"uniqueId": user.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": user} //update model(user)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProject :""
func (d *Daos) EnableProject(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProject :""
func (d *Daos) DisableProject(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProject :""
func (d *Daos) DeleteProject(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// ProjectFilter : ""
func (d *Daos) ProjectFilter(ctx *models.Context, projectfilter *models.ProjectFilter, pagination *models.Pagination) ([]models.Project, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if projectfilter != nil {
		if len(projectfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": projectfilter.Status}})
		}
		// if len(userfilter.Manager) > 0 {
		// 	query = append(query, bson.M{"managerId": bson.M{"$in": userfilter.Manager}})
		// }
		// if len(userfilter.Type) > 0 {
		// 	query = append(query, bson.M{"type": bson.M{"$in": userfilter.Type}})
		// }
		// if len(userfilter.OmitID) > 0 {
		// 	query = append(query, bson.M{"userName": bson.M{"$nin": userfilter.OmitID}})
		// }
		// if len(userfilter.OrganisationID) > 0 {
		// 	query = append(query, bson.M{"organisationId": bson.M{"$in": userfilter.OrganisationID}})
		// }

		//Regex
		// if userfilter.Regex.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: userfilter.Regex.Name, Options: "xi"}})
		// }
		// if userfilter.Regex.Contact != "" {
		// 	query = append(query, bson.M{"mobile": primitive.Regex{Pattern: userfilter.Regex.Contact, Options: "xi"}})
		// }
		// if userfilter.Regex.UserName != "" {
		// 	query = append(query, bson.M{"userName": primitive.Regex{Pattern: userfilter.Regex.UserName, Options: "xi"}})
		// }
	}
	// //Adding $match from filter
	// if len(query) > 0 {
	// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).CountDocuments(ctx.CTX, func() bson.M {
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
	// if projectfilter.GetRecentLocation {
	// 	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// 		"from": constants.COLLECTIONUSERLOCATION,
	// 		"as":   "ref.lastLocation",
	// 		"let":  bson.M{"userName": "$userName"},
	// 		"pipeline": []bson.M{
	// 			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				bson.M{"$eq": []string{"$userName", "$$userName"}},
	// 			}}}},
	// 			bson.M{"$sort": bson.M{"time": -1}},
	// 			bson.M{"$limit": 1},
	// 		},
	// 	}})
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.lastLocation": bson.M{"$arrayElemAt": []interface{}{"$ref.lastLocation", 0}}}})

	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	//Aggregation
	//d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var projectFilter []models.Project
	if err := cursor.All(context.TODO(), &projectFilter); err != nil {
		return nil, err
	}
	return projectFilter, nil
}
