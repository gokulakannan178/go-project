package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EmployeePayslipRoutes(r *mux.Router) {
	// EmployeePayslip
	r.Handle("/EmployeePayslip", Adapt(http.HandlerFunc(route.Handler.SaveEmployeePayslip))).Methods("POST")
	r.Handle("/EmployeePayslip", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeePayslip))).Methods("GET")
	r.Handle("/EmployeePayslip", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeePayslip))).Methods("PUT")
	r.Handle("/EmployeePayslip/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeePayslip))).Methods("PUT")
	r.Handle("/EmployeePayslip/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeePayslip))).Methods("PUT")
	r.Handle("/EmployeePayslip/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeePayslip))).Methods("DELETE")
	r.Handle("/EmployeePayslip/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeePayslip))).Methods("POST")

}
