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

// SaveBirthCertificate : ""
func (h *Handler) SaveBirthCertificate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	birthcertificate := new(models.BirthCertificate)
	err := json.NewDecoder(r.Body).Decode(&birthcertificate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveBirthCertificate(ctx, birthcertificate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = birthcertificate
	response.With200V2(w, "Success", m, platform)
}

// GetSingleBirthCertificate : ""
func (h *Handler) GetSingleBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	birthcertificate := new(models.RefBirthCertificate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	birthcertificate, err := h.Service.GetSingleBirthCertificate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = birthcertificate
	response.With200V2(w, "Success", m, platform)
}

// UpdateBirthCertificate : ""
func (h *Handler) UpdateBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	birthcertificate := new(models.BirthCertificate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&birthcertificate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if birthcertificate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateBirthCertificate(ctx, birthcertificate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = birthcertificate
	response.With200V2(w, "Success", m, platform)
}

//EnableBirthCertificate : ""
func (h *Handler) EnableBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableBirthCertificate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableBirthCertificate : ""
func (h *Handler) DisableBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableBirthCertificate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteBirthCertificate : ""
func (h *Handler) DeleteBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteBirthCertificate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterBirthCertificate : ""
func (h *Handler) FilterBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.BirthCertificateFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var birthcertificates []models.RefBirthCertificate
	log.Println(pagination)
	birthcertificates, err = h.Service.FilterBirthCertificate(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(birthcertificates) > 0 {
		m["birthcertificate"] = birthcertificates
	} else {
		res := make([]models.BirthCertificate, 0)
		m["birthcertificate"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) ApproveBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	birthcertificate := new(models.BirthCertificate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&birthcertificate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if birthcertificate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.ApproveBirthCertificate(ctx, birthcertificate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = birthcertificate
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) RejectBirthCertificate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	birthcertificate := new(models.BirthCertificate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&birthcertificate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if birthcertificate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.RejectBirthCertificate(ctx, birthcertificate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["birthcertificate"] = birthcertificate
	response.With200V2(w, "Success", m, platform)
}
