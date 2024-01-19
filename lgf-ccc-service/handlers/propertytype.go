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

// SavePropertyType : ""
func (h *Handler) SavePropertyType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	propertytype := new(models.PropertyType)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&propertytype)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SavePropertyType(ctx, propertytype)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertytype"] = propertytype
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePropertyType : ""
func (h *Handler) GetSinglePropertyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	propertytype := new(models.RefPropertyType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	propertytype, err := h.Service.GetSinglePropertyType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertytype"] = propertytype
	response.With200V2(w, "Success", m, platform)
}

// func (h *Handler) GetSinglePropertyTypeWithHoldingNumber(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	holdingNumber := r.URL.Query().Get("id")

// 	if holdingNumber == "" {
// 		response.With400V2(w, "id is missing", platform)
// 		return
// 	}

// 	task := new(models.RefPropertyType)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	task, err := h.Service.GetSinglePropertyTypeWithHoldingNumber(ctx, holdingNumber)
// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			response.With500mV2(w, "failed no data for this id", platform)
// 			return
// 		}
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["PropertyType"] = task
// 	response.With200V2(w, "Success", m, platform)
// }

//UpdatePropertyType : ""
func (h *Handler) UpdatePropertyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	propertytype := new(models.PropertyType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&propertytype)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if propertytype.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyType(ctx, propertytype)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertytype"] = propertytype
	response.With200V2(w, "Success", m, platform)
}

// EnablePropertyType : ""
func (h *Handler) EnablePropertyType(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.EnablePropertyType(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertytype"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePropertyType : ""
func (h *Handler) DisablePropertyType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisablePropertyType(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertytype"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePropertyType : ""
func (h *Handler) DeletePropertyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePropertyType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertytype"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyType : ""
func (h *Handler) FilterPropertyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var propertytype *models.FilterPropertyType
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
	err := json.NewDecoder(r.Body).Decode(&propertytype)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefPropertyType
	log.Println(pagination)
	fts, err = h.Service.FilterPropertyType(ctx, propertytype, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["propertytype"] = fts
	} else {
		res := make([]models.Designation, 0)
		m["propertytype"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// // GetSinglePropertyType : ""
// func (h *Handler) GetDetailPropertyType(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	UniqueID := r.URL.Query().Get("id")

// 	if UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 		return
// 	}

// 	task := new(models.RefPropertyType)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	task, err := h.Service.GetDetailPropertyType(ctx, UniqueID)
// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			response.With500mV2(w, "failed no data for this id", platform)
// 			return
// 		}
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["PropertyType"] = task
// 	response.With200V2(w, "Success", m, platform)
// }
// func (h *Handler) AssignPropertyType(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	PropertyType := new(models.PropertyType)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	err := json.NewDecoder(r.Body).Decode(&PropertyType)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	if PropertyType.UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}
// 	err = h.Service.AssignPropertyType(ctx, PropertyType)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["PropertyType"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }
