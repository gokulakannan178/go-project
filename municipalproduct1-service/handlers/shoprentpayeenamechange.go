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

// SaveShoprentPayeeNameChange : ""
func (h *Handler) SaveShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ppnc := new(models.ShoprentPayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ppnc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveShoprentPayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayeenamechange"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// UpdateShoprentPayeeNameChange :""
func (h *Handler) UpdateShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ppnc := new(models.ShoprentPayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&ppnc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if ppnc.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateShoprentPayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayeenamechange"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// EnableShoprentPayeeNameChange : ""
func (h *Handler) EnableShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableShoprentPayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayeenamechange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableShoprentPayeeNameChange : ""
func (h *Handler) DisableShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableShoprentPayeeNameChange(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayeenamechange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteShoprentPayeeNameChange : ""
func (h *Handler) DeleteShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteShoprentPayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["upshoprentpayeenamechangenc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleShoprentPayeeNameChange :""
func (h *Handler) GetSingleShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ppnc := new(models.RefShoprentPayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	ppnc, err := h.Service.GetSingleShoprentPayeeNameChange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayeenamechange"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// FilterShoprentPayeeNameChange : ""
func (h *Handler) FilterShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ShoprentPayeeNameChangeFilter
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

	var ppncs []models.RefShoprentPayeeNameChange
	log.Println(pagination)
	ppncs, err = h.Service.FilterShoprentPayeeNameChange(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ppncs) > 0 {
		m["shoprentpayeenamechange"] = ppncs
	} else {
		res := make([]models.ShoprentPayeeNameChange, 0)
		m["shoprentpayeenamechange"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// ApproveTradeLicense : ""
func (h *Handler) ApproveShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	approve := new(models.ApproveShoprentPayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&approve)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ApproveShoprentPayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayeenamechange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// NotApproveTradeLicense : ""
func (h *Handler) NotApproveShoprentPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	approve := new(models.NotApproveTradeLicense)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&approve)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.NotApproveShoprentPayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayeenamechange"] = "success"
	response.With200V2(w, "Success", m, platform)
}
