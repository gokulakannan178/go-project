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

// SaveMobileTowerRegistrationRateMaster : ""
func (d *Daos) SaveMobileTowerRegistrationRateMaster(ctx *models.Context, mobile *models.MobileTowerRegistrationRateMaster) error {
	d.Shared.BsonToJSONPrint(mobile)
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).InsertOne(ctx.CTX, mobile)
	return err
}

// GetSingleMobileTowerRegistrationRateMaster : ""
func (d *Daos) GetSingleMobileTowerRegistrationRateMaster(ctx *models.Context, UniqueID string) (*models.MobileTowerRegistrationRateMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.MobileTowerRegistrationRateMaster
	var tower *models.MobileTowerRegistrationRateMaster
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateMobileTowerRegistrationRateMaster: ""
func (d *Daos) UpdateMobileTowerRegistrationRateMaster(ctx *models.Context, mobile *models.MobileTowerRegistrationRateMaster) error {
	selector := bson.M{"uniqueId": mobile.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": mobile}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMobileTowerRegistrationRateMaster : ""
func (d *Daos) EnableMobileTowerRegistrationRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERREGISTRATIONRATEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableMobileTowerRegistrationRateMaster: ""
func (d *Daos) DisableMobileTowerRegistrationRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": "Disable"}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteMobileTowerRegistrationRateMaster : ""
func (d *Daos) DeleteMobileTowerRegistrationRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MOBILETOWERTAXSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMobileTowerRegistrationRateMaster : ""
func (d *Daos) FilterMobileTowerRegistrationRateMaster(ctx *models.Context, filter *models.MobileTowerRegistrationRateMasterFilter, pagination *models.Pagination) ([]models.RefMobileTowerRegistrationRateMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefMobileTowerRegistrationRateMaster
	if err = cursor.All(context.TODO(), &towers); err != nil {
		return nil, err
	}
	return towers, nil
}
