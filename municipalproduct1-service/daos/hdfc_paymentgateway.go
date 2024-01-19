package daos

import (
	"errors"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpsertHDFCPaymentGateway : ""
func (d *Daos) UpsertHDFCPaymentGateway(ctx *models.Context, hdfc *models.HDFCPaymentGateway) error {
	d.Shared.BsonToJSONPrint(hdfc)
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"isDefault": true}
	updateData := bson.M{"$set": hdfc}
	if _, err := ctx.DB.Collection(constants.COLLECTIONHDFCPAYMENTGATEWAY).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil

}

func (d *Daos) GetSingleDefaultHDFCPaymentGateway(ctx *models.Context) (*models.RefHDFCPaymentGateway, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHDFCPAYMENTGATEWAY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payments []models.RefHDFCPaymentGateway
	var payment *models.RefHDFCPaymentGateway
	if err = cursor.All(ctx.CTX, &payments); err != nil {
		return nil, err
	}
	if len(payments) > 0 {
		payment = &payments[0]
	}
	return payment, err
}

// GetSingleMerchantIDHDFCPaymentGateway : ""
func (d *Daos) GetSingleMerchantIDHDFCPaymentGateway(ctx *models.Context) (string, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHDFCPAYMENTGATEWAY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return "", err
	}
	var payments []models.RefHDFCPaymentGateway
	var payment *models.RefHDFCPaymentGateway
	if err = cursor.All(ctx.CTX, &payments); err != nil {
		return "", err
	}
	if len(payments) > 0 {
		payment = &payments[0]
	}
	return payment.MerchantID, err
}
