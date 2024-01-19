package handlers

import (
	"encoding/json"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

//SaveResidentialType : ""
func (h *Handler) SaveResidentialType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	residentialType := new(models.ResidentialType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&residentialType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveResidentialType(ctx, residentialType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["residentialType"] = residentialType
	response.With200V2(w, "Success", m, platform)
}

//UpdateResidentialType :""
func (h *Handler) UpdateResidentialType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	residentialType := new(models.ResidentialType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&residentialType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if residentialType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateResidentialType(ctx, residentialType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["residentialType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableResidentialType : ""
func (h *Handler) EnableResidentialType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableResidentialType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["residentialType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableResidentialType : ""
func (h *Handler) DisableResidentialType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableResidentialType(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["residentialType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteResidentialType : ""
func (h *Handler) DeleteResidentialType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteResidentialType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["residentialType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleResidentialType :""
func (h *Handler) GetSingleResidentialType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	residentialType := new(models.RefResidentialType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	residentialType, err := h.Service.GetSingleResidentialType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["residentialType"] = residentialType
	response.With200V2(w, "Success", m, platform)
}

//FilterResidentialType : ""
func (h *Handler) FilterResidentialType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var residentialType *models.ResidentialTypeFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
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
	err := json.NewDecoder(r.Body).Decode(&residentialType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var residentialTypes []models.RefResidentialType
	log.Println(pagination)
	residentialTypes, err = h.Service.FilterResidentialType(ctx, residentialType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(residentialTypes) > 0 {
		m["residentialType"] = residentialTypes
	} else {
		res := make([]models.ResidentialType, 0)
		m["residentialType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
