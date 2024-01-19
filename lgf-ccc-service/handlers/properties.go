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

// SaveProperties : ""
func (h *Handler) SaveProperties(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	properties := new(models.Properties)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&properties)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveProperties(ctx, properties)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["properties"] = properties
	response.With200V2(w, "Success", m, platform)
}

// GetSingleProperties : ""
func (h *Handler) GetSingleProperties(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefProperties)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleProperties(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["properties"] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSinglePropertiesWithHoldingNumber(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	holdingNumber := r.URL.Query().Get("id")

	if holdingNumber == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefProperties)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSinglePropertiesWithHoldingNumber(ctx, holdingNumber)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["properties"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateProperties : ""
func (h *Handler) UpdateProperties(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	properties := new(models.Properties)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&properties)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if properties.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateProperties(ctx, properties)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["properties"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableProperties : ""
func (h *Handler) EnableProperties(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.EnableProperties(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["properties"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableProperties : ""
func (h *Handler) DisableProperties(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableProperties(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["properties"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteProperties : ""
func (h *Handler) DeleteProperties(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteProperties(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["properties"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// // InitProperties : ""
// func (h *Handler) InitProperties(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	ID := r.URL.Query().Get("id")
// 	fmt.Println(r)
// 	fmt.Println(r.URL)
// 	fmt.Println(r.URL.Query())
// 	fmt.Println(r.URL.Query().Get("platform"))

// 	ctx := app.GetApp(r.Context(), h.Service.Daos)
// 	if ID == "" {
// 		response.With400V2(w, "ID is missing", platform)
// 		return
// 	}
// 	err := h.Service.InitProperties(ctx, ID)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["Properties"] = "Success"
// 	response.With200V2(w, "Success", m, platform)
// }

// // PendingProperties : ""
// func (h *Handler) PendingProperties(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	ID := r.URL.Query().Get("id")

// 	ctx := app.GetApp(r.Context(), h.Service.Daos)
// 	if ID == "" {
// 		response.With400V2(w, "ID is missing", platform)
// 		return
// 	}
// 	err := h.Service.PendingProperties(ctx, ID)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["Properties"] = "Success"
// 	response.With200V2(w, "Success", m, platform)
// }

// //InProgressProperties : ""
// func (h *Handler) InProgressProperties(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	UniqueID := r.URL.Query().Get("id")

// 	if UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}

// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	err := h.Service.InProgressProperties(ctx, UniqueID)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["Properties"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }

// //CompletedProperties : ""
// func (h *Handler) CompletedProperties(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	Properties := new(models.Properties)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	err := json.NewDecoder(r.Body).Decode(&Properties)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	if Properties.UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}
// 	err = h.Service.CompletedProperties(ctx, Properties)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["Properties"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }

// FilterProperties : ""
func (h *Handler) FilterProperties(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var properties *models.FilterProperties
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
	err := json.NewDecoder(r.Body).Decode(&properties)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefProperties
	log.Println(pagination)
	fts, err = h.Service.FilterProperties(ctx, properties, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["properties"] = fts
	} else {
		res := make([]models.Designation, 0)
		m["properties"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// // GetSingleProperties : ""
// func (h *Handler) GetDetailProperties(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	UniqueID := r.URL.Query().Get("id")

// 	if UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 		return
// 	}

// 	task := new(models.RefProperties)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	task, err := h.Service.GetDetailProperties(ctx, UniqueID)
// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			response.With500mV2(w, "failed no data for this id", platform)
// 			return
// 		}
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["Properties"] = task
// 	response.With200V2(w, "Success", m, platform)
// }
// func (h *Handler) AssignProperties(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	Properties := new(models.Properties)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	err := json.NewDecoder(r.Body).Decode(&Properties)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	if Properties.UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}
// 	err = h.Service.AssignProperties(ctx, Properties)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["Properties"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }
