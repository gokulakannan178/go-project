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

//SaveLeaseRent : ""
func (h *Handler) SaveLeaseRent(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	leaserent := new(models.LeaseRent)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&leaserent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveLeaseRent(ctx, leaserent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["leaserent"] = leaserent
	response.With200V2(w, "Success", m, platform)
}

//UpdateLeaseRent :""
func (h *Handler) UpdateLeaseRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	leaserent := new(models.LeaseRent)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&leaserent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if leaserent.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateLeaseRent(ctx, leaserent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableLeaseRent : ""
func (h *Handler) EnableLeaseRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableLeaseRent(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableLeaseRent : ""
func (h *Handler) DisableLeaseRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableLeaseRent(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteLeaseRent : ""
func (h *Handler) DeleteLeaseRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteLeaseRent(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleLeaseRent :""
func (h *Handler) GetSingleLeaseRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	leaserent := new(models.RefLeaseRent)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	leaserent, err := h.Service.GetSingleLeaseRent(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = leaserent
	response.With200V2(w, "Success", m, platform)
}

//FilterLeaseRent : ""
func (h *Handler) FilterLeaseRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var leaserent *models.LeaseRentFilter
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
	err := json.NewDecoder(r.Body).Decode(&leaserent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var leaserents []models.RefLeaseRent
	log.Println(pagination)
	leaserents, err = h.Service.FilterLeaseRent(ctx, leaserent, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(leaserents) > 0 {
		m["data"] = leaserents
	} else {
		res := make([]models.LeaseRent, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
