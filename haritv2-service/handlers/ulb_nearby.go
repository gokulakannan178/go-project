package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//ULBNearBy :""
func (h *Handler) ULBNearBy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	ulbnb := new(models.ULBNearBy)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&ulbnb)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	//km!=0
	if ulbnb.KM == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	ulbs, err := h.Service.ULBNearBy(ctx, ulbnb, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbNearBy"] = ulbs
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//UlbInTheState :""
func (h *Handler) UlbInTheState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	vars := mux.Vars(r)
	stateID := vars["state"]
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	ulbs, err := h.Service.UlbInTheState(ctx, stateID, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbNearBy"] = ulbs
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//UlbInTheStateV2 :""
func (h *Handler) UlbInTheStateV2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	sortOrder := r.URL.Query().Get("sortOrder")
	sortBy := r.URL.Query().Get("sortBy")
	sortorder := 0

	if sortorder = 1; sortOrder != "" {
		var err error
		sortorder, err = strconv.Atoi(sortOrder)
		if err != nil {
			log.Println(err)
		}

	}
	vars := mux.Vars(r)
	stateID := vars["state"]
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	ulbs, err := h.Service.UlbInTheStateV2(ctx, stateID, sortBy, sortorder, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbNearBy"] = ulbs

	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//UlbInTheStateV3 :""
func (h *Handler) UlbInTheStateV3(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	ULBStateIn := new(models.ULBStateIn)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	ulbs, err := h.Service.UlbInTheStateV3(ctx, ULBStateIn, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbNearBy"] = ulbs

	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//UlbCompostInTheState : ""
func (h *Handler) UlbCompostInTheState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	vars := mux.Vars(r)
	stateID := vars["state"]
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	ulbs, err := h.Service.UlbCompostInTheState(ctx, stateID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Compost"] = ulbs

	response.With200V2(w, "Success", m, platform)
}
