package handlers

import (
	"bpms-service/app"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//SaveRoofType : ""
func (h *Handler) SaveRoofType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	roofType := new(models.RoofType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&roofType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveRoofType(ctx, roofType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["roofType"] = roofType
	response.With200V2(w, "Success", m, platform)
}

//UpdateRoofType :""
func (h *Handler) UpdateRoofType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	roofType := new(models.RoofType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&roofType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if roofType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateRoofType(ctx, roofType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["roofType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableRoofType : ""
func (h *Handler) EnableRoofType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableRoofType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["roofType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableRoofType : ""
func (h *Handler) DisableRoofType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableRoofType(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["roofType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteRoofType : ""
func (h *Handler) DeleteRoofType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteRoofType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["roofType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleRoofType :""
func (h *Handler) GetSingleRoofType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	roofType := new(models.RefRoofType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	roofType, err := h.Service.GetSingleRoofType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["roofType"] = roofType
	response.With200V2(w, "Success", m, platform)
}

//FilterRoofType : ""
func (h *Handler) FilterRoofType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var roofType *models.RoofTypeFilter
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
	err := json.NewDecoder(r.Body).Decode(&roofType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var roofTypes []models.RefRoofType
	log.Println(pagination)
	roofTypes, err = h.Service.FilterRoofType(ctx, roofType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(roofTypes) > 0 {
		m["roofType"] = roofTypes
	} else {
		res := make([]models.RoofType, 0)
		m["roofType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
