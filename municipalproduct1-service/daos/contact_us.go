package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveUser :""
func (d *Daos) SaveContact(ctx *models.Context, contact *models.ContactUs) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).InsertOne(ctx.CTX, contact)
	return err
}

//GetSingleUser : ""
func (d *Daos) GetSingleContact(ctx *models.Context, UniqueID string) (*models.RefContactUs, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("contactUs query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var contactUss []models.RefContactUs
	var contactUs *models.RefContactUs
	if err = cursor.All(ctx.CTX, &contactUss); err != nil {
		return nil, err
	}
	if len(contactUss) > 0 {
		contactUs = &contactUss[0]
	}
	return contactUs, nil
}

//UpdateUser : ""
func (d *Daos) UpdateContact(ctx *models.Context, contact *models.ContactUs) error {
	selector := bson.M{"uniqueId": contact.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": contact}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterContactUs : ""
func (d *Daos) FilterContactUs(ctx *models.Context, contactfilter *models.FilterContactUs, pagination *models.Pagination) ([]models.RefContactUs, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if contactfilter != nil {
		if len(contactfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": contactfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("contactus query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var contactUss []models.RefContactUs
	if err = cursor.All(context.TODO(), &contactUss); err != nil {
		return nil, err
	}
	return contactUss, nil
}

//EnableUser :""
//func (d *Daos) EnableContact(ctx *models.Context, Name string) error {
func (d *Daos) EnableContact(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CONTACTUSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableUser :""
func (d *Daos) DisableContact(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CONTACTUSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteUser :""
func (d *Daos) DeleteContact(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CONTACTUSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTACTUS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
