package handlers

import (
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
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
