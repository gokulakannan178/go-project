package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AddCart : ""
func (d *Daos) AddCart(ctx *models.Context, cart *models.Cart) error {
	d.Shared.BsonToJSONPrint(cart)
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).InsertOne(ctx.CTX, cart)
	return err
}
func (d *Daos) GetSingleCartAndVendor(ctx *models.Context, customerId string, VendorId string) (*models.RefCart, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"customer.id": customerId, "customer.type": "Customer", "company.id": VendorId, "company.type": "Vendor"}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var carts []models.RefCart
	var cart *models.RefCart
	if err = cursor.All(ctx.CTX, &carts); err != nil {
		return nil, err
	}
	if len(carts) > 0 {
		cart = &carts[0]
	}
	return cart, nil
}

//UpsertCart :""
func (d *Daos) UpdateCart(ctx *models.Context, cart *models.Cart) error {
	//	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": cart.UniqueID}
	updateData := bson.M{"$set": cart}
	if _, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, updateQuery, updateData); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil
}

// GetSingleCart : ""
func (d *Daos) GetSingleCart(ctx *models.Context, UniqueID string) (*models.RefCart, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var carts []models.RefCart
	var cart *models.RefCart
	if err = cursor.All(ctx.CTX, &carts); err != nil {
		return nil, err
	}
	if len(carts) > 0 {
		cart = &carts[0]
	}
	return cart, nil
}

//GetSingleCartWithCustomerId
func (d *Daos) GetSingleCartWithCustomerId(ctx *models.Context, UniqueID string) (*models.RefCart, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"customer.id": UniqueID}})

	d.Shared.BsonToJSONPrintTag("cart query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var carts []models.RefCart
	var cart *models.RefCart
	if err = cursor.All(ctx.CTX, &carts); err != nil {
		return nil, err
	}
	if len(carts) > 0 {
		cart = &carts[0]
	}
	return cart, nil
}

// GetSingleCartWithInventoryId : ""
func (d *Daos) GetSingleCartWithInventoryId(ctx *models.Context, uniqueId string, inventoryid string) (*models.RefCart, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if uniqueId != "" {
		query = append(query, bson.M{"customer.id": uniqueId})
	}
	if uniqueId != "" {
		query = append(query, bson.M{"products.inventoryid": inventoryid})
	}
	//mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueId, "products.inventoryid": inventoryid}})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	d.Shared.BsonToJSONPrintTag("cart query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var carts []models.RefCart
	var cart *models.RefCart
	if err = cursor.All(ctx.CTX, &carts); err != nil {
		return nil, err
	}
	if len(carts) > 0 {
		cart = &carts[0]
	}
	return cart, nil
}

// UpdateCartItemQuanity : ""
func (d *Daos) UpdateCartItemQuanity(ctx *models.Context, cart *models.UpdateCart) error {
	selector := bson.M{"customer.id": cart.CustomerId, "products": bson.M{"$elemMatch": bson.M{"inventoryid": cart.InventoryID}}}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": bson.M{"products.$.quantity": cart.Quantity}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableCart : ""
func (d *Daos) EnableCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CARTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableCart : ""
func (d *Daos) DisableCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CARTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteCart : ""
func (d *Daos) DeleteCart(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CARTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCART).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCart : ""
func (d *Daos) FilterCart(ctx *models.Context, filter *models.CartFilter, pagination *models.Pagination) ([]models.RefCart, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		// if len(filter.UniqueID) > 0 {
		// 	query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		// }
		// //Regex Using searchBox Struct
		// if filter.SearchText.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchText.Name, Options: "xi"}})
		// }
		// if filter.SearchText.UniqueID != "" {
		// 	query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		// }
		// if filter.DateRange != nil {
		// 	//var sd,ed time.Time
		// 	if filter.DateRange.From != nil {
		// 		sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
		// 		ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
		// 		if filter.DateRange.To != nil {
		// 			ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
		// 		}
		// 		query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

		// 	}
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCART).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCART).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var cart []models.RefCart
	if err = cursor.All(context.TODO(), &cart); err != nil {
		return nil, err
	}
	return cart, nil
}
func (d *Daos) CalculationCart(ctx *models.Context, cart *models.Cart) error {
	for k, _ := range cart.Products {
		inventory, err := d.GetSingleInventory(ctx, cart.Products[k].InventoryID)
		if err != nil {
			return err
		}
		cart.Products[k].Price = inventory.Price.Selling
		cart.Products[k].Amount = inventory.Price.Selling * cart.Products[k].Quantity
		cart.SubTotal = cart.SubTotal + cart.Products[k].Amount
	}
	cart.TotalAmount = math.Round(cart.SubTotal)
	cart.RoundOff = cart.TotalAmount - cart.SubTotal
	return nil
}
