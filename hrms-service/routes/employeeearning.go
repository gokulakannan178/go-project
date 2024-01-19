package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeEarningRoutes : ""
func (route *Route) EmployeeEarningRoutes(r *mux.Router) {
	r.Handle("/employeeearning", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeEarning))).Methods("POST")
	r.Handle("/employeeearning", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeEarning))).Methods("GET")
	r.Handle("/employeeearning", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeEarning))).Methods("PUT")
	r.Handle("/employeeearning/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeEarning))).Methods("PUT")
	r.Handle("/employeeearning/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeEarning))).Methods("PUT")
	r.Handle("/employeeearning/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeEarning))).Methods("DELETE")
	r.Handle("/employeeearning/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeEarning))).Methods("POST")

}
