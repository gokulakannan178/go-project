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

// SaveShopRentShopSubCategory : ""
func (h *Handler) SaveShopRentShopSubCategory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	shopsubcategory := new(models.ShopRentShopSubCategory)
	err := json.NewDecoder(r.Body).Decode(&shopsubcategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveShopRentShopSubCategory(ctx, shopsubcategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleShopRentShopSubCategory : ""
func (h *Handler) GetSingleShopRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	shopsubcategory := new(models.RefShopRentShopSubCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	shopsubcategory, err := h.Service.GetSingleShopRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = shopsubcategory
	response.With200V2(w, "Success", m, platform)
}

// UpdateShopRentShopSubCategory : ""
func (h *Handler) UpdateShopRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	shopsubcategory := new(models.ShopRentShopSubCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&shopsubcategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if shopsubcategory.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateShopRentShopSubCategory(ctx, shopsubcategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableShopRentShopSubCategory : ""
func (h *Handler) EnableShopRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableShopRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableShopRentShopSubCategory : ""
func (h *Handler) DisableShopRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableShopRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteShopRentShopSubCategory : ""
func (h *Handler) DeleteShopRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteShopRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterShopRentShopSubCategory : ""
func (h *Handler) FilterShopRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ShopRentShopSubCategoryFilter
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

	var shopsubcategorys []models.RefShopRentShopSubCategory
	log.Println(pagination)
	shopsubcategorys, err = h.Service.FilterShopRentShopSubCategory(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopsubcategorys) > 0 {
		m["shopsubcategory"] = shopsubcategorys
	} else {
		res := make([]models.MobileTowerTax, 0)
		m["shopsubcategory"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
