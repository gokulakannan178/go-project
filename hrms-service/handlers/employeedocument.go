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

//SaveEmployeeDocuments : ""
func (h *Handler) SaveEmployeeDocuments(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeDocuments := new(models.EmployeeDocuments)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveEmployeeDocuments(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDocuments"] = employeeDocuments
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeDocuments :""
func (h *Handler) UpdateEmployeeDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeDocuments := new(models.EmployeeDocuments)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeDocuments.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeDocuments(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableEmployeeDocuments : ""
func (h *Handler) EnableEmployeeDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableEmployeeDocuments(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableEmployeeDocuments : ""
func (h *Handler) DisableEmployeeDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableEmployeeDocuments(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteEmployeeDocuments : ""
func (h *Handler) DeleteEmployeeDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteEmployeeDocuments(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleEmployeeDocuments :""
func (h *Handler) GetSingleEmployeeDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	employeeDocuments := new(models.RefEmployeeDocuments)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	employeeDocuments, err := h.Service.GetSingleEmployeeDocuments(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = employeeDocuments
	response.With200V2(w, "Success", m, platform)
}

//FilterEmployeeDocuments : ""
func (h *Handler) FilterEmployeeDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var employeeDocuments *models.FilterEmployeeDocuments
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
	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var employeeDocumentss []models.RefEmployeeDocuments
	log.Println(pagination)
	employeeDocumentss, err = h.Service.FilterEmployeeDocuments(ctx, employeeDocuments, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(employeeDocumentss) > 0 {
		m["data"] = employeeDocumentss
	} else {
		res := make([]models.EmployeeDocuments, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeDocumentsList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var employeeDocuments *models.FilterEmployeeDocumentslist
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var employeeDocumentss *models.EmployeeDocumentsList
	//log.Println(pagination)
	employeeDocumentss, err = h.Service.EmployeeDocumentsList(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employeeDocuments"] = employeeDocumentss
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateEmployeeDocumentsWithUpsert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeDocuments := new(models.EmployeeDocuments)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeDocuments.EmployeeID == "" {
		response.With400V2(w, "Employee is missing", platform)
	}
	err = h.Service.UpdateEmployeeDocumentsWithUpsert(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) RemoveEmployeeDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeDocuments := new(models.EmployeeDocuments)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeDocuments.EmployeeID == "" {
		response.With400V2(w, "Employee is missing", platform)
	}
	err = h.Service.RemoveEmployeeDocuments(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
