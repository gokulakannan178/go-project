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

// SaveSolidWasteUserChargeRate : ""
func (h *Handler) SaveSolidWasteUserChargeRate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	solidwasteuserchargerate := new(models.SolidWasteUserChargeRate)
	err := json.NewDecoder(r.Body).Decode(&solidwasteuserchargerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveSolidWasteUserChargeRate(ctx, solidwasteuserchargerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSolidWasteUserChargeRate : ""
func (h *Handler) GetSingleSolidWasteUserChargeRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	solidwasteuserchargerate := new(models.RefSolidWasteUserChargeRate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	solidwasteuserchargerate, err := h.Service.GetSingleSolidWasteUserChargeRate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargerate"] = solidwasteuserchargerate
	response.With200V2(w, "Success", m, platform)
}

// UpdateSolidWasteUserChargeRate : ""
func (h *Handler) UpdateSolidWasteUserChargeRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	solidwasteuserchargerate := new(models.SolidWasteUserChargeRate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&solidwasteuserchargerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if solidwasteuserchargerate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateSolidWasteUserChargeRate(ctx, solidwasteuserchargerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableSolidWasteUserChargeRate : ""
func (h *Handler) EnableSolidWasteUserChargeRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableSolidWasteUserChargeRate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableSolidWasteUserChargeRate : ""
func (h *Handler) DisableSolidWasteUserChargeRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableSolidWasteUserChargeRate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteSolidWasteUserChargeRate : ""
func (h *Handler) DeleteSolidWasteUserChargeRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteSolidWasteUserChargeRate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteuserchargerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSolidWasteUserChargeRate : ""
func (h *Handler) FilterSolidWasteUserChargeRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.SolidWasteUserChargeRateFilter
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

	var solidwasteuserchargerates []models.RefSolidWasteUserChargeRate
	log.Println(pagination)
	solidwasteuserchargerates, err = h.Service.FilterSolidWasteUserChargeRate(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(solidwasteuserchargerates) > 0 {
		m["SolidWasteUserChargeRate"] = solidwasteuserchargerates
	} else {
		res := make([]models.SolidWasteUserChargeRate, 0)
		m["SolidWasteUserChargeRate"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
