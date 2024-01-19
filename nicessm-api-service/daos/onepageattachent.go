package daos

import (
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetSingleDisease : ""
func (d *Daos) GetSingleOnePageAttachment(ctx *models.Context, UniqueID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return "", err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONEPAGEATTACHMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return "", err
	}
	var OnePageAttachments []models.RefOnePageAttachment
	var OnePageAttachment *models.RefOnePageAttachment
	if err = cursor.All(ctx.CTX, &OnePageAttachments); err != nil {
		return "", err
	}
	if len(OnePageAttachments) > 0 {
		OnePageAttachment = &OnePageAttachments[0]
	}
	OnePageAttachmentImg := OnePageAttachment.Image
	return OnePageAttachmentImg, nil
}
