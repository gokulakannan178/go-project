package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveInventory : ""
func (d *Daos) SaveInventory(ctx *models.Context, block *models.Inventory) error {
	d.Shared.BsonToJSONPrint(block)
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).InsertOne(ctx.CTX, block)
	return err
}

// UpsertProductVarients : ""
func (d *Daos) UpsertInventory(ctx *models.Context, inventory *models.Inventory) error {
	d.Shared.BsonToJSONPrint(inventory)

	opts := options.Update().SetUpsert(true)

	query := bson.M{"productVarientId": inventory.ProductVarientID, "vendorId": inventory.VendorID}
	update := bson.M{"$set": inventory}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, query, update, opts)
	if err != nil {
		return err
	}

	return nil
}

// GetSingleInventory : ""
func (d *Daos) GetSingleInventory(ctx *models.Context, UniqueID string) (*models.RefInventory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCT, "productId", "uniqueId", "ref.product", "ref.product")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "ref.product.categoryId", "uniqueId", "ref.product.ref.category", "ref.product.ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "ref.product.subCategoryId", "uniqueId", "ref.product.ref.subCategory", "ref.product.ref.subCategory")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefInventory
	var tower *models.RefInventory
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateInventory : ""
func (d *Daos) UpdateInventory(ctx *models.Context, crop *models.Inventory) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// UpdateInventoryQuantityDetails : ""
func (d *Daos) UpdateInventoryQuantityDetails(ctx *models.Context, crop *models.Inventory) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": bson.M{"quantity": crop.Quantity, "price": crop.Price}}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableInventory : ""
func (d *Daos) EnableInventory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.INVENTORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableInventory : ""
func (d *Daos) DisableInventory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.INVENTORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteInventory : ""
func (d *Daos) DeleteInventory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.INVENTORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterInventory : ""
func (d *Daos) FilterInventory(ctx *models.Context, filter *models.InventoryFilter, pagination *models.Pagination) ([]models.RefInventory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.ProductID) > 0 {
			query = append(query, bson.M{"productId": bson.M{"$in": filter.ProductID}})
		}
		if len(filter.VendorID) > 0 {
			query = append(query, bson.M{"vendorId": bson.M{"$in": filter.VendorID}})
		}

		if len(filter.PVCombination) > 0 {
			query = append(query, bson.M{"pVCombination": bson.M{"$in": filter.PVCombination}})
		}
		if filter.QuantityRange != nil {
			query = append(query, bson.M{"quantity": bson.M{"$gte": filter.QuantityRange.From, "$lte": filter.QuantityRange.To}})

			// query = append(query, bson.M{"$gte": []interface{}{"quantity", filter.QuantityRange.From}})
			// query = append(query, bson.M{"$lte": []interface{}{"quantity", filter.QuantityRange.To}})
		}
		// query = append(query, bson.M{"quantity": bson.M{"$gte": from, "$lte": to}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCT, "productId", "uniqueId", "ref.product", "ref.product")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "ref.product.categoryId", "uniqueId", "ref.product.ref.category", "ref.product.ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "ref.product.subCategoryId", "uniqueId", "ref.product.ref.subCategory", "ref.product.ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockCrop []models.RefInventory
	if err = cursor.All(context.TODO(), &blockCrop); err != nil {
		return nil, err
	}
	return blockCrop, nil
}

// imageInventory : ""
func (d *Daos) ImageInventory(ctx *models.Context, crop *models.Inventory) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{}
	//data := bson.M{"$set": bson.M{"image": crop.Image}}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// imagesInventory : ""
func (d *Daos) ImagesInventory(ctx *models.Context, crop *models.Inventory) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{}
	//	data := bson.M{"$set": bson.M{"images": crop.Images}}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) GetbyBarcodeAndVendor(ctx *models.Context, UniqueID string, vendor string) (*models.RefInventory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"barcode": UniqueID, "vendorId": vendor}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	//	Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCT, "productId", "uniqueId", "ref.product", "ref.product")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "ref.product.categoryId", "uniqueId", "ref.product.ref.category", "ref.product.ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "ref.product.subCategoryId", "uniqueId", "ref.product.ref.subCategory", "ref.product.ref.subCategory")...)
	//	Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var barcode []models.RefInventory
	var vendors *models.RefInventory
	if err = cursor.All(ctx.CTX, &barcode); err != nil {
		return nil, err
	}
	if len(barcode) > 0 {
		vendors = &barcode[0]
	}
	return vendors, nil
}
func (d *Daos) ChkUniqueness(ctx *models.Context, UniqueID string, vendor string) (*models.RefInventory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"barcode": bson.M{"$eq": UniqueID}})
	query = append(query, bson.M{"vendorId": bson.M{"$eq": vendor}})

	//	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"barcode": UniqueID, "vendorId": vendor}})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	totalCount, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).CountDocuments(ctx.CTX, func() bson.M {
		if query != nil {
			if len(query) > 0 {
				return bson.M{"$and": query}
			}
		}
		return bson.M{}
	}())
	if err != nil {
		log.Println("Error in geting pagination")
	}
	fmt.Println("count", totalCount)
	if err != nil {
		log.Println("Error in geting pagination")
	}
	if totalCount > 1 {
		return nil, errors.New("Duplicate barcode founded")
	}

	d.Shared.BsonToJSONPrintTag("inventory query =>", mainPipeline)

	//	Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCT, "productId", "uniqueId", "ref.product", "ref.product")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "ref.product.categoryId", "uniqueId", "ref.product.ref.category", "ref.product.ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBCATEGORY, "ref.product.subCategoryId", "uniqueId", "ref.product.ref.subCategory", "ref.product.ref.subCategory")...)
	//	Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var barcode []models.RefInventory
	var vendors *models.RefInventory
	if err = cursor.All(ctx.CTX, &barcode); err != nil {
		return nil, err
	}
	if len(barcode) > 0 {
		vendors = &barcode[0]
	}
	return vendors, nil
}

// UpdateInventoryQuantityForSale : ""
func (d *Daos) UpdateInventoryQuantityForSale(ctx *models.Context, inventoryId string, quantity float64) error {
	selector := bson.M{"uniqueId": inventoryId}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$inc": bson.M{"quantity": (-1 * quantity)}}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
