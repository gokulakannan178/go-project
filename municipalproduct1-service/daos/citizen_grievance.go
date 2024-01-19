package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveCitizenGrievance : ""
func (d *Daos) SaveCitizenGrievance(ctx *models.Context, mobile *models.CitizenGrievance) error {
	d.Shared.BsonToJSONPrint(mobile)
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).InsertOne(ctx.CTX, mobile)
	return err
}

// GetSingleCitizenGrievance : ""
func (d *Daos) GetSingleCitizenGrievance(ctx *models.Context, UniqueID string) (*models.RefCitizenGrievance, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var citizens []models.RefCitizenGrievance
	var citizen *models.RefCitizenGrievance
	if err = cursor.All(ctx.CTX, &citizens); err != nil {
		return nil, err
	}
	if len(citizens) > 0 {
		citizen = &citizens[0]
	}
	return citizen, nil
}

// UpdateCitizenGrievance : ""
func (d *Daos) UpdateCitizenGrievance(ctx *models.Context, mobile *models.CitizenGrievance) error {
	selector := bson.M{"uniqueId": mobile.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": mobile}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableCitizenGrievance : ""
func (d *Daos) EnableCitizenGrievance(ctx *models.Context, citizen *models.CitizenGrievance) error {
	query := bson.M{"uniqueId": citizen.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSSTATUSACTIVE,
		"activator": citizen.Activator}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableCitizenGrievance: ""
func (d *Daos) DisableCitizenGrievance(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteCitizenGrievance: ""
func (d *Daos) DeleteCitizenGrievance(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// CompletedCitizenGrievance: ""
func (d *Daos) CompletedCitizenGrievance(ctx *models.Context, citizen *models.CitizenGrievance) error {
	query := bson.M{"uniqueId": citizen.UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSSTATUSCOMPLETED,
		"completed": citizen.Completed}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// RejectedCitizenGrievance : ""
func (d *Daos) RejectedCitizenGrievance(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGRAVIANSSTATUSREJECTED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCitizenGrievance : ""
func (d *Daos) FilterCitizenGrievance(ctx *models.Context, filter *models.CitizenGrievanceFilter, pagination *models.Pagination) ([]models.RefCitizenGrievance, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Requestor) > 0 {
			query = append(query, bson.M{"requestor.by": bson.M{"$in": filter.Requestor}})
		}
		if len(filter.Activator) > 0 {
			query = append(query, bson.M{"activator.by": bson.M{"$in": filter.Activator}})
		}
		if len(filter.Completed) > 0 {
			query = append(query, bson.M{"completed.by": bson.M{"$in": filter.Completed}})
		}
		if len(filter.MobileNo) > 0 {
			query = append(query, bson.M{"mobileNo": bson.M{"$in": filter.MobileNo}})
		}
		// Requestor DateRange Filter
		//var sd,ed time.Time
		if filter.RequestedDate.From != nil {
			sd := time.Date(filter.RequestedDate.From.Year(), filter.RequestedDate.From.Month(), filter.RequestedDate.From.Day(), 0, 0, 0, 0, filter.RequestedDate.From.Location())
			ed := time.Date(filter.RequestedDate.From.Year(), filter.RequestedDate.From.Month(), filter.RequestedDate.From.Day(), 23, 59, 59, 0, filter.RequestedDate.From.Location())
			if filter.RequestedDate.To != nil {
				ed = time.Date(filter.RequestedDate.To.Year(), filter.RequestedDate.To.Month(), filter.RequestedDate.To.Day(), 23, 59, 59, 0, filter.RequestedDate.To.Location())
			}
			query = append(query, bson.M{"requestor.on": bson.M{"$gte": sd, "$lte": ed}})

		}

		// Activator Date Range Filter
		//var sd,ed time.Time
		if filter.ActivatedDate.From != nil {
			sd := time.Date(filter.ActivatedDate.From.Year(), filter.ActivatedDate.From.Month(), filter.ActivatedDate.From.Day(), 0, 0, 0, 0, filter.ActivatedDate.From.Location())
			ed := time.Date(filter.ActivatedDate.From.Year(), filter.ActivatedDate.From.Month(), filter.ActivatedDate.From.Day(), 23, 59, 59, 0, filter.ActivatedDate.From.Location())
			if filter.ActivatedDate.To != nil {
				ed = time.Date(filter.ActivatedDate.To.Year(), filter.ActivatedDate.To.Month(), filter.ActivatedDate.To.Day(), 23, 59, 59, 0, filter.ActivatedDate.To.Location())
			}
			query = append(query, bson.M{"activator.on": bson.M{"$gte": sd, "$lte": ed}})

		}

		// Completed Date Range Filter
		//var sd,ed time.Time
		if filter.CompletedDate.From != nil {
			sd := time.Date(filter.CompletedDate.From.Year(), filter.CompletedDate.From.Month(), filter.CompletedDate.From.Day(), 0, 0, 0, 0, filter.CompletedDate.From.Location())
			ed := time.Date(filter.CompletedDate.From.Year(), filter.CompletedDate.From.Month(), filter.CompletedDate.From.Day(), 23, 59, 59, 0, filter.CompletedDate.From.Location())
			if filter.CompletedDate.To != nil {
				ed = time.Date(filter.CompletedDate.To.Year(), filter.CompletedDate.To.Month(), filter.CompletedDate.To.Day(), 23, 59, 59, 0, filter.CompletedDate.To.Location())
			}
			query = append(query, bson.M{"completed.on": bson.M{"$gte": sd, "$lte": ed}})

		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"requestor.on": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var citizens []models.RefCitizenGrievance
	if err = cursor.All(context.TODO(), &citizens); err != nil {
		return nil, err
	}
	return citizens, nil
}

// UpdateCitizenGrievance : ""
func (d *Daos) UpdateCitizenGrievanceSolution(ctx *models.Context, solution *models.CitizenGrievanceSolution) error {
	selector := bson.M{"uniqueId": solution.CitizenGrievanceID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": bson.M{"solution": solution.Solution,
		"solutionImage": solution.SolutionImage,
		"solutionDate":  solution.SolutionDate,
		"status":        solution.Status,
		"solutionBy":    solution.By,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGRAVIANS).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
