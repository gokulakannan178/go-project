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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveCustomer :""
func (d *Daos) SaveCustomer(ctx *models.Context, Customer *models.Customer) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).InsertOne(ctx.CTX, Customer)
	if err != nil {
		return err
	}
	Customer.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}
func (d *Daos) UpsertCustomer(ctx *models.Context, Customer *models.Customer) error {

	d.Shared.BsonToJSONPrint(Customer)

	opts := options.Update().SetUpsert(true)

	query := bson.M{"mobile": Customer.Mobile}
	update := bson.M{"$set": Customer}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, query, update, opts)
	if err != nil {
		return err
	}

	return nil

}

//GetSingleCustomer : ""
func (d *Daos) GetSingleCustomer(ctx *models.Context, UniqueID string) (*models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Customers []models.RefCustomer
	var Customer *models.RefCustomer
	if err = cursor.All(ctx.CTX, &Customers); err != nil {
		return nil, err
	}
	if len(Customers) > 0 {
		Customer = &Customers[0]
	}
	return Customer, nil
}

//GetSingleGetUsingMobileNumber : ""
func (d *Daos) GetSingleGetUsingMobileNumber(ctx *models.Context, Mobile string) (*models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": Mobile}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Customers []models.RefCustomer
	var Customer *models.RefCustomer
	if err = cursor.All(ctx.CTX, &Customers); err != nil {
		return nil, err
	}
	if len(Customers) > 0 {
		Customer = &Customers[0]
	}
	return Customer, nil
}

//UpdateCustomer : ""
func (d *Daos) UpdateCustomer(ctx *models.Context, Customer *models.Customer) error {
	selector := bson.M{"uniqueId": Customer.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Customer, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCustomer : ""
func (d *Daos) FilterCustomer(ctx *models.Context, Customerfilter *models.CustomerFilter, pagination *models.Pagination) ([]models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Customerfilter != nil {
		if len(Customerfilter.Regex.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": Customerfilter.Regex.Name}})
		}
		if len(Customerfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Customerfilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Customer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Customers []models.RefCustomer
	if err = cursor.All(context.TODO(), &Customers); err != nil {
		return nil, err
	}
	return Customers, nil
}

//EnableCustomer :""
func (d *Daos) EnableCustomer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCustomer :""
func (d *Daos) DisableCustomer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCustomer :""
func (d *Daos) DeleteCustomer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleCustomerWithMobileNo : ""
func (d *Daos) GetSingleCustomerWithCondition(ctx *models.Context, key, value string) (*models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{key: value}})

	d.Shared.BsonToJSONPrintTag("customer query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefCustomer
	var user *models.RefCustomer
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}
