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

// SaveTask : ""
func (h *Handler) SaveTask(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tk := new(models.Task)
	err := json.NewDecoder(r.Body).Decode(&tk)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err = h.Service.SaveTask(ctx, tk)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// SaveTaskTeamMember : ""
func (h *Handler) SaveTaskTeamMember(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tm := new(models.TaskMember)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&tm)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveTaskTeamMember(ctx, tm)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = tm
	response.With200V2(w, "Success", m, platform)
}

//DisableTaskTeamMember : ""
func (h *Handler) DisableTaskTeamMember(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableTaskTeamMember(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTask : ""
func (h *Handler) GetSingleTask(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// UniqueID := r.URL.Query().Get("id")

	// if UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// 	return
	// }

	task := new(models.RefTask)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleTask(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = task
	response.With200V2(w, "Success", m, platform)
}

// EnableTask : ""
func (h *Handler) EnableTask(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.EnableTask(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableTask : ""
func (h *Handler) DisableTask(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableTask(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteTask : ""
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DeleteTask(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// UpdateTask : ""
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var task *models.Task
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&task)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if task.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateTask(ctx, task)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["task"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterTask : ""
func (h *Handler) FilterTask(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterTask
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

	var fts []models.RefTask
	log.Println(pagination)
	fts, err = h.Service.FilterTask(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["task"] = fts
	} else {
		res := make([]models.Task, 0)
		m["task"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
