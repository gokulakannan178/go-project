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

// SaveLeavePolicy : ""
func (h *Handler) SaveLeavePolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Leavepolicy := new(models.LeavePolicy)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Leavepolicy)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveLeavePolicy(ctx, Leavepolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["LeavePolicy"] = Leavepolicy
	response.With200V2(w, "Success", m, platform)
}

// GetSingleLeavePolicy : ""
func (h *Handler) GetSingleLeavePolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefLeavePolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleLeavePolicy(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["LeavePolicy"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateLeavePolicy : ""
func (h *Handler) UpdateLeavePolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	Leavepolicy := new(models.LeavePolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&Leavepolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if Leavepolicy.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateLeavePolicy(ctx, Leavepolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["LeavePolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableLeavePolicy : ""
func (h *Handler) EnableLeavePolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	fmt.Println(r)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("platform"))

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableLeavePolicy(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["LeavePolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableLeavePolicy : ""
func (h *Handler) DisableLeavePolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableLeavePolicy(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["LeavePolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteLeavePolicy : ""
func (h *Handler) DeleteLeavePolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteLeavePolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["LeavePolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterLeavePolicy : ""
func (h *Handler) FilterLeavePolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterLeavePolicy
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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
	err := json.NewDecoder(r.Body).Decode(&ft)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefLeavePolicy
	log.Println(pagination)
	fts, err = h.Service.FilterLeavePolicy(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["LeavePolicy"] = fts
	} else {
		res := make([]models.Designation, 0)
		m["LeavePolicy"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
