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

// SaveCitizenGrievance : ""
func (h *Handler) SaveCitizenGrievance(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	citizen := new(models.CitizenGrievance)
	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveCitizenGrievance(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleCitizenGrievance : ""
func (h *Handler) GetSingleCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	citizen := new(models.RefCitizenGrievance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	citizen, err := h.Service.GetSingleCitizenGrievance(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = citizen
	response.With200V2(w, "Success", m, platform)
}

// UpdateCitizenGrievance : ""
func (h *Handler) UpdateCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	citizen := new(models.CitizenGrievance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if citizen.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateCitizenGrievance(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableCitizenGrievance : ""
func (h *Handler) EnableCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// UniqueID := r.URL.Query().Get("id")
	citizen := new(models.CitizenGrievance)
	// if citizen.UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// }

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.EnableCitizenGrievance(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableCitizenGrievance : ""
func (h *Handler) DisableCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	citizen := new(models.RejectedCitizengravians)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.DisableCitizenGrievance(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteCitizenGrievance : ""
func (h *Handler) DeleteCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	citizen := new(models.RejectedCitizengravians)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.DeleteCitizenGrievance(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// CompletedCitizenGrievance : ""
func (h *Handler) CompletedCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	citizen := new(models.CitizenGrievance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.CompletedCitizenGrievance(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectedCitizenGrievance : ""
func (h *Handler) RejectedCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	citizen := new(models.RejectedCitizengravians)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	// if citizen.UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// }
	err = h.Service.RejectedCitizenGrievance(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterCitizenGrievance : ""
func (h *Handler) FilterCitizenGrievance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.CitizenGrievanceFilter
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

	var users []models.RefCitizenGrievance
	log.Println(pagination)
	users, err = h.Service.FilterCitizenGrievance(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(users) > 0 {
		m["citizenGrievance"] = users
	} else {
		res := make([]models.RefCitizenGrievance, 0)
		m["citizenGrievance"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UpdateCitizenGrievanceSolution : ""
func (h *Handler) UpdateCitizenGrievanceSolution(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	citizen := new(models.CitizenGrievanceSolution)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizen)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if citizen.CitizenGrievanceID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateCitizenGrievanceSolution(ctx, citizen)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizenGrievanceSolution"] = "success"
	response.With200V2(w, "Success", m, platform)
}
