package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePreregistration :""
func (d *Daos) SavePreregistration(ctx *models.Context, preregistration *models.Preregistration) error {
	isUniqueMobile := validateMobileAtReg(ctx, preregistration.MobileNumber)
	if isUniqueMobile == false {
		return errors.New("this mobile number is already associated with other account")
	}
	isUniqueEmail := validateEmailAtReg(ctx, preregistration.Email)
	if isUniqueEmail == false {
		return errors.New("this email is already associated with other account")
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).InsertOne(ctx.CTX, preregistration)
	return err
}

//SavePreregistrationV2 :""
func (d *Daos) SavePreregistrationV2(ctx *models.Context, preregistration *models.Preregistration) error {

	t := time.Now()
	draftLog := models.PreregistrationTimeline{
		On:   &t,
		By:   preregistration.UniqueID,
		Type: constants.PREREGISTRATIONSTATUSDRAFT,
	}
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": preregistration.UniqueID}
	updateData := bson.M{"$set": preregistration, "$addToSet": bson.M{"log": draftLog}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil
}

//SubmitPreregistration :""
func (d *Daos) SubmitPreregistration(ctx *models.Context, preregistration *models.Preregistration) error {

	submitLog := models.PreregistrationTimeline{
		On:   preregistration.Submitted.On,
		By:   preregistration.Submitted.By,
		Type: constants.PREREGISTRATIONSTATUSSUBMITTED,
	}
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": preregistration.UniqueID}
	updateData := bson.M{"$set": preregistration, "$addToSet": bson.M{"log": submitLog}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil
}

//ReapplyPreregistration :""
func (d *Daos) ReapplyPreregistration(ctx *models.Context, preregistration *models.Preregistration) error {

	log := models.PreregistrationTimeline{
		On:   preregistration.Reapplied.On,
		By:   preregistration.Reapplied.By,
		Type: constants.PREREGISTRATIONSTATUSREAPPIED,
	}
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": preregistration.UniqueID}
	updateData := bson.M{"$set": preregistration, "$addToSet": bson.M{"log": log}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil
}

//GetSinglePreregistration : ""
func (d *Daos) GetSinglePreregistration(ctx *models.Context, mobileNumber string) (*models.RefPreregistration, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNumber": mobileNumber}})
	//Lookups
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPREREGISTRATION, "uniqueId")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAPPLICANTTYPE, "typeId", "uniqueId", "ref.applicantType", "ref.applicantType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var preregistrations []models.RefPreregistration
	var preregistration *models.RefPreregistration
	if err = cursor.All(ctx.CTX, &preregistrations); err != nil {
		return nil, err
	}
	if len(preregistrations) > 0 {
		preregistration = &preregistrations[0]
	}
	return preregistration, nil
}

//UpdatePreregistration : ""
func (d *Daos) UpdatePreregistration(ctx *models.Context, preregistration *models.Preregistration) error {
	selector := bson.M{"uniqueId": preregistration.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": preregistration, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPreregistration : ""
func (d *Daos) FilterPreregistration(ctx *models.Context, preregistrationfilter *models.PreregistrationFilter, pagination *models.Pagination) ([]models.RefPreregistration, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"ref.dayssince": bson.M{
			"$trunc": bson.M{
				"$divide": []interface{}{bson.M{"$subtract": []interface{}{time.Now(), "$created.on"}}, 1000 * 60 * 60 * 24},
			},
		},
	}})
	if preregistrationfilter.IsGetExpiredDraft {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"ref.dayssince": bson.M{"$gte": 5}}})
	}
	if preregistrationfilter != nil {
		if len(preregistrationfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": preregistrationfilter.UniqueID}})
		}
		if len(preregistrationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": preregistrationfilter.Status}})
		}
		if len(preregistrationfilter.ApplicantType) > 0 {
			query = append(query, bson.M{"typeId": bson.M{"$in": preregistrationfilter.ApplicantType}})
		}
		if preregistrationfilter.Address != nil {
			if len(preregistrationfilter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": preregistrationfilter.Address.StateCode}})
			}
			if len(preregistrationfilter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": preregistrationfilter.Address.DistrictCode}})
			}
			if len(preregistrationfilter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": preregistrationfilter.Address.VillageCode}})
			}
			if len(preregistrationfilter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": bson.M{"$in": preregistrationfilter.Address.ZoneCode}})
			}
			if len(preregistrationfilter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": bson.M{"$in": preregistrationfilter.Address.WardCode}})
			}
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		cursor, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).Aggregate(ctx.CTX, mainPipeline, nil)
		if err != nil {
			return nil, err
		}
		totalCount := cursor.RemainingBatchLength()
		fmt.Println("count", totalCount)
		pagination.Count = totalCount
		d.Shared.PaginationData(pagination)
		mainPipeline = append(mainPipeline, []bson.M{{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, {"$limit": pagination.Limit}}...)

		/*
			//Getting Total count
			totalCount, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).CountDocuments(ctx.CTX, func() bson.M {
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
		*/
	}
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAPPLICANTTYPE, "typeId", "uniqueId", "ref.applicantType", "ref.applicantType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("preregistration query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var preregistrations []models.RefPreregistration
	if err = cursor.All(context.TODO(), &preregistrations); err != nil {
		return nil, err
	}
	return preregistrations, nil
}

