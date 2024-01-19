package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PmTargetRoutes(r *mux.Router) {
	// PmTarget
	r.Handle("/pmtarget", Adapt(http.HandlerFunc(route.Handler.SavePmTarget))).Methods("POST")
	r.Handle("/pmtarget", Adapt(http.HandlerFunc(route.Handler.GetSinglePmTarget))).Methods("GET")
	r.Handle("/pmtarget", Adapt(http.HandlerFunc(route.Handler.UpdatePmTarget))).Methods("PUT")
	r.Handle("/pmtarget/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePmTarget))).Methods("PUT")
	r.Handle("/pmtarget/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePmTarget))).Methods("PUT")
	r.Handle("/pmtarget/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePmTarget))).Methods("DELETE")
	r.Handle("/pmtarget/filter", Adapt(http.HandlerFunc(route.Handler.FilterPmTarget))).Methods("POST")
}
