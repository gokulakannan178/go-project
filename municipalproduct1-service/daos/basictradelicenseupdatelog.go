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

//SaveBasicPropertyUpdateLog :""
func (d *Daos) SaveBasicTradeLicenseUpdateLog(ctx *models.Context, btlul *models.BasicTradeLicenseUpdateLog) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).InsertOne(ctx.CTX, btlul)
	return err
}

// BasicUpdateTradeLicense : ""
func (d *Daos) BasicUpdateTradeLicense(ctx *models.Context, btlu *models.BasicTradeLicenseUpdateLog) error {
	selector1 := bson.M{"uniqueId": btlu.TradeLicenseID}
	data1 := bson.M{"$set": bson.M{"address": btlu.New.Address, "desc": btlu.New.Desc, "tlbtId": btlu.New.TLBTID,
		"tlctId": btlu.New.TLCTID, "mobileNo": btlu.New.MobileNo, "ownerName": btlu.New.OwnerName, "guardianName": btlu.New.GuardianName, "businessName": btlu.New.BusinessName, "photo": btlu.New.Photo}}
	res1, err1 := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, selector1, data1)
	if err1 != nil {
		return err1
	}
	fmt.Println(res1)
	// opts := options.Update().SetUpsert(true)
	// selector2 := bson.M{"uniqueId": btlu.UpdateData.OwnerName.UniqueID, "propertyId": btlu.TradeLicenseID}
	// data2 := bson.M{"$set": btlu.Owner}
	// res2, err2 := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, selector2, data2, opts)
	// if err2 != nil {
	// 	return err2
	// }
	// fmt.Println(res2)
	return nil
}
func (d *Daos) GetSingleBasicTradeLicenseUpdateLog(ctx *models.Context, uniqueID string) (*models.RefBasicTradeLicenseUpdateLog, error) {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rbtluls []models.RefBasicTradeLicenseUpdateLog
	var rbtlul *models.RefBasicTradeLicenseUpdateLog
	if err = cursor.All(ctx.CTX, &rbtluls); err != nil {
		return nil, err
	}
	if len(rbtluls) > 0 {
		rbtlul = &rbtluls[0]
	}
	return rbtlul, nil

}

// AcceptBasicTradeLicenseUpdate : ""
func (d *Daos) AcceptBasicTradeLicenseUpdate(ctx *models.Context, accept *models.AcceptBasicTradeLicenseUpdate) error {

	t := time.Now()
	query := bson.M{"uniqueId": accept.UniqueID}
	var btlul *models.BasicTradeLicenseUpdateLog
	err := ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).FindOne(ctx.CTX, query).Decode(&btlul)
	if err != nil {
		return errors.New("Not able to find the request" + err.Error())
	}
	if btlul == nil {
		return errors.New("Request in nil")
	}
	btlu := new(models.BasicTradeLicenseUpdate)
	btlu.TradeLicenseID = btlul.TradeLicenseID
	btlu.UpdateData.Address = btlul.New.Address
	// btlu.Owner = btlul.New.Owner
	// btlu.UserName = btlul.Requester.By
	// btlu.UserType = btlul.Requester.ByType
	// btlu.Proof = btlul.Proof
	// btlu.Remarks = btlul.Requester.Remarks
	//rr = d.BasicUpdateTradeLicenseUpdate(ctx, btlu)
	if err != nil {
		return errors.New("Error in updating TradeLicense" + err.Error())
	}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEBASICUPDATELOGACCEPTED,
		"action": models.Updated{
			On:      &t,
			By:      accept.UserName,
			ByType:  accept.UserType,
			Remarks: accept.Remark,
		},
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RejectBasicTradeLicenseUpdate : ""
func (d *Daos) RejectBasicTradeLicenseUpdate(ctx *models.Context, reject *models.RejectBasicTradeLicenseUpdate) error {
	t := time.Now()

	query := bson.M{"uniqueId": reject.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEBASICUPDATELOGREJECTED,
		"action": models.Updated{
			On:      &t,
			By:      reject.UserName,
			ByType:  reject.UserType,
			Remarks: reject.Remark,
		},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSinglePreviousTradeLicense(ctx *models.Context, UniqueID string) (*models.TradeLicenseUpdateData, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.TradeLicenseUpdateData
	var tower *models.TradeLicenseUpdateData
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

//FilterBasicTradeLicenseUpdateLog : ""
func (d *Daos) FilterBasicTradeLicenseUpdateLog(ctx *models.Context, filter *models.FilterBasicTradeLicenseUpdateLog, pagination *models.Pagination) ([]models.RefBasicTradeLicenseUpdateLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.TradeLicenseID) > 0 {
			query = append(query, bson.M{"tradeLicenseId": bson.M{"$in": filter.TradeLicenseID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).CountDocuments(ctx.CTX, func() bson.M {
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

	// //Aggregation
	d.Shared.BsonToJSONPrintTag("shoprent query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refs []models.RefBasicTradeLicenseUpdateLog
	if err = cursor.All(context.TODO(), &refs); err != nil {
		return nil, err
	}
	return refs, nil
}

// GetSingleBasicTradeLicenseUpdateLogV2
func (d *Daos) GetSingleBasicTradeLicenseUpdateLogV2(ctx *models.Context, uniqueID string) (*models.BasicTradeLicenseUpdateLog, error) {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "userType", "uniqueId", "ref.requestedByType", "ref.requestedByType")...)
	// Action Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "action.by", "userName", "ref.actionBy", "ref.actionBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "action.byType", "uniqueId", "ref.actionByType", "ref.actionByType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("tradelicense query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBASICTRADELICENSEUPDATELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rbtluls []models.BasicTradeLicenseUpdateLog
	var rbtlul *models.BasicTradeLicenseUpdateLog
	if err = cursor.All(ctx.CTX, &rbtluls); err != nil {
		return nil, err
	}
	if len(rbtluls) > 0 {
		rbtlul = &rbtluls[0]
	}
	return rbtlul, nil

}

//BasicTradeLicenseUpdateGetPaymentsToBeUpdated : ""
func (d *Daos) BasicTradeLicenseUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbtlul *models.RefBasicTradeLicenseUpdateLogV2) ([]models.RefTradeLicensePayments, error) {
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
	tradeLicencePaymentFindQuery := bson.M{
		"status":         constants.TRADELICENSEPAYMENRSTATUSCOMPLETED,
		"tradeLicenseId": rbtlul.TradeLicenseID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("tradeLicense query =>", tradeLicencePaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Find(ctx.CTX, tradeLicencePaymentFindQuery, nil)
	if err != nil {
		return nil, err
	}
	var tradeLicencePayments []models.RefTradeLicensePayments
	if err = cursor.All(context.TODO(), &tradeLicencePayments); err != nil {
		return nil, err
	}

	return tradeLicencePayments, nil
}
