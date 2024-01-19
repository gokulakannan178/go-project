package daos

import (
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveBankDeposit :""
func (d *Daos) SaveBankDeposit(ctx *models.Context, bankDeposits *models.BankDeposit) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSIT).InsertOne(ctx.SC, bankDeposits)
	return err
}

//Verify Bank Deposit
func (d *Daos) VerifyBankDeposit(ctx *models.Context, uniqueID string, status string, verifier *models.BankDepositVerifier) error {
	selector := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": status, "Verifier": verifier}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSIT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

//GetSingleBankDeposit
func (d *Daos) GetSingleBankDeposit(ctx *models.Context, uniqueID string) (*models.BankDeposit, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSIT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var bankDeposits []models.BankDeposit
	var bankDeposit *models.BankDeposit
	if err = cursor.All(ctx.CTX, &bankDeposits); err != nil {
		return nil, err
	}
	if len(bankDeposits) > 0 {
		bankDeposit = &bankDeposits[0]
	}
	return bankDeposit, err
}

//NonVerify Bank Deposit
func (d *Daos) NotVerifyBankDeposit(ctx *models.Context, uniqueID string, status string, verifier *models.BankDepositVerifier) error {
	selector := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": status, "Verifier": verifier}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSIT).UpdateOne(ctx.CTX, selector, update)
	if err != nil {
		fmt.Println("Not Changed" + err.Error())
		return err
	}
	return nil
}

//BankDepositFilter
func (d *Daos) BankDepositFilter(ctx *models.Context, bdf *models.BankDepositFilter, pagination *models.Pagination) ([]models.BankDeposit, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if bdf != nil {
		if len(bdf.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": bdf.Status}})
		}
		if len(bdf.UserName) > 0 {
			query = append(query, bson.M{"username": bson.M{"$in": bdf.UserName}})
		}
		if len(bdf.UserType) > 0 {
			query = append(query, bson.M{"usertype": bson.M{"$in": bdf.UserType}})
		}
		if bdf.DateRange.From != nil {
			t := *bdf.DateRange.From
			FromDate := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
			ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
			if bdf.DateRange.To != nil {
				t2 := *bdf.DateRange.To
				ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
			}
			query = append(query, bson.M{"on": bson.M{"$gte": FromDate, "$lte": ToDate}})

		}

	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)})
		mainPipeline = append(mainPipeline, bson.M{"$limit": pagination.Limit})
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSIT).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSIT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var BankDeposits []models.BankDeposit
	if err = cursor.All(ctx.CTX, &BankDeposits); err != nil {
		return nil, err
	}
	return BankDeposits, err
}
