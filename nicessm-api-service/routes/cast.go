package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CastRoutes(r *mux.Router) {
	r.Handle("/cast", Adapt(http.HandlerFunc(route.Handler.SaveCast))).Methods("POST")
	r.Handle("/cast", Adapt(http.HandlerFunc(route.Handler.GetSingleCast))).Methods("GET")
	r.Handle("/cast", Adapt(http.HandlerFunc(route.Handler.UpdateCast))).Methods("PUT")
	r.Handle("/cast/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCast))).Methods("PUT")
	r.Handle("/cast/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCast))).Methods("PUT")
	r.Handle("/cast/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCast))).Methods("DELETE")
	r.Handle("/cast/filter", Adapt(http.HandlerFunc(route.Handler.FilterCast))).Methods("POST")
}
