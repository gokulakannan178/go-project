package daos

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//VendorDashboardOrder : ""
func (d *Daos) VendorDashboardOrder(ctx *models.Context, filter *models.VendorDashBoardFilter) (*models.VendorDashBoardOrder, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{

		"noOfOrders": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{"Completed", "Pending"}}}},
			bson.M{"$count": "noOfOrders"},
		},
		"pendingAmounts": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{"Completed", "Pending"}}}},
			bson.M{"$group": bson.M{"_id": nil, "amount": bson.M{"$sum": "$payment.pendingAmount"}}},
		},
		"soldItems": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{"Completed", "Pending"}}}},
			bson.M{"$group": bson.M{"_id": nil, "soldItems": bson.M{"$sum": "$Items.quantity"}}},
		},
		"totalSales": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{"Completed", "Pending"}}}},
			bson.M{"$group": bson.M{"_id": nil, "totalSales": bson.M{"$sum": "$payment.amount"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"noOfOrders":     bson.M{"$arrayElemAt": []interface{}{"$noOfOrders", 0}},
		"pendingAmounts": bson.M{"$arrayElemAt": []interface{}{"$pendingAmounts", 0}},
		"soldItems":      bson.M{"$arrayElemAt": []interface{}{"$soldItems", 0}},
		"totalSales":     bson.M{"$arrayElemAt": []interface{}{"$totalSales", 0}},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"pendingAmounts": "$pendingAmounts.amount",
		"totalSales": "$totalSales.amount",
		"noOfOrders": "$noOfOrders.noOfOrders",
		"soldItems":  "$soldItems.quantity",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("vendordashboardorder query =>", mainPipeline)
	//	var data *models.VendorDashBoardOrder
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORDER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vendordashbords []models.VendorDashBoardOrder
	var vendordashbord *models.VendorDashBoardOrder
	if err = cursor.All(ctx.CTX, &vendordashbords); err != nil {
		return nil, err
	}
	if len(vendordashbords) > 0 {
		vendordashbord = &vendordashbords[0]
	}
	return vendordashbord, nil
}
func (d *Daos) VendorDashboardProduct(ctx *models.Context, filter *models.VendorDashBoardFilter) (*models.VendorDashBoardProduct, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{

		"products": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{"Active"}}}},
			bson.M{"$count": "products"},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"products": bson.M{"$arrayElemAt": []interface{}{"$products", 0}},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"products": "$products.products",
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("vendordashboardproduct query =>", mainPipeline)
	//var data *models.VendorDashBoardProduct
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vendordashbords []models.VendorDashBoardProduct
	var vendordashbord *models.VendorDashBoardProduct
	if err = cursor.All(ctx.CTX, &vendordashbords); err != nil {
		return nil, err
	}
	if len(vendordashbords) > 0 {
		vendordashbord = &vendordashbords[0]
	}
	return vendordashbord, nil
}
func (d *Daos) VendorDashboardLowStock(ctx *models.Context, filter *models.VendorDashBoardFilter) (*models.VendorDashBoardLowStock, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{
		"$expr": bson.M{"$and": bson.M{"$lte": []string{"$quantity", "$lowStock"}}}}},
		bson.M{"$group": bson.M{"_id": nil, "lowStocks": bson.M{"$sum": 1}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("vendordashboardlowstock query =>", mainPipeline)
	//var data *models.VendorDashBoardLowStock
	cursor, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vendordashbords []models.VendorDashBoardLowStock
	var vendordashbord *models.VendorDashBoardLowStock
	if err = cursor.All(ctx.CTX, &vendordashbords); err != nil {
		return nil, err
	}
	if len(vendordashbords) > 0 {
		vendordashbord = &vendordashbords[0]
	}
	return vendordashbord, nil
}
