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

//SaveApplicantType : ""
func (h *Handler) SaveApplicantType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	applicantType := new(models.ApplicantType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&applicantType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveApplicantType(ctx, applicantType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicantType"] = applicantType
	response.With200V2(w, "Success", m, platform)
}

//UpdateApplicantType :""
func (h *Handler) UpdateApplicantType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	applicantType := new(models.ApplicantType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&applicantType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if applicantType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateApplicantType(ctx, applicantType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicantType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableApplicantType : ""
func (h *Handler) EnableApplicantType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableApplicantType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicantType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableApplicantType : ""
func (h *Handler) DisableApplicantType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableApplicantType(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicantType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteApplicantType : ""
func (h *Handler) DeleteApplicantType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteApplicantType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicantType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleApplicantType :""
func (h *Handler) GetSingleApplicantType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	applicantType := new(models.RefApplicantType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	applicantType, err := h.Service.GetSingleApplicantType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicantType"] = applicantType
	response.With200V2(w, "Success", m, platform)
}

//FilterApplicantType : ""
func (h *Handler) FilterApplicantType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var applicantType *models.ApplicantTypeFilter
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
	err := json.NewDecoder(r.Body).Decode(&applicantType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var applicantTypes []models.RefApplicantType
	log.Println(pagination)
	applicantTypes, err = h.Service.FilterApplicantType(ctx, applicantType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(applicantTypes) > 0 {
		m["applicantType"] = applicantTypes
	} else {
		res := make([]models.ApplicantType, 0)
		m["applicantType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
