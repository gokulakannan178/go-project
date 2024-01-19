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

//SaveWhatsappLog : ""
func (h *Handler) SaveWhatsappLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	whatsapplog := new(models.WhatsappLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&whatsapplog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveWhatsappLog(ctx, whatsapplog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["whatsapplog"] = whatsapplog
	response.With200V2(w, "Success", m, platform)
}

//UpdateWhatsappLog :""
func (h *Handler) UpdateWhatsappLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	whatsapplog := new(models.WhatsappLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&whatsapplog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if whatsapplog.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateWhatsappLog(ctx, whatsapplog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["whatsapplog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableWhatsappLog : ""
func (h *Handler) EnableWhatsappLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableWhatsappLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["whatsapplog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableWhatsappLog : ""
func (h *Handler) DisableWhatsappLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableWhatsappLog(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["whatsapplog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteWhatsappLog : ""
func (h *Handler) DeleteWhatsappLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.WhatsappLog)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteWhatsappLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["whatsapplog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOrganisation :""
func (h *Handler) GetSingleWhatsappLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	whatsapplog := new(models.RefWhatsappLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	whatsapplog, err := h.Service.GetSingleWhatsappLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["whatsapplog"] = whatsapplog
	response.With200V2(w, "Success", m, platform)
}

//FilterWhatsappLog : ""
func (h *Handler) FilterWhatsappLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var whatsapplog *models.WhatsappLogFilter
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
	err := json.NewDecoder(r.Body).Decode(&whatsapplog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var whatsapplogs []models.RefWhatsappLog
	log.Println(pagination)
	whatsapplogs, err = h.Service.FilterWhatsappLog(ctx, whatsapplog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(whatsapplogs) > 0 {
		m["WhatsappLog"] = whatsapplogs
	} else {
		res := make([]models.WhatsappLog, 0)
		m["WhatsappLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
