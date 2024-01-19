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

//WardWiseShoprentReport : ""
func (h *Handler) WardWiseShoprentReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	//respType := r.URL.Query().Get("resType")
	filter := new(models.WardWiseShoprentReportFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.WardWiseShoprentReport(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["wardWiseShoprentReport"] = data
	response.With200V2(w, "Success", m, platform)
}

// ShopRentOverallDemandReport : ""
func (h *Handler) ShopRentOverallDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.ShopRentFilter
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
		data, err := h.Service.ShopRentOverallDemandReportPDF(ctx, filter, pagination)
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
		file, err := h.Service.ShopRentOverallDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentoveralldemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	data, err := h.Service.ShopRentOverallDemandReportJSON(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["shoprents"] = data
	m["pagination"] = pagination
	response.With200V2(w, "Success", m, platform)

}

// DashboardDayWiseShoprentCollectionChart : ""
func (h *Handler) DashboardDayWiseShoprentCollectionChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.DashboardDayWiseShoprentCollectionChartFilter
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
		data, err := h.Service.DashboardDayWiseShoprentCollectionChartPDF(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=daywiseshoprentcollection.pdf")
		return
	}

	if resType == "excel" {
		file, err := h.Service.DashboardDayWiseShoprentCollectionChartExcel(ctx, filter)
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
	data, err := h.Service.DashboardDayWiseShoprentCollectionChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["shoprent"] = data
	response.With200V2(w, "Success", m, platform)

}

// DayWiseShoprentDemandChart : ""
func (h *Handler) DayWiseShoprentDemandChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var filter *models.DayWiseShoprentDemandChartFilter
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
		data, err := h.Service.DayWiseShoprentDemandChartPDF(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=daywiseshoprentdemand.pdf")
		return
	}

	if resType == "excel" {
		file, err := h.Service.DayWiseShopRentDemandReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=daywiseshoprentdemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.DayWiseShoprentDemandChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["shoprent"] = data
	response.With200V2(w, "Success", m, platform)

}

// FilterWardDayWiseShopRentCollectionReport : ""
func (h *Handler) FilterWardDayWiseShopRentCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardDayWiseShopRentCollectionReportFilter
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

	var shopRents []models.WardDayWiseShopRentCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardDayWiseShopRentCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentcollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardDayWiseShopRentCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardwisecollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterWardDayWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.ShopRent, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardMonthWiseShopRentCollectionReport : ""
func (h *Handler) FilterWardMonthWiseShopRentCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardMonthWiseShopRentCollectionReportFilter
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

	var shopRents []models.WardMonthWiseShopRentCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardMonthWiseShopRentCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentcollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardMonthWiseShopRentCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardwisemonthcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterWardMonthWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.ShopRent, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardDayWiseShopRentDemandReport : ""
func (h *Handler) FilterWardDayWiseShopRentDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardDayWiseShopRentDemandReportFilter
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

	var shopRents []models.WardDayWiseShopRentDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardDayWiseShopRentDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentdemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardDayWiseShopRentDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterWardDayWiseShopRentDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.ShopRent, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardDayWiseShopRentDemandReport : ""
func (h *Handler) FilterWardMonthWiseShopRentDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardMonthWiseShopRentDemandReportFilter
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

	var shopRents []models.WardMonthWiseShopRentDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardMonthWiseShopRentDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardmonthwisedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardMonthWiseShopRentDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardmonthwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterWardMonthWiseShopRentDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.ShopRent, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamDayWiseShopRentCollectionReport : ""
func (h *Handler) FilterTeamDayWiseShopRentCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamDayWiseShopRentCollectionReportFilter
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

	var shopRents []models.TeamDayWiseShopRentCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamDayWiseShopRentCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentteamwisedaycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamDayWiseShopRentCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentteamwisedaycollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterTeamDayWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.TeamDayWiseShopRentCollectionReport, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamMonthWiseShopRentCollectionReport : ""
func (h *Handler) FilterTeamMonthWiseShopRentCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamMonthWiseShopRentCollectionReportFilter
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

	var shopRents []models.TeamMonthWiseShopRentCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamMonthWiseShopRentCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentteamwisedaycollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamMonthWiseShopRentCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentteamwisemonthcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterTeamMonthWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.TeamMonthWiseShopRentCollectionReport, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardYearWiseShopRentCollectionReport : ""
func (h *Handler) FilterWardYearWiseShopRentCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardYearWiseShopRentCollectionReportFilter
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

	var shopRents []models.WardYearWiseShopRentCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardYearWiseShopRentCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentcollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardYearWiseShopRentCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardwiseyearcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterWardYearWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.WardYearWiseShopRentCollectionReport, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterWardYearWiseShopRentDemandReport : ""
func (h *Handler) FilterWardYearWiseShopRentDemandReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.WardYearWiseShopRentDemandReportFilter
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

	var shopRents []models.WardYearWiseShopRentDemandReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterWardYearWiseShopRentDemandReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardyearwisedemandreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterWardYearWiseShopRentDemandReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentwardyearwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterWardYearWiseShopRentDemandReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.WardYearWiseShopRentDemandReport, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterTeamYearWiseShopRentCollectionReport : ""
func (h *Handler) FilterTeamYearWiseShopRentCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.TeamYearWiseShopRentCollectionReportFilter
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

	var shopRents []models.TeamYearWiseShopRentCollectionReport
	log.Println(pagination)
	if resType == "pdf" {
		data, err := h.Service.FilterTeamYearWiseShopRentCollectionReportPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentteamwiseyearcollectionreceipt.pdf")
		return
	}
	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterTeamYearWiseShopRentCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=shoprentteamwiseyearcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	shopRents, err = h.Service.FilterTeamYearWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shopRent"] = shopRents
	} else {
		res := make([]models.TeamYearWiseShopRentCollectionReport, 0)
		m["shopRent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
