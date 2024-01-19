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

// SaveDriver :""
func (d *Daos) SaveDriverDetails(ctx *models.Context, Driver *models.DriverDetails) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).InsertOne(ctx.CTX, Driver)

	return err
}

// GetSingleDriver : ""
func (d *Daos) GetSingleDriverDetails(ctx *models.Context, uniqueID string) (*models.RefDriverDetails, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONDriverLOG,
	// 		"as":   "ref.Driverlog",
	// 		"let":  bson.M{"DriverId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$status", constants.DriverASSIGNSTATUS}},
	// 				{"$eq": []string{"$DriverId", "$$DriverId"}},
	// 			}}},
	// 			},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.Driverlog": bson.M{"$arrayElemAt": []interface{}{"$ref.Driverlog", 0}}}})
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Drivers []models.RefDriverDetails
	var Driver *models.RefDriverDetails
	if err = cursor.All(ctx.CTX, &Drivers); err != nil {
		return nil, err
	}
	if len(Drivers) > 0 {
		Driver = &Drivers[0]
	}
	return Driver, nil
}

// UpdateDriver : ""
func (d *Daos) UpdateDriverDetails(ctx *models.Context, driver *models.DriverDetails) error {
	selector := bson.M{"uniqueId": driver.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": driver}
	_, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// func (d *Daos) UpdateDriverAssign(ctx *models.Context, Driver *models.DriverAssign) error {
// 	selector := bson.M{"uniqueId": Driver.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": Driver}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDriver).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// EnableDriver :""
func (d *Daos) EnableDriverDetails(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DRIVERDETAILSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableDriver :""
func (d *Daos) DisableDriverDetails(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DRIVERDETAILSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteDriver :""
func (d *Daos) DeleteDriverDetails(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DRIVERDETAILSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterDriver : ""
func (d *Daos) FilterDriverDetails(ctx *models.Context, DriverFilter *models.FilterDriverDetails, pagination *models.Pagination) ([]models.RefDriverDetails, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DriverFilter != nil {

		if len(DriverFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DriverFilter.Status}})
		}
		t := time.Now()
		ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())

		if DriverFilter.IsExpDate == "Yes" {
			query = append(query, bson.M{"licenseExpiry": bson.M{"$lte": ed}})

		}
		if DriverFilter.IsExpDate == "No" {
			query = append(query, bson.M{"licenseExpiry": bson.M{"$gte": ed}})

		}
		//Regex
		if DriverFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: DriverFilter.Regex.Name, Options: "xi"}})
		}
		if DriverFilter.Regex.LicenseType != "" {
			query = append(query, bson.M{"licenseType": primitive.Regex{Pattern: DriverFilter.Regex.LicenseType, Options: "xi"}})
		}
		if DriverFilter.Regex.Mobile != "" {
			query = append(query, bson.M{"mobile": primitive.Regex{Pattern: DriverFilter.Regex.Mobile, Options: "xi"}})
		}
	}
	if DriverFilter.DateRange.From != nil {
		t := *DriverFilter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if DriverFilter.DateRange.To != nil {
			t2 := *DriverFilter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if DriverFilter != nil {
		if DriverFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{DriverFilter.SortBy: DriverFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONDriverLOG,
	// 		"as":   "ref.Driverlog",
	// 		"let":  bson.M{"DriverId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$status", constants.DriverASSIGNSTATUS}},
	// 				{"$eq": []string{"$DriverId", "$$DriverId"}},
	// 			}}},
	// 			},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.Driverlog": bson.M{"$arrayElemAt": []interface{}{"$ref.Driverlog", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Driver query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DriversFilter []models.RefDriverDetails
	if err = cursor.All(context.TODO(), &DriversFilter); err != nil {
		return nil, err
	}
	return DriversFilter, nil
}

// func (d *Daos) DriverAssign(ctx *models.Context, Driver *models.DriverAssign) error {
// 	selector := bson.M{"uniqueId": Driver.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": Driver}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDriver).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// func (d *Daos) RevokeDriver(ctx *models.Context, Driver *models.Driver) error {
// 	selector := bson.M{"employeeId": Driver.EmployeeId}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": Driver, "status": constants.DriverREVOKESTATUS}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDriver).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// // GetSingleDriverUsingEmpID : ""
// func (d *Daos) GetSingleDriverUsingUniqueId(ctx *models.Context, UniqueID string) (*models.RefDriver, error) {
// 	mainPipeline := []bson.M{}
// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID, "status": bson.M{"$in": []string{constants.DriverSTATUSACTIVE}}}})
// 	// lookup
// 	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

// 	//Aggregation
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONDRIVER).Aggregate(ctx.CTX, mainPipeline, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var Drivers []models.RefDriver
// 	var Driver *models.RefDriver
// 	if err = cursor.All(ctx.CTX, &Drivers); err != nil {
// 		return nil, err
// 	}
// 	if len(Drivers) > 0 {
// 		Driver = &Drivers[0]
// 	}
// 	return Driver, nil
// }
