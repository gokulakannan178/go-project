package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) RouteMasterRoutes(r *mux.Router) {
	// RouteMaster
	r.Handle("/routemaster", Adapt(http.HandlerFunc(route.Handler.SaveRouteMaster))).Methods("POST")
	r.Handle("/routemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleRouteMaster))).Methods("GET")
	r.Handle("/routemaster", Adapt(http.HandlerFunc(route.Handler.UpdateRouteMaster))).Methods("PUT")
	r.Handle("/routemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRouteMaster))).Methods("PUT")
	r.Handle("/routemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRouteMaster))).Methods("PUT")
	r.Handle("/routemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRouteMaster))).Methods("DELETE")
	r.Handle("/routemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterRouteMaster))).Methods("POST")

}
