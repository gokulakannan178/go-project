// Copyright 2019 The Logikoof Technologies Private Limited Authors. All rights reserved.
// No Copy or Redistribution of any part of source code or file
//This file initiated by Solomon Arumugam (solomon@logikoof.com)
package daos

import (
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ChkCommonUniqueness : this method is to check wether a value is already available
/*
Input -
	ctx - application context
	collection - name of collection
	key - name of key to be searched
	value - searched string
Output -
	found - if available - true else false
	err - returns error
*/
//Log
//Added by Solomon Arumuhan (solomon@logikoof.com) on 07-Mar-2022
//Update and Add a New Apis  by Gokulkannan (Gokulkannan.M@logikoof.com) on 10-Mar-2022
func (d *Daos) ChkCommonUniqueness(ctx *models.Context, collection, key, value string) (found bool, err error) {
	var data []interface{}
	// err = ctx.DB.Collection(collection).FindOne(ctx.CTX, bson.M{
	// 	key: bson.M{"$regex": primitive.Regex{
	// 		Pattern: value,
	// 		Options: "i",
	// 	}},
	// }).Decode(&data)
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{key: primitive.Regex{Pattern: value, Options: "i"}})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	cursor, err := ctx.DB.Collection(collection).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return false, err
	}
	if err = cursor.All(ctx.CTX, &data); err != nil {
		return false, err
	}

	if data != nil {
		return true, nil
	}
	return false, nil
}
func (d *Daos) ChkCommonUniquenessWithoutRegex(ctx *models.Context, collection, key, value string) (found bool, err error) {
	var data []interface{}
	// err = ctx.DB.Collection(collection).FindOne(ctx.CTX, bson.M{
	// 	key: bson.M{"$regex": primitive.Regex{
	// 		Pattern: value,
	// 		Options: "i",
	// 	}},
	// }).Decode(&data)
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{key: value})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	cursor, err := ctx.DB.Collection(collection).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return false, err
	}
	if err = cursor.All(ctx.CTX, &data); err != nil {
		return false, err
	}

	if data != nil {
		return true, nil
	}
	return false, nil
}
