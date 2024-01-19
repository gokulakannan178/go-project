package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveTicketComment : ""
func (d *Daos) SaveTicketComment(ctx *models.Context, ticketComment *models.TicketComment) error {
	d.Shared.BsonToJSONPrint(ticketComment)
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).InsertOne(ctx.CTX, ticketComment)
	return err
}

// GetSingleTicketComment : ""
func (d *Daos) GetSingleTicketComment(ctx *models.Context, UniqueID string) (*models.RefTicketComment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ticketComments []models.RefTicketComment
	var ticketComment *models.RefTicketComment
	if err = cursor.All(ctx.CTX, &ticketComments); err != nil {
		return nil, err
	}
	if len(ticketComments) > 0 {
		ticketComment = &ticketComments[0]
	}
	return ticketComment, nil
}

// UpdateTicketComment : ""
func (d *Daos) UpdateTicketComment(ctx *models.Context, ticketComment *models.TicketComment) error {
	selector := bson.M{"uniqueId": ticketComment.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": ticketComment}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableTicketComment : ""
func (d *Daos) EnableTicketComment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TICKETCOMMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableTicketComment : ""
func (d *Daos) DisableTicketComment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TICKETCOMMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteTicketComment : ""
func (d *Daos) DeleteTicketComment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TICKETCOMMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTicketComment : ""
func (d *Daos) FilterTicketComment(ctx *models.Context, filter *models.TicketCommentFilter, pagination *models.Pagination) ([]models.RefTicketComment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.TicketID) > 0 {
			query = append(query, bson.M{"ticketId": bson.M{"$in": filter.TicketID}})
		}
		if filter.CreatedDateRange != nil {
			//var sd,ed time.Time
			if filter.CreatedDateRange.From != nil {
				sd := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 0, 0, 0, 0, filter.CreatedDateRange.From.Location())
				ed := time.Date(filter.CreatedDateRange.From.Year(), filter.CreatedDateRange.From.Month(), filter.CreatedDateRange.From.Day(), 23, 59, 59, 0, filter.CreatedDateRange.From.Location())
				if filter.CreatedDateRange.To != nil {
					ed = time.Date(filter.CreatedDateRange.To.Year(), filter.CreatedDateRange.To.Month(), filter.CreatedDateRange.To.Day(), 23, 59, 59, 0, filter.CreatedDateRange.To.Location())
				}
				query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTICKETCOMMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ticketComment []models.RefTicketComment
	if err = cursor.All(context.TODO(), &ticketComment); err != nil {
		return nil, err
	}
	return ticketComment, nil
}
