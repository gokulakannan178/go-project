package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OverallPropertyDemandRoutes(r *mux.Router) {
	// OverallPropertyDemand
	r.Handle("/overallpropertydemand", Adapt(http.HandlerFunc(route.Handler.SaveOverallPropertyDemand))).Methods("POST")
	r.Handle("/overallpropertydemand", Adapt(http.HandlerFunc(route.Handler.GetSingleOverallPropertyDemand))).Methods("GET")
	r.Handle("/overallpropertydemand", Adapt(http.HandlerFunc(route.Handler.UpdateOverallPropertyDemand))).Methods("PUT")
	r.Handle("/overallpropertydemand/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOverallPropertyDemand))).Methods("PUT")
	r.Handle("/overallpropertydemand/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOverallPropertyDemand))).Methods("PUT")
	r.Handle("/overallpropertydemand/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOverallPropertyDemand))).Methods("DELETE")
	r.Handle("/overallpropertydemand/filter", Adapt(http.HandlerFunc(route.Handler.FilterOverallPropertyDemand))).Methods("POST")
}
