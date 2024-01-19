package handlers

import (
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//GetDefaultPaymentGateway : ""
func (h *Handler) GetDefaultPaymentGateway(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// UniqueID := r.URL.Query().Get("id")

	// if UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// 	return
	// }

	paymentGateway := new(models.PaymentGateway)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	paymentGateway, err := h.Service.GetDefaultPaymentGateway(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["paymentGateway"] = paymentGateway
	response.With200V2(w, "Success", m, platform)
}
