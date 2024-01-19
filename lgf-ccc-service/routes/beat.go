package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// BeatRoutes : ""
func (route *Route) BeatRoutes(r *mux.Router) {
	r.Handle("/beat", Adapt(http.HandlerFunc(route.Handler.SaveBeat))).Methods("POST")
	r.Handle("/beat", Adapt(http.HandlerFunc(route.Handler.GetSingleBeat))).Methods("GET")
	r.Handle("/beat", Adapt(http.HandlerFunc(route.Handler.UpdateBeat))).Methods("PUT")
	r.Handle("/beat/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBeat))).Methods("PUT")
	r.Handle("/beat/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBeat))).Methods("PUT")
	r.Handle("/beat/filter", Adapt(http.HandlerFunc(route.Handler.FilterBeat))).Methods("POST")
	r.Handle("/beat/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBeat))).Methods("DELETE")
	r.Handle("/beat/end", Adapt(http.HandlerFunc(route.Handler.EndBeat))).Methods("PUT")
}
