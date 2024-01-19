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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveBirthCertificate : ""
func (d *Daos) SaveBirthCertificate(ctx *models.Context, birthcertificate *models.BirthCertificate) error {
	d.Shared.BsonToJSONPrint(birthcertificate)
	_, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).InsertOne(ctx.CTX, birthcertificate)
	return err
}

// GetSingleBirthCertificate : ""
func (d *Daos) GetSingleBirthCertificate(ctx *models.Context, UniqueID string) (*models.RefBirthCertificate, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONHOSPITAL, "hospitalId", "uniqueId", "ref.hospital", "ref.hospital")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBIRTHCERTIFICATESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var birthcertificates []models.RefBirthCertificate
	var birthcertificate *models.RefBirthCertificate
	if err = cursor.All(ctx.CTX, &birthcertificates); err != nil {
		return nil, err
	}
	if len(birthcertificates) > 0 {
		birthcertificate = &birthcertificates[0]
	}
	return birthcertificate, nil
}

// UpdateBirthCertificate : ""
func (d *Daos) UpdateBirthCertificate(ctx *models.Context, birthcertificate *models.BirthCertificate) error {
	selector := bson.M{"uniqueId": birthcertificate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": birthcertificate}
	_, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableBirthCertificate : ""
func (d *Daos) EnableBirthCertificate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BIRTHCERTIFICATESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableBirthCertificate : ""
func (d *Daos) DisableBirthCertificate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BIRTHCERTIFICATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteBirthCertificate : ""
func (d *Daos) DeleteBirthCertificate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BIRTHCERTIFICATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterBirthCertificate : ""
func (d *Daos) FilterBirthCertificate(ctx *models.Context, filter *models.BirthCertificateFilter, pagination *models.Pagination) ([]models.RefBirthCertificate, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueIDs}})
		}
		if len(filter.HospitalID) > 0 {
			query = append(query, bson.M{"hospitalId": bson.M{"$in": filter.HospitalID}})
		}
		if len(filter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": filter.Name}})
		}
		if len(filter.FatherName) > 0 {
			query = append(query, bson.M{"fatherName": bson.M{"$in": filter.FatherName}})
		}
		if len(filter.Gender) > 0 {
			query = append(query, bson.M{"gender": bson.M{"$in": filter.Gender}})
		}
		if len(filter.PlaceOfBirth) > 0 {
			query = append(query, bson.M{"placeOfBirth": bson.M{"$in": filter.PlaceOfBirth}})
		}
		//regex

		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.FatherName != "" {
			query = append(query, bson.M{"fatherName": primitive.Regex{Pattern: filter.Regex.FatherName, Options: "xi"}})
		}
	}
	if filter.DOB != nil {
		//var sd,ed time.Time
		if filter.DOB.From != nil {
			sd := time.Date(filter.DOB.From.Year(), filter.DOB.From.Month(), filter.DOB.From.Day(), 0, 0, 0, 0, filter.DOB.From.Location())
			ed := time.Date(filter.DOB.From.Year(), filter.DOB.To.Month(), filter.DOB.To.Day(), 23, 59, 59, 0, filter.DOB.To.Location())
			if filter.DOB.To != nil {
				ed = time.Date(filter.DOB.To.Year(), filter.DOB.To.Month(), filter.DOB.To.Day(), 23, 59, 59, 0, filter.DOB.To.Location())
			}
			query = append(query, bson.M{"dob": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).CountDocuments(ctx.CTX, func() bson.M {
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
	// //Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONHOSPITAL, "hospitalId", "uniqueId", "ref.hospital", "ref.hospital")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBIRTHCERTIFICATESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var birthcertificate []models.RefBirthCertificate
	if err = cursor.All(context.TODO(), &birthcertificate); err != nil {
		return nil, err
	}
	return birthcertificate, nil
}

func (d *Daos) ApproveBirthCertificate(ctx *models.Context, birthcertificate *models.BirthCertificate) error {
	selector := bson.M{"uniqueId": birthcertificate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"action": birthcertificate.Action, "approved.on": &t, "status": constants.BIRTHCERTIFICATESTATUSAPPROVED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) RejectBirthCertificate(ctx *models.Context, birthcertificate *models.BirthCertificate) error {
	selector := bson.M{"uniqueId": birthcertificate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": bson.M{"action": birthcertificate.Action, "action.on": &t, "status": constants.BIRTHCERTIFICATESTATUSREJECTED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBIRTHCERTIFICATE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
