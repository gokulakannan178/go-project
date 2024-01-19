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

//SaveBasicMobileTowerUpdate :""
func (d *Daos) SaveBasicMobileTowerUpdate(ctx *models.Context, btlul *models.BasicMobileTowerUpdateLog) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).InsertOne(ctx.CTX, btlul)
	return err
}

// BasicMobileTowerUpdate : ""
func (d *Daos) BasicMobileTowerUpdate(ctx *models.Context, mtul *models.BasicMobileTowerUpdateLog) error {
	selector1 := bson.M{"uniqueId": mtul.MobileTowerID}
	data1 := bson.M{"$set": bson.M{"address": mtul.New.Address, "propertyId": mtul.New.PropertyID,
		"dateFrom": mtul.New.DateFrom, "builtUpArea": mtul.New.BuiltUpArea}}
	res1, err1 := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).UpdateOne(ctx.CTX, selector1, data1)
	if err1 != nil {
		return err1
	}
	fmt.Println(res1)

	return nil
}
func (d *Daos) GetSingleBasicMobileTowerUpdate(ctx *models.Context, uniqueID string) (*models.RefBasicMobileTowerUpdateLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})

	//Old Address Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "ref.previous.address.state", "ref.previous.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "ref.previous.address.district", "ref.previous.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "ref.previous.address.village", "ref.previous.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "ref.previous.address.zone", "ref.previous.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "ref.previous.address.ward", "ref.previous.address.ward")...)

	//New Address
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "ref.new.address.state", "ref.new.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "ref.address.district", "ref.new.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "ref.address.village", "ref.new.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "ref.address.zone", "ref.new.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "ref.new.address.ward", "ref.new.address.ward")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rmtuls []models.RefBasicMobileTowerUpdateLog
	var rmtul *models.RefBasicMobileTowerUpdateLog
	if err = cursor.All(ctx.CTX, &rmtuls); err != nil {
		return nil, err
	}
	if len(rmtuls) > 0 {
		rmtul = &rmtuls[0]
	}
	return rmtul, nil

}

// AcceptBasicMobileTowerUpdate : ""
func (d *Daos) AcceptBasicMobileTowerUpdate(ctx *models.Context, accept *models.AcceptBasicMobileTowerUpdate) error {

	t := time.Now()
	query := bson.M{"uniqueId": accept.UniqueID}
	var mtuls *models.BasicMobileTowerUpdateLog
	err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).FindOne(ctx.CTX, query).Decode(&mtuls)
	if err != nil {
		return errors.New("Not able to find the request" + err.Error())
	}
	if mtuls == nil {
		return errors.New("Request in nil")
	}
	mtul := new(models.BasicMobileTowerUpdateData)
	mtuls.MobileTowerID = mtul.MobileTowerID
	mtul.UpdateData.Address = mtuls.New.Address
	if err != nil {
		return errors.New("Error in updating MobileTower" + err.Error())
	}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERBASICUPDATELOGACCEPTED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectBasicMobileTowerUpdate : ""
