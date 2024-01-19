package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeOnboardingCheckListRoutes : ""
func (route *Route) EmployeeOnboardingCheckListRoutes(r *mux.Router) {
	r.Handle("/employeeonboardingchecklist", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeOnboardingCheckList))).Methods("POST")
	r.Handle("/employeeonboardingchecklist", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeOnboardingCheckList))).Methods("GET")
	r.Handle("/employeeonboardingchecklist", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeOnboardingCheckList))).Methods("PUT")
	r.Handle("/employeeonboardingchecklist/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeOnboardingCheckList))).Methods("PUT")
	r.Handle("/employeeonboardingchecklist/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeOnboardingCheckList))).Methods("PUT")
	r.Handle("/employeeonboardingchecklist/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeOnboardingCheckList))).Methods("DELETE")
	r.Handle("/employeeonboardingchecklist/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeOnboardingCheckList))).Methods("POST")
	r.Handle("/employeeonboardingchecklist/final", Adapt(http.HandlerFunc(route.Handler.EmployeeOnboardingCheckListFinal))).Methods("GET")

}
