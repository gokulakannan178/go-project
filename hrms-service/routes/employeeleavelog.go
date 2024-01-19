package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeLeaveLogRoutes : ""
func (route *Route) EmployeeLeaveLogRoutes(r *mux.Router) {

	r.Handle("/employeeleavelog", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeLeaveLog))).Methods("POST")
	r.Handle("/employeeleavelog", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeLeaveLog))).Methods("GET")
	r.Handle("/employeeleavelog", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeLeaveLog))).Methods("PUT")
	r.Handle("/employeeleavelog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeLeaveLog))).Methods("PUT")
	r.Handle("/employeeleavelog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeLeaveLog))).Methods("PUT")
	r.Handle("/employeeleavelog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeLeaveLog))).Methods("DELETE")
	r.Handle("/employeeleavelog/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeLeaveLog))).Methods("POST")
	r.Handle("/employeeleavelog/count", Adapt(http.HandlerFunc(route.Handler.EmployeeLeaveLogCount))).Methods("POST")

}
