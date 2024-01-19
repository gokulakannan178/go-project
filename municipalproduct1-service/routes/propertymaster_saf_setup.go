package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//AVRRangeRoutes : ""
func (route *Route) AVRRangeRoutes(r *mux.Router) {
	r.Handle("/avrRange", Adapt(http.HandlerFunc(route.Handler.SaveAVRRange))).Methods("POST")
	r.Handle("/avrRange", Adapt(http.HandlerFunc(route.Handler.GetSingleAVRRange))).Methods("GET")
	r.Handle("/avrRange", Adapt(http.HandlerFunc(route.Handler.UpdateAVRRange))).Methods("PUT")
	r.Handle("/avrRange/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAVRRange))).Methods("PUT")
	r.Handle("/avrRange/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAVRRange))).Methods("PUT")
	r.Handle("/avrRange/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAVRRange))).Methods("DELETE")
	r.Handle("/avrRange/filter", Adapt(http.HandlerFunc(route.Handler.FilterAVRRange))).Methods("POST")
}

//PropertyOtherTaxRoutes : ""
func (route *Route) PropertyOtherTaxRoutes(r *mux.Router) {
	r.Handle("/propertyOtherTax", Adapt(http.HandlerFunc(route.Handler.SavePropertyOtherTax))).Methods("POST")
	r.Handle("/propertyOtherTax", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyOtherTax))).Methods("GET")
	r.Handle("/propertyOtherTax", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyOtherTax))).Methods("PUT")
	r.Handle("/propertyOtherTax/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyOtherTax))).Methods("PUT")
	r.Handle("/propertyOtherTax/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyOtherTax))).Methods("PUT")
	r.Handle("/propertyOtherTax/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyOtherTax))).Methods("DELETE")
	r.Handle("/propertyOtherTax/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyOtherTax))).Methods("POST")
}
