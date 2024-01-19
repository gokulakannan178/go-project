package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveDealer :""
func (d *Daos) SaveDealer(ctx *models.Context, dealer *models.Dealer) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDEALER).InsertOne(ctx.CTX, dealer)
	if err != nil {
		return err
	}
	dealer.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleDealer : ""
func (d *Daos) GetSingleDealer(ctx *models.Context, code string) (*models.RefDealer, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEALER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dealers []models.RefDealer
	var dealer *models.RefDealer
	if err = cursor.All(ctx.CTX, &dealers); err != nil {
		return nil, err
	}
	if len(dealers) > 0 {
		dealer = &dealers[0]
	}
	return dealer, nil
}

//UpdateDealer : ""
func (d *Daos) UpdateDealer(ctx *models.Context, dealer *models.Dealer) error {
	selector := bson.M{"_id": dealer.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dealer, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEALER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDealer : ""
func (d *Daos) FilterDealer(ctx *models.Context, filter *models.DealerFilter, pagination *models.Pagination) ([]models.RefDealer, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": filter.ActiveStatus}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": filter.GramPanchayat}})
		}
		if len(filter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": filter.Village}})
		}
		if len(filter.CertificationStatus) > 0 {
			query = append(query, bson.M{"certification.status": bson.M{"$in": filter.CertificationStatus}})
		}
		if filter.CertificationAppliedDateRange != nil {
			//var sd,ed time.Time
			if filter.CertificationAppliedDateRange.From != nil {
				sd := time.Date(filter.CertificationAppliedDateRange.From.Year(), filter.CertificationAppliedDateRange.From.Month(), filter.CertificationAppliedDateRange.From.Day(), 0, 0, 0, 0, filter.CertificationExpiryDateRange.From.Location())
				ed := time.Date(filter.CertificationAppliedDateRange.From.Year(), filter.CertificationAppliedDateRange.From.Month(), filter.CertificationAppliedDateRange.From.Day(), 23, 59, 59, 0, filter.CertificationExpiryDateRange.From.Location())
				if filter.CertificationAppliedDateRange.To != nil {
					ed = time.Date(filter.CertificationAppliedDateRange.To.Year(), filter.CertificationAppliedDateRange.To.Month(), filter.CertificationAppliedDateRange.To.Day(), 23, 59, 59, 0, filter.CertificationAppliedDateRange.To.Location())
				}
				query = append(query, bson.M{"certification.appliedDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.CertificationExpiryDateRange != nil {
			//var sd,ed time.Time
			if filter.CertificationExpiryDateRange.From != nil {
				sd := time.Date(filter.CertificationExpiryDateRange.From.Year(), filter.CertificationExpiryDateRange.From.Month(), filter.CertificationExpiryDateRange.From.Day(), 0, 0, 0, 0, filter.CertificationExpiryDateRange.From.Location())
				ed := time.Date(filter.CertificationExpiryDateRange.From.Year(), filter.CertificationExpiryDateRange.From.Month(), filter.CertificationExpiryDateRange.From.Day(), 23, 59, 59, 0, filter.CertificationExpiryDateRange.From.Location())
				if filter.CertificationExpiryDateRange.To != nil {
					ed = time.Date(filter.CertificationExpiryDateRange.To.Year(), filter.CertificationExpiryDateRange.To.Month(), filter.CertificationExpiryDateRange.To.Day(), 23, 59, 59, 0, filter.CertificationExpiryDateRange.To.Location())
				}
				query = append(query, bson.M{"certification.expiryDate": bson.M{"$gte": sd, "$lte": ed}})

			}
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

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDEALER).CountDocuments(ctx.CTX, func() bson.M {
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
	// //lookup
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Dealer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEALER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dealers []models.RefDealer
	if err = cursor.All(context.TODO(), &dealers); err != nil {
		return nil, err
	}
	return dealers, nil
}

//EnableDealer :""
func (d *Daos) EnableDealer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DEALERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDEALER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDealer :""
func (d *Daos) DisableDealer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DEALERSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDEALER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDealer :""
func (d *Daos) DeleteDealer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DEALERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDEALER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleUserWithMobileNo : ""
func (d *Daos) GetSingleDealerWithMobileNo(ctx *models.Context, mobile string) (*models.RefDealer, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": mobile}})

	d.Shared.BsonToJSONPrintTag("dealer query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEALER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dealers []models.RefDealer
	var dealer *models.RefDealer
	if err = cursor.All(ctx.CTX, &dealers); err != nil {
		return nil, err
	}
	if len(dealers) > 0 {
		dealer = &dealers[0]
	}
	return dealer, nil
}

func (d *Daos) DealerUniquenessCheckRegistration(ctx *models.Context, param string, value string) (*models.DealerUniquinessChk, error) {

	query := bson.M{
		param: value,
	}
	fmt.Println("query====>", query)

	result := new(models.DealerUniquinessChk)
	var Dealer *models.Dealer
	if err := ctx.DB.Collection(constants.COLLECTIONDEALER).FindOne(ctx.CTX, query).Decode(&Dealer); err != nil {
		return nil, err
	}
	fmt.Println("Dealerp===>", Dealer)
	if Dealer == nil {
		result.Success = true
		return result, nil
	}
	if Dealer.Status == constants.DEALERSTATUSACTIVE {
		result.Success = false
		result.Message = "Dealer alreasy exist"
		return result, nil
	}
	if Dealer.Status == constants.DEALERSTATUSDISABLED {
		result.Success = false
		result.Message = "Dealer disabled - please contact Administator"
		return result, nil
	}
	if Dealer.Status == constants.DEALERSTATUSINIT {
		result.Success = false
		result.Message = "Awaiting Activation - please contact administator"
		return result, nil
	}
	return nil, errors.New("No options Availables")
}

func (d *Daos) DealerNearBy(ctx *models.Context, dealernb *models.NearBy, pagination *models.Pagination) ([]models.RefDealer, error) {
	coordinater := []float64{dealernb.Longitude, dealernb.Latitude}

	mainPipeline := []bson.M{}
	query := []bson.M{{"$geoNear": bson.M{"near": bson.M{"type": "Point", "coordinates": coordinater},
		"maxDistance":   dealernb.KM * 1000,
		"distanceField": "dist.calculated",
		"spherical":     true,
		"includeLocs":   "dist.location",
	},
	}}
	fmt.Println(query)
	//if d.Shared.GetCmdArg(constants.ENV) != "development" {
	mainPipeline = append(mainPipeline, query...)

	//}
	// if len(ulbnb.CertificateStatus) > 0 {
	//mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"testcert.status": "Active"}})

	// }
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count

		Count, err := GetCountForAggregation(ctx, query, constants.COLLECTIONDEALER)
		if err != nil {
			log.Println(err)
		}
		pagination.Count = Count
		d.Shared.PaginationData(pagination)
	}
	// LookUps
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)

	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	// //get GP
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "address.gramPanchayatCode", "code", "ref.address.gramPanchayat", "ref.address.gramPanchayat")...)
	// //get block
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "address.blockCode", "code", "ref.address.block", "ref.address.block")...)
	// //get district
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	// //Get Inventory
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULBINVENTORY, "uniqueId", "companyId", "ref.inventory", "ref.inventory")...)
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEALER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return nil, err
	}
	var dealers []models.RefDealer
	if err = cursor.All(context.TODO(), &dealers); err != nil {
		return nil, err
	}
	//	fmt.Println(ulbs)
	return dealers, err
}

