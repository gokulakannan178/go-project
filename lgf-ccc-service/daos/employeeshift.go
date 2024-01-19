package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeShiftRequest : ""
func (d *Daos) SaveEmployeeShift(ctx *models.Context, employeeshift *models.EmployeeShift) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).InsertOne(ctx.CTX, employeeshift)
	if err != nil {
		return err
	}
	return nil
}

// GetSingleEmployeeShift : ""
func (d *Daos) GetSingleEmployeeShift(ctx *models.Context, uniqueID string) (*models.RefEmployeeShift, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeshifts []models.RefEmployeeShift
	var employeeshift *models.RefEmployeeShift
	if err = cursor.All(ctx.CTX, &employeeshifts); err != nil {
		return nil, err
	}
	if len(employeeshifts) > 0 {
		employeeshift = &employeeshifts[0]
	}
	return employeeshift, err
}

// UpdateEmployeeShift : ""
func (d *Daos) UpdateEmployeeShift(ctx *models.Context, employeeshift *models.EmployeeShift) error {
	selector := bson.M{"uniqueId": employeeshift.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeshift}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableEmployeeShift : ""
func (d *Daos) EnableEmployeeShift(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEESHIFTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeShift : ""
func (d *Daos) DisableEmployeeShift(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEESHIFTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteEmployeeShift :""
func (d *Daos) DeleteEmployeeShift(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEESHIFTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeShift : ""
func (d *Daos) FilterEmployeeShift(ctx *models.Context, filter *models.FilterEmployeeShift, pagination *models.Pagination) ([]models.RefEmployeeShift, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		// if len(property.GCID) > 0 {
		// 	query = append(query, bson.M{"gcUser.id": bson.M{"$in": property.GCID}})
		// }
		// if len(property.ManagerID) > 0 {
		// 	query = append(query, bson.M{"minUser.id": bson.M{"$in": property.ManagerID}})
		// }
		// if len(property.CitizenID) > 0 {
		// 	query = append(query, bson.M{"citizen.id": bson.M{"$in": property.CitizenID}})
		// }
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		// if property.Regex.Citizen.Name != "" {
		// 	query = append(query, bson.M{"citizen.name": primitive.Regex{Pattern: property.Regex.Citizen.Name, Options: "xi"}})
		// }
		// if property.Regex.HoldingNumber != "" {
		// 	query = append(query, bson.M{"holdingNumber": primitive.Regex{Pattern: property.Regex.HoldingNumber, Options: "xi"}})
		// }
		// if property.Regex.NfcID != "" {
		// 	query = append(query, bson.M{"nfcId": primitive.Regex{Pattern: property.Regex.NfcID, Options: "xi"}})
		// }
		// if property.Regex.EmployeeShift != "" {
		// 	query = append(query, bson.M{"EmployeeShift": primitive.Regex{Pattern: property.Regex.EmployeeShift, Options: "xi"}})
		// }
	}
	// if property.DateRange.From != nil {
	// 	t := *property.DateRange.From
	// 	FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	// 	ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	// 	if property.DateRange.To != nil {
	// 		t2 := *property.DateRange.To
	// 		ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
	// 	}
	// 	query = append(query, bson.M{"registerDate": bson.M{"$gte": FromDate, "$lte": ToDate}})

	// }
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).CountDocuments(ctx.CTX, func() bson.M {
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
	//LookUp
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESHIFT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeshift []models.RefEmployeeShift
	if err = cursor.All(context.TODO(), &employeeshift); err != nil {
		return nil, err
	}
	return employeeshift, nil
}
