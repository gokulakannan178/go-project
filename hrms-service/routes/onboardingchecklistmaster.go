package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OnboardingCheckListMasterRoutes : ""
func (route *Route) OnboardingCheckListMasterRoutes(r *mux.Router) {
	r.Handle("/onboardingchecklistmaster", Adapt(http.HandlerFunc(route.Handler.SaveOnboardingCheckListMaster))).Methods("POST")
	r.Handle("/onboardingchecklistmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleOnboardingCheckListMaster))).Methods("GET")
	r.Handle("/onboardingchecklistmaster", Adapt(http.HandlerFunc(route.Handler.UpdateOnboardingCheckListMaster))).Methods("PUT")
	r.Handle("/onboardingchecklistmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOnboardingCheckListMaster))).Methods("PUT")
	r.Handle("/onboardingchecklistmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOnboardingCheckListMaster))).Methods("PUT")
	r.Handle("/onboardingchecklistmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOnboardingCheckListMaster))).Methods("DELETE")
	r.Handle("/onboardingchecklistmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterOnboardingCheckListMaster))).Methods("POST")

}
