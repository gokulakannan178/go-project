package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CustomerRoutes(r *mux.Router) {
	// Customer
	r.Handle("/customer", Adapt(http.HandlerFunc(route.Handler.SaveCustomer))).Methods("POST")
	r.Handle("/customer", Adapt(http.HandlerFunc(route.Handler.GetSingleCustomer))).Methods("GET")
	r.Handle("/customer", Adapt(http.HandlerFunc(route.Handler.UpdateCustomer))).Methods("PUT")
	r.Handle("/customer/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCustomer))).Methods("PUT")
	r.Handle("/customer/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCustomer))).Methods("PUT")
	r.Handle("/customer/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCustomer))).Methods("DELETE")
	r.Handle("/customer/filter", Adapt(http.HandlerFunc(route.Handler.FilterCustomer))).Methods("POST")
	r.Handle("/customer/otplogin/generateotp", Adapt(http.HandlerFunc(route.Handler.CustomerLoginGenerateOTP))).Methods("POST")
	r.Handle("/customer/otplogin/validateotp", Adapt(http.HandlerFunc(route.Handler.CustomerLoginValidateOTP))).Methods("POST")
	r.Handle("/customer/profile", Adapt(http.HandlerFunc(route.Handler.GetSingleCustomerwithprofile))).Methods("GET")
	r.Handle("/customer/updateprofile", Adapt(http.HandlerFunc(route.Handler.UpdateCustomerwithprofile))).Methods("PUT")

}
func (route *Route) CustomerRegistrationRoutes(r *mux.Router) {
	r.Handle("/customer/registration/generateotp", Adapt(http.HandlerFunc(route.Handler.CustomerregistrationGenerateOTP))).Methods("POST")
	r.Handle("/customer/registration/validateotp", Adapt(http.HandlerFunc(route.Handler.CustomerregistrationValidateOTP))).Methods("POST")
}
