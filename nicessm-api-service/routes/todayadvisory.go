package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TodayAdvisoryRoutes(r *mux.Router) {
	r.Handle("/todayadvisory", Adapt(http.HandlerFunc(route.Handler.SaveTodayAdvisory))).Methods("POST")
	r.Handle("/todayadvisory", Adapt(http.HandlerFunc(route.Handler.GetSingleTodayAdvisory))).Methods("GET")
	r.Handle("/todayadvisory", Adapt(http.HandlerFunc(route.Handler.UpdateTodayAdvisory))).Methods("PUT")
	r.Handle("/todayadvisory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTodayAdvisory))).Methods("PUT")
	r.Handle("/todayadvisory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTodayAdvisory))).Methods("PUT")
	r.Handle("/todayadvisory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTodayAdvisory))).Methods("DELETE")
	r.Handle("/todayadvisory/filter", Adapt(http.HandlerFunc(route.Handler.FilterTodayAdvisory))).Methods("POST")
	r.Handle("/gettodayadvisorys", Adapt(http.HandlerFunc(route.Handler.GetTodayAdvisory))).Methods("GET")

}
