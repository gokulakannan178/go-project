package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"
)

//SaveState : ""
func (h *Handler) SaveState(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	state := new(models.State)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&state)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveState(ctx, state)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["state"] = state
	response.With200V2(w, "Success", m, platform)
}

//UpdateState :""
func (h *Handler) UpdateState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	state := new(models.State)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&state)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if state.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateState(ctx, state)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableState : ""
func (h *Handler) EnableState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableState(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableState : ""
func (h *Handler) DisableState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableState(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteState : ""
func (h *Handler) DeleteState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteState(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleState :""
func (h *Handler) GetSingleState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	state := new(models.RefState)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	state, err := h.Service.GetSingleState(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = state
	response.With200V2(w, "Success", m, platform)
}

//FilterState : ""
func (h *Handler) FilterState(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var state *models.StateFilter
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
	log.Println(pagination)
	err := json.NewDecoder(r.Body).Decode(&state)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var states []models.RefState

	states, err = h.Service.FilterState(ctx, state, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(states) > 0 {
		m["data"] = states
	} else {
		res := make([]models.State, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GeoDetatilsReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.StateFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	//	err := json.NewDecoder(r.Body).Decode(&filter)

	// if resType == "pdf" {
	// 	data, err := h.Service.GeoDetatilsReportPDF(ctx, filter)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	w.Write(data)
	// 	w.Header().Set("Content-Type", "application/pdf")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=GeoDetailReport.pdf")
	// }

	if resType == "excel" {
		file, err := h.Service.GeoDetatilsReportExcelV2(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=GeoDetailReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	//	data:=
	data, err := h.Service.GeoDetatilsReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["GeoDetailReport"] = data
	response.With200V2(w, "Success", m, platform)

}
func (h *Handler) GeoUploadExcelWithNames(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	isregex := r.URL.Query().Get("isregex")

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

	err = h.Service.GeoUploadExcelWithNames(ctx, file, func() bool {
		if isregex == "Yes" {
			return true
		}
		return false
	}())
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	if err != nil {
		m["GeoUpload"] = err
	}
	m["GeoUpload"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetWeatherDataWithSeverityType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var state *models.StateWeatherAlertMasterFilterv2
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
	err := json.NewDecoder(r.Body).Decode(&state)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var states []models.GetStateLeveWeatherDataAlert
	states, err = h.Service.GetWeatherDataWithSeverityType(ctx, state, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(states) > 0 {
		m["data"] = states
	} else {
		res := make([]models.GetStateLeveWeatherDataAlert, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
