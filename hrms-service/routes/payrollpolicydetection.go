package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PayrollPolicyDetectionRoutes : ""
func (route *Route) PayrollPolicyDetectionRoutes(r *mux.Router) {
	r.Handle("/payrollpolicydetection", Adapt(http.HandlerFunc(route.Handler.SavePayrollPolicyDetection))).Methods("POST")
	r.Handle("/payrollpolicydetection", Adapt(http.HandlerFunc(route.Handler.GetSinglePayrollPolicyDetection))).Methods("GET")
	r.Handle("/payrollpolicydetection", Adapt(http.HandlerFunc(route.Handler.UpdatePayrollPolicyDetection))).Methods("PUT")
	r.Handle("/payrollpolicydetection/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePayrollPolicyDetection))).Methods("PUT")
	r.Handle("/payrollpolicydetection/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePayrollPolicyDetection))).Methods("PUT")
	r.Handle("/payrollpolicydetection/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePayrollPolicyDetection))).Methods("DELETE")
	r.Handle("/payrollpolicydetection/filter", Adapt(http.HandlerFunc(route.Handler.FilterPayrollPolicyDetection))).Methods("POST")

}
