package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MonthSeasonRoutes(r *mux.Router) {
	r.Handle("/monthseason", Adapt(http.HandlerFunc(route.Handler.SaveMonthSeason))).Methods("POST")
	r.Handle("/monthseason", Adapt(http.HandlerFunc(route.Handler.GetSingleMonthSeason))).Methods("GET")
	r.Handle("/monthseason", Adapt(http.HandlerFunc(route.Handler.UpdateMonthSeason))).Methods("PUT")
	r.Handle("/monthseason/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMonthSeason))).Methods("PUT")
	r.Handle("/monthseason/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMonthSeason))).Methods("PUT")
	r.Handle("/monthseason/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMonthSeason))).Methods("DELETE")
	r.Handle("/monthseason/filter", Adapt(http.HandlerFunc(route.Handler.FilterMonthSeason))).Methods("POST")
}
