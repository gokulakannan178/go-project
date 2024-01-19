package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeEarningMasterRoutes : ""
func (route *Route) EmployeeEarningMasterRoutes(r *mux.Router) {
	r.Handle("/employeeearningmaster", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeEarningMaster))).Methods("POST")
	r.Handle("/employeeearningmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeEarningMaster))).Methods("GET")
	r.Handle("/employeeearningmaster", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeEarningMaster))).Methods("PUT")
	r.Handle("/employeeearningmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeEarningMaster))).Methods("PUT")
	r.Handle("/employeeearningmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeEarningMaster))).Methods("PUT")
	r.Handle("/employeeearningmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeEarningMaster))).Methods("DELETE")
	r.Handle("/employeeearningmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeEarningMaster))).Methods("POST")

}
