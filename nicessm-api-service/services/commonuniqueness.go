// Copyright 2019 The Logikoof Technologies Private Limited Authors. All rights reserved.
// No Copy or Redistribution of any part of source code or file
//This file initiated by Solomon Arumugam (solomon@logikoof.com)
package services

import (
	"nicessm-api-service/models"
)

//ChkCommonUniqueness : this method is to check wether a value is already available
//Added by Solomon Arumuhan (solomon@logikoof.com) on 07-Mar-2022
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
func (s *Service) ChkCommonUniqueness(ctx *models.Context, collection, key, value string) (found bool, err error) {
	return s.Daos.ChkCommonUniqueness(ctx, collection, key, value)
}
func (s *Service) ChkCommonUniquenessWithoutRegex(ctx *models.Context, collection, key, value string) (found bool, err error) {
	return s.Daos.ChkCommonUniquenessWithoutRegex(ctx, collection, key, value)
}
