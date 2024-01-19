package daos

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SavePopup Notification :""
func (d *Daos) SavePopupNotification(ctx *models.Context, popup *models.PopupNotification) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).InsertOne(ctx.CTX, popup)
	return err
}

//GetSinglePopup Notification : ""
func (d *Daos) GetSinglePopupNotification(ctx *models.Context, uniqueID string) (*models.RefPopupNotification, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var popups []models.RefPopupNotification
	var popup *models.RefPopupNotification
	if err = cursor.All(ctx.CTX, &popups); err != nil {
		return nil, err
	}
	if len(popups) > 0 {
		popup = &popups[0]
	}
	return popup, nil
}

//UpdatePopup Notification : ""
func (d *Daos) UpdatePopupNotification(ctx *models.Context, popup *models.PopupNotification) error {
	selector := bson.M{"uniqueId": popup.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": popup}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnablePopup Notification:""
func (d *Daos) EnablePopupNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.POPUPNOTIFICATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePopup Notification :""
func (d *Daos) DisablePopupNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.POPUPNOTIFICATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePopup Notification :""
func (d *Daos) DeletePopupNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.POPUPNOTIFICATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterPopup Notification : ""
func (d *Daos) FilterPopupNotification(ctx *models.Context, PopupNotificationfilter *models.PopupNotificationFilter, pagination *models.Pagination) ([]models.RefPopupNotification, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if PopupNotificationfilter != nil {
		if len(PopupNotificationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": PopupNotificationfilter.Status}})
		}
		if len(PopupNotificationfilter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": PopupNotificationfilter.Type}})
		}
		if len(PopupNotificationfilter.IsPop) > 0 {
			query = append(query, bson.M{"ispop": bson.M{"$in": PopupNotificationfilter.IsPop}})
		}
		//SearchText
		if PopupNotificationfilter.SearchText.Title != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: PopupNotificationfilter.SearchText.Title, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("pkg query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var popups []models.RefPopupNotification
	if err = cursor.All(context.TODO(), &popups); err != nil {
		return nil, err
	}
	return popups, nil
}
func (d *Daos) GetDefaultPopupNotification(ctx *models.Context, Type string) ([]models.RefPopupNotification, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"type": bson.M{"$in": []string{"Common", Type}}, "ispop": true}})
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var popups []models.RefPopupNotification
	if err = cursor.All(ctx.CTX, &popups); err != nil {
		return nil, err
	}
	return popups, nil
}

func (d *Daos) SetPopupNotification(ctx *models.Context, popupnotification *models.Ispoptrue) error {
	//update popup list
	filter := bson.M{
		"type": bson.M{
			"$eq": popupnotification.Type,
		},
	}
	updatemany := bson.M{"$set": bson.M{"ispop": false}}
	_, errs := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateMany(ctx.CTX, filter, updatemany)
	query := bson.M{"uniqueId": popupnotification.UniqueID}
	update := bson.M{"$set": bson.M{"ispop": true}}
	_, id := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateOne(ctx.CTX, query, update)

	if errs != nil {
		return errors.New("Not Changed" + errs.Error())
	}
	return id

}

//EnablePopupNotificationV2:""
func (d *Daos) EnablePopupNotificationV2(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.POPUPNOTIFICATIONSTATUSACTIVE, "ispop": true}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePopupNotificationV2 :""
func (d *Daos) DisablePopupNotificationV2(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.POPUPNOTIFICATIONSTATUSDISABLED, "ispop": false}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPOPUPNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
