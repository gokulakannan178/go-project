package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

//SaveApplicant :""
func (d *Daos) SaveApplicant(ctx *models.Context, Applicant *models.Applicant) error {
	Applicant.ID = primitive.NewObjectID()
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).InsertOne(ctx.CTX, Applicant)
	return err
}

//GetSingleApplicant : ""
func (d *Daos) GetSingleApplicant(ctx *models.Context, UserName string) (*models.RefApplicant, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userName": UserName}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var applicants []models.RefApplicant
	var Applicant *models.RefApplicant
	if err = cursor.All(ctx.CTX, &applicants); err != nil {
		return nil, err
	}
	if len(applicants) > 0 {
		Applicant = &applicants[0]
	}
	return Applicant, nil
}

//UpdateApplicant : ""
func (d *Daos) UpdateApplicant(ctx *models.Context, Applicant *models.Applicant) error {
	selector := bson.M{"userName": Applicant.UserName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Applicant, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterApplicant : ""
func (d *Daos) FilterApplicant(ctx *models.Context, applicantfilter *models.ApplicantFilter, pagination *models.Pagination) ([]models.RefApplicant, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if applicantfilter != nil {

		if len(applicantfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": applicantfilter.Status}})
		}
		if len(applicantfilter.ApplicantType) > 0 {
			query = append(query, bson.M{"typeId": bson.M{"$in": applicantfilter.ApplicantType}})
		}
		if applicantfilter.Address != nil {
			if len(applicantfilter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": applicantfilter.Address.StateCode}})
			}
			if len(applicantfilter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": applicantfilter.Address.DistrictCode}})
			}
			if len(applicantfilter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": applicantfilter.Address.VillageCode}})
			}
			if len(applicantfilter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": bson.M{"$in": applicantfilter.Address.ZoneCode}})
			}
			if len(applicantfilter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": bson.M{"$in": applicantfilter.Address.WardCode}})
			}
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Applicant query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var applicants []models.RefApplicant
	if err = cursor.All(context.TODO(), &applicants); err != nil {
		return nil, err
	}
	return applicants, nil
}

//EnableApplicant :""
func (d *Daos) EnableApplicant(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableApplicant :""
func (d *Daos) DisableApplicant(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteApplicant :""
func (d *Daos) DeleteApplicant(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//BlacklistApplicant :""
func (d *Daos) BlacklistApplicant(ctx *models.Context, asc *models.ApplicantStatusChange) error {
	query := bson.M{"userName": asc.UserName}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTSTATUSBLACKLISTED, "remarks": asc.Info.Remarks}, "$push": bson.M{"applicantLog": asc.Info}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//LicenseCancelApplicant :""
func (d *Daos) LicenseCancelApplicant(ctx *models.Context, asc *models.ApplicantStatusChange) error {
	query := bson.M{"userName": asc.UserName}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTSTATUSLICENSECANCELLED, "remarks": asc.Info.Remarks}, "$push": bson.M{"applicantLog": asc.Info}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//ReActivateApplicant :""
func (d *Daos) ReActivateApplicant(ctx *models.Context, asc *models.ApplicantStatusChange) error {
	query := bson.M{"userName": asc.UserName}
	update := bson.M{"$set": bson.M{"status": constants.APPLICANTSTATUSACTIVE, "remarks": asc.Info.Remarks}, "$push": bson.M{"applicantLog": asc.Info}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAPPLICANT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
