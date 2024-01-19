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

//SaveDocumentMuxMaster : ""
func (h *Handler) SaveDocumentMuxMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	documentMuxMaster := new(models.DocumentMuxMaster)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&documentMuxMaster)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveDocumentMuxMaster(ctx, documentMuxMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentMuxMaster"] = documentMuxMaster
	response.With200V2(w, "Success", m, platform)
}

//UpdateDocumentMuxMaster :""
func (h *Handler) UpdateDocumentMuxMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	documentMuxMaster := new(models.DocumentMuxMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&documentMuxMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if documentMuxMaster.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDocumentMuxMaster(ctx, documentMuxMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDocumentMuxMaster : ""
func (h *Handler) EnableDocumentMuxMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDocumentMuxMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDocumentMuxMaster : ""
func (h *Handler) DisableDocumentMuxMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDocumentMuxMaster(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDocumentMuxMaster : ""
func (h *Handler) DeleteDocumentMuxMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDocumentMuxMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDocumentMuxMaster :""
func (h *Handler) GetSingleDocumentMuxMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	documentMuxMaster := new(models.RefDocumentMuxMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	documentMuxMaster, err := h.Service.GetSingleDocumentMuxMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = documentMuxMaster
	response.With200V2(w, "Success", m, platform)
}

//FilterDocumentMuxMaster : ""
func (h *Handler) FilterDocumentMuxMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var documentMuxMaster *models.FilterDocumentMuxMaster
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
	err := json.NewDecoder(r.Body).Decode(&documentMuxMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var documentMuxMasters []models.RefDocumentMuxMaster
	log.Println(pagination)
	documentMuxMasters, err = h.Service.FilterDocumentMuxMaster(ctx, documentMuxMaster, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(documentMuxMasters) > 0 {
		m["data"] = documentMuxMasters
	} else {
		res := make([]models.DocumentMuxMaster, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
