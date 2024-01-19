package daos

import (
	"errors"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) VerifyMobileTowerRegistrationPayment(ctx *models.Context, action *models.MakeMobileTowerPaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.MobileTowerPaymentsAction,
			"status":       constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

func (d *Daos) NotVerifyMobileTowerRegistrationPayment(ctx *models.Context, action *models.MakeMobileTowerPaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.MobileTowerPaymentsAction,
			"status":       constants.MOBILETOWERPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSNOTVERIFIED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}

func (d *Daos) RejectMobileTowerRegistrationPayment(ctx *models.Context, action *models.MakeMobileTowerPaymentsAction) error {
	query := bson.M{"tnxId": action.TnxID}
	paymentData := bson.M{
		"$set": bson.M{
			"verifiedInfo": action.MobileTowerPaymentsAction,
			"status":       constants.MOBILETOWERPAYMENRSTATUSREJECTED,
		},
	}
	res, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTS).UpdateOne(ctx.CTX, query, paymentData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment resp - ", res)

	paymentFyData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSFY).UpdateMany(ctx.CTX, query, paymentFyData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment fys resp - ", res)

	paymentBasicData := bson.M{
		"$set": bson.M{
			"status": constants.MOBILETOWERPAYMENRSTATUSREJECTED,
		},
	}
	res, err = ctx.DB.Collection(constants.COLLECTIONMOBILETOWERPAYMENTSBASIC).UpdateMany(ctx.CTX, query, paymentBasicData)
	if err != nil {
		return errors.New("Error in updating payment - " + err.Error())
	}
	d.Shared.BsonToJSONPrintTag("Mobile tower payment basic resp - ", res)
	return nil
}
