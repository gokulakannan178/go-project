package daos

import (
	"context"
	"errors"
	"fmt"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveApptoken:""
func (d *Daos) SaveApptoken(ctx *models.Context, apptoken *models.Apptoken) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).InsertOne(ctx.CTX, apptoken)
	return err
}

//GetSingleApptoken : ""
func (d *Daos) GetSingleApptoken(ctx *models.Context, uniqueID string) (*models.RefApptoken, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var apptokens []models.RefApptoken
	var apptoken *models.RefApptoken
	if err = cursor.All(ctx.CTX, &apptokens); err != nil {
		return nil, err
	}
	if len(apptokens) > 0 {
		apptoken = &apptokens[0]
	}
	return apptoken, nil
}
func (d *Daos) SaveApptoken2(ctx *models.Context, apptoken *models.Apptoken) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"userid": apptoken.UserID, "usertype": apptoken.Usertype}
	updateData := bson.M{"$set": apptoken}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//UpdateApptoken : ""
func (d *Daos) UpdateApptoken(ctx *models.Context, apptoken *models.Apptoken) error {
	selector := bson.M{"uniqueId": apptoken.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": apptoken}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableApptoken:""
func (d *Daos) EnableApptoken(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.APPTOKENSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//Disableapptoken :""
func (d *Daos) DisableApptoken(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.APPTOKENSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteApptoken :""
func (d *Daos) DeleteApptoken(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.APPTOKENSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterApptoken : ""
func (d *Daos) FilterApptoken(ctx *models.Context, apptoken *models.ApptokenFilter, pagination *models.Pagination) ([]models.RefApptoken, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if apptoken != nil {
		if len(apptoken.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": apptoken.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("pkg query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var apptokens []models.RefApptoken
	if err = cursor.All(context.TODO(), &apptokens); err != nil {
		return nil, err
	}
	return apptokens, nil
}

// SendAppTokenNotification : ""
func (d *Daos) SendAppTokenNotification(ctx *models.Context, appToken *models.AppTokenNotification) (*models.RefAppTokenNotification, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active", "usertype": bson.M{"$in": appToken.Types}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "registrationtoken": bson.M{"$push": "$registrationtoken"}}})

	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"registrationtoken": bson.M{"$arrayElemAt": []interface{}{"$registrationtoken", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("pkg query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refapptoken []models.RefAppTokenNotification
	err = cursor.All(ctx.CTX, &refapptoken)
	if err != nil {
		return nil, err
	}
	if len(refapptoken) > 0 {
		res := refapptoken[0]
		return &res, nil
	}

	return nil, nil
}

func (d *Daos) GetReGTokensOfULB(ctx *models.Context, ulbIDs []string) (*models.RefAppTokenNotification, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active", "usertype": "ULB", "userid": bson.M{"$in": ulbIDs}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "registrationtoken": bson.M{"$push": "$registrationtoken"}}})

	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"registrationtoken": bson.M{"$arrayElemAt": []interface{}{"$registrationtoken", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("pkg query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refapptoken []models.RefAppTokenNotification
	err = cursor.All(ctx.CTX, &refapptoken)
	if err != nil {
		return nil, err
	}
	if len(refapptoken) > 0 {
		res := refapptoken[0]
		return &res, nil
	}

	return nil, nil
}

// GetReGTokenOfSingleULB : ""
func (d *Daos) GetReGTokenOfSingleULB(ctx *models.Context, ulbID string) (*models.RefAppTokenNotification, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active", "usertype": "ULB", "userid": ulbID}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "registrationtoken": bson.M{"$push": "$registrationtoken"}}})

	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"registrationtoken": bson.M{"$arrayElemAt": []interface{}{"$registrationtoken", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("pkg query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refapptoken []models.RefAppTokenNotification
	err = cursor.All(ctx.CTX, &refapptoken)
	if err != nil {
		return nil, err
	}
	if len(refapptoken) > 0 {
		res := refapptoken[0]
		return &res, nil
	}

	return nil, nil
}

// GetRegTokenWithParticulars : ""
func (d *Daos) GetRegTokenWithParticulars(ctx *models.Context, types string, ID string) (*models.Apptoken, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active", "usertype": types, "userid": ID}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetRegTokenWithParticulers query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refapptoken []*models.Apptoken
	err = cursor.All(ctx.CTX, &refapptoken)
	if err != nil {
		return nil, err
	}
	if len(refapptoken) > 0 {
		res := refapptoken[0]
		return res, nil
	}

	return nil, nil
}
func (d *Daos) GetRegTokenWithUserId(ctx *models.Context, ID string) (*models.Apptoken, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userid": ID}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetRegTokenWithParticulers query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var refapptoken []*models.Apptoken
	err = cursor.All(ctx.CTX, &refapptoken)
	if err != nil {
		return nil, err
	}
	if len(refapptoken) > 0 {
		res := refapptoken[0]
		return res, nil
	}

	return nil, nil
}
func (d *Daos) LogoutApptoken(ctx *models.Context, UniqueID string) error {
	// id, err := primitive.ObjectIDFromHex(UniqueID)
	// if err != nil {
	// 	return err
	// }
	query := bson.M{"uniqueId": UniqueID}
	result, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).DeleteOne(ctx.CTX, query)
	if err != nil {
		return errors.New("Not Deleted" + err.Error())
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return err
}

//GetSingleApptoken : ""
func (d *Daos) GetSingleApptokenWithUserID(ctx *models.Context, uniqueID string) (*models.RefApptoken, error) {
	mainPipeline := []bson.M{}
	userid, err := primitive.ObjectIDFromHex(uniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userid": userid}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	d.Shared.BsonToJSONPrintTag("Apptoken  query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var apptokens []models.RefApptoken
	var apptoken *models.RefApptoken
	if err = cursor.All(ctx.CTX, &apptokens); err != nil {
		return nil, err
	}
	if len(apptokens) > 0 {
		apptoken = &apptokens[0]
	}
	fmt.Println("apptoken===>", apptoken)
	return apptoken, nil
}
func (d *Daos) GetSingleApptokenWithUniqueCheck(ctx *models.Context, uniqueID string) (*models.RefApptoken, error) {
	ApptokenId, err := primitive.ObjectIDFromHex(uniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userid": ApptokenId}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	d.Shared.BsonToJSONPrintTag("Apptoken Uniquecheck query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPTOKEN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var apptokens []models.RefApptoken
	var apptoken *models.RefApptoken
	if err = cursor.All(ctx.CTX, &apptokens); err != nil {
		return nil, err
	}
	if len(apptokens) > 0 {
		apptoken = &apptokens[0]
	}
	return apptoken, nil
}
