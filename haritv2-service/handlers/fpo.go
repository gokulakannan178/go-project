package handlers

import (
	"encoding/json"
	"log"

	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
	"strconv"
)

// SaveFPO : ""
func (h *Handler) SaveFPO(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	fpo := new(models.FPO)
	err := json.NewDecoder(r.Body).Decode(&fpo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveFPO(ctx, fpo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// SaveFPORegister : ""
func (h *Handler) SaveFPORegistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	fpo := new(models.FPO)
	err := json.NewDecoder(r.Body).Decode(&fpo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveFPORegistration(ctx, fpo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = fpo
	response.With200V2(w, "Success", m, platform)
}

// GetSingleFPO : ""
func (h *Handler) GetSingleFPO(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	fpo := new(models.RefFPO)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	fpo, err := h.Service.GetSingleFPO(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = fpo
	response.With200V2(w, "Success", m, platform)
}

// UpdateFPO : ""
func (h *Handler) UpdateFPO(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	fpo := new(models.FPO)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&fpo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if fpo.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateFPO(ctx, fpo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFPO : ""
func (h *Handler) EnableFPO(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFPO(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableFPO : ""
func (h *Handler) DisableFPO(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFPO(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteFPO : ""
func (h *Handler) DeleteFPO(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFPO(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterFPO : ""
func (h *Handler) FilterFPO(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.FPOFilter
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

	var FPOs []models.RefFPO
	log.Println(pagination)
	FPOs, err = h.Service.FilterFPO(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(FPOs) > 0 {
		m["fpo"] = FPOs
	} else {
		res := make([]models.FPO, 0)
		m["fpo"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UpdateFPO : ""
func (h *Handler) UpdateFPORegistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	fpo := new(models.FPO)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&fpo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if fpo.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateFPORegistration(ctx, fpo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FPOMasterReport : ""
func (h *Handler) FPOMasterReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.FPOReportFilter)
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

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FPOMasterReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.FPOMasterReportJSON(ctx, filter, pagination)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpomasterreport"] = data
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}
	response.With200V2(w, "Success", m, platform)

}

//FPOMasterReport : ""
func (h *Handler) FPOMonthReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.FPOMothWiseeportFilter)

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FPOMonthReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.FPOMonthReportJSON(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpomasterreport"] = data
	response.With200V2(w, "Success", m, platform)

}
func (h *Handler) FBONearBy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	fbonb := new(models.FBONearBy)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&fbonb)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	//km!=0
	if fbonb.KM == 0 {
		response.With400V2(w, "id is missing", platform)
	}
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
	ulbs, err := h.Service.FBONearBy(ctx, fbonb, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fboNearBy"] = ulbs
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
