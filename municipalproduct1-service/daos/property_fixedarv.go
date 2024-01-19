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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SavePropertyFixedArv : ""
func (d *Daos) SavePropertyFixedArv(ctx *models.Context, propertyfixedarv *models.PropertyFixedArv) error {
	d.Shared.BsonToJSONPrint(propertyfixedarv)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).InsertOne(ctx.CTX, propertyfixedarv)
	return err
}

// UpsertPropertyFixedArv : ""
func (d *Daos) UpsertPropertyFixedArv(ctx *models.Context, fixedArv *models.PropertyFixedArv) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"fyId": fixedArv.FyID, "propertyId": fixedArv.PropertyID}
	updateData := bson.M{"$set": fixedArv}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in upserting - " + err.Error())
	}
	return nil
}

func (d *Daos) SavePropertyFixedARVV2(ctx *models.Context, propertyfixedArv *models.PropertyFixedArv) error {
	insertdata := []interface{}{}
	for _, v := range propertyfixedArv.FyIDs {
		insertdata = append(insertdata, v)
		propertyfixedArv.FyID = v
		fmt.Println("propertyfixedArv.FyID  =====>", propertyfixedArv.FyID)
	}
	insertdata = append(insertdata, propertyfixedArv)
	result, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).InsertMany(ctx.SC, insertdata)
	fmt.Println("insert result =>", result)
	return err

}

// GetSinglePropertyFixedArv : ""
func (d *Daos) GetSinglePropertyFixedArv(ctx *models.Context, UniqueID string) (*models.RefPropertyFixedArv, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.createdBy", "ref.createdBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "created.byType", "uniqueId", "ref.createdByType", "ref.createdByType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "fyId", "uniqueId", "ref.financialYear", "ref.financialYear")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedBy", "ref.requestedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approved.by", "userName", "ref.approvedBy", "ref.approvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejected.by", "userName", "ref.rejectedBy", "ref.rejectedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requested.byType", "uniqueId", "ref.requestedByType", "ref.requestedByType")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyfixedarvs []models.RefPropertyFixedArv
	var propertyfixedarv *models.RefPropertyFixedArv
	if err = cursor.All(ctx.CTX, &propertyfixedarvs); err != nil {
		return nil, err
	}
	if len(propertyfixedarvs) > 0 {
		propertyfixedarv = &propertyfixedarvs[0]
	}
	return propertyfixedarv, nil
}
func (d *Daos) GetSinglePropertyFixedArvWithFyID(ctx *models.Context, UniqueID string, PropertyID string) (*models.RefPropertyFixedArv, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"fyId": UniqueID, "propertyId": PropertyID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.createdBy", "ref.createdBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "created.byType", "uniqueId", "ref.createdByType", "ref.createdByType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyfixedarvs []models.RefPropertyFixedArv
	var propertyfixedarv *models.RefPropertyFixedArv
	if err = cursor.All(ctx.CTX, &propertyfixedarvs); err != nil {
		return nil, err
	}
	if len(propertyfixedarvs) > 0 {
		propertyfixedarv = &propertyfixedarvs[0]
	}
	return propertyfixedarv, nil
}

// UpdatePropertyFixedArv : ""
func (d *Daos) UpdatePropertyFixedArv(ctx *models.Context, propertyfixedarv *models.PropertyFixedArv) error {
	selector := bson.M{"uniqueId": propertyfixedarv.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": propertyfixedarv}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyFixedArv : ""
func (d *Daos) EnablePropertyFixedArv(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyFixedArv : ""
func (d *Daos) DisablePropertyFixedArv(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyFixedArv : ""
func (d *Daos) DeletePropertyFixedArv(ctx *models.Context, uniqueId string) error {
	query := bson.M{"uniqueId": uniqueId}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyFixedArv : ""
func (d *Daos) FilterPropertyFixedArv(ctx *models.Context, filter *models.PropertyFixedArvFilter, pagination *models.Pagination) ([]models.RefPropertyFixedArv, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueIDs}})
		}
		if len(filter.FyID) > 0 {
			query = append(query, bson.M{"fyId": bson.M{"$in": filter.FyID}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if len(filter.ARV) > 0 {
			query = append(query, bson.M{"arv": bson.M{"$in": filter.ARV}})
		}
		if len(filter.CreatedBy) > 0 {
			query = append(query, bson.M{"Created.by": bson.M{"$in": filter.CreatedBy}})
		}
		//regex

		if filter.Regex.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.Regex.UniqueID, Options: "xi"}})
		}
		if filter.Regex.PropertyID != "" {
			query = append(query, bson.M{"propertyId": primitive.Regex{Pattern: filter.Regex.PropertyID, Options: "xi"}})
		}
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
	}
	if filter.CreatedDate != nil {
		//var sd,ed time.Time
		if filter.CreatedDate.From != nil {
			sd := time.Date(filter.CreatedDate.From.Year(), filter.CreatedDate.From.Month(), filter.CreatedDate.From.Day(), 0, 0, 0, 0, filter.CreatedDate.From.Location())
			ed := time.Date(filter.CreatedDate.From.Year(), filter.CreatedDate.To.Month(), filter.CreatedDate.To.Day(), 23, 59, 59, 0, filter.CreatedDate.To.Location())
			if filter.CreatedDate.To != nil {
				ed = time.Date(filter.CreatedDate.To.Year(), filter.CreatedDate.To.Month(), filter.CreatedDate.To.Day(), 23, 59, 59, 0, filter.CreatedDate.To.Location())
			}
			query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookup

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.createdBy", "ref.createdBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "created.byType", "uniqueId", "ref.createdByType", "ref.createdByType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "fyId", "uniqueId", "ref.financialYear", "ref.financialYear")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "requester.by", "userName", "ref.requestedBy", "ref.requestedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approved.by", "userName", "ref.approvedBy", "ref.approvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejected.by", "userName", "ref.rejectedBy", "ref.rejectedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "requested.byType", "uniqueId", "ref.requestedByType", "ref.requestedByType")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property fixed arv query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyfixedarv []models.RefPropertyFixedArv
	if err = cursor.All(context.TODO(), &propertyfixedarv); err != nil {
		return nil, err
	}
	return propertyfixedarv, nil
}

// AcceptPropertyFixedArv : ""
func (d *Daos) AcceptPropertyFixedArv(ctx *models.Context, accept *models.AcceptPropertyFixedArv) error {
	t := time.Now()
	query := bson.M{"uniqueId": accept.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVSTATUSACTIVE,
		"approved": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectPropertyFixedArv : ""
func (d *Daos) RejectPropertyFixedArv(ctx *models.Context, reject *models.RejectPropertyFixedArv) error {
	t := time.Now()
	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFIXEDARVSTATUSREJECTED,
		"rejected": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFIXEDARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
