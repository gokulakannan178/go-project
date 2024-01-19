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

// SaveSalaryConfig : ""
func (h *Handler) SaveSalaryConfig(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	salaryConfig := new(models.SalaryConfig)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&salaryConfig)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveSalaryConfig(ctx, salaryConfig)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfig"] = salaryConfig
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSalaryConfig : ""
func (h *Handler) GetSingleSalaryConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	salaryConfig := new(models.RefSalaryConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	salaryConfig, err := h.Service.GetSingleSalaryConfig(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfig"] = salaryConfig
	response.With200V2(w, "Success", m, platform)
}

//UpdateSalaryConfig : ""
func (h *Handler) UpdateSalaryConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	salaryConfig := new(models.SalaryConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&salaryConfig)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if salaryConfig.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateSalaryConfig(ctx, salaryConfig)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableSalaryConfig : ""
func (h *Handler) EnableSalaryConfig(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableSalaryConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfig"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableSalaryConfig : ""
func (h *Handler) DisableSalaryConfig(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableSalaryConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfig"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteSalaryConfig : ""
func (h *Handler) DeleteSalaryConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteSalaryConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSalaryConfig : ""
func (h *Handler) FilterSalaryConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterSalaryConfig *models.FilterSalaryConfig
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
	err := json.NewDecoder(r.Body).Decode(&filterSalaryConfig)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterSalaryConfigs []models.RefSalaryConfig
	log.Println(pagination)
	filterSalaryConfigs, err = h.Service.FilterSalaryConfig(ctx, filterSalaryConfig, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterSalaryConfigs) > 0 {
		m["SalaryConfig"] = filterSalaryConfigs
	} else {
		res := make([]models.SalaryConfig, 0)
		m["SalaryConfig"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveSalaryConfigWithEmployeeType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	salaryConfig := new(models.SalaryConfig)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&salaryConfig)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveSalaryConfigWithEmployeeType(ctx, salaryConfig)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfig"] = salaryConfig
	response.With200V2(w, "Success", m, platform)
}
