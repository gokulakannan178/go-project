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

// SaveLeaseRentShopSubCategory : ""
func (h *Handler) SaveLeaseRentShopSubCategory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	shopsubcategory := new(models.LeaseRentShopSubCategory)
	err := json.NewDecoder(r.Body).Decode(&shopsubcategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveLeaseRentShopSubCategory(ctx, shopsubcategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleLeaseRentShopSubCategory : ""
func (h *Handler) GetSingleLeaseRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	shopsubcategory := new(models.LeaseRentShopSubCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	shopsubcategory, err := h.Service.GetSingleLeaseRentShopSubCategory(ctx, UniqueID)
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

// UpdateLeaseRentShopSubCategory : ""
func (h *Handler) UpdateLeaseRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	shopsubcategory := new(models.LeaseRentShopSubCategory)
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
	err = h.Service.UpdateLeaseRentShopSubCategory(ctx, shopsubcategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableLeaseRentShopSubCategory : ""
func (h *Handler) EnableLeaseRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableLeaseRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableLeaseRentShopSubCategory : ""
func (h *Handler) DisableLeaseRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableLeaseRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteLeaseRentShopSubCategory : ""
func (h *Handler) DeleteLeaseRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteLeaseRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shopsubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterLeaseRentShopSubCategory : ""
func (h *Handler) FilterLeaseRentShopSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.LeaseRentShopSubCategoryFilter
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

	var shopsubcategorys []models.LeaseRentShopSubCategory
	log.Println(pagination)
	shopsubcategorys, err = h.Service.FilterLeaseRentShopSubCategory(ctx, filter, pagination)
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
