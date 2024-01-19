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

// BasicSolidWasteUpdateGetPaymentsToBeUpdated : ""
func (d *Daos) BasicSolidWasteUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbsrul *models.RefBasicSolidWasteUpdateLog) ([]models.RefSolidWasteChargeMonthlyPayments, error) {
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
	solidWastePaymentFindQuery := bson.M{
		"status":                 constants.SOLIDWASTEUSERCHARGEPAYMENTSTATUSCOMPLETED,
		"solidWasteUserChargeId": rbsrul.SolidWasteID,
		"completionDate":         bson.M{"$gte": sd, "$lte": ed},
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("shopRent query =>", solidWastePaymentFindQuery)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS).Find(ctx.CTX, solidWastePaymentFindQuery, nil)
	if err != nil {
		return nil, err
	}
	var solidWastePayments []models.RefSolidWasteChargeMonthlyPayments
	if err = cursor.All(context.TODO(), &solidWastePayments); err != nil {
		return nil, err
	}

	return solidWastePayments, nil
}
