package daos

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveFPOINVENTORY : ""
func (d *Daos) SaveFPOInventory(ctx *models.Context, fpo *models.FPOInventory) error {
	d.Shared.BsonToJSONPrint(fpo)
	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).InsertOne(ctx.CTX, fpo)
	return err
}

// GetSingleFPOINVENTORY : ""
func (d *Daos) GetSingleFPOInventory(ctx *models.Context, UniqueID string) (*models.RefFPOINVENTORY, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var inventories []models.RefFPOINVENTORY
	var inventory *models.RefFPOINVENTORY
	if err = cursor.All(ctx.CTX, &inventories); err != nil {
		return nil, err
	}
	if len(inventories) > 0 {
		inventory = &inventories[0]
	}
	return inventory, nil
}

// UpdateFPOINVENTORY : ""
func (d *Daos) UpdateFPOInventory(ctx *models.Context, inventory *models.FPOInventory) error {
	selector := bson.M{"companyId": inventory.CompanyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": inventory}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableFPOINVENTORY : ""
func (d *Daos) EnableFPOInventory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FPOINVENTORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableFPOINVENTORY : ""
func (d *Daos) DisableFPOInventory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FPOINVENTORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteFPOINVENTORY : ""
func (d *Daos) DeleteFPOInventory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FPOINVENTORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// // FilterFPOINVENTORY : ""
// func (d *Daos) FilterFPOInventory(ctx *models.Context, filter *models.FPOINVENTORYFilter, pagination *models.Pagination) ([]models.RefFPOINVENTORY, error) {
// 	mainPipeline := []bson.M{}
// 	query := []bson.M{}
// 	if filter != nil {
// 		if len(filter.Status) > 0 {
// 			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
// 		}
// 	}
// 	//Adding $match from filter
// 	if len(query) > 0 {
// 		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
// 	}

// 	//Adding pagination if necessary
// 	if pagination != nil {
// 		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
// 		//Getting Total count
// 		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).CountDocuments(ctx.CTX, func() bson.M {
// 			if query != nil {
// 				if len(query) > 0 {
// 					return bson.M{"$and": query}
// 				}
// 			}
// 			return bson.M{}
// 		}())
// 		if err != nil {
// 			log.Println("Error in geting pagination")
// 		}
// 		fmt.Println("count", totalCount)
// 		pagination.Count = int(totalCount)
// 		d.Shared.PaginationData(pagination)
// 	}

// 	//Aggregation
// 	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var fpo []models.RefFPOINVENTORY
// 	if err = cursor.All(context.TODO(), &fpo); err != nil {
// 		return nil, err
// 	}
// 	return fpo, nil
// }

// FPOInventoryQuantityUpdate : ""
func (d *Daos) FPOInventoryQuantityUpdate(ctx *models.Context, inventory *models.FPOInventory) error {
	selector := bson.M{"companyId": inventory.CompanyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": inventory.Quantity}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// FPOInventoryPriceUpdate : ""
func (d *Daos) FPOInventoryPriceUpdate(ctx *models.Context, inventory *models.FPOInventory) error {
	selector := bson.M{"companyId": inventory.CompanyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"sellingPrice": inventory.Sellingprice, "buyingPrice": inventory.BuyingPrice}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleFPOINVENTORY : ""
func (d *Daos) GetSingleFPOInventoryWithCompalyID(ctx *models.Context, UniqueID string) (*models.RefFPOINVENTORY, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"companyId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("GetSingleFPOInventoryWithCompalyID query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var inventories []models.RefFPOINVENTORY
	var inventory *models.RefFPOINVENTORY
	if err = cursor.All(ctx.CTX, &inventories); err != nil {
		return nil, err
	}
	if len(inventories) > 0 {
		inventory = &inventories[0]
	}
	return inventory, nil
}

func (d *Daos) GetSingleULBInventoryWithCompalyID(ctx *models.Context, UniqueID string) (*models.ULBInventory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"companyId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("GetSingleULBInventoryWithCompalyID query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULBINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var inventories []models.ULBInventory
	var inventory *models.ULBInventory
	if err = cursor.All(ctx.CTX, &inventories); err != nil {
		return nil, err
	}
	if len(inventories) > 0 {
		inventory = &inventories[0]
	}
	return inventory, nil
}

// UpdateFPOInventoryDeliverSale : ""
func (d *Daos) UpdateFPOInventoryDeliverSale(ctx *models.Context, inventory *models.RefFPOINVENTORY) error {
	selector := bson.M{"companyId": inventory.CompanyID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"quantity": inventory.Quantity}}
	d.Shared.BsonToJSONPrintTag("UpdateFPOInventoryDeliverSale query =>", selector)
	d.Shared.BsonToJSONPrintTag("UpdateFPOInventoryDeliverSale data =>", data)

	_, err := ctx.DB.Collection(constants.COLLECTIONFPOINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
