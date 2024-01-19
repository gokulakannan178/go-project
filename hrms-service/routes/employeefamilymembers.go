package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeFamilyMembersRoutes : ""
func (route *Route) EmployeeFamilyMembersRoutes(r *mux.Router) {
	r.Handle("/employeefamilymembers", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeFamilyMembers))).Methods("POST")
	r.Handle("/employeefamilymembers", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeFamilyMembers))).Methods("GET")
	r.Handle("/employeefamilymembers", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeFamilyMembers))).Methods("PUT")
	r.Handle("/employeefamilymembers/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeFamilyMembers))).Methods("PUT")
	r.Handle("/employeefamilymembers/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeFamilyMembers))).Methods("PUT")
	r.Handle("/employeefamilymembers/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeFamilyMembers))).Methods("DELETE")
	r.Handle("/employeefamilymembers/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeFamilyMembers))).Methods("POST")

}
