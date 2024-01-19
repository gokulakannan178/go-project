package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

//SaveHonoriffic : ""
func (h *Handler) SaveHonoriffic(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	honoriffic := new(models.Honoriffic)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&honoriffic)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveHonoriffic(ctx, honoriffic)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["honoriffic"] = honoriffic
	response.With200V2(w, "Success", m, platform)
}

//UpdateHonoriffic :""
func (h *Handler) UpdateHonoriffic(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	honoriffic := new(models.Honoriffic)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&honoriffic)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if honoriffic.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateHonoriffic(ctx, honoriffic)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["honoriffic"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableHonoriffic : ""
func (h *Handler) EnableHonoriffic(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableHonoriffic(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["honoriffic"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableHonoriffic : ""
func (h *Handler) DisableHonoriffic(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableHonoriffic(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["honoriffic"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteHonoriffic : ""
func (h *Handler) DeleteHonoriffic(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteHonoriffic(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["honoriffic"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleHonoriffic :""
func (h *Handler) GetSingleHonoriffic(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	honoriffic := new(models.RefHonoriffic)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	honoriffic, err := h.Service.GetSingleHonoriffic(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["honoriffic"] = honoriffic
	response.With200V2(w, "Success", m, platform)
}

//FilterHonoriffic : ""
func (h *Handler) FilterHonoriffic(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var honoriffic *models.HonorifficFilter
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
	err := json.NewDecoder(r.Body).Decode(&honoriffic)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var honoriffics []models.RefHonoriffic
	log.Println(pagination)
	honoriffics, err = h.Service.FilterHonoriffic(ctx, honoriffic, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(honoriffics) > 0 {
		m["honoriffic"] = honoriffics
	} else {
		res := make([]models.Honoriffic, 0)
		m["honoriffic"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
