package handlers

import (
	"ecommerce-service/app"
	"ecommerce-service/models"
	"ecommerce-service/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//SaveAdvertisement : ""
func (h *Handler) SaveAdvertisement(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	advertisement := new(models.Advertisement)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&advertisement)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveAdvertisement(ctx, advertisement)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["advertisement"] = advertisement
	response.With200V2(w, "Success", m, platform)
}

//UpdateAdvertisement :""
func (h *Handler) UpdateAdvertisement(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	advertisement := new(models.Advertisement)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&advertisement)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if advertisement.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAdvertisement(ctx, advertisement)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["advertisement"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableAdvertisement : ""
func (h *Handler) EnableAdvertisement(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableAdvertisement(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["advertisement"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAdvertisement : ""
func (h *Handler) DisableAdvertisement(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableAdvertisement(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["advertisement"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAdvertisement : ""
func (h *Handler) DeleteAdvertisement(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteAdvertisement(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["advertisement"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAdvertisement :""
func (h *Handler) GetSingleAdvertisement(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	advertisement := new(models.RefAdvertisement)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	advertisement, err := h.Service.GetSingleAdvertisement(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["advertisement"] = advertisement
	response.With200V2(w, "Success", m, platform)
}

//FilterAdvertisement : ""
func (h *Handler) FilterAdvertisement(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var advertisement *models.AdvertisementFilter
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
	err := json.NewDecoder(r.Body).Decode(&advertisement)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var advertisements []models.RefAdvertisement
	log.Println(pagination)
	advertisements, err = h.Service.FilterAdvertisement(ctx, advertisement, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(advertisements) > 0 {
		m["advertisement"] = advertisements
	} else {
		res := make([]models.Advertisement, 0)
		m["advertisement"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
