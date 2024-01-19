package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DayOfWeekRoutes(r *mux.Router) {
	// DayOfWeek
	r.Handle("/dayOfWeek", Adapt(http.HandlerFunc(route.Handler.SaveDayOfWeek))).Methods("POST")
	r.Handle("/dayOfWeek", Adapt(http.HandlerFunc(route.Handler.GetSingleDayOfWeek))).Methods("GET")
	r.Handle("/dayOfWeek", Adapt(http.HandlerFunc(route.Handler.UpdateDayOfWeek))).Methods("PUT")
	r.Handle("/dayOfWeek/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDayOfWeek))).Methods("PUT")
	r.Handle("/dayOfWeek/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDayOfWeek))).Methods("PUT")
	r.Handle("/dayOfWeek/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDayOfWeek))).Methods("DELETE")
	r.Handle("/dayOfWeek/filter", Adapt(http.HandlerFunc(route.Handler.FilterDayOfWeek))).Methods("POST")

}
