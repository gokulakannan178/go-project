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

// SavePayrollPolicyEarning : ""
func (h *Handler) SavePayrollPolicyEarning(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payrollPolicyEarning := new(models.PayrollPolicyEarning)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&payrollPolicyEarning)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SavePayrollPolicyEarning(ctx, payrollPolicyEarning)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyEarning"] = payrollPolicyEarning
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePayrollPolicyEarning : ""
func (h *Handler) GetSinglePayrollPolicyEarning(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	payrollPolicyEarning := new(models.RefPayrollPolicyEarning)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	payrollPolicyEarning, err := h.Service.GetSinglePayrollPolicyEarning(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyEarning"] = payrollPolicyEarning
	response.With200V2(w, "Success", m, platform)
}

//UpdatePayrollPolicyEarning : ""
func (h *Handler) UpdatePayrollPolicyEarning(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	payrollPolicyEarning := new(models.PayrollPolicyEarning)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&payrollPolicyEarning)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if payrollPolicyEarning.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePayrollPolicyEarning(ctx, payrollPolicyEarning)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyEarning"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnablePayrollPolicyEarning : ""
func (h *Handler) EnablePayrollPolicyEarning(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnablePayrollPolicyEarning(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyEarning"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePayrollPolicyEarning : ""
func (h *Handler) DisablePayrollPolicyEarning(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisablePayrollPolicyEarning(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyEarning"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePayrollPolicyEarning : ""
func (h *Handler) DeletePayrollPolicyEarning(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePayrollPolicyEarning(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicyEarning"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPayrollPolicyEarning : ""
func (h *Handler) FilterPayrollPolicyEarning(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterPayrollPolicyEarning *models.FilterPayrollPolicyEarning
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
	err := json.NewDecoder(r.Body).Decode(&filterPayrollPolicyEarning)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterPayrollPolicyEarnings []models.RefPayrollPolicyEarning
	log.Println(pagination)
	filterPayrollPolicyEarnings, err = h.Service.FilterPayrollPolicyEarning(ctx, filterPayrollPolicyEarning, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterPayrollPolicyEarnings) > 0 {
		m["PayrollPolicyEarning"] = filterPayrollPolicyEarnings
	} else {
		res := make([]models.PayrollPolicyEarning, 0)
		m["PayrollPolicyEarning"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
