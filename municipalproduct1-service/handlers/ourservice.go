package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SaveOurService : ""
func (h *Handler) SaveOurService(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vars := mux.Vars(r)
	scenario := vars["scenario"]

	Dept := new(models.OurService)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Dept)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveOurService(ctx, scenario, Dept)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OurService"] = Dept
	response.With200V2(w, "Success", m, platform)
}

// GetSingleOurService : ""
func (h *Handler) GetSingleOurService(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	vars := mux.Vars(r)
	scenario := vars["scenario"]

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefOurService)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleOurService(ctx, scenario, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OurService"] = task
	response.With200V2(w, "Success", m, platform)
}

// UpdateOurService : ""
func (h *Handler) UpdateOurService(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	vars := mux.Vars(r)
	scenario := vars["scenario"]

	OurService := new(models.OurService)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&OurService)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if OurService.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOurService(ctx, scenario, OurService)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = OurService
	response.With200V2(w, "Success", m, platform)
}

// EnableOurService : ""
func (h *Handler) EnableOurService(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vars := mux.Vars(r)
	scenario := vars["scenario"]
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
	err := h.Service.EnableOurService(ctx, scenario, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OurService"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableOurService : ""
func (h *Handler) DisableOurService(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vars := mux.Vars(r)
	scenario := vars["scenario"]

	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableOurService(ctx, scenario, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OurService"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DeleteOurService(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	vars := mux.Vars(r)
	scenario := vars["scenario"]

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOurService(ctx, scenario, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterOurService : ""
func (h *Handler) FilterOurService(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	vars := mux.Vars(r)
	scenario := vars["scenario"]

	var ft *models.FilterOurService
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

	var fts []models.RefOurService
	log.Println(pagination)
	fts, err = h.Service.FilterOurService(ctx, scenario, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["OurService"] = fts
	} else {
		res := make([]models.OurService, 0)
		m["OurService"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
