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

//SaveOrganisation : ""
func (h *Handler) SaveOrganisation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	organisation := new(models.Organisation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&organisation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOrganisation(ctx, organisation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisation"] = organisation
	response.With200V2(w, "Success", m, platform)
}

//UpdateOrganisation :""
func (h *Handler) UpdateOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	organisation := new(models.Organisation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&organisation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if organisation.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOrganisation(ctx, organisation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOrganisation : ""
func (h *Handler) EnableOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableOrganisation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableOrganisation : ""
func (h *Handler) DisableOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableOrganisation(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOrganisation : ""
func (h *Handler) DeleteOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOrganisation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOrganisation :""
func (h *Handler) GetSingleOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	organisation := new(models.RefOrganisation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	organisation, err := h.Service.GetSingleOrganisation(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisation"] = organisation
	response.With200V2(w, "Success", m, platform)
}

//FilterOrganisation : ""
func (h *Handler) FilterOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var organisation *models.OrganisationFilter
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
	err := json.NewDecoder(r.Body).Decode(&organisation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var organisations []models.RefOrganisation
	log.Println(pagination)
	organisations, err = h.Service.FilterOrganisation(ctx, organisation, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(organisations) > 0 {
		m["organisation"] = organisations
	} else {
		res := make([]models.Organisation, 0)
		m["organisation"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DashboardOrganisationCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var organisation *models.OrganisationFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&organisation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Organisations []models.DashboardOrganisationCountReport
	Organisations, err = h.Service.DashboardOrganisationCount(ctx, organisation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Organisations) > 0 {
		m["organisation"] = Organisations
	} else {
		res := make([]models.Organisation, 0)
		m["organisation"] = res
	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleOrganisationUniqueCheck(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("name")
	//name := r.URL.Query().Get("name")
	if UniqueID == "" {
		response.With400V2(w, "name is missing", platform)
	}

	//	organisation := new(models.RefOrganisation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	organisation, err := h.Service.GetSingleOrganisationUniqueCheck(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisation"] = organisation
	response.With200V2(w, "Success", m, platform)
}
