package handlers

import (
	"bpms-service/app"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//GetSingleCRF :""
func (h *Handler) GetSingleCRF(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	crf := new(models.RefCRF)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	crf, err := h.Service.GetSingleCRF(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["crf"] = crf
	response.With200V2(w, "Success", m, platform)
}

//FilterCRF : ""
func (h *Handler) FilterCRF(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var crf *models.CRFFilter
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
	err := json.NewDecoder(r.Body).Decode(&crf)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var crfs []models.RefCRF
	log.Println(pagination)
	crfs, err = h.Service.FilterCRF(ctx, crf, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(crfs) > 0 {
		m["crf"] = crfs
	} else {
		res := make([]models.CRF, 0)
		m["crf"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//CRF Inspection APIS

//StartPlanCRFInspection : ""
func (h *Handler) StartPlanCRFInspection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	crf := new(models.PlanCRFStartInspectionReqPayload)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&crf)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.StartPlanCRFInspection(ctx, crf)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["crfInstection"] = "started"
	response.With200V2(w, "Success", m, platform)
}

//GetCRFInspectionOfPlan : ""
func (h *Handler) GetCRFInspectionOfPlan(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	planID := r.URL.Query().Get("planId")
	if planID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	deptID := r.URL.Query().Get("deptId")
	if deptID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	checklist, err := h.Service.GetCRFInspectionOfPlan(ctx, planID, deptID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	if checklist != nil {
		m["inspections"] = checklist
	} else {
		m["inspections"] = []string{}

	}
	response.With200V2(w, "Success", m, platform)
}

//SubmitCRFInspection : ""
func (h *Handler) SubmitCRFInspection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	crfi := new(models.CRFInspection)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&crfi)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SubmitCRFInspection(ctx, crfi)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["crfInstection"] = "succedd"
	response.With200V2(w, "Success", m, platform)
}

//EndPlanCRFInspection : ""
func (h *Handler) EndPlanCRFInspection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	crf := new(models.PlanCRFEndInspectionReqPayload)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&crf)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.EndPlanCRFInspection(ctx, crf)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["crfInstection"] = "completed"
	response.With200V2(w, "Success", m, platform)
}
