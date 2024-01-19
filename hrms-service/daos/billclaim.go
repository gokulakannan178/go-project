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

//SaveBillClaim :""
func (d *Daos) SaveBillClaim(ctx *models.Context, billClaim *models.BillClaim) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).InsertOne(ctx.CTX, billClaim)
	if err != nil {
		return err
	}
	billClaim.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveBillClaimV2(ctx *models.Context, billClaim *models.BillClaim) error {
	res, err := ctx.DB.Collection("billclaim2").InsertOne(ctx.CTX, billClaim)
	if err != nil {
		return err
	}
	billClaim.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleBillClaim : ""
func (d *Daos) GetSingleBillClaim(ctx *models.Context, uniqueID string) (*models.RefBillClaim, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBillClaimCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approvedBy", "userName", "ref.approvedUser", "ref.approvedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejectedBy", "userName", "ref.rejectedUser", "ref.rejectedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "userName", "ref.employee", "ref.employee")...)

	d.Shared.BsonToJSONPrintTag("BillClaim query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billClaims []models.RefBillClaim
	var billClaim *models.RefBillClaim
	if err = cursor.All(ctx.CTX, &billClaims); err != nil {
		return nil, err
	}
	if len(billClaims) > 0 {
		billClaim = &billClaims[0]
	}
	return billClaim, nil
}

//UpdateBillClaim : ""
func (d *Daos) UpdateBillClaim(ctx *models.Context, billClaim *models.BillClaim) error {
	selector := bson.M{"uniqueId": billClaim.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": billClaim}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableBillClaim :""
func (d *Daos) EnableBillClaim(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBillClaim :""
func (d *Daos) DisableBillClaim(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBillClaim :""
func (d *Daos) DeleteBillClaim(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterBillClaim : ""
func (d *Daos) FilterBillClaim(ctx *models.Context, filter *models.FilterBillClaim, pagination *models.Pagination) ([]models.RefBillClaim, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": filter.EmployeeId}})
		}
		if len(filter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationID}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approvedBy", "userName", "ref.approvedUser", "ref.approvedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejectedBy", "userName", "ref.rejectedUser", "ref.rejectedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "userName", "ref.employee", "ref.employee")...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBillClaimCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("BillClaim query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var billClaims []models.RefBillClaim
	if err = cursor.All(context.TODO(), &billClaims); err != nil {
		return nil, err
	}
	return billClaims, nil
}
func (d *Daos) ApprovedBillClaim(ctx *models.Context, approved *models.ReviewedBillClaim) error {
	query := bson.M{"uniqueId": approved.BillClaim}
	t := time.Now()
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMSTATUSAPPROVED, "approvedBy": approved.ReviewedBy, "approvedDate": &t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) RejectedBillClaim(ctx *models.Context, rejected *models.ReviewedBillClaim) error {
	query := bson.M{"uniqueId": rejected.BillClaim}
	t := time.Now()
	update := bson.M{"$set": bson.M{"status": constants.BILLCLAIMSTATUSREJECTED, "rejectedBy": rejected.ReviewedBy, "rejectedDate": &t, "remarks": rejected.Remarks}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBILLCLAIM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
