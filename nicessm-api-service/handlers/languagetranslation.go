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

//SaveLanguageTranslation : ""
func (h *Handler) SaveLanguageTranslation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	languagetranslation := new(models.LanguageTranslations)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&languagetranslation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveLanguageTranslation(ctx, languagetranslation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["languagetranslation"] = languagetranslation
	response.With200V2(w, "Success", m, platform)
}

//UpdateLanguageTranslation :""
func (h *Handler) UpdateLanguageTranslation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	LanguageTranslation := new(models.LanguageTranslations)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&LanguageTranslation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateLanguageTranslation(ctx, LanguageTranslation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["languagetranslation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableLanguageTranslation: ""
func (h *Handler) EnableLanguageTranslation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableLanguageTranslation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["languagetranslation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableLanguageTranslation : ""
func (h *Handler) DisableLanguageTranslation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableLanguageTranslation(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["languagetranslation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteLanguageTranslation : ""
func (h *Handler) DeleteLanguageTranslation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.LanguageTranslations)
	UniqueID := r.URL.Query().Get("id")

	if ID.LanguageType != "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteLanguageTranslation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["languagetranslation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleLanguageTranslation :""
func (h *Handler) GetSingleLanguageTranslation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	LanguageTranslation := new(models.RefLanguageTranslation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	LanguageTranslation, err := h.Service.GetSingleLanguageTranslation(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["languagetranslation"] = LanguageTranslation
	response.With200V2(w, "Success", m, platform)
}

//FilterLanguageTranslation : ""
func (h *Handler) FilterLanguageTranslation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var LanguageTranslation *models.LanguageTranslationFilter
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
	err := json.NewDecoder(r.Body).Decode(&LanguageTranslation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var LanguageTranslations []models.RefLanguageTranslation
	log.Println(pagination)
	LanguageTranslations, err = h.Service.FilterLanguageTranslation(ctx, LanguageTranslation, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(LanguageTranslations) > 0 {
		m["languagetranslation"] = LanguageTranslations
	} else {
		res := make([]models.LanguageTranslations, 0)
		m["languagetranslation"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleLanguageTranslationWithType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("lanaguage")

	if UniqueID == "" {
		response.With400V2(w, "lanaguage type is missing", platform)
	}

	LanguageTranslation := new(models.RefLanguageTranslation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	LanguageTranslation, err := h.Service.GetSingleLanguageTranslationWithType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["languagetranslation"] = LanguageTranslation
	response.With200V2(w, "Success", m, platform)
}
