package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeExperienceRoutes : ""
func (route *Route) EmployeeExperienceRoutes(r *mux.Router) {
	r.Handle("/employeeexperience", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeExperience))).Methods("POST")
	r.Handle("/employeeexperience", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeExperience))).Methods("GET")
	r.Handle("/employeeexperience", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeExperience))).Methods("PUT")
	r.Handle("/employeeexperience/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeExperience))).Methods("PUT")
	r.Handle("/employeeexperience/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeExperience))).Methods("PUT")
	r.Handle("/employeeexperience/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeExperience))).Methods("DELETE")
	r.Handle("/employeeexperience/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeExperience))).Methods("POST")

}
