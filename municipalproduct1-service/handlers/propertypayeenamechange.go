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

// SavePropertyPayeeNameChange : ""
func (h *Handler) SavePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ppnc := new(models.PropertyPayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ppnc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePropertyPayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertypayeenamechange"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyPayeeNameChange :""
func (h *Handler) UpdatePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ppnc := new(models.PropertyPayeeNameChange)
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
	err = h.Service.UpdatePropertyPayeeNameChange(ctx, ppnc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertypayeenamechange"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// EnablePropertyPayeeNameChange : ""
func (h *Handler) EnablePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePropertyPayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertypayeenamechange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePropertyPayeeNameChange : ""
func (h *Handler) DisablePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePropertyPayeeNameChange(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ppnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePropertyPayeeNameChange : ""
func (h *Handler) DeletePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePropertyPayeeNameChange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ppnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePropertyPayeeNameChange :""
func (h *Handler) GetSinglePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ppnc := new(models.RefPropertyPayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	ppnc, err := h.Service.GetSinglePropertyPayeeNameChange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertypayeenamechange"] = ppnc
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyPayeeNameChange : ""
func (h *Handler) FilterPropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyPayeeNameChangeFilter
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

	var ppncs []models.RefPropertyPayeeNameChange
	log.Println(pagination)
	ppncs, err = h.Service.FilterPropertyPayeeNameChange(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ppncs) > 0 {
		m["ppnc"] = ppncs
	} else {
		res := make([]models.PropertyPayeeNameChange, 0)
		m["ppnc"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// ApproveTradeLicense : ""
func (h *Handler) ApprovePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	approve := new(models.ApprovePropertyPayeeNameChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&approve)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ApprovePropertyPayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ppnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// NotApproveTradeLicense : ""
func (h *Handler) NotApprovePropertyPayeeNameChange(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.NotApprovePropertyPayeeNameChange(ctx, approve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ppnc"] = "success"
	response.With200V2(w, "Success", m, platform)
}
