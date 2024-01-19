package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OnboardingCheckListRoutes : ""
func (route *Route) OnboardingCheckListRoutes(r *mux.Router) {
	r.Handle("/onboardingchecklist", Adapt(http.HandlerFunc(route.Handler.SaveOnboardingCheckList))).Methods("POST")
	r.Handle("/onboardingchecklist", Adapt(http.HandlerFunc(route.Handler.GetSingleOnboardingCheckList))).Methods("GET")
	r.Handle("/onboardingchecklist", Adapt(http.HandlerFunc(route.Handler.UpdateOnboardingCheckList))).Methods("PUT")
	r.Handle("/onboardingchecklist/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOnboardingCheckList))).Methods("PUT")
	r.Handle("/onboardingchecklist/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOnboardingCheckList))).Methods("PUT")
	r.Handle("/onboardingchecklist/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOnboardingCheckList))).Methods("DELETE")
	r.Handle("/onboardingchecklist/filter", Adapt(http.HandlerFunc(route.Handler.FilterOnboardingCheckList))).Methods("POST")

}
