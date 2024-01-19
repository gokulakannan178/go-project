package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ThingsToKnowRoutes : ""
func (route *Route) ThingsToKnowRoutes(r *mux.Router) {
	r.Handle("/thingstoknow", Adapt(http.HandlerFunc(route.Handler.SaveThingsToKnow))).Methods("POST")
	r.Handle("/thingstoknow", Adapt(http.HandlerFunc(route.Handler.GetSingleThingsToKnow))).Methods("GET")
	r.Handle("/thingstoknow", Adapt(http.HandlerFunc(route.Handler.UpdateThingsToKnow))).Methods("PUT")
	r.Handle("/thingstoknow/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableThingsToKnow))).Methods("PUT")
	r.Handle("/thingstoknow/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableThingsToKnow))).Methods("PUT")
	r.Handle("/thingstoknow/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteThingsToKnow))).Methods("DELETE")
	r.Handle("/thingstoknow/filter", Adapt(http.HandlerFunc(route.Handler.FilterThingsToKnow))).Methods("POST")
}
