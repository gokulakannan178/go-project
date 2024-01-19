package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"
)

//SaveCommonLanguageTranslations : ""
func (h *Handler) SaveCommonLanguageTranslations(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	commonlanguagetranslations := new(models.CommonLanguageTranslationss)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&commonlanguagetranslations)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveCommonLanguageTranslations(ctx, commonlanguagetranslations)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commonlanguagetranslations"] = commonlanguagetranslations
	response.With200V2(w, "Success", m, platform)
}

//UpdateCommonLanguageTranslations :""
func (h *Handler) UpdateCommonLanguageTranslations(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	CommonLanguageTranslations := new(models.CommonLanguageTranslationss)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&CommonLanguageTranslations)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateCommonLanguageTranslations(ctx, CommonLanguageTranslations)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commonlanguagetranslations"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableCommonLanguageTranslations: ""
func (h *Handler) EnableCommonLanguageTranslations(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableCommonLanguageTranslations(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commonlanguagetranslations"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableCommonLanguageTranslations : ""
func (h *Handler) DisableCommonLanguageTranslations(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableCommonLanguageTranslations(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commonlanguagetranslations"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteCommonLanguageTranslations : ""
func (h *Handler) DeleteCommonLanguageTranslations(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.CommonLanguageTranslationss)
	UniqueID := r.URL.Query().Get("id")

	if ID.LanguageType != "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteCommonLanguageTranslations(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commonlanguagetranslations"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleCommonLanguageTranslations :""
func (h *Handler) GetSingleCommonLanguageTranslations(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	CommonLanguageTranslations := new(models.RefCommonLanguageTranslations)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	CommonLanguageTranslations, err := h.Service.GetSingleCommonLanguageTranslations(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commonlanguagetranslations"] = CommonLanguageTranslations
	response.With200V2(w, "Success", m, platform)
}

//FilterCommonLanguageTranslations : ""
func (h *Handler) FilterCommonLanguageTranslations(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var CommonLanguageTranslations *models.CommonLanguageTranslationsFilter
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
	err := json.NewDecoder(r.Body).Decode(&CommonLanguageTranslations)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var CommonLanguageTranslationss []models.RefCommonLanguageTranslations
	log.Println(pagination)
	CommonLanguageTranslationss, err = h.Service.FilterCommonLanguageTranslations(ctx, CommonLanguageTranslations, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(CommonLanguageTranslationss) > 0 {
		m["commonlanguagetranslations"] = CommonLanguageTranslationss
	} else {
		res := make([]models.CommonLanguageTranslationss, 0)
		m["commonlanguagetranslations"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleCommonLanguageTranslationsWithType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("lanaguage")

	if UniqueID == "" {
		response.With400V2(w, "lanaguage type is missing", platform)
	}

	CommonLanguageTranslations := new(models.RefCommonLanguageTranslations)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	CommonLanguageTranslations, err := h.Service.GetSingleCommonLanguageTranslationsWithType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commonlanguagetranslations"] = CommonLanguageTranslations
	response.With200V2(w, "Success", m, platform)
}
