package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EmployeeHistoryRoutes(r *mux.Router) {
	// EmployeeHistory
	r.Handle("/employeeHistory", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeHistory))).Methods("POST")
	r.Handle("/employeeHistory", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeHistory))).Methods("GET")
	r.Handle("/employeeHistory", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeHistory))).Methods("PUT")
	r.Handle("/employeeHistory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeHistory))).Methods("PUT")
	r.Handle("/employeeHistory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeHistory))).Methods("PUT")
	r.Handle("/employeeHistory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeHistory))).Methods("DELETE")
	r.Handle("/employeeHistory/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeHistory))).Methods("POST")

}
