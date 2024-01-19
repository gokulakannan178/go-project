package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveCustomer : ""
func (d *Daos) SaveCustomer(ctx *models.Context, Customer *models.Customer) error {
	d.Shared.BsonToJSONPrint(Customer)
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).InsertOne(ctx.CTX, Customer)
	return err
}

// GetSingleCustomer : ""
func (d *Daos) GetSingleCustomer(ctx *models.Context, UniqueID string) (*models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "mobile", "mobile", "ref.chairman", "ref.chairman")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefCustomer
	var tower *models.RefCustomer
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	fmt.Println(tower)
	return tower, nil

}

// UpdateCustomer : ""
func (d *Daos) UpdateCustomer(ctx *models.Context, business *models.Customer) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableCustomer : ""
func (d *Daos) EnableCustomer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMERSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableCustomer : ""
func (d *Daos) DisableCustomer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMERSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteCustomer : ""
func (d *Daos) DeleteCustomer(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMERSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCustomer : ""
func (d *Daos) FilterCustomer(ctx *models.Context, filter *models.CustomerFilter, pagination *models.Pagination) ([]models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	//Regex
	if filter.Regex.Name != "" {
		query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
	}
	if filter.Regex.Email != "" {
		query = append(query, bson.M{"email": primitive.Regex{Pattern: filter.Regex.Email, Options: "xi"}})
	}
	if filter.Regex.Mobile != "" {
		query = append(query, bson.M{"primaryContact.ph": primitive.Regex{Pattern: filter.Regex.Mobile, Options: "xi"}})
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Customer []models.RefCustomer
	if err = cursor.All(context.TODO(), &Customer); err != nil {
		return nil, err
	}
	return Customer, nil
}
func (d *Daos) GetSingleCustomerwithmobileno(ctx *models.Context, mobile string) (*models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": mobile}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "mobile", "mobile", "ref.chairman", "ref.chairman")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFPO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefCustomer
	var tower *models.RefCustomer
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	fmt.Println(tower)
	return tower, nil

}
func (d *Daos) GetSingleCustomerWithMobileNo(ctx *models.Context, mobileno string) (*models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": mobileno}})
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFPO, "companyId", "uniqueId", "ref.fpo", "ref.fpo")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
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

// GetSingleCustomer : ""
func (d *Daos) GetSingleCustomerwithprofile(ctx *models.Context, Profile string) (*models.RefCustomer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"profile": Profile}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "mobile", "mobile", "ref.chairman", "ref.chairman")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefCustomer
	var tower *models.RefCustomer
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	fmt.Println(tower)
	return tower, nil

}

// UpdateCustomer : ""
func (d *Daos) UpdateCustomerwithprofile(ctx *models.Context, Profile *models.Customer) error {
	selector := bson.M{"profile": Profile.Profile}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": Profile}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
