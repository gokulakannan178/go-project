package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeOffboardingCheckListRoutes : ""
func (route *Route) EmployeeOffboardingCheckListRoutes(r *mux.Router) {
	r.Handle("/employeeoffboardingchecklist", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeOffboardingCheckList))).Methods("POST")
	r.Handle("/employeeoffboardingchecklist", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeOffboardingCheckList))).Methods("GET")
	r.Handle("/employeeoffboardingchecklist", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeOffboardingCheckList))).Methods("PUT")
	r.Handle("/employeeoffboardingchecklist/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeOffboardingCheckList))).Methods("PUT")
	r.Handle("/employeeoffboardingchecklist/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeOffboardingCheckList))).Methods("PUT")
	r.Handle("/employeeoffboardingchecklist/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeOffboardingCheckList))).Methods("DELETE")
	r.Handle("/employeeoffboardingchecklist/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeOffboardingCheckList))).Methods("POST")
	r.Handle("/employeeoffboardingchecklist/final", Adapt(http.HandlerFunc(route.Handler.EmployeeOffboardingCheckListFinal))).Methods("GET")

}
