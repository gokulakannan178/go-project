package handlers

import (
	"encoding/json"
	"fmt"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"log"
	"net/http"
	"strconv"
)

// SaveEmployeeShift : ""
func (h *Handler) SaveEmployeeShift(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeshift := new(models.EmployeeShift)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeshift)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeShift(ctx, employeeshift)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employeeshift"] = employeeshift
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeShift : ""
func (h *Handler) GetSingleEmployeeShift(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeshift := new(models.RefEmployeeShift)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeshift, err := h.Service.GetSingleEmployeeShift(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employeeshift"] = employeeshift
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeShift : ""
func (h *Handler) UpdateEmployeeShift(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeshift := new(models.EmployeeShift)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeshift)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeshift.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeShift(ctx, employeeshift)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employeeshift"] = employeeshift
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeShift : ""
func (h *Handler) EnableEmployeeShift(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.EnableEmployeeShift(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employeeshift"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeShift : ""
func (h *Handler) DisableEmployeeShift(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeShift(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employeeshift"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteEmployeeShift : ""
func (h *Handler) DeleteEmployeeShift(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeShift(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employeeshift"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeShift : ""
func (h *Handler) FilterEmployeeShift(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var EmployeeShift *models.FilterEmployeeShift
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
	err := json.NewDecoder(r.Body).Decode(&EmployeeShift)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefEmployeeShift
	log.Println(pagination)
	fts, err = h.Service.FilterEmployeeShift(ctx, EmployeeShift, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["employeeshift"] = fts
	} else {
		res := make([]models.Designation, 0)
		m["employeeshift"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// // GetSingleEmployeeShift : ""
// func (h *Handler) GetDetailEmployeeShift(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	UniqueID := r.URL.Query().Get("id")

// 	if UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 		return
// 	}

// 	task := new(models.RefEmployeeShift)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	task, err := h.Service.GetDetailEmployeeShift(ctx, UniqueID)
// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			response.With500mV2(w, "failed no data for this id", platform)
// 			return
// 		}
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["EmployeeShift"] = task
// 	response.With200V2(w, "Success", m, platform)
// }
// func (h *Handler) AssignEmployeeShift(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	EmployeeShift := new(models.EmployeeShift)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	err := json.NewDecoder(r.Body).Decode(&EmployeeShift)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	if EmployeeShift.UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}
// 	err = h.Service.AssignEmployeeShift(ctx, EmployeeShift)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["EmployeeShift"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }
