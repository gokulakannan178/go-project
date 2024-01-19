package handlers

import (
	"encoding/json"
	"fmt"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"log"
	"net/http"
	"strconv"
)

// SaveServiceRequest : ""
func (h *Handler) SaveServiceRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	serviceRequest := new(models.ServiceRequest)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&serviceRequest)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveServiceRequest(ctx, serviceRequest)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["serviceRequest"] = serviceRequest
	response.With200V2(w, "Success", m, platform)
}

// GetSingleServiceRequest : ""
func (h *Handler) GetSingleServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefServiceRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleServiceRequest(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateServiceRequest : ""
func (h *Handler) UpdateServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	serviceRequest := new(models.ServiceRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&serviceRequest)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if serviceRequest.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateServiceRequest(ctx, serviceRequest)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableServiceRequest : ""
func (h *Handler) EnableServiceRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	fmt.Println(r)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("platform"))

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableServiceRequest(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableServiceRequest : ""
func (h *Handler) DisableServiceRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableServiceRequest(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteServiceRequest : ""
func (h *Handler) DeleteServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteServiceRequest(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// InitServiceRequest : ""
func (h *Handler) InitServiceRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	fmt.Println(r)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("platform"))

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.InitServiceRequest(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// PendingServiceRequest : ""
func (h *Handler) PendingServiceRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.PendingServiceRequest(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//InProgressServiceRequest : ""
func (h *Handler) InProgressServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.InProgressServiceRequest(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//CompletedServiceRequest : ""
func (h *Handler) CompletedServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	serviceRequest := new(models.ServiceRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&serviceRequest)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if serviceRequest.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.CompletedServiceRequest(ctx, serviceRequest)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterServiceRequest : ""
func (h *Handler) FilterServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterServiceRequest
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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
	err := json.NewDecoder(r.Body).Decode(&ft)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefServiceRequest
	log.Println(pagination)
	fts, err = h.Service.FilterServiceRequest(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["ServiceRequest"] = fts
	} else {
		res := make([]models.Designation, 0)
		m["ServiceRequest"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetSingleServiceRequest : ""
func (h *Handler) GetDetailServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefServiceRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetDetailServiceRequest(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) AssignServiceRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	serviceRequest := new(models.ServiceRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&serviceRequest)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if serviceRequest.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.AssignServiceRequest(ctx, serviceRequest)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ServiceRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}
