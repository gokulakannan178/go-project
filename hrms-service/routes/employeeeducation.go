package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeEducationRoutes : ""
func (route *Route) EmployeeEducationRoutes(r *mux.Router) {
	r.Handle("/employeeeducation", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeEducation))).Methods("POST")
	r.Handle("/employeeeducation", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeEducation))).Methods("GET")
	r.Handle("/employeeeducation", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeEducation))).Methods("PUT")
	r.Handle("/employeeeducation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeEducation))).Methods("PUT")
	r.Handle("/employeeeducation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeEducation))).Methods("PUT")
	r.Handle("/employeeeducation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeEducation))).Methods("DELETE")
	r.Handle("/employeeeducation/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeEducation))).Methods("POST")

}
