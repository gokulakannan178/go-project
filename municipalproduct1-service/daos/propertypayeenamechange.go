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

// SavePropertyPayeeNameChange : ""
func (d *Daos) SavePropertyPayeeNameChange(ctx *models.Context, ppnc *models.PropertyPayeeNameChange) error {
	d.Shared.BsonToJSONPrint(ppnc)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).InsertOne(ctx.CTX, ppnc)
	return err
}

// GetSinglePropertyPayeeNameChange : ""
func (d *Daos) GetSinglePropertyPayeeNameChange(ctx *models.Context, UniqueID string) (*models.RefPropertyPayeeNameChange, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var PropertyPayeeNameChange []models.RefPropertyPayeeNameChange
	var boringcharge *models.RefPropertyPayeeNameChange
	if err = cursor.All(ctx.CTX, &PropertyPayeeNameChange); err != nil {
		return nil, err
	}
	if len(PropertyPayeeNameChange) > 0 {
		boringcharge = &PropertyPayeeNameChange[0]
	}
	return boringcharge, nil
}

// UpdatePropertyPayeeNameChange : ""
func (d *Daos) UpdatePropertyPayeeNameChange(ctx *models.Context, ppnc *models.PropertyPayeeNameChange) error {
	selector := bson.M{"uniqueId": ppnc.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": ppnc}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyPayeeNameChange : ""
func (d *Daos) EnablePropertyPayeeNameChange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYEENAMECHANGESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyPayeeNameChange: ""
func (d *Daos) DisablePropertyPayeeNameChange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYEENAMECHANGESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyPayeeNameChange: ""
func (d *Daos) DeletePropertyPayeeNameChange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYEENAMECHANGESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyPayeeNameChange : ""
func (d *Daos) FilterPropertyPayeeNameChange(ctx *models.Context, filter *models.PropertyPayeeNameChangeFilter, pagination *models.Pagination) ([]models.RefPropertyPayeeNameChange, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppnc []models.RefPropertyPayeeNameChange
	if err = cursor.All(context.TODO(), &ppnc); err != nil {
		return nil, err
	}
	return ppnc, nil
}

// ApproveTradeLicense : ""
func (d *Daos) ApprovePropertyPayeeNameChange(ctx *models.Context, accept *models.ApprovePropertyPayeeNameChange) error {
	t := time.Now()

	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYEENAMECHANGESTATUSACTIVE,
		"approvedBy": models.Action{
			On:     &t,
			By:     accept.UserName,
			ByType: accept.UserType,
		},
	}}

	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// NotApproveTradeLicense : ""
func (d *Daos) NotApprovePropertyPayeeNameChange(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	t := time.Now()

	query := bson.M{"uniqueId": notApprove.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYEENAMECHANGESTATUSREJECTED,
		"rejectedBy": models.Updated{
			On:      &t,
			By:      notApprove.UserName,
			ByType:  notApprove.UserType,
			Remarks: notApprove.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYEENAMEHANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
