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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SavePolicyRule : ""
func (d *Daos) SavePolicyRule(ctx *models.Context, policyRule *models.PolicyRule) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).InsertOne(ctx.CTX, policyRule)
	if err != nil {
		return err
	}
	policyRule.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdatePolicyRule : ""
func (d *Daos) UpdatePolicyRule(ctx *models.Context, policyRule *models.PolicyRule) error {
	selector := bson.M{"uniqueId": policyRule.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": policyRule, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) UpdatePolicyRuleWithUpsert(ctx *models.Context, policyRule *models.PolicyRule) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"leavepolicyId": policyRule.LeavePolicyID, "leavemasterId": policyRule.LeaveMasterID}
	fmt.Println("updateQuery===>", updateQuery)
	updateData := bson.M{"$set": policyRule}
	if _, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil
}

// GetSinglePolicyRule : ""
func (d *Daos) GetSinglePolicyRule(ctx *models.Context, uniqueID string) (*models.RefPolicyRule, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leavemasterId", "uniqueId", "ref.leavemasterId", "ref.leavemasterId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEPOLICY, "leavepolicyId", "uniqueId", "ref.leavepolicyId", "ref.leavepolicyId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var policyRules []models.RefPolicyRule
	var policyRule *models.RefPolicyRule
	if err = cursor.All(ctx.CTX, &policyRules); err != nil {
		return nil, err
	}
	if len(policyRules) > 0 {
		policyRule = &policyRules[0]
	}
	return policyRule, err
}

// PolicyRuleRemoveNotPresentValue : ""
func (d *Daos) PolicyRuleRemoveNotPresentValue(ctx *models.Context, leavepolicyId string, arrayValue string) error {
	selector := bson.M{"leavepolicyId": leavepolicyId, "leavemasterId": bson.M{"$nin": arrayValue}}
	d.Shared.BsonToJSONPrintTag("selector query in policy rule =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.POLICYRULESTATUSDELETED}}
	d.Shared.BsonToJSONPrintTag("data query in policy rule =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).UpdateMany(ctx.CTX, selector, data)
	return err
}

// PolicyRuleUpsert : ""
func (d *Daos) PolicyRuleUpsert(ctx *models.Context, leavepolicyId string, arrayValue string, name string) error {
	fmt.Println("arrayValue", arrayValue)
	for _, v := range arrayValue {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"leavepolicyId": leavepolicyId, "leavemasterId": v}
		fmt.Println("updateQuery===>", updateQuery)
		policyrule := new(models.PolicyRule)
		policyrule.Status = constants.POLICYRULESTATUSACTIVE
		policyrule.Name = name
		policyrule.UniqueID = d.GetUniqueID(ctx, constants.COLLECTIONPOLICYRULE)
		fmt.Println("present added =======>", policyrule.UniqueID)
		updateData := bson.M{"$set": policyrule}
		if _, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

// EnablePolicyRule : ""
func (d *Daos) EnablePolicyRule(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.POLICYRULESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisablePolicyRule : ""
func (d *Daos) DisablePolicyRule(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.POLICYRULESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeletePolicyRule :""
func (d *Daos) DeletePolicyRule(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.POLICYRULESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPolicyRule : ""
func (d *Daos) FilterPolicyRule(ctx *models.Context, policyRule *models.FilterPolicyRule, pagination *models.Pagination) ([]models.RefPolicyRule, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if policyRule != nil {
		if len(policyRule.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": policyRule.Status}})
		}

		//Regex
		if policyRule.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: policyRule.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leavemasterId", "uniqueId", "ref.leavemasterId", "ref.leavemasterId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEPOLICY, "leavepolicyId", "uniqueId", "ref.leavepolicyId", "ref.leavepolicyId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPOLICYRULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var policyruleFilter []models.RefPolicyRule
	if err = cursor.All(context.TODO(), &policyruleFilter); err != nil {
		return nil, err
	}
	return policyruleFilter, nil
}
