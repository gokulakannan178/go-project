package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// LeavePolicy : ""
func (route *Route) LeavePolicyRoutes(r *mux.Router) {
	r.Handle("/leavepolicy", Adapt(http.HandlerFunc(route.Handler.SaveLeavePolicy))).Methods("POST")
	r.Handle("/leavepolicy", Adapt(http.HandlerFunc(route.Handler.GetSingleLeavePolicy))).Methods("GET")
	r.Handle("/leavepolicy", Adapt(http.HandlerFunc(route.Handler.UpdateLeavePolicy))).Methods("PUT")
	r.Handle("/leavepolicy/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLeavePolicy))).Methods("PUT")
	r.Handle("/leavepolicy/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLeavePolicy))).Methods("PUT")
	r.Handle("/leavepolicy/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLeavePolicy))).Methods("DELETE")
	r.Handle("/leavepolicy/filter", Adapt(http.HandlerFunc(route.Handler.FilterLeavePolicy))).Methods("POST")

}
