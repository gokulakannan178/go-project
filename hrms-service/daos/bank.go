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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveBankInformation :""
func (d *Daos) SaveBankInformation(ctx *models.Context, bankInformation *models.BankInformation) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).InsertOne(ctx.CTX, bankInformation)
	if err != nil {
		return err
	}
	bankInformation.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleBankInformation : ""
func (d *Daos) GetSingleBankInformation(ctx *models.Context, UniqueID string) (*models.RefBankInformation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var bankInformations []models.RefBankInformation
	var bankInformation *models.RefBankInformation
	if err = cursor.All(ctx.CTX, &bankInformations); err != nil {
		return nil, err
	}
	if len(bankInformations) > 0 {
		bankInformation = &bankInformations[0]
	}
	return bankInformation, nil
}

//UpdateBankInformation : ""
func (d *Daos) UpdateBankInformation(ctx *models.Context, bankInformation *models.BankInformation) error {
	selector := bson.M{"uniqueId": bankInformation.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bankInformation}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterBankInformation : ""
func (d *Daos) FilterBankInformation(ctx *models.Context, filter *models.BankInformationFilter, pagination *models.Pagination) ([]models.RefBankInformation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		//Regex
		if filter.Regex.BankName != "" {
			query = append(query, bson.M{"bankName": primitive.Regex{Pattern: filter.Regex.BankName, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("BankInformation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var bankInformations []models.RefBankInformation
	if err = cursor.All(context.TODO(), &bankInformations); err != nil {
		return nil, err
	}
	return bankInformations, nil
}

//EnableBankInformation :""
func (d *Daos) EnableBankInformation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BANKINFORMATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBankInformation :""
func (d *Daos) DisableBankInformation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BANKINFORMATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBankInformation :""
func (d *Daos) DeleteBankInformation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BANKINFORMATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) UpdateEmployeeBankInformation(ctx *models.Context, bankInformation *models.BankInformation) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"employeeId": bankInformation.EmployeeID}
	updateData := bson.M{"$set": bankInformation}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetSingleBankInformationWithEmployee(ctx *models.Context, UniqueID string) (*models.RefBankInformation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANKINFORMATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var bankInformations []models.RefBankInformation
	var bankInformation *models.RefBankInformation
	if err = cursor.All(ctx.CTX, &bankInformations); err != nil {
		return nil, err
	}
	if len(bankInformations) > 0 {
		bankInformation = &bankInformations[0]
	}
	return bankInformation, nil
}
