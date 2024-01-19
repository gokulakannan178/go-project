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

// SaveShopRentDashboardDayWise : ""
func (h *Handler) SaveShopRentDashboardDayWise(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	shopRent := new(models.ShopRentDashboardDayWise)
	err := json.NewDecoder(r.Body).Decode(&shopRent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveShopRentDashboardDayWise(ctx, shopRent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleShopRentDashboardDayWise : ""
func (h *Handler) GetSingleShopRentDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	shopRent := new(models.RefShopRentDashboardDayWise)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	shopRent, err := h.Service.GetSingleShopRentDashboardDayWise(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboardDayWise"] = shopRent
	response.With200V2(w, "Success", m, platform)
}

// UpdateShopRentDashboardDayWise : ""
func (h *Handler) UpdateShopRentDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	shopRent := new(models.ShopRentDashboardDayWise)
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
	err = h.Service.UpdateShopRentDashboardDayWise(ctx, shopRent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableShopRentDashboardDayWise : ""
func (h *Handler) EnableShopRentDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableShopRentDashboardDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableShopRentDashboardDayWise : ""
func (h *Handler) DisableShopRentDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableShopRentDashboardDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteDashBoardShopRent : ""
func (h *Handler) DeleteShopRentDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteShopRentDashboardDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopRentDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterShopRentDashboardDayWise : ""
func (h *Handler) FilterShopRentDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ShopRentDashboardDayWiseFilter
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

	var shopRentDashboardDayWises []models.RefShopRentDashboardDayWise
	log.Println(pagination)
	shopRentDashboardDayWises, err = h.Service.FilterShopRentDashboardDayWise(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRentDashboardDayWises) > 0 {
		m["shopRentDashboardDayWise"] = shopRentDashboardDayWises
	} else {
		res := make([]models.ShopRentDashboardDayWise, 0)
		m["shopRentDashboardDayWise"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
