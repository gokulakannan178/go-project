package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OnboardingPolicyRoutes : ""
func (route *Route) OnboardingPolicyRoutes(r *mux.Router) {
	r.Handle("/onboardingpolicy", Adapt(http.HandlerFunc(route.Handler.SaveOnboardingPolicy))).Methods("POST")
	r.Handle("/onboardingpolicy", Adapt(http.HandlerFunc(route.Handler.GetSingleOnboardingPolicy))).Methods("GET")
	r.Handle("/onboardingpolicy", Adapt(http.HandlerFunc(route.Handler.UpdateOnboardingPolicy))).Methods("PUT")
	r.Handle("/onboardingpolicy/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOnboardingPolicy))).Methods("PUT")
	r.Handle("/onboardingpolicy/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOnboardingPolicy))).Methods("PUT")
	r.Handle("/onboardingpolicy/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOnboardingPolicy))).Methods("DELETE")
	r.Handle("/onboardingpolicy/filter", Adapt(http.HandlerFunc(route.Handler.FilterOnboardingPolicy))).Methods("POST")

}
