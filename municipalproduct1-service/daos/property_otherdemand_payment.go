package daos

import (
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SavePropertyOtherDemandPayment :""
func (d *Daos) SavePropertyOtherDemandPayment(ctx *models.Context, payment *models.PropertyOtherDemandPayment) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENT).InsertOne(ctx.SC, payment)
	return err
}

// GetSinglePropertyOtherDemandPaymentWithTxtID :""
func (d *Daos) GetSinglePropertyOtherDemandPaymentWithTxtID(ctx *models.Context, txtID string) (*models.PropertyOtherDemandPayment, error) {
	mainPipeline := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{

		"from": constants.COLLECTIONPROPERTYOTHERDEMANDPARTPAYMENT,
		"as":   "ref.partPayments",
		"let":  bson.M{"tnxId": "$tnxId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$tnxId", "$$tnxId"}},
				{"$eq": []string{"$status", constants.PROPERTYOTHERDEMANDPAYMENTCOMPLETED}},
			}}}},
		},
	}})

	d.Shared.BsonToJSONPrintTag("property payment query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.PropertyOtherDemandPayment
	var pp *models.PropertyOtherDemandPayment
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	return pp, nil
}

// SavePropertyOtherDemandPaymentDemandBasic :""
func (d *Daos) SavePropertyOtherDemandPaymentDemandBasic(ctx *models.Context, basic *models.PropertyOtherDemandPaymentDemandBasic) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSBASIC).InsertOne(ctx.SC, basic)
	return err
}

// GetSinglePropertyOtherDemandPaymentDemandBasicWithTxtID :""
func (d *Daos) GetSinglePropertyOtherDemandPaymentDemandBasicWithTxtID(ctx *models.Context, txtID string) (*models.PropertyOtherDemandPaymentDemandBasic, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSBASIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdbs []models.PropertyOtherDemandPaymentDemandBasic
	var ppdb *models.PropertyOtherDemandPaymentDemandBasic
	if err = cursor.All(ctx.CTX, &ppdbs); err != nil {
		return nil, err
	}
	if len(ppdbs) > 0 {
		ppdb = &ppdbs[0]
	}
	return ppdb, nil
}

//SaveManyPropertyOtherDemandPaymentDemandFy :""
func (d *Daos) SaveManyPropertyOtherDemandPaymentDemandFy(ctx *models.Context, ppdfys []models.PropertyOtherDemandPaymentDemandFy) error {
	var data []interface{}
	for _, v := range ppdfys {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSFY).InsertMany(ctx.SC, data)
	return err
}

//GetPropertyOtherDemandPaymentDemandFycWithTxtID :""
func (d *Daos) GetPropertyOtherDemandPaymentDemandFycWithTxtID(ctx *models.Context, txtID string) ([]models.PropertyOtherDemandPaymentDemandFy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"tnxId": txtID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"fy.order": -1}})
	d.Shared.BsonToJSONPrintTag("property payment demand basics query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSFY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ppdfys []models.PropertyOtherDemandPaymentDemandFy
	if err = cursor.All(ctx.CTX, &ppdfys); err != nil {
		return nil, err
	}

	return ppdfys, nil
}

// CompletePropertyOtherDemandPaymentWithTxtID : ""
func (d *Daos) CompletePropertyOtherDemandPaymentWithTxtID(ctx *models.Context, payment *models.PropertyOtherDemandMakePayment, isPartPayment bool) error {
	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}

	selector := bson.M{"tnxId": payment.TnxID}
	update := bson.M{"$set": bson.M{
		"status":                    status,
		"details":                   payment.Details,
		"completionDate":            time.Now(),
		"reciptURL":                 payment.ReciptURL,
		"collectionReceived.status": "Pending",
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

// CompleteSinglePropertyOtherDemandPaymentDemandBasicWithTxtID : ""
func (d *Daos) CompleteSinglePropertyOtherDemandPaymentDemandBasicWithTxtID(ctx *models.Context, payment *models.PropertyOtherDemandMakePayment) error {
	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}
	selector := bson.M{"tnxId": payment.TnxID}
	update := bson.M{"$set": bson.M{
		"status": status,
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

// CompletePropertyOtherDemandPaymentDemandFycWithTxtID : ""
func (d *Daos) CompletePropertyOtherDemandPaymentDemandFycWithTxtID(ctx *models.Context, payment *models.PropertyOtherDemandMakePayment) error {
	status, err := d.GetMopPaymentStatus(ctx, payment.Details.MOP.Mode)
	if err != nil {
		return err
	}
	selector := bson.M{"tnxId": payment.TnxID}
	update := bson.M{"$set": bson.M{
		"status": status,
	}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSFY).UpdateMany(ctx.CTX, selector, update, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return nil
}

// PropertyOtherDemandVerifyPayment : ""
func (d *Daos) PropertyOtherDemandVerifyPayment(ctx *models.Context, vp *models.PropertyOtherDemandVerifyPayment) error {
	selector := bson.M{"tnxId": vp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDPAYMENTCOMPLETED, "remark": ""}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDPAYMENTCOMPLETED, "remark": "",
		"verifiedInfo": bson.M{"verifiedActionDate": vp.ActionDate, "verifiedDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// PropertyOtherDemandNotVerifiedPayment : ""
func (d *Daos) PropertyOtherDemandNotVerifiedPayment(ctx *models.Context, vp *models.PropertyOtherDemandNotVerifiedPayment) error {
	selector := bson.M{"tnxId": vp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDPAYMENTSTATUSNOTVERIFIED, "remark": vp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDPAYMENTSTATUSNOTVERIFIED, "remark": vp.Remarks,
		"notVerifiedInfo": bson.M{"notVerifiedActionDate": vp.ActionDate, "notVerifiedDate": vp.Date, "remark": vp.Remarks, "by": vp.By, "byType": vp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

// PropertyOtherDemandRejectPayment : ""
func (d *Daos) PropertyOtherDemandRejectPayment(ctx *models.Context, rp *models.PropertyOtherDemandRejectPayment) error {
	selector := bson.M{"tnxId": rp.TnxID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDPAYMENTSTATUSREJECTED, "rejectedRemark": rp.Remarks}}
	update2 := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDPAYMENTSTATUSREJECTED, "rejectedRemark": rp.Remarks,
		"rejectedInfo": bson.M{"rejectedActionDate": rp.ActionDate, "rejectedDate": rp.Date, "remark": rp.Remarks, "by": rp.By, "byType": rp.ByType},
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENT).UpdateOne(ctx.CTX, selector, update2)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSBASIC).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMANDPAYMENTSFY).UpdateMany(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}
