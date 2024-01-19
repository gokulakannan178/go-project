package daos

import (
	"context"
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//BasicShopRentUpdateGetPaymentsToBeUpdated : ""
func (d *Daos) BasicShopRentUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbsrul *models.RefBasicShopRentUpdateLog) ([]models.RefShopRentPayments, error) {
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
	fmt.Println("sd ====>", sd)
	fmt.Println("ed ====>", ed)
	shopRentPaymentFindQuery := bson.M{
		"status":         constants.SHOPRENTPAYMENTSTATUSCOMPLETED,
		"shopRentId":     rbsrul.ShopRentID,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("shopRent query =>", shopRentPaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Find(ctx.CTX, shopRentPaymentFindQuery, nil)
	if err != nil {
		return nil, err
	}
	var shopRentPayments []models.RefShopRentPayments
	if err = cursor.All(context.TODO(), &shopRentPayments); err != nil {
		return nil, err
	}

	return shopRentPayments, nil
}