//EnablePreregistration :""
func (d *Daos) EnablePreregistration(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PREREGISTRATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePreregistration :""
func (d *Daos) DisablePreregistration(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PREREGISTRATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePreregistration :""
func (d *Daos) DeletePreregistration(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PREREGISTRATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//ValidateMobileNumber : ""
func (d *Daos) ValidateMobileNumber(ctx *models.Context, mobileNumber, uniqueID string) (bool, *models.RefPreregistration, error) {
	isValid := false
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNumber": mobileNumber, "uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return isValid, nil, err
	}
	var preregistrations []models.RefPreregistration
	var preregistration *models.RefPreregistration
	if err = cursor.All(ctx.CTX, &preregistrations); err != nil {
		return isValid, nil, err
	}

	if len(preregistrations) > 0 {
		preregistration = &preregistrations[0]
		isValid = true
	}
	return isValid, preregistration, nil
}

// ValidateMobileAtReg : "validating only mobile number at registration time"
func validateMobileAtReg(ctx *models.Context, mobile string) bool {
	findQuery := bson.M{"mobileNumber": mobile}
	var resDoc *models.Preregistration
	res := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).FindOne(ctx.CTX, findQuery, options.FindOne())
	err := res.Decode(&resDoc)
	fmt.Println("Decode error for result of prereg", err)
	fmt.Println("res of getting mobile number ", resDoc)
	if resDoc != nil {
		return false
	}
	return true
}

// ValidateEmailAtReg : "validating only email at registration time"
func validateEmailAtReg(ctx *models.Context, email string) bool {
	var resDoc *models.Preregistration
	findQuery := bson.M{"email": email}
	res := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).FindOne(ctx.CTX, findQuery, options.FindOne())
	err := res.Decode(&resDoc)
	fmt.Println("Decode error for result of prereg", err)

	fmt.Println("res of getting email", resDoc)
	if resDoc != nil {
		return false
	}
	return true
}

//ValidateEmailAtUpdate : ""
func (d *Daos) ValidateEmailAtUpdate(ctx *models.Context, email, uniqueID string) error {
	var resDoc *models.Preregistration
	findQuery := bson.M{"email": email}
	res := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).FindOne(ctx.CTX, findQuery, options.FindOne())
	err := res.Decode(&resDoc)
	fmt.Println("Decode error for result of prereg", err)

	fmt.Println("res of getting email at update", resDoc)
	if resDoc != nil {
		if resDoc.UniqueID != uniqueID {
			return errors.New("this email is not associated with this account")
		}
	}

	return nil
}

//ValidateMobileAtUpdate : ""
func (d *Daos) ValidateMobileAtUpdate(ctx *models.Context, mobile, uniqueID string) error {
	var resDoc *models.Preregistration
	findQuery := bson.M{"mobileNumber": mobile}
	res := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).FindOne(ctx.CTX, findQuery, options.FindOne())
	err := res.Decode(&resDoc)
	fmt.Println("Decode error for result of prereg", err)
	fmt.Println("res of getting mobile at update", resDoc)
	if resDoc != nil {
		if resDoc.UniqueID != uniqueID {
			return errors.New("this mobile is not associated with this account")
		}
	}
	return nil
}

//GetSinglePreregistrationWithUniqueID : ""
func (d *Daos) GetSinglePreregistrationWithUniqueID(ctx *models.Context, uniqueID string) (*models.RefPreregistration, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAPPLICANTTYPE, "typeId", "uniqueId", "ref.applicantType", "ref.applicantType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var preregistrations []models.RefPreregistration
	var preregistration *models.RefPreregistration
	if err = cursor.All(ctx.CTX, &preregistrations); err != nil {
		return nil, err
	}
	if len(preregistrations) > 0 {
		preregistration = &preregistrations[0]
	}
	return preregistration, nil
}

//ValidateUniquenessAtUpdateV2 : ""
func (d *Daos) ValidateUniquenessAtUpdateV2(ctx *models.Context, data, uniqueID, key string) (bool, error) {
	findQuery := bson.M{key: data, "uniqueId": bson.M{"$nin": []string{uniqueID}}}
	d.Shared.BsonToJSONPrintTag("unique query =>", findQuery)
	count, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).CountDocuments(ctx.CTX, findQuery, nil)
	if err != nil {
		return false, errors.New("Error in cunting documents - " + err.Error())
	}
	if count == 0 {
		fmt.Println("Not Duplicate")
		return true, nil
	}
	fmt.Println("Duplicate")
	return false, nil
}

//PreregistrationStatusChange : ""
func (d *Daos) PreregistrationStatusChange(ctx *models.Context, psc *models.PreregistrationStatusChange) error {
	query := bson.M{"uniqueId": psc.ApplicantID}
	log := models.PreregistrationTimeline{
		On:      psc.On,
		By:      psc.By,
		ByType:  psc.ByType,
		ByName:  psc.ByName,
		Remarks: psc.Remarks,
		Type:    psc.Status,
	}
	updateData := bson.M{"$set": bson.M{"status": psc.Status}, "$addToSet": bson.M{"log": log}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPREREGISTRATION).UpdateOne(ctx.CTX, query, updateData); err != nil {
		return errors.New("Error in updating status - " + err.Error())
	}
	return nil
}
