package daos

import (
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveSurveyAndTax : ""
func (d *Daos) SaveSurveyAndTax(ctx *models.Context, sat *models.SurveyAndTax) error {
	d.Shared.BsonToJSONPrint(sat)
	_, err := ctx.DB.Collection(constants.COLLECTIONSURVEYANDTAX).InsertOne(ctx.CTX, sat)
	return err
}

//GetSingleSurveyAndTax : ""
func (d *Daos) GetSingleSurveyAndTax(ctx *models.Context, uniqueID string) (*models.RefSurveyAndTax, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSURVEYANDTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var sats []models.RefSurveyAndTax
	var sat *models.RefSurveyAndTax
	if err = cursor.All(ctx.CTX, &sats); err != nil {
		return nil, err
	}
	if len(sats) > 0 {
		sat = &sats[0]
	}
	return sat, err
}

// SurveyAndTaxFilter : ""
func (d *Daos) SurveyAndTaxFilter(ctx *models.Context, satf *models.SurveyAndTaxFilter, pagination *models.Pagination) ([]models.RefSurveyAndTax, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if satf != nil {
		if len(satf.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": satf.Status}})
			if satf.DateRange.From != nil {
				t := *satf.DateRange.From
				FromDate := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
				ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
				if satf.DateRange.To != nil {
					t2 := *satf.DateRange.To
					ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
				}
				query = append(query, bson.M{"on": bson.M{"$gte": FromDate, "$lte": ToDate}})
			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSURVEYANDTAX).CountDocuments(ctx.CTX, func() bson.M {
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

	cursor, err := ctx.DB.Collection(constants.COLLECTIONSURVEYANDTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var sats []models.RefSurveyAndTax
	if err = cursor.All(ctx.CTX, &sats); err != nil {
		return nil, err
	}
	return sats, err
}

func (d *Daos) CalcSurveyAndTaxFilter(ctx *models.Context, date *time.Time) (*models.RefSurveyAndTax, error) {
	// dateStr, err := d.Shared.UniqueDateStr(date)
	// if err != nil {
	// 	return nil, err
	// }
	// query := bson.M{"uniqueId": dateStr}
	// count, err := ctx.DB.Collection(constants.COLLECTIONSURVEYANDTAX).CountDocuments(ctx.CTX, query)

	return nil, nil

}
