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

//SaveFeedBack :""
func (d *Daos) SaveFeedBack(ctx *models.Context, FeedBack *models.FeedBack) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFEEDBACK).InsertOne(ctx.CTX, FeedBack)
	if err != nil {
		return err
	}
	FeedBack.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleFeedBack : ""
func (d *Daos) GetSingleFeedBack(ctx *models.Context, code string) (*models.RefFeedBack, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "createdBy", "_id", "ref.createdBy", "ref.createdBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONQUERY, "query", "_id", "ref.query", "ref.query")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFEEDBACK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var FeedBacks []models.RefFeedBack
	var FeedBack *models.RefFeedBack
	if err = cursor.All(ctx.CTX, &FeedBacks); err != nil {
		return nil, err
	}
	if len(FeedBacks) > 0 {
		FeedBack = &FeedBacks[0]
	}
	return FeedBack, nil
}

//UpdateFeedBack : ""
func (d *Daos) UpdateFeedBack(ctx *models.Context, FeedBack *models.FeedBack) error {
	selector := bson.M{"_id": FeedBack.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": FeedBack, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFEEDBACK).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFeedBack : ""
func (d *Daos) FilterFeedBack(ctx *models.Context, FeedBackfilter *models.FeedBackFilter, pagination *models.Pagination) ([]models.RefFeedBack, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.FeedBackAndConstructQuery(ctx, FeedBackfilter)...)
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("farmer pagenation query =>", paginationPipeline)
		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONFEEDBACK).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "createdBy", "_id", "ref.createdBy", "ref.createdBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONQUERY, "query", "_id", "ref.query", "ref.query")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("FeedBack query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFEEDBACK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var FeedBacks []models.RefFeedBack
	if err = cursor.All(context.TODO(), &FeedBacks); err != nil {
		return nil, err
	}
	return FeedBacks, nil
}

