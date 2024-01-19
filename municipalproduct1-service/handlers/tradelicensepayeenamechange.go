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

// SaveTradelicensePayeeNameChange : ""
func (h *Handler) SaveTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ppnc := new(models.TradelicensePayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ppnc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveTradelicensePayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// UpdateTradelicensePayeeNameChange :""
func (h *Handler) UpdateTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ppnc := new(models.TradelicensePayeeNameChange)
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
	err = h.Service.UpdateTradelicensePayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// EnableTradelicensePayeeNameChange : ""
func (h *Handler) EnableTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableTradelicensePayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableTradelicensePayeeNameChange : ""
func (h *Handler) DisableTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableTradelicensePayeeNameChange(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteTradelicensePayeeNameChange : ""
func (h *Handler) DeleteTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteTradelicensePayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradelicensePayeeNameChange :""
func (h *Handler) GetSingleTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ppnc := new(models.RefTradelicensePayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	ppnc, err := h.Service.GetSingleTradelicensePayeeNameChange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// FilterTradelicensePayeeNameChange : ""
func (h *Handler) FilterTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.TradelicensePayeeNameChangeFilter
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

	var ppncs []models.RefTradelicensePayeeNameChange
	log.Println(pagination)
	ppncs, err = h.Service.FilterTradelicensePayeeNameChange(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ppncs) > 0 {
		m["tpnc"] = ppncs
	} else {
		res := make([]models.TradelicensePayeeNameChange, 0)
		m["tpnc"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// ApproveTradeLicense : ""
func (h *Handler) ApproveTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	approve := new(models.ApproveTradelicensePayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&approve)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ApproveTradelicensePayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// NotApproveTradeLicense : ""
func (h *Handler) NotApproveTradelicensePayeeNameChange(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.NotApproveTradelicensePayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tpnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}
