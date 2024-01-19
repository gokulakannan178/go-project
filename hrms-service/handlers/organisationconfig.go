package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"

	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) SaveOrganisationConfig(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	organisationConfig := new(models.OrganisationConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&organisationConfig)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOrganisationConfig(ctx, organisationConfig)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationConfig"] = organisationConfig
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) UpdateOrganisationConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	organisationConfig := new(models.OrganisationConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&organisationConfig)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateOrganisationConfig(ctx, organisationConfig)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) EnableOrganisationConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableOrganisationConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DisableOrganisationConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableOrganisationConfig(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DeleteOrganisationConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.OrganisationConfig)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteOrganisationConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetSingleOrganisationConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	organisationConfig := new(models.RefOrganisationConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	organisationConfig, err := h.Service.GetSingleOrganisationConfig(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationConfig"] = organisationConfig
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) FilterOrganisationConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var organisationConfig *models.OrganisationConfigFilter
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
	err := json.NewDecoder(r.Body).Decode(&organisationConfig)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var organisationConfigs []models.RefOrganisationConfig
	log.Println(pagination)
	organisationConfigs, err = h.Service.FilterOrganisationConfig(ctx, organisationConfig, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(organisationConfigs) > 0 {
		m["organisationConfig"] = organisationConfigs
	} else {
		res := make([]models.OrganisationConfig, 0)
		m["organisationConfig"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SetdefaultOrganisationConfig(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.SetdefaultOrganisationConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationConfig"] = "success"
	response.With200V2(w, "Success", m, platform)

}

func (h *Handler) GetactiveOrganisationConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	IsDefault := true

	// if IsDefault !=nil {
	// 	response.With400V2(w, "invalid IsDefault", platform)
	// }

	var organisation *models.OrganisationConfig
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	organisation, err := h.Service.GetactiveOrganisationConfig(ctx, IsDefault)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["organisationconfig"] = organisation
	response.With200V2(w, "Success", m, platform)
}
