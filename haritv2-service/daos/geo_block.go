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
)

//SaveBlock :""
func (d *Daos) SaveBlock(ctx *models.Context, block *models.Block) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).InsertOne(ctx.CTX, block)
	return err
}

//GetSingleBlock : ""
func (d *Daos) GetSingleBlock(ctx *models.Context, uniqueID string) (*models.RefBlock, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blocks []models.RefBlock
	var block *models.RefBlock
	if err = cursor.All(ctx.CTX, &blocks); err != nil {
		return nil, err
	}
	if len(blocks) > 0 {
		block = &blocks[0]
	}
	return block, nil
}

//UpdateBlock : ""
func (d *Daos) UpdateBlock(ctx *models.Context, block *models.Block) error {
	selector := bson.M{"uniqueId": block.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": block, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterBlock : ""
func (d *Daos) FilterBlock(ctx *models.Context, blockfilter *models.BlockFilter, pagination *models.Pagination) ([]models.RefBlock, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if blockfilter != nil {
		if len(blockfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": blockfilter.UniqueID}})
		}
		if len(blockfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": blockfilter.Status}})
		}
		if len(blockfilter.DistrictCodes) > 0 {
			query = append(query, bson.M{"districtCode": bson.M{"$in": blockfilter.DistrictCodes}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("block query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blocks []models.RefBlock
	if err = cursor.All(context.TODO(), &blocks); err != nil {
		return nil, err
	}
	return blocks, nil
}

//EnableBlock :""
func (d *Daos) EnableBlock(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBlock :""
func (d *Daos) DisableBlock(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBlock :""
func (d *Daos) DeleteBlock(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
