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

//SaveTradeLicenseRebate : ""
func (h *Handler) SaveTradeLicenseRebate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tlTradeLicenseRebate := new(models.TradeLicenseRebate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&tlTradeLicenseRebate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveTradeLicenseRebate(ctx, tlTradeLicenseRebate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlTradeLicenseRebate"] = tlTradeLicenseRebate
	response.With200V2(w, "Success", m, platform)
}

//UpdateTradeLicenseRebate :""
func (h *Handler) UpdateTradeLicenseRebate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	tlTradeLicenseRebate := new(models.TradeLicenseRebate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&tlTradeLicenseRebate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if tlTradeLicenseRebate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateTradeLicenseRebate(ctx, tlTradeLicenseRebate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlTradeLicenseRebate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableTradeLicenseRebate : ""
func (h *Handler) EnableTradeLicenseRebate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableTradeLicenseRebate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlTradeLicenseRebate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableTradeLicenseRebate : ""
func (h *Handler) DisableTradeLicenseRebate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableTradeLicenseRebate(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlTradeLicenseRebate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteTradeLicenseRebate : ""
func (h *Handler) DeleteTradeLicenseRebate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteTradeLicenseRebate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlTradeLicenseRebate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleTradeLicenseRebate :""
func (h *Handler) GetSingleTradeLicenseRebate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	tlTradeLicenseRebate := new(models.RefTradeLicenseRebate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	tlTradeLicenseRebate, err := h.Service.GetSingleTradeLicenseRebate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlTradeLicenseRebate"] = tlTradeLicenseRebate
	response.With200V2(w, "Success", m, platform)
}

//FilterTradeLicenseRebate : ""
func (h *Handler) FilterTradeLicenseRebate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var tlTradeLicenseRebate *models.TradeLicenseRebateFilter
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
	err := json.NewDecoder(r.Body).Decode(&tlTradeLicenseRebate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var tlTradeLicenseRebates []models.RefTradeLicenseRebate
	log.Println(pagination)
	tlTradeLicenseRebates, err = h.Service.FilterTradeLicenseRebate(ctx, tlTradeLicenseRebate, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tlTradeLicenseRebates) > 0 {
		m["tlTradeLicenseRebate"] = tlTradeLicenseRebates
	} else {
		res := make([]models.TradeLicenseRebate, 0)
		m["tlTradeLicenseRebate"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
