package handlers

import (
	"bpms-service/app"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"net/http"
)

//PlanMakeFailScrutiny : ""
func (h *Handler) PlanMakeFailScrutiny(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	pfs := new(models.PlanFailScrutiny)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&pfs)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.PlanMakeFailScrutiny(ctx, pfs)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["pfs"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//PlanMakePassScrutiny : ""
func (h *Handler) PlanMakePassScrutiny(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	pps := new(models.PlanPassScrutiny)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&pps)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.PlanMakePassScrutiny(ctx, pps)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["pps"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ProceedPCP : ""
func (h *Handler) ProceedPCP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ppcp := new(models.ProceedPCP)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&ppcp)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ProceedPCP(ctx, ppcp)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ppcp"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//MakePCPDefective : ""
func (h *Handler) MakePCPDefective(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mpcpd := new(models.MakePCPDefective)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&mpcpd)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.MakePCPDefective(ctx, mpcpd)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mpcpd"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//PCPAccept : ""
func (h *Handler) PCPAccept(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.PCPAccept)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.PCPAccept(ctx, data)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["pcpa"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeptApprovalAccept : ""
func (h *Handler) DeptApprovalAccept(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.DeptApprovalAccept)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.DeptApprovalAccept(ctx, data)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["daa"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeptApprovalReject : ""
func (h *Handler) DeptApprovalReject(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.DeptApprovalReject)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.DeptApprovalReject(ctx, data)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//CCAccept : ""
func (h *Handler) CCAccept(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.CCAccept)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.CCAccept(ctx, data)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["cca"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//CCReject : ""
func (h *Handler) CCReject(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.CCReject)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.CCReject(ctx, data)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ccr"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//MakePayment : ""
func (h *Handler) MakePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.MakePayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.MakePayment(ctx, data)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mp"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ReapplyDefective : ""
func (h *Handler) ReapplyDefective(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.ReapplyDefective)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ReapplyDefective(ctx, data)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["res"] = "success"
	response.With200V2(w, "Success", m, platform)
}
