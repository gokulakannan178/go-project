package daos

import (
	"context"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) DashboardUserCount(ctx *models.Context, userfilter *models.DashboardUserCountFilter) ([]models.DashboardUserCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.UserReportFilter(ctx, &userfilter.UserFilter)
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"usercreator": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPECONTENTCREATOR}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "usercreator"},
		},
		"selfRegistration": []bson.M{
			bson.M{"$match": bson.M{"isSelfRegistration": true, "status": bson.M{"$in": []string{constants.USERSTATUSINIT}}}},
			bson.M{"$count": "selfRegistration"},
		},
		"callCenterAgent": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPECALLCENTERAGENT}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "callCenterAgent"},
		},
		"userManager": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPECONTENTMANAGER}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "userManager"},
		},
		"userProvider": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPECONTENTPROVIDER}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "userProvider"},
		},
		"userDisseminator": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPECONTENTDISSEMINATOR}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "userDisseminator"},
		},
		"fieldAgent": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPEFIELDAGENT}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "fieldAgent"},
		},
		"languageTranslator": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPELANGUAGETRANSLATOR}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "languageTranslator"},
		},
		"languageApprover": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPETRANSLATIONAPPROVER}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "languageApprover"},
		},
		"management": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPEMANAGEMENT}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "management"},
		},
		"moderator": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPEMODERATOR}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "moderator"},
		},
		"subjectMatterExpert": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPESUBJECTMATTEREXPERT}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "subjectMatterExpert"},
		},
		"systemAdmin": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPESYSTEMADMIN}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "systemAdmin"},
		},
		"vistorViewer": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPEVISTORVIEWER}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "vistorViewer"},
		},
		"trainer": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPETRAINER}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "trainer"},
		},
		"fieldAgentLead": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPEFIELDAGENTLEAD}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "fieldAgentLead"},
		},
		"guestUser": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPEGUESTUSER}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "guestUser"},
		},
		"superAdmin": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPESUPERADMIN}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "superAdmin"},
		},
		"districtAdmin": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.USERTYPEDISTRICTADMIN}}, "status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "districtAdmin"},
		},
		"allUser": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{constants.USERSTATUSACTIVE, constants.USERSTATUSDISABLED}}}},
			bson.M{"$count": "allUser"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"usercreator":         bson.M{"$arrayElemAt": []interface{}{"$usercreator", 0}},
			"selfRegistration":    bson.M{"$arrayElemAt": []interface{}{"$selfRegistration", 0}},
			"callCenterAgent":     bson.M{"$arrayElemAt": []interface{}{"$callCenterAgent", 0}},
			"userManager":         bson.M{"$arrayElemAt": []interface{}{"$userManager", 0}},
			"userProvider":        bson.M{"$arrayElemAt": []interface{}{"$userProvider", 0}},
			"userDisseminator":    bson.M{"$arrayElemAt": []interface{}{"$userDisseminator", 0}},
			"fieldAgent":          bson.M{"$arrayElemAt": []interface{}{"$fieldAgent", 0}},
			"languageTranslator":  bson.M{"$arrayElemAt": []interface{}{"$languageTranslator", 0}},
			"languageApprover":    bson.M{"$arrayElemAt": []interface{}{"$languageApprover", 0}},
			"management":          bson.M{"$arrayElemAt": []interface{}{"$management", 0}},
			"moderator":           bson.M{"$arrayElemAt": []interface{}{"$moderator", 0}},
			"subjectMatterExpert": bson.M{"$arrayElemAt": []interface{}{"$subjectMatterExpert", 0}},
			"systemAdmin":         bson.M{"$arrayElemAt": []interface{}{"$systemAdmin", 0}},
			"vistorViewer":        bson.M{"$arrayElemAt": []interface{}{"$vistorViewer", 0}},
			"trainer":             bson.M{"$arrayElemAt": []interface{}{"$trainer", 0}},
			"fieldAgentLead":      bson.M{"$arrayElemAt": []interface{}{"$fieldAgentLead", 0}},
			"guestUser":           bson.M{"$arrayElemAt": []interface{}{"$guestUser", 0}},
			"superAdmin":          bson.M{"$arrayElemAt": []interface{}{"$superAdmin", 0}},
			"districtAdmin":       bson.M{"$arrayElemAt": []interface{}{"$districtAdmin", 0}},
			"allUser":             bson.M{"$arrayElemAt": []interface{}{"$allUser", 0}},
		}},
		bson.M{"$addFields": bson.M{
			"contentcreator":      "$usercreator.usercreator",
			"selfRegistration":    "$selfRegistration.selfRegistration",
			"callCenterAgent":     "$callCenterAgent.callCenterAgent",
			"contentManager":      "$userManager.userManager",
			"contentProvider":     "$userProvider.userProvider",
			"contentDisseminator": "$userDisseminator.userDisseminator",
			"fieldAgent":          "$fieldAgent.fieldAgent",
			"languageTranslator":  "$languageTranslator.languageTranslator",
			"languageApprover":    "$languageApprover.languageApprover",
			"management":          "$management.management",
			"moderator":           "$moderator.moderator",
			"subjectMatterExpert": "$subjectMatterExpert.subjectMatterExpert",
			"systemAdmin":         "$systemAdmin.systemAdmin",
			"vistorViewer":        "$vistorViewer.vistorViewer",
			"trainer":             "$trainer.trainer",
			"fieldAgentLead":      "$fieldAgentLead.fieldAgentLead",
			"guestUser":           "$guestUser.guestUser",
			"superAdmin":          "$superAdmin.superAdmin",
			"districtAdmin":       "$districtAdmin.districtAdmin",
			"allUser":             "$allUser.allUser",
		}})

	d.Shared.BsonToJSONPrintTag("DashboardUserCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbucr []models.DashboardUserCountReport
	if err := cursor.All(ctx.CTX, &dbucr); err != nil {
		return nil, err
	}
	return dbucr, nil

}
func (d *Daos) DayWiseUserDemandChart(ctx *models.Context, userfilter *models.DashboardUserCountFilter) (*models.DayWiseUserDemandChartReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.UserReportFilter(ctx, &userfilter.UserFilter)
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"day": bson.M{"$dayOfMonth": "$createdDate"}, "status": "$status"},
		"count": bson.M{"$sum": 1}}})
	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "mobileTowerCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "data": bson.M{"$push": bson.M{"k": "$_id.status", "v": "$count"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data": bson.M{"$arrayToObject": "$data"}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{userfilter.CreatedFrom.StartDate.Day(), userfilter.CreatedFrom.EndDate.Day() + 1, 1}},
				"as":    "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days._id", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"_id": "$$rangeDay",
									"data": bson.M{
										"Active":   0.0,
										"Disabled": 0.0,
									}},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})
	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var user *models.DayWiseUserDemandChartReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return user, err
	}
	var data []models.DayWiseUserDemandChartReport

	if err = cursor.All(context.TODO(), &data); err != nil {
		return user, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return user, nil

}
