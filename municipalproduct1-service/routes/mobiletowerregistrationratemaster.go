package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MobileTowerRegistrationRateMaster(r *mux.Router) {

	// MobileTowerRegistrationRateMaster
	r.Handle("/mobiletowerregistrationratemaster", Adapt(http.HandlerFunc(route.Handler.SaveMobileTowerRegistrationRateMaster))).Methods("POST")
	r.Handle("/mobiletowerregistrationratemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTowerRegistrationRateMaster))).Methods("GET")
	r.Handle("/mobiletowerregistrationratemaster", Adapt(http.HandlerFunc(route.Handler.UpdateMobileTowerRegistrationRateMaster))).Methods("PUT")
	r.Handle("/mobiletowerregistrationratemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMobileTowerRegistrationRateMaster))).Methods("PUT")
	r.Handle("/mobiletowerregistrationratemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMobileTowerRegistrationRateMaster))).Methods("PUT")
	r.Handle("/mobiletowerregistrationratemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMobileTowerRegistrationRateMaster))).Methods("DELETE")
	r.Handle("/mobiletowerregistrationratemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTowerRegistrationRateMaster))).Methods("GET")
}
