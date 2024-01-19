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

//SaveFarmerCropCalendar : ""
func (h *Handler) SaveFarmerCropCalendar(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	farmerCropCalendar := new(models.FarmerCropCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&farmerCropCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveFarmerCropCalendar(ctx, farmerCropCalendar)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCropCalendar"] = farmerCropCalendar
	response.With200V2(w, "Success", m, platform)
}

//UpdateFarmerCropCalendar :""
func (h *Handler) UpdateFarmerCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	farmerCropCalendar := new(models.FarmerCropCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&farmerCropCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if farmerCropCalendar.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateFarmerCropCalendar(ctx, farmerCropCalendar)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFarmerCropCalendar : ""
func (h *Handler) EnableFarmerCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFarmerCropCalendar(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableFarmerCropCalendar : ""
func (h *Handler) DisableFarmerCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFarmerCropCalendar(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteFarmerCropCalendar : ""
func (h *Handler) DeleteFarmerCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFarmerCropCalendar(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleFarmerCropCalendar :""
func (h *Handler) GetSingleFarmerCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	farmerCropCalendar := new(models.RefFarmerCropCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	farmerCropCalendar, err := h.Service.GetSingleFarmerCropCalendar(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCropCalendar"] = farmerCropCalendar
	response.With200V2(w, "Success", m, platform)
}

//FilterFarmerCropCalendar : ""
func (h *Handler) FilterFarmerCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var farmerCropCalendar *models.FarmerCropCalendarFilter
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
	err := json.NewDecoder(r.Body).Decode(&farmerCropCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var farmerCropCalendars []models.RefFarmerCropCalendar
	log.Println(pagination)
	farmerCropCalendars, err = h.Service.FilterFarmerCropCalendar(ctx, farmerCropCalendar, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(farmerCropCalendars) > 0 {
		m["farmerCropCalendar"] = farmerCropCalendars
	} else {
		res := make([]models.FarmerCropCalendar, 0)
		m["farmerCropCalendar"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
