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

//SaveRelation : ""
func (h *Handler) SaveRelation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	relation := new(models.Relation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&relation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveRelation(ctx, relation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["relation"] = relation
	response.With200V2(w, "Success", m, platform)
}

//UpdateRelation :""
func (h *Handler) UpdateRelation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	relation := new(models.Relation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&relation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if relation.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateRelation(ctx, relation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["relation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableRelation : ""
func (h *Handler) EnableRelation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableRelation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["relation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableRelation : ""
func (h *Handler) DisableRelation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableRelation(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["relation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteRelation : ""
func (h *Handler) DeleteRelation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteRelation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["relation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleRelation :""
func (h *Handler) GetSingleRelation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	relation := new(models.RefRelation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	relation, err := h.Service.GetSingleRelation(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["relation"] = relation
	response.With200V2(w, "Success", m, platform)
}

//FilterRelation : ""
func (h *Handler) FilterRelation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var relation *models.RelationFilter
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
	err := json.NewDecoder(r.Body).Decode(&relation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var relations []models.RefRelation
	log.Println(pagination)
	relations, err = h.Service.FilterRelation(ctx, relation, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(relations) > 0 {
		m["relation"] = relations
	} else {
		res := make([]models.Relation, 0)
		m["relation"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
