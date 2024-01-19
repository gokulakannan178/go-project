package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeShiftRoutes : ""
func (route *Route) EmployeeShiftRoutes(r *mux.Router) {
	r.Handle("/employeeshift", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeShift))).Methods("POST")
	r.Handle("/employeeshift", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeShift))).Methods("GET")
	r.Handle("/employeeshift", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeShift))).Methods("PUT")
	r.Handle("/employeeshift/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeShift))).Methods("PUT")
	r.Handle("/employeeshift/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeShift))).Methods("PUT")
	r.Handle("/employeeshift/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeShift))).Methods("DELETE")
	r.Handle("/employeeshift/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeShift))).Methods("POST")

}
