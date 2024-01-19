package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"log"
	"net/http"
	"strconv"
)

//SaveDumpHistory : ""
func (h *Handler) SaveDumpHistory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	dumpHistory := new(models.DumpHistory)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&dumpHistory)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveDumpHistory(ctx, dumpHistory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dumpHistory"] = dumpHistory
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDumpHistory :""
func (h *Handler) GetSingleDumpHistory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	dumpHistory := new(models.RefDumpHistory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	dumpHistory, err := h.Service.GetSingleDumpHistory(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dumpHistory"] = dumpHistory
	response.With200V2(w, "Success", m, platform)
}

//UpdateDumpHistory :""
func (h *Handler) UpdateDumpHistory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	dumpHistory := new(models.DumpHistory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&dumpHistory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if dumpHistory.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDumpHistory(ctx, dumpHistory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dumpHistory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDumpHistory : ""
func (h *Handler) EnableDumpHistory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDumpHistory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DumpHistory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDumpHistory : ""
func (h *Handler) DisableDumpHistory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDumpHistory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DumpHistory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDumpHistory : ""
func (h *Handler) DeleteDumpHistory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDumpHistory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DumpHistory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterDumpHistory : ""
func (h *Handler) FilterDumpHistory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var DumpHistory *models.FilterDumpHistory
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
	err := json.NewDecoder(r.Body).Decode(&DumpHistory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DumpHistorys []models.RefDumpHistory
	log.Println(pagination)
	DumpHistorys, err = h.Service.FilterDumpHistory(ctx, DumpHistory, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(DumpHistorys) > 0 {
		m["DumpHistory"] = DumpHistorys
	} else {
		res := make([]models.DumpHistory, 0)
		m["DumpHistory"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetQuantityByManagerId(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var DumpHistory *models.FilterDumpHistory
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&DumpHistory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DumpHistorys []models.GetQuantity
	DumpHistorys, err = h.Service.GetQuantityByManagerId(ctx, DumpHistory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(DumpHistorys) > 0 {
		m["DumpHistory"] = DumpHistorys
	} else {
		res := make([]models.DumpHistory, 0)
		m["DumpHistory"] = res
	}

	response.With200V2(w, "Success", m, platform)
}

// func (h *Handler) DumpHistoryAssign(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	DumpHistoryAssign := new(models.DumpHistoryAssign)
// 	ctx := app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())
// 	err := json.NewDecoder(r.Body).Decode(&DumpHistoryAssign)
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	defer r.Body.Close()

// 	err = h.Service.DumpHistoryAssign(ctx, DumpHistoryAssign)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["DumpHistory"] = DumpHistoryAssign
// 	response.With200V2(w, "Success", m, platform)
// }
// func (h *Handler) RevokeDumpHistory(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	DumpHistory := new(models.DumpHistory)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())

// 	err := json.NewDecoder(r.Body).Decode(&DumpHistory)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	if DumpHistory.UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}
// 	err = h.Service.RevokeDumpHistory(ctx, DumpHistory)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["DumpHistory"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }

func (h *Handler) DateWiseDumpHistory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var DumpHistory *models.DayWiseDumpHistory
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&DumpHistory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DumpHistorys []models.GetQuantity
	DumpHistorys, err = h.Service.DateWiseDumpHistory(ctx, DumpHistory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(DumpHistorys) > 0 {
		m["DumpHistory"] = DumpHistorys
	} else {
		res := make([]models.DumpHistory, 0)
		m["DumpHistory"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
