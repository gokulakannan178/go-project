package handlers

import (
	"encoding/json"
	"log"
	"logikoof-echalan-service/app"
	"logikoof-echalan-service/models"
	"logikoof-echalan-service/response"
	"net/http"
	"strconv"
)

//SaveOffenceType : ""
func (h *Handler) SaveOffenceType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	offenceType := new(models.OffenceType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&offenceType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOffenceType(ctx, offenceType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceType"] = offenceType
	response.With200V2(w, "Success", m, platform)
}

//UpdateOffenceType :""
func (h *Handler) UpdateOffenceType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	offenceType := new(models.OffenceType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&offenceType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if offenceType.UniqueID == "" {
		response.With400V2(w, "RegNo is missing", platform)
	}
	err = h.Service.UpdateOffenceType(ctx, offenceType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOffenceType : ""
func (h *Handler) EnableOffenceType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableOffenceType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableOffenceType : ""
func (h *Handler) DisableOffenceType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableOffenceType(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOffenceType : ""
func (h *Handler) DeleteOffenceType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOffenceType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOffenceType :""
func (h *Handler) GetSingleOffenceType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	offenceType := new(models.RefOffenceType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	offenceType, err := h.Service.GetSingleOffenceType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceType"] = offenceType
	response.With200V2(w, "Success", m, platform)
}

//FilterOffenceType : ""
func (h *Handler) FilterOffenceType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var offenceType *models.OffenceTypeFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

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
	err := json.NewDecoder(r.Body).Decode(&offenceType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var offenceTypes []models.RefOffenceType
	log.Println(pagination)
	offenceTypes, err = h.Service.FilterOffenceType(ctx, offenceType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(offenceTypes) > 0 {
		m["offenceType"] = offenceTypes
	} else {
		res := make([]models.OffenceType, 0)
		m["offenceType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
