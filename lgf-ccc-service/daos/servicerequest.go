package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceRequest : ""
func (d *Daos) SaveServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).InsertOne(ctx.CTX, serviceRequest)
	if err != nil {
		return err
	}
	return nil
}

// GetSingleServiceRequest : ""
func (d *Daos) GetSingleServiceRequest(ctx *models.Context, uniqueID string) (*models.RefServiceRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ServiceRequests []models.RefServiceRequest
	var ServiceRequest *models.RefServiceRequest
	if err = cursor.All(ctx.CTX, &ServiceRequests); err != nil {
		return nil, err
	}
	if len(ServiceRequests) > 0 {
		ServiceRequest = &ServiceRequests[0]
	}
	return ServiceRequest, err
}

//UpdateServiceRequest : ""
func (d *Daos) UpdateServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	selector := bson.M{"uniqueId": serviceRequest.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": serviceRequest}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableServiceRequest : ""
func (d *Daos) EnableServiceRequest(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SERVICEREQUESTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableServiceRequest : ""
func (d *Daos) DisableServiceRequest(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SERVICEREQUESTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteServiceRequest :""
func (d *Daos) DeleteServiceRequest(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SERVICEREQUESTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// EnableServiceRequest : ""
func (d *Daos) InProgressServiceRequest(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SERVICEREQUESTSTATUSINPROGRESS}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableServiceRequest : ""
func (d *Daos) PendingServiceRequest(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SERVICEREQUESTSTATUSPENDING}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteServiceRequest :""
func (d *Daos) InitServiceRequest(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SERVICEREQUESTSTATUSINIT}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteServiceRequest :""
func (d *Daos) CompletedServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	selector := bson.M{"uniqueId": serviceRequest.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{
		"closingNotes":     serviceRequest.ClosingNotes,
		"solutionImageUrl": serviceRequest.SolutionImageurl,
		"completionDate":   &t,
		"status":           constants.SERVICEREQUESTSTATUSCOMPLETED,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err

}

// FilterServiceRequest : ""
func (d *Daos) FilterServiceRequest(ctx *models.Context, serviceRequest *models.FilterServiceRequest, pagination *models.Pagination) ([]models.RefServiceRequest, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if serviceRequest != nil {
		if len(serviceRequest.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": serviceRequest.Status}})
		}
		if len(serviceRequest.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": serviceRequest.Name}})
		}
		if len(serviceRequest.GCID) > 0 {
			query = append(query, bson.M{"gcUser.id": bson.M{"$in": serviceRequest.GCID}})
		}
		if len(serviceRequest.ManagerID) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": serviceRequest.ManagerID}})
		}
		if len(serviceRequest.CitizenID) > 0 {
			query = append(query, bson.M{"citizen.id": bson.M{"$in": serviceRequest.CitizenID}})
		}
		//Regex
		if serviceRequest.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: serviceRequest.Regex.Name, Options: "xi"}})
		}
		if serviceRequest.Regex.ManagerName != "" {
			query = append(query, bson.M{"minUser.name": primitive.Regex{Pattern: serviceRequest.Regex.ManagerName, Options: "xi"}})
		}
		if serviceRequest.Regex.GCName != "" {
			query = append(query, bson.M{"gcUser.name": primitive.Regex{Pattern: serviceRequest.Regex.GCName, Options: "xi"}})
		}
		if serviceRequest.Regex.CitizenName != "" {
			query = append(query, bson.M{"citizen.name": primitive.Regex{Pattern: serviceRequest.Regex.CitizenName, Options: "xi"}})
		}
	}
	if serviceRequest.DateRange.From != nil {
		t := *serviceRequest.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if serviceRequest.DateRange.To != nil {
			t2 := *serviceRequest.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"requstedDate": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if serviceRequest != nil {
		if serviceRequest.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{serviceRequest.SortBy: serviceRequest.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).CountDocuments(ctx.CTX, func() bson.M {
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
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ServiceRequest []models.RefServiceRequest
	if err = cursor.All(context.TODO(), &ServiceRequest); err != nil {
		return nil, err
	}
	return ServiceRequest, nil
}

func (d *Daos) GetDetailServiceRequest(ctx *models.Context, uniqueID string) (*models.RefServiceRequest, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": constants.SERVICEREQUESTSTATUSCOMPLETED}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ServiceRequests []models.RefServiceRequest
	var ServiceRequest *models.RefServiceRequest
	if err = cursor.All(ctx.CTX, &ServiceRequests); err != nil {
		return nil, err
	}
	if len(ServiceRequests) > 0 {
		ServiceRequest = &ServiceRequests[0]
	}
	return ServiceRequest, err
}
func (d *Daos) AssignServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	selector := bson.M{"uniqueId": serviceRequest.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{
		"minUser":    serviceRequest.MinUser,
		"gcUser":     serviceRequest.GCUser,
		"assignDate": &t,
		"status":     constants.SERVICEREQUESTSTATUSPENDING,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSERVICEREQUEST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
