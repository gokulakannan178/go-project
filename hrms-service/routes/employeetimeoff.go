package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// EmployeeTimeOffRoutes : ""
func (route *Route) EmployeeTimeOffRoutes(r *mux.Router) {

	r.Handle("/employeetimeoff", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeTimeOff))).Methods("POST")
	r.Handle("/employeetimeoff", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeTimeOff))).Methods("GET")
	r.Handle("/employeetimeoff", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeTimeOff))).Methods("PUT")
	r.Handle("/employeetimeoff/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeTimeOff))).Methods("PUT")
	r.Handle("/employeetimeoff/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeTimeOff))).Methods("PUT")
	r.Handle("/employeetimeoff/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeTimeOff))).Methods("DELETE")
	r.Handle("/employeetimeoff/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeTimeOff))).Methods("POST")
	r.Handle("/employeetimeoff/request", Adapt(http.HandlerFunc(route.Handler.EmployeeTimeOffRequest))).Methods("POST")
	r.Handle("/employeetimeoff/approve", Adapt(http.HandlerFunc(route.Handler.EmployeeTimeOffApprove))).Methods("PUT")
	r.Handle("/employeetimeoff/revoke", Adapt(http.HandlerFunc(route.Handler.EmployeeTimeOffRevoke))).Methods("PUT")
	r.Handle("/employeetimeoff/revoke/request", Adapt(http.HandlerFunc(route.Handler.RevokeRequestEmployeeTimeOff))).Methods("PUT")
	r.Handle("/employeetimeoff/reject", Adapt(http.HandlerFunc(route.Handler.EmployeeTimeOffReject))).Methods("PUT")
	r.Handle("/employeetimeoff/count", Adapt(http.HandlerFunc(route.Handler.EmployeeTimeOffCount))).Methods("POST")
	r.Handle("/employeetimeoff/datecheck", Adapt(http.HandlerFunc(route.Handler.EmployeeTimeoffDateCheck))).Methods("POST")

}
