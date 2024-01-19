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

//SaveQuery :""
func (d *Daos) SaveQuery(ctx *models.Context, Query *models.Query) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONQUERY).InsertOne(ctx.CTX, Query)
	if err != nil {
		return err
	}
	Query.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleQuery : ""
func (d *Daos) GetSingleQuery(ctx *models.Context, code string) (*models.RefQuery, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "assignedTo", "_id", "ref.assignedTo", "ref.assignedTo")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "resolvedBy", "_id", "ref.resolvedBy", "ref.resolvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "createdBy", "_id", "ref.createdByUser", "ref.createdByUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "createdBy", "_id", "ref.createdByFarmer", "ref.createdByFarmer")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONQUERY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Querys []models.RefQuery
	var Query *models.RefQuery
	if err = cursor.All(ctx.CTX, &Querys); err != nil {
		return nil, err
	}
	if len(Querys) > 0 {
		Query = &Querys[0]
	}
	return Query, nil
}

//UpdateQuery : ""
func (d *Daos) UpdateQuery(ctx *models.Context, Query *models.Query) error {
	selector := bson.M{"_id": Query.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Query, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONQUERY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterQuery : ""
func (d *Daos) FilterQuery(ctx *models.Context, Queryfilter *models.QueryFilter, pagination *models.Pagination) ([]models.RefQuery, error) {
	//mainPipeline := []bson.M{}
	mainPipeline, err := d.QueryFilter(ctx, Queryfilter)
	if err != nil {
		return nil, err
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("query pagenation query =>", paginationPipeline)
		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONQUERY).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "assignedTo", "_id", "ref.assignedTo", "ref.assignedTo")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "resolvedBy", "_id", "ref.resolvedBy", "ref.resolvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "createdBy", "_id", "ref.createdByUser", "ref.createdByUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.createdByFarmer", "ref.createdByFarmer")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Query query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONQUERY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Querys []models.RefQuery
	if err = cursor.All(context.TODO(), &Querys); err != nil {
		return nil, err
	}
	return Querys, nil
}

//EnableQuery :""
func (d *Daos) EnableQuery(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.QUERYSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONQUERY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableQuery :""
func (d *Daos) DisableQuery(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.QUERYSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONQUERY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteQuery :""
func (d *Daos) DeleteQuery(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.QUERYSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONQUERY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//AssignuserQuery : ""
func (d *Daos) AssignuserQuery(ctx *models.Context, Query *models.AssignUserToQuery) error {
	selector := bson.M{"_id": Query.QueryId}
	t := time.Now()
	update := bson.M{"$set": bson.M{"status": constants.QUERYSTATUSASSIGN, "assignedTo": Query.UserId, "assignedDate": &t}}

	_, err := ctx.DB.Collection(constants.COLLECTIONQUERY).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//ResolveuserQuery : ""
func (d *Daos) ResolveuserQuery(ctx *models.Context, Query *models.ResolveQuery) error {
	selector := bson.M{"_id": Query.QueryId}
	t := time.Now()
	update := bson.M{"$set": bson.M{"status": constants.QUERYSTATUSRESOLVED, "resolvedBy": Query.UserId, "resolvedDate": &t, "isSolutionSent": true, "solution": Query.Solution}}

	_, err := ctx.DB.Collection(constants.COLLECTIONQUERY).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) QueryFilter(ctx *models.Context, Queryfilter *models.QueryFilter) ([]bson.M, error) {
	var err error
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Queryfilter != nil {

		if len(Queryfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Queryfilter.ActiveStatus}})
		}
		if len(Queryfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Queryfilter.Status}})
		}

		if len(Queryfilter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": Queryfilter.Village}})
		}
		if len(Queryfilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": Queryfilter.GramPanchayat}})
		}
		if len(Queryfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": Queryfilter.Block}})
		}
		if len(Queryfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": Queryfilter.District}})
		}
		if len(Queryfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": Queryfilter.State}})
		}
		if len(Queryfilter.AssignedTo) > 0 {
			query = append(query, bson.M{"assignedTo": bson.M{"$in": Queryfilter.AssignedTo}})
		}
		if len(Queryfilter.ResolvedBy) > 0 {
			query = append(query, bson.M{"resolvedBy": bson.M{"$in": Queryfilter.ResolvedBy}})
		}
		if len(Queryfilter.Farmer) > 0 {
			query = append(query, bson.M{"farmer": bson.M{"$in": Queryfilter.Farmer}})
		}
		if len(Queryfilter.ContentID) > 0 {
			query = append(query, bson.M{"contentId": bson.M{"$in": Queryfilter.ContentID}})
		}
		//Regex
		if Queryfilter.Regex.Query != "" {
			query = append(query, bson.M{"query": primitive.Regex{Pattern: Queryfilter.Regex.Query, Options: "xi"}})
		}
		if Queryfilter.CreatedFrom.StartDate != nil {
			var sd, ed time.Time
			var sdcondition, edcondition string = "gte", "lte"
			sd = time.Date(Queryfilter.CreatedFrom.StartDate.Year(), Queryfilter.CreatedFrom.StartDate.Month(), Queryfilter.CreatedFrom.StartDate.Day(), 0, 0, 0, 0, Queryfilter.CreatedFrom.StartDate.Location())
			ed = time.Date(Queryfilter.CreatedFrom.EndDate.Year(), Queryfilter.CreatedFrom.EndDate.Month(), Queryfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, Queryfilter.CreatedFrom.EndDate.Location())

			if Queryfilter.CreatedFrom.EndDate != nil {
				ed = time.Date(Queryfilter.CreatedFrom.EndDate.Year(), Queryfilter.CreatedFrom.EndDate.Month(), Queryfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, Queryfilter.CreatedFrom.EndDate.Location())
				//edcondition = queryfilter.CreatedTo.Condition
			}
			query = append(query, bson.M{"date": bson.M{"$" + sdcondition: sd, "$" + edcondition: ed}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	return mainPipeline, err

}
