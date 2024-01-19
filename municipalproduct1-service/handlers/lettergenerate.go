package handlers

import (
	"encoding/json"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

// SaveLetterUpload : ""
func (h *Handler) SaveLetterGenerate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	letterGenerate := new(models.LetterGenerate)
	err := json.NewDecoder(r.Body).Decode(&letterGenerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveLetterGenerate(ctx, letterGenerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleLetterGenerate : ""
func (h *Handler) GetSingleLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	letterGenerate := new(models.RefLetterGenerate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	letterGenerate, err := h.Service.GetSingleLetterGenerate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = letterGenerate
	response.With200V2(w, "Success", m, platform)
}

// UpdateLetterUpload : ""
func (h *Handler) UpdateLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	letterGenerate := new(models.LetterGenerate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&letterGenerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if letterGenerate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateLetterGenerate(ctx, letterGenerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// //EnableLetterUpload : ""
// func (h *Handler) EnableLetterGenerate(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	UniqueID := r.URL.Query().Get("id")

// 	if UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}

// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())

// 	err := h.Service.EnableLetterGenerate(ctx, UniqueID)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["letterGenerate"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }

// EnableLetterGenerate : ""
func (h *Handler) EnableLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	letterGenerate := new(models.LetterGenerateAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&letterGenerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if letterGenerate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EnableLetterGenerate(ctx, letterGenerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) ApprovedLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	letterGenerate := new(models.LetterGenerateAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&letterGenerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if letterGenerate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.ApprovedLetterGenerate(ctx, letterGenerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableLetterUpload : ""
func (h *Handler) DisableLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableLetterGenerate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteLetterGenerate : ""
func (h *Handler) DeleteLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteLetterGenerate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// BlockedLetterGenerate : ""
func (h *Handler) BlockedLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	letterGenerate := new(models.LetterGenerateAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&letterGenerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if letterGenerate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.BlockedLetterGenerate(ctx, letterGenerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// SubmittedLetterGenerate : ""
func (h *Handler) SubmittedLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	letterGenerate := new(models.LetterGenerateAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&letterGenerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if letterGenerate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.SubmittedLetterGenerate(ctx, letterGenerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UploadLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	letterGenerate := new(models.LetterGenerate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&letterGenerate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if letterGenerate.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UploadLetterGenerate(ctx, letterGenerate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["letterGenerate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterLetterUpload : ""
func (h *Handler) FilterLetterGenerate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.LetterGenerateFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var lettergenerates []models.RefLetterGenerate
	log.Println(pagination)
	lettergenerates, err = h.Service.FilterLetterGenerate(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(lettergenerates) > 0 {
		m["lettergenerates"] = lettergenerates
	} else {
		res := make([]models.LetterGenerate, 0)
		m["letterGenerate"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//LetterGenerateExecute : ""
func (h *Handler) LetterGenerateExecute(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	r.Header.Set("Connection", "close")
	// 	r.Close = true
	// }()
	r.Body.Close()
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	data, lg, err := h.Service.LetterGenerateExecute(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	w.Write(data)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename="+lg.NO+".pdf")

}
