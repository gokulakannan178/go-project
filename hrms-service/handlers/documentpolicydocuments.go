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

//SaveDocumentPolicyDocuments : ""
func (h *Handler) SaveDocumentPolicyDocuments(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	documentpolicydocuments := new(models.DocumentPolicyDocuments)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&documentpolicydocuments)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveDocumentPolicyDocuments(ctx, documentpolicydocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicyDocuments"] = documentpolicydocuments
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDocumentPolicyDocuments :""
func (h *Handler) GetSingleDocumentPolicyDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	documentpolicydocuments := new(models.RefDocumentPolicyDocuments)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	documentpolicydocuments, err := h.Service.GetSingleDocumentPolicyDocuments(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicyDocuments"] = documentpolicydocuments
	response.With200V2(w, "Success", m, platform)
}

//UpdateDocumentPolicyDocuments :""
func (h *Handler) UpdateDocumentPolicyDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	documentpolicydocuments := new(models.DocumentPolicyDocuments)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&documentpolicydocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if documentpolicydocuments.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDocumentPolicyDocuments(ctx, documentpolicydocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicyDocuments"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDocumentPolicyDocuments : ""
func (h *Handler) EnableDocumentPolicyDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDocumentPolicyDocuments(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicyDocuments"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDocumentPolicyDocuments : ""
func (h *Handler) DisableDocumentPolicyDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDocumentPolicyDocuments(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicyDocuments"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDocumentPolicyDocuments : ""
func (h *Handler) DeleteDocumentPolicyDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDocumentPolicyDocuments(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DocumentPolicyDocuments"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterDocumentPolicyDocuments : ""
func (h *Handler) FilterDocumentPolicyDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var documentpolicydocuments *models.FilterDocumentPolicyDocuments
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
	err := json.NewDecoder(r.Body).Decode(&documentpolicydocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var documentpolicydocumentsFilter []models.RefDocumentPolicyDocuments
	log.Println(pagination)
	documentpolicydocumentsFilter, err = h.Service.FilterDocumentPolicyDocuments(ctx, documentpolicydocuments, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(documentpolicydocumentsFilter) > 0 {
		m["DocumentPolicyDocuments"] = documentpolicydocumentsFilter
	} else {
		res := make([]models.DocumentPolicyDocuments, 0)
		m["DocumentPolicyDocuments"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
