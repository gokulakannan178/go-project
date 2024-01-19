package daos

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePunchIn :""
func (d *Daos) SavePunchIn(ctx *models.Context, user *models.UserAttendanceAction) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERATTENDANCE).InsertOne(ctx.CTX, user)
	return err
}

//SavePunchOut :""
func (d *Daos) SavePunchOut(ctx *models.Context, user *models.UserAttendance) error {
	query := bson.M{"uniqueId": user.UniqueID}
	after := options.After
	fmt.Println(after)
	data := ctx.DB.Collection(constants.COLLECTIONUSERATTENDANCE).FindOne(ctx.CTX, query, options.FindOne())
	// update := bson.M{"$set": bson.M{"userName": user.UserName, "uniqueId": user.UniqueID}}
	update := bson.M{"$set": bson.M{"userName": user.UserName, "uniqueId": data}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERATTENDANCE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
