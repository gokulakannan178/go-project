package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeDeductionRoutes : ""
func (route *Route) EmployeeDeductionRoutes(r *mux.Router) {
	r.Handle("/employeededuction", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeDeduction))).Methods("POST")
	r.Handle("/employeededuction", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeDeduction))).Methods("GET")
	r.Handle("/employeededuction", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeDeduction))).Methods("PUT")
	r.Handle("/employeededuction/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeDeduction))).Methods("PUT")
	r.Handle("/employeededuction/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeDeduction))).Methods("PUT")
	r.Handle("/employeededuction/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeDeduction))).Methods("DELETE")
	r.Handle("/employeededuction/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeDeduction))).Methods("POST")

}
