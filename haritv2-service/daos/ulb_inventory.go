package daos

import (
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// UpdateULBInventoryDeliverSale : ""
func (d *Daos) UpdateULBInventoryDeliverSale(ctx *models.Context, inventory *models.ULBInventory) error {
	selector := bson.M{"companyId": inventory.CompanyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"quantity": inventory.Quantity}}
	d.Shared.BsonToJSONPrintTag("UpdateULBInventoryDeliverSale query =>", selector)
	d.Shared.BsonToJSONPrintTag("UpdateULBInventoryDeliverSale data =>", data)

	_, err := ctx.DB.Collection(constants.COLLECTIONULBINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) UpdateUlbInventoryDeliverSale(ctx *models.Context, inventory *models.RefULBInventory) error {
	selector := bson.M{"companyId": inventory.CompanyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"quantity": inventory.Quantity}}
	d.Shared.BsonToJSONPrintTag("UpdateULBInventoryDeliverSale query =>", selector)
	d.Shared.BsonToJSONPrintTag("UpdateULBInventoryDeliverSale data =>", data)

	_, err := ctx.DB.Collection(constants.COLLECTIONULBINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
