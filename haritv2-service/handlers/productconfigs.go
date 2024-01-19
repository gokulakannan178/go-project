package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
)

func (h *Handler) SaveProductConfig(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	productconfigs := new(models.ProductConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&productconfigs)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveProductConfig(ctx, productconfigs)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productconfigs"] = productconfigs
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) EnableProductConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableProductConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ProductConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetactiveProductConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	Status := "Active"

	if Status == "" {
		response.With400V2(w, "invalid status", platform)
	}

	var product *models.ProductConfig
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	product, err := h.Service.GetactiveProductConfig(ctx, Status)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = product
	response.With200V2(w, "Success", m, platform)
}
