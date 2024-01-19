package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OffboardingPolicyRoutes : ""
func (route *Route) OffboardingPolicyRoutes(r *mux.Router) {
	r.Handle("/offboardingpolicy", Adapt(http.HandlerFunc(route.Handler.SaveOffboardingPolicy))).Methods("POST")
	r.Handle("/offboardingpolicy", Adapt(http.HandlerFunc(route.Handler.GetSingleOffboardingPolicy))).Methods("GET")
	r.Handle("/offboardingpolicy", Adapt(http.HandlerFunc(route.Handler.UpdateOffboardingPolicy))).Methods("PUT")
	r.Handle("/offboardingpolicy/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOffboardingPolicy))).Methods("PUT")
	r.Handle("/offboardingpolicy/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOffboardingPolicy))).Methods("PUT")
	r.Handle("/offboardingpolicy/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOffboardingPolicy))).Methods("DELETE")
	r.Handle("/offboardingpolicy/filter", Adapt(http.HandlerFunc(route.Handler.FilterOffboardingPolicy))).Methods("POST")

}
