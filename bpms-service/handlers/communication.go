package handlers

import (
	"bpms-service/response"
	"net/http"
)

//TestEmailTemplate : ""
func (h *Handler) TestEmailTemplate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	templateID := r.URL.Query().Get("id")

	if templateID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	d := make(map[string]interface{})
	d["username"] = "Solomon"
	d["password"] = "password"
	err := h.Service.SendEmailWithTemplate("test subject", []string{"solomon2261993@gmail.com"}, "emailtemplates/"+templateID, d)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["email"] = "success"
	response.With200V2(w, "Success", m, platform)
}
