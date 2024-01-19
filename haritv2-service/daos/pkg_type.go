package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SavePkgType :""
func (d *Daos) SavePkgType(ctx *models.Context, pkg *models.PkgType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPKG).InsertOne(ctx.CTX, pkg)
	return err
}

//GetSinglePkgType : ""
func (d *Daos) GetSinglePkgType(ctx *models.Context, uniqueID string) (*models.RefPkgType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPKGCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPKG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pkgs []models.RefPkgType
	var pkg *models.RefPkgType
	if err = cursor.All(ctx.CTX, &pkgs); err != nil {
		return nil, err
	}
	if len(pkgs) > 0 {
		pkg = &pkgs[0]
	}
	return pkg, nil
}

//UpdatePkgType : ""
func (d *Daos) UpdatePkgType(ctx *models.Context, pkg *models.PkgType) error {
	selector := bson.M{"uniqueId": pkg.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": pkg}
	_, err := ctx.DB.Collection(constants.COLLECTIONPKG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnablePkgType :""
func (d *Daos) EnablePkgType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPKG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePkgType :""
func (d *Daos) DisablePkgType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPKG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePkgType :""
func (d *Daos) DeletePkgType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPKG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterPkgType : ""
func (d *Daos) FilterPkgType(ctx *models.Context, pkgfilter *models.PkgTypeFilter, pagination *models.Pagination) ([]models.RefPkgType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if pkgfilter != nil {
		if len(pkgfilter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": pkgfilter.Name}})
		}
		if len(pkgfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": pkgfilter.Status}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if pkgfilter != nil {
		if pkgfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{pkgfilter.SortBy: pkgfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPKG).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPKG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pkgs []models.RefPkgType
	if err = cursor.All(context.TODO(), &pkgs); err != nil {
		return nil, err
	}
	return pkgs, nil
}

//GetDefaultPkgType : ""
func (d *Daos) GetDefaultPkgType(ctx *models.Context) (*models.RefPkgType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isDefault": true}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPKG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pkgs []models.RefPkgType
	var pkg *models.RefPkgType
	if err = cursor.All(ctx.CTX, &pkgs); err != nil {
		return nil, err
	}
	if len(pkgs) > 0 {
		pkg = &pkgs[0]
	}
	return pkg, nil
}
