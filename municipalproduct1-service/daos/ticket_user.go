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

// SaveTicketUser : ""
func (d *Daos) SaveTicketUser(ctx *models.Context, ticketUser *models.TicketUser) error {
	d.Shared.BsonToJSONPrint(ticketUser)
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).InsertOne(ctx.CTX, ticketUser)
	return err
}

// GetSingleTicketUser : ""
func (d *Daos) GetSingleTicketUser(ctx *models.Context, UniqueID string) (*models.RefTicketUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ticketUsers []models.RefTicketUser
	var ticketUser *models.RefTicketUser
	if err = cursor.All(ctx.CTX, &ticketUsers); err != nil {
		return nil, err
	}
	if len(ticketUsers) > 0 {
		ticketUser = &ticketUsers[0]
	}
	return ticketUser, nil
}

// UpdateTicketUser : ""
func (d *Daos) UpdateTicketUser(ctx *models.Context, ticketUser *models.TicketUser) error {
	selector := bson.M{"uniqueId": ticketUser.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": ticketUser}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableTicketUser : ""
func (d *Daos) EnableTicketUser(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TICKETUSERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableTicketUser : ""
func (d *Daos) DisableTicketUser(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TICKETUSERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteTicketUser : ""
func (d *Daos) DeleteTicketUser(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TICKETUSERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTicketUser : ""
func (d *Daos) FilterTicketUser(ctx *models.Context, filter *models.TicketUserFilter, pagination *models.Pagination) ([]models.RefTicketUser, error) {
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
		if len(filter.UserID) > 0 {
			query = append(query, bson.M{"userId": bson.M{"$in": filter.UserID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTICKETUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ticketUser []models.RefTicketUser
	if err = cursor.All(context.TODO(), &ticketUser); err != nil {
		return nil, err
	}
	return ticketUser, nil
}
