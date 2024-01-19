package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
)

//FPOUpdateLocation :""
func (h *Handler) FPOUpdateLocation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	fpoloc := new(models.FPOUpdateLocation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&fpoloc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if fpoloc.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.FPOUpdateLocation(ctx, fpoloc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UpdateLocation"] = "success"
	response.With200V2(w, "Success", m, platform)
}
