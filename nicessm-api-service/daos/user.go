package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveUser :""
func (d *Daos) SaveUser(ctx *models.Context, user *models.User) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONUSER).InsertOne(ctx.CTX, user)
	if err != nil {
		return err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleUser : ""
func (d *Daos) GetSingleUser(ctx *models.Context, userName string) (*models.RefUser, error) {
	id, err := primitive.ObjectIDFromHex(userName)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "blockCode", "_id", "ref.block", "ref.block")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "grampanchayatCode", "_id", "ref.grampanchayat", "ref.grampanchayat")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "districtCode", "_id", "ref.district", "ref.district")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "villageCode", "_id", "ref.village", "ref.village")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "_id", "ref.state", "ref.state")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "userOrg", "_id", "ref.organisation", "ref.organisation")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISTRICT, "accessPrivilege.districts", "_id", "ref.accessDistricts", "ref.accessDistricts")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "accessPrivilege.states", "_id", "ref.accessStates", "ref.accessStates")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "accessPrivilege.villages", "_id", "ref.accessVillages", "ref.accessVillages")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONBLOCK, "accessPrivilege.blocks", "_id", "ref.accessBlocks", "ref.accessBlocks")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONGRAMPANCHAYAT, "accessPrivilege.grampanchayats", "_id", "ref.accessGrampanchayats", "ref.accessGrampanchayats")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomains", "_id", "ref.knowledgeDomains", "ref.knowledgeDomains")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSUBDOMAIN, "subDomains", "_id", "ref.subDomains", "ref.subDomains")...)
	mainPipeline = append(mainPipeline, d.UserlookupQueryConstration(ctx)...)
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROJECTUSER,
		"as":   "ref.projects",
		"let":  bson.M{"userId": "$_id"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []interface{}{"$status", "Active"}},
				{"$eq": []string{"$user", "$$userId"}},
			}}}},
			{"$lookup": bson.M{
				"from": constants.COLLECTIONPROJECT,
				"as":   "ref.project",
				"let":  bson.M{"projectId": "$project"},
				"pipeline": []bson.M{
					{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						{"$eq": []interface{}{"$status", "Active"}},
						{"$eq": []string{"$_id", "$$projectId"}},
					}}}},
				},
			}},
			{"$addFields": bson.M{"ref.project": bson.M{"$arrayElemAt": []interface{}{"$ref.project", 0}}}},
		},
	}})

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

//GetSingleUserWithUserName : ""
func (d *Daos) GetSingleUserWithUserName(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userName": UserName}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "userOrg", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	mainPipeline = append(mainPipeline, d.UserlookupQueryConstration(ctx)...)

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
	} else {
		return nil, errors.New("user not fount")
	}
	return user, nil
}

//GetSingleUser : ""
func (d *Daos) GetMobileValidation(ctx *models.Context, Mobile string) error {
	selector := bson.M{"mobile": Mobile}
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Find(ctx.CTX, selector)
	if err != nil {
		return err
	}
	var users []models.RefUser
	//var user *models.RefUser
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return err
	}
	if len(users) > 0 {
		return errors.New("mobile number already exists")
	}

	return nil

}

//UpdateUser : ""
func (d *Daos) UpdateUser(ctx *models.Context, user *models.User) error {
	selector := bson.M{"_id": user.ID}
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

//FilterUser : ""
func (d *Daos) FilterUser(ctx *models.Context, userfilter *models.UserFilter, pagination *models.Pagination) ([]models.RefUser, error) {

	paginationPipeline := []bson.M{}
	mainPipeline := d.UserReportFilter(ctx, userfilter)

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("user pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	mainPipeline = append(mainPipeline, d.UserlookupQueryConstration(ctx)...)
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
func (d *Daos) UserlookupQueryConstration(ctx *models.Context) []bson.M {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "blockCode", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "grampanchayatCode", "_id", "ref.grampanchayat", "ref.grampanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "districtCode", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "villageCode", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "stateCode", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "userOrg", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISTRICT, "accessPrivilege.districts", "_id", "ref.accessDistricts", "ref.accessDistricts")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "accessPrivilege.states", "_id", "ref.accessStates", "ref.accessStates")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "accessPrivilege.villages", "_id", "ref.accessVillages", "ref.accessVillages")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONBLOCK, "accessPrivilege.blocks", "_id", "ref.accessBlocks", "ref.accessBlocks")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONGRAMPANCHAYAT, "accessPrivilege.grampanchayats", "_id", "ref.accessGrampanchayats", "ref.accessGrampanchayats")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomains", "_id", "ref.knowledgeDomains", "ref.knowledgeDomains")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSUBDOMAIN, "subDomains", "_id", "ref.subDomains", "ref.subDomains")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "type", "uniqueId", "ref.type", "ref.type")...)

	return mainPipeline

}

