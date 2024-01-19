package handlers

import (
	"encoding/json"
	"fmt"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

//SaveHolidays : ""
func (h *Handler) SaveHolidays(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	holidays := new(models.Holidays)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&holidays)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveHolidays(ctx, holidays)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Holidays"] = holidays
	response.With200V2(w, "Success", m, platform)
}

//GetSingleHolidays :""
func (h *Handler) GetSingleHolidays(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	holidays := new(models.RefHolidays)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	holidays, err := h.Service.GetSingleHolidays(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Holidays"] = holidays
	response.With200V2(w, "Success", m, platform)
}

//UpdateHolidays :""
func (h *Handler) UpdateHolidays(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	holidays := new(models.Holidays)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&holidays)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if holidays.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateHolidays(ctx, holidays)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Holidays"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableHolidays : ""
func (h *Handler) EnableHolidays(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableHolidays(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Holidays"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableHolidays : ""
func (h *Handler) DisableHolidays(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableHolidays(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Holidays"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteHolidays : ""
func (h *Handler) DeleteHolidays(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteHolidays(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Holidays"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterHolidays : ""
func (h *Handler) FilterHolidays(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var Holidays *models.FilterHolidays
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
	err := json.NewDecoder(r.Body).Decode(&Holidays)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var holidayss []models.RefHolidays
	log.Println(pagination)
	holidayss, err = h.Service.FilterHolidays(ctx, Holidays, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(holidayss) > 0 {
		m["Holidays"] = holidayss
	} else {
		res := make([]models.Holidays, 0)
		m["Holidays"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetHolidays(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	holidays := new(models.FilterHolidays)

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&holidays)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	holidaysWeek, err := h.Service.GetHolidaysWeek(ctx, holidays)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Holidays"] = holidaysWeek
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) HolidayUpload(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	farmer := h.Service.HoildayUploadExcel(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeUpload"] = farmer
	response.With200V2(w, "Success", m, platform)
}