func GetCountForAggregation(ctx *models.Context, query []primitive.M, Collection string) (int, error) {
	mainPipeline2 := []bson.M{}

	mainPipeline2 = append(mainPipeline2, query...)

	//Getting Total count
	group := []bson.M{{"$group": bson.M{"_id": nil, "mycount": bson.M{"$sum": 1}}}}
	mainPipeline2 = append(mainPipeline2, group...)
	type countvaule struct {
		Mycount int ` bson:"mycount"`
	}
	var totalCountV []countvaule
	totalCount, err := ctx.DB.Collection(Collection).Aggregate(ctx.CTX, mainPipeline2, nil)
	if err != nil {
		log.Println("Error in geting pagination count")
	}
	if err = totalCount.All(context.TODO(), &totalCountV); err != nil {
		return 0, err
	}
	fmt.Println("count", totalCount)
	if len(totalCountV) > 0 {
		return int(totalCountV[0].Mycount), nil
	}

	return 0, nil

}
func (d *Daos) DealerCertificationApply(ctx *models.Context, dealerID string, certification *models.DealerCertification) error {
	id, err := primitive.ObjectIDFromHex(dealerID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"certification.status": constants.DEALERCERTIFICATIONSTATUSAPPLY, "certification.actionDate": certification.AppliedDate}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDEALER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) DealerCertificationApprove(ctx *models.Context, dealerID string, certification *models.DealerCertification) error {
	id, err := primitive.ObjectIDFromHex(dealerID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"certification.status": constants.DEALERCERTIFICATIONSTATUSAPPROVED, "certification.expiryDate": certification.ExpiryDate, "certification.actionDate": certification.ActionDate}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDEALER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) DealerCertificationReject(ctx *models.Context, dealerID string, certification *models.DealerCertification) error {
	id, err := primitive.ObjectIDFromHex(dealerID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"certification.status": constants.DEALERCERTIFICATIONSTATUSREJECT, "certification.actionDate": certification.ActionDate}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDEALER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
