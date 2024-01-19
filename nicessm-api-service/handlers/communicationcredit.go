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

//SaveCommunicationCredit : ""
func (h *Handler) SaveCommunicationCredit(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	CommunicationCredit := new(models.CommunicationCredit)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&CommunicationCredit)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveCommunicationCredit(ctx, CommunicationCredit)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["CommunicationCredit"] = CommunicationCredit
	response.With200V2(w, "Success", m, platform)
}

//UpdateCommunicationCredit :""
func (h *Handler) UpdateCommunicationCredit(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	CommunicationCredit := new(models.CommunicationCredit)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&CommunicationCredit)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if CommunicationCredit.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateCommunicationCredit(ctx, CommunicationCredit)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["CommunicationCredit"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableCommunicationCredit : ""
func (h *Handler) EnableCommunicationCredit(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableCommunicationCredit(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["CommunicationCredit"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableCommunicationCredit : ""
func (h *Handler) DisableCommunicationCredit(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableCommunicationCredit(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["CommunicationCredit"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteCommunicationCredit : ""
func (h *Handler) DeleteCommunicationCredit(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.CommunicationCredit)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteCommunicationCredit(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["CommunicationCredit"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOrganisation :""
func (h *Handler) GetSingleCommunicationCredit(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	CommunicationCredit := new(models.RefCommunicationCredit)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	CommunicationCredit, err := h.Service.GetSingleCommunicationCredit(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["CommunicationCredit"] = CommunicationCredit
	response.With200V2(w, "Success", m, platform)
}

//FilterCommunicationCredit : ""
func (h *Handler) FilterCommunicationCredit(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var CommunicationCredit *models.CommunicationCreditFilter
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
	err := json.NewDecoder(r.Body).Decode(&CommunicationCredit)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var CommunicationCredits []models.RefCommunicationCredit
	log.Println(pagination)
	CommunicationCredits, err = h.Service.FilterCommunicationCredit(ctx, CommunicationCredit, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(CommunicationCredits) > 0 {
		m["CommunicationCredit"] = CommunicationCredits
	} else {
		res := make([]models.CommunicationCredit, 0)
		m["CommunicationCredit"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
