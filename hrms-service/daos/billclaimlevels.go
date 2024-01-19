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

//SaveBillclaimLevels :""
func (d *Daos) SaveBillclaimLevels(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).InsertOne(ctx.CTX, billclaimlevels)
	if err != nil {
		return err
	}
	billclaimlevels.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleBillclaimLevels : ""
func (d *Daos) GetSingleBillclaimLevels(ctx *models.Context, uniqueID string) (*models.RefBillclaimLevels, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBILLCLAIM, "bill", "uniqueId", "ref.bill", "ref.bill")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employeeId", "ref.employeeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "assignedBy", "uniqueId", "ref.assignedBy", "ref.assignedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "approvedBy", "uniqueId", "ref.approvedBy", "ref.approvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "rejectedBy", "uniqueId", "ref.rejectedBy", "ref.rejectedBy")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billclaimlevelss []models.RefBillclaimLevels
	var billclaimlevels *models.RefBillclaimLevels
	if err = cursor.All(ctx.CTX, &billclaimlevelss); err != nil {
		return nil, err
	}
	if len(billclaimlevelss) > 0 {
		billclaimlevels = &billclaimlevelss[0]
	}
	return billclaimlevels, nil
}

//UpdateBillclaimLevels : ""
func (d *Daos) UpdateBillclaimLevels(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {
	selector := bson.M{"uniqueId": billclaimlevels.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": billclaimlevels}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableBillclaimLevels :""
func (d *Daos) EnableBillclaimLevels(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLEVELSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBillclaimLevels :""
func (d *Daos) DisableBillclaimLevels(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLEVELSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBillclaimLevels :""
func (d *Daos) DeleteBillclaimLevels(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLEVELSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterBillclaimLevels : ""
func (d *Daos) FilterBillclaimLevels(ctx *models.Context, billclaimlevelsFilter *models.BillclaimLevelsFilter, pagination *models.Pagination) ([]models.RefBillclaimLevels, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if billclaimlevelsFilter != nil {

		if len(billclaimlevelsFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": billclaimlevelsFilter.Status}})
		}
		if len(billclaimlevelsFilter.OmiStatus) > 0 {
			query = append(query, bson.M{"status": bson.M{"$nin": billclaimlevelsFilter.OmiStatus}})
		}
		if len(billclaimlevelsFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisation": bson.M{"$in": billclaimlevelsFilter.OrganisationID}})
		}
		if len(billclaimlevelsFilter.Employee) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": billclaimlevelsFilter.Employee}})
		}
		if len(billclaimlevelsFilter.AssignedBy) > 0 {
			query = append(query, bson.M{"assignedBy": bson.M{"$in": billclaimlevelsFilter.AssignedBy}})
		}
		if len(billclaimlevelsFilter.RejectedBy) > 0 {
			query = append(query, bson.M{"rejectedBy": bson.M{"$in": billclaimlevelsFilter.RejectedBy}})
		}
		//Regex
		if billclaimlevelsFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: billclaimlevelsFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if billclaimlevelsFilter != nil {
		if billclaimlevelsFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{billclaimlevelsFilter.SortBy: billclaimlevelsFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBILLCLAIM, "bill", "uniqueId", "ref.bill", "ref.bill")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "assignedBy", "uniqueId", "ref.assignedBy", "ref.assignedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "approvedBy", "uniqueId", "ref.approvedBy", "ref.approvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "rejectedBy", "uniqueId", "ref.rejectedBy", "ref.rejectedBy")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billclaimlevelss []models.RefBillclaimLevels
	fmt.Println("Bills", len(billclaimlevelss))
	if err = cursor.All(context.TODO(), &billclaimlevelss); err != nil {
		return nil, err
	}
	return billclaimlevelss, nil
}
func (d *Daos) ApprovedBillclaimLevels(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {
	selector := bson.M{"uniqueId": billclaimlevels.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLEVELSSTATUSAPPROVED, "approvedBy": billclaimlevels.ApprovedBy, "approvedDate": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) RejectedBillclaimLevels(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {
	selector := bson.M{"uniqueId": billclaimlevels.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLEVELSSTATUSREJECTED, "rejectedBy": billclaimlevels.RejectedBy, "rejectedDate": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetSingleBillclaimLevelsWithAssigned(ctx *models.Context, bill string, assign string, level int64) (*models.RefBillclaimLevels, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"bill": bill, "assignedBy": assign, "level": level}})
	// lookup
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billclaimlevelss []models.RefBillclaimLevels
	var billclaimlevels *models.RefBillclaimLevels
	if err = cursor.All(ctx.CTX, &billclaimlevelss); err != nil {
		return nil, err
	}
	if len(billclaimlevelss) > 0 {
		billclaimlevels = &billclaimlevelss[0]
	}
	return billclaimlevels, nil
}
func (d *Daos) PendingBillclaimLevels(ctx *models.Context, billclaimlevels string) error {
	selector := bson.M{"uniqueId": billclaimlevels}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"status": constants.BILLCLAIMLEVELSSTATUSPENDING}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIMLEVELS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
