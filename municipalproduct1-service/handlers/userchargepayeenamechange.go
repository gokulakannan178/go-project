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

// SaveUserchargePayeeNameChange : ""
func (h *Handler) SaveUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ppnc := new(models.UserchargePayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ppnc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUserchargePayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["upnc"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// UpdateUserchargePayeeNameChange :""
func (h *Handler) UpdateUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ppnc := new(models.UserchargePayeeNameChange)
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
	err = h.Service.UpdateUserchargePayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["upnc"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// EnableUserchargePayeeNameChange : ""
func (h *Handler) EnableUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUserchargePayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["upnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableUserchargePayeeNameChange : ""
func (h *Handler) DisableUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUserchargePayeeNameChange(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["upnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteUserchargePayeeNameChange : ""
func (h *Handler) DeleteUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteUserchargePayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["upnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleUserchargePayeeNameChange :""
func (h *Handler) GetSingleUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ppnc := new(models.RefUserchargePayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	ppnc, err := h.Service.GetSingleUserchargePayeeNameChange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userchargepayeenamechange"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// FilterUserchargePayeeNameChange : ""
func (h *Handler) FilterUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.UserchargePayeeNameChangeFilter
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

	var ppncs []models.RefUserchargePayeeNameChange
	log.Println(pagination)
	ppncs, err = h.Service.FilterUserchargePayeeNameChange(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ppncs) > 0 {
		m["upnc"] = ppncs
	} else {
		res := make([]models.UserchargePayeeNameChange, 0)
		m["upnc"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// ApproveTradeLicense : ""
func (h *Handler) ApproveUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	approve := new(models.ApproveUserchargePayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&approve)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ApproveUserchargePayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["upnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// NotApproveTradeLicense : ""
func (h *Handler) NotApproveUserchargePayeeNameChange(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.NotApproveUserchargePayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ppnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}
