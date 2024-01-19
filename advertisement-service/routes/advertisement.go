package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) AdvertisementRoutes(r *mux.Router) {
	// Vendor
	r.Handle("/advertisement", Adapt(http.HandlerFunc(route.Handler.SaveAdvertisement))).Methods("POST")
	r.Handle("/advertisement", Adapt(http.HandlerFunc(route.Handler.GetSingleAdvertisement))).Methods("GET")
	r.Handle("/advertisement", Adapt(http.HandlerFunc(route.Handler.UpdateAdvertisement))).Methods("PUT")
	r.Handle("/advertisement/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAdvertisement))).Methods("PUT")
	r.Handle("/advertisement/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAdvertisement))).Methods("PUT")
	r.Handle("/advertisement/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAdvertisement))).Methods("DELETE")
	r.Handle("/advertisement/filter", Adapt(http.HandlerFunc(route.Handler.FilterAdvertisement))).Methods("POST")
}
