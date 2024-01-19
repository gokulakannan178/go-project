package daos

import (
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

//GetSinglewhatsapptemplate : ""
func (d *Daos) GetSingleWhatsappTemplateWithName(ctx *models.Context, UniqueID string) (*models.SendWhatsAppText, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWHATSAPPTEMPLATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var whatsapptemplates []models.SendWhatsAppText
	var whatsapptemplate *models.SendWhatsAppText
	if err = cursor.All(ctx.CTX, &whatsapptemplates); err != nil {
		return nil, err
	}
	if len(whatsapptemplates) > 0 {
		whatsapptemplate = &whatsapptemplates[0]
	}
	return whatsapptemplate, nil
}
