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

// SavePayrollPolicyDetection : ""
func (h *Handler) SavePayrollPolicyDetection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payrollPolicyDetection := new(models.PayrollPolicyDetection)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&payrollPolicyDetection)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SavePayrollPolicyDetection(ctx, payrollPolicyDetection)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyDetection"] = payrollPolicyDetection
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePayrollPolicyDetection : ""
func (h *Handler) GetSinglePayrollPolicyDetection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	payrollPolicyDetection := new(models.RefPayrollPolicyDetection)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	payrollPolicyDetection, err := h.Service.GetSinglePayrollPolicyDetection(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyDetection"] = payrollPolicyDetection
	response.With200V2(w, "Success", m, platform)
}

//UpdatePayrollPolicyDetection : ""
func (h *Handler) UpdatePayrollPolicyDetection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	payrollPolicyDetection := new(models.PayrollPolicyDetection)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&payrollPolicyDetection)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if payrollPolicyDetection.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePayrollPolicyDetection(ctx, payrollPolicyDetection)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyDetection"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnablePayrollPolicyDetection : ""
func (h *Handler) EnablePayrollPolicyDetection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnablePayrollPolicyDetection(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyDetection"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePayrollPolicyDetection : ""
func (h *Handler) DisablePayrollPolicyDetection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisablePayrollPolicyDetection(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyDetection"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePayrollPolicyDetection : ""
func (h *Handler) DeletePayrollPolicyDetection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePayrollPolicyDetection(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyDetection"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPayrollPolicyDetection : ""
func (h *Handler) FilterPayrollPolicyDetection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterPayrollPolicyDetection *models.FilterPayrollPolicyDetection
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
	err := json.NewDecoder(r.Body).Decode(&filterPayrollPolicyDetection)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterPayrollPolicyDetections []models.RefPayrollPolicyDetection
	log.Println(pagination)
	filterPayrollPolicyDetections, err = h.Service.FilterPayrollPolicyDetection(ctx, filterPayrollPolicyDetection, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterPayrollPolicyDetections) > 0 {
		m["PayrollPolicyDetection"] = filterPayrollPolicyDetections
	} else {
		res := make([]models.PayrollPolicyDetection, 0)
		m["PayrollPolicyDetection"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
