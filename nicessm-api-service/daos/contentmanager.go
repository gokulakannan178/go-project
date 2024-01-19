package daos

import (
	"context"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//ContentManagerCount : ""
func (d *Daos) ContentManagerCount(ctx *models.Context, content *models.ContentFilter) ([]models.ContentCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = d.ContentFilter(ctx, content)
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"P": []bson.M{
			bson.M{"$count": "Total"},
		},
		"smsCount": []bson.M{
			bson.M{"$match": bson.M{"type": "S", "status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}},
			bson.M{"$count": "smsCount"},
		},
		"organicSmsCount": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"type": "S"}, bson.M{"smsContentType": "O"}, bson.M{"status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}}}},
			bson.M{"$count": "organicSmsCount"},
		},
		"inOrganicSmsCount": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"type": "S"}, bson.M{"smsContentType": "IO"}, bson.M{"status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}}}},
			bson.M{"$count": "inOrganicSmsCount"},
		},
		"monitoringSmsCount": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"type": "S"}, bson.M{"smsType": "M"}, bson.M{"status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}}}},
			bson.M{"$count": "monitoringSmsCount"},
		},
		"voiceSmsCount": []bson.M{
			bson.M{"$match": bson.M{"type": "V", "status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}},
			bson.M{"$count": "voiceSmsCount"},
		},
		"organicVoiceSmsCount": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"type": "V"}, bson.M{"smsContentType": "O"}, bson.M{"status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}}}},
			bson.M{"$count": "organicVoiceSmsCount"},
		},
		"inOrganicVoiceSmsCount": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"type": "V"}, bson.M{"smsContentType": "IO"}, bson.M{"status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}}}},
			bson.M{"$count": "inOrganicVoiceSmsCount"},
		},
		"monitoringVoiceSmsCount": []bson.M{
			bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"type": "V"}, bson.M{"smsType": "M"}, bson.M{"status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}}}},
			bson.M{"$count": "monitoringVoiceSmsCount"},
		},
		"onePageHtmlCount": []bson.M{
			bson.M{"$match": bson.M{"type": "P", "status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}},
			bson.M{"$count": "onePageHtmlCount"},
		},
		"documentCount": []bson.M{
			bson.M{"$match": bson.M{"type": "D", "status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}},
			bson.M{"$count": "documentCount"},
		},
		"videoUrlCount": []bson.M{
			bson.M{"$match": bson.M{"type": "U", "status": bson.M{"$nin": []string{constants.CONTENTCOMMENTSTATUSDELETED}}}},
			bson.M{"$count": "videoUrlCount"},
		},
	}},
		bson.M{"$project": bson.M{
			"Total":                   bson.M{"$arrayElemAt": []interface{}{"$Total.Total", 0}},
			"smsCount":                bson.M{"$arrayElemAt": []interface{}{"$smsCount.smsCount", 0}},
			"organicSmsCount":         bson.M{"$arrayElemAt": []interface{}{"$organicSmsCount.organicSmsCount", 0}},
			"inOrganicSmsCount":       bson.M{"$arrayElemAt": []interface{}{"$inOrganicSmsCount.inOrganicSmsCount", 0}},
			"monitoringSmsCount":      bson.M{"$arrayElemAt": []interface{}{"$monitoringSmsCount.monitoringSmsCount", 0}},
			"voiceSmsCount":           bson.M{"$arrayElemAt": []interface{}{"$voiceSmsCount.voiceSmsCount", 0}},
			"organicVoiceSmsCount":    bson.M{"$arrayElemAt": []interface{}{"$organicVoiceSmsCount.organicVoiceSmsCount", 0}},
			"inOrganicVoiceSmsCount":  bson.M{"$arrayElemAt": []interface{}{"$inOrganicVoiceSmsCount.inOrganicVoiceSmsCount", 0}},
			"monitoringVoiceSmsCount": bson.M{"$arrayElemAt": []interface{}{"$monitoringVoiceSmsCount.monitoringVoiceSmsCount", 0}},
			"onePageHtmlCount":        bson.M{"$arrayElemAt": []interface{}{"$onePageHtmlCount.onePageHtmlCount", 0}},
			"documentCount":           bson.M{"$arrayElemAt": []interface{}{"$documentCount.documentCount", 0}},
			"videoUrlCount":           bson.M{"$arrayElemAt": []interface{}{"$videoUrlCount.videoUrlCount", 0}},
		}},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("ContentManagerCount query =>", mainPipeline)

	var data []models.ContentCount

	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTNAME).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
