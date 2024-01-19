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

// SaveShopRentDashboard : ""
func (h *Handler) SaveShopRentDashboard(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	shopRent := new(models.ShopRentDashboard)
	err := json.NewDecoder(r.Body).Decode(&shopRent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveShopRentDashboard(ctx, shopRent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleShopRentDashboard : ""
func (h *Handler) GetSingleShopRentDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	shopRent := new(models.RefShopRentDashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	shopRent, err := h.Service.GetSingleShopRentDashboard(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboard"] = shopRent
	response.With200V2(w, "Success", m, platform)
}

// UpdateShopRentDashboard : ""
func (h *Handler) UpdateShopRentDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	shopRent := new(models.ShopRentDashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&shopRent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if shopRent.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateShopRentDashboard(ctx, shopRent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableShopRentDashboard : ""
func (h *Handler) EnableShopRentDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableShopRentDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableShopRentDashboard : ""
func (h *Handler) DisableShopRentDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableShopRentDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteShopRentDashboard : ""
func (h *Handler) DeleteShopRentDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteShopRentDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterShopRentDashboard : ""
func (h *Handler) FilterShopRentDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ShopRentDashboardFilter
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

	var shopRentDashboards []models.RefShopRentDashboard
	log.Println(pagination)
	shopRentDashboards, err = h.Service.FilterShopRentDashboard(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRentDashboards) > 0 {
		m["shopRentDashboard"] = shopRentDashboards
	} else {
		res := make([]models.ShopRentDashboard, 0)
		m["shopRentDashboard"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// DashboardShopRentDemandAndCollection : ""
func (h *Handler) DashboardShopRentDemandAndCollection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var filter *models.DashboardShopRentDemandAndCollectionFilter
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	res, err := h.Service.DashboardShopRentDemandAndCollection(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed no data in this id "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ddac"] = res
	response.With200V2(w, "Success", m, platform)

}
