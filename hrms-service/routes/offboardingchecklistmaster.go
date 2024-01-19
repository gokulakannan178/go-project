package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OffboardingCheckListMasterRoutes : ""
func (route *Route) OffboardingCheckListMasterRoutes(r *mux.Router) {
	r.Handle("/offboardingchecklistmaster", Adapt(http.HandlerFunc(route.Handler.SaveOffboardingCheckListMaster))).Methods("POST")
	r.Handle("/offboardingchecklistmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleOffboardingCheckListMaster))).Methods("GET")
	r.Handle("/offboardingchecklistmaster", Adapt(http.HandlerFunc(route.Handler.UpdateOffboardingCheckListMaster))).Methods("PUT")
	r.Handle("/offboardingchecklistmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOffboardingCheckListMaster))).Methods("PUT")
	r.Handle("/offboardingchecklistmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOffboardingCheckListMaster))).Methods("PUT")
	r.Handle("/offboardingchecklistmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOffboardingCheckListMaster))).Methods("DELETE")
	r.Handle("/offboardingchecklistmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterOffboardingCheckListMaster))).Methods("POST")

}
