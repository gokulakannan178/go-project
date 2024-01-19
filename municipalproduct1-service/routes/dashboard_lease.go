package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) LeaseRoutes(r *mux.Router) {
	// Lease
	r.Handle("/lease", Adapt(http.HandlerFunc(route.Handler.SaveLease))).Methods("POST")
	r.Handle("/lease", Adapt(http.HandlerFunc(route.Handler.GetSingleLease))).Methods("GET")
	r.Handle("/lease", Adapt(http.HandlerFunc(route.Handler.UpdateLease))).Methods("PUT")
	r.Handle("/lease/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLease))).Methods("PUT")
	r.Handle("/lease/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLease))).Methods("PUT")
	r.Handle("/lease/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLease))).Methods("DELETE")
	r.Handle("/lease/filter", Adapt(http.HandlerFunc(route.Handler.FilterLease))).Methods("POST")
}
