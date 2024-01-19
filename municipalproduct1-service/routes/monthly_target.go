package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MonthlyTargetRoutes(r *mux.Router) {
	// Monthly Target
	r.Handle("/monthlytarget", Adapt(http.HandlerFunc(route.Handler.SaveMonthlyTarget))).Methods("POST")
	r.Handle("/monthlytarget", Adapt(http.HandlerFunc(route.Handler.GetSingleMonthlyTarget))).Methods("GET")
	r.Handle("/monthlytarget", Adapt(http.HandlerFunc(route.Handler.UpdateMonthlyTarget))).Methods("PUT")
	r.Handle("/monthlytarget/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMonthlyTarget))).Methods("PUT")
	r.Handle("/monthlytarget/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMonthlyTarget))).Methods("PUT")
	r.Handle("/monthlytarget/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMonthlyTarget))).Methods("DELETE")
	r.Handle("/monthlytarget/filter", Adapt(http.HandlerFunc(route.Handler.FilterMonthlyTarget))).Methods("POST")
}
