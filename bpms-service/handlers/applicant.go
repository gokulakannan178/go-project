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

//SaveApplicant : ""
func (h *Handler) SaveApplicant(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	applicant := new(models.Applicant)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&applicant)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveApplicant(ctx, applicant)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = applicant
	response.With200V2(w, "Success", m, platform)
}

//UpdateApplicant :""
func (h *Handler) UpdateApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	applicant := new(models.Applicant)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&applicant)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if applicant.UserName == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateApplicant(ctx, applicant)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableApplicant : ""
func (h *Handler) EnableApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableApplicant(ctx, UserName)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableApplicant : ""
func (h *Handler) DisableApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableApplicant(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteApplicant : ""
func (h *Handler) DeleteApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteApplicant(ctx, UserName)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleApplicant :""
func (h *Handler) GetSingleApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "id is missing", platform)
	}

	applicant := new(models.RefApplicant)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	applicant, err := h.Service.GetSingleApplicant(ctx, UserName)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = applicant
	response.With200V2(w, "Success", m, platform)
}

//FilterApplicant : ""
func (h *Handler) FilterApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var applicant *models.ApplicantFilter
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
	err := json.NewDecoder(r.Body).Decode(&applicant)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var applicants []models.RefApplicant
	log.Println(pagination)
	applicants, err = h.Service.FilterApplicant(ctx, applicant, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(applicants) > 0 {
		m["applicant"] = applicants
	} else {
		res := make([]models.Applicant, 0)
		m["applicant"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//BlacklistApplicant : ""
func (h *Handler) BlacklistApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var asc *models.ApplicantStatusChange
	err := json.NewDecoder(r.Body).Decode(&asc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err = h.Service.BlacklistApplicant(ctx, asc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//LicenseCancelApplicant : ""
func (h *Handler) LicenseCancelApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var asc *models.ApplicantStatusChange
	err := json.NewDecoder(r.Body).Decode(&asc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err = h.Service.LicenseCancelApplicant(ctx, asc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ReActivateApplicant : ""
func (h *Handler) ReActivateApplicant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var asc *models.ApplicantStatusChange
	err := json.NewDecoder(r.Body).Decode(&asc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err = h.Service.ReActivateApplicant(ctx, asc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["applicant"] = "success"
	response.With200V2(w, "Success", m, platform)
}
