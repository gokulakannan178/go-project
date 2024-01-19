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

// SaveUserChargeCategory : ""
func (h *Handler) SaveUserChargeCategory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	userchargecategory := new(models.UserChargeCategory)
	err := json.NewDecoder(r.Body).Decode(&userchargecategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveUserChargeCategory(ctx, userchargecategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userchargecategory"] = userchargecategory
	response.With200V2(w, "Success", m, platform)
}

// GetSingleUserChargeCategory : ""
func (h *Handler) GetSingleUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	userchargecategory := new(models.RefUserChargeCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	userchargecategory, err := h.Service.GetSingleUserChargeCategory(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userchargecategory"] = userchargecategory
	response.With200V2(w, "Success", m, platform)
}

// UpdateUserChargeCategory : ""
func (h *Handler) UpdateUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	userchargecategory := new(models.UserChargeCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&userchargecategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if userchargecategory.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserChargeCategory(ctx, userchargecategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableUserChargeCategory : ""
func (h *Handler) EnableUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUserChargeCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableUserChargeCategory : ""
func (h *Handler) DisableUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUserChargeCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteUserChargeCategory : ""
func (h *Handler) DeleteUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteUserChargeCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterUserChargeCategory : ""
func (h *Handler) FilterUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.UserChargeCategoryFilter
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

	var UserChargeCategorys []models.RefUserChargeCategory
	log.Println(pagination)
	UserChargeCategorys, err = h.Service.FilterUserChargeCategory(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(UserChargeCategorys) > 0 {
		m["userchargecategory"] = UserChargeCategorys
	} else {
		res := make([]models.RefUserChargeCategory, 0)
		m["userchargecategory"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
