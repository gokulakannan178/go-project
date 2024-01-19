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
)

//SaveFinancialYear :""
func (d *Daos) SaveFinancialYear(ctx *models.Context, financialYear *models.FinancialYear) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).InsertOne(ctx.CTX, financialYear)
	return err
}

//GetSingleFinancialYear : ""
func (d *Daos) GetSingleFinancialYear(ctx *models.Context, UniqueID string) (*models.RefFinancialYear, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var financialYears []models.RefFinancialYear
	var financialYear *models.RefFinancialYear
	if err = cursor.All(ctx.CTX, &financialYears); err != nil {
		return nil, err
	}
	if len(financialYears) > 0 {
		financialYear = &financialYears[0]
	}
	return financialYear, nil
}

//UpdateFinancialYear : ""
func (d *Daos) UpdateFinancialYear(ctx *models.Context, financialYear *models.FinancialYear) error {
	selector := bson.M{"uniqueId": financialYear.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": financialYear, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFinancialYear : ""
func (d *Daos) FilterFinancialYear(ctx *models.Context, financialYearfilter *models.FinancialYearFilter, pagination *models.Pagination) ([]models.RefFinancialYear, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if financialYearfilter != nil {

		if len(financialYearfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": financialYearfilter.Status}})
		}
		if len(financialYearfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": financialYearfilter.UniqueID}})
		}
		if len(financialYearfilter.OperatingFinacialYear) > 0 {
			query = append(query, bson.M{"operatingFinacialYear": bson.M{"$in": financialYearfilter.OperatingFinacialYear}})
		}
		if financialYearfilter.DateRange != nil {
			//var sd,ed time.Time
			if financialYearfilter.DateRange.From != nil {
				sd := time.Date(financialYearfilter.DateRange.From.Year(), financialYearfilter.DateRange.From.Month(), financialYearfilter.DateRange.From.Day(), 0, 0, 0, 0, financialYearfilter.DateRange.From.Location())
				ed := time.Date(financialYearfilter.DateRange.From.Year(), financialYearfilter.DateRange.From.Month(), financialYearfilter.DateRange.From.Day(), 23, 59, 59, 0, financialYearfilter.DateRange.From.Location())
				if financialYearfilter.DateRange.To != nil {
					ed = time.Date(financialYearfilter.DateRange.To.Year(), financialYearfilter.DateRange.To.Month(), financialYearfilter.DateRange.To.Day(), 23, 59, 59, 0, financialYearfilter.DateRange.To.Location())
				}
				// query = append(query, bson.M{"to": bson.M{"$gte": sd, "$lte": ed}})
				query = append(query, bson.M{"to": bson.M{"$gte": sd}}, bson.M{"from": bson.M{"$lte": ed}})

			}
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if financialYearfilter != nil {
		if financialYearfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{financialYearfilter.SortBy: financialYearfilter.SortOrder}})

		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("financialYear query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var financialYears []models.RefFinancialYear
	if err = cursor.All(context.TODO(), &financialYears); err != nil {
		return nil, err
	}
	return financialYears, nil
}

//EnableFinancialYear :""
func (d *Daos) EnableFinancialYear(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FINANCIALYEARSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFinancialYear :""
func (d *Daos) DisableFinancialYear(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FINANCIALYEARSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFinancialYear :""
func (d *Daos) DeleteFinancialYear(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FINANCIALYEARSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//MakeCurrentFinancialYear : ""
func (d *Daos) MakeCurrentFinancialYear(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": bson.M{"$nin": []string{UniqueID}}}
	update := bson.M{"$set": bson.M{"isCurrent": false}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed I " + err.Error())
	}

	query2 := bson.M{"uniqueId": UniqueID}
	update2 := bson.M{"$set": bson.M{"isCurrent": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).UpdateOne(ctx.CTX, query2, update2)
	if err != nil {
		return errors.New("Not Changed I " + err.Error())
	}
	return err
}

//GetCurrentFinancialYear : ""
func (d *Daos) GetCurrentFinancialYear(ctx *models.Context) (*models.RefFinancialYear, error) {
	fy := new(models.RefFinancialYear)
	query2 := bson.M{"isCurrent": true}
	err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).FindOne(ctx.CTX, query2).Decode(&fy)
	return fy, err
}

// GetSingleFinancialYearUsingDate : ""
func (d *Daos) GetSingleFinancialYearUsingDate(ctx *models.Context, Date *time.Time) (*models.RefFinancialYear, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Date != nil {
		sd := time.Date(Date.Year(), Date.Month(), Date.Day(), 0, 0, 0, 0, Date.Location())

		query = append(query, bson.M{"to": bson.M{"$gte": sd}})
		query = append(query, bson.M{"from": bson.M{"$lte": sd}})

	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	d.Shared.BsonToJSONPrintTag("financial year using date query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var financialYears []models.RefFinancialYear
	var financialYear *models.RefFinancialYear
	if err = cursor.All(ctx.CTX, &financialYears); err != nil {
		return nil, err
	}
	if len(financialYears) > 0 {
		financialYear = &financialYears[0]
	}
	return financialYear, nil
}

// GetSingleFinancialYearUsingDate : ""
func (d *Daos) GetSingleFinancialYearUsingDateV2(ctx *models.Context, Date *time.Time) (*models.RefFinancialYear, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Date != nil {
		sd := time.Date(Date.Year(), Date.Month(), Date.Day(), 0, 0, 0, 0, Date.Location())
		ed := time.Date(Date.Year(), Date.Month(), Date.Day(), 23, 59, 59, 0, Date.Location())
		// sd := time.Date(Date.Year(), Date.Month(), Date.Day(), 0, 0, 0, 0, Date.Location())

		query = append(query, bson.M{"from": bson.M{"$lte": sd}})
		query = append(query, bson.M{"to": bson.M{"$gte": ed}})

	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	d.Shared.BsonToJSONPrintTag("financial year using date query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFINANCIALYEAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var financialYears []models.RefFinancialYear
	var financialYear *models.RefFinancialYear
	if err = cursor.All(ctx.CTX, &financialYears); err != nil {
		return nil, err
	}
	if len(financialYears) > 0 {
		financialYear = &financialYears[0]
	}
	return financialYear, nil
}