func (d *Daos) RejectBasicMobileTowerUpdate(ctx *models.Context, reject *models.RejectBasicMobileTowerUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERBASICUPDATELOGREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSinglePreviousMobileTower(ctx *models.Context, UniqueID string) (*models.MobileTowerUpdateData, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.MobileTowerUpdateData
	var tower *models.MobileTowerUpdateData
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

//FilterBasicMobileTowerUpdateLog : ""
func (d *Daos) FilterBasicMobileTowerUpdateLog(ctx *models.Context, filter *models.FilterBasicMobileTowerUpdateLog, pagination *models.Pagination) ([]models.RefBasicMobileTowerUpdateLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.MobileTowerID) > 0 {
			query = append(query, bson.M{"mobileTowerId": bson.M{"$in": filter.MobileTowerID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).CountDocuments(ctx.CTX, func() bson.M {
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

	// lookup
	//Old Address Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "ref.previous.address.state", "ref.previous.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "ref.previous.address.district", "ref.previous.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "ref.previous.address.village", "ref.previous.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "ref.previous.address.zone", "ref.previous.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "ref.previous.address.ward", "ref.previous.address.ward")...)

	//New Address
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "ref.new.address.state", "ref.new.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "ref.address.district", "ref.new.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "ref.address.village", "ref.new.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "ref.address.zone", "ref.new.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "ref.new.address.ward", "ref.new.address.ward")...)

	// user Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userName", "userName", "ref.requestedBy", "ref.requestedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "userType", "uniqueId", "ref.requestedByType", "ref.requestedByType")...)
	// Action Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionBy", "ref.actionBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "action.byType", "uniqueId", "ref.actionByType", "ref.actionByType")...)

	d.Shared.BsonToJSONPrintTag("mobiletower query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refs []models.RefBasicMobileTowerUpdateLog
	if err = cursor.All(context.TODO(), &refs); err != nil {
		return nil, err
	}
	return refs, nil
}

// GetSingleBasicMobileTowerUpdateLogV2
func (d *Daos) GetSingleBasicMobileTowerUpdateLogV2(ctx *models.Context, uniqueID string) (*models.BasicMobileTowerUpdateLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//Old Address Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "previous.address.stateCode", "code", "ref.previous.address.state", "ref.previous.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "previous.address.districtCode", "code", "ref.previous.address.district", "ref.previous.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "previous.address.villageCode", "code", "ref.previous.address.village", "ref.previous.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "previous.address.zoneCode", "code", "ref.previous.address.zone", "ref.previous.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "previous.address.wardCode", "code", "ref.previous.address.ward", "ref.previous.address.ward")...)

	//New Address
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "new.address.stateCode", "code", "ref.new.address.state", "ref.new.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "new.address.districtCode", "code", "ref.address.district", "ref.new.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "new.address.villageCode", "code", "ref.address.village", "ref.new.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "new.address.zoneCode", "code", "ref.address.zone", "ref.new.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "new.address.wardCode", "code", "ref.new.address.ward", "ref.new.address.ward")...)

	// user Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userName", "userName", "ref.requestedBy", "ref.requestedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "userType", "type", "ref.requestedByType", "ref.requestedByType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("mobiletower query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var mtuls []models.BasicMobileTowerUpdateLog
	var mtul *models.BasicMobileTowerUpdateLog
	if err = cursor.All(ctx.CTX, &mtuls); err != nil {
		return nil, err
	}
	if len(mtuls) > 0 {
		mtul = &mtuls[0]
	}
	return mtul, nil

}

//BasicMobileTowerUpdateGetPaymentsToBeUpdated : ""
func (d *Daos) BasicMobileTowerUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbmtul *models.RefBasicMobileTowerUpdateLogV2) ([]models.RefMobileTowerPayments, error) {
	//get current Financial year

	cfy, err := d.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	if cfy == nil {
		return nil, errors.New("current financial year is nil")
	}
	sd := time.Date(cfy.From.Year(), cfy.From.Month(), cfy.From.Day(), 0, 0, 0, 0, cfy.From.Location())
	ed := time.Date(cfy.To.Year(), cfy.To.Month(), cfy.To.Day(), 23, 59, 59, 0, cfy.To.Location())
	fmt.Println("sd ===>", sd)
	fmt.Println("ed ===>", ed)
	mobileTowerPaymentFindQuery := bson.M{
		"status":         constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		"mobileTowerId":  rbmtul.MobileTowerID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("mobileTower query =>", mobileTowerPaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).Find(ctx.CTX, mobileTowerPaymentFindQuery, nil)
	if err != nil {
		return nil, err
	}
	var mobileTowerPayments []models.RefMobileTowerPayments
	if err = cursor.All(context.TODO(), &mobileTowerPayments); err != nil {
		return nil, err
	}

	return mobileTowerPayments, nil
}
