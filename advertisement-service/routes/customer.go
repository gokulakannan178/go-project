package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CustomerRoutes(r *mux.Router) {
	// Vendor
	r.Handle("/customer", Adapt(http.HandlerFunc(route.Handler.SaveCustomer))).Methods("POST")
	r.Handle("/customer", Adapt(http.HandlerFunc(route.Handler.GetSingleCustomer))).Methods("GET")
	r.Handle("/customer/getusingmobilenumber", Adapt(http.HandlerFunc(route.Handler.GetSingleGetUsingMobileNumber))).Methods("GET")
	r.Handle("/customer", Adapt(http.HandlerFunc(route.Handler.UpdateCustomer))).Methods("PUT")
	r.Handle("/customer/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCustomer))).Methods("PUT")
	r.Handle("/customer/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCustomer))).Methods("PUT")
	r.Handle("/customer/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCustomer))).Methods("DELETE")
	r.Handle("/customer/filter", Adapt(http.HandlerFunc(route.Handler.FilterCustomer))).Methods("POST")
}

//CustomerAuthRoutes : ""
func (route *Route) CustomerAuthRoutes(r *mux.Router) {
	r.Handle("/customer/auth", Adapt(http.HandlerFunc(route.Handler.CustomerLogin))).Methods("POST")
	r.Handle("/customer/auth/otplogin/generateotp", Adapt(http.HandlerFunc(route.Handler.CustomerOTPLoginGenerateOTP))).Methods("POST")
	r.Handle("/customer/auth/otplogin/validateotp", Adapt(http.HandlerFunc(route.Handler.CustomerOTPLoginValidateOTP))).Methods("POST")
}
