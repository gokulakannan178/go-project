package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"net/http"
)

func (h *Handler) GetCollectionCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	collection := r.URL.Query().Get("collection")

	if collection == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.Dashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetCollectionCount(ctx, collection)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m[collection] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetDumbSiteCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	collection := r.URL.Query().Get("collection")

	if collection == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.DumbSiteCount)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetDumbSiteCount(ctx, collection)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m[collection] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetHousevisitedCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	collection := r.URL.Query().Get("collection")

	if collection == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.HousevisitedCount)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetHousevisitedCount(ctx, collection)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m[collection] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetvehicleCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	collection := r.URL.Query().Get("collection")

	if collection == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.Dashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetvehicleCount(ctx, collection)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m[collection] = task
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetUsertypeCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	collection := r.URL.Query().Get("collection")

	if collection == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.UserTypeCount)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetUsertypeCount(ctx, collection)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m[collection] = task
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetPropertyCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.FilterProperties
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DayWiseFareCollection *models.PropertyCount
	//log.Println(pagination)
	DayWiseFareCollection, err = h.Service.GetPropertyCount(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = DayWiseFareCollection
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetGarbaggeCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.FilterHouseVisited
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DayWiseFareCollection *models.PropertyCount
	//log.Println(pagination)
	DayWiseFareCollection, err = h.Service.GetGarbaggeCount(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["garbagge"] = DayWiseFareCollection
	response.With200V2(w, "Success", m, platform)
}
