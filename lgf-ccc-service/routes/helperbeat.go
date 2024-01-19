package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) HelperBeatRoutes(r *mux.Router) {
	r.Handle("/helperbeat/assign", Adapt(http.HandlerFunc(route.Handler.SaveHelperBeat))).Methods("POST")
	r.Handle("/helperbeat", Adapt(http.HandlerFunc(route.Handler.GetSingleHelperBeat))).Methods("GET")
	r.Handle("/helperbeat", Adapt(http.HandlerFunc(route.Handler.UpdateHelperBeat))).Methods("PUT")
	r.Handle("/helperbeat/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableHelperBeat))).Methods("PUT")
	r.Handle("/helperbeat/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableHelperBeat))).Methods("PUT")
	r.Handle("/helperbeat/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteHelperBeat))).Methods("DELETE")
	r.Handle("/helperbeat/filter", Adapt(http.HandlerFunc(route.Handler.FilterHelperBeat))).Methods("POST")
}
