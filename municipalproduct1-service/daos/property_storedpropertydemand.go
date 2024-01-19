package daos

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveStoredPropertyDemand : ""
func (d *Daos) SaveStoredPropertyDemand(ctx *models.Context, spd *models.StoredPropertyDemand) error {
	if spd == nil {
		return errors.New("from Daos - demand is nil")
	}
	if len(spd.Fys) == 0 {
		return errors.New("from Daos - no fys to calculate demand")
	}
	fmt.Println(len(spd.Fys))
	for _, v := range spd.Fys {
		fmt.Println(v.FinancialYearId)
		opts := options.Update().SetUpsert(true)
		query := bson.M{"propertyId": v.PropertyID, "financialyearId": v.FinancialYearId}
		update := bson.M{"$set": v}
		res, err := ctx.DB.Collection(constants.COLLECTIONOSTOREDPROPERTYDEMANDFYS).UpdateOne(ctx.CTX, query, update, opts)
		if err != nil {
			return err
		}
		fmt.Println(res)

	}
	return nil
}

//SaveStoredPropertyDemandFy : ""
func (d *Daos) SaveStoredPropertyDemandFy(ctx *models.Context, fydv2 *models.FinancialYearDemandV2) error {
	if fydv2 == nil {
		return errors.New("from Daos - fydv2 demand is nil")
	}

	fmt.Println(fydv2.FinancialYearId)
	opts := options.Update().SetUpsert(true)
	query := bson.M{"propertyId": fydv2.PropertyID, "financialyearId": fydv2.FinancialYearId}
	update := bson.M{"$set": fydv2}
	res, err := ctx.DB.Collection(constants.COLLECTIONOSTOREDPROPERTYDEMANDFYS).UpdateOne(ctx.CTX, query, update, opts)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}
