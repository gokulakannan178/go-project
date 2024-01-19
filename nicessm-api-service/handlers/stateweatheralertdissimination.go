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

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveStateWeatherAlertDissimination : ""
func (h *Handler) SaveStateWeatherAlertDissimination(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	StateWeatherAlertDissimination := new(models.StateWeatherAlertDissimination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertDissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveStateWeatherAlertDissimination(ctx, StateWeatherAlertDissimination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = StateWeatherAlertDissimination
	response.With200V2(w, "Success", m, platform)
}

//UpdateStateWeatherAlertDissimination :""
func (h *Handler) UpdateStateWeatherAlertDissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	StateWeatherAlertDissimination := new(models.StateWeatherAlertDissimination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertDissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if StateWeatherAlertDissimination.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateStateWeatherAlertDissimination(ctx, StateWeatherAlertDissimination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableStateWeatherAlertDissimination : ""
func (h *Handler) EnableStateWeatherAlertDissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableStateWeatherAlertDissimination(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableStateWeatherAlertDissimination : ""
func (h *Handler) DisableStateWeatherAlertDissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableStateWeatherAlertDissimination(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteStateWeatherAlertDissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.StateWeatherAlertDissimination)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteStateWeatherAlertDissimination(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleStateWeatherAlertDissimination :""
func (h *Handler) GetSingleStateWeatherAlertDissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	StateWeatherAlertDissimination := new(models.RefStateWeatherAlertDissimination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	StateWeatherAlertDissimination, err := h.Service.GetSingleStateWeatherAlertDissimination(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = StateWeatherAlertDissimination
	response.With200V2(w, "Success", m, platform)
}

//FilterStateWeatherAlertDissimination : ""
func (h *Handler) FilterStateWeatherAlertDissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var StateWeatherAlertDissimination *models.StateWeatherAlertDissiminationFilter
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
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertDissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var StateWeatherAlertDissiminations []models.RefStateWeatherAlertDissimination
	log.Println(pagination)
	StateWeatherAlertDissiminations, err = h.Service.FilterStateWeatherAlertDissimination(ctx, StateWeatherAlertDissimination, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(StateWeatherAlertDissiminations) > 0 {
		m["StateWeatherAlertDissimination"] = StateWeatherAlertDissiminations
	} else {
		res := make([]models.StateWeatherAlertDissimination, 0)
		m["StateWeatherAlertDissimination"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveStateWeatherAlertDissiminationSendNow(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	StateWeatherAlertDissimination := new(models.SendStateWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertDissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	fmt.Println("state===>", StateWeatherAlertDissimination.WeatherAlert.State.ID.Hex())
	DissseminationUserFarmer, err := h.Service.SaveStateWeatherAlertDissiminationSendNow(ctx, &StateWeatherAlertDissimination.WeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = DissseminationUserFarmer
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetWeatherAlertFarmerUserCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.StateWeatherAlertDissimination)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	StateWeatherAlertDissimination, err := h.Service.GetSingleWeatherAlertFarmerUserCount(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertDissimination"] = StateWeatherAlertDissimination
	response.With200V2(w, "Success", m, platform)
}

//FilterStateWeatherAlertDissiminationReport
func (h *Handler) FilterStateWeatherAlertDissiminationReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var StateWeatherAlertDissimination *models.StateWeatherAlertDissiminationFilter
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
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertDissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var StateWeatherAlertDissiminations []models.RefStateWeatherAlertDissimination
	log.Println(pagination)
	StateWeatherAlertDissiminations, err = h.Service.FilterStateWeatherAlertDissiminationReport(ctx, StateWeatherAlertDissimination, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	resType := r.URL.Query().Get("resType")
	if resType == "excel" {
		file, err := h.Service.FilterStateWeatherAlertDissiminationReportExcel(ctx, StateWeatherAlertDissimination, nil)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=WeatherAlertDisseminationReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	m := make(map[string]interface{})

	if len(StateWeatherAlertDissiminations) > 0 {
		m["StateWeatherAlertDissimination"] = StateWeatherAlertDissiminations
	} else {
		res := make([]models.StateWeatherAlertDissimination, 0)
		m["StateWeatherAlertDissimination"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
