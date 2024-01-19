package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

// SaveExpenseCategoryList : ""
func (h *Handler) SaveExpenseCategoryList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	expensecategorylist := new(models.ExpenseCategoryList)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&expensecategorylist)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveExpenseCategoryList(ctx, expensecategorylist)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ExpenseCategoryList"] = expensecategorylist
	response.With200V2(w, "Success", m, platform)
}

// GetSingleExpenseCategoryList : ""
func (h *Handler) GetSingleExpenseCategoryList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefExpenseCategoryList)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleExpenseCategoryList(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ExpenseCategoryList"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateExpenseCategoryList : ""
func (h *Handler) UpdateExpenseCategoryList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ExpenseCategoryList := new(models.ExpenseCategoryList)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&ExpenseCategoryList)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if ExpenseCategoryList.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateExpenseCategoryList(ctx, ExpenseCategoryList)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ExpenseCategoryList"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableExpenseCategoryList : ""
func (h *Handler) EnableExpenseCategoryList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableExpenseCategoryList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ExpenseCategoryList"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableExpenseCategoryList : ""
func (h *Handler) DisableExpenseCategoryList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableExpenseCategoryList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ExpenseCategoryList"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteExpenseCategoryList : ""
func (h *Handler) DeleteExpenseCategoryList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteExpenseCategoryList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ExpenseCategoryList"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterExpenseCategoryList : ""
func (h *Handler) FilterExpenseCategoryList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterExpenseCategoryList
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

	var fts []models.RefExpenseCategoryList
	log.Println(pagination)
	fts, err = h.Service.FilterExpenseCategoryList(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["ExpenseCategoryList"] = fts
	} else {
		res := make([]models.ExpenseCategoryList, 0)
		m["ExpenseCategoryList"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
