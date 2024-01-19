package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) VendorRoutes(r *mux.Router) {
	// Vendor
	r.Handle("/vendor", Adapt(http.HandlerFunc(route.Handler.SaveVendor))).Methods("POST")
	r.Handle("/vendor", Adapt(http.HandlerFunc(route.Handler.GetSingleVendor))).Methods("GET")
	r.Handle("/vendor", Adapt(http.HandlerFunc(route.Handler.UpdateVendor))).Methods("PUT")
	r.Handle("/vendor/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVendor))).Methods("PUT")
	r.Handle("/vendor/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVendor))).Methods("PUT")
	r.Handle("/vendor/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVendor))).Methods("DELETE")
	r.Handle("/vendor/filter", Adapt(http.HandlerFunc(route.Handler.FilterVendor))).Methods("POST")
	r.Handle("/vendor/mobileno", Adapt(http.HandlerFunc(route.Handler.GetSingleVendorWithMobileNoV2))).Methods("GET")

	// VendorAuthRoutes
	r.Handle("/vendor/auth/otplogin/generateotp", Adapt(http.HandlerFunc(route.Handler.VendorOTPLoginGenerateOTP))).Methods("POST")
	r.Handle("/vendor/auth/otplogin/validateotp", Adapt(http.HandlerFunc(route.Handler.VendorOTPLoginValidateOTP))).Methods("POST")
}
