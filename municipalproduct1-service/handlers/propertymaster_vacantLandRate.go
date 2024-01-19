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

//SaveVacantLandRate : ""
func (h *Handler) SaveVacantLandRate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vacantLandRate := new(models.VacantLandRate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&vacantLandRate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveVacantLandRate(ctx, vacantLandRate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vacantLandRate"] = vacantLandRate
	response.With200V2(w, "Success", m, platform)
}

//UpdateVacantLandRate :""
func (h *Handler) UpdateVacantLandRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	vacantLandRate := new(models.VacantLandRate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&vacantLandRate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if vacantLandRate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateVacantLandRate(ctx, vacantLandRate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vacantLandRate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableVacantLandRate : ""
func (h *Handler) EnableVacantLandRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableVacantLandRate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vacantLandRate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableVacantLandRate : ""
func (h *Handler) DisableVacantLandRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableVacantLandRate(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vacantLandRate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteVacantLandRate : ""
func (h *Handler) DeleteVacantLandRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteVacantLandRate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vacantLandRate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleVacantLandRate :""
func (h *Handler) GetSingleVacantLandRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	vacantLandRate := new(models.RefVacantLandRate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	vacantLandRate, err := h.Service.GetSingleVacantLandRate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vacantLandRate"] = vacantLandRate
	response.With200V2(w, "Success", m, platform)
}

//FilterVacantLandRate : ""
func (h *Handler) FilterVacantLandRate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var vacantLandRate *models.VacantLandRateFilter
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
	err := json.NewDecoder(r.Body).Decode(&vacantLandRate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var vacantLandRates []models.RefVacantLandRate
	log.Println(pagination)
	vacantLandRates, err = h.Service.FilterVacantLandRate(ctx, vacantLandRate, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(vacantLandRates) > 0 {
		m["vacantLandRate"] = vacantLandRates
	} else {
		res := make([]models.VacantLandRate, 0)
		m["vacantLandRate"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
