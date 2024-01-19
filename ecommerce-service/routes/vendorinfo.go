package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) VendorInfoRoutes(r *mux.Router) {
	// VendorInfo
	r.Handle("/vendorinfo", Adapt(http.HandlerFunc(route.Handler.SaveVendorInfo))).Methods("POST")
	r.Handle("/vendorinfo", Adapt(http.HandlerFunc(route.Handler.GetSingleVendorInfo))).Methods("GET")
	r.Handle("/vendorinfo", Adapt(http.HandlerFunc(route.Handler.UpdateVendorInfo))).Methods("PUT")
	r.Handle("/vendorinfo/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVendorInfo))).Methods("PUT")
	r.Handle("/vendorinfo/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVendorInfo))).Methods("PUT")
	r.Handle("/vendorinfo/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVendorInfo))).Methods("DELETE")
	r.Handle("/vendorinfo/filter", Adapt(http.HandlerFunc(route.Handler.FilterVendorInfo))).Methods("POST")
}
