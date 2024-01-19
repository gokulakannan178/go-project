package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ULBRoutes : ""
func (route *Route) ULBRoutes(r *mux.Router) {
	r.Handle("/ulb", Adapt(http.HandlerFunc(route.Handler.SaveULB))).Methods("POST")
	r.Handle("/ulb", Adapt(http.HandlerFunc(route.Handler.GetSingleULB))).Methods("GET")
	r.Handle("/ulb", Adapt(http.HandlerFunc(route.Handler.UpdateULB))).Methods("PUT")
	r.Handle("/ulb/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableULB))).Methods("PUT")
	r.Handle("/ulb/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableULB))).Methods("PUT")
	r.Handle("/ulb/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteULB))).Methods("DELETE")
	r.Handle("/ulb/filter", Adapt(http.HandlerFunc(route.Handler.FilterULB))).Methods("POST")
}
