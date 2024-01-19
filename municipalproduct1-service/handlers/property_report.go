package handlers

import (
	"encoding/json"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
	"time"
)

// PropertyOverallDemandReport : ""
func (h *Handler) PropertyOverallDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.PropertyFilter
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
		data, err := h.Service.PropertyOverallDemandReportPDF(ctx, filter, pagination)
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
		file, err := h.Service.PropertyOverallDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyoveralldemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	// header for excel file
	if resType == "excelv2" {
		file, err := h.Service.PropertyOverallDemandReportExcelV3(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyoveralldemandv2.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.PropertyOverallDemandReportJSON(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["property"] = data
	m["pagination"] = pagination
	response.With200V2(w, "Success", m, platform)

}

// DayWisePropertyCollectionChart : ""
func (h *Handler) DayWisePropertyDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.DayWisePropertyDemandChartFilter
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
		data, err := h.Service.DayWisePropertyDemandReportPDF(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisepropertydemand.pdf")
		return
	}

	if resType == "excel" {
		file, err := h.Service.DayWisePropertyDemandReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisepropertydemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.DayWisePropertyDemandReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["property"] = data
	response.With200V2(w, "Success", m, platform)

}

// FilterWardDayWisePropertyCollectionReport : ""
func (h *Handler) FilterWardDayWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardDayWisePropertyCollectionReportFilter
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

	var properties []models.WardDayWisePropertyCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardDayWisePropertyCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertycollectionreceipt.pdf")
		return
	}
	// // header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardDayWisePropertyCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardwisecollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterWardDayWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["properties"] = properties
	} else {
		res := make([]models.Property, 0)
		m["properties"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardMonthWisePropertyCollectionReport : ""
func (h *Handler) FilterWardMonthWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardMonthWisePropertyCollectionReportFilter
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

	var properties []models.WardMonthWisePropertyCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardMonthWisePropertyCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardMonthWisePropertyCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardwisemonthcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterWardMonthWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["properties"] = properties
	} else {
		res := make([]models.Property, 0)
		m["properties"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardDayWisePropertyDemandReport : ""
func (h *Handler) FilterWardDayWisePropertyDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardDayWisePropertyDemandReportFilter
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

	var propertys []models.WardDayWisePropertyDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardDayWisePropertyDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywarddaywisedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardDayWisePropertyDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywarddaywisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	propertys, err = h.Service.FilterWardDayWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertys) > 0 {
		m["property"] = propertys
	} else {
		res := make([]models.Property, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardMonthWisePropertyDemandReport : ""
func (h *Handler) FilterWardMonthWisePropertyDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardMonthWisePropertyDemandReportFilter
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

	var propertys []models.WardMonthWisePropertyDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardMonthWisePropertyDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardmonthwisedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardMonthWisePropertyDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardmonthwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	propertys, err = h.Service.FilterWardMonthWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertys) > 0 {
		m["property"] = propertys
	} else {
		res := make([]models.Property, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamDayWisePropertyCollectionReport : ""
func (h *Handler) FilterTeamDayWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamDayWisePropertyCollectionReportFilter
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

	var properties []models.TeamDayWisePropertyCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamDayWisePropertyCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyteamwisedaycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamDayWisePropertyCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyteamwisedaycollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterTeamDayWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["property"] = properties
	} else {
		res := make([]models.TeamDayWisePropertyCollectionReport, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamMonthWisePropertyCollectionReport : ""
func (h *Handler) FilterTeamMonthWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamMonthWisePropertyCollectionReportFilter
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

	var properties []models.TeamMonthWisePropertyCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamMonthWisePropertyCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyteamwisedaycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamMonthWisePropertyCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyteamwisemonthcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterTeamMonthWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["property"] = properties
	} else {
		res := make([]models.TeamMonthWisePropertyCollectionReport, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardYearWisePropertyCollectionReport : ""
func (h *Handler) FilterWardYearWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardYearWisePropertyCollectionReportFilter
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

	var properties []models.WardYearWisePropertyCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardYearWisePropertyCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardYearWisePropertyCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardwiseyearcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterWardYearWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["property"] = properties
	} else {
		res := make([]models.WardYearWisePropertyCollectionReport, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardYearWisePropertyDemandReport : ""
func (h *Handler) FilterWardYearWisePropertyDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardYearWisePropertyDemandReportFilter
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

	var properties []models.WardYearWisePropertyDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardYearWisePropertyDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardyearwisedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardYearWisePropertyDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardyearwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterWardYearWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["property"] = properties
	} else {
		res := make([]models.WardYearWisePropertyDemandReport, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamYearWisePropertyCollectionReport : ""
func (h *Handler) FilterTeamYearWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamYearWisePropertyCollectionReportFilter
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

	var properties []models.TeamYearWisePropertyCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamYearWisePropertyCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyteamwiseyearcollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamYearWisePropertyCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyteamwiseyearcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterTeamYearWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["property"] = properties
	} else {
		res := make([]models.TeamYearWisePropertyCollectionReport, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterYearWisePropertyDemandReport : ""
func (h *Handler) FilterYearWisePropertyDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.YearWisePropertyDemandReportFilter
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

	var properties []models.YearWisePropertyDemandReport
	log.Println(pagination)
	// header for pdf file
	// if resType == "pdf" {
	// 	data, err := h.Service.FilterYearWisePropertyDemandReportPDF(ctx, filter, pagination)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	w.Write(data)
	// 	w.Header().Set("Content-Type", "application/pdf")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=propertycollectionreceipt.pdf")
	// 	return
	// }
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterYearWisePropertyDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyyearwisedemandreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterYearWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["property"] = properties
	} else {
		res := make([]models.YearWisePropertyDemandReport, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterYearWisePropertyCollectionReport : ""
func (h *Handler) FilterYearWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.YearWisePropertyCollectionReportFilter
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

	var properties []models.YearWisePropertyCollectionReport
	log.Println(pagination)
	// if resType == "pdf" {
	// 	data, err := h.Service.FilterYearWisePropertyDemandReportPDF(ctx, filter, pagination)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	w.Write(data)
	// 	w.Header().Set("Content-Type", "application/pdf")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=propertycollectionreceipt.pdf")
	// 	return
	// }
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterYearWisePropertyCollectionReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyyearwisecollectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterYearWisePropertyCollectionReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(properties) > 0 {
		m["property"] = properties
	} else {
		res := make([]models.YearWisePropertyCollectionReport, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterYearWisePropertyCollectionReport : ""
func (h *Handler) FilterWardWisePropertyDemandAndCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.YearWisePropertyCollectionReportFilter
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
	var properties models.WardWisePropertyDemandAndCollectionReport
	log.Println(pagination)
	if resType == "excel" {
		file, err := h.Service.FilterWardWisePropertyDemandAndCollectionReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywardwisedemadandcollectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	properties, err = h.Service.FilterWardWisePropertyDemandAndCollectionReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collection"] = properties
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UserWisePropertyCollectionReport : ""
func (h *Handler) UserWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.UserWisePropertyCollectionReportFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	if resType == "excel" {
		file, err := h.Service.UserWisePropertyCollectionReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=userwisepropertycollectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.UserWisePropertyCollectionReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["report"] = data
	response.With200V2(w, "Success", m, platform)

}

// FilterPropertyArrearAndCurrentCollectionReport : ""
func (h *Handler) FilterPropertyArrearAndCurrentCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyArrearAndCurrentCollectionFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var res []models.PropertyArrearAndCurrentCollectionReport
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterPropertyArrearAndCurrentCollectionReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyarrearandcollectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	res, err = h.Service.FilterPropertyArrearAndCurrentCollectionReportJSON(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(res) > 0 {
		m["report"] = res
	} else {
		res := make([]models.PropertyArrearAndCurrentCollectionReport, 0)
		m["report"] = res
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterCounterReportV2 : ""
func (h *Handler) FilterCounterReportV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	response.With500mV2(w, "failed - ", platform)
	return
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

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterCounterReportV2Excel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=counterreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var payments []models.RefCounterReportV2
	log.Println(pagination)
	payments, err = h.Service.FilterCounterReportV2JSON(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["counterReportV2"] = payments
	} else {
		res := make([]models.RefCounterReportV2, 0)
		m["counterReportV2"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterUserWisePropertyCollectionReport : ""
func (h *Handler) FilterUserWisePropertyCollectionReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var filter *models.UserWisePropertyCollectionFilter
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

	// payment := new(models.RefUserWisePropertyCollection)
	// log.Println(pagination)
	// payments, err = h.Service.FilterUserWisePropertyCollectionReport(ctx, filter, pagination)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	// m := make(map[string]interface{})

	// if len(payments) > 0 {
	// 	m["collections"] = payments
	// } else {
	// 	res := make([]models.RefUserWisePropertyCollection, 0)
	// 	m["collections"] = res
	// }
	// if pagination != nil {
	// 	if pagination.PageNum > 0 {
	// 		m["pagination"] = pagination
	// 	}
	// }

	// response.With200V2(w, "Success", m, platform)
	data, err := h.Service.FilterUserWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["resData"] = data
	response.With200V2(w, "Success", m, platform)
}

// FilterHoldingWiseCollectionReportJSON : ""
func (h *Handler) FilterHoldingWiseCollectionReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	response.With500mV2(w, "failed - ", platform)
	return
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

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterHoldingWiseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=collectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var payments []models.RefHoldingWiseCollectionReport
	log.Println(pagination)
	payments, err = h.Service.FilterHoldingWiseCollectionReportJSON(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["holdingWiseCollection"] = payments
	} else {
		res := make([]models.RefHoldingWiseCollectionReport, 0)
		m["holdingWiseCollection"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) FilterPaymentCOllection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	if ctx.ProductConfig.LocationID == "Munger" {
		time.Sleep(20000)
		response.With500mV2(w, "timeout - ", platform)
		return
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.CounterReportV2(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=counterreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var payments []models.RefPropertyCollectionReport
	log.Println(pagination)
	payments, err = h.Service.FilterPaymentCOllection(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["propertyfycollection"] = payments
	} else {
		res := make([]models.RefPropertyCollectionReport, 0)
		m["propertyfycollection"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) FilterPaymentCOllectionSpare(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
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

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.CounterReportV2(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=counterreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var payments []models.RefPropertyCollectionReport
	log.Println(pagination)
	payments, err = h.Service.FilterPaymentCOllection(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["propertyfycollection"] = payments
	} else {
		res := make([]models.RefPropertyCollectionReport, 0)
		m["propertyfycollection"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyDemandAndCollectionReport : ""
func (h *Handler) FilterPropertyDemandAndCollectionReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyFilter
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

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.PropertyDemandAndCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertydemandandcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var reports []models.RefPropertyDemandAndCollectionReport
	log.Println(pagination)
	reports, err = h.Service.PropertyDemandAndCollectionReportJSON(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(reports) > 0 {
		m["dcReport"] = reports
	} else {
		res := make([]models.RefPropertyDemandAndCollectionReport, 0)
		m["dcReport"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyDemandAndCollectionReport : ""
func (h *Handler) FilterHoldingWiseCollectionReportJSONV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
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

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterHoldingWiseCollectionReportExcelV2(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=collectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var payments []models.RefHoldingWiseCollectionReportV2
	log.Println(pagination)
	payments, err = h.Service.FilterHoldingWiseCollectionReportJSONV2(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["holdingWiseCollection"] = payments
	} else {
		res := make([]models.RefHoldingWiseCollectionReport, 0)
		m["holdingWiseCollection"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) PropertyWiseDemandandCollectionExcelV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	// setting pagination
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

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.PropertyWiseDemandandCollectionExcelV3(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywisedemandandcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var wardwiseDemand []models.RefPropertyV2
	log.Println(pagination)
	wardwiseDemand, err = h.Service.PropertyWiseDemandandCollectionExcelV2(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(wardwiseDemand) > 0 {
		m["propertyDC"] = wardwiseDemand
	} else {
		res := make([]models.ResPropertyWiseDemandandCollectionV2Report, 0)
		m["propertyDC"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)

}

func (h *Handler) PropertyWiseDemandCollectionandBalanceReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	// setting pagination
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

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.PropertyWiseDemandCollectionandBalanceReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywisedemandcollectionandBalance.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var wardwiseDemand []models.RefProperty
	log.Println(pagination)
	wardwiseDemand, err = h.Service.PropertyWiseDemandCollectionandBalanceReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(wardwiseDemand) > 0 {
		m["propertyDC"] = wardwiseDemand
	} else {
		res := make([]models.ResPropertyWiseDemandandCollectionV2Report, 0)
		m["propertyDC"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)

}
