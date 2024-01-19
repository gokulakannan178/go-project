package daos

import (
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//FPOUpdateLocation : ""
func (d *Daos) FPOUpdateLocation(ctx *models.Context, fpoloc *models.FPOUpdateLocation) error {
	updated := new(models.Updated)
	t := time.Now()
	updated.On = &t
	updated.By = fpoloc.By
	updated.ByType = fpoloc.ByType
	updated.Scenario = "addingLocation"
	selector := bson.M{"uniqueId": fpoloc.UniqueID}
	updateInterface := bson.M{"$set": bson.M{"address.location": fpoloc.Location, "updated": updated}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFPO).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
