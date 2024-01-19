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

// SaveMobileTower : ""
func (h *Handler) SaveMobileTower(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mobileTower := new(models.PropertyMobileTower)
	err := json.NewDecoder(r.Body).Decode(&mobileTower)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveMobileTower(ctx, mobileTower)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTower"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleMobileTower : ""
func (h *Handler) GetSingleMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	mobileTower := new(models.RefPropertyMobileTower)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mobileTower, err := h.Service.GetSingleMobileTower(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDashboard"] = mobileTower
	response.With200V2(w, "Success", m, platform)
}

// UpdateMobileTower : ""
func (h *Handler) UpdateMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mobileTower := new(models.PropertyMobileTower)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mobileTower)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if mobileTower.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateMobileTower(ctx, mobileTower)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableMobileTower : ""
func (h *Handler) EnableMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableMobileTower(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableMobileTower : ""
func (h *Handler) DisableMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableMobileTower(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteMobileTower : ""
func (h *Handler) DeleteMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteMobileTower(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterMobileTower : ""
func (h *Handler) FilterMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyMobileTowerFilter
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

	var mobileTowerDashboards []models.RefPropertyMobileTower
	log.Println(pagination)
	mobileTowerDashboards, err = h.Service.FilterMobileTower(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(mobileTowerDashboards) > 0 {
		m["mobileTowerDashboard"] = mobileTowerDashboards
	} else {
		res := make([]models.MobileTowerDashboard, 0)
		m["mobileTowerDashboard"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// DashboardMobileTowerDemandAndCollection : ""
func (h *Handler) DashboardMobileTowerDemandAndCollection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var filter *models.DashboardMobileTowerDemandAndCollectionFilter
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	res, err := h.Service.DashboardMobileTowerDemandAndCollection(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed no data in this id "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ddac"] = res
	response.With200V2(w, "Success", m, platform)

}
