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

//SavePropertyRequiredDocument : ""
func (h *Handler) SavePropertyRequiredDocument(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	propertyrequireddocument := new(models.PropertyRequiredDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&propertyrequireddocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePropertyRequiredDocument(ctx, propertyrequireddocument)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PropertyRequiredDocument"] = propertyrequireddocument
	response.With200V2(w, "Success", m, platform)
}

//UpdatePropertyRequiredDocument :""
func (h *Handler) UpdatePropertyRequiredDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	propertyrequireddocument := new(models.PropertyRequiredDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&propertyrequireddocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if propertyrequireddocument.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyRequiredDocument(ctx, propertyrequireddocument)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PropertyRequiredDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePropertyRequiredDocument : ""
func (h *Handler) EnablePropertyRequiredDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePropertyRequiredDocument(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PropertyRequiredDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePropertyRequiredDocument : ""
func (h *Handler) DisablePropertyRequiredDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePropertyRequiredDocument(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PropertyRequiredDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePropertyRequiredDocument : ""
func (h *Handler) DeletePropertyRequiredDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePropertyRequiredDocument(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PropertyRequiredDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePropertyRequiredDocument :""
func (h *Handler) GetSinglePropertyRequiredDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	propertyrequireddocument := new(models.RefPropertyRequiredDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	propertyrequireddocument, err := h.Service.GetSinglePropertyRequiredDocument(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PropertyRequiredDocument"] = propertyrequireddocument
	response.With200V2(w, "Success", m, platform)
}

//FilterPropertyRequiredDocument : ""
func (h *Handler) FilterPropertyRequiredDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyRequiredDocumentFilter
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

	var propertyrequireddocuments []models.RefPropertyRequiredDocument
	log.Println(pagination)
	propertyrequireddocuments, err = h.Service.FilterPropertyRequiredDocument(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertyrequireddocuments) > 0 {
		m["propertyrequireddocument"] = propertyrequireddocuments
	} else {
		res := make([]models.PropertyRequiredDocument, 0)
		m["propertyrequireddocument"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
