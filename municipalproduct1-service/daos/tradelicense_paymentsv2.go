package daos

import (
	"errors"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//UpdateExpiryDate :""
func (d *Daos) UpdateExpiryDate(ctx *models.Context, uniqueID string, expiryDate *time.Time, fromDate *time.Time, status string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"licenseExpiryDate": expiryDate, "status": status, "licenseDate": fromDate}}
	d.Shared.BsonToJSONPrintTag("query==>", query)
	d.Shared.BsonToJSONPrintTag("update==>", update)
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return nil
}

// // GetTradeLicenseCurrentFinancialYearMarketYear
// func (d *Daos) GetTradeLicenseCurrentFinancialYearMarketYear(ctx *models.Context, fromDate *time.Time, toDate *time.Time) error {
// 	return nil
// }
