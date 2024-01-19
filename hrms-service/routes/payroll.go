package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PayrollRoutes(r *mux.Router) {
	// Payroll
	r.Handle("/payroll", Adapt(http.HandlerFunc(route.Handler.SavePayroll))).Methods("POST")
	r.Handle("/payroll", Adapt(http.HandlerFunc(route.Handler.GetSinglePayroll))).Methods("GET")
	r.Handle("/payroll", Adapt(http.HandlerFunc(route.Handler.UpdatePayroll))).Methods("PUT")
	r.Handle("/payroll/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePayroll))).Methods("PUT")
	r.Handle("/payroll/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePayroll))).Methods("PUT")
	r.Handle("/payroll/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePayroll))).Methods("DELETE")
	r.Handle("/payroll/filter", Adapt(http.HandlerFunc(route.Handler.FilterPayroll))).Methods("POST")
	r.Handle("/payroll/employee", Adapt(http.HandlerFunc(route.Handler.SavePayrollWithEmployee))).Methods("POST")
	r.Handle("/payroll/getemployee", Adapt(http.HandlerFunc(route.Handler.GetSinglePayrollWithEmployee))).Methods("GET")

}
