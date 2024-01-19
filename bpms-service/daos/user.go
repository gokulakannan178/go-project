package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "organisationId", "uniqueId", "ref.tempULB", "ref.tempULB")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "organisationId", "uniqueId", "ref.dept", "ref.dept")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ref.dept.organisationId", "uniqueId", "ref.tempDeptULB", "ref.tempDeptULB")...)

	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"ref.ulb": bson.M{
			"$cond": bson.M{"if": bson.M{"$eq": []string{"$type", "ULB"}}, "then": "$ref.tempULB", "else": "$ref.tempDeptULB"},
		},

			"ref.department": bson.M{
				"$cond": bson.M{"if": bson.M{"$eq": []string{"$type", "Department"}}, "then": "$ref.dept", "else": nil},
			},
		}})

	fmt.Println("GetSingleUser query", mainPipeline)
	d.Shared.BsonToJSONPrintTag("GetSingleUser query", mainPipeline)
	fmt.Println("DB", ctx.DB.Name())

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var users []models.RefUser
	var User *models.RefUser
	if err = cursor.All(ctx.CTX, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		User = &users[0]
	}
	return User, nil
}

//UpdateUser : ""
func (d *Daos) UpdateUser(ctx *models.Context, User *models.User) error {
	selector := bson.M{"userName": User.UserName}
	// t := time.Now()
	// update := models.Updated{}
	// update.On = &t
	// update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": User}
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
		if len(userfilter.UserType) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": userfilter.UserType}})
		}
		if len(userfilter.Organisation) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": userfilter.Organisation}})
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

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "organisationId", "uniqueId", "ref.tempULB", "ref.tempULB")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "organisationId", "uniqueId", "ref.dept", "ref.dept")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ref.dept.organisationId", "uniqueId", "ref.tempDeptULB", "ref.tempDeptULB")...)

	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"ref.ulb": bson.M{
			"$cond": bson.M{"if": bson.M{"$eq": []string{"$type", "ULB"}}, "then": "$ref.tempULB", "else": "$ref.tempDeptULB"},
		},

			"ref.department": bson.M{
				"$cond": bson.M{"if": bson.M{"$eq": []string{"$type", "Department"}}, "then": "$ref.dept", "else": nil},
			},
		}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("User query =>", mainPipeline)
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
