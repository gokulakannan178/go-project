package daos

import (
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveBankDeposit :""
func (d *Daos) CollectionSubmissionRequest(ctx *models.Context, csr *models.CollectionSubmissionRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSITCOLLECTIONSUBMISSIONREQUEST).InsertOne(ctx.SC, csr)
	return err
}

//CollectionSubmissionRequest
func (d *Daos) CollectionSubmissionRequestFilter(ctx *models.Context, csr *models.CollectionSubmissionRequestFilter, pagination *models.Pagination) ([]models.RefCollectionSubmissionRequest, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if csr != nil {
		if len(csr.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": csr.Status}})
		}
		// if len(csr.Date) > 0 {
		// 	query = append(query, bson.M{"date": bson.M{"$in": csr.Date}})
		// }
		if len(csr.Actioner) > 0 {
			query = append(query, bson.M{"actioner": bson.M{"$in": csr.Actioner}})
		}
		if len(csr.Requestor) > 0 {
			query = append(query, bson.M{"requestor": bson.M{"$in": csr.Requestor}})
		}
		// if csr.DateRange.From != nil {
		// 	t := *bdf.DateRange.From
		// 	FromDate := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
		// 	ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		// 	if csr.DateRange.To != nil {
		// 		t2 := *csr.DateRange.To
		// 		ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		// 	}
		// 	query = append(query, bson.M{"on": bson.M{"$gte": FromDate, "$lte": ToDate}})

		// }

	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)})
		mainPipeline = append(mainPipeline, bson.M{"$limit": pagination.Limit})
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSITCOLLECTIONSUBMISSIONREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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

	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANKDEPOSITCOLLECTIONSUBMISSIONREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var RefCollectionSubmissionRequest []models.RefCollectionSubmissionRequest
	if err = cursor.All(ctx.CTX, &RefCollectionSubmissionRequest); err != nil {
		return nil, err
	}
	return RefCollectionSubmissionRequest, err
}
