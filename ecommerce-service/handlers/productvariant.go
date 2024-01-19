package handlers

import (
	"ecommerce-service/app"
	"ecommerce-service/models"
	"ecommerce-service/response"
	"encoding/json"
	"log"

	"net/http"
	"strconv"
)

// SaveProductVariant : ""
func (h *Handler) SaveProductVariant(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	block := new(models.ProductVariant)
	err := json.NewDecoder(r.Body).Decode(&block)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveProductVariant(ctx, block)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productVariant"] = block
	response.With200V2(w, "Success", m, platform)
}

// GetSingleProductVariant : ""
func (h *Handler) GetSingleProductVariant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	crop := new(models.RefProductVariant)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	crop, err := h.Service.GetSingleProductVariant(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productVariant"] = crop
	response.With200V2(w, "Success", m, platform)
}

// UpdateProductVariant : ""
func (h *Handler) UpdateProductVariant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	crop := new(models.ProductVariant)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&crop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if crop.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateProductVariant(ctx, crop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productVariant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableProductVariant : ""
func (h *Handler) EnableProductVariant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableProductVariant(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productVariant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableProductVariant : ""
func (h *Handler) DisableProductVariant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableProductVariant(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productVariant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteProductVariant : ""
func (h *Handler) DeleteProductVariant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteProductVariant(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productVariant"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterProductVariant : ""
func (h *Handler) FilterProductVariant(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ProductVariantFilter
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

	var ProductVariants []models.RefProductVariant
	log.Println(pagination)
	ProductVariants, err = h.Service.FilterProductVariant(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ProductVariants) > 0 {
		m["productVariant"] = ProductVariants
	} else {
		res := make([]models.ProductVariant, 0)
		m["productVariant"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) ProductVariantRegister(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	product := new(models.RegProductVariant)
	err := json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.ProductVariantRegister(ctx, product)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productVariant"] = product
	response.With200V2(w, "Success", m, platform)
}

// GetMyInventory : ""
func (h *Handler) GetMyInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ProductVariantInventoryFilter
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

	var ProductVariants []models.RefProductVariant
	log.Println(pagination)
	ProductVariants, err = h.Service.GetMyInventory(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ProductVariants) > 0 {
		m["productVariant"] = ProductVariants
	} else {
		res := make([]models.ProductVariant, 0)
		m["productVariant"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
