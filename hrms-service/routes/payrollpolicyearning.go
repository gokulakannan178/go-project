package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PayrollPolicyEarningRoutes : ""
func (route *Route) PayrollPolicyEarningRoutes(r *mux.Router) {
	r.Handle("/payrollpolicyearning", Adapt(http.HandlerFunc(route.Handler.SavePayrollPolicyEarning))).Methods("POST")
	r.Handle("/payrollpolicyearning", Adapt(http.HandlerFunc(route.Handler.GetSinglePayrollPolicyEarning))).Methods("GET")
	r.Handle("/payrollpolicyearning", Adapt(http.HandlerFunc(route.Handler.UpdatePayrollPolicyEarning))).Methods("PUT")
	r.Handle("/payrollpolicyearning/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePayrollPolicyEarning))).Methods("PUT")
	r.Handle("/payrollpolicyearning/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePayrollPolicyEarning))).Methods("PUT")
	r.Handle("/payrollpolicyearning/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePayrollPolicyEarning))).Methods("DELETE")
	r.Handle("/payrollpolicyearning/filter", Adapt(http.HandlerFunc(route.Handler.FilterPayrollPolicyEarning))).Methods("POST")

}
