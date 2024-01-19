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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveUser :""
func (d *Daos) SaveUser(ctx *models.Context, user *models.User) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).InsertOne(ctx.CTX, user)
	return err
}

//GetSingleUser : ""
func (d *Daos) GetSingleUser(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userName": UserName}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "managerId", "type", "ref.userType", "ref.userType")...)

	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISTRICT, "accessPrivilege.districts", "code", "ref.accessDistricts", "ref.accessDistricts")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "accessPrivilege.states", "code", "ref.accessStates", "ref.accessStates")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONZONE, "accessPrivilege.zones", "code", "ref.accessZones", "ref.accessZones")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONWARD, "accessPrivilege.wards", "code", "ref.accessWards", "ref.accessWards")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "accessPrivilege.villages", "code", "ref.accessVillages", "ref.accessVillages")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefUser
	var user *models.RefUser
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//UpdateUser : ""
func (d *Daos) UpdateUser(ctx *models.Context, user *models.User) error {
	selector := bson.M{"userName": user.UserName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": user}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// UpdateForcedLogout : ""
func (d *Daos) UpdateForcedLogout(ctx *models.Context, userName string, Action string) error {
	selector := bson.M{"userName": userName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"isForcedLogout": Action}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// UpdateUserToken : ""
func (d *Daos) UpdateUserToken(ctx *models.Context, userName string, token string) error {
	selector := bson.M{"userName": userName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"token": token}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterUser : ""
func (d *Daos) FilterUser(ctx *models.Context, userfilter *models.UserFilter, pagination *models.Pagination) ([]models.RefUser, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if userfilter != nil {
		if len(userfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": userfilter.Status}})
		}
		if len(userfilter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": userfilter.UserName}})
		}
		if len(userfilter.Manager) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": userfilter.Manager}})
		}
		if len(userfilter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": userfilter.Type}})
		}
		if len(userfilter.OmitID) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$nin": userfilter.OmitID}})
		}
		if len(userfilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": userfilter.OrganisationID}})
		}
		if len(userfilter.AccessWardCode) > 0 {
			query = append(query, bson.M{"accessPrivilege.wards": bson.M{"$in": userfilter.AccessWardCode}})
		}
		if len(userfilter.MpinStatus) > 0 {
			query = append(query, bson.M{"mobileAuth.mpinStatus": bson.M{"$in": userfilter.MpinStatus}})
		}

		//Regex
		if userfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: userfilter.Regex.Name, Options: "xi"}})
		}
		if userfilter.Regex.Contact != "" {
			query = append(query, bson.M{"mobile": primitive.Regex{Pattern: userfilter.Regex.Contact, Options: "xi"}})
		}
		if userfilter.Regex.UserName != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: userfilter.Regex.UserName, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSER).CountDocuments(ctx.CTX, func() bson.M {
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
	if userfilter.GetRecentLocation {
		mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
			"from": constants.COLLECTIONUSERLOCATION,
			"as":   "ref.lastLocation",
			"let":  bson.M{"userName": "$userName"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$userName", "$$userName"}},
				}}}},
				bson.M{"$sort": bson.M{"time": -1}},
				bson.M{"$limit": 1},
			},
		}})
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.lastLocation": bson.M{"$arrayElemAt": []interface{}{"$ref.lastLocation", 0}}}})

	}

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "managerId", "type", "ref.userType", "ref.userType")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISTRICT, "accessPrivilege.districts", "code", "ref.accessDistricts", "ref.accessDistricts")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "accessPrivilege.states", "code", "ref.accessStates", "ref.accessStates")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONZONE, "accessPrivilege.zones", "code", "ref.accessZones", "ref.accessZones")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONWARD, "accessPrivilege.wards", "code", "ref.accessWards", "ref.accessWards")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "accessPrivilege.villages", "code", "ref.accessVillages", "ref.accessVillages")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefUser
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

