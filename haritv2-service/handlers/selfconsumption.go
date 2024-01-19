package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"log"
	"net/http"
	"strconv"
)

//SaveSelfConsumption : ""
func (h *Handler) SaveSelfConsumption(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	selfconsumption := new(models.SelfConsumption)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&selfconsumption)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveSelfConsumption(ctx, selfconsumption)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["selfconsumption"] = selfconsumption
	response.With200V2(w, "Success", m, platform)
}

//UpdateSelfConsumption :""
func (h *Handler) UpdateSelfConsumption(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	selfconsumption := new(models.SelfConsumption)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&selfconsumption)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if selfconsumption.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateSelfConsumption(ctx, selfconsumption)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["selfconsumption"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableSelfConsumption : ""
func (h *Handler) EnableSelfConsumption(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableSelfConsumption(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["selfconsumption"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableSelfConsumption : ""
func (h *Handler) DisableSelfConsumption(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableSelfConsumption(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["selfconsumption"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteSelfConsumption : ""
func (h *Handler) DeleteSelfConsumption(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteSelfConsumption(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["selfconsumption"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleSelfConsumption :""
func (h *Handler) GetSingleSelfConsumption(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	selfconsumption := new(models.RefSelfConsumption)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	selfconsumption, err := h.Service.GetSingleSelfConsumption(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["selfconsumption"] = selfconsumption
	response.With200V2(w, "Success", m, platform)
}

//FilterSelfConsumption : ""
func (h *Handler) FilterSelfConsumption(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var selfconsumption *models.SelfConsumptionFilter
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
	err := json.NewDecoder(r.Body).Decode(&selfconsumption)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var SelfConsumptions []models.RefSelfConsumption
	log.Println(pagination)
	if resType == "excel" {
		file, err := h.Service.SelfConsumptionExcel(ctx, selfconsumption, nil)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=SelfConsumption.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	SelfConsumptions, err = h.Service.FilterSelfConsumption(ctx, selfconsumption, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(SelfConsumptions) > 0 {
		m["selfconsumption"] = SelfConsumptions
	} else {
		res := make([]models.SelfConsumption, 0)
		m["selfconsumption"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DecreaseInventoryForULBandFPO(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	selfconsumption := new(models.SelfConsumption)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&selfconsumption)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.DecreaseInventoryForULBandFPO(ctx, selfconsumption)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["selfconsumption"] = "success"
	response.With200V2(w, "Success", m, platform)
}
