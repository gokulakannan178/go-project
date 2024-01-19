package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CronLogRoutes(r *mux.Router) {
	// CronLog
	r.Handle("/cronlog", Adapt(http.HandlerFunc(route.Handler.SaveCronLog))).Methods("POST")
	r.Handle("/cronlog", Adapt(http.HandlerFunc(route.Handler.GetSingleCronLog))).Methods("GET")
	r.Handle("/cronlog", Adapt(http.HandlerFunc(route.Handler.UpdateCronLog))).Methods("PUT")
	r.Handle("/cronlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCronLog))).Methods("PUT")
	r.Handle("/cronlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCronLog))).Methods("PUT")
	r.Handle("/cronlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCronLog))).Methods("DELETE")
	r.Handle("/cronlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterCronLog))).Methods("POST")
}
