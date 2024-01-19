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

// SavePropertyUserCharge : ""
func (d *Daos) SavePropertyUserCharge(ctx *models.Context, propertyusercharge *models.PropertyUserCharge) error {
	d.Shared.BsonToJSONPrint(propertyusercharge)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).InsertOne(ctx.CTX, propertyusercharge)
	return err
}

// GetSinglePropertyUserCharge : ""
func (d *Daos) GetSinglePropertyUserCharge(ctx *models.Context, UniqueID string) (*models.RefPropertyUserCharge, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefPropertyUserCharge
	var tower *models.RefPropertyUserCharge
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdatePropertyUserCharge : ""
func (d *Daos) UpdatePropertyUserCharge(ctx *models.Context, business *models.PropertyUserCharge) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyUserCharge : ""
func (d *Daos) EnablePropertyUserCharge(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYUSERCHARGESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyUserCharge : ""
func (d *Daos) DisablePropertyUserCharge(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYUSERCHARGESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyUserCharge : ""
func (d *Daos) DeletePropertyUserCharge(ctx *models.Context, rp *models.UserChargeAction) error {
	query := bson.M{"uniqueId": rp.UniqueId}
	update := bson.M{"$set": bson.M{"userCharge.status": constants.PROPERTYUSERCHARGESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyUserCharge : ""
func (d *Daos) FilterPropertyUserCharge(ctx *models.Context, filter *models.PropertyUserChargeFilter, pagination *models.Pagination) ([]models.RefPropertyUserCharge, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYUSERCHARGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyusercharge []models.RefPropertyUserCharge
	if err = cursor.All(context.TODO(), &propertyusercharge); err != nil {
		return nil, err
	}
	return propertyusercharge, nil
}

// VerifyPayment : ""
func (d *Daos) VerifyPropertyUserCharge(ctx *models.Context, vp *models.UserChargeAction) error {
	selector := bson.M{"uniqueId": vp.UniqueId}
	//update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTCOMPLETED, "remark": ""}}
	update2 := bson.M{"$set": bson.M{"userCharge.status": constants.PROPERTYUSERCHARGESTATUSACTIVE,
		"userchargeverifiedInfo": bson.M{"actionDate": vp.ActionDate, "actualDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}

	return nil
}

// NotVerifiedPayment : ""
func (d *Daos) RejectPropertyUserCharge(ctx *models.Context, vp *models.UserChargeAction) error {
	selector := bson.M{"uniqueId": vp.UniqueId}
	//update := bson.M{"$set": bson.M{"status": constants.PROPERTYPAYMENTNOTVERIFIED, "remark": vp.Remarks}}
	update2 := bson.M{"$set": bson.M{"userCharge.status": constants.PROPERTYUSERCHARGESTATUSREJECTED,
		"userchargerejectedInfo": bson.M{"actionDate": vp.ActionDate, "actualDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}
