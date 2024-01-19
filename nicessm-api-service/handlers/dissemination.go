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

//SaveDissemination : ""
func (h *Handler) SaveDissemination(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	dissemination := new(models.Dissemination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&dissemination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDissemination(ctx, dissemination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = dissemination
	response.With200V2(w, "Success", m, platform)
}

//SaveDissemination : ""
func (h *Handler) SaveSendNow(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	dissemination := new(models.Dissemination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&dissemination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveSendNow(ctx, dissemination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = dissemination
	response.With200V2(w, "Success", m, platform)
}

//SaveDissemination : ""
func (h *Handler) SaveSendLater(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	dissemination := new(models.Dissemination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&dissemination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveSendLater(ctx, dissemination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = dissemination
	response.With200V2(w, "Success", m, platform)
}

//UpdateDissemination :""
func (h *Handler) UpdateDissemination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	dissemination := new(models.Dissemination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&dissemination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateDissemination(ctx, dissemination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDissemination: ""
func (h *Handler) EnableDissemination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDissemination(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDissemination : ""
func (h *Handler) DisableDissemination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDissemination(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDissemination : ""
func (h *Handler) DeleteDissemination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.Dissemination)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDissemination(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDissemination :""
func (h *Handler) GetSingleDissemination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	dissemination := new(models.RefDissemination)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	dissemination, err := h.Service.GetSingleDissemination(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dissemination"] = dissemination
	response.With200V2(w, "Success", m, platform)
}

//FilterDissemination : ""
func (h *Handler) FilterDissemination(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var dissemination *models.DisseminationFilter
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
	err := json.NewDecoder(r.Body).Decode(&dissemination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var disseminations []models.RefDissemination
	log.Println(pagination)
	disseminations, err = h.Service.FilterDissemination(ctx, dissemination, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(disseminations) > 0 {
		m["dissemination"] = disseminations
	} else {
		res := make([]models.Dissemination, 0)
		m["dissemination"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//DisseminationPDF : ""
func (h *Handler) DisseminationPDF(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	if resType == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.DisseminationFilter)
	if resType == "pdf" {
		data, err := h.Service.DisseminationPDF(ctx, filter, nil)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=disseminations.pdf")
		w.Write(data)

	}
}
func (h *Handler) DisseminationReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var dissemination *models.DisseminationReportFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&dissemination)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	resType := r.URL.Query().Get("resType")
	if resType == "excel" {
		file, err := h.Service.DisseminationReportExcel(ctx, dissemination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=DisseminationReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	var disseminations []models.RefDisseminationReport
	//	log.Println(pagination)
	disseminations, err = h.Service.DisseminationReport(ctx, dissemination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(disseminations) > 0 {
		m["dissemination"] = disseminations
	} else {
		res := make([]models.Dissemination, 0)
		m["dissemination"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
