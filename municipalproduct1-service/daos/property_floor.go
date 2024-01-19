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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePropertyFloor :""
func (d *Daos) SavePropertyFloor(ctx *models.Context, propertyFloor *models.PropertyFloor) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).InsertOne(ctx.CTX, propertyFloor)
	return err
}

//SavePropertyFloor :""
func (d *Daos) SavePropertyFloorV2(ctx *models.Context, db *mongo.Database, sc *mongo.SessionContext, propertyFloor *models.PropertyFloor) error {
	_, err := db.Collection(constants.COLLECTIONPROPERTYFLOOR).InsertOne(*sc, propertyFloor)
	return err
}

//SavePropertyFloor :""
func (d *Daos) SavePropertyFloorV3(db *mongo.Database, sc mongo.SessionContext, propertyFloor *models.PropertyFloor) error {
	_, err := db.Collection(constants.COLLECTIONPROPERTYFLOOR).InsertOne(sc, propertyFloor)
	return err
}

//SavePropertyFloors :""
func (d *Daos) SavePropertyFloors(ctx *models.Context, propertyFloors []models.PropertyFloor) error {
	insertdata := []interface{}{}
	for _, v := range propertyFloors {
		insertdata = append(insertdata, v)
	}
	result, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).InsertMany(ctx.SC, insertdata)
	fmt.Println("insert result =>", result)
	return err
}

//SavePropertyFloors :""
func (d *Daos) SavePropertyFloorsV2(ctx *models.Context, db *mongo.Database, sc *mongo.SessionContext, propertyFloors []models.PropertyFloor) error {
	insertdata := []interface{}{}
	for _, v := range propertyFloors {
		insertdata = append(insertdata, v)
	}
	result, err := db.Collection(constants.COLLECTIONPROPERTYFLOOR).InsertMany(ctx.CTX, insertdata)
	fmt.Println("insert result =>", result)
	return err
}

//SavePropertyFloors :""
func (d *Daos) SavePropertyFloorsV3(db *mongo.Database, sc mongo.SessionContext, propertyFloors []models.PropertyFloor) error {
	insertdata := []interface{}{}
	for _, v := range propertyFloors {
		insertdata = append(insertdata, v)
	}
	result, err := db.Collection(constants.COLLECTIONPROPERTYFLOOR).InsertMany(sc, insertdata)
	fmt.Println("insert result =>", result)
	return err
}

//GetSinglePropertyFloor : ""
func (d *Daos) GetSinglePropertyFloor(ctx *models.Context, UniqueID string) (*models.RefPropertyFloor, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyFloors []models.RefPropertyFloor
	var propertyFloor *models.RefPropertyFloor
	if err = cursor.All(ctx.CTX, &propertyFloors); err != nil {
		return nil, err
	}
	if len(propertyFloors) > 0 {
		propertyFloor = &propertyFloors[0]
	}
	return propertyFloor, nil
}

//UpdatePropertyFloor : ""
func (d *Daos) UpdatePropertyFloor(ctx *models.Context, propertyFloor *models.PropertyFloor) error {
	selector := bson.M{"uniqueId": propertyFloor.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyFloor, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//UpdatePropertyFloorV2 : ""
func (d *Daos) UpdatePropertyFloorV2(ctx *models.Context, propertyFloors []models.PropertyFloor) error {
	floorIDs := []string{}
	propertyIDs := []string{}
	if len(propertyFloors) > 0 {
		for _, v := range propertyFloors {
			floorIDs = append(floorIDs, v.UniqueID)
			propertyIDs = append(propertyIDs, v.PropertyID)
			selector := bson.M{"uniqueId": v.UniqueID}
			opts := options.Update().SetUpsert(true)
			data := bson.M{"$set": v}
			if _, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).UpdateOne(ctx.CTX, selector, data, opts); err != nil {
				return errors.New("Error in updating floor : " + "unique ID => " + v.UniqueID + " floor no => " + v.No + " - " + err.Error())
			}
		}
		if len(propertyIDs) > 0 {
			deleteOpts := options.Update().SetUpsert(false)
			deleteSelector := bson.M{"uniqueId": bson.M{"$nin": floorIDs}, "propertyId": bson.M{"$in": propertyIDs}}
			d.Shared.BsonToJSONPrintTag("delete Floor query =>", deleteSelector)

			deleteData := bson.M{"$set": bson.M{"status": constants.PROPERTYFLOORSTATUSDELETED}}
			if _, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).UpdateMany(ctx.CTX, deleteSelector, deleteData, deleteOpts); err != nil {
				return errors.New("Error in updating deleted floor - " + err.Error())
			}
		}
	}
	return nil
}

//FilterPropertyFloor : ""
func (d *Daos) FilterPropertyFloor(ctx *models.Context, propertyFloorfilter *models.PropertyFloorFilter, pagination *models.Pagination) ([]models.RefPropertyFloor, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyFloorfilter != nil {

		if len(propertyFloorfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyFloorfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("propertyFloor query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyFloors []models.RefPropertyFloor
	if err = cursor.All(context.TODO(), &propertyFloors); err != nil {
		return nil, err
	}
	return propertyFloors, nil
}

//EnablePropertyFloor :""
func (d *Daos) EnablePropertyFloor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFLOORSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyFloor :""
func (d *Daos) DisablePropertyFloor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFLOORSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyFloor :""
func (d *Daos) DeletePropertyFloor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYFLOORSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) GetFloorsOfProperty(ctx *models.Context, propertyID string) ([]models.RefPropertyFloor, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$in": []string{constants.PROPERTYFLOORSTATUSACTIVE}}})
	query = append(query, bson.M{"propertyId": bson.M{"$in": []string{propertyID}}})
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFLOORTYPE, "no", "uniqueId", "ref.floorNo", "ref.floorNo")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONSTRUCTIONTYPE, "constructionType", "uniqueId", "ref.constructionType", "ref.constructionType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSAGETYPE, "usageType", "uniqueId", "ref.usageType", "ref.usageType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("propertyFloor query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyFloors []models.RefPropertyFloor
	if err = cursor.All(context.TODO(), &propertyFloors); err != nil {
		return nil, err
	}
	return propertyFloors, nil
}

// UpdateFloorPropertyID :""
func (d *Daos) UpdateFloorPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYFLOOR).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
