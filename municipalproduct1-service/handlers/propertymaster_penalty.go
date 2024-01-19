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

//SavePenalty : ""
func (h *Handler) SavePenalty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	penalty := new(models.Penalty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&penalty)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePenalty(ctx, penalty)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["penalty"] = penalty
	response.With200V2(w, "Success", m, platform)
}

//UpdatePenalty :""
func (h *Handler) UpdatePenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	penalty := new(models.Penalty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&penalty)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if penalty.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePenalty(ctx, penalty)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["penalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePenalty : ""
func (h *Handler) EnablePenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePenalty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["penalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePenalty : ""
func (h *Handler) DisablePenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePenalty(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["penalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePenalty : ""
func (h *Handler) DeletePenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePenalty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["penalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePenalty :""
func (h *Handler) GetSinglePenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	penalty := new(models.RefPenalty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	penalty, err := h.Service.GetSinglePenalty(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["penalty"] = penalty
	response.With200V2(w, "Success", m, platform)
}

//FilterPenalty : ""
func (h *Handler) FilterPenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var penalty *models.PenaltyFilter
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
	err := json.NewDecoder(r.Body).Decode(&penalty)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var penaltys []models.RefPenalty
	log.Println(pagination)
	penaltys, err = h.Service.FilterPenalty(ctx, penalty, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(penaltys) > 0 {
		m["penalty"] = penaltys
	} else {
		res := make([]models.Penalty, 0)
		m["penalty"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
