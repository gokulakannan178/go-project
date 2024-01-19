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

// SaveTradelicensePayeeNameChange : ""
func (d *Daos) SaveTradelicensePayeeNameChange(ctx *models.Context, ppnc *models.TradelicensePayeeNameChange) error {
	d.Shared.BsonToJSONPrint(ppnc)
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).InsertOne(ctx.CTX, ppnc)
	return err
}

// GetSingleTradelicensePayeeNameChange : ""
func (d *Daos) GetSingleTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) (*models.RefTradelicensePayeeNameChange, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var TradelicensePayeeNameChange []models.RefTradelicensePayeeNameChange
	var boringcharge *models.RefTradelicensePayeeNameChange
	if err = cursor.All(ctx.CTX, &TradelicensePayeeNameChange); err != nil {
		return nil, err
	}
	if len(TradelicensePayeeNameChange) > 0 {
		boringcharge = &TradelicensePayeeNameChange[0]
	}
	return boringcharge, nil
}

// UpdateTradelicensePayeeNameChange : ""
func (d *Daos) UpdateTradelicensePayeeNameChange(ctx *models.Context, ppnc *models.TradelicensePayeeNameChange) error {
	selector := bson.M{"uniqueId": ppnc.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": ppnc}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableTradelicensePayeeNameChange : ""
func (d *Daos) EnableTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEPAYEENAMECHANGESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableTradelicensePayeeNameChange: ""
func (d *Daos) DisableTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEPAYEENAMECHANGESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteTradelicensePayeeNameChange: ""
func (d *Daos) DeleteTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEPAYEENAMECHANGESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTradelicensePayeeNameChange : ""
func (d *Daos) FilterTradelicensePayeeNameChange(ctx *models.Context, filter *models.TradelicensePayeeNameChangeFilter, pagination *models.Pagination) ([]models.RefTradelicensePayeeNameChange, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.ApprovedBy) > 0 {
			query = append(query, bson.M{"approvedBy.by": bson.M{"$in": filter.ApprovedBy}})
		}
		if len(filter.CreatedBy) > 0 {
			query = append(query, bson.M{"createdBy.by": bson.M{"$in": filter.CreatedBy}})
		}
		if len(filter.RejectedBy) > 0 {
			query = append(query, bson.M{"rejectedBy.by": bson.M{"$in": filter.RejectedBy}})
		}
		if len(filter.PropertyId) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyId}})
		}
		if len(filter.ReceiptNo) > 0 {
			query = append(query, bson.M{"receiptNo": bson.M{"$in": filter.ReceiptNo}})
		}

	}
	if filter.DateRange != nil {
		//var sd,ed time.Time
		if filter.DateRange.From != nil {
			sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
			ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
			if filter.DateRange.To != nil {
				ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
			}
			query = append(query, bson.M{"createdOn.on": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"requestor.on": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppnc []models.RefTradelicensePayeeNameChange
	if err = cursor.All(context.TODO(), &ppnc); err != nil {
		return nil, err
	}
	return ppnc, nil
}

// ApproveTradeLicense : ""
func (d *Daos) ApproveTradelicensePayeeNameChange(ctx *models.Context, accept *models.ApproveTradelicensePayeeNameChange) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEPAYEENAMECHANGESTATUSACTIVE,
		"approvedBy": models.Action{
			On:     &t,
			By:     accept.UserName,
			ByType: accept.UserType,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// NotApproveTradeLicense : ""
func (d *Daos) NotApproveTradelicensePayeeNameChange(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	t := time.Now()

	query := bson.M{"uniqueId": notApprove.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEPAYEENAMECHANGESTATUSREJECTED,
		"rejectedBy": models.Updated{
			On:      &t,
			By:      notApprove.UserName,
			ByType:  notApprove.UserType,
			Remarks: notApprove.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
