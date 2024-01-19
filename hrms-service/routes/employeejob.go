package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EmployeeJobRoutes(r *mux.Router) {
	// EmployeeJob
	r.Handle("/employeeJob", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeJob))).Methods("POST")
	r.Handle("/employeeJob", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeJob))).Methods("GET")
	r.Handle("/employeeJob", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeJob))).Methods("PUT")
	r.Handle("/employeeJob/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeJob))).Methods("PUT")
	r.Handle("/employeeJob/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeJob))).Methods("PUT")
	r.Handle("/employeeJob/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeJob))).Methods("DELETE")
	r.Handle("/employeeJob/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeJob))).Methods("POST")

}
