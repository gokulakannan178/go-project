package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveDistrictweatheralertdissimination : ""
func (h *Handler) SaveDistrictweatheralertdissimination(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Districtweatheralertdissimination := new(models.DistrictWeatherAlertDissimination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Districtweatheralertdissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDistrictweatheralertdissimination(ctx, Districtweatheralertdissimination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = Districtweatheralertdissimination
	response.With200V2(w, "Success", m, platform)
}

//UpdateDistrictweatheralertdissimination :""
func (h *Handler) UpdateDistrictweatheralertdissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	Districtweatheralertdissimination := new(models.DistrictWeatherAlertDissimination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&Districtweatheralertdissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if Districtweatheralertdissimination.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateDistrictweatheralertdissimination(ctx, Districtweatheralertdissimination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDistrictweatheralertdissimination : ""
func (h *Handler) EnableDistrictweatheralertdissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDistrictweatheralertdissimination(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDistrictweatheralertdissimination : ""
func (h *Handler) DisableDistrictweatheralertdissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDistrictweatheralertdissimination(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteDistrictweatheralertdissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.DistrictWeatherAlertDissimination)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDistrictweatheralertdissimination(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDistrictweatheralertdissimination :""
func (h *Handler) GetSingleDistrictweatheralertdissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Districtweatheralertdissimination := new(models.RefDistrictWeatherAlertDissimination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Districtweatheralertdissimination, err := h.Service.GetSingleDistrictweatheralertdissimination(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = Districtweatheralertdissimination
	response.With200V2(w, "Success", m, platform)
}

//FilterDistrictweatheralertdissimination : ""
func (h *Handler) FilterDistrictweatheralertdissimination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var Districtweatheralertdissimination *models.DistrictWeatherAlertDissiminationFilter
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
	err := json.NewDecoder(r.Body).Decode(&Districtweatheralertdissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Districtweatheralertdissiminations []models.RefDistrictWeatherAlertDissimination
	log.Println(pagination)
	Districtweatheralertdissiminations, err = h.Service.FilterDistrictweatheralertdissimination(ctx, Districtweatheralertdissimination, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Districtweatheralertdissiminations) > 0 {
		m["Districtweatheralertdissimination"] = Districtweatheralertdissiminations
	} else {
		res := make([]models.DistrictWeatherAlertDissimination, 0)
		m["Districtweatheralertdissimination"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveDistrictweatheralertdissiminationSendNow(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Districtweatheralertdissimination := new(models.SendDistrictWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Districtweatheralertdissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDistrictweatheralertdissiminationSendNow(ctx, &Districtweatheralertdissimination.WeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetDistrictWeatherAlertFarmerUserCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.DistrictWeatherAlertDissimination)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Districtweatheralertdissimination, err := h.Service.GetSingleDistrictWeatherAlertFarmerUserCount(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Districtweatheralertdissimination"] = Districtweatheralertdissimination
	response.With200V2(w, "Success", m, platform)
}

//FilterDistrictweatheralertdissiminationReport
func (h *Handler) FilterDistrictweatheralertdissiminationReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var Districtweatheralertdissimination *models.DistrictWeatherAlertDissiminationFilter
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
	err := json.NewDecoder(r.Body).Decode(&Districtweatheralertdissimination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Districtweatheralertdissiminations []models.RefDistrictWeatherAlertDissimination
	log.Println(pagination)
	Districtweatheralertdissiminations, err = h.Service.FilterDistrictweatheralertdissiminationReport(ctx, Districtweatheralertdissimination, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	resType := r.URL.Query().Get("resType")
	if resType == "excel" {
		file, err := h.Service.FilterDistrictweatheralertdissiminationReportExcel(ctx, Districtweatheralertdissimination, nil)
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

	if len(Districtweatheralertdissiminations) > 0 {
		m["Districtweatheralertdissimination"] = Districtweatheralertdissiminations
	} else {
		res := make([]models.DistrictWeatherAlertDissimination, 0)
		m["Districtweatheralertdissimination"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
