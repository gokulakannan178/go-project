package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TodayTipsRoutes(r *mux.Router) {
	// TodayTips
	r.Handle("/todaytips", Adapt(http.HandlerFunc(route.Handler.SaveTodayTips))).Methods("POST")
	r.Handle("/todaytips", Adapt(http.HandlerFunc(route.Handler.GetSingleTodayTips))).Methods("GET")
	r.Handle("/todaytips", Adapt(http.HandlerFunc(route.Handler.UpdateTodayTips))).Methods("PUT")
	r.Handle("/todaytips/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTodayTips))).Methods("PUT")
	r.Handle("/todaytips/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTodayTips))).Methods("PUT")
	r.Handle("/todaytips/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTodayTips))).Methods("DELETE")
	r.Handle("/todaytips/filter", Adapt(http.HandlerFunc(route.Handler.FilterTodayTips))).Methods("POST")
	r.Handle("/today/todaytips", Adapt(http.HandlerFunc(route.Handler.GetTodayTips))).Methods("GET")

}
