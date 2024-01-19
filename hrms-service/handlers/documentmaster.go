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

//SaveDocumentMaster : ""
func (h *Handler) SaveDocumentMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	documentMaster := new(models.DocumentMaster)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&documentMaster)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveDocumentMaster(ctx, documentMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentMaster"] = documentMaster
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDocumentMaster :""
func (h *Handler) GetSingleDocumentMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	documentmaster := new(models.RefDocumentMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	documentmaster, err := h.Service.GetSingleDocumentMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentMaster"] = documentmaster
	response.With200V2(w, "Success", m, platform)
}

//UpdateDocumentMaster :""
func (h *Handler) UpdateDocumentMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	documentmaster := new(models.DocumentMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&documentmaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if documentmaster.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDocumentMaster(ctx, documentmaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDocumentMaster : ""
func (h *Handler) EnableDocumentMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDocumentMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDocumentMaster : ""
func (h *Handler) DisableDocumentMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDocumentMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDocumentMaster : ""
func (h *Handler) DeleteDocumentMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDocumentMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterDocumentMaster : ""
func (h *Handler) FilterDocumentMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var documentMaster *models.FilterDocumentMaster
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
	err := json.NewDecoder(r.Body).Decode(&documentMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var documentMasters []models.RefDocumentMaster
	log.Println(pagination)
	documentMasters, err = h.Service.FilterDocumentMaster(ctx, documentMaster, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(documentMasters) > 0 {
		m["DocumentMaster"] = documentMasters
	} else {
		res := make([]models.DocumentMaster, 0)
		m["DocumentMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
