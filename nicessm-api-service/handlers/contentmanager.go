package handlers

import (
	"encoding/json"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

// ContentManagerCount : ""
func (h *Handler) ContentManagerCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var content *models.ContentFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&content)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Contents []models.ContentCount
	Contents, err = h.Service.ContentManagerCount(ctx, content)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Contents) > 0 {
		m["content"] = Contents[0]
	} else {
		res := new(models.Content)
		m["content"] = res
	}

	response.With200V2(w, "Success", m, platform)
}

// func (h *Handler) ContentManagerCount(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	var content *models.Content

// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())
// 	var Contents []models.RefContentCount

// 	Contents, err := h.Service.ContentManagerCount(ctx, content)
// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			response.With500mV2(w, "failed no data for this id", platform)
// 			return
// 		}
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["contentmanagercount"] = Contents
// 	response.With200V2(w, "Success", m, platform)
// }
