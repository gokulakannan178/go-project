package handlers

import (
	"encoding/json"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

func (h *Handler) DashboardContentSmsCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var content *models.DashboardContentCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&content)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Contents []models.DashboardContentCountReport
	Contents, err = h.Service.DashboardContentSmsCount(ctx, content)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Contents) > 0 {
		m["content"] = Contents
	} else {
		res := make([]models.Content, 0)
		m["content"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DashboardContentVoiceCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var content *models.DashboardContentCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&content)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Contents []models.DashboardContentCountReport
	Contents, err = h.Service.DashboardContentVoiceCount(ctx, content)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Contents) > 0 {
		m["content"] = Contents
	} else {
		res := make([]models.Content, 0)
		m["content"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DashboardContentVideoCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var content *models.DashboardContentCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&content)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Contents []models.DashboardContentCountReport
	Contents, err = h.Service.DashboardContentVideoCount(ctx, content)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Contents) > 0 {
		m["content"] = Contents
	} else {
		res := make([]models.Content, 0)
		m["content"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DashboardContentPosterCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var content *models.DashboardContentCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&content)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Contents []models.DashboardContentCountReport
	Contents, err = h.Service.DashboardContentPosterCount(ctx, content)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Contents) > 0 {
		m["content"] = Contents
	} else {
		res := make([]models.Content, 0)
		m["content"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DashboardContentDocmentCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var content *models.DashboardContentCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&content)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Contents []models.DashboardContentCountReport
	Contents, err = h.Service.DashboardContentDocmentCount(ctx, content)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Contents) > 0 {
		m["content"] = Contents
	} else {
		res := make([]models.Content, 0)
		m["content"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DayWiseContentDemandChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var filter *models.DashboardContentCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// DayWiseContentDemandChart
	data, err := h.Service.DayWiseContentDemandChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["content"] = data
	response.With200V2(w, "Success", m, platform)

}
