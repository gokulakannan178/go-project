package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
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
func (d *Daos) GetSingleUser(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userName": UserName}})
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

//GetSingleUserWithLogin : ""
func (d *Daos) GetSingleUserWithLoginId(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"loginId": UserName})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$or": query}})
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

//GetSingleUserWithLogin : ""
func (d *Daos) GetSingleUserWithLogin(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"userName": UserName})
	query = append(query, bson.M{"mobile": UserName})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$or": query}})
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

//GetSingleUserbyemployeeid : ""
func (d *Daos) GetSingleUserbyemployeeid(ctx *models.Context, EmployeeID string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": EmployeeID}})
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)

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
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	//mainPipeline = append(mainPipeline, d.UserlookupQueryConstration(ctx)...)

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
func (d *Daos) GetSingleUserWithEmployedID(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": UserName}})
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "managerId", "userName", "ref.manager", "ref.manager")...)
	//mainPipeline = append(mainPipeline, d.UserlookupQueryConstration(ctx)...)

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

//UpdateUserbyemployeeId : ""
func (d *Daos) UpdateUserbyemployeeId(ctx *models.Context, user *models.User) error {
	selector := bson.M{"employeeId": user.EmployeeId}
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
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if userfilter != nil {
		if len(userfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": userfilter.Status}})
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
			log.Println("Error in getting pagination")
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
func (d *Daos) UpdateUserProfileImage(ctx *models.Context, Employee *models.UpdateBioData, UniqueID string) error {
	selector := bson.M{"userName": UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"profileImg": Employee.ProfileImg}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
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

//New Password : ""
func (d *Daos) ForgetPasswordNewPassword(ctx *models.Context, userName string, password string) error {
	selector := bson.M{"userName": userName}
	updateInterface := bson.M{"$set": bson.M{"password": password}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetSingleUserWithUniqueId(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userName": UserName}})
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
func (d *Daos) UserUniquenessCheckRegistration(ctx *models.Context, orgID string, param string, value string) (*models.UserUniquinessChk, error) {
	id, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	query := bson.M{
		"organisationId": id,
		param:            value,
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": query})
	}
	fmt.Println("query====>", query)

	result := new(models.UserUniquinessChk)
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
	fmt.Println("userp===>", user)
	if user == nil {
		result.Success = true
		result.Message = "User Not Found"
		return result, nil
	}
	if user != nil {
		result.Success = false
		result.Message = "User  Found"
		return result, nil
	}

	return result, nil
}
func (d *Daos) UpdateUserLoginId(ctx *models.Context, username string, loginid string) error {
	selector := bson.M{"userName": username}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"loginId": loginid}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
