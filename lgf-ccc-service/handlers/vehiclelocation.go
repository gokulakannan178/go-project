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

// SaveVehicleLocation : ""
func (h *Handler) SaveVehicleLocation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Dept := new(models.VehicleLocation)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Dept)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveVehicleLocation(ctx, Dept)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelocation"] = Dept
	response.With200V2(w, "Success", m, platform)
}

// GetSingleVehicleLocation : ""
func (h *Handler) GetSingleVehicleLocation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefVehicleLocation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleVehicleLocation(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelocation"] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateVehicleLocation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	dept := new(models.VehicleLocation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&dept)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if dept.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateVehicleLocation(ctx, dept)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelocation"] = dept
	response.With200V2(w, "Success", m, platform)
}

// EnableVehicleLocation : ""
func (h *Handler) EnableVehicleLocation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	fmt.Println(r)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("platform"))

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableVehicleLocation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelocation"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableVehicleLocation : ""
func (h *Handler) DisableVehicleLocation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableVehicleLocation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelocation"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DeleteVehicleLocation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteVehicleLocation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelocation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterVehicleLocation : ""
func (h *Handler) FilterVehicleLocation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterVehicleLocation
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

	var fts []models.RefVehicleLocation
	log.Println(pagination)
	fts, err = h.Service.FilterVehicleLocation(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["vehiclelocation"] = fts
	} else {
		res := make([]models.VehicleLocation, 0)
		m["vehiclelocation"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
