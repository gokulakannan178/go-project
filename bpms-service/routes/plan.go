package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PlanRoutes : ""
func (route *Route) PlanRoutes(r *mux.Router) {
	r.Handle("/plan", Adapt(http.HandlerFunc(route.Handler.SavePlan))).Methods("POST")
	r.Handle("/plan", Adapt(http.HandlerFunc(route.Handler.GetSinglePlan))).Methods("GET")
	r.Handle("/plan", Adapt(http.HandlerFunc(route.Handler.UpdatePlan))).Methods("PUT")
	r.Handle("/plan/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePlan))).Methods("PUT")
	r.Handle("/plan/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePlan))).Methods("PUT")
	r.Handle("/plan/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePlan))).Methods("DELETE")
	r.Handle("/plan/filter", Adapt(http.HandlerFunc(route.Handler.FilterPlan))).Methods("POST")
}

//PlanRegistrationTypeRoutes : ""
func (route *Route) PlanRegistrationTypeRoutes(r *mux.Router) {
	r.Handle("/planregistrationtype", Adapt(http.HandlerFunc(route.Handler.SavePlanRegistrationType))).Methods("POST")
	r.Handle("/planregistrationtype", Adapt(http.HandlerFunc(route.Handler.GetSinglePlanRegistrationType))).Methods("GET")
	r.Handle("/planregistrationtype", Adapt(http.HandlerFunc(route.Handler.UpdatePlanRegistrationType))).Methods("PUT")
	r.Handle("/planregistrationtype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePlanRegistrationType))).Methods("PUT")
	r.Handle("/planregistrationtype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePlanRegistrationType))).Methods("PUT")
	r.Handle("/planregistrationtype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePlanRegistrationType))).Methods("DELETE")
	r.Handle("/planregistrationtype/filter", Adapt(http.HandlerFunc(route.Handler.FilterPlanRegistrationType))).Methods("POST")
}

//PlanDepartmentApprovalRoutes : ""
func (route *Route) PlanDepartmentApprovalRoutes(r *mux.Router) {
	r.Handle("/plandepartmentapproval", Adapt(http.HandlerFunc(route.Handler.SavePlanDepartmentApproval))).Methods("POST")
	r.Handle("/plandepartmentapprovals", Adapt(http.HandlerFunc(route.Handler.SaveMultiplePlanDepartmentApproval))).Methods("POST")
	// r.Handle("/plandepartmentapprovals/v2", Adapt(http.HandlerFunc(route.Handler.GetAPlanDeptsApprov?alV2))).Methods("POST")
	r.Handle("/plandepartmentapproval", Adapt(http.HandlerFunc(route.Handler.GetSinglePlanDepartmentApproval))).Methods("GET")
	r.Handle("/plandepartmentapproval", Adapt(http.HandlerFunc(route.Handler.UpdatePlanDepartmentApproval))).Methods("PUT")
	r.Handle("/plandepartmentapproval/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePlanDepartmentApproval))).Methods("PUT")
	r.Handle("/plandepartmentapproval/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePlanDepartmentApproval))).Methods("PUT")
	r.Handle("/plandepartmentapproval/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePlanDepartmentApproval))).Methods("DELETE")
	r.Handle("/plandepartmentapproval/filter", Adapt(http.HandlerFunc(route.Handler.FilterPlanDepartmentApproval))).Methods("POST")
	r.Handle("/plandepartmentapproval/singledept", Adapt(http.HandlerFunc(route.Handler.GetAPlanDeptsApproval))).Methods("GET")
	r.Handle("/plandepartmentapproval/singledept/v2", Adapt(http.HandlerFunc(route.Handler.GetAPlanDeptsApprovalV2))).Methods("GET")
	r.Handle("/plandepartmentapproval/singledept/v3", Adapt(http.HandlerFunc(route.Handler.GetAPlanDeptsApprovalV3))).Methods("GET")
}

//PlanDepartmentFlowRoutes : ""
func (route *Route) PlanDepartmentFlowRoutes(r *mux.Router) {
	r.Handle("/plan/makefailscrutiny", Adapt(http.HandlerFunc(route.Handler.PlanMakeFailScrutiny))).Methods("POST")
	r.Handle("/plan/makepassscrutiny", Adapt(http.HandlerFunc(route.Handler.PlanMakePassScrutiny))).Methods("POST")
	r.Handle("/plan/proceedpcp", Adapt(http.HandlerFunc(route.Handler.ProceedPCP))).Methods("POST")
	r.Handle("/plan/makepcpdefective", Adapt(http.HandlerFunc(route.Handler.MakePCPDefective))).Methods("POST")
	r.Handle("/plan/pcpaccept", Adapt(http.HandlerFunc(route.Handler.PCPAccept))).Methods("POST")
	r.Handle("/plan/deptapprovalaccept", Adapt(http.HandlerFunc(route.Handler.DeptApprovalAccept))).Methods("POST")
	r.Handle("/plan/deptapprovalreject", Adapt(http.HandlerFunc(route.Handler.DeptApprovalReject))).Methods("POST")
	r.Handle("/plan/ccaccept", Adapt(http.HandlerFunc(route.Handler.CCAccept))).Methods("POST")
	r.Handle("/plan/ccreject", Adapt(http.HandlerFunc(route.Handler.CCReject))).Methods("POST")
	r.Handle("/plan/makepayment", Adapt(http.HandlerFunc(route.Handler.MakePayment))).Methods("POST")
	r.Handle("/plan/reapplydefective", Adapt(http.HandlerFunc(route.Handler.ReapplyDefective))).Methods("POST")
}

//PlanCRFRoutes : ""
func (route *Route) PlanCRFRoutes(r *mux.Router) {
	r.Handle("/plan/crf/filter", Adapt(http.HandlerFunc(route.Handler.FilterCRF))).Methods("POST")
	r.Handle("/plan/crf", Adapt(http.HandlerFunc(route.Handler.GetSingleCRF))).Methods("GET")
	r.Handle("/plan/crf/startinspection", Adapt(http.HandlerFunc(route.Handler.StartPlanCRFInspection))).Methods("POST")
	r.Handle("/plan/crf/getinspection", Adapt(http.HandlerFunc(route.Handler.GetCRFInspectionOfPlan))).Methods("GET")
	r.Handle("/plan/crf/submitinspection", Adapt(http.HandlerFunc(route.Handler.SubmitCRFInspection))).Methods("POST")
	r.Handle("/plan/crf/endinspection", Adapt(http.HandlerFunc(route.Handler.EndPlanCRFInspection))).Methods("POST")

	r.Handle("/plan/crf/crfaccept", Adapt(http.HandlerFunc(route.Handler.PlanCRFAccept))).Methods("POST")
	r.Handle("/plan/crf/postinspectionaccept", Adapt(http.HandlerFunc(route.Handler.PlanCRFPostInspectionAccept))).Methods("POST")
	r.Handle("/plan/crf/postinspectionreject", Adapt(http.HandlerFunc(route.Handler.PlanCRFPostInspectionReject))).Methods("POST")
	r.Handle("/plan/crf/certificatecomplete", Adapt(http.HandlerFunc(route.Handler.PlanCRFCertificateComplete))).Methods("POST")
	r.Handle("/plan/crf/crfreapply", Adapt(http.HandlerFunc(route.Handler.PlanCRFReapply))).Methods("POST")
	r.Handle("/plan/crf/crfreject", Adapt(http.HandlerFunc(route.Handler.PlanCRFReject))).Methods("POST")

}

//PlanReqDocumentRoutes : ""
func (route *Route) PlanReqDocumentRoutes(r *mux.Router) {
	r.Handle("/plan/reqdocument", Adapt(http.HandlerFunc(route.Handler.SavePlanReqDocument))).Methods("POST")
	r.Handle("/plan/reqdocument", Adapt(http.HandlerFunc(route.Handler.GetSinglePlanReqDocument))).Methods("GET")
	r.Handle("/plan/reqdocument", Adapt(http.HandlerFunc(route.Handler.UpdatePlanReqDocument))).Methods("PUT")
	r.Handle("/plan/reqdocument/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePlanReqDocument))).Methods("PUT")
	r.Handle("/plan/reqdocument/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePlanReqDocument))).Methods("PUT")
	r.Handle("/plan/reqdocument/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePlanReqDocument))).Methods("DELETE")
	r.Handle("/plan/reqdocument/filter", Adapt(http.HandlerFunc(route.Handler.FilterPlanReqDocument))).Methods("POST")
}

//PlanDocumentRoutes : ""
func (route *Route) PlanDocumentRoutes(r *mux.Router) {
	r.Handle("/plan/document", Adapt(http.HandlerFunc(route.Handler.SavePlanDocument))).Methods("POST")
	r.Handle("/plan/document", Adapt(http.HandlerFunc(route.Handler.GetSinglePlanDocument))).Methods("GET")
	r.Handle("/plan/document", Adapt(http.HandlerFunc(route.Handler.UpdatePlanDocument))).Methods("PUT")
	r.Handle("/plan/document/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePlanDocument))).Methods("PUT")
	r.Handle("/plan/document/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePlanDocument))).Methods("PUT")
	r.Handle("/plan/document/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePlanDocument))).Methods("DELETE")
	r.Handle("/plan/document/filter", Adapt(http.HandlerFunc(route.Handler.FilterPlanDocument))).Methods("POST")
	r.Handle("/plan/document/pending", Adapt(http.HandlerFunc(route.Handler.GetPendingDocuments))).Methods("POST")
}
