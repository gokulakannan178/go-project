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

// SaveSolidWasteUserChargeCategory : ""
func (h *Handler) SaveSolidWasteUserChargeCategory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	solidwasteuserchargecategory := new(models.SolidWasteUserChargeCategory)
	err := json.NewDecoder(r.Body).Decode(&solidwasteuserchargecategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveSolidWasteUserChargeCategory(ctx, solidwasteuserchargecategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSolidWasteUserChargeCategory : ""
func (h *Handler) GetSingleSolidWasteUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	solidwasteuserchargecategory := new(models.RefSolidWasteUserChargeCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	solidwasteuserchargecategory, err := h.Service.GetSingleSolidWasteUserChargeCategory(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargecategory"] = solidwasteuserchargecategory
	response.With200V2(w, "Success", m, platform)
}

// UpdateSolidWasteUserChargeCategory : ""
func (h *Handler) UpdateSolidWasteUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	solidwasteuserchargecategory := new(models.SolidWasteUserChargeCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&solidwasteuserchargecategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if solidwasteuserchargecategory.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateSolidWasteUserChargeCategory(ctx, solidwasteuserchargecategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableSolidWasteUserChargeCategory : ""
func (h *Handler) EnableSolidWasteUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableSolidWasteUserChargeCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableSolidWasteUserChargeCategory : ""
func (h *Handler) DisableSolidWasteUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableSolidWasteUserChargeCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteSolidWasteUserChargeCategory : ""
func (h *Handler) DeleteSolidWasteUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteSolidWasteUserChargeCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargecategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSolidWasteUserChargeCategory : ""
func (h *Handler) FilterSolidWasteUserChargeCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.SolidWasteUserChargeCategoryFilter
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

	var solidwasteuserchargecategorys []models.RefSolidWasteUserChargeCategory
	log.Println(pagination)
	solidwasteuserchargecategorys, err = h.Service.FilterSolidWasteUserChargeCategory(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(solidwasteuserchargecategorys) > 0 {
		m["solidwasteuserchargecategory"] = solidwasteuserchargecategorys
	} else {
		res := make([]models.SolidWasteUserChargeCategory, 0)
		m["solidwasteuserchargecategory"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
