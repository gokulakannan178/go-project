package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MobileTowerRegistrationTaxRoutes(r *mux.Router) {
	// MobileTowerRegistrationTax
	r.Handle("/mobiletowerregistrationtax", Adapt(http.HandlerFunc(route.Handler.SaveMobileTowerRegistrationTax))).Methods("POST")
	r.Handle("/mobiletowerregistrationtax", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTowerRegistrationTax))).Methods("GET")
	r.Handle("/mobiletowerregistrationtax/default", Adapt(http.HandlerFunc(route.Handler.GetSingleDefaultMobileTowerRegistrationTax))).Methods("GET")
	r.Handle("/mobiletowerregistrationtax", Adapt(http.HandlerFunc(route.Handler.UpdateMobileTowerRegistrationTax))).Methods("PUT")
	r.Handle("/mobiletowerregistrationtax/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMobileTowerRegistrationTax))).Methods("PUT")
	r.Handle("/mobiletowerregistrationtax/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMobileTowerRegistrationTax))).Methods("PUT")
	r.Handle("/mobiletowerregistrationtax/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMobileTowerRegistrationTax))).Methods("DELETE")
	r.Handle("/mobiletowerregistrationtax/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTowerRegistrationTax))).Methods("POST")
}
