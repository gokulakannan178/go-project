package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveULB :""
func (d *Daos) SaveULB(ctx *models.Context, ULB *models.ULB) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).InsertOne(ctx.CTX, ULB)
	return err
}

//GetSingleULB : ""
func (d *Daos) GetSingleULB(ctx *models.Context, UniqueID string) (*models.RefULB, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var ulbs []models.RefULB
	var ULB *models.RefULB
	if err = cursor.All(ctx.CTX, &ulbs); err != nil {
		return nil, err
	}
	if len(ulbs) > 0 {
		ULB = &ulbs[0]
	}
	return ULB, nil
}

//UpdateULB : ""
func (d *Daos) UpdateULB(ctx *models.Context, ULB *models.ULB) error {
	selector := bson.M{"uniqueId": ULB.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ULB, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterULB : ""
func (d *Daos) FilterULB(ctx *models.Context, ulbfilter *models.ULBFilter, pagination *models.Pagination) ([]models.RefULB, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if ulbfilter != nil {

		if len(ulbfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": ulbfilter.Status}})
		}
		if len(ulbfilter.State) > 0 {
			query = append(query, bson.M{"address.stateCode": bson.M{"$in": ulbfilter.State}})
		}
		if len(ulbfilter.District) > 0 {
			query = append(query, bson.M{"address.districtCode": bson.M{"$in": ulbfilter.District}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONULB).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("ULB query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ulbs []models.RefULB
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	return ulbs, nil
}

//EnableULB :""
func (d *Daos) EnableULB(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ULBSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableULB :""
func (d *Daos) DisableULB(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ULBSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteULB :""
func (d *Daos) DeleteULB(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ULBSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
