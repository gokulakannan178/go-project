package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SeasonRoutes(r *mux.Router) {
	r.Handle("/season", Adapt(http.HandlerFunc(route.Handler.SaveSeason))).Methods("POST")
	r.Handle("/season", Adapt(http.HandlerFunc(route.Handler.GetSingleSeason))).Methods("GET")
	r.Handle("/season", Adapt(http.HandlerFunc(route.Handler.UpdateSeason))).Methods("PUT")
	r.Handle("/season/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSeason))).Methods("PUT")
	r.Handle("/season/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSeason))).Methods("PUT")
	r.Handle("/season/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSeason))).Methods("DELETE")
	r.Handle("/season/filter", Adapt(http.HandlerFunc(route.Handler.FilterSeason))).Methods("POST")
}
