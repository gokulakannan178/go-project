package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PayrollPolicyRoutes : ""
func (route *Route) PayrollPolicyRoutes(r *mux.Router) {
	r.Handle("/payrollpolicy", Adapt(http.HandlerFunc(route.Handler.SavePayrollPolicy))).Methods("POST")
	r.Handle("/payrollpolicy", Adapt(http.HandlerFunc(route.Handler.GetSinglePayrollPolicy))).Methods("GET")
	r.Handle("/payrollpolicy", Adapt(http.HandlerFunc(route.Handler.UpdatePayrollPolicy))).Methods("PUT")
	r.Handle("/payrollpolicy/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePayrollPolicy))).Methods("PUT")
	r.Handle("/payrollpolicy/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePayrollPolicy))).Methods("PUT")
	r.Handle("/payrollpolicy/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePayrollPolicy))).Methods("DELETE")
	r.Handle("/payrollpolicy/filter", Adapt(http.HandlerFunc(route.Handler.FilterPayrollPolicy))).Methods("POST")
	r.Handle("/salarycalc", Adapt(http.HandlerFunc(route.Handler.GetSalaryCalc))).Methods("GET")
	r.Handle("/salarycalc/employeeType", Adapt(http.HandlerFunc(route.Handler.GetSalaryCalcWithEmployeeType))).Methods("GET")

}
