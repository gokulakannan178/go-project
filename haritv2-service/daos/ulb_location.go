package daos

import (
	"context"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//UpdateULB : ""
func (d *Daos) UpdateUlbLocation(ctx *models.Context, ulbloc *models.UpdateLocation) error {
	updated := new(models.Updated)
	t := time.Now()
	updated.On = &t
	updated.By = ulbloc.By
	updated.ByType = ulbloc.ByType
	updated.Scenario = "addingLocation"
	selector := bson.M{"uniqueId": ulbloc.UniqueID}
	updateInterface := []bson.M{}
	updateInterface = append(updateInterface, bson.M{"$set": bson.M{"address.location": ulbloc.Location, "isLocationUpdated": true, "updated": updated}})
	d.Shared.BsonToJSONPrint(updateInterface)
	_, err := ctx.DB.Collection(constants.COLLECTIONULB).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) GutULBNotUpdatedLocation(ctx *models.Context) ([]models.RefULB, error) {

	query := bson.M{"status": "Active", "isLocationUpdated": bson.M{"$ne": true}}
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Find(ctx.CTX, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var ulbs []models.RefULB
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	return ulbs, nil
}
func (d *Daos) GutULBNotUpdatedProfile(ctx *models.Context) ([]models.RefULB, error) {

	query := bson.M{"status": "Active", "isProfileUpdated": bson.M{"$ne": true}}
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Find(ctx.CTX, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var ulbs []models.RefULB
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	return ulbs, nil
}
