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

// SaveSolidWasteUserChargeSubCategory : ""
func (h *Handler) SaveSolidWasteUserChargeSubCategory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	solidwasteuserchargesubcategory := new(models.SolidWasteUserChargeSubCategory)
	err := json.NewDecoder(r.Body).Decode(&solidwasteuserchargesubcategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveSolidWasteUserChargeSubCategory(ctx, solidwasteuserchargesubcategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargesubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSolidWasteUserChargeSubCategory : ""
func (h *Handler) GetSingleSolidWasteUserChargeSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	solidwasteuserchargesubcategory := new(models.RefSolidWasteUserChargeSubCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	solidwasteuserchargesubcategory, err := h.Service.GetSingleSolidWasteUserChargeSubCategory(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargesubcategory"] = solidwasteuserchargesubcategory
	response.With200V2(w, "Success", m, platform)
}

// UpdateSolidWasteUserChargeSubCategory : ""
func (h *Handler) UpdateSolidWasteUserChargeSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	solidwasteuserchargesubcategory := new(models.SolidWasteUserChargeSubCategory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&solidwasteuserchargesubcategory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if solidwasteuserchargesubcategory.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateSolidWasteUserChargeSubCategory(ctx, solidwasteuserchargesubcategory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargesubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableSolidWasteUserChargeSubCategory : ""
func (h *Handler) EnableSolidWasteUserChargeSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableSolidWasteUserChargeSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargesubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableSolidWasteUserChargeSubCategory : ""
func (h *Handler) DisableSolidWasteUserChargeSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableSolidWasteUserChargeSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargesubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteSolidWasteUserChargeSubCategory : ""
func (h *Handler) DeleteSolidWasteUserChargeSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteSolidWasteUserChargeSubCategory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargesubcategory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSolidWasteUserChargeSubCategory : ""
func (h *Handler) FilterSolidWasteUserChargeSubCategory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.SolidWasteUserChargeSubCategoryFilter
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

	var solidwasteuserchargesubcategorys []models.RefSolidWasteUserChargeSubCategory
	log.Println(pagination)
	solidwasteuserchargesubcategorys, err = h.Service.FilterSolidWasteUserChargeSubCategory(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(solidwasteuserchargesubcategorys) > 0 {
		m["solidwasteuserchargesubcategory"] = solidwasteuserchargesubcategorys
	} else {
		res := make([]models.SolidWasteUserChargeSubCategory, 0)
		m["solidwasteuserchargesubcategory"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
