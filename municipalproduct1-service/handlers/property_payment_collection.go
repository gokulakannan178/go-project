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

//DashboardTotalCollectionChart : ""
func (h *Handler) DashboardTotalCollectionChart(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.DashboardTotalCollectionChartFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	if resType == "excel" {
		file, err := h.Service.DayWisePropertyCollectionReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisepropertycollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	if resType == "pdf" {
		data, err := h.Service.DayWisePropertyCollectionReportPDF(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisepropertycollection.pdf")
	}

	data, err := h.Service.DashboardTotalCollectionChart(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collection"] = data
	response.With200V2(w, "Success", m, platform)
}

//DashboardTotalCollectionOverview : ""
func (h *Handler) DashboardTotalCollectionOverview(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	filter := new(models.DashboardTotalCollectionOverviewFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.DashboardTotalCollectionOverview(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collection"] = data
	response.With200V2(w, "Success", m, platform)
}

//DashboardDayWiseCollectionChart : ""
func (h *Handler) DashboardDayWiseCollectionChart(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.DashboardDayWiseCollectionChartFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	if resType == "excel" {
		file, err := h.Service.DashboardDayWiseCollectionChartExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisepropertycollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	if resType == "daywisecollection" {
		file, err := h.Service.DayWiseCollectionChartExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisecollectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	if resType == "pdf" {
		data, err := h.Service.DashboardDayWiseCollectionChartPDF(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=daywisepropertycollection.pdf")
	}
	data, err := h.Service.DashboardDayWiseCollectionChart(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collection"] = data
	response.With200V2(w, "Success", m, platform)
}

//WardWiseCollectionReport : ""
func (h *Handler) WardWiseCollectionReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	respType := r.URL.Query().Get("resType")
	filter := new(models.WardWiseCollectionReportFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
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

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	if respType == "excel" {
		file, err := h.Service.WardWiseCollectionReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=wardwisecollection.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	log.Println(pagination)
	data, err := h.Service.WardWiseCollectionReport(ctx, filter, pagination)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collection"] = data
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}
	response.With200V2(w, "Success", m, platform)
}

//TCCollectionSummaryReport : ""
func (h *Handler) TCCollectionSummaryReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	respType := r.URL.Query().Get("resType")
	filter := new(models.TCCollectionSummaryFilter)

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
	if respType == "excel" {
		file, err := h.Service.TCCollectionSummaryReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=teamwisecollection.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	// // header for pdf file

	/*// if respType == "pdf" {
	// 	file, err := h.Service.TCCollectionSummaryReportPdf(ctx, filter)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/octet-stream")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	// 	w.Header().Set("Content-Transfer-Encoding", "binary")
	// 	file.Write(w)
	// 	return
	// }*/
	data, err := h.Service.TCCollectionSummaryReport(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	data1, err := h.Service.TCCollectionSummaryReportV2(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collection"] = data
	m["totalAmount"] = data1.TotalAmount
	m["totalConsumer"] = data1.TotalConsumer
	response.With200V2(w, "Success", m, platform)

}

// PropertyIDWiseReport : ""
func (h *Handler) PropertyIDWiseReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var property *models.PropertyFilter
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
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.PropertyIDWiseReport(ctx, property, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyidwisereport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var propertys []models.RefProperty
	log.Println(pagination)
	propertys, err = h.Service.FilterProperty(ctx, property, pagination)
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

// ZoneAndWardWiseCollection : ""
func (h *Handler) ZoneAndWardWiseCollection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var zwFilter *models.ZoneAndWardWiseReportFilter
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
	err := json.NewDecoder(r.Body).Decode(&zwFilter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.ZoneAndWardWiseReportExcel(ctx, zwFilter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=zoneandwardwisereport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var report []models.ZoneAndWardWiseReport
	log.Println(pagination)
	report, err = h.Service.ZoneAndWardWiseReport(ctx, zwFilter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(report) > 0 {
		m["report"] = report
	} else {
		res := make([]models.ZoneAndWardWiseReport, 0)
		m["report"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyMonthWiseCollectionReport : ""
func (h *Handler) FilterPropertyMonthWiseCollectionReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// resType := r.URL.Query().Get("resType")
	var filter *models.PropertyMonthWiseCollectionReportFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// pageNo := r.URL.Query().Get("pageno")
	// Limit := r.URL.Query().Get("limit")

	// var pagination *models.Pagination
	// if pageNo != "no" {
	// 	pagination = new(models.Pagination)
	// 	if pagination.PageNum = 1; pageNo != "" {
	// 		page, err := strconv.Atoi(pageNo)
	// 		if pagination.PageNum = 1; err == nil {
	// 			pagination.PageNum = page
	// 		}
	// 	}
	// 	if pagination.Limit = 10; Limit != "" {
	// 		limit, err := strconv.Atoi(Limit)
	// 		if pagination.Limit = 10; err == nil {
	// 			pagination.Limit = limit
	// 		}
	// 	}
	// }
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var shopRents []models.PropertyMonthWiseCollectionReport
	// log.Println(pagination)
	// if resType == "pdf" {
	// 	data, err := h.Service.FilterWardYearWiseShopRentCollectionReportPDF(ctx, filter, pagination)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	w.Write(data)
	// 	w.Header().Set("Content-Type", "application/pdf")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=shoprentcollectionreceipt.pdf")
	// 	return
	// }
	// header for excel file
	// if resType == "excel" {
	// 	file, err := h.Service.FilterPropertyMonthWiseCollectionReportExcel(ctx, filter)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/octet-stream")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=propertymonthwisecollectionreport.xlsx")
	// 	w.Header().Set("ocntent-Transfer-Encoding", "binary")
	// 	file.Write(w)
	// 	return
	// }

	shopRents, err = h.Service.FilterPropertyMonthWiseCollectionReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["propertyReport"] = shopRents
	} else {
		res := make([]models.PropertyMonthWiseCollectionReport, 0)
		m["propertyReport"] = res
	}
	// if pagination != nil {
	// 	if pagination.PageNum > 0 {
	// 		m["pagination"] = pagination
	// 	}
	// }

	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyWiseCollectionReport : ""
func (h *Handler) FilterPropertyWiseCollectionReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.PropertyWiseCollectionReportFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	if resType == "excel" {
		file, err := h.Service.FilterPropertyWiseCollectionReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywisecollectionreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	var res []models.PropertyWiseCollectionReport

	res, err = h.Service.FilterPropertyWiseCollectionReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(res) > 0 {
		m["collection"] = res
	} else {
		res := make([]models.PropertyWiseCollectionReport, 0)
		m["collection"] = res
	}
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyWiseDemandReportExcel : ""
func (h *Handler) FilterPropertyWiseDemandReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.PropertyFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	if resType == "excel" {
		file, err := h.Service.FilterPropertyWiseDemandReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywisedemandreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	// var res []models.PropertyWiseCollectionReport

	// res, err = h.Service.PropertyOverallDemandReportJSON(ctx, filter)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["demandReport"] = "success"

	// if len(res) > 0 {
	// 	m["collection"] = res
	// } else {
	// 	res := make([]models.PropertyWiseCollectionReport, 0)
	// 	m["collection"] = res
	// }
	response.With200V2(w, "Success", m, platform)
}
