package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OffboardingCheckListRoutes : ""
func (route *Route) OffboardingCheckListRoutes(r *mux.Router) {
	r.Handle("/offboardingchecklist", Adapt(http.HandlerFunc(route.Handler.SaveOffboardingCheckList))).Methods("POST")
	r.Handle("/offboardingchecklist", Adapt(http.HandlerFunc(route.Handler.GetSingleOffboardingCheckList))).Methods("GET")
	r.Handle("/offboardingchecklist", Adapt(http.HandlerFunc(route.Handler.UpdateOffboardingCheckList))).Methods("PUT")
	r.Handle("/offboardingchecklist/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOffboardingCheckList))).Methods("PUT")
	r.Handle("/offboardingchecklist/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOffboardingCheckList))).Methods("PUT")
	r.Handle("/offboardingchecklist/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOffboardingCheckList))).Methods("DELETE")
	r.Handle("/offboardingchecklist/filter", Adapt(http.HandlerFunc(route.Handler.FilterOffboardingCheckList))).Methods("POST")

}
