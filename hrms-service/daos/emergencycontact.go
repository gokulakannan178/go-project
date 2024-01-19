package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveEmergencyContact :""
func (d *Daos) SaveEmergencyContact(ctx *models.Context, emergencyContact *models.EmergencyContact) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).InsertOne(ctx.CTX, emergencyContact)
	if err != nil {
		return err
	}
	emergencyContact.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleEmergencyContact : ""
func (d *Daos) GetSingleEmergencyContact(ctx *models.Context, uniqueID string) (*models.RefEmergencyContact, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmergencyContactCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var emergencyContacts []models.RefEmergencyContact
	var emergencyContact *models.RefEmergencyContact
	if err = cursor.All(ctx.CTX, &emergencyContacts); err != nil {
		return nil, err
	}
	if len(emergencyContacts) > 0 {
		emergencyContact = &emergencyContacts[0]
	}
	return emergencyContact, nil
}

//UpdateEmergencyContact : ""
func (d *Daos) UpdateEmergencyContact(ctx *models.Context, emergencyContact *models.EmergencyContact) error {
	selector := bson.M{"uniqueId": emergencyContact.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": emergencyContact}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableEmergencyContact :""
func (d *Daos) EnableEmergencyContact(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMERGENCYCONTACTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEmergencyContact :""
func (d *Daos) DisableEmergencyContact(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMERGENCYCONTACTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEmergencyContact :""
func (d *Daos) DeleteEmergencyContact(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMERGENCYCONTACTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterEmergencyContact : ""
func (d *Daos) FilterEmergencyContact(ctx *models.Context, filter *models.FilterEmergencyContact, pagination *models.Pagination) ([]models.RefEmergencyContact, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		//Regex
		if filter.Regex.FullName != "" {
			query = append(query, bson.M{"fullName": primitive.Regex{Pattern: filter.Regex.FullName, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEmergencyContactCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("EmergencyContact query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMERGENCYCONTACT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmergencyContacts []models.RefEmergencyContact
	if err = cursor.All(context.TODO(), &EmergencyContacts); err != nil {
		return nil, err
	}
	return EmergencyContacts, nil
}
