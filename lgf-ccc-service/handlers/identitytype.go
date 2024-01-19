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

// SaveIdentityType : ""
func (h *Handler) SaveIdentityType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	identitytype := new(models.IdentityType)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&identitytype)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveIdentityType(ctx, identitytype)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["identitytype"] = identitytype
	response.With200V2(w, "Success", m, platform)
}

// GetSingleIdentityType : ""
func (h *Handler) GetSingleIdentityType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	identitytype := new(models.RefIdentityType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	identitytype, err := h.Service.GetSingleIdentityType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["identitytype"] = identitytype
	response.With200V2(w, "Success", m, platform)
}

//UpdateIdentityType : ""
func (h *Handler) UpdateIdentityType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	identitytype := new(models.IdentityType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&identitytype)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if identitytype.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateIdentityType(ctx, identitytype)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["identitytype"] = identitytype
	response.With200V2(w, "Success", m, platform)
}

// EnableIdentityType : ""
func (h *Handler) EnableIdentityType(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.EnableIdentityType(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["IdentityType"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableIdentityType : ""
func (h *Handler) DisableIdentityType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableIdentityType(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["identitytype"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteIdentityType : ""
func (h *Handler) DeleteIdentityType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteIdentityType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["identitytype"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterIdentityType : ""
func (h *Handler) FilterIdentityType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var IdentityType *models.FilterIdentityType
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
	err := json.NewDecoder(r.Body).Decode(&IdentityType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefIdentityType
	log.Println(pagination)
	fts, err = h.Service.FilterIdentityType(ctx, IdentityType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["identitytype"] = fts
	} else {
		res := make([]models.Designation, 0)
		m["identitytype"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// // GetSingleIdentityType : ""
// func (h *Handler) GetDetailIdentityType(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	UniqueID := r.URL.Query().Get("id")

// 	if UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 		return
// 	}

// 	task := new(models.RefIdentityType)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	task, err := h.Service.GetDetailIdentityType(ctx, UniqueID)
// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			response.With500mV2(w, "failed no data for this id", platform)
// 			return
// 		}
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["IdentityType"] = task
// 	response.With200V2(w, "Success", m, platform)
// }
// func (h *Handler) AssignIdentityType(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	IdentityType := new(models.IdentityType)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)

// 	err := json.NewDecoder(r.Body).Decode(&IdentityType)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	if IdentityType.UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}
// 	err = h.Service.AssignIdentityType(ctx, IdentityType)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["IdentityType"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }
