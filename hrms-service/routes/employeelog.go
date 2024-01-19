package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeLogRoutes : ""
func (route *Route) EmployeeLogRoutes(r *mux.Router) {
	r.Handle("/employeelog", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeLog))).Methods("POST")
	r.Handle("/employeelog", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeLog))).Methods("GET")
	r.Handle("/employeelog", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeLog))).Methods("PUT")
	r.Handle("/employeelog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeLog))).Methods("PUT")
	r.Handle("/employeelog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeLog))).Methods("PUT")
	r.Handle("/employeelog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeLog))).Methods("DELETE")
	r.Handle("/employeelog/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeLog))).Methods("POST")

}
