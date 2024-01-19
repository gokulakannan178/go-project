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

// SaveVechile :""
func (d *Daos) SaveVechile(ctx *models.Context, Vechile *models.Vechile) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).InsertOne(ctx.CTX, Vechile)

	return err
}

// GetSingleVechile : ""
func (d *Daos) GetSingleVechile(ctx *models.Context, uniqueID string) (*models.RefVechile, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVEHICLELOG, "uniqueId", "vehicle.uniqueId", "ref.vehiclelog", "ref.vehiclelog")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVEHICLETYPE, "vechileTypeId", "uniqueId", "ref.vechileTypeId", "ref.vechileTypeId")...)
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Vechiles []models.RefVechile
	var Vechile *models.RefVechile
	if err = cursor.All(ctx.CTX, &Vechiles); err != nil {
		return nil, err
	}
	if len(Vechiles) > 0 {
		Vechile = &Vechiles[0]
	}
	return Vechile, nil
}

// UpdateVechile : ""
func (d *Daos) UpdateVechile(ctx *models.Context, Vechile *models.Vechile) error {
	selector := bson.M{"uniqueId": Vechile.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Vechile}
	_, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdateDriverWithVehicleId(ctx *models.Context, Vechile *models.Vechile) error {
	selector := bson.M{"uniqueId": Vechile.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Vechile}
	_, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// func (d *Daos) UpdateVechileAssign(ctx *models.Context, Vechile *models.VechileAssign) error {
// 	selector := bson.M{"uniqueId": Vechile.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": Vechile}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONVechile).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// EnableVechile :""
func (d *Daos) EnableVechile(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VECHILESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableVechile :""
func (d *Daos) DisableVechile(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VECHILESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteVechile :""
func (d *Daos) DeleteVechile(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VECHILESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterVechile : ""
func (d *Daos) FilterVechile(ctx *models.Context, VechileFilter *models.FilterVechile, pagination *models.Pagination) ([]models.RefVechile, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if VechileFilter != nil {

		if len(VechileFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": VechileFilter.Status}})
		}
		if len(VechileFilter.ManagerID) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": VechileFilter.ManagerID}})
		}
		if len(VechileFilter.DumbsiteID) > 0 {
			query = append(query, bson.M{"dumbSite.id": bson.M{"$in": VechileFilter.DumbsiteID}})
		}
		t := time.Now()
		ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())

		if VechileFilter.IsExpDate == "Yes" {
			query = append(query, bson.M{"insuranceValid": bson.M{"$lte": ed}})

		}
		if VechileFilter.IsExpDate == "No" {
			query = append(query, bson.M{"insuranceValid": bson.M{"$gte": ed}})

		}
		//Regex
		if VechileFilter.Regex.Name != "" {
			query = append(query, bson.M{"vechileName": primitive.Regex{Pattern: VechileFilter.Regex.Name, Options: "xi"}})
		}
		if VechileFilter.Regex.LicenseNo != "" {
			query = append(query, bson.M{"licenseNo": primitive.Regex{Pattern: VechileFilter.Regex.LicenseNo, Options: "xi"}})
		}
		if VechileFilter.Regex.PUCNo != "" {
			query = append(query, bson.M{"pUCNo": primitive.Regex{Pattern: VechileFilter.Regex.PUCNo, Options: "xi"}})
		}
		if VechileFilter.Regex.RcNO != "" {
			query = append(query, bson.M{"rcNO": primitive.Regex{Pattern: VechileFilter.Regex.RcNO, Options: "xi"}})
		}

	}
	if VechileFilter.DateRange.From != nil {
		t := *VechileFilter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if VechileFilter.DateRange.To != nil {
			t2 := *VechileFilter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if VechileFilter != nil {
		if VechileFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{VechileFilter.SortBy: VechileFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVEHICLETYPE, "vechileTypeId", "uniqueId", "ref.vechileTypeId", "ref.vechileTypeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVEHICLELOG, "uniqueId", "vehicle.uniqueId", "ref.vehiclelog", "ref.vehiclelog")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Vechile query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var VechilesFilter []models.RefVechile
	if err = cursor.All(context.TODO(), &VechilesFilter); err != nil {
		return nil, err
	}
	return VechilesFilter, nil
}

func (d *Daos) VechileAssign(ctx *models.Context, assign *models.VechileAssign) error {
	selector := bson.M{"uniqueId": assign.VehicleID}
	refdriver, err := d.GetSingleDriver(ctx, assign.DriverID)
	if err != nil {
		return errors.New("error in getting the Vechilelog- " + err.Error())
	}
	fmt.Println("refdriver===========>>>>>>>", refdriver.UniqueID)
	Vechile := new(models.Driver)
	if refdriver != nil {
		fmt.Println("refdriver.UniqueID===========", refdriver.UniqueID)
		Vechile.ID = assign.DriverID
		Vechile.Name = refdriver.Name
	}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"driver": Vechile, "isDriverAssign": "Yes"}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVECHILE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// func (d *Daos) RevokeVechile(ctx *models.Context, Vechile *models.Vechile) error {
// 	selector := bson.M{"employeeId": Vechile.EmployeeId}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": Vechile, "status": constants.VechileREVOKESTATUS}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONVechile).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// GetSingleVechileUsingEmpID : ""
func (d *Daos) GetSingleVechileUsingUniqueId(ctx *models.Context, UniqueID string) (*models.RefVechile, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID, "status": bson.M{"$in": []string{constants.VECHILESTATUSACTIVE}}}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Vechiles []models.RefVechile
	var Vechile *models.RefVechile
	if err = cursor.All(ctx.CTX, &Vechiles); err != nil {
		return nil, err
	}
	if len(Vechiles) > 0 {
		Vechile = &Vechiles[0]
	}
	return Vechile, nil
}

func (d *Daos) UpdateVechileLocation(ctx *models.Context, Vechile *models.VehicleLocationUpdate) error {
	selector := bson.M{"uniqueId": Vechile.VehicleID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"location": Vechile.Location}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVECHILE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
