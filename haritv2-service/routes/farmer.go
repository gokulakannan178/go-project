package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) FarmerRoutes(r *mux.Router) {
	// Farmer
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.SaveFarmer))).Methods("POST")
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmer))).Methods("GET")
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.UpdateFarmer))).Methods("PUT")
	r.Handle("/farmer/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmer))).Methods("PUT")
	r.Handle("/farmer/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmer))).Methods("PUT")
	r.Handle("/farmer/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmer))).Methods("DELETE")
	r.Handle("/farmer/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmer))).Methods("POST")

	// Farmer Login / Generate OTP
	r.Handle("/user/auth/farmer/otplogin/generateotp", Adapt(http.HandlerFunc(route.Handler.FarmerLoginGenerateOTP))).Methods("POST")
	r.Handle("/user/auth/farmer/otplogin/validateotp", Adapt(http.HandlerFunc(route.Handler.FarmerLoginValidateOTP))).Methods("POST")
	r.Handle("/user/auth/farmerregistration/generateotp", Adapt(http.HandlerFunc(route.Handler.FarmerRegistrationLoginGenerateOTP))).Methods("POST")
	r.Handle("/user/auth/farmerregistration/validateotp", Adapt(http.HandlerFunc(route.Handler.FarmerRegistrationLoginValidateOTP))).Methods("POST")

}
