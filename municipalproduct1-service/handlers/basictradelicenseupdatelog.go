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

// BasicUpdateTradeLicense : ""
func (h *Handler) BasicUpdateTradeLicense(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	btlu := new(models.BasicTradeLicenseUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&btlu)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.BasicUpdateTradeLicense(ctx, btlu)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basictradelicenseupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// AcceptBasicTradeLicenseUpdate : ""
func (h *Handler) AcceptBasicTradeLicenseUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.AcceptBasicTradeLicenseUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AcceptBasicTradeLicenseUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basictradelicenseupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectBasicTradeLicenseUpdate : ""
func (h *Handler) RejectBasicTradeLicenseUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectBasicTradeLicenseUpdate)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectBasicTradeLicenseUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basictradelicenseupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterBasicTradeLicenseUpdateLog : ""
func (h *Handler) FilterBasicTradeLicenseUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.FilterBasicTradeLicenseUpdateLog
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

	var refs []models.RefBasicTradeLicenseUpdateLog
	log.Println(pagination)
	refs, err = h.Service.FilterBasicTradeLicenseUpdateLog(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(refs) > 0 {
		m["data"] = refs
	} else {
		res := make([]models.BasicTradeLicenseUpdateLog, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicenseUpdateLogv2 : ""
func (h *Handler) GetSingleBasicTradeLicenseUpdateLogV2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	trade, err := h.Service.GetSingleBasicTradeLicenseUpdateLogV2(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicense"] = trade
	response.With200V2(w, "Success", m, platform)
}

// BasicTradeLicenseUpdateGetPaymentsToBeUpdated : ""
func (h *Handler) BasicTradeLicenseUpdateGetPaymentsToBeUpdated(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	rbtlul := new(models.RefBasicTradeLicenseUpdateLogV2)
	err := json.NewDecoder(r.Body).Decode(&rbtlul)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	payments, err := h.Service.BasicTradeLicenseUpdateGetPaymentsToBeUpdated(ctx, rbtlul)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payments"] = payments
	if len(payments) < 1 {
		m["payments"] = []interface{}{}

	}
	response.With200V2(w, "Success", m, platform)
}
