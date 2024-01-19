package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//DesignationRoutes : ""
func (route *Route) DesignationRoutes(r *mux.Router) {
	r.Handle("/designation", Adapt(http.HandlerFunc(route.Handler.SaveDesignation))).Methods("POST")
	r.Handle("/designation", Adapt(http.HandlerFunc(route.Handler.GetSingleDesignation))).Methods("GET")
	r.Handle("/designation", Adapt(http.HandlerFunc(route.Handler.UpdateDesignation))).Methods("PUT")
	r.Handle("/designation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDesignation))).Methods("PUT")
	r.Handle("/designation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDesignation))).Methods("PUT")
	r.Handle("/designation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDesignation))).Methods("DELETE")
	r.Handle("/designation/filter", Adapt(http.HandlerFunc(route.Handler.FilterDesignation))).Methods("POST")

}
