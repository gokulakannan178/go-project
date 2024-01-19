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

// SaveTradeLicenseDashboard : ""
func (h *Handler) SaveTradeLicenseDashboard(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tradeLicense := new(models.TradeLicenseDashboard)
	err := json.NewDecoder(r.Body).Decode(&tradeLicense)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveTradeLicenseDashboard(ctx, tradeLicense)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicenseDashboard : ""
func (h *Handler) GetSingleTradeLicenseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	tradeLicense := new(models.RefTradeLicenseDashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	tradeLicense, err := h.Service.GetSingleTradeLicenseDashboard(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboard"] = tradeLicense
	response.With200V2(w, "Success", m, platform)
}

// UpdateTradeLicenseDashboard : ""
func (h *Handler) UpdateTradeLicenseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	tradeLicense := new(models.TradeLicenseDashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&tradeLicense)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if tradeLicense.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateTradeLicenseDashboard(ctx, tradeLicense)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableTradeLicenseDashboard : ""
func (h *Handler) EnableTradeLicenseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableTradeLicenseDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableTradeLicenseDashboard : ""
func (h *Handler) DisableTradeLicenseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableTradeLicenseDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteDashBoardTradeLicense : ""
func (h *Handler) DeleteDashBoardTradeLicense(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDashBoardTradeLicense(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterTradeLicenseDashboard : ""
func (h *Handler) FilterTradeLicenseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.TradeLicenseDashboardFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var tradeLicenseDashboards []models.RefTradeLicenseDashboard
	log.Println(pagination)
	tradeLicenseDashboards, err = h.Service.FilterTradeLicenseDashboard(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenseDashboards) > 0 {
		m["tradeLicenseDashboard"] = tradeLicenseDashboards
	} else {
		res := make([]models.TradeLicenseDashboard, 0)
		m["tradeLicenseDashboard"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// DashboardTradeLicenseDemandAndCollection : ""
func (h *Handler) DashboardTradeLicenseDemandAndCollection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var filter *models.DashboardTradeLicenseDemandAndCollectionFilter
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	res, err := h.Service.DashboardTradeLicenseDemandAndCollection(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed no data in this id "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ddac"] = res
	response.With200V2(w, "Success", m, platform)

}

// GetUserChargeSAFDashboard : ""
func (h *Handler) UserwiseTradelicenseReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.UserFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if resType == "excel" {
		file, err := h.Service.UserwiseTradeLicenseReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=userwisetradelicense.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	result, err := h.Service.UserwiseTradelicenseReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlSafResult"] = result
	response.With200V2(w, "Success", m, platform)
}