//EnableUser :""
func (d *Daos) EnableUser(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": constants.USERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableUser :""
func (d *Daos) DisableUser(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": constants.USERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteUser :""
func (d *Daos) DeleteUser(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": constants.USERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//ResetUserPassword : ""
func (d *Daos) ResetUserPassword(ctx *models.Context, userName string, password string) error {
	selector := bson.M{"userName": userName}
	updateInterface := bson.M{"$set": bson.M{"password": password}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//PasswordUpdate : ""
func (d *Daos) PasswordUpdate(ctx *models.Context, user *models.RefPassword) error {

	selector := bson.M{"userName": user.UniqueID}
	updateInterface := bson.M{"$set": bson.M{"password": user.Password}}

	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//UserCollectionLimit : ""
func (d *Daos) UserCollectionLimit(ctx *models.Context, UserName string, cl *models.CollectionLimit) error {

	selector := bson.M{"userName": UserName}
	updateInterface := bson.M{"$set": bson.M{"collectionLimit.cash": cl.Cash}}

	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//GetSingleUserWithUniqueID : ""
func (d *Daos) GetSingleUserWithUniqueID(ctx *models.Context, userName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"userName": bson.M{"$regex": "^" + userName + "$", "$options": "i"}})
	orQuery = append(orQuery, bson.M{"email": userName})
	orQuery = append(orQuery, bson.M{"mobile": userName})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$or": orQuery}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefUser
	var user *models.RefUser
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//GetSingleUserWithUniqueID : ""
func (d *Daos) GetSingleActiveUserWithUniqueID(ctx *models.Context, userName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"userName": bson.M{"$regex": "^" + userName + "$", "$options": "i"}})
	orQuery = append(orQuery, bson.M{"email": userName})
	orQuery = append(orQuery, bson.M{"mobile": userName})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$or": orQuery}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": "Active"}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefUser
	var user *models.RefUser
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//UpdateAccessPrivilege : ""
func (d *Daos) UpdateAccessPrivilege(ctx *models.Context, user *models.User) error {
	selector := bson.M{"userName": user.UserName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": user}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) UpdateAppVersionUser(ctx *models.Context, user *models.AppVersionUser) error {
	selector := bson.M{"userName": user.UserName, "mobile": user.MobileNo}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"appVersion": user.AppVersion}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) UpdateAppVersionUserV2(ctx *models.Context, user *models.User) error {
	selector := bson.M{"mobileAuth.deviceId": user.MobileAuth.DeviceID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"appVersion": user.AppVersion}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) UserMpinRegistration(ctx *models.Context, user *models.User) error {
	selector := bson.M{"userName": user.UserName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"mobileAuth.deviceId": user.MobileAuth.DeviceID, "mobileAuth.date": &t, "mobileAuth.mpinStatus": constants.MPINSTATUSACTIVE, "mobileAuth.mpin": user.MobileAuth.Mpin}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// VerifyUserMpinRegistration : "This api is used for mpin verification"
func (d *Daos) VerifyUserMpinRegistration(ctx *models.Context, user *models.User) error {
	selector := bson.M{"userName": user.UserName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"mobileAuth.mpinStatus": constants.MPINSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//GetSingleUser : ""
func (d *Daos) GetSingleUserWithUsername(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"userName": UserName})
	orQuery = append(orQuery, bson.M{"mobile": UserName})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$or": orQuery}})
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefUser
	var user *models.RefUser
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	log.Println("user =========>", user)
	return user, nil

}

// GetSingleUserWithDeviceID : ""
func (d *Daos) GetSingleUserWithDeviceID(ctx *models.Context, DeviceID string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileAuth.deviceId": DeviceID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "managerId", "type", "ref.userType", "ref.userType")...)

	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISTRICT, "accessPrivilege.districts", "code", "ref.accessDistricts", "ref.accessDistricts")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "accessPrivilege.states", "code", "ref.accessStates", "ref.accessStates")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONZONE, "accessPrivilege.zones", "code", "ref.accessZones", "ref.accessZones")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONWARD, "accessPrivilege.wards", "code", "ref.accessWards", "ref.accessWards")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "accessPrivilege.villages", "code", "ref.accessVillages", "ref.accessVillages")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefUser
	var user *models.RefUser
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, nil
}

//
func (d *Daos) RemovedUserToken(ctx *models.Context, UniqueID string) error {
	selector := bson.M{"userName": UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"token": ""}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
