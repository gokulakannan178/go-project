package daos

import (
	"errors"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//SavePropertyFyDemandLog :""
func (d *Daos) SavePropertyFyDemandLog(ctx *models.Context, propertyID string, fyd []models.FinancialYearDemand) error {
	query := bson.M{"propertyId": propertyID}
	count, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFYDEMANDLOG).CountDocuments(ctx.CTX, query)
	if err != nil {
		return errors.New("Error in geting Prev Fys - " + err.Error())
	}
	if count > 0 {
		delRes, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFYDEMANDLOG).DeleteMany(ctx.CTX, query)
		if err != nil {
			return errors.New("Error in Del Prev Fys - " + err.Error())
		}
		d.Shared.BsonToJSONPrintTag("Delete fy res ==>", delRes)
	}
	if len(fyd) == 0 {
		return nil
	}
	var temp []interface{}
	for _, v := range fyd {
		temp = append(temp, v)
	}
	insertRes, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFYDEMANDLOG).InsertMany(ctx.CTX, temp)
	d.Shared.BsonToJSONPrintTag("Insert fy res ==>", insertRes)

	return err
}

func (d *Daos) GetPropertyFyDemandLog(ctx *models.Context, propertyID string) ([]models.FinancialYearDemand, error) {
	query := bson.M{"propertyId": propertyID}
	//Find
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFYDEMANDLOG).Find(ctx.CTX, query, nil)
	if err != nil {
		return nil, errors.New("GetPropertyFyDemandLog Find - " + err.Error())
	}
	var financialYears []models.FinancialYearDemand

	if err := cursor.All(ctx.CTX, &financialYears); err != nil {
		return nil, errors.New("GetPropertyFyDemandLog All - " + err.Error())
	}
	// fmt.Println(len(financialYears))
	return financialYears, nil
}

// UpdatePropertyFYDemandLogPropertyID :""
func (d *Daos) UpdatePropertyFYDemandLogPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFYDEMANDLOG).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