//EnableFeedBack :""
func (d *Daos) EnableFeedBack(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FEEDBACKSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFEEDBACK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFeedBack :""
func (d *Daos) DisableFeedBack(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FEEDBACKSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFEEDBACK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFeedBack :""
func (d *Daos) DeleteFeedBack(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FEEDBACKSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFEEDBACK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) FeedBackAndConstructQuery(ctx *models.Context, FeedBackfilter *models.FeedBackFilter) []bson.M {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)
	//Adding $match from filter

	if FeedBackfilter != nil {

		if len(FeedBackfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": FeedBackfilter.ActiveStatus}})
		}
		if len(FeedBackfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": FeedBackfilter.Status}})
		}

		if len(FeedBackfilter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": FeedBackfilter.Village}})
		}
		if len(FeedBackfilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": FeedBackfilter.GramPanchayat}})
		}
		if len(FeedBackfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": FeedBackfilter.Block}})
		}
		if len(FeedBackfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": FeedBackfilter.District}})
		}
		if len(FeedBackfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": FeedBackfilter.State}})
		}
		if len(FeedBackfilter.FeedbackType) > 0 {
			query = append(query, bson.M{"feedbackType": bson.M{"$in": FeedBackfilter.FeedbackType}})
		}
		if len(FeedBackfilter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": FeedBackfilter.Type}})
		}
		if len(FeedBackfilter.Content) > 0 {
			query = append(query, bson.M{"content": bson.M{"$in": FeedBackfilter.Content}})
		}
		if len(FeedBackfilter.ContentOrganisation) > 0 {
			query = append(query, bson.M{"ref.content.organisation": bson.M{"$in": FeedBackfilter.ContentOrganisation}})
		}
		if len(FeedBackfilter.ContentProject) > 0 {
			query = append(query, bson.M{"ref.content.project": bson.M{"$in": FeedBackfilter.ContentProject}})
		}
		if FeedBackfilter.DateRange != nil {
			//var sd,ed time.Time
			if FeedBackfilter.DateRange.From != nil {
				sd := time.Date(FeedBackfilter.DateRange.From.Year(), FeedBackfilter.DateRange.From.Month(), FeedBackfilter.DateRange.From.Day(), 0, 0, 0, 0, FeedBackfilter.DateRange.From.Location())
				ed := time.Date(FeedBackfilter.DateRange.From.Year(), FeedBackfilter.DateRange.From.Month(), FeedBackfilter.DateRange.From.Day(), 23, 59, 59, 0, FeedBackfilter.DateRange.From.Location())
				if FeedBackfilter.DateRange.To != nil {
					ed = time.Date(FeedBackfilter.DateRange.To.Year(), FeedBackfilter.DateRange.To.Month(), FeedBackfilter.DateRange.To.Day(), 23, 59, 59, 0, FeedBackfilter.DateRange.To.Location())
				}
				query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

			}
		}

		//Regex
		if FeedBackfilter.Regex.Type != "" {
			query = append(query, bson.M{"type": primitive.Regex{Pattern: FeedBackfilter.Regex.Type, Options: "xi"}})
		}
		if FeedBackfilter.Regex.Feedback != "" {
			query = append(query, bson.M{"feedback": primitive.Regex{Pattern: FeedBackfilter.Regex.Feedback, Options: "xi"}})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	return mainPipeline
}

func (d *Daos) ConsolidatedFeedBack(ctx *models.Context, FeedBackfilter *models.FeedBackFilter) ([]models.FeedBackRating, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.FeedBackAndConstructQuery(ctx, FeedBackfilter)...)
	mainPipeline = append(mainPipeline, bson.M{
		"$facet": bson.M{

			"1": []bson.M{

				bson.M{"$match": bson.M{"rating": 1}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"2": []bson.M{

				bson.M{"$match": bson.M{"rating": 2}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"3": []bson.M{

				bson.M{"$match": bson.M{"rating": 3}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"4": []bson.M{

				bson.M{"$match": bson.M{"rating": 4}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"5": []bson.M{

				bson.M{"$match": bson.M{"rating": 5}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"6": []bson.M{

				bson.M{"$match": bson.M{"rating": 6}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"7": []bson.M{

				bson.M{"$match": bson.M{"rating": 7}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"8": []bson.M{

				bson.M{"$match": bson.M{"rating": 8}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"9": []bson.M{

				bson.M{"$match": bson.M{"rating": 9}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},

			"10": []bson.M{

				bson.M{"$match": bson.M{"rating": 10}},

				bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{

		"1": bson.M{"$arrayElemAt": []interface{}{"$1", 0}},

		"2": bson.M{"$arrayElemAt": []interface{}{"$2", 0}},

		"3": bson.M{"$arrayElemAt": []interface{}{"$3", 0}},

		"4": bson.M{"$arrayElemAt": []interface{}{"$4", 0}},

		"5": bson.M{"$arrayElemAt": []interface{}{"$5", 0}},

		"6": bson.M{"$arrayElemAt": []interface{}{"$6", 0}},

		"7": bson.M{"$arrayElemAt": []interface{}{"$7", 0}},

		"8": bson.M{"$arrayElemAt": []interface{}{"$8", 0}},
		"9": bson.M{"$arrayElemAt": []interface{}{"$9", 0}},

		"10": bson.M{"$arrayElemAt": []interface{}{"$10", 0}},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{

		"1":  "$1.count",
		"2":  "$2.count",
		"3":  "$3.count",
		"4":  "$4.count",
		"5":  "$5.count",
		"6":  "$6.count",
		"7":  "$7.count",
		"8":  "$8.count",
		"9":  "$9.count",
		"10": "$10.count",
	}})

	d.Shared.BsonToJSONPrintTag("feedback query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFEEDBACK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var feedbacks []models.FeedBackRating
	if err = cursor.All(context.TODO(), &feedbacks); err != nil {
		return nil, err
	}
	return feedbacks, nil

}
