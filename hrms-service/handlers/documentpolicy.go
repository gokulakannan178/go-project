package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

//SaveDocumentPolicy : ""
func (h *Handler) SaveDocumentPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	documentPolicy := new(models.DocumentPolicy)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&documentPolicy)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveDocumentPolicy(ctx, documentPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicy"] = documentPolicy
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDocumentPolicy :""
func (h *Handler) GetSingleDocumentPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	documentPolicy := new(models.RefDocumentPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	documentPolicy, err := h.Service.GetSingleDocumentPolicy(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicy"] = documentPolicy
	response.With200V2(w, "Success", m, platform)
}

//UpdateDocumentPolicy :""
func (h *Handler) UpdateDocumentPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	documentPolicy := new(models.DocumentPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&documentPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if documentPolicy.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDocumentPolicy(ctx, documentPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDocumentPolicy : ""
func (h *Handler) EnableDocumentPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDocumentPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDocumentPolicy : ""
func (h *Handler) DisableDocumentPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDocumentPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDocumentPolicy : ""
func (h *Handler) DeleteDocumentPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDocumentPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterDocumentPolicy : ""
func (h *Handler) FilterDocumentPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var documentPolicy *models.FilterDocumentPolicy
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
	err := json.NewDecoder(r.Body).Decode(&documentPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var documentPolicys []models.RefDocumentPolicy
	log.Println(pagination)
	documentPolicys, err = h.Service.FilterDocumentPolicy(ctx, documentPolicy, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(documentPolicys) > 0 {
		m["DocumentPolicy"] = documentPolicys
	} else {
		res := make([]models.DocumentPolicy, 0)
		m["DocumentPolicy"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