//EnableUser :""
func (d *Daos) EnableUser(ctx *models.Context, UserName string) error {
	id, err := primitive.ObjectIDFromHex(UserName)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableUser :""
func (d *Daos) DisableUser(ctx *models.Context, UserName string) error {
	id, err := primitive.ObjectIDFromHex(UserName)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteUser :""
func (d *Daos) DeleteUser(ctx *models.Context, UserName string) error {
	id, err := primitive.ObjectIDFromHex(UserName)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, query, update)
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

	selector := bson.M{"mobileNumber": user.UniqueID}
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

//GetSingleUserWithMobileNo : ""
func (d *Daos) GetSingleUserWithMobileNo(ctx *models.Context, mobileno string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNumber": mobileno}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "userOrg", "_id", "ref.organisation", "ref.organisation")...)
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

//UpdateUserTypeV2 : ""
func (d *Daos) UpdateUserTypeV2(ctx *models.Context, user *models.User) error {

	selector := bson.M{"_id": user.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"type": user.Type, "role": user.Role}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateMany(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//UpdateUserPassword : ""
func (d *Daos) UpdateUserPassword(ctx *models.Context, user *models.User) error {

	selector := bson.M{"_id": user.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"userName": user.UserName, "password": user.Pass}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UserUniquenessCheckRegistration(ctx *models.Context, orgID string, param string, value string) (*models.UserUniquinessChk, error) {
	id, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return nil, err
	}
	query := bson.M{
		"userOrg": id,
		param:     value,
	}
	result := new(models.UserUniquinessChk)
	var user *models.User
	if err = ctx.DB.Collection(constants.COLLECTIONUSER).FindOne(ctx.CTX, query).Decode(&user); err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	if user == nil {
		result.Success = true
		return result, nil
	}
	if user.Status == constants.USERSTATUSACTIVE {
		result.Success = false
		result.Message = "user status alreasy exist"
		return result, nil
	}
	if user.Status == constants.USERSTATUSDISABLED {
		result.Success = false
		result.Message = "user disabled - please contact Administator"
		return result, nil
	}
	if user.Status == constants.USERSTATUSINIT {
		result.Success = false
		result.Message = "Awaiting Activation - please contact administator"
		return result, nil
	}
	return nil, errors.New("No options Availables")
}

//RejectUser :""
func (d *Daos) RejectUser(ctx *models.Context, UserName string) error {
	id, err := primitive.ObjectIDFromHex(UserName)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.USERSTATUSREJECT}}
	_, err = ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// GetSingleUserWithMobileNoAndEmailID : ""
func (d *Daos) GetSingleUserWithMobileNoAndEmailID(ctx *models.Context, mobileNo string, emailID string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNumber": mobileNo, "email": emailID}})
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
func (d *Daos) GetContentDisseminationUser(ctx *models.Context, UserFilter *models.UserFilter) ([]models.DissiminateUser, error) {
	mainPipeline := d.UserReportFilter(ctx, UserFilter)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAPPTOKEN, "_id", "userid", "appRegistrationToken", "appRegistrationToken")...)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"appRegistrationToken": "$appRegistrationToken.registrationtoken"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"name": bson.M{"$concat": []string{"$firstName", " ", "$lastname"}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{"_id": 1, "name": 1, "mobileNumber": 1, "email": 1, "appRegistrationToken": 1, "farmerID": 1},
	})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})

	d.Shared.BsonToJSONPrintTag("GetContentDisseminationUser", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.DissiminateUser
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil

}
func (d *Daos) GetSingleUserWithQueryCount(ctx *models.Context, KD string, SD string) (*models.RefUser, error) {

	KDid, err := primitive.ObjectIDFromHex(KD)
	if err != nil {
		return nil, err
	}
	SDid, err := primitive.ObjectIDFromHex(SD)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{
		"knowledgeDomains": bson.M{"$eq": KDid},
		"subDomains":       bson.M{"$eq": SDid},
		"type":             "Subject_Matter_Expert"}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONQUERY,
		"as":   "queryCount",
		"let":  bson.M{"userCode": "$_id"},
		"pipeline": []bson.M{

			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

				{"$eq": []string{"$status", "O"}},
				{"$eq": []string{"$assignedTo", "$$userCode"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"queryCount": bson.M{"$arrayElemAt": []interface{}{"$queryCount", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"queryCount.count": 1}})
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
	} else {
		return nil, errors.New("user not fount")
	}
	return user, nil
}
