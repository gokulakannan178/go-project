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

// DashboardDayWiseTradelicenseCollectionChart : ""
func (h *Handler) DashboardDayWiseTradelicenseCollectionChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.DashboardDayWiseTradeLicenseCollectionChartFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	if resType == "pdf" {
		data, err := h.Service.DashboardDayWiseTradeLicenseCollectionChartPDF(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisetradelicensecollection.pdf")
		return
	}

	if resType == "excel" {
		file, err := h.Service.DashboardDayWiseTradeLicenseCollectionChartExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisetradelicensecollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.DashboardDayWiseTradelicenseCollectionChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["tradelicense"] = data
	response.With200V2(w, "Success", m, platform)

}

// DayWiseTradelicenseDemandChart : ""
func (h *Handler) DayWiseTradelicenseDemandChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.DayWiseTradeLicenseDemandChartFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	if resType == "pdf" {
		data, err := h.Service.DayWiseTradeLicenseDemandReportPDF(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisetradelicensedemand.pdf")
		return
	}

	if resType == "excel" {
		file, err := h.Service.DayWiseTradeLicenseDemandReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisetradelicensedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.DayWiseTradeLicenseDemandChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["tradelicense"] = data
	response.With200V2(w, "Success", m, platform)

}

// TradeLicenseOverallDemandReport : ""
func (h *Handler) TradeLicenseOverallDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.TradeLicenseFilter
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
	if resType == "pdf" {
		data, err := h.Service.TradeLicenseOverallDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=demandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.TradeLicenseOverallDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicenseoveralldemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	data, err := h.Service.TradeLicenseOverallDemandReportJSON(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["tradelicense"] = data
	m["pagination"] = pagination
	response.With200V2(w, "Success", m, platform)

}

// FilterWardDayWiseTradeLicenseCollectionReport : ""
func (h *Handler) FilterWardDayWiseTradeLicenseCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardDayWiseTradeLicenseCollectionReportFilter
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

	var tradeLicenses []models.WardDayWiseTradeLicenseCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardDayWiseTradeLicenseCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensecollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardDayWiseTradeLicenseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardwisecollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterWardDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.TradeLicense, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardMonthWiseTradeLicenseCollectionReport : ""
func (h *Handler) FilterWardMonthWiseTradeLicenseCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardMonthWiseTradeLicenseCollectionReportFilter
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

	var tradeLicenses []models.WardMonthWiseTradeLicenseCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardMonthWiseTradeLicenseCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensecollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardMonthWiseTradeLicenseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardwisemonthcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterWardMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.TradeLicense, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardDayWiseTradeLicenseDemandReport : ""
func (h *Handler) FilterWardDayWiseTradeLicenseDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardDayWiseTradeLicenseDemandReportFilter
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

	var tradeLicenses []models.WardDayWiseTradeLicenseDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardDayWiseTradeLicenseDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardDayWiseTradeLicenseDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterWardDayWiseTradeLicenseDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.TradeLicense, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardDayWiseTradeLicenseDemandReport : ""
func (h *Handler) FilterWardMonthWiseTradeLicenseDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardMonthWiseTradeLicenseDemandReportFilter
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

	var tradeLicenses []models.WardMonthWiseTradeLicenseDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardMonthWiseTradeLicenseDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardmonthwisedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardMonthWiseTradeLicenseDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardmonthwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterWardMonthWiseTradeLicenseDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.WardMonthWiseTradeLicenseDemandReport, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamDayWiseTradeLicenseCollectionReport : ""
func (h *Handler) FilterTeamDayWiseTradeLicenseCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamDayWiseTradeLicenseCollectionReportFilter
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

	var tradeLicenses []models.TeamDayWiseTradeLicenseCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamDayWiseTradeLicenseCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicenseteamwisedaycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamDayWiseTradeLicenseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicenseteamwisedaycollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterTeamDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.TeamDayWiseTradeLicenseCollectionReport, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamMonthWiseTradeLicenseCollectionReport : ""
func (h *Handler) FilterTeamMonthWiseTradeLicenseCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamMonthWiseTradeLicenseCollectionReportFilter
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

	var tradeLicenses []models.TeamMonthWiseTradeLicenseCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamMonthWiseTradeLicenseCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicenseteamwisedaycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamMonthWiseTradeLicenseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicenseteamwisemonthcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterTeamMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.TeamMonthWiseTradeLicenseCollectionReport, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardYearWiseTradeLicenseCollectionReport : ""
func (h *Handler) FilterWardYearWiseTradeLicenseCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardYearWiseTradeLicenseCollectionReportFilter
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

	var tradeLicenses []models.WardYearWiseTradeLicenseCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardYearWiseTradeLicenseCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensecollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardYearWiseTradeLicenseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardwiseyearcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterWardYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.WardYearWiseTradeLicenseCollectionReport, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardYearWiseTradeLicenseDemandReport : ""
func (h *Handler) FilterWardYearWiseTradeLicenseDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardYearWiseTradeLicenseDemandReportFilter
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

	var tradeLicenses []models.WardYearWiseTradeLicenseDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardYearWiseTradeLicenseDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardyearwisedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardYearWiseTradeLicenseDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensewardyearwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterWardYearWiseTradeLicenseDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.WardYearWiseTradeLicenseDemandReport, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamYearWiseTradeLicenseCollectionReport : ""
func (h *Handler) FilterTeamYearWiseTradeLicenseCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamYearWiseTradeLicenseCollectionReportFilter
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

	var tradeLicenses []models.TeamYearWiseTradeLicenseCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamYearWiseTradeLicenseCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicenseteamwiseyearcollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamYearWiseTradeLicenseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicenseteamwiseyearcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	tradeLicenses, err = h.Service.FilterTeamYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenses) > 0 {
		m["tradeLicense"] = tradeLicenses
	} else {
		res := make([]models.TeamYearWiseTradeLicenseCollectionReport, 0)
		m["tradeLicense"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
