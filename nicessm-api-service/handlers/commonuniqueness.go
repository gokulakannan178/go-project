// Copyright 2019 The Logikoof Technologies Private Limited Authors. All rights reserved.
// No Copy or Redistribution of any part of source code or file
//This file initiated by Solomon Arumugam (solomon@logikoof.com)
package handlers

import (
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

//ChkCommonUniqueness : this method is to check wether a value is already available
//Added by Solomon Arumuhan (solomon@logikoof.com) on 07-Mar-2022
/*
From Request -
	from - name of collection
	key - name of key to be searched
	value - searched string
Response -
	1) Input error
	2) Logic error
	3) Failed response
	4) Success response
*/
//Log
//Added by Solomon Arumuhan (solomon@logikoof.com) on 07-Mar-2022
//Update and Add a New Apis  by Gokulkannan (Gokulkannan.M@logikoof.com) on 10-Mar-2022
func (h *Handler) ChkCommonUniqueness(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	from := r.URL.Query().Get("from")
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if from == "" || key == "" || value == "" {
		response.With400V2(w, "from or key or value cannot be empty", platform)
		return
	}
	found, err := h.Service.ChkCommonUniqueness(ctx, from, key, value)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	if found {
		m["duplicate"] = "failed"

	} else {
		m["duplicate"] = "success"
	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) ChkCommonUniquenessWithoutRegex(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	from := r.URL.Query().Get("from")
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if from == "" || key == "" || value == "" {
		response.With400V2(w, "from or key or value cannot be empty", platform)
		return
	}
	found, err := h.Service.ChkCommonUniquenessWithoutRegex(ctx, from, key, value)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	if found {
		m["duplicate"] = "failed"

	} else {
		m["duplicate"] = "success"
	}
	response.With200V2(w, "Success", m, platform)
}
